package main

import (
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Song struct {
	Name     string
	Duration time.Duration
}

type node struct {
	song *Song
	next *node
	prev *node
}

type Playlist struct {
	head    *node
	tail    *node
	current *node
	mutex   sync.Mutex
}

func main() {
	router := gin.Default()

	router.GET("/songs", getSongs)
	router.GET("/songs/:id", getSong)
	router.POST("/songs", createSong)
	router.PUT("/songs/:id", updateSong)
	router.DELETE("/songs/:id", deleteSong)
	router.POST("/songs/play", playSong)
	router.POST("/songs/pause", pauseSong)
	router.POST("/songs/next", nextSong)
	router.POST("/songs/previous", previousSong)

	router.Run(":8080")
}

func getSongs(c *gin.Context) {
	var songs []Song
	db.Find(&songs)

	c.JSON(200, songs)
}

func getSong(c *gin.Context) {
	var song Song
	id := c.Param("id")
	db.First(&song, id)

	if song.ID == 0 {
		c.JSON(404, gin.H{"message": "Song not found"})
		return
	}

	c.JSON(200, song)
}

func createSong(c *gin.Context) {
	var song Song
	c.BindJSON(&song)

	db.Create(&song)

	c.JSON(201, song)
}

func updateSong(c *gin.Context) {
	var song Song
	id := c.Param("id")
	db.First(&song, id)

	if song.ID == 0 {
		c.JSON(404, gin.H{"message": "Song not found"})
		return
	}

	c.BindJSON(&song)
	db.Save(&song)

	c.JSON(200, song)
}

func deleteSong(c *gin.Context) {
	var song Song
	id := c.Param("id")
	db.First(&song, id)

	if song.ID == 0 {
		c.JSON(404, gin.H{"message": "Song not found"})
		return
	}

	db.Delete(&song)

	c.JSON(200, gin.H{"message": "Song deleted"})
}

func (p *Playlist) playSong() {
    p.mutex.Lock()
    defer p.mutex.Unlock()

    if p.current != nil {
        return
    }

    p.current = p.head
    for p.current != nil {
        go func(current *node) {
            time.Sleep(current.song.Duration)
            if current == p.current {
                p.current = p.current.next
                p.playSong()
            }
        }(p.current)

        return
    }
}

func pauseSong(c *gin.Context) {
	// Implement logic for pausing a song
}

func nextSong(c *gin.Context) {

}

func previousSong(c *gin.Context) {

}
