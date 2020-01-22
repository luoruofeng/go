package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"math/rand"
	"net/http"
	"runtime"
)

var roomManager *Manager

func ConfigRuntime() {
	nuCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nuCPU)
	fmt.Printf("Running with %d CPUs\n", nuCPU)
}

func main() {
	ConfigRuntime()

	roomManager = NewRoomManager()

	g := gin.Default()
	g.Static("static", "static/")

	g.LoadHTMLGlob("./tmp/*")

	g.GET("/room/:room_id", roomGET)
	g.GET("/index", indexGET)
	g.POST("/room/:room_id", roomPOST)
	g.DELETE("/room/:room_id", roomDELETE)
	g.GET("/stream/:room_id", stream)

	g.Run(":8080")
}

func roomGET(c *gin.Context) {
	rid := c.Param("room_id")
	userid := rand.Int31()
	c.HTML(http.StatusCreated, "index.tmpl", gin.H{
		"roomid": rid,
		"userid": userid,
	})
}

func indexGET(c *gin.Context) {
	c.HTML(http.StatusCreated, "index.tmpl", gin.H{})
}

func roomPOST(c *gin.Context) {
	rid := c.Param("room_id")
	uid := c.PostForm("user")
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
