// Code generated. DO NOT EDIT.

package main

import (
	"github.com/spf13/cobra"

	"fmt"

	"github.com/golang/protobuf/jsonpb"

	kioskpb "github.com/googleapis/kiosk/rpc"

	"os"
)

var GetSignInput kioskpb.GetSignRequest

var GetSignFromFile string

func init() {
	DisplayServiceCmd.AddCommand(GetSignCmd)

	GetSignCmd.Flags().Int32Var(&GetSignInput.Id, "id", 0, "Required. Sign id.")

	GetSignCmd.Flags().StringVar(&GetSignFromFile, "from_file", "", "Absolute path to JSON file containing request payload")

}

var GetSignCmd = &cobra.Command{
	Use:   "get-sign",
	Short: "Get a sign.",
	Long:  "Get a sign.",
	PreRun: func(cmd *cobra.Command, args []string) {

		if GetSignFromFile == "" {

			cmd.MarkFlagRequired("id")

		}

	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {

		in := os.Stdin
		if GetSignFromFile != "" {
			in, err = os.Open(GetSignFromFile)
			if err != nil {
				return err
			}
			defer in.Close()

			err = jsonpb.Unmarshal(in, &GetSignInput)
			if err != nil {
				return err
			}

		}

		if Verbose {
			printVerboseInput("Display", "GetSign", &GetSignInput)
		}
		resp, err := DisplayClient.GetSign(ctx, &GetSignInput)
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
