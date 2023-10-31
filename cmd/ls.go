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
	Use:   "ls",
	Short: "List all installed PHP versions on system",
	Long:  `It will list all the installed versions of PHP on your systems. `,
	Run: func(cmd *cobra.Command, args []string) {
		listCommand := exec.Command("update-alternatives", "--list", "php")
		output, err := listCommand.Output()
		if err != nil {
			fmt.Println("Failed to execute the command")
			return
		}
		installedPhpVersions := strings.Split(helper.RemoveWhitespace(string(output)), "\n")

		phpVersionPattern := `php\d{1,2}\.\d{1,2}`
		phpVersionRegex := regexp.MustCompile(phpVersionPattern)

		// TODO: Add support for showing current version of PHP

		sort.Strings(installedPhpVersions)

		for _, version := range installedPhpVersions {
			match := phpVersionRegex.FindString(version)
			fmt.Printf(" %s\n", match)
		}
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
}
