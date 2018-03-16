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
	"path/filepath"

	"github.com/yujiahaol68/rossy/config"
	"github.com/yujiahaol68/rossy/feed"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all rss feed source name",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		s, err := feed.ReadExistSource(filepath.Join(config.Get("dataDir"), "source.json"))
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println("Exist feed source list below:")
		for i, source := range s {
			fmt.Printf("%d. %s\n- %s\n", i+1, source.Alias, source.URL)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
