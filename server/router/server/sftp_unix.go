//go:build !windows
// +build !windows

package server

func getFileInfo(file os.FileInfo, remoteDir string) map[string]interface{} {
	_type := "f"
	if file.IsDir() {
		_type = "d"
	}

	stat, ok := file.Sys().(*syscall.Stat_t)
	if !ok {
		render.ParamError(c, "Failed to assert type")
		return
	}

	// 获取用户信息
	usr, err := user.LookupId(fmt.Sprint(stat.Uid))
	if err != nil {
		render.ParamError(c, "Failed to lookup user: "+err.Error())
		return
	}

	// 获取组信息
	grp, err := user.LookupGroupId(fmt.Sprint(stat.Gid))
	if err != nil {
		render.ParamError(c, "Failed to lookup group: "+err.Error())
		return
	}

	// 获取 user 和 group 的逻辑
	// 例如：
	// user := "username"
	// group := "groupname"

	return map[string]interface{}{
		"path":     filepath.Join(remoteDir, file.Name()),
		"name":     file.Name(),
		"mode":     file.Mode().String(),
		"size":     file.Size(),
		"mod_time": file.ModTime().Format("2006-01-02 15:04:05"),
		"type":     _type,
		"user":     usr.Name,
		"group":    grp.Name,
	}
}
