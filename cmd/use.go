package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/mukul98s/boggart/helper"
	"github.com/spf13/cobra"
)

// useCmd represents the use command
var useCmd = &cobra.Command{
	Use:     "use",
	Short:   "Switch php version",
	Long:    `Switch php version by mentioning version number or by mentioning php version in .phpversion file`,
	Example: "boggart use 8.2",
	Args:    cobra.MaximumNArgs(1),
	Run:     runUseCommand,
}

func runUseCommand(cmd *cobra.Command, args []string) {
	var phpVersionNumber string

	file, _ := os.ReadFile(".phpversion")
	if len(file) > 0 {
		phpVersionNumber = helper.RemoveWhitespace(string(file))
	}

	if len(args) == 0 && phpVersionNumber == "" {
		_ = cmd.Help()
		return
	}

	if len(args) > 0 {
		phpVersionNumber = helper.RemoveWhitespace(args[0])
	}

	if !helper.IsValidPhpVersion(phpVersionNumber) {
		fmt.Printf("Please Enter a valid version number\n\n")
		_ = cmd.Help()
		return
	}

	phpVersionCheck := fmt.Sprintf("php%v", phpVersionNumber)
	command := exec.Command("which", phpVersionCheck)
	output, err := command.Output()
	if err != nil || string(output) == "" {
		fmt.Printf("PHP v%s is not installed on your system\n", phpVersionNumber)
		_, err := helper.InstallPhp(phpVersionNumber)
		if err != nil {
			fmt.Println(err)
		}
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
}

func init() {
	rootCmd.AddCommand(useCmd)
}
