package main

import (
	"text-streaming-service/db"
	health "text-streaming-service/health"
	query "text-streaming-service/query"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {
	db.Init()
	server := echo.New()

	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()
	go func() {
		for {
			select {
			case <-ticker.C:
				{
					health.GetResponse()
				}
			}
		}
	}()

	server.GET("/query", query.GetResponse)
	server.GET("/1/health", health.GetFirstProvider)
	server.GET("/2/health", health.GetSecondProvider)
	server.GET("/3/health", health.GetThirdProvider)

	server.RouteNotFound("/*", func(c echo.Context) error {
		return c.JSON(404, map[string]string{
			"error": "endpoint does not exist",
		})
	})

	server.Logger.Fatal(server.Start(":8000"))

}
