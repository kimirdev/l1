package main

import "fmt"

type Computer interface {
	InsertInSquarePort()
}

type Mac struct {
}

func (m *Mac) InsertInSquarePort() {
	fmt.Println("Insert square port into mac machine")
}

type Windows struct {
}

func (w *Windows) InsertInCirclePort() {
	fmt.Println("Insert circle port into windows machine")
}

type WindowsAdapter struct {
	win *Windows
}

func (a *WindowsAdapter) InsertInSquarePort() {
	a.win.InsertInCirclePort()
}

func InsertSquarePortInComputer(c Computer) {
	c.InsertInSquarePort()
}

func main() {
	mac := &Mac{}
	windows := &Windows{}
	windowsAdapter := &WindowsAdapter{
		win: windows,
	}

	InsertSquarePortInComputer(mac)
	InsertSquarePortInComputer(windowsAdapter)
}
