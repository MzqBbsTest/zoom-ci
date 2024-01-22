// Copyright 2021 Zoom Author. All Rights Reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gofile

import (
	"os"
)

func CreateFile(filePath string, data []byte, perm os.FileMode) error {
	return os.WriteFile(filePath, data, perm)
}
