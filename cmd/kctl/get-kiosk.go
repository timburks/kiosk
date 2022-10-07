// Code generated. DO NOT EDIT.

package main

import (
	"github.com/spf13/cobra"

	"fmt"

	"github.com/golang/protobuf/jsonpb"

	kioskpb "github.com/googleapis/kiosk/rpc"

	"os"
)

var GetKioskInput kioskpb.GetKioskRequest

var GetKioskFromFile string

func init() {
	DisplayServiceCmd.AddCommand(GetKioskCmd)

	GetKioskCmd.Flags().Int32Var(&GetKioskInput.Id, "id", 0, "Required. Kiosk id.")

	GetKioskCmd.Flags().StringVar(&GetKioskFromFile, "from_file", "", "Absolute path to JSON file containing request payload")

}

var GetKioskCmd = &cobra.Command{
	Use:   "get-kiosk",
	Short: "Get a kiosk.",
	Long:  "Get a kiosk.",
	PreRun: func(cmd *cobra.Command, args []string) {

		if GetKioskFromFile == "" {

			cmd.MarkFlagRequired("id")

		}

	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {

		in := os.Stdin
		if GetKioskFromFile != "" {
			in, err = os.Open(GetKioskFromFile)
			if err != nil {
				return err
			}
			defer in.Close()

			err = jsonpb.Unmarshal(in, &GetKioskInput)
			if err != nil {
				return err
			}

		}

		if Verbose {
			printVerboseInput("Display", "GetKiosk", &GetKioskInput)
		}
		resp, err := DisplayClient.GetKiosk(ctx, &GetKioskInput)
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
