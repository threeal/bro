package main

import (
	"fmt"

	"github.com/threeal/bro/pkg/utils"
)

func main() {
	backendConfig := utils.InitializeBackendConfig()
	fmt.Println(*backendConfig.ListenAddr)
}
