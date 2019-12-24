package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	times int
	payload string
	method string
)

func Execute() {
	cmdCall.PersistentFlags().IntVarP(&times, "times", "t", 1, "number of concurrent calls")
	cmdCall.PersistentFlags().StringVarP(&payload, "payload", "p", "", "payload for POST requests")
	cmdCall.PersistentFlags().StringVarP(&method, "method", "m", "post", "define the http method")

	var rootCmd = &cobra.Command{Use: "concurl"}
	rootCmd.AddCommand(cmdCall)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
