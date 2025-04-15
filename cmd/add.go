package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/hepem/gig/constants"
	"github.com/hepem/gig/utils"
	"github.com/spf13/cobra"
)

func add(path string) {
	if !utils.DirExists(constants.GigDir) {
		fmt.Println("Gig directory does not exist.")
		return
	}

	if path == "" {
		fmt.Println("Please, provide a path.")
		return
	}

	file, err := utils.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	sha1 := utils.FileToSHA1(file)
	blob, err := utils.Deflate(file)
	if err != nil {
		fmt.Println("Error deflating file:", err)
		return
	}

	object_path := fmt.Sprintf("%s/%s", constants.ObjectDir, sha1[0:2])
	blob_path := fmt.Sprintf("%s/%s", object_path, sha1[2:])

	err = utils.CreateDir(object_path)
	if err != nil {
		fmt.Println("Error creating object directory:", err)
		return
	}

	err = utils.WriteFile(blob_path, blob)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	f, err := os.OpenFile(constants.IndexPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if _, err := fmt.Fprintf(f, "%s %s\n", sha1, path); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Added %s to the repository.\n", path)
}

var addCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"a"},
	Short:   "Add files to the repository.",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Adding files to the repository ...")
		add(args[0])
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
