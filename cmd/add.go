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
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/yujiahaol68/rossy/config"
	"github.com/yujiahaol68/rossy/feed"
	"github.com/yujiahaol68/rossy/logger"
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

		cc := new(feed.CmdController)
		tunnel := make(chan *logger.Message)

		wg.Add(1)
		go func() {
			defer wg.Done()
			for m := range tunnel {
				m.ShowInCmd()
			}
		}()

		s, err := cc.AddNewSource(tunnel, args...)
		if err != nil {
			log.Fatalln(err)
		}

		wg.Wait()

		err = feed.SaveAsJSON(s, filepath.Join(config.Get("dataDir"), "source.json"))
		if err != nil {
			log.Fatalln(err)
			os.Exit(1)
		}

		fmt.Printf("\nNew feed source has saved successfully!\n")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
