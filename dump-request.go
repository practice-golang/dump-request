package main // import "dump-request"

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Response : Response set
type Response struct {
	Status  int
	Message string
}

func bodyDumpHandler(c echo.Context, reqBody, resBody []byte) {
	fmt.Printf("Request Body: %v\n", string(reqBody))
	fmt.Printf("Response Body: %v\n", string(resBody))
}

func main() {
	e := echo.New()

	e.Use(middleware.BodyDump(bodyDumpHandler))

	e.POST("/yay", func(c echo.Context) error {
		return c.JSON(http.StatusOK, Response{
			Status:  http.StatusOK,
			Message: "Yay!",
		})
	})

	e.Logger.Fatal(e.Start("127.0.0.1:1323"))
}
