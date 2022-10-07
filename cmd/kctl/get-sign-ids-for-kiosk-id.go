// Code generated. DO NOT EDIT.

package main

import (
	"github.com/spf13/cobra"

	"fmt"

	"github.com/golang/protobuf/jsonpb"

	"io"

	kioskpb "github.com/googleapis/kiosk/rpc"

	"os"
)

var GetSignIdsForKioskIdInput kioskpb.GetSignIdForKioskIdRequest

var GetSignIdsForKioskIdFromFile string

func init() {
	DisplayServiceCmd.AddCommand(GetSignIdsForKioskIdCmd)

	GetSignIdsForKioskIdCmd.Flags().Int32Var(&GetSignIdsForKioskIdInput.KioskId, "kiosk_id", 0, "Required. Kiosk id.")

	GetSignIdsForKioskIdCmd.Flags().StringVar(&GetSignIdsForKioskIdFromFile, "from_file", "", "Absolute path to JSON file containing request payload")

}

var GetSignIdsForKioskIdCmd = &cobra.Command{
	Use:   "get-sign-ids-for-kiosk-id",
	Short: "Get signs that should be displayed on a kiosk....",
	Long:  "Get signs that should be displayed on a kiosk. Streams.",
	PreRun: func(cmd *cobra.Command, args []string) {

		if GetSignIdsForKioskIdFromFile == "" {

			cmd.MarkFlagRequired("kiosk_id")

		}

	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {

		in := os.Stdin
		if GetSignIdsForKioskIdFromFile != "" {
			in, err = os.Open(GetSignIdsForKioskIdFromFile)
			if err != nil {
				return err
			}
			defer in.Close()

			err = jsonpb.Unmarshal(in, &GetSignIdsForKioskIdInput)
			if err != nil {
				return err
			}

		}

		if Verbose {
			printVerboseInput("Display", "GetSignIdsForKioskId", &GetSignIdsForKioskIdInput)
		}
		resp, err := DisplayClient.GetSignIdsForKioskId(ctx, &GetSignIdsForKioskIdInput)
		if err != nil {
			return err
		}

		var item *kioskpb.GetSignIdResponse
		for {
			item, err = resp.Recv()
			if err != nil {
				break
			}

			if Verbose {
				fmt.Print("Output: ")
			}
			printMessage(item)
		}

		if err == io.EOF {
			return nil
		}

		return err
	},
}
