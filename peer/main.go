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
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/op/go-logging"
	"github.com/spf13/viper"

	_ "net/http/pprof"

	"github.com/hyperledger/fabric/core/crypto"
)

// Constants go here.
const (
	nodeFuncName        = "node"
	networkFuncName     = "network"
	chainFuncName       = "chaincode"
	cmdRoot             = "core"
	undefinedParamValue = ""
)

var logger = logging.MustGetLogger("main")

// login related variables.
var (
	loginPW string
)

// Chaincode-related variables.
var (
	chaincodeLang           string
	chaincodeCtorJSON       string
	chaincodePath           string
	chaincodeName           string
	chaincodeDevMode        bool
	chaincodeUsr            string
	chaincodeQueryRaw       bool
	chaincodeQueryHex       bool
	chaincodeAttributesJSON string
	customIDGenAlg          string
)

var (
	stopPidFile string
)

func main() {
	// For environment variables.
	viper.SetEnvPrefix(cmdRoot)
	viper.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)


	var alternativeCfgPath = os.Getenv("PEER_CFG_PATH")
	if alternativeCfgPath != "" {
		logger.Info("User defined config file path: %s", alternativeCfgPath)
		viper.AddConfigPath(alternativeCfgPath) // Path to look for the config file in
	} else {
		viper.AddConfigPath("./") // Path to look for the config file in
		// Path to look for the config file in based on GOPATH
		gopath := os.Getenv("GOPATH")
		for _, p := range filepath.SplitList(gopath) {
			peerpath := filepath.Join(p, "src/github.com/hyperledger/fabric/peer")
			viper.AddConfigPath(peerpath)
		}
	}

	// Now set the configuration file.
	viper.SetConfigName(cmdRoot) // Name of config file (without extension)

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error when reading %s config file: %s\n", cmdRoot, err))
	}

	runtime.GOMAXPROCS(viper.GetInt("peer.gomaxprocs"))

	// Init the crypto layer
	if err := crypto.Init(); err != nil {
		panic(fmt.Errorf("Failed to initialize the crypto layer: %s", err))
	}

	mainCmd := GetCommands()
	// On failure Cobra prints the usage message and error string, so we only
	// need to exit with a non-0 status
	if mainCmd.Execute() != nil {
		//os.Exit(1)
	}
	logger.Info("Exiting.....")
}
