package cmd

import (
	"os"

	"github.com/Yakiyo/go-template/meta"
	cc "github.com/ivanpirog/coloredcobra"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     meta.AppName,
	Short:   "Shorter description",
	Long:    `Longer description`,
	Version: meta.Version,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {

		return nil
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cc.Init(&cc.Config{
		RootCmd:         rootCmd,
		Headings:        cc.HiCyan + cc.Bold + cc.Underline,
		Commands:        cc.HiYellow + cc.Bold,
		Example:         cc.Bold,
		ExecName:        cc.Bold,
		Flags:           cc.Bold,
		FlagsDataType:   cc.Italic + cc.HiBlue,
		NoExtraNewlines: true,
	})

	// allow users to set custom config files
	rootCmd.PersistentFlags().StringP("config", "c", "", "Path to config file")
}
