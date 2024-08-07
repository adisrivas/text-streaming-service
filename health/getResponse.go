package health

import (
	"errors"
	"fmt"
	"text-streaming-service/controllers"
	"text-streaming-service/err"
	"text-streaming-service/models"
	"text-streaming-service/stubs"
	"time"
)

func GetResponse() {
	var defaultQuestion string = "What is the capital of India"
	for i := 0; i < 3; i++ {

		requestId := controllers.AddRequest(i+1, 0, defaultQuestion)

		var isAvailable int = 1

		var statusCode int
		var err1 error

		start := time.Now().Unix()
		if i == 0 {
			statusCode, _, err1 = stubs.FirstProvider(defaultQuestion)
		} else if i == 1 {
			statusCode, _, err1 = stubs.SecondProvider(defaultQuestion)
		} else {
			statusCode, _, err1 = stubs.ThirdProvider(defaultQuestion)
		}
		end := time.Now().Unix()

		var e error
		if err1 != nil || statusCode != 200 || end-start > 5 {
			isAvailable = 0
			e = err1
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
			Provider:    i + 1,
			Start:       int(start),
			End:         int(end),
			UserId:      0,
			RequestId:   int(requestId),
			IsAvailable: isAvailable,
		})
	}
}
