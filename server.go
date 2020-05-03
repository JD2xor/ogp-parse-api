package main

import (
	"fmt"
	"net/http"
	"os"
	
	"github.com/labstack/echo"
	"github.com/otiai10/opengraph"
)

func main() {
	e := echo.New()
	e.GET("/ogp", getOgp)

	port := os.Getenv("PORT")
	if port == "" {
		port = "1323"
	}
	e.Logger.Fatal(e.Start(":" + port))
}

func getOgp(ctx echo.Context) error {
	url := ctx.QueryParam("url")
	if len(url) == 0 {
		return ctx.String(http.StatusBadRequest, "Url is empty")
	}

	og, err := opengraph.Fetch(url)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, fmt.Sprintf("Error: %s", err.Error()))
	}
	
	og = og.ToAbsURL()

	return ctx.JSON(http.StatusOK, og)	
}
