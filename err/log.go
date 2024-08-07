package err

import "text-streaming-service/db"

func Log(err error, requestId int) {
	errorStr := err.Error()
	db.Conn.Exec("INSERT INTO error_log (error, request_id) VALUES (?,?)", errorStr, requestId)
}
