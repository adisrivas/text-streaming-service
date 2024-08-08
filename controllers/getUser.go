package controllers

import (
	"text-streaming-service/db"
	"text-streaming-service/models"
)

func GetUser() *models.User {
	result, err := db.Conn.Query("SELECT id, name, email, plan_id FROM users ORDER BY id DESC LIMIT 1")
	if err != nil {
		panic(err)
	}
	defer result.Close()

	var user *models.User = new(models.User)
	for result.Next() {
		err := result.Scan(&user.Id, &user.Name, &user.Email, &user.PlanId)
		if err != nil {
			panic(err)
		}
		user.HideId()
	}

	return user
}
