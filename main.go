package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo"

	firebase "firebase.google.com/go"
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

	e.GET("/addmemo", func(c echo.Context) error {
		_, _, err = client.Collection("Users").Add(ctx, map[string]interface{}{
			"name":  "First User",
			"age":   11,
			"email": "first@example.com",
		})
		if err != nil {
			fmt.Println(err)
		}
		return c.String(http.StatusOK, "Hello World...!")
	})

	type Room struct {
		ID string `json:"id"`
	}

	e.GET("/createroom", func(c echo.Context) error {
		ref := client.Collection("Users").NewDoc()

		_, err := ref.Set(ctx, map[string]interface{}{})

		if err != nil {
			fmt.Println("An error has occurred: %s", err)
		}

		p := &Room{
			ID: ref.ID,
		}

		return c.JSON(http.StatusOK, p)
	})
	e.Logger.Fatal(e.Start(":8080"))
}
