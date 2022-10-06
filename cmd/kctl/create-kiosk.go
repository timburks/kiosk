// Code generated. DO NOT EDIT.

package main

import (
	"github.com/spf13/cobra"

	"fmt"

	"github.com/golang/protobuf/jsonpb"

	kioskpb "github.com/googleapis/kiosk/rpc"

	latlngpb "google.golang.org/genproto/googleapis/type/latlng"

	"os"

	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

var CreateKioskInput kioskpb.Kiosk

var CreateKioskFromFile string

func init() {
	DisplayServiceCmd.AddCommand(CreateKioskCmd)

	CreateKioskInput.Size = new(kioskpb.ScreenSize)

	CreateKioskInput.Location = new(latlngpb.LatLng)

	CreateKioskInput.CreateTime = new(timestamppb.Timestamp)

	CreateKioskCmd.Flags().Int32Var(&CreateKioskInput.Id, "id", 0, "Output only.")

	CreateKioskCmd.Flags().StringVar(&CreateKioskInput.Name, "name", "", "Required.")

	CreateKioskCmd.Flags().Int32Var(&CreateKioskInput.Size.Width, "size.width", 0, "")

	CreateKioskCmd.Flags().Int32Var(&CreateKioskInput.Size.Height, "size.height", 0, "")

	CreateKioskCmd.Flags().Float64Var(&CreateKioskInput.Location.Latitude, "location.latitude", 0.0, "The latitude in degrees. It must be in the range...")

	CreateKioskCmd.Flags().Float64Var(&CreateKioskInput.Location.Longitude, "location.longitude", 0.0, "The longitude in degrees. It must be in the range...")

	CreateKioskCmd.Flags().Int64Var(&CreateKioskInput.CreateTime.Seconds, "create_time.seconds", 0, "Represents seconds of UTC time since Unix epoch ...")

	CreateKioskCmd.Flags().Int32Var(&CreateKioskInput.CreateTime.Nanos, "create_time.nanos", 0, "Non-negative fractions of a second at nanosecond...")

	CreateKioskCmd.Flags().StringVar(&CreateKioskFromFile, "from_file", "", "Absolute path to JSON file containing request payload")

}

var CreateKioskCmd = &cobra.Command{
	Use:   "create-kiosk",
	Short: "Create a kiosk. This enrolls the kiosk for sign...",
	Long:  "Create a kiosk. This enrolls the kiosk for sign display.",
	PreRun: func(cmd *cobra.Command, args []string) {

		if CreateKioskFromFile == "" {

		}

	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {

		in := os.Stdin
		if CreateKioskFromFile != "" {
			in, err = os.Open(CreateKioskFromFile)
			if err != nil {
				return err
			}
			defer in.Close()

			err = jsonpb.Unmarshal(in, &CreateKioskInput)
			if err != nil {
				return err
			}

		}

		if Verbose {
			printVerboseInput("Display", "CreateKiosk", &CreateKioskInput)
		}
		resp, err := DisplayClient.CreateKiosk(ctx, &CreateKioskInput)
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
