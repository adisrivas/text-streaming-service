package health

import (
	"net/http"
	"text-streaming-service/db"
	"time"
)

var URLs []string = []string{
	"http://localhost:8000/1/health",
	"http://localhost:8000/2/health",
	"http://localhost:8000/3/health",
}

func GetResponse() {
	for i, url := range URLs {
		result, err := db.Conn.Exec("INSERT INTO requests (provider, user_id) VALUES (?,?)", i+1, 0)
		if err != nil {
			panic(err)
		}
		requestId, _ := result.LastInsertId()
		start := time.Now().Unix()
		resp, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		end := time.Now().Unix()
		Log(i+1, start, end, int(requestId))
	}
}
