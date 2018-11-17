package main

import (
	"flag"
	"fmt"
	"github.com/lucasb-eyer/go-colorful"
	"github.com/ottojo/blnk2"
)

var filename = flag.String("c", "/home/jonas/clients.json", "blnk System config file")

func main() {
	flag.Parse()
	system := blnk2.CreateFromFile(*filename)
	go system.Discovery()
	fmt.Println("Clients connected")

	for {
		for pixel := system.Stage.First; pixel != nil; pixel = pixel.Next {
			pixel.Data.Color = colorful.LinearRgb(0, 0, 0)
		}
		system.Commit()
	}
}
