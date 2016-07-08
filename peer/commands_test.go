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
	"testing"

	"github.com/spf13/cobra"
)

func TestBuildCommandHeirachySingleCommand(t *testing.T) {
	command := &cobra.Command{Use: "command"}

	commandHeirachy := buildCommandHeirachyFromCobraCommand(command)

	assertSlicesEqual(t, []string{"command"}, commandHeirachy)
}

func TestBuildCommandHeirachyNilCommand(t *testing.T) {
	var command *cobra.Command

	commandHeirachy := buildCommandHeirachyFromCobraCommand(command)

	assertSlicesEqual(t, []string{}, commandHeirachy)
}

func TestBuildCommandHeirachyTwoCommands(t *testing.T) {
	rootCommand := &cobra.Command{Use: "rootcommand"}
	childCommand := &cobra.Command{Use: "childcommand"}
	rootCommand.AddCommand(childCommand)

	commandHeirachy := buildCommandHeirachyFromCobraCommand(childCommand)

	assertSlicesEqual(t, []string{"rootcommand", "childcommand"}, commandHeirachy)
}

func TestBuildCommandHeirachyMultipleCommands(t *testing.T) {
	rootCommand := &cobra.Command{Use: "rootcommand"}
	childCommand := &cobra.Command{Use: "childcommand"}
	leafCommand := &cobra.Command{Use: "leafCommand"}
	rootCommand.AddCommand(childCommand)
	childCommand.AddCommand(leafCommand)

	commandHeirachy := buildCommandHeirachyFromCobraCommand(leafCommand)

	assertSlicesEqual(t, []string{"rootcommand", "childcommand", "leafCommand"}, commandHeirachy)
}

func assertSlicesEqual(t *testing.T, expected []string, actual []string) {
	if len(expected) != len(actual) {
		t.Errorf("Incorrect length given. Expected %v, got %v", len(expected), len(actual))
	}

	for i := range expected {
		if expected[i] != actual[i] {
			t.Errorf("Element %v not correct. Expected %v, got %v", i, expected[i], actual[i])
		}
	}
}
