package cmd

import (
	"bufio"
	"bytes"
	"compress/zlib"
	"crypto/sha1"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/hepem/gig/constants"
	"github.com/spf13/cobra"
)

var commitMessage string

func ReadIndex() (map[string]string, error) {
	index := make(map[string]string)

	file, err := os.Open(constants.IndexPath)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.SplitN(scanner.Text(), " ", 2)
		if len(parts) != 2 {
			continue
		}
		index[parts[1]] = parts[0]
	}
	return index, scanner.Err()
}

func BuildTree(entries map[string]string) (string, error) {
	tree := make(map[string]interface{})
	for path, sha := range entries {
		parts := strings.Split(path, "/")
		current := tree
		for i, part := range parts {
			if i == len(parts)-1 {
				current[part] = sha
			} else {
				if _, ok := current[part]; !ok {
					current[part] = make(map[string]interface{})
				}
				current = current[part].(map[string]interface{})
			}
		}
	}

	return WriteTree("root", tree)
}

func WriteTree(name string, tree map[string]interface{}) (string, error) {
	var buffer bytes.Buffer

	for key, value := range tree {
		switch v := value.(type) {
		case string:
			buffer.WriteString(fmt.Sprintf("blob %s %s\n", v, key))
		case map[string]interface{}:
			sha, err := WriteTree(key, v)
			if err != nil {
				return "", err
			}
			buffer.WriteString(fmt.Sprintf("tree %s %s\n", sha, key))
		}
	}

	content := buffer.Bytes()
	sha := fmt.Sprintf("%x", sha1.Sum(content))
	err := writeObject(sha, content)
	return sha, err
}

func writeObject(sha string, content []byte) error {
	dir := filepath.Join(constants.GigDir, "objects", sha[:2])
	file := filepath.Join(dir, sha[2:])
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()
	w := zlib.NewWriter(f)
	_, err = w.Write(content)
	w.Close()
	return err
}

func BuildCommit(treeSHA, message string) (string, error) {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("tree %s\n", treeSHA))
	buffer.WriteString(fmt.Sprintf("author user\n\n%s\n", message))
	content := buffer.Bytes()
	sha := fmt.Sprintf("%x", sha1.Sum(content))
	err := writeObject(sha, content)
	return sha, err
}

func UpdateRef(commitSHA string) error {
	headContent, err := os.ReadFile(fmt.Sprintf("%s/HEAD", constants.GigDir))
	if err != nil {
		return err
	}
	ref := strings.TrimSpace(strings.Split(string(headContent), " ")[1])
	refPath := filepath.Join(constants.GigDir, ref)
	return os.WriteFile(refPath, []byte(commitSHA), 0644)
}

func ClearIndex() {
	err := os.Truncate(constants.IndexPath, 0)
	if err != nil {
		log.Fatal(err)
	}
}

func Commit() {
	entries, err := ReadIndex()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading index:", err)
		os.Exit(1)
	}

	if len(entries) == 0 {
		fmt.Fprintln(os.Stderr, "Nothing to commit")
		os.Exit(1)
	}

	treeSHA, err := BuildTree(entries)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error building tree:", err)
		os.Exit(1)
	}

	commitSHA, err := BuildCommit(treeSHA, commitMessage)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error building commit:", err)
		os.Exit(1)
	}

	if err := UpdateRef(commitSHA); err != nil {
		fmt.Fprintln(os.Stderr, "Error updating ref:", err)
		os.Exit(1)
	}

	ClearIndex()

	fmt.Println("Commit successful:", commitSHA)
}

var commitCmd = &cobra.Command{
	Use:     "commit",
	Aliases: []string{"c"},
	Short:   "Commit your changes to the repository.",
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Commiting files to the repository ...")
		Commit()
	},
}

func init() {
	commitCmd.Flags().StringVarP(&commitMessage, "message", "m", "", "The commit message")
	commitCmd.MarkFlagRequired("message")
	rootCmd.AddCommand(commitCmd)
}
