package internal

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss/tree"
)

const limitenclosure = 3

func getFiles(startDir string) ([]os.FileInfo, error) {
	dir, err := os.Open(startDir)
	if err != nil {
		return nil, err
	}
	defer dir.Close()

	files, err := dir.Readdir(-1)
	if err != nil {
		return nil, err
	}

	return files, nil
}

func getLevel(startDir string) int {
	level := 1

	for _, char := range startDir {
		if string(char) != "/" {
			continue
		}

		level++
	}

	return level
}

func syncExecute(startDir string, enclosure int, limitFiles int) chan *tree.Tree {
	ch := make(chan *tree.Tree)

	go func() {
		treeFiles, _ := buildTree(startDir, enclosure, limitFiles)
		ch <- treeFiles
		close(ch)
	}()

	return ch
}

func getTreeRoot(startDir string) *tree.Tree {
	treeFiles := tree.Root(startDir)

	if startDir == "." {
		treeFiles.Root("⁜")
	} else {
		lastIndex := strings.LastIndex(startDir, "/")
		treeFiles.Root(startDir[lastIndex+1:])
	}

	return treeFiles
}

func buildTree(startDir string, enclosure int, limitFiles int) (*tree.Tree, error) {
	files, err := getFiles(startDir)
	if err != nil {
		return nil, err
	}

	treeFiles := getTreeRoot(startDir)

	var folders []chan *tree.Tree
	filesMap := make(map[string]int)

	for _, file := range files {

		typeFile := strings.ToLower(filepath.Ext(file.Name()))

		filesMap[typeFile]++

		if filesMap[typeFile] > limitFiles && startDir != "." {
			continue
		}

		if file.IsDir() && getLevel(startDir) < enclosure {
			child := syncExecute(startDir+"/"+file.Name(), enclosure, limitFiles)

			folders = append(folders, child)

		} else {
			treeFiles.Child(file.Name())
		}
	}

	for _, folder := range folders {
		treeFiles.Child(<-folder)
	}

	if startDir != "." {
		hideFiles(filesMap, limitFiles, treeFiles)
	}

	return treeFiles, nil
}

func GetTree(startDir string, enclosure int, limitFiles int, searchPattern string) (*tree.Tree, error) {
	return buildTree(startDir, enclosure, limitFiles)
}

func hideFiles(filesMap map[string]int, limitFiles int, treeFiles *tree.Tree) {
	for key, value := range filesMap {
		if value < limitFiles {
			continue
		}

		treeFiles.Child("and more " + strconv.Itoa(value-limitFiles) + " " + key)
	}
}
