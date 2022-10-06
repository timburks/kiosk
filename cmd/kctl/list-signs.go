// Code generated. DO NOT EDIT.

package main

import (
	"github.com/spf13/cobra"

	emptypb "google.golang.org/protobuf/types/known/emptypb"

	"fmt"
)

var ListSignsInput emptypb.Empty

func init() {
	DisplayServiceCmd.AddCommand(ListSignsCmd)

}

var ListSignsCmd = &cobra.Command{
	Use:   "list-signs",
	Short: "List active signs.",
	Long:  "List active signs.",
	PreRun: func(cmd *cobra.Command, args []string) {

	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {

		if Verbose {
			printVerboseInput("Display", "ListSigns", &ListSignsInput)
		}
		resp, err := DisplayClient.ListSigns(ctx, &ListSignsInput)
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
