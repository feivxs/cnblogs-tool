package utils

import (
	"fmt"
	"os"
)

func HandleErr(error error) {
	if error != nil {
		fmt.Println(error.Error())
		os.Exit(1)
	}
}
