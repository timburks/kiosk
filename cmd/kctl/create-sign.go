// Code generated. DO NOT EDIT.

package main

import (
	"github.com/spf13/cobra"

	"fmt"

	"github.com/golang/protobuf/jsonpb"

	kioskpb "github.com/googleapis/kiosk/rpc"

	"os"
)

var CreateSignInput kioskpb.Sign

var CreateSignFromFile string

func init() {
	DisplayServiceCmd.AddCommand(CreateSignCmd)

	CreateSignCmd.Flags().StringVar(&CreateSignInput.Name, "name", "", "Required. Name of sign.")

	CreateSignCmd.Flags().StringVar(&CreateSignInput.Text, "text", "", "Text to display.")

	CreateSignCmd.Flags().BytesHexVar(&CreateSignInput.Image, "image", []byte{}, "Image to display.")

	CreateSignCmd.Flags().StringVar(&CreateSignFromFile, "from_file", "", "Absolute path to JSON file containing request payload")

}

var CreateSignCmd = &cobra.Command{
	Use:   "create-sign",
	Short: "Create a sign and enroll the sign for sign...",
	Long:  "Create a sign and enroll the sign for sign display.",
	PreRun: func(cmd *cobra.Command, args []string) {

		if CreateSignFromFile == "" {

			cmd.MarkFlagRequired("name")

		}

	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {

		in := os.Stdin
		if CreateSignFromFile != "" {
			in, err = os.Open(CreateSignFromFile)
			if err != nil {
				return err
			}
			defer in.Close()

			err = jsonpb.Unmarshal(in, &CreateSignInput)
			if err != nil {
				return err
			}

		}

		if Verbose {
			printVerboseInput("Display", "CreateSign", &CreateSignInput)
		}
		resp, err := DisplayClient.CreateSign(ctx, &CreateSignInput)
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
