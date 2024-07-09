package main

import (
	"context"
	"fmt"
	"time"

	data "desktop.rythm/core/data"
	testData "desktop.rythm/core/test"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) CreateCandle(time time.Time, open float64, high float64, low float64, close float64) data.Candle {
	return data.NewCandle(time, open, high, low, close)
}

func (a *App) ExportCandleHistory() []data.FrontEndCandle {
	return testData.Simple_Wave_Frontend
}

// []models.CandleStick
func (a *App) OpenStream() []data.FrontEndCandle  {
	// candles := data.QueryCandles()
	newData := data.QueryCandleHistory()
	// for _, c := range candles {
	// 	feC := data.NewFrontEndCandle(c.Time, c.Mid.O, c.Mid.H, c.Mid.L, c.Mid.C)
	// 	newData = append(newData, feC)
	// }
	return newData
}

func (a *App) CloseStream() {
	// data.QuitWatching()
}