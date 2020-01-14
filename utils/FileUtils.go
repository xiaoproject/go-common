package utils

import "os"

func FileMkdir(filePath string) bool {

	exist := isExist(filePath)

	if exist {
		return true
	} else {
		err := createFile(filePath)
		if err == nil {
			return true
		} else {
			return false
		}
	}
}

//调用os.MkdirAll递归创建文件夹
func createFile(filePath string) error {
	if !isExist(filePath) {
		err := os.MkdirAll(filePath, os.ModePerm)
		return err
	}
	return nil
}

// 判断所给路径文件/文件夹是否存在(返回true是存在)
func isExist(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
