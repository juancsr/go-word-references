package uter

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "github.com/juancsr/go-word-references/managers/wordManager"
)


var router *gin.Engine

func initRouting() {
    router = gin.Default()
    wordRoute()
    userRoute()
}

func wordRoute() {
    router.GET("/word/:id", func(c *gin.Context) {
        id := c.Param("id")
        print("router id: ", id,"\n")
        c.JSON(200, wordManager.GetOne(id))
    })
}

func userRoute() {
    router.GET("/use/:name", func(c *gin.Context) {
        name := c.Param(":name")
		c.String(http.StatusOK, "Hello %s", name)
	})
}

func Run() {
    initRouting()
    router.Run(":3001")
}


