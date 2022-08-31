package main

import (
	"fmt"
	"github.com/odinit/global/gos"
)

func main() {
	p := "/Volumes/OdinDisk/Develop/github.com/odinit/global/gcache"
	err := gos.ClearDir(p)
	if err != nil {
		fmt.Println(err.Error())
	}
}
