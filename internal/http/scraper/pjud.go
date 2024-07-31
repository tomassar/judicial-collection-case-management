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
	thirdDivButtonSel := `div.container:nth-child(3) > div:nth-child(1) > div:nth-child(3) > div:nth-child(1) > button:nth-child(1)`
	competenciaSel := `html body div.wrapper div#content div#contMain div.container-4 div.card.shadow.mb-4 div.cajaContentIzq.card.border-left-success.shadow.h-100.py-2 section.row div.panel-body2 div div.tab-content.hidden-sm.hidden-xs div#busRit.tab-pane.fade.in.active div.jumboTabs form#formConsulta div.form-group.col-md-4 select#competencia.form-control`
	corteSel := `html body div.wrapper div#content div#contMain div.container-4 div.card.shadow.mb-4 div.cajaContentIzq.card.border-left-success.shadow.h-100.py-2 section.row div.panel-body2 div div.tab-content.hidden-sm.hidden-xs div#busRit.tab-pane.fade.in.active div.jumboTabs form#formConsulta div.form-group.col-md-4 select#conCorte.form-control`
	tribunalSel := `html body div.wrapper div#content div#contMain div.container-4 div.card.shadow.mb-4 div.cajaContentIzq.card.border-left-success.shadow.h-100.py-2 section.row div.panel-body2 div div.tab-content.hidden-sm.hidden-xs div#busRit.tab-pane.fade.in.active div.jumboTabs form#formConsulta div.form-group.col-md-4 select#conTribunal.form-control`
	libroTipoSel := `html body div.wrapper div#content div#contMain div.container-4 div.card.shadow.mb-4 div.cajaContentIzq.card.border-left-success.shadow.h-100.py-2 section.row div.panel-body2 div div.tab-content.hidden-sm.hidden-xs div#busRit.tab-pane.fade.in.active div.jumboTabs form#formConsulta div div.col-md-4.row-2 select#conTipoCausa.form-control`
	rolSel := `html body div.wrapper div#content div#contMain div.container-4 div.card.shadow.mb-4 div.cajaContentIzq.card.border-left-success.shadow.h-100.py-2 section.row div.panel-body2 div div.tab-content.hidden-sm.hidden-xs div#busRit.tab-pane.fade.in.active div.jumboTabs form#formConsulta div div#rolConTipoCausa div.col-md-2.row-3 input#conRolCausa.form-control.inptNum`
	yearSel := `html body div.wrapper div#content div#contMain div.container-4 div.card.shadow.mb-4 div.cajaContentIzq.card.border-left-success.shadow.h-100.py-2 section.row div.panel-body2 div div.tab-content.hidden-sm.hidden-xs div#busRit.tab-pane.fade.in.active div.jumboTabs form#formConsulta div div#rolConTipoCausa div.col-md-2.row-4 input#conEraCausa.form-control.inptNum.ex`
	submitSel := `html body div.wrapper div#content div#contMain div.container-4 div.card.shadow.mb-4 div.cajaContentIzq.card.border-left-success.shadow.h-100.py-2 section.row div.panel-body2 div div.tab-content.hidden-sm.hidden-xs div#busRit.tab-pane.fade.in.active div.jumboTabs form#formConsulta div.text-center button#btnConConsulta.btn.btn-primary`

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
		chromedp.Sleep(2*time.Second),
		chromedp.SetValue(corteSel, data.Corte, chromedp.ByQuery),
		chromedp.Sleep(2*time.Second),
		chromedp.SetValue(tribunalSel, data.Tribunal, chromedp.ByQuery),
		chromedp.Sleep(2*time.Second),
		chromedp.SetValue(libroTipoSel, data.LibroTipo),
		chromedp.Sleep(2*time.Second),
		chromedp.SetValue(rolSel, fmt.Sprint(data.Rol)),
		chromedp.SetValue(yearSel, fmt.Sprint(data.Year)),

		// Click the submit button
		chromedp.Click(submitSel),
		chromedp.Sleep(5*time.Second),

		// Wait for navigation to complete (or for a specific element to be visible)
		/* chromedp.WaitVisible(`#resultDiv`, chromedp.ByID), */
	)
}
