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

	var cnt int64
	if err := db.Find(&models.Personel{}).Count(&cnt).Error; err != nil {
		return nil, nil, err
	}

	var personels []models.Personel
	if err := db.Find(&personels).Preload("Role").Offset(offset).Limit(limit).Error; err != nil {
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
	}

	return &data, nil
}
