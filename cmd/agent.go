/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/brandtkeller/dev-container-manager/cmd/common"
	"github.com/robfig/cron/v3"
	"github.com/spf13/cobra"
	"golang.org/x/exp/slog"
)

type agentFlags struct {
	flags
	schedule string // -s --schedule
	test     string // -t --test
}

var agentOpts = &agentFlags{}

// agentCmd represents the agent command
var agentCmd = &cobra.Command{
	Use:   "agent",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		slog.Info("agent called")

		if agentOpts.test != "" {
			slog.Info(fmt.Sprintf("test: %s\n", agentOpts.test))
		} else {
			slog.Info("No test variable configured")
		}

		c := cron.New()

		slog.Info(fmt.Sprintf("schedule: %s\n", agentOpts.schedule))

		c.AddFunc(agentOpts.schedule, func() { slog.Info("hello world") })
		c.Start()

		// sleep for 60 seconds
		time.Sleep(60 * time.Second)

		c.Stop() // Stop the scheduler (does not stop any jobs already running).
	},
}

func init() {
	rootCmd.AddCommand(agentCmd)

	agentCmd.Flags().StringVarP(&agentOpts.test, "test", "t", "", "a test message")
	agentCmd.Flags().StringVarP(&agentOpts.schedule, "schedule", "s", "", "cron format for schedule")

	// Here you will define your flags and configuration settings.
	v := common.InitViper(agentOpts.cfgFile)
	common.BindFlags(agentCmd, v)

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// agentCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// agentCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
