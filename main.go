package main

import (
	"fmt"
  "github.com/labstack/echo/v4"
  "github.com/labstack/echo/v4/middleware"
  "net/http"
)

func main() {
  // Echo instance
  e := echo.New()

  // Middleware
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  // Routes
  e.GET("/", hello)
  e.GET("/users/:id", getUser)

  // Start server
  e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
  return c.String(http.StatusOK, "Hello, World!")
}

func getUser(c echo.Context) error {
	for key, values := range c.Request().Header {
		fmt.Println(key)
		for _,value := range values {
			fmt.Println(value)
		}
	}
	// User ID from path `users/:id`
	print(c.Request(),"\n")
	id := c.Param("id")
  return c.String(http.StatusOK, id)
}
