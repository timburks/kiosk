// Code generated. DO NOT EDIT.

package main

import (
	"github.com/spf13/cobra"

	emptypb "google.golang.org/protobuf/types/known/emptypb"

	"fmt"
)

var ListKiosksInput emptypb.Empty

func init() {
	DisplayServiceCmd.AddCommand(ListKiosksCmd)

}

var ListKiosksCmd = &cobra.Command{
	Use:   "list-kiosks",
	Short: "List active kiosks.",
	Long:  "List active kiosks.",
	PreRun: func(cmd *cobra.Command, args []string) {

	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {

		if Verbose {
			printVerboseInput("Display", "ListKiosks", &ListKiosksInput)
		}
		resp, err := DisplayClient.ListKiosks(ctx, &ListKiosksInput)
		if err != nil {
			return err
		}

		if Verbose {
			fmt.Print("Output: ")
		}
		printMessage(resp)

		return err
	},
}
