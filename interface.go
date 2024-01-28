package main

import "fmt"

type HandPhone interface {
	OpenApp() string
}

type CodingApps struct {
	Name string
}

type Games struct {
	Name string
}

func (apps CodingApps) OpenApp() string {
	return apps.Name
}

func (games Games) OpenApp() string {
	return games.Name
}

func MonitorRam(handPhone HandPhone) {
	fmt.Println("Show Usable Ram in APP : " + handPhone.OpenApp())
}

func main() {
	apps := CodingApps{Name: "Vscode"}
	game := Games{Name: "Apex Legend"}
	MonitorRam(apps)
	MonitorRam(game)
}
