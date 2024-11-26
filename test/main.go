package main

import (
	"fmt"
	"github.com/YuanJey/goconf/pkg/config"
)

func main() {
	fmt.Println(config.Config)
	fmt.Println(config.Config.Test.GoVERSION)
}
