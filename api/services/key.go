package services

import (
	"errors"
	"fmt"
	"server/api/template"
	"server/models"
	"server/setup"

	"gorm.io/gorm"
)

func GetKeys(p template.SearchParameter, keyword string, status string, notOwned bool) ([]template.KeyData, *template.Pagination, error) {
	db := setup.DB

	offset := (p.Page - 1) * p.Limit
	limit := p.Limit
	keyword = "%" + keyword + "%"

	var queryString string
	var keys []template.KeyData
	var cnt int64

	if notOwned {
		var ids []uint64
		if err := db.Raw("SELECT key_id FROM personels WHERE key_id != 0").Scan(&ids).Error; err != nil {
			return nil, nil, err
		}

		queryString = `
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
			personels.id NOT IN ? AND 
			(
				keys.label LIKE ? OR
				keys.key_id LIKE ? OR
				personels.name LIKE ?
			)`
		if err := db.Raw(
			fmt.Sprintf("SELECT COUNT(*) AS cnt FROM ( %s ) AS t", queryString),
			keyword, keyword, keyword).
			Scan(&cnt).Error; err != nil {
			return nil, nil, err
		}

		if p.Limit < 0 {
			if err := db.Raw(
				fmt.Sprintf("%s ORDER BY keys.created_at DESC", queryString),
				keyword, keyword, keyword).
				Scan(&keys).Error; err != nil {
				return nil, nil, err
			}
		} else {
			if err := db.Raw(
				fmt.Sprintf("%s ORDER BY keys.created_at DESC OFFSET ? LIMIT ?", queryString),
				keyword, keyword, keyword, offset, limit).
				Scan(&keys).Error; err != nil {
				return nil, nil, err
			}
		}
	} else {
		if status == "any" {
			queryString = `
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
				personels.name LIKE ?`

			if err := db.Raw(
				fmt.Sprintf("SELECT COUNT(*) AS cnt FROM ( %s ) AS t", queryString),
				keyword, keyword, keyword).
				Scan(&cnt).Error; err != nil {
				return nil, nil, err
			}

			if p.Limit < 0 {
				if err := db.Raw(
					fmt.Sprintf("%s ORDER BY keys.created_at DESC", queryString),
					keyword, keyword, keyword).
					Scan(&keys).Error; err != nil {
					return nil, nil, err
				}
			} else {
				if err := db.Raw(
					fmt.Sprintf("%s ORDER BY keys.created_at DESC OFFSET ? LIMIT ?", queryString),
					keyword, keyword, keyword, offset, limit).
					Scan(&keys).Error; err != nil {
					return nil, nil, err
				}
			}
		} else {
			queryString = `
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
				keys.status = ? AND
				(
					keys.label LIKE ? OR
					keys.key_id LIKE ? OR
					personels.name LIKE ?
				)`

			if err := db.Raw(
				fmt.Sprintf("SELECT COUNT(*) AS cnt FROM ( %s ) AS t", queryString),
				status == "active", keyword, keyword, keyword).
				Scan(&cnt).Error; err != nil {
				return nil, nil, err
			}

			if p.Limit < 0 {
				if err := db.Raw(
					fmt.Sprintf("%s ORDER BY keys.created_at DESC", queryString),
					status == "active", keyword, keyword, keyword).
					Scan(&keys).Error; err != nil {
					return nil, nil, err
				}
			} else {
				if err := db.Raw(
					fmt.Sprintf("%s ORDER BY keys.created_at DESC OFFSET ? LIMIT ?", queryString),
					status == "active", keyword, keyword, keyword, offset, limit).
					Scan(&keys).Error; err != nil {
					return nil, nil, err
				}
			}
		}
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

	key.Status = e.Status

	if err := db.Save(&key).Error; err != nil {
		return nil, err
	}

	return &key, nil
}
