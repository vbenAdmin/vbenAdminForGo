package main

import (
	"goVben/Routers"
)

func main() {
	r := Routers.SetupRouter()
	// Listen and Server in 0.0.0.0:8080
	err := r.Run(":39874")
	if err != nil {

	}

}
