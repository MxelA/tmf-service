package cmd

import (
	"github.com/MxelA/tmf-service/internal/app"
	"github.com/spf13/cobra"
)

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Run the API server",
	Run: func(cmd *cobra.Command, args []string) {
		app.New("host", "port").Serve()
	},
}
