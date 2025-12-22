package controller

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Otavio-Fina/live-websocket/models"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"golang.org/x/exp/slices"
)

func HandlePollMsg(pollID string, selectedOption string, user string) (success bool) {

	poll := models.Pools[pollID]

	if !slices.Contains(poll.Options, selectedOption) {
		println("no option on poll")
		return
	}

	models.Pools[pollID].Votes[user] = selectedOption

	return true
}

func AddConnectionToPoll(conn *websocket.Conn, pollID string) {
	if !slices.Contains(models.Pools[pollID].Connections, conn) {
		models.Pools[pollID].Connections = append(models.Pools[pollID].Connections, conn)
	}
}

func CreatePoll(d *models.PostPollContract) *models.Poll {
	idPoll := uuid.NewString()

	poll := &models.Poll{
		ID:          idPoll,
		Name:        d.Name,
		Question:    d.Question,
		Options:     d.Options,
		Votes:       make(map[string]string),
		Connections: []*websocket.Conn{},
	}

	models.Pools[idPoll] = poll

	return poll
}

func GetPollByID(pollID string) (*models.Poll, bool) {
	poll, exists := models.Pools[pollID]
	return poll, exists
}

func StartConnection(conn *websocket.Conn, mt int, pollID string) {

	err := conn.WriteMessage(mt, []byte("conectado"))
	if err != nil {
		log.Println("writeMsg error: ", err)
	} else {
		AddConnectionToPoll(conn, pollID)
	}
}

func Vote(pollID string, vote string, userIDStr string, mt int, conn *websocket.Conn) {
	success := HandlePollMsg(pollID, vote, userIDStr)

	pollJson, _ := json.Marshal(models.Pools[pollID])
	if success {
		for _, connI := range models.Pools[pollID].Connections {
			if connI == nil {
				continue
			}

			responseMsg := []byte(fmt.Sprintf("200: %s", string(pollJson)))
			err := connI.WriteMessage(mt, responseMsg)
			if err != nil {
				log.Println("writeMsg error: ", err)
				models.Pools[pollID].Connections = slices.DeleteFunc(models.Pools[pollID].Connections, func(c *websocket.Conn) bool { return c == connI })
			}
		}
	} else {
		responseMsg := []byte(fmt.Sprintf("400: %s", string(pollJson)))
		err := conn.WriteMessage(mt, responseMsg)
		if err != nil {
			log.Println("writeMsg error: ", err)
			models.Pools[pollID].Connections = slices.DeleteFunc(models.Pools[pollID].Connections, func(c *websocket.Conn) bool { return c == conn })
		}
	}
}

func ChangeOptions(pollID string, changeOptionsParams map[string]string, mt int, conn *websocket.Conn) {
	poll, exists := GetPollByID(pollID)

	if !exists {
		responseMsg := []byte("400: poll not found")
		err := conn.WriteMessage(mt, responseMsg)
		if err != nil {
			log.Println("writeMsg error: ", err)
		}
		return
	}

	for option, AddOrDel := range changeOptionsParams {
		switch AddOrDel {
		case "add":
			poll.Options = append(poll.Options, option)
		case "del":
			poll.Options = slices.DeleteFunc(poll.Options, func(s string) bool { return s == option })
		}
	}

	pollJson, err := json.Marshal(poll)
	if err != nil {
		log.Println("json marshal error: ", err)
	}
	for _, connI := range models.Pools[pollID].Connections {
		if connI == nil {
			continue
		}

		responseMsg := []byte(fmt.Sprintf("200: %s", string(pollJson)))
		err := connI.WriteMessage(mt, responseMsg)
		if err != nil {
			log.Println("writeMsg error: ", err)
			models.Pools[pollID].Connections = slices.DeleteFunc(models.Pools[pollID].Connections, func(c *websocket.Conn) bool { return c == connI })
		}
	}
}
