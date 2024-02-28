package main

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:                   "teachable",
		Short:                 "teachable",
		Version:               "1.0.0",
		DisableFlagsInUseLine: true,
	}

	rootCmd.AddCommand(CoursesCmd())
	rootCmd.AddCommand(StudentsCmd())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(context.Background(), "teachable - unable to run this command", err)
		os.Exit(1)
	}

}
