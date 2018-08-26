package main

import "fmt"

func main() {
	fmt.Println("Running on 8080....")
	a := App{}
	a.Initialize()
	a.Run("localhost:8080")
}
