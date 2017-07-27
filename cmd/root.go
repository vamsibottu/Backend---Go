// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
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
	"os"
	"path"

	start "github.com/Datamigration/cmd/Start"
	"github.com/Datamigration/cmd/migrate"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

//persistent flags
var (
	verbose bool
	cfgFile string
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "places view",
	Short: "get places and shows in json formate",
	Long:  `migrate the formated data to postgres and connected to front end`,
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.lub-api.yaml)")
	RootCmd.PersistentFlags().BoolVar(&verbose, "verbose", false, "more verbose error reporting")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	RootCmd.AddCommand(migrate.MigrateCmd)
	RootCmd.AddCommand(start.StartCmd)

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if wd, err := os.Getwd(); err == nil {
		viper.SetConfigFile(path.Dir(wd) + `/config.yaml`)
	} else {
		log.Fatalf("config.yaml missing!")
	}
	viper.SetEnvPrefix("LUB_API")
	for _, envVar := range viper.GetStringSlice("env") {
		viper.BindEnv(envVar)
	}
	viper.AutomaticEnv() // read in environment variables that match
	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
