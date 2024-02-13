package controllers

import (
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/e1ehpark/go-fiber-postgresql-backend/initializers"
	"github.com/e1ehpark/go-fiber-postgresql-backend/models"
	"gorm.io/gorm"
)

func CreateNoteHandler(c *fiber.Ctx) error {
	varr payload *models.CreageNoteSchema

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	errors := models.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	now := time.Now()
	newNote := models.Note{
		Title:     payload.Title,
		Content:   payload.Content,
		Category:  payload.Category,
		Published: payload.Published,
		CreatedAt: now,
		UpdateAt:  now,
	}

	result := initializers.DB.Create(&newNote)

	if result.Error != nil && strings.Contains(result.Error.Error(),"duplicate key value violates unique") {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"status": "fail", "message": "Title already exist, please use another title"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": fiber.Map{"node": newNote}})
}

func FindNotes(c *fiber.Ctx) error {
	noteId := c.Params("noteId")

	var note models.Note
	result := initializers.DB.First(&note, "id = ?", noteId)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No note with that Id exists"})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"note": note}})
}

func DeleteNote(c *fiber.Ctx) error {
	noteId := c.Params("noteId")

	result := initializers.DB.Delete(&models.Note{}, "id = ?", noteId)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No note with that Id exists"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)

}