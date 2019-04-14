package main

import (
	"github.com/mlctrez/crystal/rdisplay"
)

func main() {

	d, err := rdisplay.Connect("10.0.0.230:8266")
	failError(err)
	defer d.Close()

	failError(d.Clear())
	failError(d.Print("Hello World\r\n"))

}

func failError(err error) {
	if err != nil {
		panic(err)
	}
}
