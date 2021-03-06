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
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/apeiron242/todo/db"
	"github.com/spf13/cobra"
)

// detailsCmd represents the details command
var detailsCmd = &cobra.Command{
	Use:   "details",
	Short: "Show details of a single Todo",
	Run: func(cmd *cobra.Command, args []string) {
		getDetails()
	},
}

func getDetails() {
	var title string
	var details string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("\nTitle : ")
	scanner.Scan()
	title = scanner.Text()
	db, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	row, err := db.Query("SELECT details FROM data WHERE title = ?", title)
	if err != nil {
		log.Fatal(err)
	}

	for row.Next() {
		row.Scan(&details)
	}

	if len(details) == 0 {
		log.Fatalf("Title %v doens't exist", title)
	}

	fmt.Printf("Details : %v\n\n", details)
}

func init() {
	rootCmd.AddCommand(detailsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// detailsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// detailsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
