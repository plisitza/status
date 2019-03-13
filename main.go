package main

import "fmt"

// keep user oriented status info in a potentially threaded environment and provide hooks for that to be retrieved and displayed in a display layer
// or display itself via command line-style lines of logging, or something like https://github.com/gizak/termui
// for providing info to a more gui display layer, start a goroutine that listens on a configured port

func main() {
	fmt.Println("hi!")
}
