package cmd

import (
	"fmt"
	"github.com/mukul98s/boggart/helper"
	"github.com/spf13/cobra"
	"os/exec"
)

// useCmd represents the use command
var useCmd = &cobra.Command{
	Use:     "use",
	Short:   "Switch php version",
	Long:    `Switch php version by mentioning version number`,
	Example: "boggart use 8.2",
	Args:    cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			err := cmd.Help()
			if err != nil {
				return
			}
			return
		}
		phpVersionNumber := helper.RemoveWhitespace(args[0])

		if !helper.IsValidPhpVersion(phpVersionNumber) {
			fmt.Printf("Please Enter a valid version number\n\n")
			err := cmd.Help()
			if err != nil {
				return
			}
			return
		}

		phpVersionCheck := fmt.Sprintf("php%v", phpVersionNumber)
		command := exec.Command("which", phpVersionCheck)
		output, err := command.Output()
		if err != nil || string(output) == "" {
			fmt.Printf("PHP v%s is not installed on your system\n", phpVersionNumber)
			return
		}

		phpPath := string(output)
		phpPath = helper.RemoveWhitespace(phpPath)

		changePhpVersionCommand := exec.Command("sudo", "update-alternatives", "--set", "php", phpPath)
		output, err = changePhpVersionCommand.Output()
		if err != nil {
			fmt.Printf("Unable to change PHP to v%s\n", phpVersionNumber)
			return
		}
		if string(output) == "" {
			fmt.Printf("Already using using PHP v%s\n", phpVersionNumber)
			return
		}

		fmt.Printf("Now using PHP version v%s\n", phpVersionNumber)
	},
}

func init() {
	rootCmd.AddCommand(useCmd)
}