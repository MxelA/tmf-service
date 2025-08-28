package cmd

import (
	"github.com/MxelA/tmf-service/internal/app"
	"github.com/spf13/cobra"
	"log"
)

var dbCmd = &cobra.Command{
	Use:   "db:service-inventory:set-index",
	Short: "Run db set index",
	Run: func(cmd *cobra.Command, args []string) {
		err := app.ServiceInventoryPkgSetMongoIndex()
		if err != nil {
			log.Fatal("Failed ServiceInventoryPkgSetIndex:", err)
		}
	},
}

var dbSeedServiceCmd = &cobra.Command{
	Use:   "seed:service-inventory",
	Short: "Run db seed service: --seedName=[withRelationshipTo]|None --count=1 --deep=2 --rootServiceName=''",
	Run: func(cmd *cobra.Command, args []string) {
		var err error

		switch seedName {
		case "withRelationshipTo":
			err = app.SeedServicesWithRelationshipTo(seedCount, deep, rootServiceName)
		default:
			err = app.SeedServices(seedCount)
		}

		if err != nil {
			log.Fatal("Failed Seed Services:", err)
		}
	},
}
