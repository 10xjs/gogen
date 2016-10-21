package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func cli() (*cobra.Command, error) {

	// configuration parameters
	cfgParams, err := getConfigParams()
	if err != nil {
		return nil, err
	}

	// set the root command
	rootCmd := new(cobra.Command)

	// overwrites the configuration parameters with the ones specified in the command line (if any)
	rootCmd.Flags().StringVarP(&appParams.logLevel, "logLevel", "o", cfgParams.logLevel, "Log level: panic, fatal, error, warning, info, debug")
	rootCmd.Flags().IntVarP(&appParams.quantity, "quantity", "r", cfgParams.quantity, "Number of results to return")

	rootCmd.Use = "~#PROJECT#~"
	rootCmd.Short = "~#SHORTDESCRIPTION#~"
	rootCmd.Long = `~#PROJECT#~ - ~#SHORTDESCRIPTION#~`
	rootCmd.RunE = func(cmd *cobra.Command, args []string) error {
		// check values
		err := checkParams(appParams)
		if err != nil {
			return err
		}

		// get results
		for i := 0; i < appParams.quantity; i++ {
			getResult()
		}

		return nil
	}

	// sub-command to print the version
	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "print this program version",
		Long:  `print this program version`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(ProgramVersion)
		},
	}
	rootCmd.AddCommand(versionCmd)

	return rootCmd, nil
}
