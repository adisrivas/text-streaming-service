package health

import (
	"fmt"
	"strings"

	"text-streaming-service/db"
)

func Log(providerInt int, start int64, end int64, requestId int) {
	var provider string
	if providerInt == 1 {
		provider = "first"
	} else if providerInt == 2 {
		provider = "second"
	} else {
		provider = "third"
	}

	var placeholders string
	var params []interface{}
	for i := 0; i < 2; i++ {
		placeholders += "(?,?,?,?),"
	}

	params = append(params, 0, int(start), requestId, 0, 1, int(end), requestId, 0)

	placeholders = strings.TrimSuffix(placeholders, ",")

	_, err := db.Conn.Exec(fmt.Sprintf("INSERT INTO %s_provider (type,created_at,request_id,user_id) VALUES %s", provider, placeholders), params...)
	if err != nil {
		panic(err)
	}
}
