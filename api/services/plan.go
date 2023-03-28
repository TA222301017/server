package services

import (
	"errors"
	"fmt"
	"os"
	"server/api/template"
	"server/models"
	"server/setup"
	"strings"
)

func GetPlans(p template.SearchParameter, keyword string) ([]models.Plan, *template.Pagination, error) {
	db := setup.DB

	offset := (p.Page - 1) * p.Limit
	limit := p.Limit
	keyword = "%" + keyword + "%"

	var cnt int64
	if err := db.Where("name LIKE ?", keyword).
		Find(&models.Plan{}).Count(&cnt).Error; err != nil {
		return nil, nil, err
	}

	var plans []models.Plan
	if err := db.
		Where("name LIKE ?", keyword).Order("created_at DESC").
		Offset(offset).Limit(limit).Find(&plans).Error; err != nil {
		return nil, nil, err
	}

	last := cnt / int64(limit)
	pagination := template.Pagination{
		Page:  p.Page,
		Limit: limit,
		Total: int(cnt),
		Last:  int(last),
	}

	return plans, &pagination, nil
}

func GetPlan(planID uint64) (*template.PlanData, error) {
	db := setup.DB

	var plan models.Plan
	if err := db.First(&plan, planID).Error; err != nil {
		return nil, err
	}

	var locks []models.Lock = make([]models.Lock, 0)
	if err := db.Order("created_at DESC").Find(&locks, "plan_id = ?", planID).Error; err != nil {
		return nil, err
	}

	data := template.PlanData{
		ID:          plan.ID,
		Name:        plan.Name,
		Width:       plan.Width,
		Height:      plan.Height,
		ImageURL:    plan.ImageURL,
		ImageHeight: plan.ImageHeight,
		ImageWidth:  plan.ImageWidth,
		Locks:       locks,
	}

	return &data, nil
}

func CreatePlan(r template.CreatePlanRequest) (*models.Plan, error) {
	db := setup.DB

	filename, img, err := r.SaveImage()
	if err != nil {
		return nil, err
	}

	plan := models.Plan{
		Name:        r.Name,
		Width:       r.Width,
		Height:      r.Height,
		ImageWidth:  img.Bounds().Size().X,
		ImageHeight: img.Bounds().Size().Y,
		ImageURL:    fmt.Sprintf("/api/plan/images/%s", filename),
	}

	if err := db.Create(&plan).Error; err != nil {
		return nil, err
	}

	return &plan, nil
}

func EditPlan(r template.CreatePlanRequest, planID uint64) (*models.Plan, error) {
	db := setup.DB

	var plan models.Plan
	if err := db.First(&plan, planID).Error; err != nil {
		return nil, err
	}

	if r.Name != "" {
		plan.Name = r.Name
	}

	if r.Height != 0 {
		plan.Height = r.Height
	}

	if r.Width != 0 {
		plan.Width = r.Width
	}

	if r.ImageBase64 != "" {
		r.Name = plan.Name
		filename, img, err := r.SaveImage()
		if err != nil {
			return nil, err
		}

		plan.ImageURL = fmt.Sprintf("/api/plan/images/%s", filename)
		plan.ImageWidth = img.Bounds().Size().X
		plan.ImageHeight = img.Bounds().Size().Y
	}

	if err := db.Save(&plan).Error; err != nil {
		return nil, err
	}

	return &plan, nil
}

func DeletePlan(planID uint64) error {
	db := setup.DB

	var plan models.Plan
	if err := db.First(&plan, planID).Error; err != nil {
		return err
	}

	if plan.ID == 0 {
		return errors.New("plan not found")
	}

	db.Delete(&plan)

	temp := strings.Split(plan.ImageURL, "/")
	os.Remove(fmt.Sprintf("plans/%s", temp[len(temp)-1]))

	return nil
}

func AddLockToPlan(r template.AddLockToPlanRequest, planID uint64) error {
	db := setup.DB

	var plan models.Plan
	if err := db.First(&plan, planID).Error; err != nil {
		return err
	}

	if plan.ID == 0 {
		return errors.New("plan not found")
	}

	var lock models.Lock
	if err := db.First(&lock, r.LockID).Error; err != nil {
		return err
	}

	if lock.ID == 0 {
		return errors.New("lock not found")
	}

	if lock.PlanID != 0 {
		return errors.New("lock already used in another plan")
	}

	lock.PlanID = planID
	lock.CoordX = r.CoordX
	lock.CoordY = r.CoordY

	if err := db.Save(&lock).Error; err != nil {
		return err
	}

	return nil
}

func EditLockInPlan(r template.AddLockToPlanRequest, planID uint64, lockID uint64) error {
	db := setup.DB

	var plan models.Plan
	if err := db.First(&plan, planID).Error; err != nil {
		return err
	}

	if plan.ID == 0 {
		return errors.New("plan not found")
	}

	var lock models.Lock
	if err := db.First(&lock, lockID).Error; err != nil {
		return err
	}

	if lock.ID == 0 {
		return errors.New("lock not found")
	}

	if lock.PlanID != planID {
		return errors.New("lock not used in this plan")
	}

	lock.CoordX = r.CoordX
	lock.CoordY = r.CoordY

	if err := db.Save(&lock).Error; err != nil {
		return err
	}

	return nil
}

func RemoveLockFromPlan(planID uint64, lockID uint64) error {
	db := setup.DB

	var plan models.Plan
	if err := db.First(&plan, planID).Error; err != nil {
		return err
	}

	if plan.ID == 0 {
		return errors.New("plan not found")
	}

	var lock models.Lock
	if err := db.First(&lock, lockID).Error; err != nil {
		return err
	}

	if lock.ID == 0 {
		return errors.New("lock not found")
	}

	if lock.PlanID != planID {
		return errors.New("lock not used in this plan")
	}

	lock.PlanID = 0
	lock.CoordX = 0
	lock.CoordY = 0

	if err := db.Save(&lock).Error; err != nil {
		return err
	}

	return nil
}
