package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/tg112/go/go-admin/database"
	"github.com/tg112/go/go-admin/models"
)

func AllOrders(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	return c.JSON(models.Paginate(database.DB, &models.Order{}, page))
}
