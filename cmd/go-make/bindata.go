// Code generated by go-bindata.
// sources:
// templates/editorconfig.tpl
// templates/gitignore.tpl
// templates/vscode-settings.tpl
// DO NOT EDIT!

package main

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

var _templatesEditorconfigTpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x3c\xcc\xb1\xaa\x02\x41\x0c\x85\xe1\x3e\x4f\x11\xb8\xdd\x2d\xdc\xc6\x4a\xf0\x29\x2c\x45\x86\xd5\x39\xd9\x0d\x8c\x89\x64\x22\xa2\x4f\x2f\x23\x62\x79\xfe\x0f\xce\x1f\x1f\x00\x5e\x33\x6f\xbb\x69\x42\xd5\xf4\xb8\xb8\x89\x2e\x1b\x8f\x85\xc5\x83\xaf\x1e\x60\x35\x71\x76\xe3\x5c\xb5\xb3\x68\x03\x51\xb8\x27\xef\x39\xe3\x0e\xa2\xe3\xff\x89\xd4\x2a\x2c\x4b\xcf\x67\xc3\x80\xf9\xfc\x4b\xfa\x1a\x65\x4b\xb0\x5a\x5c\x4a\x53\x1b\xbb\x09\xa9\x75\x44\x16\x51\x9b\x5b\x31\x3c\xbe\xf2\x39\x7d\x07\x00\x00\xff\xff\xb8\x6c\x50\x7b\x9b\x00\x00\x00")

func templatesEditorconfigTplBytes() ([]byte, error) {
	return bindataRead(
		_templatesEditorconfigTpl,
		"templates/editorconfig.tpl",
	)
}

func templatesEditorconfigTpl() (*asset, error) {
	bytes, err := templatesEditorconfigTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/editorconfig.tpl", size: 155, mode: os.FileMode(420), modTime: time.Unix(1482882698, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesGitignoreTpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x44\x8a\x41\x0a\x85\x30\x0c\x44\xf7\xbd\xc5\xdf\xfe\x45\x3c\x93\xb4\x43\xad\x68\x06\x9a\x18\xf4\xf6\x82\x28\xdd\xcd\xbc\xf7\x24\x2c\xb3\x60\xfa\xa7\xdf\x37\x71\x3a\xd4\x1a\xd5\x64\x35\xea\x10\x06\xf7\xa6\xf5\xc5\x01\x2d\xec\x29\x33\xd0\xe7\x0a\x59\x7c\xdf\xc6\x7b\x92\xca\x4c\x0d\x5c\xc2\xc3\xd3\x1d\x00\x00\xff\xff\x76\x39\x33\xe4\x6a\x00\x00\x00")

func templatesGitignoreTplBytes() ([]byte, error) {
	return bindataRead(
		_templatesGitignoreTpl,
		"templates/gitignore.tpl",
	)
}

func templatesGitignoreTpl() (*asset, error) {
	bytes, err := templatesGitignoreTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/gitignore.tpl", size: 106, mode: os.FileMode(420), modTime: time.Unix(1482882921, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesVscodeSettingsTpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xaa\xe6\xe2\x54\x4a\xcf\xd7\x4b\xcf\x2f\x48\x2c\xc9\x50\xb2\x52\x50\x52\xa9\x2e\xcf\x2f\xca\x2e\x2e\x48\x4c\x4e\x0d\xca\xcf\x2f\xa9\xd5\xaf\xae\xd6\x73\xcf\x0f\x48\x2c\xc9\x08\x4a\xcd\xa9\xad\x55\xe2\xaa\xe5\x02\x04\x00\x00\xff\xff\x22\xc7\x32\xce\x34\x00\x00\x00")

func templatesVscodeSettingsTplBytes() ([]byte, error) {
	return bindataRead(
		_templatesVscodeSettingsTpl,
		"templates/vscode-settings.tpl",
	)
}

func templatesVscodeSettingsTpl() (*asset, error) {
	bytes, err := templatesVscodeSettingsTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/vscode-settings.tpl", size: 52, mode: os.FileMode(420), modTime: time.Unix(1482883802, 0)}
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
	"templates/editorconfig.tpl": templatesEditorconfigTpl,
	"templates/gitignore.tpl": templatesGitignoreTpl,
	"templates/vscode-settings.tpl": templatesVscodeSettingsTpl,
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
	"templates": &bintree{nil, map[string]*bintree{
		"editorconfig.tpl": &bintree{templatesEditorconfigTpl, map[string]*bintree{}},
		"gitignore.tpl": &bintree{templatesGitignoreTpl, map[string]*bintree{}},
		"vscode-settings.tpl": &bintree{templatesVscodeSettingsTpl, map[string]*bintree{}},
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

