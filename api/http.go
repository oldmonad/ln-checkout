package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/oldmonad/ln-checkout.git/domain"
)

type Handler struct {
	service domain.Service
}

func (h *Handler) CreatePaymentLink(ctx *fiber.Ctx) error {
	// timeout, err := strconv.ParseInt(os.Getenv("DB_HOST"), 0, 64)

	// if err != nil {
	// 	return ctx.Status(500).JSON(err)
	// }

	// context, cancel := context.WithTimeout(context.Background(), time.Duration(timeout))
	// defer cancel()

	p := &domain.Payment{}

	if err := ctx.BodyParser(&p); err != nil {
		fmt.Println(err)
		return ctx.Status(500).JSON(err)
	}

	data, err := h.service.CreatePaymentLink(p)

	if err != nil {
		fmt.Println(err)
		return ctx.Status(500).JSON(err)
	}

	return ctx.JSON(data)
}

func (h *Handler) FetchPaymentLink(ctx *fiber.Ctx) error {

	reference := ctx.Params("unique_reference")

	p, err := h.service.FetchPaymentLink(reference)

	if err != nil {
		return ctx.Status(404).JSON(nil)
	}

	return ctx.JSON(&p)
}

// NewHandler  New handler instantiates a http handler for our product service
func NewHandler(service domain.Service) *Handler {
	return &Handler{service: service}
}
