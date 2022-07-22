package engine

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type InputType int

type View interface {
	Render()
}

type InputHandler interface {
	Run() error
	GetType() InputType
}

type KeyboardHandler interface {
	InputHandler
	GetKey() ebiten.Key
	GetJustPressed() bool
}

type MouseButtonHandler interface {
	InputHandler
	GetButton() ebiten.MouseButton
	GetJustPressed() bool
}

type MouseScrollHandler interface {
	InputHandler
}

const (
	Keyboard    InputType = 0
	MouseButton           = 1
	MouseScroll           = 2
)

var views = []View{}
var keyboardHandlers = map[ebiten.Key][]KeyboardHandler{}
var mouseButtonHandlers = map[ebiten.MouseButton][]MouseButtonHandler{}
var mouseScrollHandlers = []MouseScrollHandler{}

func init() {
	// Load images
	//
}

func Run() {
	engine := Engine{}

	err := ebiten.RunGame(&engine)

	if err != nil {
		log.Fatal(err)
	}
}

func AddView(view View) {
	views = append(views, view)
}

func AddInputHandler(handler InputHandler) {
	inputType := handler.GetType()

	switch inputType {
	case Keyboard:
		handler := handler.(KeyboardHandler)
		key := handler.GetKey()
		keyboardHandlers[key] = append(keyboardHandlers[key], handler)
	case MouseButton:
		handler := handler.(MouseButtonHandler)
		button := handler.GetButton()
		mouseButtonHandlers[button] = append(mouseButtonHandlers[button], handler)
	case MouseScroll:
		handler := handler.(MouseScrollHandler)
		mouseScrollHandlers = append(mouseScrollHandlers, handler)
	}
}

func handleInput() error {
	for key, handlers := range keyboardHandlers {
		if !ebiten.IsKeyPressed(key) {
			continue
		}

		justPressed := inpututil.IsKeyJustPressed(key)

		for _, handler := range handlers {
			handlerJustPressed := handler.GetJustPressed()

			if justPressed != handlerJustPressed {
				continue
			}

			err := handler.Run()

			if err != nil {
				return err
			}
		}
	}

	return nil
}

type Engine struct{}

func (e *Engine) Update() error {
	err := handleInput()
	return err
}

func (e *Engine) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, world!")
}

func (e *Engine) Layout(windowWidth, windowHeight int) (screenWidth, screenHeight int) {
	return windowWidth, windowHeight
}
