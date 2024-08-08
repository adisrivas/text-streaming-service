package controllers

import (
	"text-streaming-service/db"
)

func UpdateRequest(provider int, requestId int) {
	_, err := db.Conn.Exec("UPDATE requests SET provider=? WHERE id=?", provider, requestId)
	if err != nil {
		panic(err)
	}
}
