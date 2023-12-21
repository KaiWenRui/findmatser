package main

import (
	"fmt"
	"os"
	"path/filepath"
	"syscall"
)

func main() {
	readDir, err := os.ReadDir("./")
	if err != nil {
		fmt.Println("Read Filed")
		panic(err)
	}
	fmt.Printf("权限 \t\t大小(KB) \t创建日期 \t\t名称\n")
	for _, v := range readDir {
		FileInfo, _ := v.Info()
		t := FileInfo.ModTime()
		fss := FileInfo.Sys().(*syscall.Stat_t)

		fmt.Println(fss.Uid)
		fmt.Printf("%s\t%4d\t\t%s\t%s\n", FileInfo.Mode(), FileInfo.Size()/8, t.Format("2006-01-02 15:04:05"), FileInfo.Name())
		filepath.Glob(FileInfo.Name())
	}

}
