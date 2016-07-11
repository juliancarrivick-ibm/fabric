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
	"strings"

	"github.com/hyperledger/fabric/core"
	"github.com/spf13/cobra"
)

// getPeerCommandFromCobraCommand retreives the peer command from the cobra command struct.
// i.e. for a command of `peer node start`, this should return "node"
// For invalid commands (i.e. the main/root command and nil commands) this will return an empty string
func getPeerCommandFromCobraCommand(command *cobra.Command) string {
	var commandName string

	if command == nil {
		return commandName
	}

	if peerCommand, ok := findChildOfRootCommand(command); ok {
		commandName = peerCommand.Name()
	}

	return commandName
}

func findChildOfRootCommand(command *cobra.Command) (*cobra.Command, bool) {
	for command.HasParent() {
		if !command.Parent().HasParent() {
			return command, true
		}

		command = command.Parent()
	}

	return nil, false
}

// The main command describes the service and
// defaults to printing the help message.
var mainCmd = &cobra.Command{
	Use: "peer",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// Init Crypt here
		// Init Logging here
		return core.CacheConfiguration()
	},
}

var nodeCmd = &cobra.Command{
	Use:   nodeFuncName,
	Short: fmt.Sprintf("%s specific commands.", nodeFuncName),
	Long:  fmt.Sprintf("%s specific commands.", nodeFuncName),
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		core.LoggingInit(nodeFuncName)
	},
}

var nodeStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the node.",
	Long:  `Starts a node that interacts with the network.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return serve(args)
	},
}

var nodeStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Returns status of the node.",
	Long:  `Returns the status of the running node.`,
	Run: func(cmd *cobra.Command, args []string) {
		status()
	},
}

var nodeStopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stops the running node.",
	Long:  `Stops the running node, disconnecting from the network.`,
	Run: func(cmd *cobra.Command, args []string) {
		stop()
	},
}

var networkCmd = &cobra.Command{
	Use:   networkFuncName,
	Short: fmt.Sprintf("%s specific commands.", networkFuncName),
	Long:  fmt.Sprintf("%s specific commands.", networkFuncName),
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		core.LoggingInit(networkFuncName)
	},
}

var networkLoginCmd = &cobra.Command{
	Use:   "login",
	Short: "Logs in user to CLI.",
	Long:  `Logs in the local user to CLI. Must supply username as a parameter.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return networkLogin(args)
	},
}

var networkListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "Lists all network peers.",
	Long:    `Returns a list of all existing network connections for the target peer node, includes both validating and non-validating peers.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return networkList()
	},
}

var chaincodeCmd = &cobra.Command{
	Use:   chainFuncName,
	Short: fmt.Sprintf("%s specific commands.", chainFuncName),
	Long:  fmt.Sprintf("%s specific commands.", chainFuncName),
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		core.LoggingInit(chainFuncName)
	},
}

var chaincodePathArgumentSpecifier = fmt.Sprintf("%s_PATH", strings.ToUpper(chainFuncName))

var chaincodeDeployCmd = &cobra.Command{
	Use:       "deploy",
	Short:     fmt.Sprintf("Deploy the specified %s to the network.", chainFuncName),
	Long:      fmt.Sprintf(`Deploy the specified %s to the network.`, chainFuncName),
	ValidArgs: []string{"1"},
	RunE: func(cmd *cobra.Command, args []string) error {
		return chaincodeDeploy(cmd, args)
	},
}

var chaincodeInvokeCmd = &cobra.Command{
	Use:       "invoke",
	Short:     fmt.Sprintf("Invoke the specified %s.", chainFuncName),
	Long:      fmt.Sprintf(`Invoke the specified %s.`, chainFuncName),
	ValidArgs: []string{"1"},
	RunE: func(cmd *cobra.Command, args []string) error {
		return chaincodeInvoke(cmd, args)
	},
}

var chaincodeQueryCmd = &cobra.Command{
	Use:       "query",
	Short:     fmt.Sprintf("Query using the specified %s.", chainFuncName),
	Long:      fmt.Sprintf(`Query using the specified %s.`, chainFuncName),
	ValidArgs: []string{"1"},
	RunE: func(cmd *cobra.Command, args []string) error {
		return chaincodeQuery(cmd, args)
	},
}
