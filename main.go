package main

import (
	"embed"
	"runtime"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	wailsRuntime "github.com/wailsapp/wails/v2/pkg/runtime"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	app := NewApp()
	AppMenu := menu.NewMenu()
	if runtime.GOOS == "darwin" {
		AppMenu.Append(menu.AppMenu())
	}
	FileMenu := AppMenu.AddSubmenu("File")
	FileMenu.AddText("Quit", keys.CmdOrCtrl("q"), func(_ *menu.CallbackData) {
		wailsRuntime.Quit(app.ctx)
	})
	ViewMenu := AppMenu.AddSubmenu("View")
	ViewMenu.AddCheckbox("Show CPU", true, nil, func(_ *menu.CallbackData) {
		app.SendToggleEvent("cpu")
	})
	ViewMenu.AddCheckbox("Show RAM", true, nil, func(_ *menu.CallbackData) {
		app.SendToggleEvent("ram")
	})
	ViewMenu.AddCheckbox("Show Disk", true, nil, func(_ *menu.CallbackData) {
		app.SendToggleEvent("disk")
	})
	ViewMenu.AddSeparator()
	ViewMenu.AddText("Reload Window", keys.CmdOrCtrl("r"), func(_ *menu.CallbackData) {
		wailsRuntime.WindowReload(app.ctx)
	})
	HelpMenu := AppMenu.AddSubmenu("Help")
	HelpMenu.AddText("About", nil, func(_ *menu.CallbackData) {
		wailsRuntime.MessageDialog(app.ctx, wailsRuntime.MessageDialogOptions{
			Type:    wailsRuntime.InfoDialog,
			Title:   "About",
			Message: "System Monitor v1.0\nBuilt with Wails & Vue",
		})
	})
	if runtime.GOOS == "darwin" {
		AppMenu.Append(menu.EditMenu())
	}
	err := wails.Run(&options.App{
		Title:  "sysmon",
		Width:  1200,
		Height: 800,
		Menu:   AppMenu,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})
	if err != nil {
		println("Error:", err.Error())
	}
}
