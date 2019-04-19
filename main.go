package main

import (
	"fmt"
	"time"

	"github.com/Gitter-Robot/config"
	"github.com/Gitter-Robot/robot"
	gitter "github.com/go-gitter"
)

const (
	chatRoom = "MultiVAC/MTV-Dev"
)

func main() {

	// Load config
	cfg := config.LoadConfig()

	roomMap := make(map[string]gitter.Room)
	api := gitter.New(cfg.GitterToken)
	user, err := api.GetUser()
	if err != nil {
		fmt.Println(err)
	}

	rooms, err := api.GetRooms()
	for _, v := range rooms {
		fmt.Println(v.Name)
		roomMap[v.Name] = v
	}
	multiVAC := roomMap[chatRoom].ID
	stream := api.Stream(multiVAC)
	go api.Listen(stream)

	// Roboat instance
	rob := robot.NewRobot("qingyun")

	for {
		event := <-stream.Event
		switch ev := event.Data.(type) {
		case *gitter.MessageReceived:
			// log
			fmt.Println(ev.Message.From.Username + ": " + ev.Message.Text)
			// Ask turing
			res, err := rob.Ask(ev.Message.Text)
			if err != nil {
				fmt.Println("Ask error: ", err)
			}

			// Other people
			if ev.Message.From.ID != user.ID {
				api.SendMessage(multiVAC, res)
			} else {
				api.SendMessage(multiVAC, "我的主人说："+ev.Message.Text)
			}
			time.Sleep(time.Second * 3)
		case *gitter.GitterConnectionClosed:
			// connection was closed
		}
	}

}
