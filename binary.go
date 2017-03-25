// Code generated by go-bindata.
// sources:
// assets/public/js/application.js
// DO NOT EDIT!

package hugo

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data, name string) ([]byte, error) {
	gz, err := gzip.NewReader(strings.NewReader(data))
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

var _publicJsApplicationJs = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe4\x57\x5b\x6f\xdb\xb8\x12\x7e\x0f\xd0\xff\x30\xd0\x39\x27\x94\x4e\x6c\x39\x6d\x1f\x16\x70\x2c\x03\x6d\xd2\x45\xbb\xdb\x1b\xda\x2e\xfa\xb0\x58\xa0\xb4\x38\x92\xd8\x50\xa4\x96\xa4\x9c\x1a\x69\xfe\xfb\x82\x12\x25\xdb\x8a\x2f\xdb\xa2\x4f\xbb\x79\x89\x4d\xce\xe5\xe3\x7c\x1f\x87\x63\x52\x1b\x04\x63\x35\x4f\x2d\xb9\x78\x70\xf2\xe0\x84\xa9\xb4\x2e\x51\xda\x98\x32\xf6\x6c\x89\xd2\xbe\xe4\xc6\xa2\x44\x1d\x92\xab\x37\xaf\x2e\x95\xb4\x6e\x4d\x51\x86\x8c\x8c\x00\x9d\x05\x24\x73\xb8\x7d\x70\x02\xd0\xfb\xfe\x59\xa3\x5e\xbd\x47\x81\xa9\x55\x3a\x24\xff\xb1\xaa\x1a\x2f\xa8\x86\x39\x30\xbe\x84\x39\x54\xd3\x8c\x6b\x63\xc7\x69\xc1\x05\x23\x51\xcc\xa5\x44\xfd\xfc\xc3\xab\x97\x90\x00\x79\x5e\xe7\x0a\x32\xa5\xe1\x92\x32\xb6\x22\x87\x02\x67\x4a\x59\xd4\x5b\x01\xce\x12\x20\xf0\x91\xdb\x02\x28\x64\x82\x2e\x55\xad\x41\x65\x30\xa3\xa0\x51\x24\x81\x54\xaa\x72\xa7\x01\xa9\x34\x66\xa8\x35\xea\x00\x0a\x8d\x59\x12\x14\xd6\x56\x66\x3a\x99\xe4\xdc\x16\xf5\x22\x4e\x55\x39\x29\x68\xca\x38\x35\x93\xd4\x21\x19\x17\x75\xae\x82\xb9\x83\x37\x9b\xd0\x79\xdc\xd6\xeb\xc0\xa9\x17\xca\x5a\x55\xba\x83\xcf\xff\x3f\x38\xb0\xb1\x2b\x81\x71\x49\xbf\x7c\xe4\xcc\x16\x90\x40\x90\x52\x91\x86\x0f\xcf\xcf\xff\x07\x63\x78\xf4\x13\x96\x51\xd0\x46\x17\x68\x41\x70\x79\x0d\x09\x2c\xa8\xc1\xdf\xde\xbd\x84\x33\x08\x26\x06\xad\xe5\x32\x37\x93\x60\x00\x22\x47\xfb\x4c\xa0\xfb\xf8\x74\xf5\x82\x85\x84\xcb\x4c\x35\x05\x32\xa8\xed\x13\xf6\x99\xa6\x28\xad\xab\x54\x48\x16\x98\x29\x8d\x0b\xcc\xb9\x24\x23\xf8\x34\xa3\xbe\x10\xff\xbd\x75\x09\xef\x82\xb9\x0b\xec\xfe\x66\x8e\xb5\x54\x50\x63\x92\x80\xa6\x96\x2b\xb9\xde\x83\x19\xef\xb6\x4a\x6a\x51\x73\x2a\xc6\x3c\x55\xd2\x04\xf3\x0e\xe3\x6c\xc2\x37\xcc\x4d\x45\xe5\xfc\x7d\xbf\xd5\x7c\xed\x13\x4d\x18\x5f\x76\xdf\x5c\x91\x3f\x45\xbe\xc8\x3c\x0b\x17\xb5\xb5\x4a\x9a\x58\xe2\x0d\x9c\x9e\xc2\x0d\x97\x4c\xdd\xc4\x42\xa5\xd4\x21\x8a\x2b\x6a\x0b\x49\x4b\x84\x24\xd9\x2e\x55\xda\x6a\x76\x12\x44\xad\x4a\x01\x36\x22\xc5\x1a\x4b\xb5\xc4\x81\xd2\x53\xc1\xd3\x6b\x32\x02\xc1\x8d\x83\xe9\x0c\x7f\xe6\x02\x9f\x36\x7e\x0e\xd2\x30\xca\xfd\xcb\xd2\x85\x70\xa2\xd9\xe5\x7f\x77\xef\x5c\x86\x2e\xb1\x87\xe8\x68\x5f\xa8\x2f\x90\xec\x67\x36\xe3\x02\xc7\x4a\x8a\x15\xe9\x8a\x04\xce\xe5\x00\xd5\x28\x59\x43\xf4\x7d\x3a\x81\xb3\x24\xa8\xea\x85\xe0\xa6\xd8\xe0\xf6\x18\xbf\x92\x6d\x71\xbb\xe6\xf7\x6d\x1b\x6a\x40\xef\x9a\xe2\x4f\x1b\x90\xfd\xf1\x7d\xf6\x43\x27\xf6\x26\x24\xda\xe9\x79\x8c\x04\x6f\x16\x75\x89\x79\x16\x86\x7b\x53\x31\x6a\x91\x44\xf0\xf5\xeb\x51\x34\xad\x65\x04\xa7\xa7\xdd\x31\xf7\x7a\x20\xe3\x56\xb9\x76\xc5\xa8\xa5\x06\x6d\x7c\xcd\x25\x83\xc4\xdd\x7e\x55\x56\x02\x2d\xb6\x1a\xed\x02\xfd\x2d\x36\x77\xdd\xce\x86\x4e\x93\x16\xc8\x6a\x81\x03\x3e\x0f\x31\x4a\x05\xd5\xe5\x3d\x4a\xfb\x4b\xeb\x03\xee\x60\x75\x07\xaf\x6b\x7e\x3a\x20\x87\xa8\xed\x6c\x3a\x6e\xef\x7b\x1f\xa3\xb7\xb3\xf3\x01\xee\x3a\x1c\x3f\xbe\x3d\x3f\x3e\x6f\xdb\x73\x93\xe4\xae\x3d\xf1\x92\xea\x06\x06\x24\x70\x7b\xd7\xac\xdc\xbf\xf8\x90\x40\x56\xcb\x86\x21\x08\x9b\x97\xd3\x5f\xf7\xe6\x73\x5c\xe9\xe6\xff\x15\x66\xb4\x16\x36\x6c\xdb\x44\xbb\x65\xac\xaa\xde\x6a\x55\xd1\xbc\xe9\x75\x61\x57\x65\xd7\x25\x52\xa1\xe4\x56\x69\x79\x59\x29\x6d\x5f\x2b\x86\xa1\xc5\xb2\x12\xd4\xa2\x71\x87\x37\x4d\x9b\xf4\xdd\x70\x04\x56\xd7\xd8\xe6\x68\x22\x0c\xcb\x53\x3c\x1e\xbe\xcb\xaf\xf1\x06\x5c\xcf\x21\xfb\x9d\xaa\xa1\xcf\x33\xc9\xe0\xa6\x7d\x8c\xad\xa6\x5c\x70\x99\x83\x11\xd4\x14\x60\x15\xa4\x1a\xa9\x45\xa0\xc0\xb8\x8e\xe1\x83\x02\x37\x8b\x50\x09\x54\xa7\x05\xda\x55\x85\xa3\x66\x65\x96\x2a\x86\x73\x97\xf9\xf7\x69\xbf\xf5\xc7\x6c\xd2\x2c\xc7\x07\xd0\xc4\xea\x7a\x88\xe7\xb2\x49\x79\xc0\x27\x53\xba\x24\xd1\x0e\xad\x99\x7a\x51\x72\x3b\x68\xe8\x6f\xb5\x2a\x2b\x1b\x1d\x19\x04\x16\x8a\xad\x5c\xcc\xaa\x42\xc9\x2e\x9d\xc4\xc2\x26\x77\x74\xc8\x29\x56\x4b\xd4\x82\x3a\xc7\xe6\xb2\x3a\x24\x0e\x56\x48\xdc\x1d\x5f\x22\x69\xb9\xdb\xeb\x5e\x35\xd0\x0e\x7a\xdf\x0d\x65\xda\x1e\xe7\x7b\x64\xda\x5f\x56\x6c\x66\x43\x2e\xf3\x90\x48\xbc\x21\x9b\x42\x5d\x52\x51\x3b\xa1\xb6\x61\xd2\x5a\x6b\x94\xf6\x03\xd5\x39\xde\x43\xcf\x65\x55\x3b\xec\x8d\xcb\xc8\xb7\x6b\xc9\xd0\xbd\x87\xcd\x5a\x2c\xa8\xb1\x2f\xdc\xca\x9b\x2c\x24\x53\x12\x79\xa3\x76\x00\xf0\x36\xa6\x5e\xb8\xc9\x56\xe6\xe1\xf9\xa8\x75\xef\xcc\x7a\x19\xed\xb0\x6d\xf3\x9c\xc1\xc3\x51\x97\x0a\x65\x6e\x8b\x8d\x41\xc4\x4f\x19\x10\x04\x51\x97\xb0\x0f\x78\xe1\x6d\x3c\xda\x04\xc6\x0f\xa3\xad\x74\x41\xe0\x03\xdd\xe0\x82\xd1\xa5\x2b\x7d\xb8\x77\x94\x39\x6b\xe2\x8f\x80\x90\x51\x37\x18\x00\x10\x47\x95\x5b\x27\xd3\x76\xbb\xdf\x78\xd2\xe5\x21\xd3\x75\x4e\xdf\x0e\x7d\x5b\x8c\x6d\x81\x32\x0c\xa3\x7e\x66\xdf\x6a\xb5\x68\xaf\x94\xc4\x35\x75\xed\xf6\x00\x1e\x24\xfb\x67\xaf\x16\xf0\xc5\x20\x67\x4a\x6d\x5a\x84\xb8\x95\xd3\x3d\x38\x4a\x60\x2c\x54\x1e\xe2\x3a\xd7\x4e\x28\x23\xc8\xa8\x30\xbd\xd5\x5d\xc7\x45\x2a\x94\xf1\xa2\xf5\x42\x6d\x2c\x34\xda\x5a\xcb\xd6\x67\x4b\xe6\xeb\x11\xe3\x1b\x05\xee\x49\xdd\x3f\x2d\x68\x9a\x59\x37\x04\xdc\x0e\xde\x9c\xa1\xe1\x42\xa8\xf4\x7a\xec\xcd\xfd\xec\x19\x6e\x0d\x84\xbb\x6e\x52\x3f\xf7\x6c\xdc\x26\x37\x42\x40\x02\xbf\xbc\x7f\xf3\x3a\x6e\x95\xcb\xb3\x55\xe8\xda\xd7\xa3\xcf\x26\xdc\xff\x83\xc9\xb5\xb7\xa8\xbb\x07\x05\x52\x86\xda\xb8\x47\xab\x17\xd1\xaf\x5c\x32\x32\xfd\xc6\x21\x66\xad\xc1\x77\x98\xbb\x86\xe9\xba\xec\x14\x88\x7b\x64\x88\x67\x6d\x5b\xf7\x55\x6d\xf7\xea\x7e\xd4\x9c\x6e\xd4\xc1\xfb\x26\xe9\x6e\xd6\xea\x87\x89\xb0\x0b\xda\x09\x71\x1d\x7a\x2d\xae\x8d\x29\xe7\x3b\xd4\xe5\x29\x3d\x38\x23\xb5\x93\x66\xdb\x14\x2f\x8e\x49\xd2\x43\xbe\xf2\xd3\x69\x27\xcc\x23\x29\xb6\xbc\x36\x32\xfd\x43\xc4\xd9\x4d\xb0\xff\x0e\x69\x9e\xfc\x15\x00\x00\xff\xff\xd8\xc8\x3e\xe2\xdd\x11\x00\x00"

func publicJsApplicationJsBytes() ([]byte, error) {
	return bindataRead(
		_publicJsApplicationJs,
		"public/js/application.js",
	)
}

func publicJsApplicationJs() (*asset, error) {
	bytes, err := publicJsApplicationJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "public/js/application.js", size: 4573, mode: os.FileMode(438), modTime: time.Unix(1483478471, 0)}
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
	"public/js/application.js": publicJsApplicationJs,
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
	"public": &bintree{nil, map[string]*bintree{
		"js": &bintree{nil, map[string]*bintree{
			"application.js": &bintree{publicJsApplicationJs, map[string]*bintree{}},
		}},
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

