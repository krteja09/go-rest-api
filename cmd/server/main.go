package main

import "fmt"


//App -struct which contains things like pointers to database connections
type App struct {

}

func (app *App)Run() error {
	fmt.Println("Setting up Error")
	return nil
}
func main()  {
	fmt.Println("Go Rest API ")
	app := App{}
	if err := app.Run(); err !=nil {
		fmt.Println("Error starting up our REST API")
		fmt.Println(err)
	}
}
