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

    e.GET("/add", func(c echo.Context) error {
        _, _, err = client.Collection("Users").Add(ctx, map[string]interface{}{
            "name": "First User",
            "age": 11,
            "email": "first@example.com",
        })
        if err != nil {
            fmt.Println(err)
        }
        return c.String(http.StatusOK, "Hello World...!")
    })
    e.Logger.Fatal(e.Start(":8080"))
}
