package services

import (
	"server/api/template"
	"server/api/utils"
	"server/models"
	"server/setup"
)

func CreateAdmin(r template.CreateAdminRequest) (*models.User, error) {
	db := setup.DB

	var cnt int64
	if err := db.Find(&models.User{}).Where("username = ?", r.Username).Count(&cnt).Error; err != nil {
		return nil, err
	}

	var admin models.User
	admin = models.User{
		Name:     r.Name,
		Username: admin.Username,
		Password: utils.HashPassword(admin.Password),
	}

	if err := db.Create(&admin).Error; err != nil {
		return nil, err
	}

	return &admin, nil
}

func GetAdmins(p template.SearchParameter, keyword string) ([]models.User, *template.Pagination, error) {
	db := setup.DB

	offset := (p.Page - 1) * p.Limit
	page := p.Page
	limit := p.Limit
	keyword = "%" + keyword + "%"

	var cnt int64
	if err := db.Find(&models.User{}).Where("name LIKE ?", keyword).Count(&cnt).Error; err != nil {
		return nil, nil, err
	}

	var admins []models.User
	if err := db.Find(&admins).Where("name LIKE ?", keyword).Offset(offset).Limit(limit).Error; err != nil {
		return nil, nil, err
	}

	last := cnt / int64(limit)
	pagination := template.Pagination{
		Page:  page,
		Limit: limit,
		Last:  int(last),
		Total: int(cnt),
	}

	return admins, &pagination, nil
}

func GetAdmin(adminID uint64) (*models.User, error) {
	db := setup.DB

	var admin models.User
	if err := db.First(&admin, adminID).Error; err != nil {
		return nil, err
	}

	return &admin, nil
}

func EditAdmin(adminID uint64, r template.EditAdminRequest) (*models.User, error) {
	db := setup.DB

	var admin models.User
	if err := db.First(&admin, adminID).Error; err != nil {
		return nil, err
	}

	if r.Name != "" {
		admin.Name = r.Name
	}

	if r.Username != "" {
		admin.Username = r.Username
	}

	if r.Password != "" {
		admin.Password = utils.HashPassword(r.Password)
	}

	if err := db.Save(&admin).Error; err != nil {
		return nil, err
	}

	return &admin, nil
}

func DeleteAdmin(adminID uint64) error {
	db := setup.DB

	if err := db.Delete(&models.User{}, "id = ?", adminID).Error; err != nil {
		return err
	}

	return nil
}
