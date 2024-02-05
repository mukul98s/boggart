package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "boggart",
	Short: "boggart uses magic to manage your versions",
	Long: `Seamlessly switch between PHP versions, manage extensions, and perform compatibility 
checks on the fly, all while keeping system-wide changes at bay. Enjoy the flexibility of 
managing PHP versions with ease, optimizing your server's performance, and ensuring your 
codebase remains compatible.`,
}

func Execute() {
	rootCmd.CompletionOptions.HiddenDefaultCmd = true
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
