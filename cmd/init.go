package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var gigDir = ".gig"
var objectDir = fmt.Sprintf("%s/objects", gigDir)
var refsDir = fmt.Sprintf("%s/refs", gigDir)

func buildObjectDir() {
	err := os.Mkdir(objectDir, 0755)
	if err != nil {
		fmt.Println("Error creating object directory:", err)
		os.Exit(1)
	}

	err = os.Mkdir(fmt.Sprintf("%s/info", objectDir), 0755)
	if err != nil {
		fmt.Println("Error creating info directory:", err)
		os.Exit(1)
	}

	err = os.Mkdir(fmt.Sprintf("%s/pack", objectDir), 0755)
	if err != nil {
		fmt.Println("Error creating pack directory:", err)
		os.Exit(1)
	}
}

func buildRefsDir() {
	err := os.Mkdir(refsDir, 0755)
	if err != nil {
		fmt.Println("Error creating refs directory:", err)
		os.Exit(1)
	}

	err = os.Mkdir(fmt.Sprintf("%s/heads", refsDir), 0755)
	if err != nil {
		fmt.Println("Error creating heads directory:", err)
		os.Exit(1)
	}

	err = os.Mkdir(fmt.Sprintf("%s/tags", refsDir), 0755)
	if err != nil {
		fmt.Println("Error creating tags directory:", err)
		os.Exit(1)
	}
}

func initialize() {
	err := os.Mkdir(gigDir, 0755)

	if err != nil {
		fmt.Println("Error creating the repository:", err)
		os.Exit(1)
	}

	buildObjectDir()
	buildRefsDir()

	file, err := os.Create(fmt.Sprintf("%s/HEAD", gigDir))
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
		fmt.Println("Initializing repository...")
		initialize()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
