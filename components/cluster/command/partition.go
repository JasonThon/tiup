package command

import (
	"github.com/pingcap/tiup/pkg/cliutil"
	"github.com/spf13/cobra"
)

func newPartitionCmd() *cobra.Command {
	cmd := &cobra.Command {
		Use: 	"partition <cluster-name> <instance-name>",
		Short: 	"Move an instance with name into a stand-alone network partition",
		RunE:  func(cmd *cobra.Command, args []string) error {
			shouldContinue, err := cliutil.CheckCommandArgsAndMayPrintHelp(cmd, args, 2)
			if err != nil {
				return err
			}
			if !shouldContinue {
				return nil
			}
			clusterName := args[0]
			instanceName := args[1]
		},
	}
	return cmd
}

func newRemovePartitionCmd() *cobra.Command {
	cmd := &cobra.Command {
		Use: 	"departition <cluster-name> <instance-name>",
		Short: 	"Remove a instance from partition",
		RunE: func(cmd *cobra.Command, args []string) error {

		},
	}
}
