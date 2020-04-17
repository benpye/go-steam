// A simple example that uses the modules from the gsbot package and go-steam to log on
// to the Steam network.
package main

import (
	"flag"
	"fmt"

	"github.com/benpye/go-steam"
	"github.com/benpye/go-steam/gsbot"
	"github.com/benpye/go-steam/protocol/steamlang"
)

func main() {
	username := flag.String("username", "", "Account username")
	password := flag.String("password", "", "Account password")
	authcode := flag.String("authcode", "", "Auth code")
	twofactor := flag.String("twofactor", "", "Two factor code")

	flag.Parse()

	if *username == "" || *password == "" {
		flag.PrintDefaults()
		return
	}

	bot := gsbot.Default()
	client := bot.Client
	auth := gsbot.NewAuth(bot, &gsbot.LogOnDetails{
		Username:      *username,
		Password:      *password,
		AuthCode:      *authcode,
		TwoFactorCode: *twofactor,
	}, "sentry.bin")
	debug, err := gsbot.NewDebug(bot, "debug")
	if err != nil {
		panic(err)
	}
	client.RegisterPacketHandler(debug)
	serverList := gsbot.NewServerList(bot, "serverlist.json")
	serverList.Connect()

	for event := range client.Events() {
		auth.HandleEvent(event)
		debug.HandleEvent(event)
		serverList.HandleEvent(event)

		switch e := event.(type) {
		case error:
			fmt.Printf("Error: %v", e)
		case *steam.LoggedOnEvent:
			client.Social.SetPersonaState(steamlang.EPersonaState_Online)
		}
	}
}
