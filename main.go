package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"

	"github.com/lncm/lnd-rpc/v0.10.0/lnrpc"
	"github.com/oldmonad/ln-checkout.git/api"
	"github.com/oldmonad/ln-checkout.git/config"
	"github.com/oldmonad/ln-checkout.git/service"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbuser := os.Getenv("DB_USER")
	dbPassord := os.Getenv("DB_USER")

	dbInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbuser, dbPassord, dbName)

	repo, err := config.NewDbRepository(dbInfo, 5)

	if err != nil {
		log.Fatal(err)
	}

	service := service.NewService(repo)
	handler := api.NewHandler(service)

	r := fiber.New()
	r.Use(recover.New())
	r.Use(logger.New(logger.Config{
		Format: "[${time}] ${ip}  ${status} - ${latency} ${method} ${path}\n",
	}))

	r.Post("/payment_link", handler.CreatePaymentLink)
	r.Get("/payment_link/:unique_reference", handler.FetchPaymentLink)

	client, err := config.NewLnConnection()

	if err != nil {
		log.Fatal(err)
	}

	go func() {
		stream, err := client.SubscribeInvoices(context.Background(), &lnrpc.InvoiceSubscription{})
		if err != nil {
			log.Fatalf("could not subscribe to invoices: %v", err)
		}

		for {
			invoice, err := stream.Recv()
			if err != nil {
				log.Fatalf("error receiving invoice: %v", err)
			}

			if invoice.Settled {
				// log.Printf("invoice %s settled", invoice.PaymentRequest)
				// log.Printf("invoice %s settled", invoice.RHash)
				log.Printf("invoice %s settled", base64.StdEncoding.EncodeToString(invoice.RHash))
			} else {
				// log.Printf("invoice %s has been created", invoice.PaymentRequest)
				// log.Printf("invoice %s has been created", invoice.RHash)
				fmt.Printf("invoice %v\n has been created", base64.StdEncoding.EncodeToString(invoice.RHash))
			}
		}
	}()
	r.Listen(":4004")
}
