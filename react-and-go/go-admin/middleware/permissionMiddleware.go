package middleware

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/tg112/go/go-admin/database"
	"github.com/tg112/go/go-admin/models"
	jwt "github.com/tg112/go/go-admin/util"
)

func IsAuthorized(c *fiber.Ctx, page string) error {
	cookies := c.Cookies("jwt")

	Id, err := jwt.ParseJwt(cookies)

	if err != nil {
		return err
	}

	userId, _ := strconv.Atoi(Id)

	user := models.User{
		Id: uint(userId),
	}

	database.DB.Preload("Role").Find(&user)

	role := models.Role{
		Id: user.RoleId,
	}

	database.DB.Preload("Permissions").Find(&role)

	if c.Method() == "GET" {
		for _, permission := range role.Permissions {
			if permission.Name == "view_"+page || permission.Name == "edit_"+page {
				return nil
			}
		}
	} else {
		for _, permission := range role.Permissions {
			if permission.Name == "edit_"+page {
				return nil
			}
		}
	}

	c.Status(fiber.StatusUnauthorized)
	return errors.New("unauthorized")

}
