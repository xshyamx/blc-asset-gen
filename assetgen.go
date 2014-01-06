package main

import (
	"fmt"
	"os"
	"io"
	"path/filepath"
	"strings"
	"crypto/md5"
)
 	
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func copy(src, dest string) bool {
	sf, err := os.Open(src)
	check(err)
	err = os.MkdirAll(filepath.Dir(dest), os.ModeDir)
	check(err)
	df, err := os.Create(dest)
	check(err)
	bytes, err := io.Copy(df, sf)
	return bytes > 0 && err == nil
}

func main() {
	rootDir := "E:/workspaces/activ-wsp/broadleaf/core/src/main/resources/cms/static/img"
	outDir  := "E:/workspaces/activ-wsp/broadleaf/site/target/assets"
	fileCount := 0
	filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		url := strings.Replace(path, "\\", "/", -1)
		url = strings.Replace(url, rootDir, "/img", 1)
		if !info.IsDir() {
			md5 := md5.New()
			md5.Write([]byte(url))
			digest := fmt.Sprintf("%x", md5.Sum(nil))
			dest := filepath.Join(outDir, digest[0:2], digest[2:4], filepath.Base(path))
			if copy(path, dest) {
				fileCount++
			}
		}
		return nil
	})
	defer func() {
		fmt.Printf("Copied %d static assets\n", fileCount)
	}()
}
