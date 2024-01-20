package controllers

import (
	"context"
	"github.com/Jneville0815/fitness-app-backend-v2/app/models"
	"github.com/Jneville0815/fitness-app-backend-v2/pkg/utils"
	"github.com/Jneville0815/fitness-app-backend-v2/platform/cache"
	"github.com/google/uuid"
	"time"

	"github.com/gofiber/fiber/v2"
)

func RenewTokens(c *fiber.Ctx) error {
	now := time.Now().Unix()

	renew := &models.Renew{}

	if err := c.BodyParser(renew); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	expiresRefreshToken, err := utils.ParseRefreshToken(renew.RefreshToken)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	if now < expiresRefreshToken {
		userID, err := uuid.Parse(c.Params("user_id"))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		}

		connRedis, err := cache.RedisConnection()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		}

		redisRefresh, err := connRedis.Get(context.Background(), userID.String()).Result()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		}
		if redisRefresh != renew.RefreshToken {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": true,
				"msg":   "no token on server to match token submitted",
			})
		}

		tokens, err := utils.GenerateNewTokens(userID.String())
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		}

		errRedis := connRedis.Set(context.Background(), userID.String(), tokens.Refresh, 0).Err()
		if errRedis != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": true,
				"msg":   errRedis.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"error": false,
			"msg":   nil,
			"tokens": fiber.Map{
				"access":  tokens.Access,
				"refresh": tokens.Refresh,
			},
		})
	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "unauthorized, refresh token expired",
		})
	}
}
