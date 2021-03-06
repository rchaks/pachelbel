// Copyright © 2017 ben dewan <benj.dewan@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "pachelbel",
	Short: "Compose deployment provisioner and deprovisioner",
	Long: `pachelbel provisions and deprovisions deployments of compose resources
in an idempotent fashion.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	RootCmd.PersistentFlags().StringP("api-key", "a", "",
		`A Compose API Key to use for authentication with the API.
			 If this flag is not set pachelbel will attempt to
			 read the COMPOSE_API_KEY environment variable.`)
	RootCmd.PersistentFlags().StringP("log-file", "l", "",
		`If specified pachelbel will enable logging for
				all Compose API requests and write them, as well
				as the reponses, to the specified log file`)
	RootCmd.PersistentFlags().BoolP("dry-run", "n", false,
		`Simulate a pachelbel command run without making any
				 real changes.`)

	viper.BindPFlag("dry-run", RootCmd.PersistentFlags().Lookup("dry-run"))
	viper.BindPFlag("log-file", RootCmd.PersistentFlags().Lookup("log-file"))
	viper.BindPFlag("api-key", RootCmd.PersistentFlags().Lookup("api-key"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.BindEnv("api-key", "COMPOSE_API_KEY")

	viper.AutomaticEnv() // read in environment variables that match
}
