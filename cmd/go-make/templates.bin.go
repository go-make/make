// Code generated for package main by go-bindata DO NOT EDIT. (@generated)
// sources:
// templates/editorconfig.tpl
// templates/gitignore.tpl
// templates/golangci.yml.tpl
// templates/makefile.tpl
// templates/vscode-settings.tpl
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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _templatesEditorconfigTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x8e\xc1\x8a\xc2\x40\x0c\x40\xef\xf3\x15\x81\xbd\x2d\xa5\x85\x65\x4f\x0b\x7b\xf3\x0f\x3c\x8a\x0c\xb5\x93\x69\xa3\x33\x49\xc9\x44\xa4\x8a\xff\x2e\x03\xc5\x83\x5e\x3c\xe6\x3d\xf2\x92\x2f\xd8\x22\xc2\x64\x36\xff\x75\x1d\x06\x32\xd1\x41\x38\xd2\xd8\x8a\x8e\x10\x45\x21\x8b\x22\x10\x47\x01\x61\xb0\x89\x0a\x44\x4a\xe8\x9c\x8a\x18\xfc\x83\xe9\x19\x9d\xdb\x7d\xef\x1d\x71\x40\x36\x5f\x6c\x49\x58\x45\x7f\x78\x22\xba\x56\xf2\xeb\x90\x83\x97\xe8\x13\x71\x9d\x53\x74\xc4\x05\xd5\x7c\x24\xee\x93\x67\xbc\xac\x66\x8d\x6e\x64\x38\xa1\xd6\x73\x6f\xf5\x32\xf7\x03\xbe\xf4\x7f\xea\x1f\xed\x6d\xc9\xa9\x31\xc9\xa9\x39\x16\xe1\x26\x87\xfb\x87\xcb\x8f\x00\x00\x00\xff\xff\x98\xb0\x88\x9d\x0b\x01\x00\x00")

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

	info := bindataFileInfo{name: "templates/editorconfig.tpl", size: 267, mode: os.FileMode(420), modTime: time.Unix(1541150130, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesGitignoreTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2c\x8c\x5d\xaa\x83\x40\x0c\x46\xdf\xb3\x8a\xc8\x7d\x13\xcc\x5d\x52\xd1\x49\x3a\x0d\xd5\xa4\x8c\x89\xad\xbb\x2f\x38\x7d\xf9\x7e\xe0\x70\x68\x84\x81\xaa\x86\x56\xf3\x26\x7d\xaf\xf3\x32\x15\xa5\x73\x5b\x61\x20\x61\x0d\x6f\xc5\xed\xae\x15\x06\x3a\xf6\xe2\x7c\x71\x3e\x6d\xf3\x53\x00\x0e\x31\xf6\x06\xb7\xde\xd3\x08\xe0\x19\x00\x7f\x28\x1f\x29\x19\xf3\xb2\xca\x8e\xaf\xe6\x9c\x45\x18\xdf\x0f\x31\x3c\x3d\xb1\xa5\xfd\xb3\x2c\x59\x51\x0d\x7f\xd6\xeb\xf7\xa4\x90\x3d\xe0\x1b\x00\x00\xff\xff\xf3\xb5\xe6\x0c\x9e\x00\x00\x00")

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

	info := bindataFileInfo{name: "templates/gitignore.tpl", size: 158, mode: os.FileMode(420), modTime: time.Unix(1541150130, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesGolangciYmlTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x5a\x7b\x8f\xdc\x46\x72\xff\x3b\xf3\x29\x0a\xbb\x02\x56\x12\x86\x9c\x5d\x2b\x06\x12\x0a\x4e\xe0\xc7\xda\x11\xce\x46\x02\xcb\xf2\x19\x88\x72\xda\x1e\xb2\x48\xf6\x4d\xb3\x9b\xe8\xc7\xcc\xd2\x51\xf2\xd9\x83\xaa\xee\xe6\x90\xbb\xab\x8b\x0f\x67\x24\xfa\x43\x3b\x43\x76\x57\x55\xd7\xf3\x57\xd5\x73\x09\x3f\xf5\xd2\x41\x2b\x15\x42\x6d\xb4\x17\x52\x3b\x10\x4a\x81\x38\x0a\xa9\xc4\x3e\x3e\x6e\x65\x17\xac\xf0\xd2\x68\x30\x23\xfd\x71\x9b\x4b\x38\x49\xdf\x83\xef\x51\x5a\x68\xb0\x15\x41\x79\x38\x0a\x15\xd0\x95\x9b\xcd\x65\x5e\x07\xad\xb1\x20\xb4\x50\x93\x93\x0e\x6c\xd0\x5a\xea\x6e\x63\x83\xae\x36\x00\x97\xf3\xc6\xda\xe8\x3a\x58\x8b\xba\x9e\x40\x3a\x10\x0b\xf6\x5f\xff\xdb\x3b\xd0\x61\xd8\xa3\xdd\xc0\x72\x5d\x05\x7f\xbf\x61\x1a\x5e\x0e\x68\x82\x5f\x71\xda\x02\x96\x5d\x09\xaf\xae\xdd\x16\x3e\x1f\xb6\x33\x1f\xe9\xe0\x66\xd8\x00\x34\x28\x1a\x25\x35\x56\xf4\x95\xa9\xe0\xbd\x24\x31\x1a\x84\x53\x8f\x1a\x84\x07\x85\xc2\x79\x30\x1a\x41\x3a\x17\x10\x4e\x82\x4e\x13\x74\xb3\x26\xb7\x81\xf8\xde\x15\x44\xa2\x20\x12\x15\xdc\x44\xa2\x52\xd7\x2a\x34\x08\x1e\x9d\x67\x1d\x3b\x30\x16\xb4\xf1\x2b\x12\xde\x06\xdc\x00\x2f\x72\x55\xfc\xc6\xbb\x95\x24\xfe\x2d\xec\x83\x54\x0d\x78\xd1\xb9\x2d\x9b\x46\x49\xed\xd1\x3a\x08\x0e\x41\xfa\x12\xbe\x39\x93\xc2\x61\xf4\x13\x6f\x2c\x37\x10\x37\x16\xb4\x91\xb4\x0d\x50\xc0\x30\x79\xd1\x45\xea\xa7\x5e\xd6\x3d\x34\xd2\x3a\xf0\x06\xdc\x41\x8e\x15\x59\x73\x82\x93\xd1\x57\x1e\xf6\x18\x95\xf9\x2b\x36\xaf\x79\x7d\x2d\x34\x73\xb4\xd8\xe1\xfd\x08\x3d\x5a\xac\xa0\x43\x8d\x56\x78\x6c\xca\x97\xdb\xfc\x86\xec\x37\x8e\x4a\x62\x03\x46\x43\x1b\x94\x82\x51\xf8\xfe\xf5\xca\xe0\xec\x29\x6b\x89\xb7\xb0\x0f\x1e\x34\xde\xfb\x28\x95\xb0\x08\x42\x9d\xc4\xe4\x58\xba\x11\x1b\x90\xba\xc1\x11\x75\x83\xda\xab\x89\xe9\xb5\xd6\x0c\xe0\xc9\x85\xa3\xc3\x5d\xb9\x48\x3a\xba\x17\xc0\xdf\x1d\x51\x37\xc6\x3e\xdb\xd2\x22\xdb\x7c\x18\x85\xf5\x13\x7d\x43\xe7\x1b\xe1\xc5\xb3\x2d\xe0\xbd\x18\x46\x85\xee\xd9\x16\xbe\x33\x0d\x8e\xf4\x81\x34\xe7\xa5\x7e\xb6\x01\xe6\x5d\x90\x40\x59\x87\xce\xd6\x3b\xbc\xf7\x68\xb5\x50\x1f\x94\xdc\xbb\xf4\x5c\x04\x6f\x66\x7d\x7c\xd8\x4f\x1f\x86\x89\x5e\x2f\xb5\x1d\x5d\xe0\x81\xba\xa5\x52\x4b\x6d\x47\x35\x44\x97\xca\xc7\xc3\x21\x12\xc9\x96\xb1\x38\x1a\x4b\x5a\x9f\x6d\xff\x49\x7d\x7a\x32\x14\x48\xc7\x04\xb4\x01\x8d\xd8\x90\x04\xd9\x35\x39\xd6\x97\x92\x47\x21\xb7\x70\x4a\x91\x1f\x95\x0d\x16\x6b\xd3\x69\xf9\x2b\x32\xa1\x27\x76\x94\xf0\xa6\x05\xe9\xaf\x1c\xf9\x37\x8c\x14\x3d\x08\x0a\x3d\x04\x07\x07\x6d\x4e\x65\xd6\x25\xaf\xce\xca\xbc\x28\x5f\xbe\x7f\x5f\x0e\xd3\xfb\xf7\x65\x67\x9e\x5d\xa4\xa7\x4a\xee\x77\x7b\xd1\x94\x9d\x89\xda\xdb\x4f\x8b\x78\x21\x1d\x38\xf4\xcc\xcf\xa1\x27\x41\x47\xe1\x1c\x48\x4f\xe7\xba\xe8\x4c\x0c\x9c\x62\x30\xcd\x17\xff\x19\xbd\xe2\xbf\x2e\x4a\xf8\x96\x54\x49\x6f\x7b\x54\x23\x0c\xa6\x09\x0a\xdd\x45\xf4\x13\x92\x5c\x1f\xcd\x01\x9b\x98\xd5\x78\xaf\x45\xd1\x18\xad\x26\x72\x1d\x84\xce\x40\x6d\x86\x41\xe8\x86\xb4\xdc\x48\x27\x94\x32\x27\x3a\x7c\x32\x11\xc8\x61\x54\xb2\x96\x7e\x56\xd0\x20\xbc\xac\x21\x8c\x8d\xf0\x52\x77\x14\xc9\x9d\x29\x07\xd3\x40\x83\xae\xb6\x72\x8f\x0d\x88\xbd\x39\x62\x09\x6f\xb4\xf3\x28\x9a\x2d\x9d\xa1\x15\x52\xb9\x94\x84\xf4\x04\x75\x2f\x74\x87\xd1\x7c\xde\x64\x0a\x14\x1b\x64\x49\x72\x01\x4e\xe0\x0e\x3d\x33\x91\x0e\x06\xe3\x48\xe9\xd8\x06\x45\x3b\xea\x1e\xeb\x03\xf8\x5e\xf8\x99\xbd\xc1\xec\x0e\x3e\xfa\x03\xcb\x48\x46\x77\xa1\xee\x41\x38\x90\x1a\x04\xd7\x03\xa9\x83\x09\xf4\xdd\x63\x97\x2a\x00\xa9\x80\xc2\x87\xd8\xb9\xc9\x79\x1c\xca\x4f\x2b\x31\x86\xdf\x23\x15\x0a\xe7\xc2\x40\xa1\x40\x62\xd1\xbb\xb8\x2e\xa6\x08\x69\xb1\xf6\xc6\x4e\xd0\x1b\xd5\x38\x7e\x5d\x1b\x4b\x0f\xa1\x36\xa3\xa4\x24\xda\x42\x4e\x04\x35\x7d\x67\xab\x74\xda\xd8\xac\xa8\x1e\xcf\x0b\xa6\xa4\xef\x54\x90\xa4\x4e\x7a\x88\x52\x27\x47\x28\x1a\x73\xd2\xca\x88\x86\xa4\xc6\x0a\xb2\xf1\x3f\x5a\x64\x4f\xfe\x98\x04\xe4\xd2\x16\xfc\x18\xfc\x27\xea\x62\x7c\x19\xdd\xaa\x36\xca\x58\x6c\x0a\xaa\x34\x45\xac\x61\x1f\x97\x9f\xff\xec\x8c\xfe\xe8\xc5\xfe\x23\xdb\xc8\xf9\x49\xe1\x47\xaa\x1f\x45\xad\xe4\x20\x3c\xae\xaa\xc4\xc5\x13\xd4\x28\x5e\x5a\x63\x07\xe1\xab\xa7\x98\xc5\xe0\x19\xad\xd4\x9e\x8a\x46\x54\x5c\xac\x71\x64\x20\xce\x31\x4f\x55\x22\xde\x51\xf0\xeb\x48\x6f\x55\x96\x66\x7a\x1e\x2d\x68\x31\x20\xa9\x94\x34\x8e\xba\x21\x06\xb1\x5a\x7a\xbc\x7f\xb2\xca\x45\xda\x71\x77\x41\xbb\x33\xe9\xcd\xe5\x03\xdc\x91\x5c\x9a\x85\x76\x23\xd6\xb2\x95\x75\xae\x7d\x9b\xf4\xb7\xc8\x8b\x48\xe1\x68\x2d\x2b\x32\xe6\x96\xcb\x94\x25\x29\xc6\xa8\xae\x18\x1f\x43\x21\x85\x22\x5a\x6b\x2c\x7b\x83\x9f\x46\x24\x87\x44\x36\x60\x05\x77\x02\xaa\x2f\x60\x5f\x3e\xff\x61\x7a\xeb\x6d\xa8\xfd\x8b\xbb\xd7\x89\xe2\xe2\x38\xad\x50\x0e\xab\x18\x30\xb5\x70\xc8\xf5\x8a\x52\x53\xce\xcd\x8b\xa4\x55\xf2\x76\xe6\x5e\x10\xb7\x82\xb8\xd9\xc4\x8e\x09\x6d\x9e\x12\x59\x38\x27\x3b\x3d\xa0\xf6\x0b\x81\xbd\x81\xbd\x12\xfa\x00\x9c\x98\x65\x2b\xd1\x56\x70\xa7\xc3\xb0\x85\x0f\x24\xb7\xf3\xb6\x36\xfa\x58\x7e\xe9\x8d\x7c\xae\xc3\xf0\xd6\xdb\xdf\x55\x7e\x66\xfe\x40\xea\x7f\x6f\x70\xb4\x58\x53\x2d\xf8\x8f\x18\xe2\x85\xc3\x51\xc4\xe2\x90\x51\xcc\x28\xa8\xa8\x9b\x96\x5d\x85\x9c\x16\xc6\x43\x57\x31\x66\x48\x64\xe8\x05\x7f\x27\xf9\x82\x4b\x75\x8a\xc3\x9a\xfd\xcc\xb1\xdb\x4a\x4d\x1b\x4b\x78\x9e\x0f\x73\xd1\x0e\xbe\x2a\x5f\x5e\xbc\x28\x13\x1d\x87\x08\xbd\xf7\xa3\xab\x76\xbb\x4e\xfa\x3e\xec\xcb\xda\x0c\xbb\x83\x74\x12\xd5\x61\x97\xdd\xe4\xd2\xf7\x58\x9c\x25\x2f\x06\xf4\xbd\x69\x18\x46\x36\xe8\x29\x0b\x33\xbd\x28\x40\x05\x91\xcb\x56\x9a\x9d\x34\xc1\x4b\x55\xfd\xe9\x47\x14\x4d\xf9\x32\x2b\x81\x30\x0e\x09\x2c\x56\x60\x9a\xbc\x4d\xcc\x3a\x68\x83\xae\x63\x16\xf2\x06\xf0\x3e\x56\x60\x2e\x20\xd9\x35\xff\xca\x23\x44\x1a\x52\x77\xc5\x99\xf4\xc3\x03\x5c\x66\x4e\x15\xec\x48\xc8\x9d\x37\x3b\x12\xb1\xf4\xf7\x54\xab\x3a\x73\x44\xff\x64\xc8\xb8\x5e\x34\x5c\xe2\x8e\xc2\x4a\x0a\x49\xb7\xf0\x83\xf8\x52\xea\xee\x9c\x1a\xa2\xe0\x29\x68\x47\xb4\x19\xd2\x58\x7e\xb7\x8c\x54\xfa\xc7\x89\xa0\xad\xa8\x56\xa6\x65\x6c\xe4\x2d\x35\x09\x70\xd7\x19\xf0\xc6\x28\x38\xa2\xe7\x82\x7d\xc7\xa8\x09\x13\x5e\x49\x1b\x5c\x22\x05\xac\x58\x47\xb4\x9e\xdc\x9c\x78\x9d\x69\x3c\xce\x31\xa4\xb4\xbb\xbc\x6c\x25\x77\xfc\x57\xc0\xf3\x85\x1d\x3a\xa3\x84\xee\x6a\x39\x7f\xe0\x84\xb6\x1b\x0f\xdd\x4e\x99\x8e\xbc\xc3\x95\xdf\x9b\xee\x45\xf9\x46\xb7\xa6\xfd\xdb\xc9\xfc\x51\x58\xfd\x3b\x90\xb9\xa5\x14\xf2\x3b\xd0\xf9\x56\x78\xa1\x5a\xf6\x1d\x5a\x91\x9d\x67\x90\x5a\x0e\x42\xcd\xb0\xb1\xe6\x38\x4f\x50\x76\x55\x0b\xae\xcb\x7f\xe0\x3d\x83\xd4\xc5\x79\x75\x95\x9e\x77\x86\x62\x2d\x3b\x14\xc3\xaa\x76\x82\xd8\x58\xf1\xbb\x58\xbe\xee\x0a\x77\x97\xca\xee\x96\x7d\x70\x91\xb1\xa2\xc7\xa5\xad\x55\xae\x3d\x9d\x91\x03\xf9\xb7\xcb\xc4\xa9\x8c\xa7\x47\xb0\xc7\x4e\x72\x77\x1a\xa9\x8f\x16\x5b\x79\x0f\xa2\xa5\x2a\xf7\xca\x36\x05\x77\x0e\x30\x8a\xfa\x20\x3a\x74\x39\xa3\x32\xe0\x15\x9f\x4e\x7a\x4c\x25\x05\x8e\x32\xb5\x50\x45\x7e\x54\xc1\x42\xf5\xc6\x76\xbb\xd1\x9a\x3f\x63\x1d\x43\xb2\x9e\x6a\x65\x1e\xeb\xb5\xa1\xbc\x42\xcd\xca\xbd\xf4\x13\x79\x73\x8c\xd7\x2d\xbc\xba\x5e\x62\xe4\xe7\x04\xfc\x4f\xc8\x98\x7d\x18\xa8\x2c\xdf\x5c\x17\x9f\x5d\xbf\x58\xe8\x3c\x13\xa9\xe0\xd5\xf5\x06\x60\x10\x4a\x76\x1a\x9b\x59\x31\x5c\xe3\x1d\x97\xbf\xa8\x8f\x81\xb2\x30\xb6\x2d\xd6\x5e\x1e\x11\x06\x1c\x08\xa1\x29\x31\x51\xa6\xc8\x4d\x2d\x57\x85\x47\x66\x08\x5d\x87\xce\x17\x1a\x4f\xb3\x25\x9a\x30\xaa\xcc\xcb\x9b\x03\x6a\x07\xb5\x09\x9a\x11\xbc\xb7\xb2\xeb\xd0\x66\x78\x72\xf3\xf9\xf5\x43\x8a\xbe\xb7\xe8\x08\x1c\x56\x70\x73\x7d\x1d\x15\x66\xb4\x7b\xe4\x88\x0a\x75\xe7\x7b\xc6\x0f\xde\x92\x65\x79\x95\xd0\xa4\xb0\x87\x34\x49\x2d\x0a\x75\x05\xaf\x1e\x10\x31\x75\x9a\x38\xe0\x63\x19\x3f\x41\x67\xb1\x25\xd2\x6b\x70\xec\x82\xb0\x49\xbd\xe4\x1c\x8c\x03\x2a\x2a\xe5\xf5\x81\xbe\xc7\x52\x13\x7b\xb2\xa2\x33\x85\x35\xc6\xe7\x2a\xcb\xd9\x32\xf9\x5d\xce\x9e\xc5\xd2\x7b\x1a\x71\xc4\xba\xeb\x77\x9d\x29\xdc\x88\xa7\x1d\xfd\x47\x46\x95\xce\x8d\xa8\x66\x45\x7f\x9d\x40\x34\x3f\xe4\xac\x17\x1c\x69\x85\xfd\x12\xd9\x57\x31\x1d\x94\xe2\xf6\xdd\x5b\xb2\xeb\xbb\x3f\xe4\xb2\xba\x18\x33\x78\xc3\xa3\x00\x01\x1a\x83\xb7\x42\x71\x85\x40\x3f\x91\xaa\x6f\x75\xa7\xa4\xeb\xf3\xae\xb7\xa9\x31\x49\x4c\xbc\x21\xba\xdc\xf3\x66\x4c\x4f\x65\xff\x2b\x2b\xbd\x74\xfd\x2c\x1a\x11\xba\x22\x44\x1b\xec\x15\xed\xe1\xcf\xf6\xaa\x3c\xc7\x11\x56\xf0\xee\x0f\x39\x08\xb9\x44\x17\x27\x63\x9b\x39\xb4\xb9\x53\x37\x03\xd2\xc3\x0d\x80\x3a\xab\x61\x10\xf7\x8c\x86\x93\x7b\x6c\x13\x34\x56\x46\x93\xd7\xe5\x76\xfc\x71\x8b\x2d\x1d\xdc\x7c\x76\x9d\xcf\x75\xf5\xde\x5f\xd1\x23\x76\x09\xea\xe4\x1c\xdc\x50\xb3\x66\x45\x4d\x19\xe3\xec\x15\x5b\xee\x4e\x6a\xa1\x89\x6a\xec\xe6\x9a\x79\x76\x06\x5e\xec\x8b\x93\x6c\xc8\x4b\x39\x93\x25\x07\xd1\x58\x44\xe9\x2a\xb8\xe1\x00\xe5\x40\x11\x7b\x88\x6b\xa5\x06\x37\x8a\x9a\x9a\xee\x2c\x9d\x37\x70\x13\x65\x9b\x49\x56\x3c\x9c\x0a\x9a\xa0\xd4\x1c\x6c\x16\x45\x9a\x73\x09\x4a\x5c\xa3\x35\x9d\x15\x03\x3c\x27\x7c\x4c\x40\x65\x6f\x85\x9d\x5e\xb0\xcc\x09\x0c\x44\x02\x80\xf7\x09\x17\x9e\x81\xa7\x7b\xfd\x08\x51\x66\xf5\xfc\xf2\xcb\x2f\x15\xc8\x16\x26\x13\x00\x35\xd7\x5a\xbf\xe8\x53\xb7\x99\x2a\xab\x3b\xa3\x0e\x50\x26\xc2\x24\xa2\x54\x8c\xc6\x49\xca\x35\x11\xa4\xe3\xbd\x07\x6c\xa4\x37\x76\x36\xb1\x4c\xf3\x86\x5a\x28\x85\x11\xb8\xb9\xb0\x6f\xa4\x25\x1a\x7c\x34\xca\xa7\xd4\x53\xd7\x82\x90\x6d\x2b\x75\x13\x91\x02\x04\x47\xc1\x54\xc2\x97\x4a\x2d\x49\x2f\xfb\xdc\x0c\x9e\xd8\x54\xab\x8a\xc8\x0c\x89\x2c\x35\xc3\x8b\x76\x75\x36\x6a\xb6\x32\xa3\xac\x05\x66\xca\x2a\x3c\x07\x76\xd0\x54\x31\x86\x7c\xa2\x37\x9a\x1a\x1d\x7f\xd6\xf5\x8c\xea\xb6\x8f\x35\x4d\xa1\x15\x73\x51\x40\xd2\x85\x26\x48\x19\xc7\x52\xd9\xac\xbb\x64\xcf\xb9\xc8\x4d\x26\x58\xb6\xfe\x5f\x61\x27\x96\xf0\xff\xde\x50\xf3\x59\xb8\xd3\x6b\xa3\xb7\xff\x7f\xdb\x4b\x8b\x03\x36\xf6\x8c\x97\x07\x71\x40\x10\x3a\xf5\xbc\x32\x22\x7c\xe8\x85\x8b\xf5\x72\xdd\x74\xfb\x5e\xe8\x95\x76\xe3\xdc\xc2\xf3\x7a\xa6\x0c\x16\x7d\xb0\x7a\x1d\x57\x29\xfe\x07\x71\xcf\x20\x3f\x37\xe4\xfc\x78\xb4\x28\x94\x32\x75\xb5\x34\xe7\x09\xa1\x31\xb1\x95\xcb\xf5\x3f\x26\x7a\xe6\x9d\x1a\xf7\x3d\xb6\x24\x61\x63\xe8\xc5\x88\x96\xc7\x08\x04\xd9\x46\x6b\x5a\x49\x39\x38\xbb\xc8\xb7\xc6\xc6\x99\x52\x72\xaa\x14\x3d\x09\xde\x30\xfb\xf3\x10\x93\x1e\x0d\xc2\x07\x8b\x9c\xd0\x06\xf9\x2b\x1b\xa7\xcc\xfd\xc1\x8f\xd1\x83\xf2\xc6\x38\x3b\x49\x08\x81\xbb\x17\xa3\xd5\x44\x86\x62\xf0\x86\xa0\x8c\x19\xd3\x98\xa8\x17\x47\x24\x27\x4f\x3a\xda\xed\x2d\x8a\x83\xdb\xa5\x29\x15\xba\x5d\x67\xbc\x71\x69\x18\x31\x64\xe9\x7f\x5a\x03\xc3\xf2\x8c\x0c\x71\x46\x23\x00\x96\x1c\xa0\x60\x66\xf1\xe9\x6f\x13\x35\xee\x8b\x42\x3e\x8d\x41\x5b\x63\x33\xd9\x08\x8e\x7e\x1b\x5d\x0a\x91\x44\xf5\x09\x4c\xd5\x99\x9a\xea\xe5\x6c\xf5\x3f\xf2\x70\x39\x0e\x91\xc0\xf5\x26\xa8\x86\x8c\x11\xe3\xb9\x79\x9d\xa2\x6a\xcf\xe0\x71\x2f\x75\xae\x3f\x57\x8d\x74\xbc\xa2\x88\x5b\xaf\x5e\xcf\x35\x7b\xd1\x7d\x9a\x22\x32\x2b\x13\xe0\x90\x66\x67\x8e\x68\x8f\x12\x4f\x97\x71\x5f\x91\xbf\x67\x95\xe7\xa1\xe3\x69\x29\x97\xb0\xb3\x44\xb1\x4b\xfb\xee\xfb\x0f\xdf\xdc\x7e\xf5\xee\xbb\x2f\xf2\x71\x1e\x44\xad\x0d\xfa\x2e\x51\xfc\xea\x8c\x6d\x33\xb2\x76\x3e\x5e\x4e\x45\xe2\x69\x60\x90\x8d\x9e\xf8\xa4\x73\x2d\x21\x01\x1b\xec\x67\xa1\xbe\x36\xe3\xb4\xf9\xcb\xea\xcb\xda\xf9\xb4\xfe\xd6\x6c\xae\x56\x51\xcb\x73\xf8\x3c\x76\x59\xeb\x79\x25\x0f\xdf\x99\xfc\x10\x9c\xcf\xd2\xdc\xc6\x2c\x3c\x04\xe5\xe5\x78\x3e\xe1\x7e\x4a\x37\x40\xbf\x45\x77\x77\xbc\x6a\xd1\x31\xd3\xd6\x88\x43\x98\x5a\xd6\xd3\xed\x7c\x59\xb0\x0c\x91\xb5\xfd\x97\xfd\x61\x72\x85\xf3\xa7\xcb\x98\x0a\x8a\x7f\x02\x87\x5c\xa8\xe0\xe2\x27\xd1\xb9\x8b\xc8\x20\xeb\xe7\x7c\x03\x45\x47\x5e\x24\x9b\xcd\x7a\x32\xb0\x9a\x1f\x08\x97\x26\x40\xf9\x88\x89\x40\x2d\x46\xff\x3d\x21\x40\x5a\x3e\x04\xc7\x86\x39\x0a\x25\x9b\xd9\xbd\xa2\xf7\x69\x31\xe0\xdc\xeb\x72\x11\x73\xff\xaa\xd5\xb4\x88\x7a\x58\xb9\x43\x35\x2f\x76\xf2\x57\xfc\xe9\xdc\x63\xbc\xfa\x6c\xb3\xb9\xcc\x93\xc7\x6a\x73\x39\x9f\x2c\x7e\xe6\x7b\x35\xec\x04\x73\x9d\x9f\xf0\x58\x65\xb1\xb6\x10\x4a\xe5\x3a\x42\x4f\x93\x4f\x2c\x48\xa4\x16\x6c\x7e\x90\x13\xc4\x72\xf9\x43\x2a\xa3\x45\x87\xde\x9d\xa9\xec\x43\xe7\xe6\x2f\x11\x63\xf1\xd7\x56\xb8\xb9\x9f\xd8\x6c\x62\x4b\x1e\x87\xd4\xdf\xa7\x80\x8a\x9e\xe8\xd6\x13\xdc\xe5\xa0\x6a\xbb\xb8\x5b\x7a\x90\x51\x2f\xe1\x2b\xea\xa4\x97\xb7\x73\x8f\x6e\xe6\xa8\x2c\x51\xd3\x90\x63\x24\x8f\xbf\x46\xe1\xa9\xd0\xbb\x6d\xbc\x31\xf5\x19\x2a\xe7\xa0\x21\x56\x77\x69\x71\x11\x1c\x16\x89\x40\x3a\xcd\x5d\x49\xf9\x86\x65\x12\x4a\x6d\x16\xe3\xae\xe5\x04\x73\xe6\x02\x78\x8f\x75\xf0\x08\x77\x8f\x92\x0d\x14\x05\x4f\x9a\x36\x30\x0f\xcc\xf2\x25\x55\x4a\x0a\x5c\x81\x62\x49\xf5\xf0\xbc\x7c\x09\x2f\xfe\xd9\x50\x25\x85\xa0\x33\x54\xb8\x88\x33\xf2\xdb\x3c\x97\x7b\x70\x4b\x30\xa2\x2d\x46\x41\x4d\x07\x7d\x8a\x2e\x15\x3f\x33\xaa\xa1\x08\xa5\x2f\xce\x04\x5b\xe3\x59\x90\xc2\x86\xf9\xce\x2c\x13\x47\x6e\x6f\xe6\xcb\x60\x56\x77\xba\x5f\xa7\x0a\xc2\xb7\xc9\xe9\x66\x2e\x1d\x83\x18\x57\xf0\x81\xde\xbc\x2f\x3b\x93\xdc\x7d\xf6\xeb\xf3\x50\x28\xcd\x1f\x16\x4f\xf2\xc0\x71\xf1\x88\x5a\xf8\xd5\x1e\x87\xf5\xe6\x81\x84\x07\x6d\x4e\x7a\x2d\xe2\x28\xac\x97\x42\xa9\x09\x7a\x61\x9b\x22\x5e\xab\x50\xcc\x9a\x06\xb7\x19\xc2\x71\x42\x96\x8e\x61\xab\x73\x72\x1f\x1b\xc7\xec\x31\x47\x29\xe0\x42\xf3\xe8\xe9\x22\x9b\xe3\xc1\x21\x99\xa3\x16\x6a\xd7\x0f\xa2\xde\x25\x31\x49\xc5\x15\x5c\x9c\x50\x1c\xa0\xb6\xd3\xe8\x09\xce\x8c\xbd\xac\x61\xb4\x72\x60\x08\x7b\xf1\x97\x94\xf2\xc4\x01\xd9\x04\xce\x0b\x2f\xeb\x98\x74\x06\x74\xdc\x5a\xcc\x17\x99\x8f\xe8\x2c\x56\xaf\xe5\x7a\xfb\xe5\x3f\x5e\x5f\xbf\xaa\x2e\x1e\xf2\x50\x04\x5e\xd3\x75\x30\xc3\x02\xea\xa6\x19\x58\x26\xa8\x5b\xe5\x9b\xd8\x4f\x33\x55\x2a\xdb\x2a\xba\x56\x05\x17\x7f\xa2\xda\x3e\x6f\x85\xe4\xb8\x6f\x1e\xc7\x70\x0a\xdf\x1c\x83\x77\x7f\x4b\x20\x2f\xf2\xc1\xef\x1e\xb7\x31\x11\xad\x2f\xc4\x49\x61\xcb\x1c\x94\xee\xa1\xca\x45\x68\x3d\x4e\x29\x51\x13\x3f\x88\x7b\x39\x84\x21\xab\x3e\x8e\x7e\x46\xb4\xfc\x8b\x90\xa8\xe2\xb9\x05\xbb\xa6\xff\xd2\x39\x57\x03\x83\xcf\x79\x5e\x40\xd0\x3d\xfd\x46\xe4\x1c\xf6\x15\x5c\xaf\x19\x45\x0e\x39\xfd\xba\x73\x7b\xe2\xc4\x10\xb3\xf1\xff\xce\xef\x55\x66\x47\x7b\x12\xcf\x99\xd1\xdb\xde\x9c\x22\xbe\xd6\x78\x82\xfc\x52\xb6\xe9\x37\x02\x04\xd0\x82\x76\x5e\x74\x5c\x42\xf9\xda\x19\x8c\x85\xa0\xbd\x15\xf5\x61\xfe\x65\x00\xd3\x62\x32\xbe\x37\x0e\xe7\xa5\xfc\x7b\x8d\xf9\x47\x0c\x48\xa8\x95\x57\xe5\xf7\x52\xc3\xbf\xdc\x7e\xf9\xcd\x7f\xaf\x16\xa6\xab\xe3\x38\x45\x75\x81\xd4\x93\xae\xae\x93\xc5\xda\x75\x97\x17\x6f\xd1\x97\x1e\x20\x35\xa7\x06\xc9\x17\xd2\xf1\x77\x33\xc2\x76\xc8\x09\x65\x2f\xa8\x4f\x7e\x33\xff\x28\xc1\x8a\xda\xcb\x5a\xf0\xbd\x38\x4f\x77\x95\x9a\xb7\x66\xb5\xa7\x0b\xe9\xc1\x50\x5a\x89\x67\x6d\x97\x12\x54\x30\x84\xba\x87\x3d\x92\x63\xa6\x46\x8b\x7f\x0b\x90\x09\x48\xcd\xfa\xcd\x5d\xf6\x6a\x70\x36\xcf\x48\x78\x18\xba\xf0\xb6\xa7\x6c\x03\xb5\x45\x9e\x26\xc7\x21\x74\x27\xa9\xa3\x3b\x4a\xc7\xd1\xf8\xe3\xed\xcf\x77\xf1\x12\x1f\x4f\x05\xc5\x69\x61\xf1\x58\xc1\x8f\xb7\x3f\xff\x06\x82\x52\x33\xb5\x51\xf8\xba\x8f\x7e\xe6\x30\xfe\x40\x89\xf3\x67\xb9\x26\xcc\xcb\x2a\xc8\x77\x47\xfc\x95\x6f\x90\x36\xff\x13\x00\x00\xff\xff\xbe\x87\x9e\x00\x3e\x26\x00\x00")

func templatesGolangciYmlTplBytes() ([]byte, error) {
	return bindataRead(
		_templatesGolangciYmlTpl,
		"templates/golangci.yml.tpl",
	)
}

func templatesGolangciYmlTpl() (*asset, error) {
	bytes, err := templatesGolangciYmlTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/golangci.yml.tpl", size: 9790, mode: os.FileMode(420), modTime: time.Unix(1587136682, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesMakefileTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x54\x51\x6f\xe2\x38\x10\x7e\xae\x7f\xc5\x68\x93\x13\x89\x04\xc9\x7b\x24\xa4\xe3\x5a\x8e\x45\x6d\x09\x62\xb9\x87\x7d\xaa\x4c\x32\x21\x5e\x1c\x3b\xb2\xc7\xb0\x15\xe2\xbf\x9f\x9c\xa4\xc7\x75\x4f\x27\x6d\x59\x1e\x6c\x79\xe6\xe3\x1b\xcf\x37\x5f\x1c\x4c\x3e\xf8\x63\x01\x0b\xee\x2c\x92\x6b\x01\xd5\x51\x18\xad\x1a\x54\xe4\xa3\x1f\x66\x62\xf8\xbd\xd5\x86\x60\x91\xaf\x67\xdb\xcf\xd9\x34\x8c\x0c\x72\xd9\x72\xaa\x21\x8c\x6c\x8d\x52\x42\x7b\x2a\xe3\xf4\x7c\x4e\x16\x7a\xcd\xa9\xde\xa0\xbc\x5c\x62\xb6\xcd\xd7\xff\x0f\x0e\xa3\x52\x18\x08\xa3\x4a\x18\x4b\x27\x6d\x4a\x08\xa3\xe7\xd9\xe3\xfc\xcf\xe5\xd3\xfc\xe5\x69\xf9\x65\x1b\xc7\x71\xcc\xd6\xb3\xfb\xc7\xd9\x62\xee\x69\xac\xdb\x59\x82\x30\xea\x6f\x11\xa7\xd6\x14\xe9\x78\x1c\x46\xdb\x7c\x1d\xc7\xec\x86\xb6\x3a\x81\x88\xab\x92\x9b\x12\x8c\x93\x68\x6f\x53\x27\x80\x6d\x8d\xd0\xf5\x01\xc4\xcd\x1e\x09\x4a\xac\x84\xc2\x12\x84\x05\xaa\xd1\x1f\xb9\x93\x04\xa2\x02\xa5\xdf\x30\xc2\xb2\x00\x6c\x8b\x85\xa8\x04\x96\xa0\x55\x07\x2d\x74\xd3\x70\x55\x82\x14\x0a\x13\x80\x67\x7e\x40\xb0\xce\x20\x50\x2d\x2c\x94\x1a\xad\x1a\x11\x0b\x80\x7c\x82\xb4\x06\xa9\xd5\x1e\x48\x83\x71\x6a\x0c\x56\x03\xd5\x9c\xa0\x45\xdd\x4a\x84\x93\x90\xd2\x27\x40\x90\xe7\xc7\x23\x9a\x57\x16\xc0\xce\x09\x59\x26\x2c\x59\x7f\xce\x57\x5f\x33\xa8\xb8\x25\xe6\x97\xac\xcf\x40\xa1\x8f\x68\xf8\x1e\x27\xb6\xf6\x63\x97\x42\xd1\xa4\x03\xb1\x00\x66\xd2\xea\xa1\xbd\xee\xbe\x9f\x2a\x27\x25\x54\x9c\x3e\x75\x12\xf6\xe5\xfd\x35\xfb\x6a\x54\x0b\xb5\xff\xa7\x12\x97\xb2\x2f\xc1\xb8\x94\x3f\x56\x1b\xea\x38\x29\x19\x9b\x08\x55\x48\x57\x22\x24\x7b\x3d\x69\xf8\x01\xd3\x6e\xd9\x71\x22\x34\x02\x6d\xd2\x1c\xd8\xfb\xd4\x6f\x19\xbb\xdb\x0b\x82\x42\x6a\x85\x50\x13\xb5\x36\x4b\xd3\xbd\xa0\xda\xed\x92\x42\x37\xe9\xbf\xd1\x89\x47\x0e\x06\xfc\x3d\x66\x77\xa6\x81\x89\xa9\xae\x91\xd4\x03\x6e\xb5\x94\xd2\x27\x3f\x0f\x5e\x90\xe3\x52\xbe\x0e\x5d\x5a\x72\x55\x95\x24\xc9\xaf\x58\x6c\xd4\x51\x8d\xae\x06\x82\x53\x8d\x06\xe1\x55\x3b\x28\x9c\x25\xdd\x08\x8b\x50\x69\xe3\x23\x06\x5a\xa3\xbf\x61\x41\x09\x0b\x60\xa9\x7a\xff\x58\xd1\x78\x5f\xe0\x77\xee\xf7\x31\x9c\x70\x24\x25\x7c\x73\x96\x40\x28\x4b\xdc\x6b\xff\x36\xac\x7e\x50\xdd\x9a\x5d\xb3\xe7\xf3\xc4\xdb\x38\x79\xd0\xc5\x01\xcd\xe5\x72\xa3\x48\x65\xf7\x77\xb0\xae\xf5\x2f\xcb\x6d\xa2\xfc\xd7\x24\x3d\xeb\xb0\xbd\xf3\xc8\x90\xf9\x69\x97\x0c\x14\x3f\xe7\x93\x87\xfc\xfe\x71\xbe\x79\xd9\xcc\x17\xcb\x2f\xdb\xcd\xd7\x6c\xfa\x16\xc9\x37\x8b\x6c\xca\xee\xf3\xd5\x76\xb6\x5c\xcd\x37\x2f\xab\xd9\xf3\x3c\x9b\x56\x5a\x33\x16\x46\x78\xe4\x12\xc2\xa8\xf0\x1f\xc5\x80\xff\xe3\xaf\xe5\xd3\xc3\xb8\x97\xb6\x12\x12\xc7\x61\xf4\x03\x77\x7c\x0d\xe5\x9b\x85\x3f\xbd\x67\xf7\x6f\xa6\x1f\x11\xaa\xf2\x72\x61\x7f\x07\x00\x00\xff\xff\x16\x1b\xa6\x09\x35\x06\x00\x00")

func templatesMakefileTplBytes() ([]byte, error) {
	return bindataRead(
		_templatesMakefileTpl,
		"templates/makefile.tpl",
	)
}

func templatesMakefileTpl() (*asset, error) {
	bytes, err := templatesMakefileTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/makefile.tpl", size: 1589, mode: os.FileMode(420), modTime: time.Unix(1587136682, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesVscodeSettingsTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x24\xc8\xb1\x0e\x82\x30\x10\x06\xe0\xd9\x3e\xc5\x9f\xc6\x44\x5d\x60\xf7\x05\x5c\x0d\x6f\x70\x29\x07\x57\x41\xae\x69\x4f\x09\x69\xfa\xee\x0e\xae\x5f\x75\x00\xd0\xf7\x30\x89\x05\xb1\x40\x68\x1b\x0f\xc4\x09\xd7\x35\x2e\x8c\x37\xdf\x70\xe8\xe7\x92\x19\x84\x6f\x09\x3a\x32\x26\xda\xb0\x8b\x42\xa8\x80\x30\x6b\x22\x13\x24\xce\x48\x59\x5f\x1c\xcc\x9d\xfc\xac\xdd\xdf\xfd\x1d\xfe\x5c\x77\xcd\x4b\x49\x14\x78\x50\xb5\xd6\xd7\xda\x3d\xf4\x49\x26\x03\xaf\xad\x79\xd7\xdc\x2f\x00\x00\xff\xff\x94\x6a\x60\xbe\x87\x00\x00\x00")

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

	info := bindataFileInfo{name: "templates/vscode-settings.tpl", size: 135, mode: os.FileMode(420), modTime: time.Unix(1541150130, 0)}
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
	"templates/editorconfig.tpl":    templatesEditorconfigTpl,
	"templates/gitignore.tpl":       templatesGitignoreTpl,
	"templates/golangci.yml.tpl":    templatesGolangciYmlTpl,
	"templates/makefile.tpl":        templatesMakefileTpl,
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
		"editorconfig.tpl":    &bintree{templatesEditorconfigTpl, map[string]*bintree{}},
		"gitignore.tpl":       &bintree{templatesGitignoreTpl, map[string]*bintree{}},
		"golangci.yml.tpl":    &bintree{templatesGolangciYmlTpl, map[string]*bintree{}},
		"makefile.tpl":        &bintree{templatesMakefileTpl, map[string]*bintree{}},
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
