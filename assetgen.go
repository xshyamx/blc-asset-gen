package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func copy(src, dest string) bool {
	sf, err := os.Open(src)
	check(err)
	err = os.MkdirAll(filepath.Dir(dest), os.ModeDir|0755)
	check(err)
	df, err := os.Create(dest)
	check(err)
	bytes, err := io.Copy(df, sf)
	return bytes > 0 && err == nil
}

func main() {
	debug := flag.Bool("debug", false, "Print debug messages")
	rootDir := flag.String("src", "", "Source folder to copy assets")
	outDir := flag.String("dest", "target", "Target folder to copy assets")
	flag.Parse()
	*rootDir, _ = filepath.Abs(*rootDir)
	*outDir, _ = filepath.Abs(*outDir)
	if *debug {
		fmt.Println("rootDir:", *rootDir)
		fmt.Println("outDir:", *outDir)
	}
	if *rootDir == "" {
		flag.PrintDefaults()
		return
	}
	fileCount := 0
	filepath.Walk(*rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		url := strings.Replace(path, *rootDir, "/img", 1)
		url = strings.Replace(url, "\\", "/", -1)
		filename := filepath.Base(path)
		if !info.IsDir() {
			if filename[0] == '.' && len(filename) > 2 {
				return filepath.SkipDir
			}
			md5 := md5.New()
			md5.Write([]byte(url))
			digest := fmt.Sprintf("%x", md5.Sum(nil))
			dest := filepath.Join(*outDir, digest[0:2], digest[2:4], filename)
			if *debug {
				fmt.Println(path, "->", dest, "url:", url, "["+digest+"]")
			}
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
