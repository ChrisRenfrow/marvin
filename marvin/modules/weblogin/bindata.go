// Code generated by go-bindata.
// sources:
// layout.html.tmpl
// assets/styles.css
// templates/factoid-info.html
// templates/factoid-list.html
// DO NOT EDIT!

package weblogin

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

var _layoutHtmlTmpl = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x55\x5b\x73\xda\x38\x14\x7e\xcf\xaf\x38\xd5\xf6\x2d\x8b\x35\x40\x9a\xb2\x1d\xe3\x59\x12\x68\x4b\x43\x2e\xdb\x10\x32\xdb\xb7\x83\x2d\xdb\x02\x59\x26\x92\x6c\x60\x3d\xfe\xef\x3b\xbe\x01\x61\x69\xbb\x53\x5e\x2c\x9f\xf3\xf9\x3b\xb7\x4f\x87\x2c\x03\x8f\xf9\x5c\x32\x20\x02\xb7\x71\x62\x08\xe4\xf9\x99\xfd\x66\x78\x7f\x3d\xfd\xfb\x61\x04\xa1\x89\x84\x73\x66\x37\x0f\x86\x9e\x73\x06\x00\x90\x65\x30\x17\xb1\xbb\x04\x62\xb8\x11\x8c\x80\x05\x79\x6e\x97\x67\x27\xcb\xb8\x0f\xd6\xb4\x38\xe7\x79\x96\xed\x4f\x4c\x68\x96\xe7\xb7\xa8\x52\x2e\xe1\x99\xcd\x61\x2c\x0d\x53\x3e\xba\x2c\xcb\x98\xf4\xf2\xdc\xa6\x0d\x03\x30\xe9\x15\x99\xbc\x8e\xe5\x31\x1f\x13\x61\x5a\xda\x6c\x05\xd3\x45\xd0\x56\x8d\xb1\x05\x97\x4b\x50\x4c\xf4\x49\xe5\x0c\x19\x33\x04\x42\xc5\xfc\x3e\x09\x8d\x59\xe9\x0f\x94\x46\xb8\x71\x3d\x69\xcd\xe3\xd8\x68\xa3\x70\x55\xbc\xb8\x71\x44\x77\x06\xda\xb5\xba\xd6\x25\x75\xb5\xde\xdb\xac\x88\x4b\xcb\xd5\x9a\x94\x71\xaa\x1f\x97\x86\x05\x8a\x9b\x6d\x9f\xe8\x10\xbb\xbd\x8b\x56\xfb\xa5\x17\x4d\xbf\xdc\x0f\x1e\x37\xbd\x45\x7b\x90\x9c\xe3\xbb\xe7\xe1\x4c\x3e\xf0\x8e\x58\x7e\xf4\xd7\xeb\xd1\x00\x7b\xe1\x70\xe8\x2d\xbe\x89\xd5\x84\x05\x9b\x70\x31\xbb\x1d\xb5\xfd\x60\xf1\xfc\xf0\x29\x5a\xfe\xa3\xdf\x13\x70\x55\xac\x75\xac\x78\xc0\x65\x9f\xa0\x8c\xe5\x36\x8a\x13\x4d\xaa\x76\xdb\xda\x55\x7c\x65\x40\x2b\x77\x5f\x8e\xeb\xc9\x85\xb6\x5c\x11\x27\x9e\x2f\x50\xb1\xb2\x16\x5c\xe0\x86\x0a\x3e\xd7\x74\xf1\x92\x30\xb5\xa5\x1d\xab\x63\x75\xeb\x97\xb2\x96\xc5\x4f\x4a\x19\x5f\x7e\x7c\x77\x7f\x33\xba\x9e\xcc\xcc\x0d\xbd\x9a\x9c\xf7\xf8\xe3\x64\x38\xfa\x1c\xaf\x1f\x07\xfe\x53\xfc\xfe\xf2\xdb\xe4\x8f\xf3\xe5\xa7\x41\x30\xfd\xea\xf1\xab\xed\xf8\xee\xe6\x0b\xbe\x4c\x1f\x3e\xd3\xbf\x66\x77\x8f\xed\xd9\x70\xfe\xfd\x52\x6c\x5a\x95\xe1\xfc\x7c\x64\x14\xb5\x66\x46\xd3\xca\x53\xf6\xbf\x11\x5e\xeb\xb4\x36\xf6\x9a\x28\xa4\xd6\x60\x6c\x5a\x69\xd6\x9e\xc7\xde\xd6\x39\xdb\xc3\x25\xa6\x73\x54\x8d\x84\x6c\x89\x29\xb8\x02\xb5\xee\xd7\x1e\xa8\x1e\x2d\x2e\x53\xa6\x34\x6b\xe6\xe0\xf1\x1d\xce\x8d\xa5\x41\x2e\x99\x6a\xf9\x22\xe1\xde\x09\x44\x4d\x51\x64\xc0\x54\xed\x2f\x31\x78\x84\x98\x2b\x94\xde\xae\x74\xe2\x54\x57\xc4\xa6\x58\x73\x52\x8f\xa7\x07\xf4\xdc\xeb\xef\xf2\xdf\x25\x23\x04\xae\x34\x6b\xd2\x6e\xde\x0f\xa3\x26\xe2\x20\x6c\x03\x94\x98\x1e\x60\xb2\x0c\xd6\xdc\x84\x4d\x13\xdf\xba\x89\x52\x4c\x1a\xf8\xd0\x07\xeb\xae\xfc\xe0\xba\xb6\x94\x7e\x85\x32\x60\x8d\x67\x6c\x58\xa4\x9b\xb9\x1c\x10\x72\x1f\xd8\x4b\x01\x8a\xd8\x9e\xb0\x75\x84\xb3\x05\x6f\x92\x43\xd7\xf0\x94\x11\xc7\xc6\xba\x21\x59\x66\x3d\x7d\x9d\xe4\x39\x71\xb2\xac\xa4\x29\x16\x05\x3a\x36\x15\xdc\x39\x0a\xd6\x82\x62\xc7\x9c\x62\xff\x55\xba\xbd\xd4\xea\x72\x2a\xcb\xeb\x0d\x55\x0d\x29\x11\xff\xa3\xd9\xcd\x51\xf1\x20\x34\xe4\x38\x60\xd1\x2c\xab\x6e\xf1\x93\x66\xea\x64\x25\xaf\x0c\xa5\x91\x47\x01\x84\xac\x60\xec\x93\x6e\x87\xc0\x9a\x7b\x26\xac\x8e\xe5\xc6\xc8\xb2\x43\x52\x6b\x90\xa2\x41\x05\xdd\x4e\x9e\x93\xff\x90\x95\x3f\xad\x5c\xcd\xcc\x8f\x3e\x84\xf6\xe6\x77\x38\xed\xbe\xbc\xc8\x73\xe8\x7c\xd7\xdd\xee\xf4\xf2\x1c\x2e\x36\x04\xe8\x89\x52\x9a\x29\xfd\x46\x1c\x2d\xd0\x5d\xfe\x79\x44\xb2\x1f\xd8\xeb\xbe\xfc\xa2\x18\x68\x8c\x89\x09\x69\x19\x8a\x6a\x83\xca\x10\x67\x12\x07\x30\x96\x3f\xd4\xc4\xc9\xb9\x1f\x5e\xd3\xea\x68\x53\x89\x69\xb1\x75\x76\x42\x6a\x2e\xd8\xdb\x40\xc4\x73\x14\xe5\xc5\xaa\xe4\x54\xef\xa5\x62\xab\x30\x69\x08\x58\x57\xb1\xb7\x1d\xa2\xc1\x72\x8d\x85\x6d\xe7\x96\x6b\xcd\x65\x00\x35\xe2\x8d\x4d\xc3\x76\xb9\xd1\x8e\x15\x69\xd3\x6a\xd9\xd9\xb4\xfa\xdb\xde\x79\xfe\x0d\x00\x00\xff\xff\xc4\xde\x68\x63\xed\x07\x00\x00"

func layoutHtmlTmplBytes() ([]byte, error) {
	return bindataRead(
		_layoutHtmlTmpl,
		"layout.html.tmpl",
	)
}

func layoutHtmlTmpl() (*asset, error) {
	bytes, err := layoutHtmlTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "layout.html.tmpl", size: 2029, mode: os.FileMode(420), modTime: time.Unix(1481178153, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsStylesCss = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x53\x4f\x6f\xdb\x3e\x0c\xbd\xe7\x53\x3c\xb4\xb7\xb4\x4e\x9c\x34\xbf\xfe\x3a\x17\x2b\xb0\x65\x3b\x0c\x18\x76\xd8\x8a\x5d\x86\x1d\x64\x89\x8e\x85\xca\xa2\x2b\xd1\x4d\x83\xa1\xdf\x7d\xf0\x9f\xb6\x46\x96\x0c\x3b\x8a\x7c\x7c\x7c\x7c\xa4\x72\x36\x3b\xfc\x9a\x00\x40\xae\xf4\xdd\x26\x70\xe3\x4d\xa2\xd9\x71\xc8\x70\x5a\x50\xa1\x0b\x7d\x3d\x79\x9a\xcc\xbc\x7a\xc8\x55\x48\xbc\x7a\xb8\x71\xf6\xc6\x56\x9b\x1f\x25\xd9\x4d\x29\x6f\x4f\x2e\x96\x27\x3f\x07\x8a\x4a\x85\x8d\xf5\x49\x68\x13\x19\xfe\xab\x1f\xaf\xbb\xb0\xb1\xb1\x76\x6a\x97\xc1\x7a\x67\x3d\x25\xb9\x63\x7d\xf7\x2f\xac\x67\x50\x03\xf3\x11\x8a\x36\x55\x2b\x63\xac\xdf\x24\x8e\x0a\xc9\x90\x5e\x8f\x95\xf4\xb1\x4e\xc8\xd3\x64\x32\x9f\xe2\x1b\x69\xb1\xec\x33\xac\x95\x73\xdc\x48\x6c\xd1\xef\x99\x25\x4a\x50\x35\x3e\xb0\xee\x22\x6b\xae\x77\xdd\x14\x58\xa6\x8b\x45\xb2\x4c\x17\x97\xb8\xdd\x5a\x11\x0a\xe7\xf8\xe4\xf5\xac\x05\x7d\xb6\x9a\x7c\xa4\x0c\xeb\x35\xde\x89\x04\x9b\x37\x2d\x37\x2e\x66\x29\xa6\xf3\xb6\xdd\x04\xd3\x51\xa3\xee\xf9\x85\x05\xf7\x8d\x15\x82\x72\x14\x24\x9e\x23\x6f\x04\xba\x89\xc2\x15\x94\x37\x28\xc9\xd5\x45\xe3\xe0\x59\x28\xa2\xe0\x80\x82\xdd\x5d\x44\x20\xd5\x8e\x09\x29\x09\x86\x75\x9c\xb5\x6c\x5f\xe9\xbe\xb1\x81\x22\x14\x72\x15\xa9\x23\xa8\xd8\xd8\xc2\x52\x80\x76\x2a\x76\xb0\x4e\x0b\xd6\x5c\x55\xec\x11\x65\xe7\x06\x62\xe5\x1c\x64\x57\x53\x6c\x21\xb3\x3c\x26\xba\xd7\xda\x79\x3e\xd8\x9a\x61\x99\xf6\x7b\xec\x2d\xed\xdf\xbd\xcd\x39\x07\x43\x21\xc3\xa2\x7e\x44\x64\x67\x0d\x4e\x89\xe8\x35\xd3\xd9\x9f\x6c\xad\x91\xf2\xe5\x1a\x86\x4c\x50\xc6\x36\x31\xc3\x45\xbf\x9a\x71\xf3\x72\xd5\xf5\x1f\x36\x28\x5c\x0f\x4b\x1d\x02\x39\x8b\x70\xf5\xb2\xd4\x71\x65\x9d\x39\x15\x25\xd1\xa5\x75\x66\xcc\xf1\x5c\x92\xee\x17\x68\x36\xd4\x01\x0f\xab\x6a\x5d\xbb\x6d\xaf\x80\x3c\x9a\x1a\xb1\x56\x9a\x90\x93\x6c\x89\x3c\xaa\xc6\x89\xad\x1d\x61\x20\xfb\xc3\xc4\x33\xec\x5b\x3a\x1e\x29\x19\x1d\xe5\x77\x15\xac\x6a\x4f\x67\x9f\x23\x31\xca\x6f\x28\x8c\x25\x76\x96\x3e\xff\x4f\x4d\xab\xab\xd5\x6a\x6f\xaa\xe7\xa2\xc1\xc7\xbf\x63\xb7\x2a\xf8\xf6\xaa\x8e\x75\x50\xea\xf2\xff\xf4\xea\x58\xd5\x5e\x8b\xc3\x60\xeb\x0b\x3e\xca\xbf\xc8\xaf\xd2\x37\x74\xb0\x64\x8f\xfc\x15\x39\x9f\xe2\xa3\x37\x2f\x1f\x0b\xd3\xf9\xef\x00\x00\x00\xff\xff\x61\xab\x21\xdd\xc6\x04\x00\x00"

func assetsStylesCssBytes() ([]byte, error) {
	return bindataRead(
		_assetsStylesCss,
		"assets/styles.css",
	)
}

func assetsStylesCss() (*asset, error) {
	bytes, err := assetsStylesCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/styles.css", size: 1222, mode: os.FileMode(420), modTime: time.Unix(1481179878, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesFactoidInfoHtml = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x54\x8f\xcd\x6a\x2b\x31\x0c\x85\xf7\x7e\x0a\x31\xfb\x9b\x21\xdb\x8b\xeb\x4d\xa1\x50\x28\x5d\xf4\xe7\x01\x84\x7d\xd2\x31\x9d\xd8\xc3\xc8\x4d\x09\x46\xef\x5e\xe2\x78\x52\xba\x3b\xfa\xfc\x49\x96\x6a\x0d\x38\xc4\x04\x1a\xa4\x9c\x67\xc8\xa0\x6a\x6c\x8b\xce\x18\x3b\xf6\x54\x2b\x52\x50\x35\xbf\xb6\xcf\xa9\x20\x95\xa6\x87\x78\x22\x3f\xb3\xc8\x5d\xc3\x1c\x13\xd6\xc1\xfd\xe1\x0b\x7f\xe0\xdf\x04\x0e\xed\x85\x88\xc8\x4e\x7b\xf7\xc0\xbe\xe4\x18\xfe\x53\xad\xbb\x9e\x9f\xf9\x08\x55\x3b\x4e\x7b\x67\xec\x18\xe2\xc9\x99\xab\xbe\xb8\xa7\xec\x3f\xd1\xe4\x78\xa0\xdd\xa3\x5c\x6b\xd5\x33\xa4\x56\xcc\x02\xd5\x94\xfb\xa6\x76\x5c\xdc\xad\x8f\xa5\x10\x42\x2c\xed\x9f\x4b\xf5\x2e\x58\x55\x29\xa6\x0d\xdc\x4f\x9c\x12\x66\x55\xe2\xb2\xb1\xb7\x78\x84\x14\x3e\x2e\x7d\x5a\x1f\xb7\x62\xbb\x49\xf2\xd7\xea\x31\x38\xeb\x73\x80\xab\x75\xf7\xc2\xdf\xaf\x8d\x5d\x3a\x1a\xb4\xe3\xb2\xe2\x76\x48\xdf\xed\x27\x00\x00\xff\xff\x0f\x6b\x9d\x95\x73\x01\x00\x00"

func templatesFactoidInfoHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesFactoidInfoHtml,
		"templates/factoid-info.html",
	)
}

func templatesFactoidInfoHtml() (*asset, error) {
	bytes, err := templatesFactoidInfoHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/factoid-info.html", size: 371, mode: os.FileMode(420), modTime: time.Unix(1481185646, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesFactoidListHtml = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x55\x4d\x6f\xe3\x36\x10\xbd\xeb\x57\xcc\xfa\xb0\x68\x01\x4b\xc2\x7e\x9c\x52\xae\xd0\x62\x81\x02\x01\xb2\x3d\x34\x1b\x14\x3d\x35\x63\x69\x64\xb1\xe1\x87\xca\xa1\xec\x08\x02\xff\x7b\x41\x9a\x72\xec\x66\xf7\xd2\x9b\x34\xe4\xcc\x7b\xf3\xe6\x69\xb4\x2c\x1d\xf5\xd2\x10\x6c\xd8\xcf\x8a\x78\x13\x42\x21\xd2\x63\x53\x4c\xaa\xea\xb1\xf5\x56\x76\xa5\x92\xec\x61\x29\x8a\xf0\x2a\xa8\x24\x2c\x05\x00\x40\x7c\x2b\x53\xe6\x0d\x18\x6b\xe8\xa7\x14\xd5\xe8\xf6\xd2\x94\x8a\x7a\x7f\x03\xef\xc7\xe7\xab\xa8\x93\xfb\x61\x0d\x87\xe2\x55\xdd\x4a\x21\xfb\x52\xdb\x4e\xf6\x92\xba\x0c\xd3\x2b\x8b\xfe\x06\x52\x6a\xcc\x12\x75\x66\xbb\x2c\x64\xba\x10\x8a\x97\x8e\x5a\x6b\x3c\x19\x9f\x5a\xea\xe4\x01\x5a\x85\xcc\x9f\x52\x18\xa5\x21\xb7\x69\xae\xe2\x23\xee\xa9\x1c\x08\xbb\x74\x12\xb1\xc4\xf0\xae\xf9\xf5\x44\x8a\x45\x3d\xbc\x6b\x40\xb0\x46\xa5\x1a\x84\x23\xce\xe0\x2d\xb0\xb7\x8e\x00\x4d\x07\x8e\xbc\x93\x74\x20\xf0\xf4\xec\x45\x7d\xba\x57\x88\xba\x93\x87\x6b\x98\x1d\x97\x2d\x2a\x65\x27\x0f\x2f\x8f\xa5\x34\xbd\x3d\xa3\x7e\x6c\x1e\x58\x9a\x3d\x5c\x60\x7f\xcc\x67\x63\x73\x8f\x33\x88\xd6\x76\xd4\xbc\xc9\x82\x19\xd4\x24\xea\x14\x02\x69\x00\xcd\x0c\xed\x80\xc6\x90\x82\xe3\x40\x8e\xe0\x0b\xba\x83\x34\x20\x39\x52\x76\x93\x01\x84\x9c\x5a\xc1\x6d\x0f\x7e\xa0\xf5\x1d\x3c\x3e\x11\x03\xba\xfd\xa4\xc9\x78\xde\x26\xd4\xd1\xd9\x83\xec\x28\x5e\xd4\x80\xbd\x27\x97\x72\x22\xee\x16\x78\xc4\x96\x4a\xa6\x11\x1d\x7a\xea\x2a\xf8\x3a\xd0\x4a\x70\x65\x25\x48\x37\x7a\x62\x2f\x6a\xd2\x0d\xec\x08\xd0\xa7\x0a\x07\x72\x33\xec\x68\x2f\x8d\x89\xfd\xda\xc4\xe5\xe4\x10\x62\xc6\x3d\xc5\x6a\x92\x23\x73\x54\x6c\x21\xaa\x45\x1d\x20\xec\xd0\xec\xcb\xb5\x07\x51\x8f\x67\x75\xfe\xb4\x13\xb4\x68\x4e\xd7\xaf\x5a\x85\xdd\x0c\x8c\x73\xc4\x39\xb1\xfb\x59\x9f\x64\x59\xcf\xf7\xe4\xe1\xb5\xa2\x15\xfc\xc2\xb0\xa3\xde\x3a\xda\x9e\x85\x38\xeb\x93\x60\x91\xc1\x58\xa7\x51\x9d\x98\xfc\xdf\x99\xff\xe1\xa4\xff\xee\xd4\x1f\x38\xc9\xff\x1f\xea\x8e\x34\xe9\x1d\xb9\x4b\xde\x90\x6d\xbf\x6a\xdf\x5a\xad\xa3\x41\xa3\x5b\xc9\x5f\x8c\xfe\x52\xb6\xdb\x1e\x66\x3b\x81\x21\xea\xae\xec\x90\x6b\xc5\xe4\xfc\xd9\x80\xa1\xa3\x92\x86\x78\x0b\x9d\x05\x3f\x48\xbe\x39\x75\x3d\x3a\x6a\xbe\xc3\x8e\x9e\x51\x8f\x8a\xe0\xed\xbb\xe2\xf1\xf1\xb1\xf8\x32\x29\x2f\x63\x8d\xe2\xf3\xa9\x7c\x8c\x66\xba\xa2\x8e\x75\xbe\x45\x4b\xc7\xac\x58\x45\xaf\xe9\x30\xa2\xf3\xbc\x85\xbf\x27\xf6\x80\x5d\xf4\xc5\xdb\xf7\xd9\x31\x47\xeb\x9e\x18\x7a\xeb\x00\x95\x5a\x35\xe0\xf8\x79\xe4\x8f\xa1\x84\xdd\xe4\xc1\x58\x1f\x2f\x25\xbc\x4b\x4f\xf1\x36\x4e\xd5\x0f\x34\x03\xaa\x23\xce\x0c\x53\x1e\x00\x19\x2f\x1d\xad\xfe\x7c\x73\x39\xf1\x42\x0c\x1f\xd6\x7d\x01\x77\x32\xda\x7d\xf8\xd0\x14\x62\x52\xab\x0f\x2e\x37\xdc\x26\xae\x2c\x70\x68\xf6\x04\x55\xbc\x0d\x21\x14\x00\x42\xc9\x46\x20\x0c\x8e\xfa\x4f\x9b\x7a\xa5\x53\x2f\x8b\xec\x81\xfe\x81\xea\xbe\xb5\x23\x7d\xce\x9f\xf7\x66\x13\xc2\x5f\xcb\x42\x8a\x29\x84\x65\xb9\x3a\x8c\x81\xb4\x11\xeb\x65\xa9\x32\xab\xdf\x50\x53\x08\xd9\x73\x90\xcd\xb4\x2c\x70\x79\x0e\x21\xe4\x59\x24\x4c\x43\xdf\xc0\x14\x3c\xa2\x59\x9b\x92\x5c\xe6\x75\x53\x5a\xa3\xe6\x4d\xf3\x8a\x88\xa8\xe3\xfd\x26\xf3\x11\x35\xae\x04\x12\x42\x75\xcb\x77\xb6\x7d\xa2\x2e\x84\x1f\x54\x7a\xf8\x71\xdd\xe5\x99\xe6\x25\xda\xd5\x3f\x61\xd3\xc4\x57\x38\xff\x22\x76\x33\x2c\x4b\x75\x87\xec\x1f\x98\x5c\x08\x71\xcf\xe4\xc0\x57\xa9\x89\x3d\xea\xf1\x4c\x67\xad\x3e\x3a\x5a\x8b\xb3\x9d\x5c\x4b\x9b\x66\x15\xa6\xfa\x1d\x8f\xf7\x29\x76\x16\xe5\x6c\x50\x51\x2b\x99\x46\x48\xa6\x8b\xa3\x13\xf5\xf4\xb2\xf0\x73\x03\xff\x06\x00\x00\xff\xff\x39\x1c\xf1\xad\x5e\x07\x00\x00"

func templatesFactoidListHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesFactoidListHtml,
		"templates/factoid-list.html",
	)
}

func templatesFactoidListHtml() (*asset, error) {
	bytes, err := templatesFactoidListHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/factoid-list.html", size: 1886, mode: os.FileMode(420), modTime: time.Unix(1481185626, 0)}
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
	"layout.html.tmpl": layoutHtmlTmpl,
	"assets/styles.css": assetsStylesCss,
	"templates/factoid-info.html": templatesFactoidInfoHtml,
	"templates/factoid-list.html": templatesFactoidListHtml,
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
	"assets": &bintree{nil, map[string]*bintree{
		"styles.css": &bintree{assetsStylesCss, map[string]*bintree{}},
	}},
	"layout.html.tmpl": &bintree{layoutHtmlTmpl, map[string]*bintree{}},
	"templates": &bintree{nil, map[string]*bintree{
		"factoid-info.html": &bintree{templatesFactoidInfoHtml, map[string]*bintree{}},
		"factoid-list.html": &bintree{templatesFactoidListHtml, map[string]*bintree{}},
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
