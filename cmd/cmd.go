package cmd

import "github.com/spf13/cobra"

var (
	name string

	host string
	port string

	ext string
)

var rootCmd = &cobra.Command{Use: "app"}

func init() {
	rootCmd.AddCommand(apiCmd)

	apiCmd.Flags().StringVar(&port, "port", "8080", "Port to run the API server")
	apiCmd.Flags().StringVar(&host, "host", "0.0.0.0", "Host to run the API server")
	rootCmd.Execute()

}

func Execute() error {
	return rootCmd.Execute()
}
