package controllers

import (
	"fmt"
	"text-streaming-service/db"
	"text-streaming-service/models"
	"time"
)

func IsRateLimitExceeded(user *models.User) bool {
	res, err := db.Conn.Query("SELECT plan_limit FROM plans WHERE id=?", user.PlanId)
	if err != nil {
		panic(err)
	}

	var limit int
	for res.Next() {
		err = res.Scan(&limit)
		if err != nil {
			panic(err)
		}
	}
	currentDT := time.Now().UTC()
	currentDate := currentDT.Format("2006-01-02")
	startDate := fmt.Sprintf("%s 00:00:00", currentDate)
	endDate := fmt.Sprintf("%s 23:59:59", currentDate)

	res, err = db.Conn.Query("SELECT COUNT(id) AS request_count FROM requests WHERE user_id=? AND created_at>=? AND created_at<=?", user.Id, startDate, endDate)
	if err != nil {
		panic(err)
	}
	var requestCount int
	for res.Next() {
		err = res.Scan(&requestCount)
		if err != nil {
			panic(err)
		}
	}

	return requestCount < limit
}
