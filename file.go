package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Create("data/a.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(file)
		fmt.Println("Done")
	}

}
