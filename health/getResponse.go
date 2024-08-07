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
	var defaultQuestion string = "What is the capital of India"
	for i, url := range URLs {
		result, err := db.Conn.Exec("INSERT INTO requests (provider, user_id) VALUES (?,?)", i+1, 0)
		if err != nil {
			panic(err)
		}
		requestId, _ := result.LastInsertId()
		
		var IsAvailable int = 1
		
		var statusCode int
		var response string
		var err error
		
		start := time.Now().Unix()
		if i == 0 {
			statusCode, response, err = stubs.FirstProvider(defaultQuestion)
		} else if i == 1 {
			statusCode, response, err = stubs.SecondProvider(defaultQuestion)
		} else {
			statusCode, response, err = stubs.ThirdProvider(defaultQuestion)
		}
		end := time.Now().Unix()

		var e error
		if err != nil || statusCode != 200 || end-start > 5 {
			isAvailable = 0
			e = err
		}
		if statusCode != 200 {
			e = errors.New(fmt.Sprintf("some error in response. status code- %d", statusCode))
		}
		if end-start > 5 {
			e = errors.New("experiencing delay")
		}

		if e != nil {
			err.Log(e, int(requestId))
		}
		
		controllers.UpdateProviderTable(&models.UpdateProviderTableInput{
			Provider: i+1,
			Start: start,
			End: end,
			UserId: 0,
			RequestId: int(requestId),
			IsAvailable: isAvailable
		})
	}
}
