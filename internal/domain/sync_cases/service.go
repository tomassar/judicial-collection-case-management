package sync_cases

import (
	"context"

	"github.com/chromedp/chromedp"
	"github.com/tomassar/judicial-collection-case-management/internal/http/scraper"
)

type Service interface {
	SyncCases(ctx context.Context) error
}

type service struct {
}

func NewService() Service {
	return &service{}
}

func (s *service) SyncCases(ctx context.Context) error {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
		chromedp.Flag("disable-gpu", false),
	)

	// Create a new allocator context
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	// Create a new context
	ctx, cancel = chromedp.NewContext(allocCtx)
	defer cancel()

	scraper.FillForm(ctx, scraper.FormData{
		Competencia: "3",
		Corte:       "50",
		Tribunal:    "197", //198, 406
		LibroTipo:   "C",   //E
		Rol:         2,
		Year:        2024,
	})

	return nil
}
