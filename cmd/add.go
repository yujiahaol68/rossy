// Copyright © 2018 Tecker.Yu <tecker_yuknigh@163.com>
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
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [URL] [URL] ...",
	Short: "add new XML feed by URL",
	Long:  `add new XML feed by URL, can be multiple`,
	Run: func(cmd *cobra.Command, args []string) {
		for _, u := range args {
			_, err := url.ParseRequestURI(u)
			if err != nil {
				log.Fatal(err)
				fmt.Printf("Invalid URL: %s\n", u)
				os.Exit(1)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
