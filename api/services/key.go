package services

import (
	"errors"
	"server/api/template"
	"server/models"
	"server/setup"

	"gorm.io/gorm"
)

func GetKeys(p template.SearchParameter, keyword string, status bool) ([]template.KeyData, *template.Pagination, error) {
	db := setup.DB

	offset := (p.Page - 1) * p.Limit
	limit := p.Limit
	keyword = "%" + keyword + "%"

	var cnt int64
	if err := db.
		Where("label LIKE ? OR key_id LIKE ?", keyword, keyword).
		Find(&models.Key{}).Count(&cnt).Error; err != nil {
		return nil, nil, err
	}

	var keys []template.KeyData
	if err := db.Raw(`
		SELECT 
			keys.id AS id, 
			keys.key_id AS key_id, 
			keys.status AS status, 
			keys.label AS name,
			personels.name AS owner,
			personels.id AS owner_id 
		FROM keys 
		LEFT JOIN personels
		ON keys.id = personels.key_id
		WHERE
			keys.label LIKE ? OR
			keys.key_id LIKE ? OR
			personels.name LIKE ?
		ORDER BY keys.created_at DESC
		OFFSET ? LIMIT ?
	`, keyword, keyword, keyword, offset, limit).
		Scan(&keys).Error; err != nil {
		return nil, nil, err
	}

	last := cnt / int64(limit)
	pagination := template.Pagination{
		Page:  p.Page,
		Limit: p.Limit,
		Last:  int(last),
		Total: int(cnt),
	}

	return keys, &pagination, nil
}

func GetKey(keyID uint64) (*models.Key, error) {
	db := setup.DB

	var key models.Key
	if err := db.First(&key, keyID).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, RecordsNotFound
	}

	return &key, nil
}

func AddKey(a template.AddKeyRequest) (*models.Key, error) {
	db := setup.DB

	key := models.Key{
		Label:       a.Name,
		KeyID:       a.KeyID,
		Status:      a.Status,
		Description: a.Description,
	}

	if err := db.Create(&key).Error; err != nil {
		return nil, err
	}

	return &key, nil
}

func EditKey(e template.EditKeyRequest, keyID uint64) (*models.Key, error) {
	db := setup.DB

	var key models.Key
	if err := db.First(&key, keyID).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, RecordsNotFound
	}

	if e.Description != "" {
		key.Description = e.Description
	}

	if e.KeyID != "" {
		key.KeyID = e.KeyID
	}

	if e.Name != "" {
		key.Label = e.Name
	}

	if err := db.Save(&key).Error; err != nil {
		return nil, err
	}

	return &key, nil
}
