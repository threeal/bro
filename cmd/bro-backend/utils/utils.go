package utils

import (
	"github.com/threeal/bro/cmd/bro-backend/config"
	"github.com/threeal/bro/pkg/tcp"
	"github.com/threeal/bro/pkg/utils"
)

func CreateServer() (*tcp.Server, error) {
	config := &config.Config{}
	err := utils.InitializeConfig(config)
	if err != nil {
		return nil, err
	}
	addr := config.ListenAddr
	return tcp.NewServer(addr)
}
