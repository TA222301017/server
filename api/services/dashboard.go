package services

import (
	"server/models"
	"server/setup"
)

func DashboardData() (keyCnt int64, lockCnt int64, personelCnt int64) {
	db := setup.DB

	db.Find(&models.Key{}).Count(&keyCnt)
	db.Find(&models.Lock{}).Count(&lockCnt)
	db.Find(&models.Personel{}).Count(&personelCnt)

	return
}
