package delete

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


import (
"fmt"
"github.com/arielireni/flotta-dev-cli/internal/resources"
"github.com/spf13/cobra"
)

// edgedeviceCmd represents the edgedevice command
var edgedeviceCmd = &cobra.Command{
	Use:   "edgedevice",
	Short: "Adds a new edgedevice",
	Long: `Adds a new edgedevice"`,
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

		err = device.Unregister()
		if err != nil {
			fmt.Printf("Unregister failed: %v\n", err)
			return
		}

		err = device.Remove()
		if err != nil {
			fmt.Printf("Remove failed: %v\n", err)
			return
		}

		fmt.Printf("edgedevice '%v' was deleted \n", device.GetName())
	},
}

func init() {
	// subcommand of delete
	deleteCmd.AddCommand(edgedeviceCmd)
}
