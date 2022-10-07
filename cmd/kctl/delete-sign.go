// Code generated. DO NOT EDIT.

package main

import (
	"github.com/spf13/cobra"

	"github.com/golang/protobuf/jsonpb"

	kioskpb "github.com/googleapis/kiosk/rpc"

	"os"
)

var DeleteSignInput kioskpb.DeleteSignRequest

var DeleteSignFromFile string

func init() {
	DisplayServiceCmd.AddCommand(DeleteSignCmd)

	DeleteSignCmd.Flags().Int32Var(&DeleteSignInput.Id, "id", 0, "Required. Sign id.")

	DeleteSignCmd.Flags().StringVar(&DeleteSignFromFile, "from_file", "", "Absolute path to JSON file containing request payload")

}

var DeleteSignCmd = &cobra.Command{
	Use:   "delete-sign",
	Short: "Delete a sign.",
	Long:  "Delete a sign.",
	PreRun: func(cmd *cobra.Command, args []string) {

		if DeleteSignFromFile == "" {

			cmd.MarkFlagRequired("id")

		}

	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {

		in := os.Stdin
		if DeleteSignFromFile != "" {
			in, err = os.Open(DeleteSignFromFile)
			if err != nil {
				return err
			}
			defer in.Close()

			err = jsonpb.Unmarshal(in, &DeleteSignInput)
			if err != nil {
				return err
			}

		}

		if Verbose {
			printVerboseInput("Display", "DeleteSign", &DeleteSignInput)
		}
		err = DisplayClient.DeleteSign(ctx, &DeleteSignInput)
		if err != nil {
			return err
		}

		return err
	},
}
