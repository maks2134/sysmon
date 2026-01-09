package main

import (
	"context"
	"fmt"
	"sort"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/process"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type App struct {
	ctx context.Context
}

type ProcessInfo struct {
	PID  int32   `json:"pid"`
	Name string  `json:"name"`
	RAM  float64 `json:"ram"`
}

type Stats struct {
	CPU       float64       `json:"cpu"`
	RAM       float64       `json:"ram"`
	Disk      float64       `json:"disk"`
	Processes []ProcessInfo `json:"processes"`
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	go a.startSystemMonitor()
}

func (a *App) startSystemMonitor() {
	for {
		if a.ctx.Err() != nil {
			return
		}
		cpuPercent, _ := cpu.Percent(1*time.Second, false)
		vmStat, _ := mem.VirtualMemory()
		diskStat, _ := disk.Usage("/")
		stats := Stats{
			CPU:       0,
			RAM:       0,
			Disk:      0,
			Processes: []ProcessInfo{},
		}
		if len(cpuPercent) > 0 {
			stats.CPU = cpuPercent[0]
		}
		if vmStat != nil {
			stats.RAM = vmStat.UsedPercent
		}
		if diskStat != nil {
			stats.Disk = diskStat.UsedPercent
		}
		procs, _ := process.Processes()
		var processesList []ProcessInfo
		for _, p := range procs {
			name, err := p.Name()
			if err != nil {
				continue
			}
			memP, err := p.MemoryPercent()
			if err != nil {
				continue
			}
			processesList = append(processesList, ProcessInfo{
				PID:  p.Pid,
				Name: name,
				RAM:  float64(memP),
			})
		}
		sort.Slice(processesList, func(i, j int) bool {
			return processesList[i].RAM > processesList[j].RAM
		})
		if len(processesList) > 20 {
			processesList = processesList[:20]
		}
		stats.Processes = processesList
		fmt.Printf("CPU: %.1f | RAM: %.1f | Procs: %d\n", stats.CPU, stats.RAM, len(stats.Processes))
		runtime.EventsEmit(a.ctx, "system_stats", stats)
	}
}
