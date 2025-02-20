// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/microsoft/fabrikate/logger"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "The version of Fabrikate being used",
	Long:  "The version of Fabrikate being used",
	Run: func(cmd *cobra.Command, args []string) {
		PrintVersion()
	},
}

// PrintVersion prints the current version of Fabrikate being used.
func PrintVersion() {
	logger.Info("fab version 0.16.2")
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
