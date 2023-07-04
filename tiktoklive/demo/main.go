package main

import (
	"fmt"

	"github.com/Davincible/gotiktoklive"
)

var username = "honekrasae_official"

func main() {
	tiktok := gotiktoklive.NewTikTok()
	live, err := tiktok.TrackUser(username)
	if err != nil {
		panic(err)
	}

	if err := live.DownloadStream(); err != nil {
		panic(err)
	}

	for event := range live.Events {
		switch e := event.(type) {
		case gotiktoklive.UserEvent:
			fmt.Printf("%T: %s %s\n", e, e.Event, e.User.Username)
		case gotiktoklive.ViewersEvent:
			fmt.Printf("%T: %d\n", e, e.Viewers)
		default:
			fmt.Printf("%T: %+v\n", e, e)
		}
	}
}
