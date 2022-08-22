package main

import (
	"context"
	"fmt"
	"net/http"

	firebase "firebase.google.com/go"
	"github.com/labstack/echo"
	"google.golang.org/api/option"
)

func main() {

	ctx := context.Background()
	opt := option.WithCredentialsFile("path/to/serviceAccountKey.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		fmt.Println(err)
	}

	client, err := app.Firestore(ctx)

	if err != nil {
		fmt.Println(err)
	}

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World...!")
	})

	type Memo struct {
		Name     string `json:"name"`
		Birthday string `json:"birthday"`
	}

	e.POST("/addmemo", func(c echo.Context) error {
		param := new(Memo)
		if err := c.Bind(param); err != nil {
			return err
		}

		_, _, err = client.Collection("Users").Add(ctx, map[string]interface{}{
			"name":     param.Name,
			"birthday": param.Birthday,
		})
		if err != nil {
			fmt.Println(err)
		}
		p := &Memo{
			Name:     "AC_TLE",
			Birthday: "1/22",
		}
		return c.JSON(http.StatusOK, p)
	})
	e.Logger.Fatal(e.Start(":8080"))
}
