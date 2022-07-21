/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

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

package add

import (
	"fmt"
	"github.com/project-flotta/flotta-dev-cli/internal/resources"
	"github.com/spf13/cobra"
)

// deviceCmd represents the device command
var deviceCmd = &cobra.Command{
	Use:   "device",
	Short: "Add new device",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := resources.NewClient()
		if err != nil {
			fmt.Printf("NewClient failed: %v\n", err)
			return
		}

		device, err := resources.NewEdgeDevice(client, args[0])
		if err != nil {
			fmt.Printf("NewEdgeDevice failed: %v\n", err)
			return
		}

		err = device.Register()
		if err != nil {
			// if device.Register() failed, remove the container
			err := device.Remove()
			if err != nil {
				fmt.Printf("Remove device that failed to register failed: %v\n", err)
				return
			}

			fmt.Printf("Register failed: %v\n", err)
			return
		}

		fmt.Printf("device '%v' was added \n", device.GetName())
	},
}

func init() {
	// subcommand of add
	addCmd.AddCommand(deviceCmd)
}
