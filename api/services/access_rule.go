package services

import (
	"errors"
	"fmt"
	"server/api/template"
	"server/models"
	"server/setup"
	"server/udp/usecases"

	"gorm.io/gorm"
)

func GetAccessRules(p template.SearchParameter, keyword string) ([]template.AccessRuleData, *template.Pagination, error) {
	db := setup.DB

	keyword = "%" + keyword + "%"
	offset := (p.Page - 1) * p.Limit
	limit := p.Limit

	var cnt int64
	var queryString string
	var rules []template.AccessRuleData = make([]template.AccessRuleData, 0)

	queryString = `
	SELECT
		access_rules.id,
		access_rules.personel_id,
		access_rules.lock_id,
		access_rules.key_id,
		access_rules.starts_at,
		access_rules.ends_at,
		personels.name AS personel,
		locks.label AS lock,
		locks.location AS location,
		keys.label AS key
	FROM access_rules
	LEFT JOIN personels ON personels.id = access_rules.personel_id
	LEFT JOIN locks ON locks.id = access_rules.lock_id
	LEFT JOIN keys ON keys.id = access_rules.key_id
	WHERE
		access_rules.starts_at >= ? AND
		access_rules.ends_at <= ? AND (
			personels.name LIKE ? OR
			locks.label LIKE ? OR
			locks.location LIKE ? OR
			keys.label LIKE ?
		)`

	if err := db.Raw(
		fmt.Sprintf("SELECT COUNT(*) AS cnt FROM ( %s ) AS t", queryString),
		p.StartDate, p.EndDate, keyword, keyword, keyword, keyword).
		Scan(&cnt).Error; err != nil {
		return nil, nil, err
	}

	if p.Limit < 0 {
		if err := db.Raw(
			fmt.Sprintf("%s ORDER BY access_rules.created_at DESC", queryString),
			p.StartDate, p.EndDate, keyword, keyword, keyword, keyword).
			Scan(&rules).Error; err != nil {
			return nil, nil, err
		}
	} else {
		if err := db.Raw(
			fmt.Sprintf("%s ORDER BY access_rules.created_at DESC OFFSET ? LIMIT ?", queryString),
			p.StartDate, p.EndDate, keyword, keyword, keyword, keyword, offset, limit).
			Scan(&rules).Error; err != nil {
			return nil, nil, err
		}
	}

	last := cnt / int64(limit)
	pagination := template.Pagination{
		Page:  p.Page,
		Limit: p.Limit,
		Last:  int(last),
		Total: int(cnt),
	}

	return rules, &pagination, nil
}

func GetPersonelAccessRules(p template.SearchParameter, personelID uint64) ([]template.AccessRuleData, *template.Pagination) {
	db := setup.DB

	offset := (p.Page - 1) * p.Limit
	limit := p.Limit

	var cnt int64
	db.Where("personel_id = ?", personelID).Find(&models.AccessRule{}).Count(&cnt)

	var accessRules []models.AccessRule
	db.Offset(offset).Limit(limit).
		Where("personel_id = ?", personelID).
		Preload("Lock").Preload("Personel").Find(&accessRules)

	var accessRuleData []template.AccessRuleData
	for _, a := range accessRules {
		accessRuleData = append(accessRuleData, template.AccessRuleData{
			ID:         a.ID,
			StartsAt:   a.StartsAt,
			EndsAt:     a.EndsAt,
			Lock:       a.Lock.Label,
			Location:   a.Lock.Location,
			LockID:     a.LockID,
			PersonelID: a.PersonelID,
			Personel:   a.Personel.Name,
			KeyID:      a.KeyID,
		})
	}

	last := cnt / int64(limit)
	pagination := template.Pagination{
		Page:  p.Page,
		Limit: p.Limit,
		Last:  int(last),
		Total: int(cnt),
	}

	return accessRuleData, &pagination
}

func AddAccessRule(a template.AddAccessRule, userID uint64) (*template.AccessRuleData, error) {
	db := setup.DB

	var lock models.Lock
	if err := db.First(&lock, a.LockID).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("lock not found")
	}

	var personel models.Personel
	if err := db.First(&personel, a.PersonelID).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("personel not found")
	}

	var key models.Key
	if err := db.First(&key, personel.KeyID).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("key not found")
	}

	accessRule := models.AccessRule{
		PersonelID: a.PersonelID,
		LockID:     a.LockID,
		KeyID:      key.ID,
		CreatorID:  userID,
		StartsAt:   a.StartsAt,
		EndsAt:     a.EndsAt,
	}

	if err := db.Create(&accessRule).Error; err != nil {
		return nil, err
	}

	if _, err := usecases.AddAccessRule(accessRule, lock, key); err != nil {
		db.Delete(&accessRule)
		return nil, err
	}

	return &template.AccessRuleData{
		ID:         accessRule.ID,
		StartsAt:   accessRule.StartsAt,
		EndsAt:     accessRule.EndsAt,
		Lock:       lock.Label,
		LockID:     accessRule.LockID,
		Location:   lock.Location,
		PersonelID: accessRule.PersonelID,
		Personel:   personel.Name,
		KeyID:      accessRule.KeyID,
	}, nil
}

func EditAccessRule(e template.EditAccessRule, userID uint64, accessRuleID uint64) (*template.AccessRuleData, error) {
	db := setup.DB

	var accessRule models.AccessRule
	if err := db.Preload("Personel").First(&accessRule, accessRuleID).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("access rule not found")
	}

	var lock models.Lock
	if e.LockID != 0 {
		if err := db.First(&lock, e.LockID).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("lock not found")
		}
		accessRule.LockID = e.LockID
	} else {
		if err := db.First(&lock, accessRule.LockID).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("lock not found")
		}
	}

	var key models.Key
	if err := db.First(&key, accessRule.KeyID).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("key not found")
	}

	if !e.StartsAt.IsZero() {
		accessRule.StartsAt = e.StartsAt
	}

	if !e.EndsAt.IsZero() {
		accessRule.EndsAt = e.EndsAt
	}

	accessRule.CreatorID = userID

	if _, err := usecases.EditAccessRule(accessRule, lock, key); err != nil {
		return nil, err
	}

	if err := db.Save(&accessRule).Error; err != nil {
		return nil, err
	}

	return &template.AccessRuleData{
		ID:         accessRuleID,
		StartsAt:   accessRule.StartsAt,
		EndsAt:     accessRule.EndsAt,
		Lock:       lock.Label,
		Personel:   accessRule.Personel.Name,
		LockID:     accessRule.LockID,
		Location:   lock.Location,
		KeyID:      accessRule.KeyID,
		PersonelID: accessRule.PersonelID,
	}, nil
}

func DeleteAccessRule(accessRuleID uint64) error {
	db := setup.DB

	var accessRule models.AccessRule
	if err := db.First(&accessRule, accessRuleID).Error; err != nil {
		return err
	}

	var lock models.Lock
	if err := db.First(&lock, accessRule.LockID).Error; err != nil {
		return err
	}

	if _, err := usecases.DeleteAccessRule(accessRuleID, lock.IpAddress); err != nil {
		return err
	}

	return db.Delete(&models.AccessRule{}, accessRuleID).Error
}
