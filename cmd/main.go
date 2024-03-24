package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	dat, err := os.ReadFile("out/notion-magic-link.html")

	if err != nil {
		fmt.Println(err.Error())
	}
	fileAsStr := string(dat)

	fmt.Println(strings.Replace(fileAsStr, "URL_PLACEHOLDER", "https://dartlsy.app", -1))
	fmt.Println("Hallo Welt2")
}
