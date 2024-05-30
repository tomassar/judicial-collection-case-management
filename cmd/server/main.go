package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/chromedp/chromedp"
	"github.com/joho/godotenv"
	"github.com/tomassar/judicial-collection-case-management/internal/database"
	"github.com/tomassar/judicial-collection-case-management/internal/domain/auth"
	"github.com/tomassar/judicial-collection-case-management/internal/domain/cases"
	"github.com/tomassar/judicial-collection-case-management/internal/domain/dashboard"
	"github.com/tomassar/judicial-collection-case-management/internal/domain/lawyers"
	"github.com/tomassar/judicial-collection-case-management/internal/domain/users"
	"github.com/tomassar/judicial-collection-case-management/internal/http/rest"
	"github.com/tomassar/judicial-collection-case-management/internal/http/scraper"
	"github.com/tomassar/judicial-collection-case-management/internal/storage/postgres"
)

func main() {
	initService()
}

func initService() {
	//init env variables with godotenv
	err := godotenv.Load()
	if err != nil {
		panic("failed to load env variables")
	}

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
		chromedp.Flag("disable-gpu", false),
	)

	// Create a new allocator context
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	// Create a new context
	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	// Create a new data instance
	data := scraper.FormData{
		Competencia: "3",
		Corte:       "50",
		Tribunal:    "197",
		LibroTipo:   "C",
		Rol:         2,
		Year:        2024,
	}

	// Run the form filling function
	if err := scraper.FillForm(ctx, data); err != nil {
		log.Fatal(err)
	}

	fmt.Println("DB: ", os.Getenv("DB"))
	dsn := os.Getenv("DB")
	db, err := database.Connect(dsn)
	if err != nil {
		panic("failed to connect database")
	}

	storage := postgres.NewStorage(db)

	caseService := cases.NewService(storage.Cases)
	userService := users.NewService(storage.Users)
	lawyerService := lawyers.NewLawyerService(storage.Lawyers)
	authService := auth.NewService(userService, lawyerService)
	dashboardService := dashboard.NewService(caseService)

	router := rest.
		NewHandler(caseService, userService, authService, lawyerService, dashboardService).
		Init()

	router.Run(os.Getenv("HOST_ADDR"))

}
