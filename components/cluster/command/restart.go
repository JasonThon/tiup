// Copyright 2020 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package command

import (
	"github.com/spf13/cobra"
)

func newRestartClusterCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "restart <cluster-name>",
		Short: "Restart a TiDB cluster",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return cmd.Help()
			}

			if err := validRoles(gOpt.Roles); err != nil {
				return err
			}

			clusterName := args[0]
			teleCommand = append(teleCommand, scrubClusterName(clusterName))

			return manager.RestartCluster(clusterName, gOpt)
		},
	}

	cmd.Flags().StringSliceVarP(&gOpt.Roles, "role", "R", nil, "Only restart specified roles")
	cmd.Flags().StringSliceVarP(&gOpt.Nodes, "node", "N", nil, "Only restart specified nodes")

	return cmd
}

func newRestartInstanceCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use: "restart <cluster-name> <instance-names>",
		Short: "Restart TiDB/TiKV/PD instances in a cluster",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) <= 1 {
				return cmd.Help()
			}
			clusterName := args[0]
			instanceNames := args[1:]
			teleCommand = append(teleCommand, scrubClusterName(clusterName))

			return manager.RestartInstances(clusterName, instanceNames, gOpt)
		},
	}
	cmd.Flags().StringSliceVarP(&gOpt.Roles, "role", "R", nil, "Only restart specified roles")
	cmd.Flags().StringSliceVarP(&gOpt.Nodes, "node", "N", nil, "Only restart specified nodes")

	return cmd
}
