package main

import (
	"files/utils"
	"fmt"
	"os"
)

func main() {
	root, _ := os.Getwd()
	path := root + "/data/text.txt"
	c, e := utils.ReadTextFile(path)
	if e != nil {
		fmt.Println("Error", e)
	} else {
		newContent := fmt.Sprintf("Original: %s\nDouble the Original: %s, %s", c, c, c)
		utils.WriteToFile(path+".out", newContent)
		fmt.Println(c)
	}
}
