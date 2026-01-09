package main

import (
	"context"
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

type Stats struct {
	CPU float64 `json:"cpu"`
	RAM float64 `json:"ram"`
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	go a.start()
}

func (a *App) start() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-a.ctx.Done():
			return
		case <-ticker.C:
			cpuPercent, err := cpu.Percent(0, false)
			if err != nil {
				fmt.Println("Error getting CPU stats: ", err)
			}
			ram, err := mem.VirtualMemory()
			if err != nil {
				fmt.Println("Error getting memory stats: ", err)
			}
			stats := Stats{
				CPU: cpuPercent[0],
				RAM: ram.UsedPercent,
			}
			runtime.EventsEmit(a.ctx, "system_stats", stats)
		}
	}
}
