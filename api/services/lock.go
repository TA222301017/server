package services

import (
	"errors"
	"server/api/template"
	"server/models"
	"server/setup"

	"gorm.io/gorm"
)

func GetLocks(p template.SearchParameter, keyword string, status string) ([]models.Lock, *template.Pagination, error) {
	db := setup.DB

	offset := (p.Page - 1) * p.Limit
	limit := p.Limit
	keyword = "%" + keyword + "%"

	var query *gorm.DB
	if status == "any" {
		query = db.Where("label LIKE ? OR location LIKE ?", keyword, keyword)
	} else if status == "unused" {
		query = db.Where("(label LIKE ? OR location LIKE ?) AND (plan_id = 0 OR plan_id = NULL)", keyword, keyword)
	} else {
		query = db.Where("(label LIKE ? OR location LIKE ?) AND status = ?", keyword, keyword, status == "active")
	}

	var cnt int64
	if err := query.Find(&models.Lock{}).Count(&cnt).Error; err != nil {
		return nil, nil, err
	}

	if p.Limit > 0 {
		query = query.Limit(limit).Offset(offset)
	}

	var locks []models.Lock
	if err := query.Preload("Plan").Order("created_at DESC").Find(&locks).Error; err != nil {
		return nil, nil, err
	}

	last := cnt / int64(limit)
	pagination := template.Pagination{
		Page:  p.Page,
		Limit: p.Limit,
		Last:  int(last),
		Total: int(cnt),
	}

	return locks, &pagination, nil
}

func GetLock(lockID uint64) (*models.Lock, error) {
	db := setup.DB

	var lock models.Lock
	if err := db.First(&lock, lockID).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, RecordsNotFound
	}

	return &lock, nil
}

func EditLock(e template.EditLockRequest, lockID uint64) (*models.Lock, error) {
	db := setup.DB

	var lock models.Lock
	if err := db.First(&lock, lockID).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, RecordsNotFound
	}

	if e.Description != "" {
		lock.Description = e.Description
	}

	if e.Location != "" {
		lock.Location = e.Location
	}

	if e.Name != "" {
		lock.Label = e.Name
	}

	if err := db.Save(&lock).Error; err != nil {
		return nil, err
	}

	return &lock, nil
}

func DeleteLock(lockID uint64) error {
	db := setup.DB

	var lock models.Lock
	if err := db.First(&lock, lockID).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return RecordsNotFound
	}

	if err := db.Delete(&models.AccessRule{}, "lock_id = ?", lockID).Error; err != nil {
		return err
	}

	if err := db.Delete(&models.Lock{}, lockID).Error; err != nil {
		return err
	}

	return nil
}
