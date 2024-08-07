package controllers

import "text-streaming-service/db"

func AddRequest(provider int, userId int, prompt string) int64 {
	result, err := db.Conn.Exec("INSERT INTO requests (provider, user_id, prompt) VALUES (?,?,?)", provider, userId, prompt)
	if err != nil {
		panic(err)
	}
	requestId, _ := result.LastInsertId()

	return requestId
}
