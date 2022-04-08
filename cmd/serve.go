package cmd

import (
	"log"
	"net/http"

	"github.com/lateralusd/laserver/db"
	"github.com/lateralusd/laserver/handler"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the server",
	RunE: func(cmd *cobra.Command, args []string) error {
		dbPath, err := cmd.Flags().GetString("db")
		if err != nil {
			return err
		}
		addr, err := cmd.Flags().GetString("iface")
		if err != nil {
			return err
		}

		log.Printf("Starting the server on %s", addr)
		log.Printf("Using database %s", dbPath)

		db := db.NewDB(dbPath)
		defer db.Close()

		h := &handler.Handler{
			DB: db,
		}
		http.Handle("/", h)
		return http.ListenAndServe(addr, nil)
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)
	serveCmd.Flags().StringP("db", "d", "data.db", "path to database")
	serveCmd.Flags().StringP("iface", "i", ":80", "which interface:port to bind to")
}
