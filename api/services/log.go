package services

import (
	"errors"
	"server/api/template"
	"server/models"
	"server/setup"

	"gorm.io/gorm"
)

func GetAccessLog(p template.SearchParameter, location string, personel string) ([]models.AccessLog, *template.Pagination, error) {
	db := setup.DB

	location = "%" + location + "%"
	personel = "%" + personel + "%"
	offset := (p.Page - 1) * p.Limit
	limit := p.Limit

	var cnt int64
	if err := db.
		Where("location LIKE ? OR personel_name = ? OR personel_id_number = ?", location, personel, personel).
		Find(&models.AccessLog{}).Count(&cnt).Error; err != nil {
		return nil, nil, err
	}

	var data []models.AccessLog
	if err := db.
		Limit(limit).Offset(offset).
		Where("location LIKE ? OR personel_name = ? OR personel_id_number = ?", location, personel, personel).
		Find(&data).Error; err != nil {
		return nil, nil, err
	}

	last := cnt / int64(limit)
	pagination := template.Pagination{
		Page:  p.Page,
		Limit: p.Limit,
		Last:  int(last),
		Total: int(cnt),
	}

	return data, &pagination, nil
}

func GetRSSILog(p *template.SearchParameter, personelID uint64) ([]template.RSSILogData, *template.Pagination, error) {
	db := setup.DB

	offset := (p.Page - 1) * p.Limit
	limit := p.Limit

	var personel models.Personel
	if err := db.First(&personel, personelID).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil, err
	}

	var cnt int64
	if err := db.
		Limit(limit).Offset(offset).
		Where("key_id = ? AND timestamp BETWEEN ? AND ?", personel.KeyID, p.StartDate, p.EndDate).
		Count(&cnt).Error; err != nil {
		return nil, nil, err
	}

	var logs []models.RSSILog
	if err := db.
		Limit(limit).Offset(offset).
		Where("key_id = ? AND timestamp BETWEEN ? AND ?", personel.KeyID, p.StartDate, p.EndDate).
		Preload("Lock").Find(&logs).Error; err != nil {
		return nil, nil, err
	}

	var data []template.RSSILogData
	for _, l := range logs {
		data = append(data, template.RSSILogData{
			Timestamp:    l.Timestamp,
			LockName:     l.Lock.Label,
			LockLocation: l.Lock.Location,
		})
	}

	last := cnt / int64(limit)
	pagination := template.Pagination{
		Page:  p.Page,
		Limit: p.Limit,
		Last:  int(last),
		Total: int(cnt),
	}

	return data, &pagination, nil
}

func GetHealthcheckLog(p *template.SearchParameter, personelID uint64) ([]template.RSSILogData, *template.Pagination, error) {
	db := setup.DB

	offset := (p.Page - 1) * p.Limit
	limit := p.Limit

	var personel models.Personel
	if err := db.First(&personel, personelID).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil, err
	}

	var cnt int64
	if err := db.
		Limit(limit).Offset(offset).
		Where("key_id = ? AND timestamp BETWEEN ? AND ?", personel.KeyID, p.StartDate, p.EndDate).
		Count(&cnt).Error; err != nil {
		return nil, nil, err
	}

	var logs []models.RSSILog
	if err := db.
		Limit(limit).Offset(offset).
		Where("key_id = ? AND timestamp BETWEEN ? AND ?", personel.KeyID, p.StartDate, p.EndDate).
		Preload("Lock").Find(&logs).Error; err != nil {
		return nil, nil, err
	}

	var data []template.RSSILogData
	for _, l := range logs {
		data = append(data, template.RSSILogData{
			Timestamp:    l.Timestamp,
			LockName:     l.Lock.Label,
			LockLocation: l.Lock.Location,
		})
	}

	last := cnt / int64(limit)
	pagination := template.Pagination{
		Page:  p.Page,
		Limit: p.Limit,
		Last:  int(last),
		Total: int(cnt),
	}

	return data, &pagination, nil
}
