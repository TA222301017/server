package services

import (
	"errors"
	"server/api/template"
	"server/models"
	"server/setup"

	"gorm.io/gorm"
)

func RegisterPersonel(a template.AddPersonelRequest) (*template.PersonelData, error) {
	db := setup.DB

	var role models.Role
	if err := db.First(&role, a.RoleID).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("role not found")
	}

	if err := db.First(&models.Key{}, a.KeyID).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("key not found")
	}

	personel := models.Personel{
		IDNumber:    a.PersonelID,
		Name:        a.Name,
		Status:      a.Status,
		RoleID:      a.RoleID,
		KeyID:       a.KeyID,
		Description: a.Description,
	}

	if err := db.Create(&personel).Error; err != nil {
		return nil, err
	}

	personelData := template.PersonelData{
		ID:          personel.ID,
		Name:        personel.Name,
		PersonelID:  personel.IDNumber,
		Status:      personel.Status,
		Role:        role.Name,
		Description: personel.Description,
	}

	return &personelData, nil
}

func GetPersonels(p template.SearchParameter, status bool, keyword string) ([]template.PersonelData, *template.Pagination, error) {
	db := setup.DB

	offset := (p.Page - 1) * p.Limit
	limit := p.Limit
	keyword = "%" + keyword + "%"

	var cnt int64
	if err := db.
		Where("name LIKE ? OR id_number LIKE ?", keyword, keyword).
		Find(&models.Personel{}).Count(&cnt).Error; err != nil {
		return nil, nil, err
	}

	var personels []models.Personel
	if err := db.
		Limit(limit).Offset(offset).
		Where("name LIKE ? OR id_number LIKE ?", keyword, keyword).
		Preload("Role").Find(&personels).Error; err != nil {
		return nil, nil, err
	}

	last := cnt / int64(limit)
	pagination := template.Pagination{
		Page:  p.Page,
		Limit: p.Limit,
		Last:  int(last),
		Total: int(cnt),
	}

	var data []template.PersonelData
	for _, p := range personels {
		data = append(data, template.PersonelData{
			ID:          p.ID,
			Name:        p.Name,
			PersonelID:  p.IDNumber,
			Status:      p.Status,
			Role:        p.Role.Name,
			Description: p.Description,
			RoleID:      p.RoleID,
			KeyID:       p.KeyID,
		})
	}

	return data, &pagination, nil
}

func GetPersonel(personelID uint64) (*template.PersonelData, error) {
	db := setup.DB

	var p models.Personel
	if err := db.First(&p, personelID).Preload("Role").Error; err != nil {
		return nil, err
	}

	data := template.PersonelData{
		ID:          p.ID,
		Name:        p.Name,
		PersonelID:  p.IDNumber,
		Status:      p.Status,
		Role:        p.Role.Name,
		Description: p.Description,
		RoleID:      p.RoleID,
		KeyID:       p.KeyID,
	}

	return &data, nil
}

func EditPersonel(e template.EditPersonelRequest, personelID uint64) (*template.PersonelData, error) {
	db := setup.DB

	var (
		p    models.Personel
		role models.Role
	)

	if err := db.First(&p, personelID).Preload("Role").Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, RecordsNotFound
	}
	role = p.Role

	if e.KeyID != 0 {
		p.KeyID = e.KeyID
	}

	if e.Description != "" {
		p.Description = e.Description
	}

	if e.Name != "" {
		p.Name = e.Name
	}

	if e.PersonelID != "" {
		p.IDNumber = e.PersonelID
	}

	if e.RoleID != 0 {
		if err := db.First(&role, e.RoleID).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("role not found")
		}

		p.RoleID = e.RoleID
	}

	p.Status = e.Status

	if err := db.Save(&p).Error; err != nil {
		return nil, err
	}

	data := template.PersonelData{
		ID:          p.ID,
		Name:        p.Name,
		PersonelID:  p.IDNumber,
		Status:      p.Status,
		Role:        role.Name,
		Description: p.Description,
		RoleID:      p.RoleID,
		KeyID:       p.KeyID,
	}

	return &data, nil
}

func GetRoles() []models.Role {
	db := setup.DB

	var roles []models.Role
	db.Find(&roles)

	return roles
}
