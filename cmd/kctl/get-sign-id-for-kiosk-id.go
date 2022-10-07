// Code generated. DO NOT EDIT.

package main

import (
	"github.com/spf13/cobra"

	"fmt"

	"github.com/golang/protobuf/jsonpb"

	kioskpb "github.com/googleapis/kiosk/rpc"

	"os"
)

var GetSignIdForKioskIdInput kioskpb.GetSignIdForKioskIdRequest

var GetSignIdForKioskIdFromFile string

func init() {
	DisplayServiceCmd.AddCommand(GetSignIdForKioskIdCmd)

	GetSignIdForKioskIdCmd.Flags().Int32Var(&GetSignIdForKioskIdInput.KioskId, "kiosk_id", 0, "Required. Kiosk id.")

	GetSignIdForKioskIdCmd.Flags().StringVar(&GetSignIdForKioskIdFromFile, "from_file", "", "Absolute path to JSON file containing request payload")

}

var GetSignIdForKioskIdCmd = &cobra.Command{
	Use:   "get-sign-id-for-kiosk-id",
	Short: "Get the sign that should be displayed on a kiosk.",
	Long:  "Get the sign that should be displayed on a kiosk.",
	PreRun: func(cmd *cobra.Command, args []string) {

		if GetSignIdForKioskIdFromFile == "" {

			cmd.MarkFlagRequired("kiosk_id")

		}

	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {

		in := os.Stdin
		if GetSignIdForKioskIdFromFile != "" {
			in, err = os.Open(GetSignIdForKioskIdFromFile)
			if err != nil {
				return err
			}
			defer in.Close()

			err = jsonpb.Unmarshal(in, &GetSignIdForKioskIdInput)
			if err != nil {
				return err
			}

		}

		if Verbose {
			printVerboseInput("Display", "GetSignIdForKioskId", &GetSignIdForKioskIdInput)
		}
		resp, err := DisplayClient.GetSignIdForKioskId(ctx, &GetSignIdForKioskIdInput)
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
