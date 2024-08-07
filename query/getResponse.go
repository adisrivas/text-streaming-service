package query

import (
	"db"
	"text-streaming-service/controllers"
	"time"

	"github.com/labstack/echo/v4"
)

func GetResponse(c echo.Context) error {
	user := controllers.GetUser() // user information can be retrived through cookies later
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
	var wg = &sync.WaitGroup
	var mutex &sync.RWMutex{} 

	for i:=1; i<4; i++ {
		wg.Add(1)
		go func(i int) {
			var providerstr string
			if i == 1 {
				providerStr = "first"
			} else if i ==2 {
				providerStr = "second"
			} else {
				providerStr = "third"
			}
			result, err := db.Conn.Query(fmt.Sprintf("SELECT created_at,is_available FROM %s_provider ORDER BY id DESC LIMIT 2", providerStr))
			if err != nil {
				panic(err)
			}
			var end bool = true
			var responseTimeInt int = 0
			for result.Next() {
				var createdAt int
				var isAvailable int
				err = result.Scan(&createdAt, &isAvailable)
				if err != nil {
					panic(err)
				}
	
				if isAvailable {
					if end {
						responseTimeInt = createdAt
						end = false
					} else {
						responseTimeInt -= createdAt
					}
				}
			}
			mutex.Lock()
			availableProviders[i] = isAvailable > 0
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
			var providerFunc
			if responseTime <= 5 {
				requestId := controllers.AddRequest(provider, user.Id)
				start := time.Now().Unix()
				var statusCode int
				var response string
				var err error
				if provider == 1 {
					statusCode, response, err = stubs.FirstProvider(prompt)
				} else if provider == 2 {
					statusCode, response, err = stubs.SecondProvider(prompt)
				} else {
					statusCode, response, err = stubs.ThirdProvider(prompt)
				}
				end := time.Now().Unix()
				if err != nil {
					err.Log(err, requestId)
					newAvailabilityStatus = 0
				}
				if statusCode != 200 {
					err.Log(errors.New(fmt.Sprintf("some error in response. status code- %d", statusCode)), requestId)
					newAvailabilityStatus = 0
				}
				responseData = response
				isProcessed = true
				break
			} else {
				err.Log(errors.New("experiencing delay"), requestId)
				newAvailabilityStatus = 0
			}
			controllers.updateProviderTable(&models.UpdateProviderTableInput{
				Provider: provider,
				Start: int(start),
				End: int(end),
				RequestId: int(requestId),
				UserId: userId,
				isAvailable: newAvailabilityStatus
			})
		}
	}

	if isProcessed {
		c.JSON(200, map[string]interface{}{
			"message": "generated response",
			"data":    responseData,
		})
	} else {
		c.JSON(204, map[string]interface{}{
			"message": "text inference providing services are experiencing delay",
			"data":    "",
		})
	}
	return nil
}
