//go:build windows
// +build windows

package server

import (
	"os"
	"path"
)

func getFileInfo(file os.FileInfo, remoteDir string) map[string]interface{} {
	_type := "f"
	if file.IsDir() {
		_type = "d"
	}

	// Windows上不获取user和group
	return map[string]interface{}{
		"path":     path.Join(remoteDir, file.Name()),
		"name":     file.Name(),
		"mode":     file.Mode().String(),
		"size":     file.Size(),
		"mod_time": file.ModTime().Format("2006-01-02 15:04:05"),
		"type":     _type,
		"user":     "-",
		"group":    "-",
	}
}
