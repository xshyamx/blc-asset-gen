package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	rootDir := "E:/workspaces/activ-wsp/broadleaf/core/src/main/resources/cms/static/img"
	filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		path = strings.Replace(path, "\\", "/", -1)
		path = strings.Replace(path, rootDir, "/img", 1)
		if info.IsDir() {
			fmt.Println(path + "/");
		} else {
			fmt.Println(path);
		}
		return nil
	})
}
