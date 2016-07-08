/*
Copyright IBM Corp. 2016 All Rights Reserved.

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

package main

import (
	"fmt"
	"sync"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var initialiseOnce sync.Once

// GetCommands gets the Command struct that should be utilised by Cobra
func GetCommands() *cobra.Command {
	initialiseOnce.Do(initialiseCommands)
	return mainCmd
}

func initialiseCommands() {
	mainCmd.AddCommand(nodeCmd)
	mainCmd.AddCommand(networkCmd)
	mainCmd.AddCommand(chaincodeCmd)

	mainFlags := mainCmd.PersistentFlags()
	mainFlags.String("logging-level", "", "Default logging level and overrides, see core.yaml for full syntax")
	viper.BindPFlag("logging_level", mainFlags.Lookup("logging-level"))
	fmt.Println("LogLevel: " + viper.GetString("logging_level"))
	testCoverProfile := ""
	mainFlags.StringVarP(&testCoverProfile, "test.coverprofile", "", "coverage.cov", "Done")

	initialiseNodeCmd()
	initialiseNetworkCmd()
	initialiseChainCodeCmd()
}

func initialiseNodeCmd() {
	nodeCmd.AddCommand(nodeStartCmd)
	nodeCmd.AddCommand(nodeStopCmd)
	nodeCmd.AddCommand(nodeStatusCmd)
	nodeStopCmd.Flags().StringVar(&stopPidFile, "stop-peer-pid-file", viper.GetString("peer.fileSystemPath"), "Location of peer pid local file, for forces kill")

	// Set the flags on the node start command.
	flags := nodeStartCmd.Flags()
	flags.BoolVarP(&chaincodeDevMode, "peer-chaincodedev", "", false, "Whether peer in chaincode development mode")

}

func initialiseNetworkCmd() {
	networkCmd.AddCommand(networkListCmd)
	networkCmd.AddCommand(networkLoginCmd)
	// Set the flags on the login command.
	networkLoginCmd.PersistentFlags().StringVarP(&loginPW, "password", "p", undefinedParamValue, "The password for user. You will be requested to enter the password if this flag is not specified.")

}

func initialiseChainCodeCmd() {
	chaincodeCmd.AddCommand(chaincodeDeployCmd)
	chaincodeCmd.AddCommand(chaincodeInvokeCmd)
	chaincodeCmd.AddCommand(chaincodeQueryCmd)

	chaincodeCmd.PersistentFlags().StringVarP(&chaincodeLang, "lang", "l", "golang", fmt.Sprintf("Language the %s is written in", chainFuncName))
	chaincodeCmd.PersistentFlags().StringVarP(&chaincodeCtorJSON, "ctor", "c", "{}", fmt.Sprintf("Constructor message for the %s in JSON format", chainFuncName))
	chaincodeCmd.PersistentFlags().StringVarP(&chaincodeAttributesJSON, "attributes", "a", "[]", fmt.Sprintf("User attributes for the %s in JSON format", chainFuncName))
	chaincodeCmd.PersistentFlags().StringVarP(&chaincodePath, "path", "p", undefinedParamValue, fmt.Sprintf("Path to %s", chainFuncName))
	chaincodeCmd.PersistentFlags().StringVarP(&chaincodeName, "name", "n", undefinedParamValue, fmt.Sprintf("Name of the chaincode returned by the deploy transaction"))
	chaincodeCmd.PersistentFlags().StringVarP(&chaincodeUsr, "username", "u", undefinedParamValue, fmt.Sprintf("Username for chaincode operations when security is enabled"))
	chaincodeCmd.PersistentFlags().StringVarP(&customIDGenAlg, "tid", "t", undefinedParamValue, fmt.Sprintf("Name of a custom ID generation algorithm (hashing and decoding) e.g. sha256base64"))

	chaincodeQueryCmd.Flags().BoolVarP(&chaincodeQueryRaw, "raw", "r", false, "If true, output the query value as raw bytes, otherwise format as a printable string")
	chaincodeQueryCmd.Flags().BoolVarP(&chaincodeQueryHex, "hex", "x", false, "If true, output the query value byte array in hexadecimal. Incompatible with --raw")

}
