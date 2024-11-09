package main

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type DirEntry = fs.DirEntry
type FileInfo struct {
	size      int64
	Path      string
	isLegal   bool
	SizeLabel string
}

const TARTGET_DIR = "node_modules"

var IGNORE_DIR = []string{".git", ".github", ".vscode", ".idea"}

var TotalSize int64 = 0

func includes(s []string, e string) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}
func GetDirList() ([]FileInfo, error) {
	currentDir, err := os.Getwd()

	if err != nil {
		return nil, err
	}
	dirEntries, err := os.ReadDir(currentDir)
	if err != nil {
		return nil, err
	}
	dirChan := make(chan fs.DirEntry, 10)
	var targetDirInfoList []FileInfo
	var wg sync.WaitGroup
	//遍历根目录最外层文件夹
	for _, dir := range dirEntries {
		if dir.IsDir() {
			wg.Add(1)
			dirChan <- dir
			go GetNodeModulesList(currentDir, dirChan, &targetDirInfoList, &wg)
		}

	}
	wg.Wait()
	return targetDirInfoList, nil
}
func GetNodeModulesList(currentDir string, dirChan chan fs.DirEntry, targetDirInfoList *[]FileInfo, wg *sync.WaitGroup) {
	defer wg.Done()
	dir := <-dirChan
	basePath := filepath.Join(currentDir, dir.Name())
	dirEntries, err := os.ReadDir(basePath)
	if err != nil {
		log.Fatalln(err)
		return
	}
	getTargetDirInfo(basePath, dirEntries, targetDirInfoList)

}
func getTargetDirInfo(basePath string, dirEntries []DirEntry, targetDirInfoList *[]FileInfo) {
	for _, dirEntry := range dirEntries {
		if dirEntry.IsDir() {
			if includes(IGNORE_DIR, dirEntry.Name()) {
				continue
			}

			dirPath := filepath.Join(basePath, dirEntry.Name())
			if dirEntry.Name() == TARTGET_DIR {
				labelPath := strings.ReplaceAll(dirPath, "\\", "/")
				size, err := getDirSize(dirPath)
				if err != nil {
					log.Fatalln(err)
				}
				sizeLabel := transferUnit(size)
				TotalSize = TotalSize + size
				*targetDirInfoList = append(*targetDirInfoList, FileInfo{
					isLegal:   true,
					size:      size,
					Path:      labelPath,
					SizeLabel: sizeLabel,
				})
			} else {
				childDirEntries, err := os.ReadDir(dirPath)
				if err != nil {
					log.Fatalln(err)
					return
				}
				getTargetDirInfo(dirPath, childDirEntries, targetDirInfoList)
			}
		}

	}
}
func getDirSize(path string) (int64, error) {
	var size int64
	err := filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		size += info.Size()
		return nil
	})
	return size, err
}
