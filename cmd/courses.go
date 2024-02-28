package main

import (
	"fmt"
	"log"
	"teachable/config"
	"teachable/courses"
	"teachable/pkg/teachable"

	"github.com/spf13/cobra"
)

func CoursesCmd() *cobra.Command {
	coursesCommand := &cobra.Command{
		Use:                   "get-published-courses-info",
		Short:                 "Get published courses info",
		Version:               "1.0.0",
		DisableFlagsInUseLine: true,
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()
			requiredInfo := args[0]
			if requiredInfo == "" || (requiredInfo != "names" && requiredInfo != "headings") {
				log.Fatal("Please input \"names\" or \"headings\" to retrieve your published courses infos")
			}
			fmt.Println("\nYou have requested your published courses", requiredInfo)

			configs := config.LoadEnvVars()

			teachableService := teachable.NewService(configs.TeachableURL, configs.TeachableAPIKey)
			coursesService := courses.NewService(teachableService)

			response, err := coursesService.GetCoursesInfo(ctx, requiredInfo)
			if err != nil {
				log.Fatal("Unable to fetch courses infos")
			}

			fmt.Println("Here they are:")
			for _, r := range response {
				fmt.Println(" -", r)
			}
		},
	}
	return coursesCommand
}
