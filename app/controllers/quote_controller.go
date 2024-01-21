package controllers

import (
	"github.com/Jneville0815/fitness-app-backend-v2/app/models"
	"github.com/Jneville0815/fitness-app-backend-v2/pkg/utils"
	"github.com/Jneville0815/fitness-app-backend-v2/platform/database"
	"github.com/gofiber/fiber/v2"
)

func CreateQuote(c *fiber.Ctx) error {
	claims, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	quote := &models.Quote{}
	quote.UserID = claims.UserID

	if err := c.BodyParser(quote); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	validate := utils.NewValidator()

	if err := validate.Struct(quote); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}

	if err := database.DB.CreateQuote(quote); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"quote": quote,
	})
}

func GetQuotes(c *fiber.Ctx) error {
	claims, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	userId := claims.UserID

	quotes, err := database.DB.GetQuotes(userId.String())
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error":  false,
		"msg":    nil,
		"quotes": quotes,
	})
}

func DeleteQuote(c *fiber.Ctx) error {
	claims, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	quoteId := c.Params("quote_id")
	userId := claims.UserID

	err = database.DB.DeleteQuote(userId.String(), quoteId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error": false,
		"msg":   "successfully deleted quote",
	})
}
