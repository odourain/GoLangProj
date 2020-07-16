package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string] Vertex{
	"Bell Labs": Vertex {
		40.52312, -74.39929,
	},
	"Google" : Vertex {
		37.23552, -122.07542,
	},
}

func main()  {
	fmt.Println(m)
}
