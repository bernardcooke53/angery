/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "angery",
	Short: "Text just not quite getting the message across? Make it A N G E R Y",
	Long: `
Emphasise your point by transorming your text into something that makes it clear what you are trying to say.

Make it A N G E R Y

https://www.youtube.com/watch?v=5jO2PLqEdUY

Vegetals are not included`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },

	Run: angeryMain,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.angery.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func angeryMain(cmd *cobra.Command, args []string) {
	text, err := readStdin()
	cobra.CheckErr(err)

	transformed := transformText(text)

	fmt.Println(transformed)
}

func readStdin() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Can i haz some txt?: ")
	text, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return text, nil
}

func transformText(text string) string {
	chars := strings.Split(text, "")
	return strings.Join(chars, " ")
}
