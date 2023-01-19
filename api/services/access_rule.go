package services

import (
	"errors"
	"server/api/template"
	"server/models"
	"server/setup"
	"server/udp/usecases"

	"gorm.io/gorm"
)

func GetPersonelAccessRules(personelID uint64) []template.AccessRuleData {
	db := setup.DB

	var accessRules []models.AccessRule
	db.Select("id", "starts_at", "ends_at").
		Find(&accessRules).
		Where("personel_id = ?", personelID).
		Preload("Lock")

	var accessRuleData []template.AccessRuleData
	for _, a := range accessRules {
		accessRuleData = append(accessRuleData, template.AccessRuleData{
			ID:       a.ID,
			StartsAt: a.StartsAt,
			EndsAt:   a.EndsAt,
			Lock:     a.Lock.Label,
		})
	}

	return accessRuleData
}

func AddAccessRule(a template.AddAccessRule, userID uint64) (*template.AccessRuleData, error) {
	db := setup.DB

	var lock models.Lock
	if err := db.First(&lock, a.LockID).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("lock not found")
	}

	if err := db.First(&models.Personel{}, a.PersonelID).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("personel not found")
	}

	accessRule := models.AccessRule{
		PersonelID: a.PersonelID,
		LockID:     a.LockID,
		Lock:       lock,
		CreatorID:  userID,
		StartsAt:   a.StartsAt,
		EndsAt:     a.EndsAt,
	}

	if _, err := usecases.AddAccessRule(accessRule); err != nil {
		return nil, err
	}

	if err := db.Create(&accessRule).Error; err != nil {
		return nil, err
	}

	return &template.AccessRuleData{
		ID:       accessRule.ID,
		StartsAt: accessRule.StartsAt,
		EndsAt:   accessRule.EndsAt,
		Lock:     lock.Label,
	}, nil
}

func EditAccessRule(e template.EditAccessRule, userID uint64, accessRuleID uint64) (*template.AccessRuleData, error) {
	db := setup.DB

	var accessRule models.AccessRule
	if err := db.First(&accessRule, accessRuleID).Error; errors.Is(err, gorm.ErrRecordNotFound) {
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

	if !e.StartsAt.IsZero() {
		accessRule.StartsAt = e.StartsAt
	}

	if !e.EndsAt.IsZero() {
		accessRule.EndsAt = e.EndsAt
	}

	accessRule.CreatorID = userID

	if _, err := usecases.EditAccessRule(accessRule); err != nil {
		return nil, err
	}

	if err := db.Save(&accessRule).Error; err != nil {
		return nil, err
	}

	return &template.AccessRuleData{
		ID:       accessRuleID,
		StartsAt: accessRule.StartsAt,
		EndsAt:   accessRule.EndsAt,
		Lock:     lock.Label,
	}, nil
}

func DeleteAccessRule(accessRuleID uint64) error {
	db := setup.DB

	var accessRule models.AccessRule
	if err := db.First(&accessRule, accessRuleID).Preload("Lock").Error; err != nil {
		return err
	}

	if _, err := usecases.DeleteAccessRule(accessRuleID, accessRule.Lock.IpAddress); err != nil {
		return err
	}

	return db.Delete(&models.AccessRule{}, accessRuleID).Error
}
