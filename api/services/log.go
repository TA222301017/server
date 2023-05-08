package services

import (
	"errors"
	"fmt"
	"server/api/template"
	"server/models"
	"server/setup"
	"server/udp/usecases"
	"strings"
	"sync"
	"time"

	"gorm.io/gorm"
)

func GetAccessLog(p template.SearchParameter, location string, personel string) ([]template.AccessLogData, *template.Pagination, error) {
	db := setup.DB

	location = "%" + location + "%"
	personel = "%" + personel + "%"
	offset := (p.Page - 1) * p.Limit
	limit := p.Limit

	var cnt int64
	var query *gorm.DB
	var logs []models.AccessLog
	var data []template.AccessLogData = make([]template.AccessLogData, 0)

	query = db.Where(`
			(location LIKE ? OR 
			personel_name LIKE ? OR 
			personel_id_number LIKE ?) AND 
			timestamp BETWEEN ? AND ?
		`, location, personel, personel, p.StartDate, p.EndDate)

	if err := query.Find(&models.AccessLog{}).Count(&cnt).Error; err != nil {
		return nil, nil, err
	}

	if p.Limit < 0 {
		if err := query.
			Preload("Lock").Preload("Key").
			Order("created_at DESC").Find(&logs).Error; err != nil {
			return nil, nil, err
		}
	} else {
		if err := query.
			Limit(limit).Offset(offset).
			Preload("Lock").Preload("Key").
			Order("created_at DESC").Find(&logs).Error; err != nil {
			return nil, nil, err
		}
	}

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
	var queryString string
	var data []template.RSSILogData

	queryString = `
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
		rssi_logs.timestamp BETWEEN ? AND ? AND (
			personels.name LIKE ? OR 
			locks.label LIKE ? OR 
			keys.label LIKE ?
		)`

	if err := db.Raw(
		fmt.Sprintf("SELECT COUNT(*) AS cnt FROM ( %s ) AS t", queryString),
		p.StartDate, p.EndDate, keyword, keyword, keyword).
		Scan(&cnt).Error; err != nil {
		return nil, nil, err
	}

	if p.Limit < 0 {
		if err := db.Raw(
			fmt.Sprintf("%s ORDER BY rssi_logs.created_at DESC", queryString),
			p.StartDate, p.EndDate, keyword, keyword, keyword).
			Scan(&data).Error; err != nil {
			return nil, nil, err
		}
	} else {
		if err := db.Raw(
			fmt.Sprintf("%s ORDER BY rssi_logs.created_at DESC OFFSET ? LIMIT ?", queryString),
			p.StartDate, p.EndDate, keyword, keyword, keyword, offset, limit).
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

		if p.Limit < 0 {
			if err := db.Raw(
				fmt.Sprintf("%s ORDER BY healthcheck_logs.timestamp DESC", queryString),
				p.StartDate, p.EndDate, location, location).
				Scan(&data).Error; err != nil {
				return nil, nil, err
			}
		} else {
			if err := db.Raw(
				fmt.Sprintf("%s ORDER BY healthcheck_logs.timestamp DESC OFFSET ? LIMIT ?", queryString),
				p.StartDate, p.EndDate, location, location, offset, limit).
				Scan(&data).Error; err != nil {
				return nil, nil, err
			}
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

		if p.Limit < 0 {
			if err := db.Raw(
				fmt.Sprintf("%s ORDER BY healthcheck_logs.timestamp DESC", queryString),
				p.StartDate, p.EndDate, status == "active", location, location).
				Scan(&data).Error; err != nil {
				return nil, nil, err
			}
		} else {
			if err := db.Raw(
				fmt.Sprintf("%s ORDER BY healthcheck_logs.timestamp DESC OFFSET ? LIMIT ?", queryString),
				p.StartDate, p.EndDate, status == "active", location, location, offset, limit).
				Scan(&data).Error; err != nil {
				return nil, nil, err
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

	lock.Status = err == nil

	if err := db.Save(&lock).Error; err != nil {
		return nil, err
	}

	if err := db.Create(&status).Error; err != nil {
		return nil, err
	}

	return &status, nil
}

func MatchRSSILogEvent(data *models.RSSILog, keyword string) (*template.RSSILogData, error) {
	if strings.Contains(data.Personel.Name, keyword) ||
		strings.Contains(data.Lock.Label, keyword) ||
		strings.Contains(data.Key.Label, keyword) {
		return &template.RSSILogData{
			Timestamp:  data.Timestamp,
			PersonelID: data.PersonelID,
			Personel:   data.Personel.Name,
			LockID:     data.LockID,
			Lock:       data.Lock.Label,
			KeyID:      data.KeyID,
			Key:        data.Key.Label,
			Location:   data.Lock.Location,
			RSSI:       data.RSSI,
		}, nil
	}

	return nil, errors.New("this event doesnt match with the keyword")
}

func MatchAccessLogEvent(data *models.AccessLog, keyword string) (*template.AccessLogData, error) {
	if strings.Contains(data.PersonelName, keyword) ||
		strings.Contains(data.Lock.Label, keyword) ||
		strings.Contains(data.Key.Label, keyword) {
		return &template.AccessLogData{
			ID:         data.ID,
			Timestamp:  data.Timestamp,
			PersonelID: 0, // TODO : BENERIN JADI PAKE ID BENERAN
			Personel:   data.PersonelName,
			LockID:     data.LockID,
			Lock:       data.Lock.Label,
			KeyID:      data.KeyID,
			Key:        data.Key.Label,
			Location:   data.Lock.Location,
		}, nil
	}

	return nil, errors.New("this event doesnt match with the keyword")
}
