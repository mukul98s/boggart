/*
Copyright © 2023 Mukul sharma mymukul112@gmail.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
