package services

import (
	"fmt"
	"server/api/template"
	"server/models"
	"server/setup"
	"server/udp/usecases"
	"sync"
	"time"
)

func GetAccessLog(p template.SearchParameter, location string, personel string) ([]template.AccessLogData, *template.Pagination, error) {
	db := setup.DB

	location = "%" + location + "%"
	personel = "%" + personel + "%"
	offset := (p.Page - 1) * p.Limit
	limit := p.Limit

	var cnt int64
	if err := db.
		Where(`
			location LIKE ? OR 
			personel_name LIKE ? OR 
			personel_id_number LIKE ? AND 
			timestamp BETWEEN ? AND ?
		`, location, personel, personel, p.StartDate, p.EndDate).
		Find(&models.AccessLog{}).Count(&cnt).Error; err != nil {
		return nil, nil, err
	}

	var logs []models.AccessLog
	if err := db.
		Limit(limit).Offset(offset).
		Where(`
			location LIKE ? OR 
			personel_name LIKE ? OR 
			personel_id_number LIKE ? AND 
			timestamp BETWEEN ? AND ?
		`, location, personel, personel, p.StartDate, p.EndDate).
		Preload("Lock").Preload("Key").
		Find(&logs).Error; err != nil {
		return nil, nil, err
	}

	var data []template.AccessLogData = make([]template.AccessLogData, 0)
	for _, l := range logs {
		data = append(data, template.AccessLogData{
			ID:         l.ID,
			Personel:   l.PersonelName,
			PersonelID: 0,
			Lock:       l.Lock.Label,
			LockID:     l.LockID,
			Key:        l.Key.Label,
			KeyID:      l.KeyID,
			Location:   l.Location,
			Timestamp:  l.Timestamp,
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

func GetRSSILog(p *template.SearchParameter, keyword string) ([]template.RSSILogData, *template.Pagination, error) {
	db := setup.DB

	keyword = "%" + keyword + "%"
	offset := (p.Page - 1) * p.Limit
	limit := p.Limit

	var cnt int64
	if err := db.Raw(`
	SELECT cnt FROM (
		SELECT
			COUNT(*) AS cnt
		FROM rssi_logs
			LEFT JOIN personels ON rssi_logs.personel_id = personels.id
			LEFT JOIN locks ON rssi_logs.lock_id = locks.id
			LEFT JOIN keys ON rssi_logs.key_id = keys.id
		WHERE
			rssi_logs.timestamp BETWEEN ? AND ? OR
			personels.name LIKE ? OR 
			locks.label LIKE ? OR 
			keys.label LIKE ?
	) AS t
	`, p.StartDate, p.EndDate, keyword, keyword, keyword).
		Scan(&cnt).Error; err != nil {
		return nil, nil, err
	}

	var data []template.RSSILogData
	if err := db.Raw(`
	SELECT
		rssi_logs.*,
		personels.name AS personel, 
		locks.label AS lock, 
		locks.location AS location, 
		keys.label AS key
	FROM rssi_logs
		LEFT JOIN personels ON rssi_logs.personel_id = personels.id
		LEFT JOIN locks ON rssi_logs.lock_id = locks.id
		LEFT JOIN keys ON rssi_logs.key_id = keys.id
	WHERE
		rssi_logs.timestamp BETWEEN ? AND ? OR
		personels.name LIKE ? OR 
		locks.label LIKE ? OR 
		keys.label LIKE ?
	ORDER BY rssi_logs.created_at DESC
	OFFSET ? LIMIT ?
	`, p.StartDate, p.EndDate, keyword, keyword, keyword, offset, limit).
		Scan(&data).Error; err != nil {
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

func GetHealthcheckLog(p *template.SearchParameter, location string, status string) ([]template.HealthcheckLogData, *template.Pagination, error) {
	db := setup.DB

	location = "%" + location + "%"
	offset := (p.Page - 1) * p.Limit
	limit := p.Limit

	var queryString string
	var cnt int64
	var data []template.HealthcheckLogData

	if status == "any" {
		queryString = `
		SELECT
			healthcheck_logs.id AS id,
			healthcheck_logs.status AS status,
			healthcheck_logs.lock_id AS device_id,
			healthcheck_logs.timestamp AS timestamp,
			locks.location AS location,
			locks.label AS device
		FROM healthcheck_logs
		LEFT JOIN locks
		ON healthcheck_logs.lock_id = locks.id
		WHERE
			(healthcheck_logs.timestamp BETWEEN ? AND ?) AND
			(locks.location LIKE ? OR locks.label LIKE ?)`

		if err := db.Raw(
			fmt.Sprintf("SELECT COUNT(*) AS cnt FROM ( %s ) AS t", queryString),
			p.StartDate, p.EndDate, location, location).
			Scan(&cnt).Error; err != nil {
			return nil, nil, err
		}

		if err := db.Raw(
			fmt.Sprintf("%s ORDER BY healthcheck_logs.timestamp DESC OFFSET ? LIMIT ?", queryString),
			p.StartDate, p.EndDate, location, location, offset, limit).
			Scan(&data).Error; err != nil {
			return nil, nil, err
		}
	} else {
		queryString = `
		SELECT
			healthcheck_logs.id AS id,
			healthcheck_logs.status AS status,
			healthcheck_logs.lock_id AS device_id,
			healthcheck_logs.timestamp AS timestamp,
			locks.location AS location,
			locks.label AS device
		FROM healthcheck_logs
		LEFT JOIN locks
		ON healthcheck_logs.lock_id = locks.id
		WHERE
			(healthcheck_logs.timestamp BETWEEN ? AND ?) AND
			healthcheck_logs.status = ? AND
			(locks.location LIKE ? OR locks.label LIKE ?)`

		if err := db.Raw(
			fmt.Sprintf("SELECT COUNT(*) AS cnt FROM ( %s ) AS t", queryString),
			p.StartDate, p.EndDate, status == "active", location, location).
			Scan(&cnt).Error; err != nil {
			return nil, nil, err
		}

		if err := db.Raw(
			fmt.Sprintf("%s ORDER BY healthcheck_logs.timestamp DESC OFFSET ? LIMIT ?", queryString),
			p.StartDate, p.EndDate, status == "active", location, location, offset, limit).
			Scan(&data).Error; err != nil {
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

	return data, &pagination, nil
}

func CheckLocks() ([]models.HealthcheckLog, error) {
	db := setup.DB

	var locks []models.Lock
	if err := db.Find(&locks).Error; err != nil {
		return nil, err
	}

	var status []models.HealthcheckLog
	wg := new(sync.WaitGroup)
	wg.Add(len(locks))
	for _, l := range locks {
		go func(l models.Lock) {
			_, err := usecases.RequestHealthcheck(&l)
			status = append(status, models.HealthcheckLog{
				LockID:    l.ID,
				Timestamp: time.Now(),
				Status:    err == nil,
			})

			l.Status = err == nil
			db.Save(&l)

			wg.Done()
		}(l)
	}

	wg.Wait()

	if err := db.Create(&status).Error; err != nil {
		return nil, err
	}

	return status, nil
}

func CheckLock(lockID uint64) (*models.HealthcheckLog, error) {
	db := setup.DB

	var lock models.Lock
	if err := db.First(&lock, lockID).Error; err != nil {
		return nil, err
	}

	var status models.HealthcheckLog
	_, err := usecases.RequestHealthcheck(&lock)
	status = models.HealthcheckLog{
		LockID:    lock.ID,
		Timestamp: time.Now(),
		Status:    err == nil,
	}

	if err := db.Create(&status).Error; err != nil {
		return nil, err
	}

	return &status, nil
}
