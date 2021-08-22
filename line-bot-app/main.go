package main

import (
    "log"
    //"net/http"
	//echo "github.com/labstack/echo/v4"
	//"github.com/labstack/echo/v4/middleware"
    "github.com/line/line-bot-sdk-go/v7/linebot"
)

func main() {
    // TODO: チャネルシークレットとチャネルトークンを追加
    bot, err := linebot.New("<channel secret>", "<channel access token>")
    if err != nil {
        log.Fatal(err)
    }
    message := linebot.NewTextMessage("hello, world")
    if _, err := bot.BroadcastMessage(message).Do(); err != nil {
        log.Fatal(err)
    }



	// e := echo.New()
    //
	// e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	// 	Format: `"time":"${time_rfc3339}","remote_ip":"${remote_ip}","host":"${host}",` +
	// 		`"method":"${method}","uri":"${uri}","status":${status},"error":"${error}"` + "\n",
	// }))
	// e.Use(middleware.Recover())
    //
	// e.GET("/", func(c echo.Context) error {
    //     return c.String(http.StatusOK, "hello world")
	// })
    //
	// e.Logger.Fatal(e.Start(":1234"))
}
