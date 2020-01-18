package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

var roomManager *Manager

func main() {
	roomManager = NewRoomManager()

	g := gin.Default()
	g.Static("static", "static/")
	g.SetHTMLTemplate(html)

	g.GET("/room/:room_id", roomGET)
	g.POST("/room/:room_id", roomPOST)
	g.DELETE("/room/:room_id", roomDELETE)
	g.GET("/stream/:room_id", stream)

	g.Run(":8080")
}

func roomGET(c *gin.Context) {
	rid := c.Param("room_id")
	userid := rand.Int31()
	c.HTML(http.StatusCreated, "chat_room", gin.H{
		"roomid": rid,
		"userid": userid,
	})
}

func roomPOST(c *gin.Context) {
	rid := c.Param("room_id")
	uid := c.PostForm("userid")
	m := c.PostForm("message")

	roomManager.Submit(uid, rid, m)

	c.JSON(http.StatusOK, gin.H{
		"message": m,
	})
}

func roomDELETE(c *gin.Context) {
	roomid := c.Param("room_id")
	roomManager.DeleteListener(roomid)
}

func stream(c *gin.Context) {
	roomid := c.Param("room_id")
	listener := roomManager.OpenListener(roomid)
	defer roomManager.CloseListener(roomid, listener)

	clientGone := c.Writer.CloseNotify()
	c.Stream(func(w io.Writer) bool {
		select {
		case <-clientGone:
			fmt.Println("Client Gone!")
			return false
		case m := <-listener:
			c.SSEvent("message", m)
			fmt.Println(m)
			return true
		}
	})
}
