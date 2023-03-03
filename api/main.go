package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Song struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Length int    `json:"lenth"`
}

var db *sql.DB

func init() {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres sslmode=disable password=goLANGn1nja")
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
}

func main() {

	http.HandleFunc("/songs", getSongs)
	http.HandleFunc("/songs/:id", getSong)
	// http.HandleFunc("/songs", createSong)
	// http.HandleFunc("/songs/:id", updateSong)
	// http.HandleFunc("/songs/:id", deleteSong)


	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

	// router := gin.Default()

	// router.GET("/songs", getSongs)
	// router.GET("/songs/:id", getSong)
	// router.POST("/songs", createSong)
	// router.PUT("/songs/:id", updateSong)
	// router.DELETE("/songs/:id", deleteSong)

	// router.Run(":8080")
}

func getSongs(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("select * from playlist")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	songs := make([]Song, 0)

	for rows.Next() {
		s := Song{}
		err := rows.Scan(&s.Id, &s.Title, &s.Artist, &s.Length)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		songs = append(songs, s)
		resp, err := json.Marshal(songs)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	
		w.Write(resp)

	}

}

func getSong(w http.ResponseWriter, r *http.Request) {
	rows := db.QueryRow("select * from playlist where $1 = id", r.Header.Values("id"))
	s := Song{}
	rows.Scan(&s.Id, &s.Title, &s.Artist, &s.Length)
	resp, err := json.Marshal(s)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(resp)
}

func createSong(http.ResponseWriter, *http.Request) {
	// var song Song
	// c.BindJSON(&song)

	// db.Create(&song)

	// c.JSON(201, song)
}

func updateSong(http.ResponseWriter, *http.Request) {
	// var song Song
	// id := c.Param("id")
	// db.First(&song, id)

	// if song.ID == 0 {
	// 	c.JSON(404, gin.H{"message": "Song not found"})
	// 	return
	// }

	// c.BindJSON(&song)
	// db.Save(&song)

	// c.JSON(200, song)
}

func deleteSong(http.ResponseWriter, *http.Request) {
	// var song Song
	// id := c.Param("id")
	// db.First(&song, id)

	// if song.ID == 0 {
	// 	c.JSON(404, gin.H{"message": "Song not found"})
	// 	return
	// }

	// db.Delete(&song)

	// c.JSON(200, gin.H{"message": "Song deleted"})
}
