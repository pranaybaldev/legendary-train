package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "magic",
	Short: "CLI for some random testing",
	Long:  `Magician testing some stuff out`,
	// Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
