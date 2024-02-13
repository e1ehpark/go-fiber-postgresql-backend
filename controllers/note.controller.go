package controllers

import (
	"strconv"
	"strings"
	"time"

	"github.com/gofiver/fiber/v2"
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