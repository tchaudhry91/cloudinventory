/*
Copyright 2019 Adobe. All rights reserved.
This file is licensed to you under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License. You may obtain a copy
of the License at http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under
the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR REPRESENTATIONS
OF ANY KIND, either express or implied. See the License for the specific language
governing permissions and limitations under the License.
*/

package cmd

import (
	"github.com/spf13/cobra"
)

// dumpCmd represents the dump command
var dumpCmd = &cobra.Command{
	Use:   "dump",
	Short: "Dumps the inventory for the given options",
}

func init() {
	rootCmd.AddCommand(dumpCmd)
	dumpCmd.PersistentFlags().StringP("filter", "f", "", "limit dump to a particular cloud service, e.g ec2/rds")
	dumpCmd.PersistentFlags().StringP("path", "p", "cloudinventory.json", "file path to dump the inventory in")

}
