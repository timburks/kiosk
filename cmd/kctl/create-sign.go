// Code generated. DO NOT EDIT.

package main

import (
	"github.com/spf13/cobra"

	"fmt"

	"github.com/golang/protobuf/jsonpb"

	kioskpb "github.com/googleapis/kiosk/rpc"

	"os"

	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

var CreateSignInput kioskpb.Sign

var CreateSignFromFile string

func init() {
	DisplayServiceCmd.AddCommand(CreateSignCmd)

	CreateSignInput.CreateTime = new(timestamppb.Timestamp)

	CreateSignCmd.Flags().Int32Var(&CreateSignInput.Id, "id", 0, "Output only.")

	CreateSignCmd.Flags().StringVar(&CreateSignInput.Name, "name", "", "Required.")

	CreateSignCmd.Flags().StringVar(&CreateSignInput.Text, "text", "", "")

	CreateSignCmd.Flags().BytesHexVar(&CreateSignInput.Image, "image", []byte{}, "")

	CreateSignCmd.Flags().Int64Var(&CreateSignInput.CreateTime.Seconds, "create_time.seconds", 0, "Represents seconds of UTC time since Unix epoch ...")

	CreateSignCmd.Flags().Int32Var(&CreateSignInput.CreateTime.Nanos, "create_time.nanos", 0, "Non-negative fractions of a second at nanosecond...")

	CreateSignCmd.Flags().StringVar(&CreateSignFromFile, "from_file", "", "Absolute path to JSON file containing request payload")

}

var CreateSignCmd = &cobra.Command{
	Use:   "create-sign",
	Short: "Create a sign. This enrolls the sign for sign...",
	Long:  "Create a sign. This enrolls the sign for sign display.",
	PreRun: func(cmd *cobra.Command, args []string) {

		if CreateSignFromFile == "" {

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
