package main

import (
	"fmt"
	"lemIn/tools"
)

func main() {
	tools.GetMaxEndedPath()

	fmt.Println(tools.FindChunk(16))
}
