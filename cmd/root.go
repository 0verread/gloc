package cmd

import(
		"fmt"
		"os"
		
		"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
		Use:			"gloc",
		Version:		"0.1.0",
		Short:			"gloc: count line of code in a beautiful way",
		Run: func(cmd *cobra.Command, args []string) {
				fmt.Println("running gloc...")
		},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}


