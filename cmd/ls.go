package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
	"regexp"
	"sort"
	"strings"

	"github.com/mukul98s/boggart/helper"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Args:  cobra.ExactArgs(0),
	Use:   "ls",
	Short: "List all installed PHP versions on system",
	Long:  `It will list all the installed versions of PHP on your systems. `,
	Run:   runLsCommand,
}

func runLsCommand(cmd *cobra.Command, args []string) {
	phpVersionPattern := `(?mi)php\s{0,1}\d{1,2}\.\d{1,2}`
	phpVersionRegex := regexp.MustCompile(phpVersionPattern)

	phpVersionCommand := exec.Command("php", "--version")
	output, err := phpVersionCommand.Output()
	if err != nil || len(output) == 0 {
		fmt.Println("No version of PHP is installed on your system!")
		cmd.ErrOrStderr()
		return
	}

	currentPhpVersion := phpVersionRegex.FindString(string(output))
	currentPhpVersion = strings.Join(strings.Split(strings.ToLower(currentPhpVersion), " "), "")

	listCommand := exec.Command("update-alternatives", "--list", "php")
	output, err = listCommand.Output()
	if err != nil {
		fmt.Println("Failed to execute the command")
		return
	}

	installedPhpVersions := strings.Split(helper.RemoveWhitespace(string(output)), "\n")
	sort.Strings(installedPhpVersions)

	for _, version := range installedPhpVersions {
		match := phpVersionRegex.FindString(version)
		isActiveVersion := " "
		if match == currentPhpVersion {
			isActiveVersion = "$"
		}

		fmt.Printf("%v %s\n", isActiveVersion, match)
	}
}

func init() {
	rootCmd.AddCommand(lsCmd)
}
