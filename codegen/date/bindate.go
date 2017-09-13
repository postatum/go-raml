// Code generated by go-bindata.
// sources:
// codegen/date/date_only.go
// codegen/date/datetime.go
// codegen/date/datetime_only.go
// codegen/date/datetime_rfc2616.go
// codegen/date/time_only.go
// DO NOT EDIT!

package date

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

var _date_onlyGo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\x92\xd1\x6b\xdb\x30\x10\xc6\x9f\x7d\x7f\xc5\xcd\x4f\x52\x1a\xa7\x59\x03\x83\x0d\xfc\x30\x36\xf2\x30\xd6\x75\xb4\xd9\xd3\x18\x54\x89\xcf\xad\xa8\x25\x99\x93\x5c\xd0\x46\xfe\xf7\x21\x39\x89\x5b\x5a\xa8\x5f\x0c\x77\xdf\x77\xf7\xbb\xcf\xee\xd5\xee\x41\xdd\x11\x36\x2a\x10\x80\x36\xbd\xe3\x80\x02\x8a\x32\x68\x43\x25\x48\x80\x47\xc5\xa9\x90\x04\x57\xb6\x8b\x6b\x13\x70\x7c\x6a\x2c\x2f\x96\xcb\x0f\xd5\xf2\x7d\xb5\xbc\x28\x9f\x29\x36\x7a\xf7\x40\x0d\xd6\x78\x5b\xde\xe2\x19\x3e\xf5\x9e\xa5\x5a\x9a\x7b\x7e\x8e\x5f\x0f\x75\x64\xea\x99\x3c\xd9\x80\xd7\x9f\x2f\xbf\x67\x7d\xe5\x52\x23\xc4\x9e\x92\x72\x73\x4f\x58\xb6\x43\xd7\x55\xa9\x57\xa2\x75\x41\x05\xed\x2c\xba\x16\xaf\xd7\x5f\x56\xab\xd5\xc7\x39\x5a\x65\xa8\x8b\x18\x63\x8c\x95\x31\x55\xd3\x2c\xf2\x12\x47\x3e\xe9\xd1\x0f\x7d\xbe\x2e\x5d\x86\x8e\xc7\xf7\x5f\x67\xa9\x72\x6d\xeb\x29\x9c\x86\x2e\x20\xad\x9d\xe8\x92\x70\xb1\xd1\x86\x32\xf4\xa5\x62\x7f\xaf\xba\x6f\x37\x57\x3f\xd0\x3d\x12\xb3\x6e\x08\xcd\x54\x84\x76\xb0\x3b\x14\x8d\xc3\xd9\x71\x82\x7c\x6a\x12\x12\xc5\xef\x3f\xdb\x18\x68\x8e\xc4\xec\x58\xe2\x3f\x28\x98\xc2\xc0\x16\xc7\x86\x38\x6d\x14\xb3\xc6\xc9\xc5\xda\xb1\x51\x41\xbc\x08\x58\xca\x39\x5a\xdd\xc1\x3e\x83\xfd\xb2\xe6\x35\xb4\xc1\xbe\x05\xf7\xcc\x28\xb6\x07\x08\x39\xd2\x25\xb8\xe0\x33\x2a\x7e\xaa\xc7\x2c\x7e\x2a\xf6\xf4\x12\x67\x8e\x3e\xb0\xb6\x77\x62\x2b\x25\x14\xba\xcd\x9e\x77\x75\x42\x4c\x53\x8e\x37\x12\x33\x14\x7b\x80\x62\xd6\x38\xac\x4f\x31\x8b\xe0\xe5\x29\x87\xe9\xaa\x9b\x3c\x12\xc7\xba\x3f\x6c\x98\x7e\x98\xfc\xc1\x5e\xbd\x6a\x34\x0a\x79\xb4\x4c\x21\xbf\x99\xae\x84\x3d\xfc\x0f\x00\x00\xff\xff\x10\x9a\x87\x5d\x19\x03\x00\x00")

func date_onlyGoBytes() ([]byte, error) {
	return bindataRead(
		_date_onlyGo,
		"date_only.go",
	)
}

func date_onlyGo() (*asset, error) {
	bytes, err := date_onlyGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "date_only.go", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _datetimeGo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\x90\x4f\x6b\x1b\x31\x10\xc5\xcf\x3b\x9f\x62\xba\x97\x4a\x8e\xed\x6c\xe2\xa6\x10\xc3\x9e\x5a\x7c\x28\xf4\x0f\x8d\x7b\x69\x29\x44\xb6\x66\x53\x91\x48\xbb\x8c\xc6\x81\x50\xfc\xdd\x83\x76\xbd\xde\x18\x1b\xac\x93\x78\xd2\x7b\xf3\x9b\xd7\x98\xf5\xa3\x79\x20\xb4\x46\x08\xc0\xf9\xa6\x66\x41\x05\x59\x2e\xce\x53\x0e\x1a\xe0\xd9\x70\x12\xd2\x87\xa5\xf3\xb4\xf0\x82\xdd\x29\x31\xbf\x2e\x8a\x8f\x93\xe2\x6a\x52\x5c\x2f\xaf\x6e\xe6\xc5\x87\x79\x71\x33\xbd\xed\xcf\xef\xfc\xc0\xb5\x74\xeb\x47\xb2\x58\xe2\x7d\x7e\x8f\x17\xf8\x36\xef\x22\x69\x69\xd6\xe5\x25\x7e\xde\xe9\xe8\x22\x26\x86\x28\xc6\x37\xe8\x02\xe6\xc9\x31\x69\xb1\xb0\xaa\xd9\x1b\x41\x4b\x95\x0b\x64\xd3\xeb\xcf\xc5\xa7\xd9\x6c\x76\x0b\xf2\xd2\xd0\x90\x91\x7e\x4f\xd3\xad\x8d\xfe\x6a\x38\xfe\x33\x4f\x5f\xee\xbe\x7f\xc3\xfa\x99\x98\x9d\x25\xf4\x83\x08\xd5\x26\xac\x51\x59\xc1\x51\x9f\xa0\xdf\x9a\x94\x46\xf5\xe7\xef\xea\x45\x68\x8c\xc4\x5c\xb3\xc6\xff\x90\x31\xc9\x86\x03\x76\x0f\x6a\x3f\x51\x8d\xac\xe8\xe9\xa2\x05\x55\x47\x35\x68\x3d\xc6\xe0\x9e\x60\xdb\x82\xfd\x0a\xfe\x14\xda\x26\x9c\x83\x3b\x30\xaa\xd5\x0e\x42\x77\x74\x09\x4e\x62\x8b\x8a\xf3\xb2\xeb\xe2\x87\xe1\x48\xc7\x38\x63\x8c\xc2\x2e\x3c\xa8\x95\xd6\x90\xb9\xaa\xf5\xbc\x2b\x13\x62\x4a\xe9\x77\x24\x66\xc8\xb6\x00\xd9\xc8\x0a\x96\xfb\x9a\x95\x44\xbd\xef\x61\xd8\xea\xae\x8d\xc4\x4e\x8f\xe8\xe4\x7d\xdc\x8d\x41\xa6\x86\x29\x52\x10\x23\xae\x0e\x27\x57\xeb\xdc\x4a\xf7\x96\xa1\xe9\xb3\x15\x6b\xd8\xc2\x6b\x00\x00\x00\xff\xff\xb5\x10\x2c\xd8\xd8\x02\x00\x00")

func datetimeGoBytes() ([]byte, error) {
	return bindataRead(
		_datetimeGo,
		"datetime.go",
	)
}

func datetimeGo() (*asset, error) {
	bytes, err := datetimeGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "datetime.go", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _datetime_onlyGo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x92\x41\x4f\xdc\x30\x10\x85\xcf\xf1\xaf\x98\xe6\x64\xc3\x26\x04\x54\x2a\x11\x29\x87\xaa\x88\x43\x55\x4a\x55\xd2\x13\x42\xc2\xbb\x99\xb0\x16\x6b\x3b\x1a\xcf\x52\xa5\xd5\xfe\xf7\xca\x09\x6c\x76\xdb\xad\x44\x4e\xd1\x8c\xdf\xf3\xf7\x5e\xd2\xe9\xc5\x93\x7e\x44\x68\x34\xa3\x10\xc6\x76\x9e\x18\xa4\x48\x52\x36\x16\x53\xa1\x84\x78\xd6\x14\x07\xf1\x40\x9c\xdd\xb8\x55\x7f\x65\x19\xc6\xa7\x82\xf4\xac\x28\x3e\x64\xc5\x69\x56\x9c\xd5\xa7\xe7\x65\xf1\xbe\x2c\xce\xf3\x8b\x8b\xf4\x1f\x45\x6d\x16\x4f\xd8\x40\x05\x0f\xe9\x03\x1c\xc3\xdf\x7e\xc7\x71\x1e\xef\x3b\x39\x81\xcb\x9d\x1d\x10\x76\x84\x01\x1d\xc3\xf7\x8f\xd7\x5f\xb6\xba\xcc\xc7\x25\xf7\x1d\x46\xc5\x27\x6f\xe7\xc6\x61\x33\xac\xc7\x95\x76\x0d\x4c\x07\x7f\x1a\x5e\x82\x86\x80\x9d\x26\xcd\x9e\xc0\xb7\x90\xd6\xe9\x2c\x8a\x9d\xb6\xb8\xea\xa1\xef\xfb\x3e\xb3\x36\x6b\x9a\x7a\xb9\x2c\xad\x2d\x43\xb8\xcb\xdb\x36\xcf\xf3\xfb\x1c\x2e\x3d\x06\x70\x9e\x21\xac\xbb\xa1\x23\x3d\x98\xc3\x2f\xef\x10\x7c\xdb\x06\xe4\x5c\x44\x9a\x7d\xf8\xf8\x92\xd7\xc6\xe2\x90\xeb\x5a\x53\x58\xea\xd5\xe7\xdb\x9b\xaf\xe0\x9f\x91\xc8\x34\x08\x76\x1a\x8a\x76\xed\x16\x20\x1b\xf6\x70\xb4\x6b\xa3\x76\x95\x52\x81\xbc\xbb\x9f\xf7\x8c\x33\x40\x22\x4f\x0a\x7e\x8b\x84\x90\xd7\xe4\x60\x5c\xc8\xed\xb5\xf2\xa8\x61\xaf\xf2\x2b\x4f\x56\xb3\x3c\xf8\x45\x94\x9a\x81\x33\x2b\xb1\x19\x10\x7f\x38\x7b\x08\x72\xed\xde\x84\xb9\xa7\x96\xf3\x17\x1c\x35\x72\x46\x4c\x0e\x03\x34\x94\xd5\x58\xcd\x37\x4d\x01\x0f\x73\xcd\x20\x30\x19\xf7\x28\xe7\x4a\x89\xc4\xb4\x83\xee\x5d\x15\x59\xa3\xd3\x6b\x62\x24\x12\xc9\x46\x88\x24\x26\x85\x6a\xaf\x7e\xc9\x41\x6d\xab\x99\x32\xde\x0e\xbe\x30\xce\xc3\xcb\x35\xd3\x7f\xa6\xd9\x78\xf7\xff\x8c\xa3\x5a\xaa\x57\xdd\x54\xfe\xdb\x5a\x57\x62\x23\xfe\x04\x00\x00\xff\xff\xb1\x36\x35\x97\x76\x03\x00\x00")

func datetime_onlyGoBytes() ([]byte, error) {
	return bindataRead(
		_datetime_onlyGo,
		"datetime_only.go",
	)
}

func datetime_onlyGo() (*asset, error) {
	bytes, err := datetime_onlyGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "datetime_only.go", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _datetime_rfc2616Go = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x91\x41\x6b\xdb\x40\x10\x85\xcf\x9a\x5f\xf1\xaa\x4b\x77\x1d\x91\x28\xa6\xf1\xc1\xa0\x53\x8b\x0f\x01\xb7\xa5\x76\x4f\xa5\x90\xb5\x35\x4e\x97\x44\x2b\x31\x3b\x0e\x84\xe2\xff\x5e\x56\x8a\xe3\x26\x75\x8a\xf7\xf8\x66\xdf\x9b\x6f\x66\x3a\xb7\xbe\x73\xb7\x8c\xda\x29\x13\xf9\xa6\x6b\x45\x61\x28\xcb\xd5\x37\x9c\x93\x25\x7a\x70\x92\x84\xf4\x61\xe9\x1b\xfe\x36\xfb\x38\x9e\x5c\x4e\x66\x8d\x62\x78\x15\xf2\x79\x1b\x0a\x94\x63\x5c\xbb\x80\x71\x59\x4e\x70\x79\x35\x2d\x3f\x4c\xcb\x2b\xcc\x17\xcb\xfc\x98\x77\xe9\xd7\x77\x5c\xa3\xc2\x4d\x7e\x83\x33\x1c\x09\x3f\x4b\xa5\xd4\xff\xe2\x02\x9f\x5e\x96\xe1\x23\x12\x5e\x54\xd7\x74\xf0\x01\x7b\x7d\xd3\x4a\xe3\x94\xf4\xb1\xe3\x7f\x3c\xc9\x70\x9e\x84\x3e\x71\xee\x24\xfe\x72\xf7\xd7\x8b\x2f\x9f\xd1\x3e\xb0\x88\xaf\x19\xcd\x41\xa4\xcd\x36\xac\x61\x6a\xc5\xe8\x55\x90\xfd\xdb\x6b\x2c\xcc\x8f\x9f\xab\x47\xe5\x02\x2c\xd2\x8a\xc5\x6f\xca\x84\x75\x2b\x01\x43\xc1\x3c\x37\x36\xa3\x5a\xed\xf9\xac\x67\x34\x6f\xad\xc4\xda\x02\xc1\xdf\xd3\xae\xc7\xfc\x1e\x9a\x63\xa0\xdb\x70\x22\xea\x0b\xbf\x59\x3d\x21\xd9\x81\x35\xa1\x6a\xec\xc1\x31\xad\x86\x05\x7d\x75\x12\xf9\x4d\xb8\x02\x51\xc5\x87\x5b\xb3\xb2\x96\x32\xbf\xe9\xad\xef\xaa\x04\x9c\xc2\xf6\x83\xb3\x08\x65\x3b\xa2\x6c\x54\x2b\xaa\xd7\x97\x30\x1a\xed\xf3\x8e\x0e\xa3\x2e\xfa\x64\x0c\x7a\x84\xd7\xf7\xf1\xa9\x1b\x84\x3b\xe1\xc8\x41\x9d\xfa\x36\xfc\x6f\xde\x21\xc4\xd8\xbd\xf3\x70\x8c\x53\xaf\x60\x69\x47\x7f\x02\x00\x00\xff\xff\xe4\xf1\xdc\x20\x14\x03\x00\x00")

func datetime_rfc2616GoBytes() ([]byte, error) {
	return bindataRead(
		_datetime_rfc2616Go,
		"datetime_rfc2616.go",
	)
}

func datetime_rfc2616Go() (*asset, error) {
	bytes, err := datetime_rfc2616GoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "datetime_rfc2616.go", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _time_onlyGo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\x92\x41\x4f\x1b\x31\x10\x85\xcf\xeb\x5f\x31\xdd\x93\x1d\x12\x43\x95\x72\xc8\x4a\x7b\xa8\x5a\xe5\x50\x95\x52\x01\x3d\x21\x24\x9c\x64\x96\x58\xac\xed\xd5\x78\x82\x94\x56\xf9\xef\x95\xbd\x49\x16\x04\x52\xf6\xe6\xf1\xbc\x37\xdf\xbc\x75\x67\x96\xcf\xe6\x09\x61\x65\x18\x85\xb0\xae\x0b\xc4\x20\x45\x51\xb2\x75\x58\x0a\x25\xc4\x8b\xa1\x54\x48\xe7\x6b\xdf\x6e\xe7\x8e\xa1\xff\x6a\x28\x3f\x5f\x56\x17\x5f\xaa\x8b\x4b\x3d\x9b\x95\x6f\x5a\xee\xec\xf2\x19\x57\x50\xc3\x63\xf9\x08\x67\xf0\x5a\x7c\x96\x6a\xc9\xf8\xfc\x1c\xee\xf6\x75\x20\xec\x08\x23\x7a\x86\x9b\xaf\x57\x3f\x73\xff\x24\xa4\x0b\xde\x76\xa8\x73\xeb\x1a\xa1\xec\x0c\xb1\x35\xed\x24\xc3\x81\x0f\x6c\xd8\x06\x0f\xa1\x81\x9b\xf9\xb7\xe9\x74\x3a\x1b\x83\x37\x0e\xdb\x2d\xac\xd7\x95\x73\x55\x8c\xf7\xba\x69\xb4\xd6\x0f\xd9\xe3\x7b\xc0\x98\x54\x10\x37\x5d\x5e\x34\x6d\x0d\x81\xf2\x3c\xf8\x1b\x3c\x4e\x42\xd3\x44\xe4\xa3\xb5\x16\x09\x60\xe0\x4c\x8d\x3a\x9d\x32\xfe\x95\xa1\xb8\x36\xed\x8f\xdb\xeb\x5f\x10\x5e\x90\xc8\xae\x10\xdc\x50\x14\xcd\xc6\x2f\x41\x72\x80\xd1\xc1\x41\xbd\x16\x49\x05\xf2\xfe\x61\xb1\x65\x1c\x03\x12\x05\x52\xf0\x4f\x14\x84\xbc\x21\x0f\xfd\x85\x3c\x4e\x94\x23\x0e\x4a\xcf\x03\x39\xc3\xf2\x5d\xd4\x4a\x8d\xc1\xdb\x56\xec\x32\xd8\x1f\xef\x3e\x42\xdb\xf8\x53\x70\x6f\x84\x72\xb1\x87\x50\x3d\x5d\x82\xe3\x98\x51\xa1\xaa\xfb\x2c\x7e\x1b\x8a\xf8\x1e\x67\x0c\x91\xc9\xfa\x27\xb9\x50\x4a\x14\xb6\xc9\x9a\x4f\x75\x42\x4c\x2e\x87\x1d\x91\x48\x14\x3b\x21\x8a\x11\x07\xa8\x8f\x31\x4b\x8e\xea\x98\xc3\xb0\xd5\x6d\xb6\x84\xbe\x1e\xf7\x13\x86\xa7\x93\x7f\xd8\x87\x5b\xf5\x42\xa9\x0e\x92\x21\xe4\x93\xe9\x2a\xb1\x13\xff\x03\x00\x00\xff\xff\x87\x02\xb2\xbe\x24\x03\x00\x00")

func time_onlyGoBytes() ([]byte, error) {
	return bindataRead(
		_time_onlyGo,
		"time_only.go",
	)
}

func time_onlyGo() (*asset, error) {
	bytes, err := time_onlyGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "time_only.go", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"date_only.go":        date_onlyGo,
	"datetime.go":         datetimeGo,
	"datetime_only.go":    datetime_onlyGo,
	"datetime_rfc2616.go": datetime_rfc2616Go,
	"time_only.go":        time_onlyGo,
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
	"date_only.go":        &bintree{date_onlyGo, map[string]*bintree{}},
	"datetime.go":         &bintree{datetimeGo, map[string]*bintree{}},
	"datetime_only.go":    &bintree{datetime_onlyGo, map[string]*bintree{}},
	"datetime_rfc2616.go": &bintree{datetime_rfc2616Go, map[string]*bintree{}},
	"time_only.go":        &bintree{time_onlyGo, map[string]*bintree{}},
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
