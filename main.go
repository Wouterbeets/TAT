package main

import (
	"TAT/klant"
	"fmt"
	"gopkg.in/qml.v1"
	"os"
)

func main() {
	if err := qml.Run(run); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	engine := qml.NewEngine()

	qml.RegisterTypes("GoExtensions", 1, 0, []qml.TypeSpec{{
		Init: func(k *klant.Klant, obj qml.Object) {
			k.Object = obj
		},
	}})

	component, err := engine.LoadFile("gui.qml")
	if err != nil {
		return err
	}

	win := component.CreateWindow(nil)
	win.Set("x", 560)
	win.Set("y", 320)
	win.Show()
	//obj := win.Root()
	//	obj.On("clicked", func() {
	//		fmt.Println("awesome")
	//	})
	win.Wait()
	return nil
}
