package main

import (
	"fmt"
	"log"
	"strconv"
	"teachable/config"
	"teachable/courses"
	"teachable/pkg/teachable"

	"github.com/spf13/cobra"
)

func StudentsCmd() *cobra.Command {
	coursesCommand := &cobra.Command{
		Use:                   "get-students",
		Short:                 "Get students from a course",
		Version:               "1.0.0",
		DisableFlagsInUseLine: true,
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()
			courseID := args[0]
			if courseID == "" {
				log.Fatal("Please input \"{course-id}\" to retrieve the course students")
			}
			fmt.Println("\nYou have requested students from course", courseID)

			configs := config.LoadEnvVars()

			teachableService := teachable.NewService(configs.TeachableURL, configs.TeachableAPIKey)
			coursesService := courses.NewService(teachableService)

			courseIDInt, err := strconv.Atoi(courseID)
			if err != nil {
				log.Fatal("Unable to convert course id to int")
			}

			response, err := coursesService.GetStudentsInACourse(ctx, courseIDInt)
			if err != nil {
				log.Fatal("Unable to fetch students infos")
			}

			fmt.Println("Here are the students:")
			for _, r := range response {
				fmt.Println(" -", r.Name, " Email:", r.Email)
			}
		},
	}
	return coursesCommand
}
