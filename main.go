package main

import (
	"fmt"
	"log"
	"runtime"

	"github.com/felixangell/phi-editor/cfg"
	"github.com/felixangell/phi-editor/gui"
	"github.com/felixangell/strife"
)

const (
	PRINT_FPS bool = true
)

type PhiEditor struct {
	gui.BaseComponent
	running     bool
	defaultFont *strife.Font
}

func (n *PhiEditor) init(cfg *cfg.TomlConfig) {
	n.AddComponent(gui.NewView(1280, 720, cfg))

	font, err := strife.LoadFont("./res/firacode.ttf", 14)
	if err != nil {
		panic(err)
	}
	n.defaultFont = font
}

func (n *PhiEditor) dispose() {
	for _, comp := range n.GetComponents() {
		gui.Dispose(comp)
	}
}

func (n *PhiEditor) update() bool {
	needsRender := false
	for _, comp := range n.GetComponents() {
		dirty := gui.Update(comp)
		if dirty {
			needsRender = true
		}
	}
	return needsRender
}

func (n *PhiEditor) render(ctx *strife.Renderer) {
	ctx.Clear()
	ctx.SetFont(n.defaultFont)

	for _, child := range n.GetComponents() {
		gui.Render(child, ctx)
	}

	ctx.Display()
}

func main() {
	config := cfg.Setup()

	ww, wh := 1280, 720
	window := strife.SetupRenderWindow(ww, wh, strife.DefaultConfig())
	window.SetTitle("Hello world!")
	window.SetResizable(true)

	editor := &PhiEditor{running: true}
	window.HandleEvents(func(evt strife.StrifeEvent) {
		switch evt.(type) {
		case *strife.CloseEvent:
			window.Close()
		}
	})

	window.Create()

	{
		size := "16"
		switch runtime.GOOS {
		case "windows":
			size = "256"
		case "darwin":
			size = "512"
		case "linux":
			size = "96"
		default:
			log.Println("unrecognized runtime ", runtime.GOOS)
		}

		icon, err := strife.LoadImage("./res/icons/icon" + size + ".png")
		if err != nil {
			panic(err)
		}
		window.SetIconImage(icon)
		defer icon.Destroy()
	}

	editor.init(&config)

	timer := strife.CurrentTimeMillis()
	num_frames := 0

	ctx := window.GetRenderContext()

	editor.render(ctx)
	for {
		window.PollEvents()
		if window.CloseRequested() {
			break
		}

		if editor.update() {
			editor.render(ctx)
		}

		num_frames += 1

		if strife.CurrentTimeMillis()-timer > 1000 {
			timer = strife.CurrentTimeMillis()
			if PRINT_FPS {
				fmt.Println("frames: ", num_frames)
			}
			num_frames = 0
		}
	}

	editor.dispose()
}
