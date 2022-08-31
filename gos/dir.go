package gos

import (
	"github.com/odinit/global/gvar"
	"os"
	"path"
)

func ClearDir(p string) (err error) {
	pStat, err := os.Stat(p)
	if err != nil {
		return
	}

	if !pStat.IsDir() {
		return gvar.ErrPathIsNotDir
	}

	ps, err := os.ReadDir(p)
	if err != nil {
		return
	}
	for pIndex := range ps {
		_ = os.RemoveAll(path.Join(p, ps[pIndex].Name()))
	}

	return nil
}
