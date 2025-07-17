package cmd

import (
	"github.com/MxelA/tmf-service/internal/app"
	"github.com/spf13/cobra"
	"log"
)

var dbCmd = &cobra.Command{
	Use:   "db-set-index",
	Short: "Run db set index",
	Run: func(cmd *cobra.Command, args []string) {
		err := app.ServiceInventoryPkgSetMongoIndex()
		if err != nil {
			log.Fatal("Failed ServiceInventoryPkgSetIndex:", err)
		}
	},
}
