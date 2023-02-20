package database

import (
	"errors"
	"html"
	"strings"

	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	Title       string `gorm:"size:255;not null"`
	Description string `gorm:"size:255"`
	UserID      uint   `gorm:"not null"`
}

func GetNote(ID string) (Note, error) {
	var note Note

	if err := database.Where("id = ?", ID).First(&note).Error; err != nil {
		return note, errors.New("note not found")
	}

	return note, nil
}

func GetNotes(userID uint) ([]Note, error) {
	var notes []Note

	if err := database.Where("user_id = ?", userID).Find(&notes).Error; err != nil {
		return notes, errors.New("user not found")
	}

	return notes, nil
}

func AddNote(title, description string, userID uint) error {
	note := Note{
		Title:       html.EscapeString(strings.TrimSpace(title)),
		Description: html.EscapeString(strings.TrimSpace(description)),
		UserID:      userID,
	}

	if err := database.Create(&note).Error; err != nil {
		return errors.New("could not add note")
	}

	return nil
}

func UpdateNote(note Note, title, description string) error {
	note.Title = html.EscapeString(strings.TrimSpace(title))
	note.Description = html.EscapeString(strings.TrimSpace(description))

	if err := database.Save(&note).Error; err != nil {
		return errors.New("could not update note")
	}

	return nil
}

func DeleteNote(note Note) error {
	if err := database.Delete(&note).Error; err != nil {
		return errors.New("could not delete note")
	}

	return nil
}
