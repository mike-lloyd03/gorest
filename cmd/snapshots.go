/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// snapshotsCmd represents the snapshots command
var snapshotsCmd = &cobra.Command{
	Use:   "snapshots",
	Short: "Lists all available snapshots in a repo",
	Long: `List all snapshots currently in repository. By default,
this command will show all snapshots in all configured repositories.
To view a specific repository, use the --repo flag.`,
	Run: func(cmd *cobra.Command, args []string) {
		config, err := ReadConfig("repos.yaml")
		if err != nil {
			fmt.Println("error")
		}
		for _, repo := range config.Repos {
			// fmt.Println(repo.Name)
			// fmt.Println(repo.RepoPath)
			// fmt.Println(repo.PasswordFile)
			cmd := exec.Command("restic", "snapshots", "--repo", repo.RepoPath, "--password-file", repo.PasswordFile)
			cmd.Stdout = os.Stdout
			// fmt.Println(cmd.Stdout)
			cmd.Run()
		}
	},
}

func init() {
	rootCmd.AddCommand(snapshotsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// snapshotsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// snapshotsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func getSnapshots() {

}
