package common

import (
	"fmt"
	"log/slog"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	// Viper instance used by commands
	v *viper.Viper
)

func InitViper() *viper.Viper {
	v := viper.New()

	// if cfgFile != "" {
	// 	// Use config file from the flag.
	// 	v.SetConfigFile(cfgFile)
	// } else {
		// Find home directory.
	// home, err := os.UserHomeDir()
	// cobra.CheckErr(err)

	// Search config in home directory with name ".dev-container-manager" (without extension).
	v.AddConfigPath(".")
	v.SetConfigType("yaml")
	v.SetConfigName(".dev-container-manager")

	v.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := v.ReadInConfig(); err == nil {
		slog.Info(fmt.Sprintf("Using config file: %s", v.ConfigFileUsed()))
	}

	return v
}

// Bind each cobra flag to its associated viper configuration (config file and environment variable)
func BindFlags(cmd *cobra.Command, v *viper.Viper) {
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		// Determine the naming convention of the flags when represented in the config file
		configName := f.Name
		// If using camelCase in the config file, replace hyphens with a camelCased string.
		// Since viper does case-insensitive comparisons, we don't need to bother fixing the case, and only need to remove the hyphens.
		// if replaceHyphenWithCamelCase {
		// 	configName = strings.ReplaceAll(f.Name, "-", "")
		// }
		slog.Info(fmt.Sprintf("configName: %s / changed: %v / set: %v\n", configName, f.Changed, v.IsSet(configName)))

		// Apply the viper config value to the flag when the flag is not set and viper has a value
		if !f.Changed && v.IsSet(configName) {
			val := v.Get(configName)
			cmd.Flags().Set(f.Name, fmt.Sprintf("%v", val))
		}
	})
}

func GetViper() *viper.Viper {
	return v
}
