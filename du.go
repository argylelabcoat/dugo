package dugo

import (
	"fmt"
	"os"
)

func DiskUsage(currPath string, info os.FileInfo, depth int) int64 {
	var size int64

	dir, err := os.Open(currPath)
	if err != nil {
		fmt.Println(err)
		return size
	}
	defer dir.Close()

	files, err := dir.Readdir(-1)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, file := range files {
		if file.IsDir() {
			size += DiskUsage(fmt.Sprintf("%s/%s", currPath, file.Name()), file, depth+1)
		} else {
			size += file.Size()
		}
	}

	return size
}
