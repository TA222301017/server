package services

import (
	"errors"
	"server/api/template"
	"server/models"
	"server/setup"

	"gorm.io/gorm"
)

func GetLocks(p template.SearchParameter, keyword string, status bool) ([]models.Lock, *template.Pagination, error) {
	db := setup.DB

	offset := (p.Page - 1) * p.Limit
	limit := p.Limit
	keyword = "%" + keyword + "%"

	var cnt int64
	if err := db.
		Where("label LIKE ? OR location LIKE ?", keyword, keyword).
		Find(&models.Lock{}).Count(&cnt).Error; err != nil {
		return nil, nil, err
	}

	var locks []models.Lock
	if err := db.
		Limit(limit).Offset(offset).
		Where("label LIKE ? OR location LIKE ?", keyword, keyword).
		Find(&locks).Error; err != nil {
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
