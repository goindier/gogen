package pkg

import (
	"os"
	"path/filepath"
)

func FindTplFiles(root string) ([]string, error) {
	var tplFiles []string

	// 遍历文件夹
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 检查是否是文件且后缀为 .tpl
		if !info.IsDir() && filepath.Ext(path) == ".tpl" {
			tplFiles = append(tplFiles, path)
		}
		return nil
	})

	return tplFiles, err
}
