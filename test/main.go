package main

import (
	"encoding/json"
	"fmt"
	"github.com/YuanJey/goconf/pkg/config"
	"os"
)

func main() {
	//os.Setenv("test2", "{\"a\":\"12\"}")
	err := json.Unmarshal([]byte("{\"a\":\"12\"}"), &config.Config.Test2)
	if err != nil {
	}
	fmt.Println(config.Config)
	fmt.Println(config.Config.Test.GoVERSION)
	fmt.Println(config.Config.Test2)

	fmt.Println(os.Getenv("test2"))
}
