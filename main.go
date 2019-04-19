package main

import (
	"encoding/json"
	"fmt"
	"time"

	gitter "github.com/go-gitter"
)

const (
	token    = ""
	multivac = "MultiVAC/MTV-Dev"
	robaot   = "http://openapi.tuling123.com/openapi/api/v2"
)

func main() {

	roomMap := make(map[string]gitter.Room)
	api := gitter.New(token)
	user, err := api.GetUser()
	if err != nil {
		fmt.Println(err)
	}

	rooms, err := api.GetRooms()
	for _, v := range rooms {
		fmt.Println(v.Name)
		roomMap[v.Name] = v
	}
	multiVAC := roomMap[multivac].ID
	stream := api.Stream(multiVAC)
	go api.Listen(stream)

	// Roboat instance
	turing := NewTuringRoboat()

	for {
		event := <-stream.Event
		switch ev := event.Data.(type) {
		case *gitter.MessageReceived:
			// log
			fmt.Println(ev.Message.From.Username + ": " + ev.Message.Text)
			// Ask turing
			res, err := turing.AskQingYun(ev.Message.Text)
			if err != nil {
				fmt.Println("Ask error: ", err)
			}
			result := &QingYunResult{}
			json.Unmarshal([]byte(res), &result)
			// Other people
			if ev.Message.From.ID != user.ID {
				api.SendMessage(multiVAC, result.Content)
			} else {
				api.SendMessage(multiVAC, "我的主人说："+ev.Message.Text)
			}
			time.Sleep(time.Second * 3)
		case *gitter.GitterConnectionClosed:
			// connection was closed
		}
	}

}
