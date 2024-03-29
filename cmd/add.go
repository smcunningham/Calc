/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

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
	"strconv"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add numbers together",
	Long: `This command takes any amount of numbers, and adds them together.
	Nothing too complicated`,
	Run: func(cmd *cobra.Command, args []string) {
		// get the flag value, its default value is false
		status, _ := cmd.Flags().GetBool("float")

		if status {
			addFloat(args)
		} else {
			addInt(args)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().BoolP("float", "f", false, "Add Floating Numbers")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func addInt(args []string) {
	var sum int

	// iterate over the argunents
	// the first return value is index of args, we can omit it using _
	for _, val := range args {
		temp, err := strconv.Atoi(val)
		if err != nil {
			fmt.Println(err)
		}
		sum += temp
	}
	fmt.Printf("Addition of numbers %s is %d", args, sum)
}

func addFloat(args []string) {
	var sum float64

	// iterate over the argunents
	for _, val := range args {
		temp, err := strconv.ParseFloat(val, 64)
		if err != nil {
			fmt.Println(err)
		}
		sum += temp
	}
	fmt.Printf("Addition of numbers %s is %f", args, sum)
}
