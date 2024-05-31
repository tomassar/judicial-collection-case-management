package scraper

import (
	"context"
	"fmt"
	"time"

	"github.com/chromedp/chromedp"
)

type FormData struct {
	Competencia string
	Corte       string
	Tribunal    string
	LibroTipo   string
	Rol         int
	Year        int
}

func FillForm(ctx context.Context, data FormData) error {
	// Define the selectors
	thirdDivButtonSel := `div.container:nth-child(4) > div:nth-child(1) > div:nth-child(3) > div:nth-child(1) > button:nth-child(1)`
	competenciaSel := `#competencia`
	corteSel := `#conCorte`
	tribunalSel := `#conTribunal`
	libroTipoSel := `#conTipoCausa`
	rolSel := `#conRolCausa`
	yearSel := `#conEraCausa`
	submitSel := `#btnConConsulta`

	// Run the tasks
	return chromedp.Run(ctx,
		// Navigate to the page
		chromedp.Navigate(`https://oficinajudicialvirtual.pjud.cl/indexN.php`),
		chromedp.WaitVisible(thirdDivButtonSel),
		chromedp.Click(thirdDivButtonSel),
		// Wait for the form to be visible
		chromedp.WaitVisible(competenciaSel),
		// Fill out the form fields
		chromedp.SetValue(competenciaSel, data.Competencia, chromedp.ByQuery),
		/* 		chromedp.SetValue(corteSel, data.Corte),
		   		chromedp.SetValue(tribunalSel, data.Tribunal),
		   		chromedp.SetValue(libroTipoSel, data.LibroTipo),
		   		chromedp.SetValue(rolSel, fmt.Sprint(data.Rol)),
		   		chromedp.SetValue(yearSel, fmt.Sprint(data.Year)), */

		// Click the submit button
		/* chromedp.Click(submitSel), */

		// Wait for navigation to complete (or for a specific element to be visible)
		/* chromedp.WaitVisible(`#resultDiv`, chromedp.ByID), */
	)
}
