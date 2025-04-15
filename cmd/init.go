package cmd

import (
	"fmt"
	"os"

	"github.com/hepem/gig/constants"
	"github.com/hepem/gig/utils"
	"github.com/spf13/cobra"
)

func buildObjectDir() {
	err := utils.CreateDir(constants.ObjectDir)
	if err != nil {
		fmt.Println("Error creating object directory:", err)
		os.Exit(1)
	}

	err = utils.CreateDir(fmt.Sprintf("%s/info", constants.ObjectDir))
	if err != nil {
		fmt.Println("Error creating info directory:", err)
		os.Exit(1)
	}

	err = utils.CreateDir(fmt.Sprintf("%s/pack", constants.ObjectDir))
	if err != nil {
		fmt.Println("Error creating pack directory:", err)
		os.Exit(1)
	}
}

func buildRefsDir() {
	err := utils.CreateDir(constants.RefsDir)
	if err != nil {
		fmt.Println("Error creating refs directory:", err)
		os.Exit(1)
	}

	err = utils.CreateDir(fmt.Sprintf("%s/heads", constants.RefsDir))
	if err != nil {
		fmt.Println("Error creating heads directory:", err)
		os.Exit(1)
	}

	err = utils.CreateDir(fmt.Sprintf("%s/tags", constants.RefsDir))
	if err != nil {
		fmt.Println("Error creating tags directory:", err)
		os.Exit(1)
	}
}

func initialize() {
	err := utils.CreateDir(constants.GigDir)

	if err != nil {
		fmt.Println("Skipping. Repository already exists.")
		return
	}

	buildObjectDir()
	buildRefsDir()

	file, err := os.Create(fmt.Sprintf("%s/HEAD", constants.GigDir))
	if err != nil {
		fmt.Println("Error creating HEAD file:", err)
		os.Exit(1)
	}

	defer file.Close()

	_, err = file.WriteString("ref: refs/heads/master\n")
	if err != nil {
		fmt.Println("Error writing HEAD file:", err)
		os.Exit(1)
	}
}

var initCmd = &cobra.Command{
	Use:     "init",
	Aliases: []string{"i"},
	Short:   "Initialize a new repository",
	Long:    "Initialize a new repository in the current directory.",
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Initializing repository ...")
		initialize()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
