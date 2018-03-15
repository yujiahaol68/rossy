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
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	wg      sync.WaitGroup
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "rossy",
	Short: "A CLI that can remind you to read your favorite feed",
	Long:  ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.rossy.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.SetVersionTemplate(version)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		dataDir := filepath.Join(home, "rossy_data")
		if _, e := os.Stat(dataDir); os.IsNotExist(e) {
			err = os.Mkdir(dataDir, os.ModePerm)
			if err != nil {
				log.Fatalln(err)
				fmt.Println("Permission deny when try to create rossy_data Dir under $HOME")
				os.Exit(1)
			}
		}

		defaultCfg := []byte(fmt.Sprintf("dataDir: %s", dataDir))
		err = ioutil.WriteFile(filepath.Join(home, ".rossy.yaml"), defaultCfg, 0644)
		if err != nil {
			log.Fatalln(err)
		}

		// Search config in home directory with name ".rossy" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".rossy")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("If you have .rossy.yaml config file. Please make sure it is under your $HOME dir or you need to specify its path by --config <PATH>")
		fmt.Println("Config file name must be .rossy.yaml")
		os.Exit(1)
	}
}
