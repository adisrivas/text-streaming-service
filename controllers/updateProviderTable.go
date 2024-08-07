package controllers

import (
	"fmt"
	"strings"
	"text-streaming-service/db"
	"text-streaming-service/models"
)

func UpdateProviderTable(inputValues *models.UpdateProviderTableInput) {
	var provider string
	if inputValues.Provider == 1 {
		provider = "first"
	} else if inputValues.Provider == 2 {
		provider = "second"
	} else {
		provider = "third"
	}
	var placeholders string
	var params []interface{}
	for i := 0; i < 2; i++ {
		placeholders += "(?,?,?,?,?),"
		if i == 0 {
			params = append(params, i, inputValues.Start, inputValues.UserId, inputValues.RequestId, inputValues.IsAvailable)
		} else {
			params = append(params, i, inputValues.End, inputValues.UserId, inputValues.RequestId, inputValues.IsAvailable)
		}
	}

	placeholders = strings.TrimSuffix(placeholders, ",")

	_, err := db.Conn.Exec(fmt.Sprintf("INSERT INTO %s_provider (type,created_at,user_id,request_id,is_available) VALUES %s", provider, placeholders), params...)
	if err != nil {
		panic(err)
	}
}
