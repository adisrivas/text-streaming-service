package query

import (
	"errors"
	"fmt"
	"sync"
	"text-streaming-service/controllers"
	"text-streaming-service/db"
	"text-streaming-service/err"
	"text-streaming-service/models"
	"text-streaming-service/stubs"
	"time"

	"github.com/labstack/echo/v4"
)

func GetResponse(c echo.Context) error {
	user := controllers.GetUser() // user information can be retrived through cookies later
	userId := user.Id
	prompt := c.QueryParam("prompt")
	if len(prompt) == 0 {
		c.JSON(200, map[string]string{
			"message": "please provide valid prompt",
			"data":    "",
		})
		return nil
	}

	var availableProviders map[int]bool = make(map[int]bool)
	var responseTime map[int]int = make(map[int]int)
	var wg sync.WaitGroup
	var mutex = &sync.RWMutex{}

	for i := 1; i < 4; i++ {
		wg.Add(1)
		go func(i int) {
			var providerStr string
			if i == 1 {
				providerStr = "first"
			} else if i == 2 {
				providerStr = "second"
			} else {
				providerStr = "third"
			}
			result, err1 := db.Conn.Query(fmt.Sprintf("SELECT created_at,is_available FROM %s_provider ORDER BY id DESC LIMIT 2", providerStr))
			if err1 != nil {
				panic(err1)
			}
			var end bool = true
			var responseTimeInt int = 0
			for result.Next() {
				var createdAt int
				var isAvailable int
				err1 = result.Scan(&createdAt, &isAvailable)
				if err1 != nil {
					panic(err1)
				}

				if isAvailable > 0 {
					if end {
						responseTimeInt = createdAt
						end = false
					} else {
						responseTimeInt -= createdAt
					}
				}
				if _, ok := availableProviders[i]; !ok {
					mutex.Lock()
					availableProviders[i] = isAvailable > 0
					mutex.Unlock()
				}
			}
			mutex.Lock()
			responseTime[i] = responseTimeInt
			mutex.Unlock()
			wg.Done()
		}(i)
	}
	wg.Wait()

	var responseData string
	var isProcessed bool = false
	for provider, isAvailable := range availableProviders {
		var newAvailabilityStatus int = 1
		if isAvailable {
			var start int64
			var end int64
			requestId := controllers.AddRequest(provider, user.Id, prompt)
			if responseTime[provider] <= 5 {
				start = time.Now().Unix()
				var statusCode int
				var response string
				var err1 error
				if provider == 1 {
					statusCode, response, err1 = stubs.FirstProvider(prompt)
				} else if provider == 2 {
					statusCode, response, err1 = stubs.SecondProvider(prompt)
				} else {
					statusCode, response, err1 = stubs.ThirdProvider(prompt)
				}
				end = time.Now().Unix()
				if err1 != nil {
					err.Log(err1, int(requestId))
					newAvailabilityStatus = 0
				}
				if statusCode != 200 {
					err.Log(errors.New(fmt.Sprintf("some error in response. status code- %d", statusCode)), int(requestId))
					newAvailabilityStatus = 0
				}
				responseData = response
				isProcessed = true
			} else {
				err.Log(errors.New("experiencing delay"), int(requestId))
				newAvailabilityStatus = 0
			}
			controllers.UpdateProviderTable(&models.UpdateProviderTableInput{
				Provider:    provider,
				Start:       int(start),
				End:         int(end),
				RequestId:   int(requestId),
				UserId:      userId,
				IsAvailable: newAvailabilityStatus,
			})
			if isProcessed {
				break
			}
		}
	}

	if isProcessed {
		c.JSON(200, map[string]interface{}{
			"message": "generated response",
			"data":    responseData,
		})
	} else {
		c.JSON(204, map[string]interface{}{
			"message": "something went wrong",
			"data":    "",
		})
	}
	return nil
}
