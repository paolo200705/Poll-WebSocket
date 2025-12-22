package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Otavio-Fina/live-websocket/controller"
	"github.com/Otavio-Fina/live-websocket/models"

	"github.com/gin-gonic/gin"
)

func PostPoll(c *gin.Context) {
	var data models.PostPollContract
	c.BindJSON(&data)

	newPoll := controller.CreatePoll(&data)

	c.JSON(http.StatusOK, newPoll)
}

func GetAllPolls(c *gin.Context) {
	c.JSON(http.StatusOK, models.Pools)
}

func GetPoll(c *gin.Context) {
	pollID, err := c.GetQuery("pollID")

	if err {
		c.JSON(http.StatusBadRequest, gin.H{"error": "pollID is required"})
		return
	}

	poll, exist := controller.GetPollByID(pollID)

	if !exist {
		c.JSON(http.StatusNotFound, gin.H{"error": "Poll not found"})
		return
	}
	c.JSON(http.StatusOK, poll)
}

func ConectPoll(c *gin.Context) {
	pollID := c.Param("pollID")

	if _, exists := controller.GetPollByID(pollID); !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Poll not found"})
	}

	conn, err := models.WsUpgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("upgrade", err)
	}

	controller.AddConnectionToPoll(conn, pollID)
	defer conn.Close()

	for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("readMsg", err)
			break
		}

		var wsMsg models.WSContract
		err = json.Unmarshal(message, &wsMsg)
		if err != nil {
			log.Println("unmarshal msg ws", err)
			break
		}

		userID, _ := c.Get("user_id")
		userIDStr, ok := userID.(string)
		if !ok {
			log.Println("erro ao converter user_id para string")
			break
		}

		switch wsMsg.MesaggeType {
		case 1:
			controller.StartConnection(conn, mt, pollID)
			continue
		case 2:
			controller.Vote(pollID, wsMsg.Vote, userIDStr, mt, conn)
			continue
		case 3:
			controller.ChangeOptions(pollID, wsMsg.ChangeOptionsParams, mt, conn)
			continue
		default:
			log.Println("tipo de mensagem desconhecido", wsMsg.MesaggeType)
		}

	}
}
