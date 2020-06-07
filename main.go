package main

import (
	"fmt"
	"github.com/kostaspt/gouted/cmd"
	"os"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
