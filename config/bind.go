package config

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// bind flags to viper
// all config flags should be registered as `PersistentFlags` in the root command
// for this to work
func BindFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().VisitAll(func(f *pflag.Flag) {
		key := f.Name

		// replace dashes with underscores since toml files use _
		// if we used camelcase, we could have just removed - only
		// since viper does case-insensitive searches
		key = strings.ReplaceAll(key, "-", "_")

		// bind the flag to viper
		viper.BindPFlag(key, f)
	})
}
