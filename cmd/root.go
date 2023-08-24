package cmd

import (
	"fmt"
	"os"

	"github.com/Yakiyo/go-template/config"
	logger "github.com/Yakiyo/go-template/log"
	"github.com/Yakiyo/go-template/meta"
	cc "github.com/ivanpirog/coloredcobra"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:     meta.AppName,
	Short:   "Shorter description",
	Long:    `Longer description`,
	Version: meta.Version,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		config.BindFlags(cmd)
		logger.SetLevel(viper.GetString("log_level"))
		fmt.Println("In pre-run")
		fmt.Printf("%#v\n", viper.AllSettings())
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

	f := rootCmd.PersistentFlags()
	// allow users to set custom config files
	f.StringP("config", "c", "", "Path to config file")
	// dont mention debug level, usually users dont need, only on the dev side
	f.String("log-level", "", "Set log level [info, warn, error, fatal]")
	f.String("color", "", "Set color output [always, auto, never]")
}
