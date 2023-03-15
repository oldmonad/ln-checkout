package api

import (
	"github.com/gofiber/fiber/v2"
)

type PaymenHandler interface {
	CreatePaymentLink(ctx *fiber.Ctx)
	FetchPaymentLink(ctx *fiber.Ctx)
}
