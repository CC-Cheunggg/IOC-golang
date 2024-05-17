package gen

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func GetWorkingDirectory() (string, error) {
	// 获取当前工作目录
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	paths := make([]string, 0)
	// 递归遍历当前目录及其子目录
	err = filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		// 如果是目录，则打印出来
		if d.IsDir() {

			entries, err := os.ReadDir(path)
			if err != nil {
				return err
			}

			if len(entries) > 0 {
				paths = append(paths, strings.ReplaceAll(path, dir, ".")+"/...")
			}
		}
		return nil
	})
	if err != nil {
		return "", err
	}

	return strings.Join(paths, ";"), nil
}
