// Code generated. DO NOT EDIT.

package main

import (
	"github.com/spf13/cobra"

	"github.com/golang/protobuf/jsonpb"

	kioskpb "github.com/googleapis/kiosk/rpc"

	"os"
)

var SetSignIdForKioskIdsInput kioskpb.SetSignIdForKioskIdsRequest

var SetSignIdForKioskIdsFromFile string

func init() {
	DisplayServiceCmd.AddCommand(SetSignIdForKioskIdsCmd)

	SetSignIdForKioskIdsCmd.Flags().Int32SliceVar(&SetSignIdForKioskIdsInput.KioskIds, "kiosk_ids", []int32{}, "Required. Kiosk ids.")

	SetSignIdForKioskIdsCmd.Flags().Int32Var(&SetSignIdForKioskIdsInput.SignId, "sign_id", 0, "Required. Sign id.")

	SetSignIdForKioskIdsCmd.Flags().StringVar(&SetSignIdForKioskIdsFromFile, "from_file", "", "Absolute path to JSON file containing request payload")

}

var SetSignIdForKioskIdsCmd = &cobra.Command{
	Use:   "set-sign-id-for-kiosk-ids",
	Short: "Set a sign for display on one or more kiosks.",
	Long:  "Set a sign for display on one or more kiosks.",
	PreRun: func(cmd *cobra.Command, args []string) {

		if SetSignIdForKioskIdsFromFile == "" {

			cmd.MarkFlagRequired("kiosk_ids")

			cmd.MarkFlagRequired("sign_id")

		}

	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {

		in := os.Stdin
		if SetSignIdForKioskIdsFromFile != "" {
			in, err = os.Open(SetSignIdForKioskIdsFromFile)
			if err != nil {
				return err
			}
			defer in.Close()

			err = jsonpb.Unmarshal(in, &SetSignIdForKioskIdsInput)
			if err != nil {
				return err
			}

		}

		if Verbose {
			printVerboseInput("Display", "SetSignIdForKioskIds", &SetSignIdForKioskIdsInput)
		}
		err = DisplayClient.SetSignIdForKioskIds(ctx, &SetSignIdForKioskIdsInput)
		if err != nil {
			return err
		}

		return err
	},
}
