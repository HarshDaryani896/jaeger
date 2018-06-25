// Code generated by go-bindata.
// sources:
// web_assets/index.html
// DO NOT EDIT!

package frontend

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"github.com/elazarl/go-bindata-assetfs"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _web_assetsIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x57\xff\x53\xdb\x38\x16\xff\x3d\x7f\xc5\x1b\x6d\xef\xec\x5c\xb0\x9d\x10\x28\x34\xc4\xb9\xe1\x48\x97\xd2\xbd\x2e\xbd\x40\xbb\xd3\xdb\xe9\x0f\xb2\xf4\x62\x2b\xd8\x92\x2b\xc9\x81\x2c\xc3\xff\x7e\x23\xdb\x09\x09\x94\x9b\xfd\x71\x9b\x19\x40\x7a\xdf\x3f\xef\x8b\x5e\x18\x67\xb6\xc8\x27\x1d\x80\x71\x81\x96\x02\xcb\xa8\x36\x68\x63\x72\x71\x75\x19\x1c\x1f\x1f\xbe\x09\x06\xe4\x91\x2b\x69\x81\x31\x59\x0a\xbc\x2d\x95\xb6\x04\x98\x92\x16\xa5\x8d\xc9\xad\xe0\x36\x8b\x39\x2e\x05\xc3\xa0\xbe\xec\x81\x90\xc2\x0a\x9a\x07\x86\xd1\x1c\xe3\x41\xd8\xdf\x83\x82\xde\x89\xa2\x2a\xb6\x49\x95\x41\x5d\xdf\x69\x92\x63\xdc\x6f\x9c\x65\x48\xb9\x3b\x00\x8c\xad\xb0\x39\x4e\xde\x29\x3b\xbb\x9c\x42\x00\x33\xc1\xd1\xc0\xa5\x84\x29\x16\x54\xf2\x71\xd4\xf0\x1b\x59\xc3\xb4\x28\x2d\x18\xcd\x62\x92\x59\x5b\x9a\x51\x14\x31\xc5\x31\x5c\x7c\xab\x50\xaf\x42\xa6\x8a\xa8\x39\x06\xc3\x70\x10\x0e\xc2\x42\xc8\x70\x61\xc8\x64\x1c\x35\xaa\xad\x9d\x5c\xc8\x1b\xd0\x98\xc7\xc4\xd8\x55\x8e\x26\x43\xb4\x04\x32\x8d\xf3\x47\xbb\x05\xbd\x63\x5c\x86\x89\x52\xd6\x58\x4d\x4b\x77\x71\xf6\x37\x84\x68\x18\x0e\xc3\xa3\x88\x19\xf3\x48\xab\x1d\x32\x63\x08\x08\x69\x31\xd5\xc2\xae\x62\x62\x32\x3a\x3c\x3e\x08\xfe\xf5\xf9\x8b\x10\x57\x17\x3f\xe3\x2f\x03\x7e\x5e\xbc\x9f\x9d\xde\xac\x58\xf5\xee\xf4\xdd\x2c\x1d\xee\x5f\x16\x9f\xd8\xed\xed\x91\x92\xc3\xd9\x17\x9e\x1e\x7c\xa6\xbd\x8f\xc5\xd5\xb5\xf9\x23\xfa\xe5\xf5\xf1\x32\xe1\x6f\x17\xd9\x41\x45\x80\x69\x65\x8c\xd2\x22\x15\x32\x26\x54\x2a\xb9\x2a\x54\x65\xc8\xff\x49\xce\x9f\x05\xb1\x78\x8a\x61\xf1\x5d\x08\xd7\xec\xf0\xe2\x3f\x22\xe9\xef\x1f\x7d\x5b\xae\x16\x57\x1f\xe6\xef\x16\x97\x1f\xe8\xbf\x6f\xe6\xd5\x6f\x9f\xef\xfe\x7b\xf7\xe9\xa3\x3c\x7b\x7f\x7a\x94\xef\x17\x67\xbf\xfd\x7a\x51\x9e\xbf\x29\xce\xcf\xa6\xc7\xb7\xe7\xbf\x5e\xb0\x8f\xd3\xa3\xeb\x3b\xfa\x32\x84\xc7\x02\xb5\x60\x5c\x5d\x26\x9d\xb0\xaa\x04\x87\x7b\x28\xa8\x4e\x85\x0c\xac\x2a\x47\x30\x38\x2c\xef\x4e\xe0\xa1\x13\x66\xca\x6a\xc5\x83\xa4\xb2\x56\x49\xb8\x87\x92\x72\x2e\x64\x3a\x82\xfd\xbe\x93\x60\x95\x36\x4a\x8f\xa0\x54\x0e\x88\x3e\xd9\x35\xd2\xff\x9e\x91\x51\xa6\x96\xa8\xe1\xfe\xb9\xee\x5c\xe4\x16\xf5\x08\x12\x2d\xd2\xcc\x4a\x34\xc6\x3f\x3e\xfc\x5b\xd7\x99\xf8\xa9\x35\x91\xab\xf4\x85\x48\x7f\xb2\xa2\x7c\x81\x55\x83\x8d\x5a\xb4\x6e\x32\xa2\xf5\x68\x8c\x13\xc5\x57\x6d\x69\xb9\x58\x02\xcb\xa9\x31\x31\x71\x13\x49\x85\x44\xdd\x96\x7d\x97\x5b\xa7\x8b\xe6\xa8\x6d\xf3\x3b\x10\x72\xae\x5c\x76\xb9\x58\x6e\xe4\x19\x3a\x4c\xeb\xab\x9b\xc6\x81\x9b\x3f\x98\x85\x97\xe1\x34\x1c\x47\xd9\x60\x9b\x77\x30\x19\x63\x31\x79\x36\x96\x58\x4c\xc6\x51\x76\xb0\x25\xb9\x15\x86\x56\xb7\xe4\x91\xf3\x1c\x42\x1e\x14\x3c\x18\x82\x3b\x98\x22\x78\xfd\x44\xb6\x69\x80\x92\xca\x67\x54\xf7\x69\x8d\x24\x56\x42\x62\x65\x0d\xb0\x3e\x24\xb9\x62\x37\xb0\x53\x4e\xf2\x5d\x03\x9c\x5a\x1a\xb0\xca\x58\x55\xa0\x8e\xc9\x60\x7f\x48\x26\x33\xca\x32\xcc\x3d\x03\x3f\xe7\x4a\xd3\x1c\xa6\x68\x44\x2a\xcd\x38\x72\x61\x3c\x41\xb2\x9d\xcb\xbf\x3a\xb8\xe1\x9b\x7d\x32\xb9\xd6\xaa\x80\xb3\x4c\x31\x95\x53\x2b\x50\xff\xf0\xa8\x8e\x86\x03\x32\x79\x4f\x4b\x2a\xd1\xa0\xab\x15\x6a\xfb\xe3\xd7\xea\xf0\xf5\x11\x99\x9c\x16\xf4\x0f\x21\x53\x38\x53\xf3\x39\x22\xcc\x14\x35\x16\xf5\x9f\x01\xf7\xf4\xea\x70\x0a\x1e\x13\x2b\x4a\x32\x39\xcb\x05\xbb\x01\x25\x61\xed\xae\xde\xf4\x40\x13\xb5\x44\xb0\x0a\x94\xe6\xa8\x81\x02\xa3\x3a\x7c\xc9\xd0\xe3\x33\x47\xd6\xd8\x73\xa4\xfc\xe9\xeb\x12\x6d\x3f\x2f\x1b\xd6\x38\x6a\x9e\xb3\xce\x66\x51\x4d\x3a\x9d\x79\x25\x99\x15\x4a\xc2\x5c\xe9\x82\xda\x69\xa5\xa9\xbb\xfa\xbc\x3d\x74\xe1\xbe\x03\xb0\xa4\x1a\x38\xc4\xb0\xa6\x42\x04\xfe\xa0\x5f\x7f\xe0\x1f\x30\x68\xfe\xbc\xee\x77\x4f\x5a\xd9\x4a\x0a\x6b\x20\x06\xaf\x10\xd2\x73\x44\x8d\xb6\xd2\x12\x3e\x50\x9b\x85\x5a\x55\x92\xfb\xbc\x0b\xbd\x46\xee\xa4\xf3\xd0\xe9\x38\x2d\x96\x0b\x94\xf6\xd3\xa7\x8b\x29\xc4\xdb\xa2\xcd\x91\x4a\xae\x0a\xbf\xdb\xfa\x73\xbe\x9c\x4e\x4e\x8d\x9d\xe1\xb7\x0a\x8d\xad\xd5\xfa\x27\x9d\xce\x2b\x9f\xd4\x4b\x8b\x74\x43\xf7\xc5\xcb\x27\x5f\x54\xa5\xe1\x16\x93\xd6\x83\x67\x40\xf0\x91\x5b\x70\x5a\xc9\x74\x42\xa0\xb7\xed\xba\x07\xc4\x6d\x83\x86\xd5\x6d\xcd\xed\xb6\x52\x37\x64\xae\x98\xfe\x3a\x79\x3e\x2e\x6d\x93\xa8\x9d\x70\x7a\xbd\x75\x3e\xf4\x56\x84\xbb\xae\x02\xe7\x7d\x47\x6b\xad\x33\xd7\x68\xb2\x33\xaa\x21\x86\x57\xfe\x2b\x9f\x6c\xed\x38\xd2\x0d\x4b\x8d\x25\x4a\xee\x7b\xdb\xc3\x54\xab\x04\x8c\x6a\x52\x2f\x8c\xa9\x30\x25\xb5\x2c\x73\xcd\xdc\xf4\x55\x18\xfe\xae\xf1\xdb\x08\xbc\xde\x26\xa2\x9e\xf7\xb5\xdd\x24\xae\x4d\xbc\x6e\xc8\x32\x91\x73\x8d\xd2\xef\xfe\xde\xff\xba\xa9\xe8\xa6\x69\x63\xc0\xa5\x0d\x2d\xd5\x29\xda\xd0\x8d\x8f\x41\x1b\xae\xb9\x4e\xda\x2d\x4f\xd4\xae\xfa\xf7\x6d\x3f\x7a\x0b\x8a\x29\xea\x20\xa1\x69\x4a\x53\xf4\x46\xe0\x19\x34\x46\x28\x19\x7b\x4f\x73\xef\xed\xad\x93\x55\xf3\x36\x61\x76\x00\x1e\x9c\x75\xa6\xa4\x51\x39\x86\xb9\x4a\xfd\xd6\xd3\x26\xc6\x04\xe7\x4a\x23\xc4\x30\xa5\x16\x43\xa9\x6e\xfd\x9a\xf5\x2a\xa4\x0b\x7a\xe7\x7b\x11\x6f\xf3\xf1\xcf\xcd\xc0\xd7\xee\xd7\xc8\x7a\xe0\xfd\x5d\x2a\x69\xb0\x26\xef\x34\xdd\x5e\x8b\xa5\xf5\x38\x5a\x1f\xf6\x6a\x6a\x81\x36\x53\x7c\x04\xde\xf9\xdb\x6b\xaf\x21\x99\x8a\x31\x34\x66\x04\x9b\x16\x71\xa9\xda\x03\x8b\x77\xf6\xca\x52\x5b\x99\xee\x26\x3d\x2e\x74\x3a\xb7\x75\x6e\x77\x23\xaf\x9f\xb8\x2d\xc0\xce\xc6\x86\x51\xcf\xe4\x7a\x1e\xe3\x67\xf3\x4b\x2d\x0d\xdf\x5e\x9f\x6e\xc4\xd7\xed\xd4\x8c\x84\xd7\x7e\xd9\x1f\x27\x13\x07\xb6\x96\x9e\x6a\xb1\x6c\xd2\x30\x8e\x92\x09\x50\xad\xc5\xd2\xb5\x8e\x90\x50\xcb\xac\x7d\xf5\xc0\x83\xb6\x8d\xb6\x2b\xd4\x14\x2f\xa7\x16\x25\x5b\x35\x3c\xbf\x86\x15\x34\x75\x71\xf3\xee\x15\xe6\xab\xd7\x86\xf4\xe0\x32\xf5\xd0\x3d\xe9\xb8\x9f\xfa\x79\xda\x3c\x4a\xe3\xa8\xf9\x87\xe9\x7f\x01\x00\x00\xff\xff\xb4\x93\x5f\x7b\x38\x0d\x00\x00")

func web_assetsIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_web_assetsIndexHtml,
		"web_assets/index.html",
	)
}

func web_assetsIndexHtml() (*asset, error) {
	bytes, err := web_assetsIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "web_assets/index.html", size: 3384, mode: os.FileMode(420), modTime: time.Unix(1519405039, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"web_assets/index.html": web_assetsIndexHtml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"web_assets": &bintree{nil, map[string]*bintree{
		"index.html": &bintree{web_assetsIndexHtml, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

func assetFS() *assetfs.AssetFS {
	assetInfo := func(path string) (os.FileInfo, error) {
		return os.Stat(path)
	}
	for k := range _bintree.Children {
		return &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: assetInfo, Prefix: k}
	}
	panic("unreachable")
}