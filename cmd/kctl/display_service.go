// Code generated. DO NOT EDIT.

package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"google.golang.org/api/option"
	"google.golang.org/grpc"

	gapic "github.com/googleapis/kiosk/gapic"
)

var DisplayConfig *viper.Viper
var DisplayClient *gapic.DisplayClient
var DisplaySubCommands []string = []string{
	"create-kiosk",
	"list-kiosks",
	"get-kiosk",
	"delete-kiosk",
	"create-sign",
	"list-signs",
	"get-sign",
	"delete-sign",
	"set-sign-id-for-kiosk-ids",
	"get-sign-id-for-kiosk-id",
	"get-sign-ids-for-kiosk-id",
}

func init() {
	rootCmd.AddCommand(DisplayServiceCmd)

	DisplayConfig = viper.New()
	DisplayConfig.SetEnvPrefix("KCTL_DISPLAY")
	DisplayConfig.AutomaticEnv()

	DisplayServiceCmd.PersistentFlags().Bool("insecure", false, "Make insecure client connection. Or use KCTL_DISPLAY_INSECURE. Must be used with \"address\" option")
	DisplayConfig.BindPFlag("insecure", DisplayServiceCmd.PersistentFlags().Lookup("insecure"))
	DisplayConfig.BindEnv("insecure")

	DisplayServiceCmd.PersistentFlags().String("address", "", "Set API address used by client. Or use KCTL_DISPLAY_ADDRESS.")
	DisplayConfig.BindPFlag("address", DisplayServiceCmd.PersistentFlags().Lookup("address"))
	DisplayConfig.BindEnv("address")

	DisplayServiceCmd.PersistentFlags().String("token", "", "Set Bearer token used by the client. Or use KCTL_DISPLAY_TOKEN.")
	DisplayConfig.BindPFlag("token", DisplayServiceCmd.PersistentFlags().Lookup("token"))
	DisplayConfig.BindEnv("token")

	DisplayServiceCmd.PersistentFlags().String("api_key", "", "Set API Key used by the client. Or use KCTL_DISPLAY_API_KEY.")
	DisplayConfig.BindPFlag("api_key", DisplayServiceCmd.PersistentFlags().Lookup("api_key"))
	DisplayConfig.BindEnv("api_key")
}

var DisplayServiceCmd = &cobra.Command{
	Use:       "display",
	Short:     "The Kiosk Display service.",
	Long:      "The Kiosk Display service.",
	ValidArgs: DisplaySubCommands,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) (err error) {
		var opts []option.ClientOption

		address := DisplayConfig.GetString("address")
		if address != "" {
			opts = append(opts, option.WithEndpoint(address))
		}

		if DisplayConfig.GetBool("insecure") {
			if address == "" {
				return fmt.Errorf("Missing address to use with insecure connection")
			}

			conn, err := grpc.Dial(address, grpc.WithInsecure())
			if err != nil {
				return err
			}
			opts = append(opts, option.WithGRPCConn(conn))
		}

		if token := DisplayConfig.GetString("token"); token != "" {
			opts = append(opts, option.WithTokenSource(oauth2.StaticTokenSource(
				&oauth2.Token{
					AccessToken: token,
					TokenType:   "Bearer",
				})))
		}

		if key := DisplayConfig.GetString("api_key"); key != "" {
			opts = append(opts, option.WithAPIKey(key))
		}

		DisplayClient, err = gapic.NewDisplayClient(ctx, opts...)
		return
	},
}
