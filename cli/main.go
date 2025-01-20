package main

import "github.com/spf13/cobra"

var root = cobra.Command{
	Use:   "spiffe",
	Short: "SPIFFE CLI",
}

var initCmd = cobra.Command{
	Use:   "init",
	Short: "Initialize the SPIFFE sidecar",
	RunE: func(cmd *cobra.Command, args []string) error {

	},
}

func main() {
	root.AddCommand(&initCmd)
	root.Execute()
}
