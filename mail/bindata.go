// Code generated by go-bindata.
// sources:
// attachment.go
// attachment_test.go
// mail.go
// mail_test.go
// responsive.html
// DO NOT EDIT!

package mail

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

var _attachmentGo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x51\xcd\x6a\xe3\x30\x10\x3e\x4b\x4f\x31\xeb\x93\x0d\xc6\xbe\x2f\xe4\xb0\x64\xff\x58\x76\x97\x40\x0f\x3d\x84\x1c\x14\x67\x6c\x0f\xb1\x24\x57\x19\x97\x86\x92\x77\xef\x4c\xed\xb4\xf4\xd0\x53\x7b\x11\x9f\x34\xfa\xfe\xa4\xd1\x35\x47\xd7\x21\x78\x47\x83\xb5\xe4\xc7\x98\x18\x72\x6b\x32\x8a\x35\xc5\x89\x69\xc8\x64\x13\x4f\xba\x8e\x8e\xfb\xba\xa5\x01\x15\x64\x56\x4e\x3a\xe2\x7e\xda\x57\x4d\xf4\xf5\xc3\x74\x47\xa1\x6b\x31\x74\xf5\x40\x2d\x67\xef\x4e\xd5\xc9\xbb\x50\xcf\xda\x85\xb5\xed\x14\x1a\xb8\x71\xf7\xf8\x8d\xd9\x35\xbd\xc7\xc0\xb9\xba\xac\x63\x60\xc1\xb0\xdd\xed\xcf\x8c\x25\x70\x3c\x62\x28\x41\x47\xff\x9d\x47\x38\x71\x12\xcd\x02\x30\xa5\x98\xe0\x51\xf2\xf4\xd1\xe3\x77\x4a\xf0\x75\x05\xaa\x5e\xfd\x42\xfe\x3d\x1f\xe5\x85\x35\x07\x4a\x1b\x09\xae\xd3\x6b\x89\xea\x4f\xa4\x90\x2f\xb4\x72\x26\x89\x6d\x4b\x9d\xde\xdc\x66\xec\x47\x05\xd9\x6e\x71\x17\x95\xba\x86\x7f\xf1\x80\x1b\x4c\xde\x1a\xb1\x56\x39\x2d\x5c\xad\x13\x3a\x56\x1d\x6c\x38\xa6\x73\xbe\xd8\x95\x10\x4f\xd5\x95\x21\x7c\x6a\x35\x30\x7c\x59\x41\xa0\x41\x52\x1b\xf3\xec\xfa\x53\x12\xfd\x8d\x5d\xf5\x43\xcb\xe4\x72\x63\x41\x85\x50\x4c\x42\x9e\x52\x50\x9e\x35\x17\x6b\xdc\xcb\x3b\x7d\xb4\xcf\xeb\x6b\x16\x73\x99\x15\xcc\x9f\x5e\xdd\x26\x62\xd4\x50\xf9\x5b\xb7\x99\xb1\x7c\xcd\xe7\x77\xb3\xd7\xbd\x08\xd8\x8b\x7d\x0a\x00\x00\xff\xff\x9d\x31\x2d\xea\x9e\x02\x00\x00")

func attachmentGoBytes() ([]byte, error) {
	return bindataRead(
		_attachmentGo,
		"attachment.go",
	)
}

func attachmentGo() (*asset, error) {
	bytes, err := attachmentGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "attachment.go", size: 670, mode: os.FileMode(420), modTime: time.Unix(1464849676, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _attachment_testGo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xa4\x90\x41\x4b\xc4\x30\x10\x85\xcf\xcd\xaf\x18\x03\x42\x2a\xa5\xf5\xac\xf4\x20\x2a\x8a\x07\x11\x76\x6f\xcb\x22\xd9\x92\xb4\x61\x9b\xa4\xb6\xd3\x45\x91\xfd\xef\xce\xb4\xbb\xa2\x0b\x9e\x84\x52\x26\x93\xf9\xde\x7b\x99\x4e\x57\x5b\x5d\x1b\xf0\xda\xb5\x42\x38\xdf\xc5\x1e\x41\x89\x44\xc6\x41\xd2\xbf\xd3\xd8\x14\xd6\xb5\x86\x0b\x6e\xa0\x19\xd0\x85\x5a\x0a\xaa\x6b\x87\xcd\xb8\xc9\xab\xe8\x8b\xf7\xf1\x8d\xba\xd6\x84\xba\x60\x25\xaf\x43\x31\xa2\x6b\xa5\x48\x85\xb0\x63\xa8\x60\x49\xdc\x42\xef\xcc\x0d\xa2\xae\x1a\x6f\x02\x2a\x84\x8b\x83\x5a\xbe\x4c\xe1\x93\x14\xd9\xe8\x36\x06\xa4\x5b\xb8\x2a\x61\xb5\xde\x7c\xa0\x51\x12\x1b\x37\x00\x7d\xd5\x7c\x25\x53\x91\x60\xdc\x9a\xc0\x33\xf2\x52\xce\xdc\xb3\xf6\x66\x6a\xb0\x26\xc7\x33\x7d\xcf\xe7\x13\xd3\x1f\x16\x19\x4c\x2a\x19\x1c\x71\xd2\x75\x16\x98\x3b\x2b\x21\xb8\x96\x32\x25\x09\xe6\xf7\x7d\x1f\x7b\xab\xe4\x89\x52\x0a\x96\x5e\x0a\xe7\x3b\x99\x31\x43\xf0\x9e\x5c\x9b\xe8\xcd\x9d\x9b\x9c\x79\x01\xf9\x83\xc1\xc7\xb9\xa5\x68\x42\x7f\xe3\x2f\xb4\x4f\x1e\x3a\xee\x36\x7f\x8a\x2e\xa8\x03\x9d\xcd\x2c\xe5\xb4\xae\xe6\xc9\x95\x44\xdf\x71\x21\xd7\x7f\xa4\x7e\x9d\x42\x40\x09\x71\xc8\x17\xa8\x51\xfd\xb6\x4a\xaf\xff\xf5\xae\xbd\xf8\x0a\x00\x00\xff\xff\x66\x4e\xe2\x59\x28\x02\x00\x00")

func attachment_testGoBytes() ([]byte, error) {
	return bindataRead(
		_attachment_testGo,
		"attachment_test.go",
	)
}

func attachment_testGo() (*asset, error) {
	bytes, err := attachment_testGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "attachment_test.go", size: 552, mode: os.FileMode(420), modTime: time.Unix(1464849719, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _mailGo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x56\x51\x6f\xe3\x36\x0c\x7e\xb6\x7f\x85\xe6\x87\xc1\xc6\x7a\xf2\x75\x03\x86\xa1\x43\x07\xf4\x72\xed\x76\xc3\x72\x08\xd6\x00\x3b\xa0\x28\x76\x8e\x2c\xa7\xba\x5a\x92\x27\xc9\x69\x83\xa2\xff\x7d\xa4\x2c\xc7\x76\x12\x2c\xb7\x3c\xb4\x12\xf5\x91\x22\xf9\x91\x94\x9b\x82\x3d\x16\x6b\x4e\x64\x21\xea\x38\x16\xb2\xd1\xc6\x91\x34\x8e\x92\xd5\xd6\x71\x9b\xc0\x42\xe8\x5c\xe8\xd6\x89\x1a\x37\x4d\xe1\x1e\xf2\x4a\xd4\x1c\x17\x28\xb0\xce\x30\xad\x36\xb8\x74\xfc\xd9\xe5\x8e\xcb\xa6\x2e\x1c\x4f\x62\x90\xac\x85\x7b\x68\x57\x94\x69\x99\x9b\xd6\x5a\xa3\xad\xcd\x57\x35\xdc\x58\x19\x51\x16\xdb\x64\x0a\x79\x6e\xff\x11\x6a\x5d\x71\xb5\xce\xd1\x1b\x59\xa8\xbc\x60\x4c\xb7\xca\x9d\x06\x5a\xe9\x9a\xd3\xa8\x3e\x8a\xb5\x6e\x1e\xd7\x54\xa8\x7c\xad\xf1\x88\x6e\xbe\x4f\xe2\x2c\x8e\xdd\xb6\xe1\x64\x0e\x02\x02\x51\xb5\xcc\x91\x97\x38\xba\x6d\x57\x5f\x38\x2c\xc7\x3f\x38\x05\xcb\xb0\xf8\xfc\xc5\x6a\x75\x91\xd8\x0e\x93\x7c\x8e\xa3\xa5\x26\x07\xbf\xbb\xfb\xa0\x10\xe0\x4e\x23\x72\xc6\x4e\x23\x19\x43\xe4\x8d\xd1\x72\x1f\xb9\xef\x42\x05\x18\xc4\x2e\x8c\xd0\x46\xb8\xed\x18\xbb\xd2\xba\xf6\x8b\x80\x6d\x02\x06\xf1\xef\x74\xb9\x3d\x65\x7b\x05\x98\x2e\xb6\x47\xae\x4e\x60\x1d\x62\x10\x7c\xe5\x5c\xc1\x1e\x24\x57\xee\x06\xaa\xe5\x63\x21\xb9\x3d\x08\xaf\x38\xc4\x80\xea\x6b\x1c\x57\xad\x62\xe4\x96\xab\x12\xc9\x48\x91\x22\x4f\x4b\x46\xb8\x31\xda\x00\x2d\x71\x14\x2a\xe3\x0c\x45\xe4\xe2\x92\x84\x3d\xfd\x95\xbb\xab\x6e\xf9\x41\x55\xda\xeb\x52\x4c\x60\x16\x47\xa2\xf2\xe0\x6f\x2e\x89\x02\x83\xc0\x6d\x64\xb8\x6b\x8d\x42\x69\x1c\xbd\x02\xd7\xf3\xe5\xe2\x96\x9b\x0d\x37\x3b\xb3\x58\x57\x68\x73\x38\xfa\x7f\x26\x25\x1a\x09\x55\xf6\x91\x3f\xcd\xb9\xb5\xd0\x6d\x29\xa8\x4a\x7a\xcb\xdd\x6f\xbc\x28\xc1\x64\x12\xaa\x2c\x39\xf3\x7d\x48\xc3\x76\x1f\xb5\xd4\x3d\x60\xa9\x29\xa5\x50\xb1\x51\x9e\x13\x28\x24\x61\x09\xb4\x9d\xdb\x7a\x87\x6a\xae\x3a\x1f\x67\x2c\x23\xbf\x90\xb7\xde\x2d\xc0\xc9\xb6\x76\xa2\xa9\x39\x61\x0c\x04\x13\xc3\x33\xd6\x1b\x9e\x31\x6f\x18\x5d\x07\x15\xa7\x4b\x4d\x9e\xa0\xa5\x88\x02\x72\xc8\x1b\xe0\xc4\x5d\x95\xa5\x81\x28\x3a\x55\x44\x79\x4b\x13\xe9\xd4\xe0\xdd\xdb\x7b\x7a\x8d\xeb\x89\x04\xd9\xde\x8f\x0f\x93\x0a\x7a\x3d\x95\x5e\x29\x04\xd9\xd7\xac\x8f\xd0\x9b\xd9\x55\xfa\xcb\x7e\x34\x9f\xde\xf4\x67\x60\x2c\x39\x4f\x7c\x38\xde\xca\x50\x70\x71\xe4\x64\xf3\x5e\x78\x8e\x71\x2a\x20\xc7\x4b\x2f\x41\x6e\x06\x5c\x80\xf4\x13\x8f\xfe\xae\x85\x4a\x3b\xd5\x1d\x15\x50\xf1\xd9\x34\xf3\x47\x8a\x7f\xa0\xa2\x82\x0a\xfe\xfb\x8c\x6c\xd0\xae\x29\x54\x98\xbd\xc7\x74\x3c\x7c\xe4\xcc\x02\x3c\x38\xf4\x66\xe2\x2c\xd8\xcd\x50\x49\x06\x7b\xe9\x54\x1b\xcf\x5e\xbb\x74\xc0\xd4\x76\x20\x45\x7b\x8b\xc2\x58\x8e\x0d\x36\xeb\x64\x5d\x10\x38\x1a\x30\xfb\x3e\xb7\xb8\x49\xbb\x09\xff\xe0\x64\x0d\x89\x0d\xfa\x88\xc0\x37\x63\xe8\x98\xee\x41\x80\xfb\xb5\x48\x87\xb6\xa1\x0b\x00\x7d\x4d\xc7\x94\xd3\x8e\x79\x2f\x8a\x1a\x48\x1d\x19\xea\x3b\xb4\xbb\x75\x52\x2c\xc3\x76\x51\x58\xfb\xa4\x4d\x39\xdc\x78\x49\x4a\x8a\xc6\xae\x54\x89\x93\x25\x95\xd9\xcf\xfb\x9e\xf8\x42\xc0\xec\xff\xa1\xd7\xf4\x1a\x47\x4d\x0a\x88\xb0\xca\xb2\x43\x5f\xc3\x16\xf4\x71\x6a\x41\x81\xdd\x7c\xf8\x34\xbf\xbe\x20\xe7\x3f\xe6\x3f\xe5\x3f\x9c\x77\x73\xec\x20\xbb\x38\x4f\xc3\xe0\xcc\x48\xda\xf3\xd0\x0b\x5e\xba\x52\x95\x85\x79\x2c\xf5\x93\x22\x0d\xaa\x43\x8e\xf1\x5f\x39\x1b\x48\x1b\xbd\xa3\x74\x1e\xc0\x33\x2d\xa5\x56\xe9\xdd\x3d\x3e\xde\xfe\x1e\xf4\xba\xbf\xe1\x32\xdc\x91\x4e\x6c\x01\x81\xc0\x31\xb8\xb7\x0c\x2f\x77\x90\xef\x08\xed\x1e\x7f\xfa\x27\x74\x17\xe6\x26\xf5\xbb\xf9\x48\xc1\x97\xe5\x77\x24\xc9\x13\xf8\x7b\x70\xba\xc4\x47\x15\x4e\xa9\xaf\x1b\xb8\x0d\x9f\x8d\xff\xba\xef\xca\x5a\xee\xd2\x24\x39\x5a\x2c\x13\x8a\xfe\x2a\x8c\x3a\xce\x10\xb0\x43\x78\x6d\xb9\xd7\xf1\xcf\xba\x1c\x08\x18\xbd\xee\x51\x34\x9b\xa4\xbf\x6b\x8f\xc8\x35\xf5\xce\x9d\xfe\x83\x06\x8b\x31\x4d\xd0\x4c\x92\x51\xcf\x69\x1a\xd2\x79\x24\x18\xef\xca\xa1\xf7\xa7\xdd\xdf\xf9\x3f\x0e\x20\xda\x14\x86\xac\xda\x8a\xf8\x6f\x32\xfa\xae\xad\x2a\x1c\xbd\xd0\xe5\x0c\x3d\x1c\x85\xe6\xd1\x3d\xe1\x67\xb8\xc1\x70\xa2\xae\xfe\x21\x2a\x7a\xfd\xcc\x59\x0b\xb5\xf1\x2d\x98\x83\x01\xc6\xfc\x9d\x47\x1c\xfd\x0a\x4f\x07\x57\x27\xbe\x8e\xca\x0d\xee\xa0\xb7\x5d\x8e\xb2\xde\x97\x30\x7d\x82\xee\x6b\xfc\x6f\x00\x00\x00\xff\xff\xe7\xf6\x00\xe0\x80\x0a\x00\x00")

func mailGoBytes() ([]byte, error) {
	return bindataRead(
		_mailGo,
		"mail.go",
	)
}

func mailGo() (*asset, error) {
	bytes, err := mailGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "mail.go", size: 2688, mode: os.FileMode(420), modTime: time.Unix(1472690805, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _mail_testGo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\x93\x41\x8b\xdb\x30\x10\x85\xcf\xd1\xaf\x98\xea\x64\x95\xad\xbc\xbb\x24\xa1\x14\x02\x5d\xda\x1c\x03\x4b\x93\x5b\xe9\x41\x6b\x8f\xb3\xa6\x96\xe4\x4a\xe3\x6d\x4a\xc8\x7f\xaf\x24\xbb\x89\x83\x03\xbd\xc4\x19\x69\xe6\x7b\x6f\x9e\x71\xab\x8a\x9f\x6a\x8f\xa0\x55\xdd\x30\x56\xeb\xd6\x3a\x82\x8c\xcd\xb8\xf5\x3c\xfc\x12\x7a\xaa\xcd\x9e\xb3\xf0\x7f\x5f\xd3\x6b\xf7\x22\x0b\xab\xf3\x43\xf7\x2b\x9c\x56\x68\xf6\x79\x1c\xd4\xca\xe4\xaa\x28\x6c\x67\x88\xff\xb7\xd1\x6b\x6a\x39\x13\x8c\x55\x9d\x29\x60\x17\x04\x36\xaa\x36\x99\x86\xf7\x83\x98\xdc\x08\x38\x06\xc1\x81\x38\x8f\x83\xf0\x69\x05\x43\x2d\x9f\xfa\xe7\x91\xcd\x66\xbc\x72\x56\x7f\xc6\x83\xd2\x6d\x83\x51\x90\xdf\xc5\xd3\x56\x79\xff\xdb\xba\x32\x56\x27\x36\x8b\x8a\x67\x4a\x2c\xe4\x76\xb3\x7b\xde\xa2\x7b\x43\x97\x28\x53\x42\xea\x9a\x9c\x3e\x2e\x06\xe2\x3f\x2b\x5b\xf5\x86\x83\x9d\x6c\x6c\x57\xf4\xa2\xe9\xfe\xa2\x95\x9d\x8d\x84\x7b\x2d\xbf\x75\x26\x13\x17\xd6\x57\x6c\x90\x6e\xd2\xe4\x7a\xcc\xec\xfb\x6e\x51\xe5\x53\x59\x3a\xf4\x3e\x74\x5a\x2f\xd7\x87\x9a\xb2\x7b\xc1\x4e\xa3\xa4\xb7\x68\xca\x90\x76\x93\xd1\x25\xed\x5d\x9f\xf6\xb0\xec\x66\x88\x29\x3e\x53\x36\xc3\x79\x0a\xe0\xfb\x0f\x4f\x2e\xcc\x1c\x39\xd9\xab\xcc\x4e\x57\xb7\xa9\xba\xfd\x6a\x2a\xd5\x78\xbc\x1b\x71\xe1\xc5\x96\x7f\xfa\x74\x1f\xe6\xcb\x87\xe5\x62\x71\xff\x71\xf9\x38\x5f\xf2\x09\x31\xc4\x8e\xce\x45\x6f\xe7\x2d\x46\x9e\xc3\xce\x75\x05\xb1\x61\xb5\x02\x13\x76\x88\xe6\x49\xae\x9d\xb3\x2e\xe3\xe7\x09\x01\xfe\xd5\x76\x4d\x09\x55\xa8\xb8\x88\xd4\x71\x3e\xcf\xca\xf9\x84\xfb\x62\x0d\x61\x78\x0d\xd3\x9c\xf2\x3c\x7d\x2b\x1f\x08\x83\xb4\x22\x04\x63\x09\xf0\x50\x7b\x4a\x0e\x38\x87\x77\x2b\x98\x70\x38\x17\xd7\x8e\x26\x1d\xe2\xca\xd2\xdf\x00\x00\x00\xff\xff\xde\xbd\x4e\x0f\x97\x03\x00\x00")

func mail_testGoBytes() ([]byte, error) {
	return bindataRead(
		_mail_testGo,
		"mail_test.go",
	)
}

func mail_testGo() (*asset, error) {
	bytes, err := mail_testGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "mail_test.go", size: 919, mode: os.FileMode(420), modTime: time.Unix(1470023560, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _responsiveHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x59\xdd\x73\xe3\x34\x10\x7f\xcf\x5f\xa1\x33\xc3\x14\x6e\xea\xd8\x49\x3f\xd2\xe6\xa3\x0c\x1c\x77\x73\xcc\x50\xe0\xa1\x30\xc3\xa3\x6c\xc9\xb1\x5a\xcb\xf2\xc9\x72\xd2\x52\xfa\xbf\xb3\x92\xed\xd4\x72\xdc\x94\x32\x85\x98\x99\x53\x1f\x1a\x6b\xe5\xd5\xee\x6f\x3f\xb3\x99\xbf\x71\x5d\xf4\x43\x9a\xb0\x94\x4a\xf4\x5d\xc1\x12\x82\x7e\xa3\x32\x67\x22\x45\xc7\x47\x67\x7e\x30\x99\x1c\x8f\x82\x60\x72\x72\x4e\x4e\xc3\xe0\xfc\x7c\x72\x72\x7c\x12\x1d\x61\x42\xc6\x23\x4c\x8e\xcf\x22\x7f\xe4\x07\xc8\x75\x2f\x06\xf3\x37\xdf\xff\xfc\xee\xea\xf7\x5f\xde\xa3\x58\xf1\x04\x9e\xf5\x3f\x94\xab\xbb\x84\x2e\x9c\x48\xa4\xca\x8d\x30\x67\xc9\xdd\x14\x1d\x7c\xa4\xc9\x8a\x2a\x16\x62\xf4\x13\x2d\xe8\xc1\x61\x63\x07\x1e\x36\x9f\x0f\xd1\xb7\x92\xe1\xe4\x10\xe5\x38\xcd\xdd\x9c\x4a\x16\xcd\x90\xe1\x94\xb3\x3f\xe8\x14\x8d\x7c\xff\xcb\x19\xd2\x82\xbb\x31\x65\xcb\x58\xc1\xd6\xf0\x94\xf2\x19\xe2\x58\x2e\x59\x3a\x45\xfe\x0c\x65\x20\x2a\x4b\x97\xfa\xb3\xa3\x85\xa2\x98\x5c\x0c\x10\xac\x39\xa7\x0a\xa3\x30\xc6\x32\xa7\x6a\xe1\x14\x2a\x72\xcf\x9c\x26\x29\xc5\x1c\x24\x5f\x31\xba\xce\x84\x54\x0e\x0a\xe1\x66\x9a\xc2\xd1\x35\x23\x2a\x5e\x10\xba\x62\x21\x75\xcd\x83\xf5\x5e\xac\x54\xe6\xd2\x4f\x05\x5b\x2d\x9c\x77\xe5\x3b\xee\xd5\x5d\x46\x1b\x1c\x14\xbd\x55\x9e\xc6\x67\xb6\x11\xe0\xd7\xab\x0f\x8f\xf7\x1b\xd4\xca\xcf\x7a\x79\x6f\x97\x4c\xc5\x45\xf0\xd6\x1b\x34\xf6\xc2\x22\x57\x82\xc3\x5e\xbd\x45\xd8\x6a\x58\x5d\x81\xee\x37\xbb\x7a\xbd\x1a\xfc\xdb\x5c\x1b\xa6\xb0\x88\x1d\x66\xb1\xe8\x84\xe5\x59\x82\x41\x9c\x20\x11\xe1\x8d\x4d\xe3\xf8\xb6\xc4\x75\x8a\xce\xdb\x7c\x37\xa6\x45\xb8\x50\xc2\xa6\x35\x6c\xbd\xd9\x7f\x18\x58\xf8\x04\x82\xdc\xf5\x1f\x1c\xf3\xf2\xba\xa2\xa7\x42\x72\x9c\x3c\x85\x82\x0f\xec\xb3\xdb\x17\xe1\x10\x14\x4a\x41\x74\xf7\x0b\x85\xb1\x4d\x0b\x45\x22\xe4\x14\x7d\x11\x99\x65\xd3\x02\x21\x09\x95\xae\xc4\x84\x15\x39\xbc\x79\xd2\xd6\x7f\xe3\x5b\xcc\xa4\x35\xb7\xc3\xc5\xc2\x42\xe6\xfa\x82\x4c\x30\x08\x18\xb9\x03\xfe\x40\x24\xc4\x26\xeb\xe8\x75\x09\x0d\x85\xc4\x0a\xd2\xa4\xb6\x50\x4a\x5b\x22\xe2\xf0\x66\x29\x45\x91\x12\xd0\xe1\xe8\xf8\x8c\x12\xfc\x84\x01\x9f\xb7\x5c\x43\xe5\x1a\x95\x2e\x8e\xd5\x11\x93\x3a\xa6\x28\x17\x09\x23\x9d\x07\xaa\xc0\xd2\x5e\x83\xc6\x96\xeb\x34\x5c\x04\xf7\xcc\x3b\x3a\x62\x64\x17\x16\x2f\x40\xb7\xa1\x34\xe3\xcb\x21\xe5\xe2\x9a\xb5\x94\xdf\x08\xd1\x16\xa1\x86\xb2\xbd\xff\x18\x9d\x43\xff\x84\x72\xfd\x7f\xeb\xcc\x8a\x4a\x8d\x51\xe2\xe2\x84\x2d\xe1\xac\xeb\xdb\x67\x1e\xca\x4a\xe0\x35\x4a\xc1\x5c\x31\x05\x9f\x2f\x31\x4b\x38\x4e\xd1\x9f\xe8\xe3\xd5\xe5\x8f\xe8\x3d\x87\x67\x74\x45\x39\xb8\xbc\xa2\x73\xaf\x3c\x54\xd5\x91\x50\xb2\x4c\xa1\x5c\x86\x0b\xc7\xf3\xd4\xda\x68\x37\x84\xf4\x1a\x92\x14\x4a\x05\xf7\xc6\x8f\x9b\x2c\x1d\x5e\xe7\xdf\x8c\x87\xa3\xe1\xc8\xb9\x80\x7b\xcd\xab\x50\x31\xbd\xb2\x64\xce\x4d\xe6\x0c\x96\x06\xf6\x85\xa3\x03\xd3\xa9\x84\xdd\x4f\x8d\x87\x00\x0d\x6e\x18\xdc\x68\x0e\x73\x21\x54\x6c\x6c\x8b\x53\x05\x0c\x19\xce\x29\x99\x3d\x9a\xce\xb0\xa9\x5f\x31\x01\xac\xf9\xbb\x98\x5c\x43\x15\xad\x22\x78\x63\x4f\x38\x8c\xde\x30\xae\xcb\x3e\x70\xdb\xd1\x4e\x0c\x4c\xf3\x64\x90\x31\x2d\x90\xc2\x41\x42\x51\x98\xe0\x3c\x5f\x38\x7a\xdb\x5d\x4b\x9c\x39\x0d\xd8\x3e\x98\x55\x23\xb7\x27\xe8\x1a\x7a\x76\x2b\x67\xf2\x42\xdd\x8c\x28\xd9\xb7\x2e\xae\x8e\x91\xb9\x22\x3d\x13\x0d\xc2\x8f\xd8\xf2\x55\xce\xa0\xfb\x32\xac\x5b\xec\x1d\xce\xb0\x47\x8f\x08\x13\x8a\xa5\xae\x75\x2a\xb6\x3c\xdf\xee\xd2\x2c\xd2\xa5\xdd\x88\xb5\x9c\xa7\xaa\x36\x70\x01\x94\x19\x53\x8c\xa0\x96\xfb\xfa\xcf\xc4\x4d\x53\x63\x13\x43\x75\xdf\xaa\xc3\xc8\x22\x42\xd3\xd6\xc4\x10\xce\x38\xf6\x89\x12\x68\x13\x78\x3d\x8d\x26\xcb\x67\x6d\xb1\x7b\x1b\x5a\xdb\xa2\xf6\x2d\xd4\x9e\x14\xd5\x88\x8b\x75\x75\xb5\x73\xb1\xe6\x74\x43\xc4\x3a\x75\xf5\xd3\x33\x0c\xf4\xba\xbf\x1f\x56\xdf\xdf\x1e\x1e\x76\xdf\xe6\x55\xd7\xed\xc0\xcf\x4a\x0c\x2d\x8a\xec\xf0\x68\xcf\xb8\x74\x2b\x18\x3c\x88\x86\x8b\xed\xe0\xf1\x9a\xd1\x33\xe8\xbe\xb2\x7f\x06\x7c\x94\xaf\x84\x60\xa3\x72\xa9\x53\x5d\x55\xcd\x53\x04\xe5\x9d\xca\x8e\x22\x5b\x12\xca\x32\xbb\xdf\xa2\xfa\x54\x0a\xfd\xfb\xe9\xa1\xc7\xe9\xa0\xcf\xde\x53\xcb\xb7\x55\x69\xff\x9f\x75\xb5\x31\xfc\x38\xf5\xf5\x97\xb4\xce\x4e\xb4\x5d\x74\x5f\xa1\xac\x6e\x67\xa7\xfd\x60\xd6\x9a\x0c\x6d\x01\xf2\x1c\x0a\x5b\xa9\xf4\x73\x73\xf0\x1f\x34\x07\xe6\x6b\x2c\x38\x13\xd5\xe3\x94\x0e\x5f\xda\xbb\x5b\xfd\xb3\x3e\x22\xfb\x77\xa5\x1d\x6b\x77\xee\x4c\x1c\xd5\x74\xe3\xd4\xac\x59\xe7\x44\x70\x7b\x0a\xf8\x02\xcd\xf4\x5a\x31\x0c\x9d\x12\x8a\x25\x8d\x16\x8e\x1e\x5b\xe7\x53\xcf\x2b\x47\xcd\x66\x3a\x70\x5b\x7c\x02\x66\x11\x4d\x97\x1e\x2f\x87\x0e\xbb\x0d\xdb\x0b\x3b\xd7\xc8\x9d\x9b\xf5\xb4\xdd\x2b\x8d\x9e\x6b\xec\x9e\xf1\x0f\x2f\xfb\xdc\xf3\xbd\x5e\xcf\xd7\x6c\xf3\xaa\xd9\x95\xba\xcb\x68\xf5\x83\xc9\x35\x5e\xe1\x72\xb7\xf2\xed\x7a\x76\xa5\xa5\x41\x0b\x74\x30\x19\xdf\x4e\xc6\x07\x33\x8b\x96\xe9\x9f\x57\xbe\x22\x22\x2c\x38\xa0\x65\xe6\xff\x5f\xcf\x06\xcd\xf9\x96\xde\x32\x73\x2e\xf3\xb3\xd5\x5f\x01\x00\x00\xff\xff\xad\xc1\x3a\x74\x0f\x1b\x00\x00")

func responsiveHtmlBytes() ([]byte, error) {
	return bindataRead(
		_responsiveHtml,
		"responsive.html",
	)
}

func responsiveHtml() (*asset, error) {
	bytes, err := responsiveHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "responsive.html", size: 6927, mode: os.FileMode(420), modTime: time.Unix(1470022268, 0)}
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
	"attachment.go":      attachmentGo,
	"attachment_test.go": attachment_testGo,
	"mail.go":            mailGo,
	"mail_test.go":       mail_testGo,
	"responsive.html":    responsiveHtml,
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
	"attachment.go":      &bintree{attachmentGo, map[string]*bintree{}},
	"attachment_test.go": &bintree{attachment_testGo, map[string]*bintree{}},
	"mail.go":            &bintree{mailGo, map[string]*bintree{}},
	"mail_test.go":       &bintree{mail_testGo, map[string]*bintree{}},
	"responsive.html":    &bintree{responsiveHtml, map[string]*bintree{}},
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