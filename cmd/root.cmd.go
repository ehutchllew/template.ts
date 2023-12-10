package cmd

import (
	"fmt"
	"os"

	"github.com/ehutchllew/template.ts/cmd/models"
	"github.com/ehutchllew/template.ts/cmd/utils"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use: "tmplts",
		RunE: func(cmd *cobra.Command, args []string) error {
			printWelcomeMessage()
			checkForFlags(cmd)

			return nil
		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func checkForFlags(cmd *cobra.Command) error {
	allFlag, allFlagErr := cmd.Flags().GetBool("all")
	if allFlagErr != nil {
		return allFlagErr
	}

	if allFlag {
		questionnaire := models.UserAnswers{
			EsLint:     true,
			Jest:       true,
			Swc:        true,
			Typescript: true,
		}

		writeFiles(&questionnaire)
	} else {
		requestUserInput()
	}

	return nil
}

func init() {
	rootCmd.Flags().BoolP("all", "a", false, "tmplts -a")
}

func printWelcomeMessage() {
	fmt.Println("**********************************************")
	fmt.Println("*                                            *")
	fmt.Println("* Welcome to Template.TS! Let's get started. *")
	fmt.Println("*                                            *")
	fmt.Println("**********************************************")
}

func requestUserInput() {
	/*
	 * To be filled out with custom functionality in the future
	 */
	// questionnaire := models.UserAnswers{}
	// writeFiles(&questionnaire)
}

func writeFiles(userInput *models.UserAnswers) {
	// TODO: Will contain the actual text to write instead of hard coding it inside GenerateRoot
	answers := models.UserAnswers{
		AppName:    "testingggg",
		EsLint:     true,
		Jest:       true,
		Swc:        true,
		Typescript: true,
	}

	dir, _ := os.Getwd()
	fmt.Printf("\nCWD:::%v", dir)
	utils.GenerateAll(&answers, dir)
}
