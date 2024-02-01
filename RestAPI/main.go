package main

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	user     = ""
	password = ""
	host     = ""
	port     = 0
	dbname   = ""
)

type Note struct {
	ID        int    `json:"id"`
	NoteTitle string `json:"note_title"`
	NoteText  string `json:"note_text"`
	CreatedBy string `json:"created_by"`
	Favorite  bool   `json:"favorite"`
}

var notes = []Note{
	{ID: 1, NoteTitle: "Groceries", NoteText: "Milk, Bread, Eggs", CreatedBy: "Oscar", Favorite: false},
	{ID: 2, NoteTitle: "Consoles To Play", NoteText: "PS2, PS5, XBox, Sega Genesis", CreatedBy: "Oscar", Favorite: false},
	{ID: 3, NoteTitle: "Classes", NoteText: "Math, English 101, Algebra, Intro To CS 2", CreatedBy: "Oscar", Favorite: false},
}

func getNotes(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, notes)

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

	return nil, errors.New("Error Getting Note")
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
			currNote.Favorite = !currNote.Favorite
			return currNote, nil
		}
	}

	return nil, errors.New("Error Updating Note")
}

func connect() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	rows, err := db.Query("SELECT ID, CREATED_AT, CREATED_BY, NOTE_TEXT, NOTE_TITLE, FAVORITED FROM NOTES")

	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		err = rows.Scan()
		if err != nil {
			// handle this error
			panic(err)
		}
	}

	fmt.Println("Successfully connected!")

}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
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
