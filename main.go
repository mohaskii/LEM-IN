package main

import (
	"fmt"
	"lemIn/tools"
)

func main() {
	tools.PutAntsOnPaths()
	for _, p := range tools.ValidePaths {
		fmt.Println(p)
		fmt.Println("")
		fmt.Println("===============================================================================================")
	}
}

/* exemple that return empty Board
exemple 02
*/
