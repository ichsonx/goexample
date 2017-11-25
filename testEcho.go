package main

import (
	echo2 "github.com/labstack/echo"
	"net/http"
)

func main() {
	e := echo2.New()
	e.GET("/", func(context echo2.Context) error {
		return context.String(http.StatusOK, "hello echo")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
