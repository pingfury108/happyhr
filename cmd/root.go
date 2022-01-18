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
	"path"
	"runtime"

	_ "happyhr/controllers"
	_ "happyhr/db"
	"happyhr/router"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var cfgFile string
var logLevel string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "happyhr",
	Short: "A happy hr",
	Long:  `A happy hr.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	Run: func(cmd *cobra.Command, args []string) {
		r := router.Route
		r.Run(":8085")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig, initLog)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./happyhr.toml)")
	rootCmd.PersistentFlags().StringVar(&logLevel, "logLevel", "Info", "Loglevel")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {

		// Search config in home directory with name ".happyhr" (without extension).
		viper.SetConfigType("toml")
		viper.SetConfigName("happyhr")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

func initLog() {
	formatter := &log.TextFormatter{
		FullTimestamp: true,
	}

	switch logLevel {
	case "Debug":
		log.SetLevel(log.DebugLevel)
		formatter.CallerPrettyfier = setLog
		log.SetReportCaller(true)
	case "Info":
		log.SetLevel(log.InfoLevel)
	case "Error":
		log.SetLevel(log.ErrorLevel)
		formatter.CallerPrettyfier = setLog
		log.SetReportCaller(true)
	case "Fatal":
		log.SetLevel(log.FatalLevel)
		formatter.CallerPrettyfier = setLog
		log.SetReportCaller(true)
	}
	log.SetFormatter(formatter)
}

func setLog(f *runtime.Frame) (string, string) {
	filename := path.Base(f.File)
	return f.Function, fmt.Sprintf("%s:%d", filename, f.Line)
}
