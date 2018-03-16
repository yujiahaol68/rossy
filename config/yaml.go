package config

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
)

const defaultFileName = ".rossy.yaml"

var (
	ErrFileNoExist = errors.New("Cfg File .rossy.yaml doesn't exist!")
	ErrFileFormat  = errors.New("Cfg format is <Key>: <value>")
	m              map[string]string
)

// ReadYAML just replace viper to simply read config from file, only for self-use
func ReadYAML(filePath string) error {
	if filePath == "" {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		filePath = filepath.Join(home, defaultFileName)

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
			os.Exit(1)
		}
	}

	file, err := os.Open(filePath)
	if err != nil {
		return err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	m = make(map[string]string)

	for scanner.Scan() {
		c := strings.Split(scanner.Text(), ":")
		if len(c) != 2 {
			return ErrFileFormat
		}

		m[strings.Trim(c[0], " ")] = strings.Trim(c[1], " ")
	}

	err = scanner.Err()
	return err
}

func Get(key string) string {
	return m[key]
}
