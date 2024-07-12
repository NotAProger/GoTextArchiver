package cmd

import "github.com/spf13/cobra"

var unpackCmd = &cobra.Command{
	Use:   "unpack",
	Short: "Unpack file useing variable-length code",
}

func init() {
	rootCmd.AddCommand(unpackCmd)
}
