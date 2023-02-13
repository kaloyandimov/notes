package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"notes/backend/services/database"
	"notes/backend/utilities/token"
)

type NoteInput struct {
	Title       string `json:"title" binding:"required,gte=1"`
	Description string `json:"description"`
}

func Index(c *gin.Context) {
	userID, err := token.ExtractUserID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "data": err.Error()})
		return
	}

	var notes []database.Note
	notes, err = database.GetNotes(userID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "data": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "success", "data": notes})
	}
}

func Create(c *gin.Context) {
	var input NoteInput
	var err error

	if err = c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "data": err.Error()})
		return
	}

	var userID uint
	userID, err = token.ExtractUserID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "data": err.Error()})
		return
	}

	err = database.AddNote(input.Title, input.Description, userID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "data": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "success", "data": "Note added successfully"})
	}
}

func Show(c *gin.Context) {
	note, err := database.GetNote(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "data": err.Error()})
		return
	}

	var userID uint
	userID, err = token.ExtractUserID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "data": err.Error()})
		return
	}

	if userID != note.UserID {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "data": "Unauthorized"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": note})
}

func Update(c *gin.Context) {
	var note database.Note
	var input NoteInput
	var err error

	if err = c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "data": err.Error()})
		return
	}

	note, err = database.GetNote(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "data": err.Error()})
		return
	}

	var userID uint
	userID, err = token.ExtractUserID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "data": err.Error()})
		return
	}

	if userID != note.UserID {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "data": "Unauthorized"})
		return
	}

	err = database.UpdateNote(note, input.Title, input.Description)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "data": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "success", "data": "Note updated successfully"})
	}
}

func Destroy(c *gin.Context) {
	note, err := database.GetNote(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "data": err.Error()})
		return
	}

	var userID uint
	userID, err = token.ExtractUserID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "data": err.Error()})
		return
	}

	if userID != note.UserID {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "data": "Unauthorized"})
		return
	}

	if err = database.DeleteNote(note); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "data": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": "Note deleted successfully"})
}
