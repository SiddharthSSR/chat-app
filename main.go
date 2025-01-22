package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/glebarez/go-sqlite"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Channel struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Message struct {
	ID        int    `json:"id"`
	ChannelID int    `json:"channel_id"`
	UserID    int    `json:"user_id"`
	UserName  string `json:"user_name"`
	Text      string `json:"text"`
}

func main() {
	// Get the working directory
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	// Print the working directory
	fmt.Println("Working Directory:", wd)

	// Open the SQLite database file
	db, err := sql.Open("sqlite", wd+"/database.db")

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)

	// Create the Gin router
	r := gin.Default()

	if err != nil {
		log.Fatal(err)
	}

	// Creation endpoints
	r.POST("/users", func(c *gin.Context) { createUser(c, db) })
	r.POST("/channels", func(c *gin.Context) { createChannel(c, db) })
	r.POST("/messages", func(c *gin.Context) { createMessage(c, db) })

	// Listing endpoints
	r.GET("/channels", func(c *gin.Context) { listChannels(c, db) })
	r.GET("/messages", func(c *gin.Context) { listMessages(c, db) })

	// Login endpoint
	r.POST("/login", func(c *gin.Context) { login(c, db) })

	err = r.RUN(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
