package main

import "fmt"

func main() {

	fmt.Println("Hello Web Server.")
	app := NewWebServer(80, "")
	app.Run()
	//app.RunWithHttp()
}
