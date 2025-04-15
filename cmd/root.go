package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gig",
	Short: "Gig (Git in Golang)",
	Long:  "Gig is a reimplementation of basic Git functions in Golang.",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oops. An error while executing gig '%s'\n", err)
		os.Exit(1)
	}
}
