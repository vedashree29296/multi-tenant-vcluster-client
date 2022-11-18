package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:     "mtc",
	Version: "1.0",
	Short:   "some help text",
	Long:    ``,
}

func Execute() {
	rootCmd.AddCommand(createTenantCmd)
	rootCmd.AddCommand(deployCmd)
	_ = rootCmd.Execute()
}
