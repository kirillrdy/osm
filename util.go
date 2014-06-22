package osm

import (
	"path"
	"runtime"
)

func packageDir() string {
	_, current_file, _, _ := runtime.Caller(0)
	package_dir := path.Dir(current_file)
	return package_dir
}
