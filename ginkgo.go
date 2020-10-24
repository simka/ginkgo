package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	args := []string{"."}

	if len(os.Args) > 1 {
		args = os.Args[1:]
	}

	for _, arg := range args {
		err := tree(arg, "")

		if err != nil {
			fmt.Printf("Tree of %s: %v\n", arg, err)
		}
	}
}

func tree(root, indent string) error {
	fi, err := os.Stat(root)

	if err != nil {
		return fmt.Errorf("Couldn't stat %s because %v", root, err)
	}

	fmt.Println(fi.Name())

	if !fi.IsDir() {
		return nil
	}

	fis, err := ioutil.ReadDir(root)

	if err != nil {
		return fmt.Errorf("Could not read dir %s because %v", root, err)
	}

	var fileNames []string

	for _, fi := range fis {
		if fi.Name()[0] != '.' {
			fileNames = append(fileNames, fi.Name())
		}
	}

	for i, name := range fileNames {
		addIndent := "│   "

		if i == len(fileNames)-1 {
			fmt.Printf(indent + "└── ")
			addIndent = "    "
		} else {
			fmt.Printf(indent + "├── ")
		}

		if err := tree(filepath.Join(root, name), indent+addIndent); err != nil {
			return err
		}
	}

	return nil
}
