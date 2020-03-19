package cmd

import (
	"github.com/jkonarze/concurl/internal"
	"github.com/spf13/cobra"
)

var cmdCall = &cobra.Command{
	Use:   "call [url]",
	Short: "Call given url",
	Long: `An easy way to call a given url by default once`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		svc := internal.NewSvc(times, args[0], method, payload, headers)
		svc.Init()
	},
}
