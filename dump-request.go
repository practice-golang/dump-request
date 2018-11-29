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

func dumpHandler(c echo.Context, reqBody, resBody []byte) {
	fmt.Println("Request Header:")
	for name, headers := range c.Request().Header {
		// name = strings.ToLower(name)
		for _, header := range headers {
			fmt.Println(name + " = " + header)
		}
	}

	fmt.Println("")
	fmt.Printf("Request Body:\n%v\n", string(reqBody))

	fmt.Println("")
	fmt.Printf("Response Body:\n%v\n", string(resBody))
}

func main() {
	e := echo.New()

	e.Use(middleware.BodyDump(dumpHandler))

	e.POST("/yay", func(c echo.Context) error {
		return c.JSON(http.StatusOK, Response{
			Status:  http.StatusOK,
			Message: "Yay!",
		})
	})

	e.Logger.Fatal(e.Start("127.0.0.1:1323"))
}
