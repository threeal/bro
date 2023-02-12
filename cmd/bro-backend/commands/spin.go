package commands

import (
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"
	backendUtils "github.com/threeal/bro/cmd/bro-backend/utils"
	"github.com/threeal/bro/pkg/db/sqlite/log_system"
	"github.com/threeal/bro/pkg/schema"
	"github.com/threeal/bro/pkg/service"
)

func getSpinCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "spin",
		Short: "Spin command",
		Long:  `A command to run Bro backend process..`,
		Args:  cobra.NoArgs,
		PreRun: func(cmd *cobra.Command, args []string) {
			db := log_system.InitDB()
			log_system.QueryDB(db, log_system.InsertToSessionsTableSQL)
			sessionID, err := log_system.QueryLastID(db)
			if err != nil {
				log.Fatalf("cannot get session id, %v", err)
			}
			dbWriter := &log_system.DBWriter{}
			dbWriter.SetDB(db)
			dbWriter.SetSessionID(sessionID)
			mw := io.MultiWriter(dbWriter, os.Stdout)
			log.SetOutput(mw)
		},
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
