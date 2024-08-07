package controllers

import "text-streaming-service/db"

func AddRequest(provider int, userId int) int64 {
	result, err := db.Conn.Exec("INSERT INTO requests (provider, user_id) VALUES (?,?)", provider, userId)
	if err != nil {
		panic(err)
	}

	return result.LastInsertId()
}
