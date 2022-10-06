// Code generated. DO NOT EDIT.

package main

import (
	"github.com/spf13/cobra"

	"github.com/golang/protobuf/jsonpb"

	kioskpb "github.com/googleapis/kiosk/rpc"

	"os"
)

var DeleteKioskInput kioskpb.DeleteKioskRequest

var DeleteKioskFromFile string

func init() {
	DisplayServiceCmd.AddCommand(DeleteKioskCmd)

	DeleteKioskCmd.Flags().Int32Var(&DeleteKioskInput.Id, "id", 0, "Required.")

	DeleteKioskCmd.Flags().StringVar(&DeleteKioskFromFile, "from_file", "", "Absolute path to JSON file containing request payload")

}

var DeleteKioskCmd = &cobra.Command{
	Use:   "delete-kiosk",
	Short: "Delete a kiosk.",
	Long:  "Delete a kiosk.",
	PreRun: func(cmd *cobra.Command, args []string) {

		if DeleteKioskFromFile == "" {

		}

	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {

		in := os.Stdin
		if DeleteKioskFromFile != "" {
			in, err = os.Open(DeleteKioskFromFile)
			if err != nil {
				return err
			}
			defer in.Close()

			err = jsonpb.Unmarshal(in, &DeleteKioskInput)
			if err != nil {
				return err
			}

		}

		if Verbose {
			printVerboseInput("Display", "DeleteKiosk", &DeleteKioskInput)
		}
		err = DisplayClient.DeleteKiosk(ctx, &DeleteKioskInput)
		if err != nil {
			return err
		}

		return err
	},
}
