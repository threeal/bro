package commands

import (
	"log"

	"github.com/spf13/cobra"
	backendUtils "github.com/threeal/bro/cmd/bro-backend/utils"
	"github.com/threeal/bro/pkg/schema"
	"github.com/threeal/bro/pkg/service"
)

func getSpinCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "spin",
		Short: "Spin command",
		Long:  `A command to run Bro backend process..`,
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			server, err := backendUtils.CreateServer()
			if err != nil {
				log.Fatalf("failed to create a new server: %v", err)
			}
			schema.RegisterEchoServer(server, &service.EchoServer{})
			log.Printf("server listening at %v", server.Addr())
			if err := server.Serve(); err != nil {
				log.Fatalf("failed to serve: %v", err)
			}
		},
	}
}
