package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Note struct {
	ID        int    `json:"id"`
	NoteTitle string `json:"note_title"`
	NoteText  string `json:"note_text"`
	CreatedBy string `json:"created_by"`
	Favorited bool   `json:"favorite"`
}

var notes = []Note{}

func getNotes(c *gin.Context) {
	db, err := getDbConnection()

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Issue connecting to db")
	}

	db.Find(&notes)

	c.IndentedJSON(http.StatusOK, notes)
}

func getDbConnection() (*gorm.DB, error) {
	dsn := ""
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, errors.New("issue connecting to the database")
	}

	return db, nil
}

func createNote(c *gin.Context) {
	var newNote Note

	if err := c.BindJSON(&newNote); err != nil {
		return
	}

	notes = append(notes, newNote)
	c.IndentedJSON(http.StatusCreated, newNote)
}

func noteById(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, "Invalid ID")
		return
	}
	note, err := getNoteById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, "No Book Found")
		return
	}

	c.IndentedJSON(http.StatusOK, note)

}

func getNoteById(id int) (*Note, error) {
	for i, n := range notes {
		if n.ID == id {
			return &notes[i], nil
		}
	}

	return nil, errors.New("error getting note")
}

func favoriteNoteById(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)

	if err != nil {
		return
	}

	updatedNote, err := favoriteNote(id)

	if err != nil {
		return
	}

	c.IndentedJSON(http.StatusAccepted, updatedNote)
}

func favoriteNote(id int) (*Note, error) {
	for i, n := range notes {
		if n.ID == id {
			currNote := &notes[i]
			currNote.Favorited = !currNote.Favorited
			return currNote, nil
		}
	}

	return nil, errors.New("error updating note")
}

func main() {
	// connect()
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/notes", getNotes)
	router.POST("/notes", createNote)
	router.GET("/notes/:id", noteById)
	router.PATCH("/notes/favorites/:id", favoriteNoteById)
	router.Run("localhost:8080")
}
