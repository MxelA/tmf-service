package cmd

import "github.com/spf13/cobra"

var (
	name string

	host            string
	port            string
	ext             string
	seedName        string
	seedCount       int
	deep            int
	rootServiceName string
)

var rootCmd = &cobra.Command{Use: "app"}

func init() {
	rootCmd.AddCommand(apiCmd)

	apiCmd.Flags().StringVar(&port, "port", "8080", "Port to run the API server")
	apiCmd.Flags().StringVar(&host, "host", "0.0.0.0", "Host to run the API server")

	rootCmd.AddCommand(dbCmd)

	rootCmd.AddCommand(dbSeedServiceCmd)
	dbSeedServiceCmd.Flags().StringVar(&seedName, "seedName", "single", "Service seeder name")
	dbSeedServiceCmd.Flags().IntVar(&seedCount, "count", 1, "Service seeder count")
	dbSeedServiceCmd.Flags().IntVar(&deep, "deep", 2, "How many service in relationship")
	dbSeedServiceCmd.Flags().StringVar(&rootServiceName, "rootServiceName", "", "Name of the root service to seed")

	rootCmd.Execute()

}

func Execute() error {
	return rootCmd.Execute()
}
