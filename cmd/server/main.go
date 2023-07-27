package main

import "fmt"

// App - the struct holds the db connection
type App struct{}

func (app *App) Run() error {
	fmt.Println("Setting up the app")
	return nil
}

func main() {
	fmt.Println("Demo REST API!")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting up our REST API")
		fmt.Println(err)
	}
}
