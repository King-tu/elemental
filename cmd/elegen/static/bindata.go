// Code generated by go-bindata.
// sources:
// templates/README.md
// templates/identities_registry.gotpl
// templates/model.gotpl
// templates/relationships_registry.gotpl
// DO NOT EDIT!

package static

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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz) // #nosec
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
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

var _templatesReadmeMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x04\xc0\xcb\x11\x02\x21\x10\x45\xd1\x7d\x47\xf1\x2c\x53\x22\x81\xc6\xbe\x0a\xc5\x47\x6a\x60\x33\xd9\xcf\x79\x2b\x31\x56\xf7\xc3\x36\x4b\x85\x8d\xbe\xb5\xa3\x09\xa1\xf3\x57\x46\x8c\x4c\x04\xa1\x3a\x75\x0a\xca\x75\xfa\x75\xbf\xcc\x24\x69\x78\x43\xcb\x3f\xcd\x7f\xd8\x13\x00\x00\xff\xff\xaa\x97\xff\x85\x4d\x00\x00\x00")

func templatesReadmeMdBytes() ([]byte, error) {
	return bindataRead(
		_templatesReadmeMd,
		"templates/README.md",
	)
}

func templatesReadmeMd() (*asset, error) {
	bytes, err := templatesReadmeMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/README.md", size: 77, mode: os.FileMode(420), modTime: time.Unix(1563338401, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesIdentities_registryGotpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x58\x4d\x6b\xe4\x38\x13\xbe\xfb\x57\xd4\x1b\xc2\x8b\x0d\x1d\xe7\xb2\xec\x21\x4b\x0e\x61\x98\x81\x86\xcd\x10\x12\xd8\x4b\x93\x83\xc6\x5d\x76\xc4\xca\x92\x91\xe4\xcc\x34\x46\xff\x7d\x91\x2c\x7f\xc8\xed\xfe\x4a\x32\x3d\x30\xb9\xa4\xdb\xaa\xa7\xaa\xf4\x94\xea\x51\xb9\x2b\x92\xfd\x4b\x0a\x84\xa6\x81\xf4\x09\x75\xfa\x49\xf0\x9c\x16\xb5\x24\x9a\x0a\x9e\x7e\x25\x25\x82\x31\x51\x44\xcb\x4a\x48\x0d\x17\x85\x48\x49\x25\x24\x6a\x91\x52\x71\x8d\x0c\x4b\xe4\x9a\xb0\x8b\x28\x7a\x25\x12\xe2\x08\x00\x80\xae\x91\x6b\xaa\x37\x16\xac\xee\x49\x05\xb7\x50\x92\x6a\xa5\xb4\xa4\xbc\x78\xee\x31\xe9\xd2\xdb\x41\xe3\x60\xf6\xaf\x69\xae\x40\x12\x5e\x60\x9b\xcc\x53\x85\x19\xcd\x69\xe6\x92\x51\x36\x91\xc1\x10\x68\x0e\xea\x45\xd4\x6c\xfd\x88\x05\x55\x1a\x65\x60\x0d\x29\x5c\xa6\x0f\xf5\x37\x46\xb3\x7b\xb1\xc6\x10\x7b\x05\x97\x43\x8a\x70\x73\x0b\xa9\xb5\x61\xe9\xe7\xe1\xe1\xd5\x08\x70\xd1\x34\xde\xe0\x11\x95\xb6\xcb\xc6\x5c\xdc\xd8\x1c\xc6\x6e\x8c\xe9\x36\xb4\x08\x42\x21\x5f\x4f\xa3\x8f\x1e\x99\x28\xe0\x2c\x23\x1a\x0b\x21\xe9\x6f\x48\x9c\xa8\x65\x86\x3f\x85\x3c\xc2\x28\x51\x3f\x91\xb1\xab\x33\x52\xd6\x34\x3e\xab\x4b\xba\x80\x4b\xb7\xb3\x11\xe8\xae\xdd\x29\x84\x1c\x77\x76\x1f\x47\xec\x9e\x83\xca\xd7\xf8\x63\x86\xeb\xd5\xf3\xea\xb9\xfd\x78\x76\x8e\x2d\x03\x93\xfe\xec\xa9\xa0\x39\x30\xe4\x90\x2e\xdb\xb4\xc1\x98\x51\x7a\x61\x8a\x8e\xf0\x4c\x94\x95\xa8\xf9\xda\x71\x3e\x80\x42\xc8\x18\xd4\x03\x8c\x71\x79\xd8\xff\x8b\x81\x3a\x30\x7b\x69\x37\x8b\xa6\x01\x64\xca\x26\xcc\x29\x5b\xbc\xb1\x32\x49\x14\x5d\x5f\x83\xa3\xe0\x1f\x94\xca\xd2\x25\x51\xd7\x92\x2b\xd0\x2f\x08\x59\x2d\x25\x72\x0d\xaf\x7e\x4d\xe4\xee\x71\xe9\x28\x8b\xf2\x9a\x67\x01\x36\x4e\x20\x67\x82\xe8\x3f\xff\x80\xc6\xfb\xe9\xaf\x87\xbb\x87\xe5\x92\xe7\x22\xed\xc2\xd8\x1d\x46\x91\xde\x54\xde\xdd\x3d\xe1\xa4\x40\x09\x4a\xcb\x3a\xd3\x8d\x89\x9c\xfb\x38\x0f\x56\x13\xe8\xce\xe4\x17\x29\x4a\x5b\xaf\x98\xdb\xa2\xb5\xe7\x27\x81\xd9\xbe\x75\x5b\xf5\xd9\x4c\x6f\x98\x95\x85\x3f\x47\xc7\x44\xfb\xd4\x0a\xec\x26\xf6\x4a\xbb\x39\x3d\x6a\xa0\xd1\xab\xce\xcf\x71\xe1\x5d\x03\xc7\x6d\xbb\x1e\x1d\x78\xd0\xb7\x95\xfb\x78\x64\x28\xbe\x89\x09\x1f\xf6\x17\xd3\x99\x48\x49\x17\x8a\xe6\x40\xe1\x16\xf2\x74\xab\x34\x84\x6f\x92\xbf\xe0\x7f\x34\x5d\xaa\xcf\x65\xa5\x37\x71\x32\x6a\xa1\x8e\x9a\x40\x22\xe6\x5c\xf5\xbc\x9f\xec\xce\x3f\x0b\xdd\x79\x1e\xf9\x26\x39\xc0\x45\x4e\xc9\x37\x86\x71\x57\xbb\x59\x0a\xa6\xcf\x5a\x4c\xc7\x8c\xfa\x4e\x75\xf6\xd2\x57\xdf\x67\xdb\xeb\xf4\x1e\x65\x7b\xb3\xaa\x65\x44\xb5\x13\xd9\xd6\x55\x31\xc8\xf9\xcd\x94\xb4\xaf\xf8\x7d\x07\x24\x4e\xa2\x19\xd9\x98\x7c\x5d\x63\x4e\x6a\xa6\xb7\xdc\x72\xca\x7c\x35\x76\x11\xfd\x54\x11\xa9\xf0\x2d\x74\x6f\x23\x7f\x21\xe9\x1e\xc8\x85\xee\x48\x5c\xaa\x47\x21\xf4\x7b\x8b\xd2\x6e\xf2\x3d\xa5\xf9\xb0\x4a\xf9\x0b\x6d\x7f\x79\x82\x9b\x3c\xd0\xbf\xfe\xea\x5f\x75\x0e\xdc\x0b\xc2\x21\x39\x6a\x2b\x6b\xbb\xf6\xc9\xb9\x0d\x54\x69\x7f\xef\x4d\x7a\xdf\x9f\xae\x89\x14\xb4\x3a\x97\x1c\xa7\x04\x07\x36\x3f\x9f\x8e\xfa\xb0\x63\xb9\xf3\x74\x9d\x57\x2b\xfe\x3f\x07\x78\x60\xb5\x24\x0c\x8c\xf9\x9b\x2a\x7b\x75\x9f\xf1\x60\x6e\x0b\xc1\xd1\x75\x9a\x81\xfe\x76\xd5\xda\x2d\x21\xbf\xb0\x66\x01\xe5\x27\x75\xb7\xda\xdb\xde\xea\xe4\xfe\x7e\x44\xd6\x96\xef\x85\x56\x2a\x1e\x47\x0d\x56\xda\x3a\xc9\xe9\x74\x25\xe7\x6c\x6c\xac\x57\x22\xa1\xf4\xf3\xec\x6d\x10\xd2\xce\xb5\x76\xe0\xf6\x8b\xe3\x59\xdb\x99\x8d\x32\xb8\x1f\xc1\xba\x59\xbb\xfd\x16\x24\x3a\x36\x1b\xa6\xee\x2e\xba\x89\x5c\xbc\x3b\xc6\x3c\x31\x14\x55\x1f\x95\x30\x06\xf8\x83\x2a\x6d\x15\x9b\xf6\xeb\x3e\x58\x80\x89\xad\xba\x1f\x1c\x37\xe7\x4c\xce\xff\xde\xbc\xb7\x45\x4e\xfb\x91\xc0\x78\xf6\xdc\x14\xfd\x45\xc8\x7e\xdf\x63\x0a\x6d\xf1\xfc\xa0\x0d\xb9\x90\xee\x7b\x41\x5f\x71\x98\xfb\x7b\x46\xa7\x7e\x0e\xdd\xa7\xe1\x6d\xba\x4b\x95\x8e\xa0\xf5\xbc\xb2\xd3\x25\xde\xbc\xf7\x47\x89\xfd\xaf\xc0\x47\x88\x54\xf8\x26\x60\xa5\xc9\x44\xff\x05\x00\x00\xff\xff\xb4\x15\x94\xdc\xb9\x14\x00\x00")

func templatesIdentities_registryGotplBytes() ([]byte, error) {
	return bindataRead(
		_templatesIdentities_registryGotpl,
		"templates/identities_registry.gotpl",
	)
}

func templatesIdentities_registryGotpl() (*asset, error) {
	bytes, err := templatesIdentities_registryGotplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/identities_registry.gotpl", size: 5305, mode: os.FileMode(420), modTime: time.Unix(1585855989, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesModelGotpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x3d\x69\x73\xdc\x36\xb2\x9f\x3d\xbf\x02\x99\x52\x6c\x8e\x77\x44\xe5\xf3\x64\xb5\x55\x8e\x8f\x3c\xd5\xf3\xa1\xb2\xbc\xd9\xaa\xe7\x72\xad\x20\x12\x1c\x21\xe6\x00\x0c\x00\x4a\xd6\x9b\xe2\x7f\x7f\x85\x83\x24\xc0\x13\x1c\xcd\x38\xca\x5b\xf9\x43\x22\xe1\x68\xf4\x85\x46\x77\x03\x6c\x65\x30\xfa\x0a\xd7\x08\x6c\xb7\x20\xbc\x40\x22\x7c\x49\x49\x82\xd7\x39\x83\x02\x53\x12\xbe\x87\x1b\x04\x8a\x62\x36\xc3\x9b\x8c\x32\x01\x82\x19\x00\x00\xcc\x93\x8d\x98\xeb\x9f\xd6\x34\x84\x19\x65\x48\xd0\x10\xd3\x13\x94\xa2\x0d\x22\x02\xa6\x65\x2f\x16\xd7\xf9\x55\x18\xd1\xcd\xc9\x3a\xa5\x57\x30\xe5\x78\x4d\x4e\x36\x6b\x7a\x72\xc5\x29\x69\x0f\xda\x60\x11\x5d\xa3\x34\xbd\x3e\x89\x68\x76\xc7\x05\xcb\x23\x91\x33\xa4\x07\x6e\xb7\xc7\x80\x41\xb2\x46\x20\xbc\xc8\x50\x14\x7e\xba\xcb\xd0\x39\xa3\x37\x38\x46\x8c\x4b\x24\x15\x34\x49\x07\x28\x8a\x7a\x0a\x22\x31\x38\x36\xbd\x4d\x10\xbf\xc1\x14\xc7\x8a\x52\x4f\x40\x45\x31\x5b\xcc\x66\xdb\x2d\x38\x4a\xa1\x40\x5c\xfc\x86\x18\xc7\x94\x80\xd5\xa9\x81\xf8\x56\x35\xbf\x10\x82\xe1\xab\x5c\x20\x5e\x0e\x90\x3c\xdc\x6e\xcd\xe2\x47\x78\x09\x8e\x10\xc9\x37\x72\xde\x55\x8e\xd3\xf8\x35\xc9\x37\x5c\x83\x68\x82\x2e\x8a\xd9\xc9\x89\x14\x8f\x9a\xa1\xa8\x06\x45\x01\x18\xca\x18\xe2\x88\x08\x0e\xc4\x35\x02\x19\xe5\x1c\x5f\xa5\x08\xdc\xc0\x34\x47\x1c\x24\x94\x01\x58\x62\xa1\x88\xd1\xd3\x2b\xcc\x8c\x64\xe7\xe1\x4c\x48\x88\x2d\xf8\x5c\x30\x4c\xd6\xb3\x59\x44\x09\x2f\xe5\x5e\xb3\xef\x88\xc0\x0d\x5a\x82\x23\xb5\x9a\xa4\x42\x4f\xfe\x4d\x2f\x6e\x58\x68\xd0\x26\x7a\xa5\x26\xc6\x7a\xaa\x1c\xa0\x7f\x2a\x8a\xd0\x2c\x52\x4f\x69\x61\x75\xaa\x49\x29\x67\x94\xc2\xa9\x65\x53\xff\x3c\x93\xd8\xe2\x04\x10\x2a\x8c\x6c\xde\xd1\x18\xa5\xe1\x2b\x24\x60\x74\x8d\xe2\x9a\xb1\x76\xef\x6b\x22\xb0\xb8\x33\xcc\x39\x8b\x91\xfa\xb5\x89\x7a\xd5\x4e\x13\xf5\x3b\xbd\xfa\x1d\x45\x22\x9c\xdd\x40\xe6\x07\xef\x14\x54\x3b\x25\xac\x1a\xb7\x8a\x18\x39\x74\x05\x2a\x0d\xb4\x40\x7d\x44\x5c\xc8\xde\xa2\x98\x2f\xd5\xd0\x97\x50\xa0\x35\x65\x77\xab\xae\xa1\x34\x67\x51\x25\x64\x3d\xfe\x5c\x6f\xf5\x55\x1b\xb4\xe9\xa9\x47\x32\x7c\x03\x85\x1c\xd9\x1c\xa8\x3b\x8a\x62\x39\x2b\x26\xf2\x7a\xbb\xed\x1a\x71\xc6\x3f\x52\x2a\xc6\x64\x71\x9e\xe6\x0c\xa6\xa0\x28\xde\x62\x2e\x6c\x69\x40\x90\xca\x16\x9a\x78\xcc\xad\x14\xdd\x67\x8d\xcf\x5f\x9e\xf7\x8e\x94\x04\x9f\x9c\x00\x4b\x3b\x44\xce\x88\x56\x0d\xdc\xa9\x1a\x1c\x60\xa2\x7e\x95\xd8\x86\xb3\x24\x27\x11\x08\xa8\x27\x2e\x8b\x6a\xa5\x60\xd1\xad\x37\x4a\x66\x1a\x8b\x7e\x98\xb5\xfa\xcd\x34\xfe\x2f\x69\x56\xe3\x0e\x41\x46\x31\x11\x88\x01\x41\x01\x04\xd2\xfc\x2a\x84\xfd\x50\x9c\x4e\x92\x5c\xbc\x83\x9c\x04\xc3\xab\x14\xf1\x92\x26\x85\xc6\xea\x14\xc0\x2c\x43\x24\x0e\xfc\x80\x6f\x8b\x25\xa0\x61\x18\x2e\x6c\xb6\x3c\x95\xa0\x0c\xe1\x2f\x14\x34\x03\x94\x3b\x62\x12\x54\xfd\x0a\x01\x41\xb7\x7a\x75\x23\xc7\x43\xf1\x41\xe3\x12\x94\xeb\x87\x61\xd8\xcd\x92\x51\x56\xd1\x5c\xdc\x93\x53\xf2\xc8\xf8\xf7\x52\xb2\x42\x02\xd2\x76\xbe\xc4\x4b\xdb\xa6\x72\x9d\x6a\x19\x9a\x0b\x35\x21\x0c\x86\x76\xcb\x42\xc3\x2f\x1c\x3d\xa5\xb9\x30\xe2\x50\xfb\x2d\xa2\xe4\x06\x31\x61\x4b\x43\x69\x22\xe9\xa3\x7b\x37\x76\xcb\xff\xf6\xab\x9d\xc2\xc4\xe5\xe7\x06\x7e\x45\xc1\xc0\xf0\x25\x48\x11\x09\xe8\xa2\x66\x21\x96\xd3\x7e\xfa\x19\x60\xf0\x77\xd3\xf7\x33\xc0\x7f\xfb\x9b\xcb\xc2\xcf\xf8\x0b\x38\x05\xf4\x33\xfe\x32\xc8\x9a\x57\x28\x81\x79\x2a\x3e\xb0\x18\x31\xc7\xcc\xc4\xba\x03\x50\xd9\x83\xc9\x1a\x24\x18\xa5\x31\x2f\xb5\x35\xa2\x44\x20\xb2\x03\x7f\xec\x05\x83\x05\xf8\xfc\x45\xbb\x01\x0d\x1b\x53\x36\xd7\x24\x55\xae\x8d\x5e\xc5\xc1\xbb\x74\xbe\x1c\xaf\x6a\x69\x4f\x35\xa7\x88\xe6\x84\xa6\xfc\x13\xbd\xc8\x20\xe3\xc8\xa1\xda\xd3\x76\x1b\x5d\x42\xb1\xd4\x20\x0d\xc6\x77\xfb\x9e\x9c\x80\x0f\x6d\x8b\x0d\x6e\x71\x9a\x02\x4a\xd2\x3b\xc5\x59\x68\xba\xd6\xf8\x06\x11\xc3\xf9\x10\xbc\xa7\xfa\x47\xb0\x41\x90\x70\x20\xf5\x84\x21\xd3\xc4\xd1\x0e\xb2\x28\x59\x10\x18\xd9\x86\x61\xa8\xd9\xee\x6b\x0b\x94\xee\x4e\xa1\xff\xde\xca\x1c\x36\x70\x96\xb6\x25\x0c\x9e\x8f\xe0\x00\x8a\x62\xd8\x42\x94\xae\xb0\xad\x0b\x37\xa6\xed\xbe\x1a\x6f\x60\x07\x0b\x80\x89\xe8\x3a\x4b\x91\x08\x5f\x9c\x9f\x9d\x91\x84\x86\x96\x4b\xae\xdd\x79\xa3\xb8\x56\x74\xa0\x7e\x3e\xc2\xfc\x35\x89\xd8\x5d\x26\xa4\x58\x24\x0b\x13\x98\x72\xe4\xe1\x71\x36\x3d\xcd\x8d\x1c\x22\x69\x84\xcd\x69\xa5\x37\x38\xee\xd8\x18\x6f\x3e\x8f\x24\x79\xcd\x08\xa8\x8e\x54\x9a\x71\xc7\x71\x51\x18\xaf\x2e\x34\xc4\x28\x3f\xae\x83\xbe\x53\x20\x98\x72\xc7\x2d\x3e\x6c\xb7\x2a\xfc\xf8\x44\xdf\xa8\x0d\x70\x24\xf9\x68\xb8\x10\x36\x59\x26\x39\xae\x90\x2e\x97\x96\x92\xb8\xfc\x9d\x53\xb2\x9a\x1f\xcf\xc1\x86\xaf\x65\x90\xaa\x7e\xbe\x52\x8d\xff\x56\x6c\x31\x1a\x30\xbf\x34\x5a\xf2\x1e\xdd\x8e\xb0\xb6\x74\x75\xe4\xe1\xde\x7f\x60\x49\x9c\x94\x16\x8d\x00\x0c\x16\xc3\x40\x1a\xca\xf4\x74\x68\x6c\xbd\x9f\x6c\x46\xac\x06\x34\x70\x39\xb3\x2c\xe8\xb1\x1d\x58\x4a\xbe\x4b\x9d\xe3\x94\x59\x81\x28\x08\x46\x04\xbe\x70\x0c\xb5\x11\x3d\xbf\xa6\x79\x1a\xff\x8b\x61\x81\xce\x08\x16\x18\xa6\xf8\x7f\x11\x93\xe2\x54\x91\xaa\x5c\x4a\xe7\x08\x1a\xca\x73\x14\x9e\xe7\x57\x29\x8e\x24\x35\xa0\x01\xf6\x08\x13\xac\xec\xd3\x6d\x07\x58\x24\x1c\xe0\xcd\xb9\x38\x31\xd3\x9d\xf6\xae\xb6\x63\xfb\x54\xf1\x6b\x32\xbb\xda\x27\x6a\xec\xf4\xfd\xfb\xc2\xc2\xd2\x24\x0d\x6a\xcb\x9e\xbc\x7c\xe0\xb8\xf9\x4d\xc3\xe4\x4b\x58\x82\x1b\xfe\x86\x0e\xb8\x1d\xba\x9e\x71\x90\x13\xfc\x47\x5e\xc6\x3c\x72\xce\x44\x5a\xe5\x94\x60\x01\x5c\x1f\x43\xc7\x89\x7a\xae\x85\x4d\xa9\x9c\xe5\xe1\x10\x56\x0b\xd4\x83\xc2\x97\xe5\xc9\x5f\xee\xe3\x4a\xca\xd2\xf0\x34\x40\xcc\x5b\xa9\x9d\x5d\x18\x76\x81\x84\x85\x25\x47\xe2\x30\x0c\x73\x96\x09\x70\x0c\x4a\x57\xc0\x8f\x6b\x7e\xec\x02\xa7\x00\xc7\xc3\x4c\x39\x39\x01\xbf\x22\xf1\xcb\xc5\x87\xf7\x00\x6f\x32\xad\xa6\x9a\x62\x69\x9a\xc1\x06\x32\x7e\x0d\x53\x29\x4e\x15\x4d\x26\x30\x42\xca\xab\xfa\x74\x8d\x39\xc0\x1c\xe4\x5c\xbb\x65\x82\x41\xc2\x33\xc8\x10\x11\xda\xab\x92\x88\x80\xb3\x57\xb2\xef\x1d\x25\x6b\xfa\xea\x97\xb3\x57\x00\x72\xf0\xe1\x0a\x45\xe2\xec\x95\x37\xa3\x0c\x76\xc1\x02\x04\x15\x06\x32\xce\x41\x8c\x51\x56\xb1\x0b\x27\x80\x82\xd3\x53\x40\x70\x6a\xf9\x32\x46\x31\x08\x4e\x97\xf2\x3f\xb6\x4f\xc2\xa5\xc1\x7a\xba\x91\x98\xd5\x16\x74\xd0\xa2\x97\xca\xe7\x7f\xdc\xce\x2c\x2b\x17\x5e\x08\xca\x50\x0c\x1a\xad\x96\x68\x4d\x8f\xa4\x44\x09\xb7\x25\xcc\x1f\x4e\xc1\x7c\x6e\x51\xc7\xbb\x87\x9d\x2a\xc9\x85\xda\xed\x3d\x8b\xff\x0b\x7d\x0b\xba\x01\x96\x3e\x9a\xb3\xa7\x0c\x16\xbd\xb0\xbb\x41\x35\x75\x6c\xf8\x57\x7b\xd3\x72\x2d\x19\xad\x89\x17\x0f\x5a\x13\x0d\x76\x01\x83\xb7\x9a\xc5\x1f\xe1\xed\x42\xeb\xa1\xaf\x1a\xee\x43\x03\x71\x22\xd7\xd4\x21\xfd\x6d\xf8\x4f\x62\x18\x13\xf0\xc5\xcf\xaa\xe3\x87\x9e\xe5\x11\x63\x8e\xc0\x0f\xac\xc7\x3d\x4a\x7c\xda\xa3\x5a\xa1\xd4\xd3\x45\xa7\x2e\x4e\x84\xb4\xbb\x2e\x1a\x45\xf4\x3b\x27\xba\x62\x98\x6b\xc8\xe2\x88\xc6\x28\x6e\x46\x33\xca\xbf\xf5\x56\xb4\x9d\x43\x98\x86\x61\xff\x25\x45\x37\x48\x65\xda\x9b\x1b\x4a\x76\x84\x2f\x53\xc8\xb9\x96\xd9\x59\xbd\xa3\x3c\x71\xac\x60\xb7\xce\xfb\xf2\x34\xee\x8f\x6f\xe6\xfe\x5c\xee\x4d\x98\x94\x69\xe2\x9e\xc4\x89\x17\x1d\x8a\x90\x07\x91\x21\x99\xe8\xa6\xd8\xce\x41\xd9\xcb\x23\x86\x33\x51\x5f\x30\xbd\xa2\x91\x9b\x61\xa2\x51\xae\x7c\x50\x35\x26\xa1\xcc\xf2\x64\x7c\x85\xfe\x8a\x46\x7d\xe2\xbe\x94\x44\xf1\xe8\x17\x18\x7d\x15\x38\xfa\xca\x07\xb0\xbb\x74\xae\x1a\x26\x92\xee\x6b\xab\x15\x8e\x7d\xc8\x26\x1b\x11\x5e\x64\x0c\x13\x91\x04\xf3\xbf\xff\xc8\x57\x3f\xf2\x7f\xcc\x97\x80\x86\xb5\xcb\xae\xa2\xa0\xba\x49\x7b\xb6\x8b\x96\xa8\x3c\x8d\x68\x25\x33\x1d\x7f\xfd\x8a\x08\x62\x50\xa0\x5f\x91\x10\x88\x81\xb0\x15\x5e\x69\xaf\xac\xd3\xec\x35\xf3\x67\xad\x01\xc6\xe4\x30\x14\x21\x7c\xd3\xf4\x48\x8f\x86\x3d\xad\x2e\x80\xc1\xc2\x5d\xa7\xbc\xba\x73\x59\xda\xe3\x17\x34\xd2\x2a\x6d\x16\x5c\x0c\xb0\xe0\xa2\x87\x05\x95\x53\x9e\x31\x9a\x21\x26\x83\x29\x0f\x46\x80\x9c\x4b\x4d\xa8\x13\x7d\xca\xa5\xf7\x67\x4f\x0f\x36\x2a\x35\xff\xbe\xbe\xdf\x6c\x31\xaa\xf2\x51\x7b\xcf\x31\x0b\x42\x63\x6b\x34\x77\x06\x24\x31\x08\xfa\xb6\xc7\xa2\xdd\xa5\x6f\xe3\x16\x86\x9f\x9d\x39\x58\xae\x9b\xba\x0f\x2c\xe5\x5e\x95\xe3\x51\x5c\x26\xf2\xf7\x9a\x3e\x1d\xd9\xc9\x5e\x59\x53\x3d\xc4\xce\x9d\x5a\x1e\x59\x8a\x88\x99\xbc\x90\xbe\xd9\x4f\x96\x6b\x74\x72\x02\x08\x4d\x31\x11\x2b\xb0\xa6\xfa\x49\x04\x6f\xfa\x4d\x4f\xc7\xb3\x9d\x16\x44\xd0\xf1\x2a\x61\xd0\x2e\x34\x27\xe2\x04\x5c\x43\x7e\xce\x50\x82\xbf\x81\x40\xe7\xdc\x94\x26\x39\x29\xb7\x05\x98\x3f\x9f\xb7\xa7\xb7\xd5\x6b\xd5\xa3\x76\xcb\xd6\xc2\xb6\xcb\x35\x0c\xf1\xa9\x37\x48\x37\x3d\xd3\xd3\x5c\x38\x5e\x71\xa6\xdc\x62\x1f\x9e\x17\xf6\x2d\x57\x52\xdf\x71\x19\x45\xb1\x02\xa5\x5b\x2c\xa2\x6b\x90\xec\x4b\x4c\x11\xe4\xfa\x09\x46\xb9\x6b\xe7\x2b\xa7\x7f\x0f\xa2\xd4\xac\x68\xb3\x79\x2c\x06\xf3\x11\xea\x20\xec\xa7\x83\xc1\xe2\x1e\x04\x5c\xc6\x7d\x99\xe5\x04\x36\xf2\xcf\xda\x5a\x99\x16\x4b\x2a\x48\xb7\x68\xbb\x05\xeb\xf6\x0d\x64\x5f\x51\x2c\x43\xba\x4b\x54\x66\xb6\x2f\x5b\xe6\xbe\xec\xf2\xcf\xd1\xb4\x30\x08\x2a\x18\x96\xed\xa9\xba\xcb\xac\x3a\x5b\x80\x40\x06\x62\x55\x86\x02\x4c\x89\xb7\x1a\x81\x95\x9d\xaa\x1f\xc8\x0f\x14\x3a\x23\x02\x4e\x2d\x32\xcd\x54\xe3\x0a\x75\x4e\x1a\x09\x19\xa5\x9f\xf4\x5a\x52\x91\x04\xf3\x9c\x28\xd9\x08\x5a\xae\x60\x3d\x47\x7a\xd6\x01\xfa\x99\xda\x99\xcf\x06\x0f\xd5\x67\x20\xf8\x91\x2f\x56\xe0\x47\x3e\x6f\xba\x5a\x8a\x9c\x56\x86\xc2\x37\x84\x53\x91\x43\x53\x7d\x62\x74\x1f\xf5\x31\xb3\x27\xa8\x4f\x0b\x83\xbf\x96\xfa\x18\xf4\xf7\xae\x3e\x86\x91\x0f\x59\x7d\xec\x5e\xa9\x4b\xe7\x50\x9e\x1f\x30\xcb\x52\xfd\x88\x86\x50\x35\xb4\x4e\x0a\x43\xe0\x71\x27\x5a\x3e\x46\x99\x78\x8d\xa0\x16\x0f\x8c\x9b\x36\xe4\xf2\x58\xb9\xe3\x63\x7d\xdb\xc2\xeb\x77\x8c\xbe\x1a\x63\xe6\xd5\xda\xf2\x83\x5e\xd9\x0e\x8d\xce\xf8\xeb\x3f\x72\x98\x06\x76\xbc\xb4\xb0\xc4\x9f\x41\x82\xa3\x60\x1e\x41\x22\xfd\xd1\x4c\x31\x2f\x61\x74\x03\x20\xd0\x54\xdc\x62\x71\x0d\x62\x9c\x24\x88\x21\x22\xaa\x37\x56\x73\xe7\xd6\x98\x53\x75\xe9\xa5\x57\xf7\xbb\x73\x9e\xb9\xc7\x7a\x8b\x16\xde\x9f\x59\x75\x15\xf8\x1e\xa7\x77\x7f\xb6\x6a\xe4\xd8\xee\x3a\xae\x7b\x81\x3d\xf7\x82\x66\x27\x19\x5a\x6d\x43\x57\x02\xaf\x10\xca\x1a\xcf\xc9\x62\x84\x32\xfd\x82\x0a\x8f\xbc\xa0\x52\x2f\x3f\xbd\x6d\xa4\x5e\xc8\xf7\xee\xd5\x4d\xb0\x3e\x79\x62\xed\xdb\x27\xc5\x6c\xf6\xc4\xbc\x94\x18\xbe\x9b\x2d\x66\x4f\x68\x58\xae\x7c\x46\x04\x0d\x68\x2e\x16\xb3\xd9\x93\x8e\xf7\x3a\xf5\x20\x49\x3c\x46\xdc\x8d\x29\x31\x31\x9b\x5a\x1f\x12\x83\x34\x4c\x66\x4a\x89\xda\xd8\xf8\xed\x6c\xf6\x44\x40\xb6\x46\x62\x59\xa6\x86\x9d\xe7\xd6\xa1\xe2\x30\x5d\xcc\x9e\x98\xdc\xf1\x0f\x35\x03\xf5\x5e\x75\x12\x22\xff\xb4\x4c\x35\xca\x94\xc8\x87\xd6\x37\xf6\x57\xda\xdb\x85\x16\xc2\x73\xfd\xa6\xec\xb9\xc6\x69\xe8\x2d\x99\xda\xb5\xe6\x4d\x88\x7e\xba\xad\x6e\xda\x70\x6c\xf8\x1c\xe5\x4c\x5b\x08\x92\x50\xb6\xd1\xa9\x2b\xae\x13\xd0\x15\xe7\x6b\x32\xbd\xf3\xab\x66\xa9\xa0\x91\xbd\x57\xbf\x28\x9b\x59\x9b\x59\x75\x7e\xf1\x6d\x79\xd1\xf8\x47\x8e\x19\x8a\x5f\x0f\x0d\xdc\xf1\xbc\xae\x12\x5f\x9f\x18\x24\x1c\x4b\xaa\x9d\xbe\xf0\xf5\xb7\x8c\x72\x54\x1f\x59\xa6\xf9\xa3\xc1\xc9\x1d\x8d\xfe\x00\xfa\x8d\xf5\x5c\x07\xcb\x73\xcb\x0a\x1a\x15\xa9\x51\x2f\xf9\x51\x82\x32\x47\xbe\x13\xe1\x2c\x7b\x6c\x51\xbf\x0b\xe0\xb0\xea\xb4\xd1\x10\x9a\x77\x92\xad\x53\xda\x39\x95\x35\x2d\x94\x81\xa0\xa6\x07\x91\x7c\x33\x5f\xdc\x9f\x1c\xde\xeb\xd7\x48\xaa\xbe\x03\x59\x35\x49\x02\x6f\xd0\x24\x01\x7d\xc2\x1b\xf4\x50\xc5\xf3\x4d\x20\x46\x60\x3a\x5f\xd8\xad\x29\xe6\x62\xbe\x98\x40\xe1\x6b\x03\xe6\xc1\x50\x59\xd3\x82\x89\x40\x6b\xc4\x26\x09\xec\x8c\x88\x07\x48\x49\x92\x52\x28\x26\xd1\xf1\x46\xce\x78\x18\x94\x0c\x11\xc6\x50\x32\xf7\xba\x4f\x77\x51\xab\xe9\xfe\x88\x38\x12\xe6\x4a\xe7\x0d\x65\xff\x83\x18\xd5\x9f\xc2\x8c\x66\x47\x6a\x2e\x76\x8f\x0c\xeb\xc3\xa7\x87\x41\xd6\x49\x74\x6a\x7e\x68\x31\x05\x58\x59\x15\xcf\x8d\xc9\x50\xf2\x56\xed\xc2\x46\xe3\x3b\x98\xd5\xe6\xd4\x24\xd3\x78\x7e\x65\x3d\x19\xef\xe6\xde\xd6\x26\x59\x4e\x68\x5d\x7b\x03\xf5\xe0\x9f\x08\x4c\x72\xd4\xc0\xda\x97\xdb\x3c\xbf\xea\x62\x2d\xcf\xaf\xbe\x23\x1f\xc3\x17\x69\x4a\x6f\x51\xfc\xf2\x9a\xe2\x08\x71\x9f\xfd\xa2\x8f\x9c\x33\xa2\x9e\xa7\x4f\x3a\x78\x96\xf5\x55\xa3\x9c\xf7\x3b\xc5\xa4\x85\xc0\xe5\x7c\x09\xe6\x97\x12\x5a\xb1\x54\xae\xd9\x8b\x5c\xd0\xb5\xb9\x50\x89\x07\xf6\xde\x18\x3b\xbc\x98\x00\x99\x17\x0b\xce\xa1\x90\x26\xdc\xcf\x58\x2c\xd5\xfd\x61\x73\x8d\xcb\x8e\xe6\x77\x88\x73\xb8\x46\xba\x57\x76\x5a\xfe\xcf\x01\xc8\x5e\x0b\x10\xbe\x83\xdf\xde\x22\xb2\x16\xd7\xe0\x27\x1f\xc2\xdf\xc1\x6f\x78\x93\x6f\xf4\x14\x5f\xf2\x65\x6b\xbd\x8e\x6c\x51\xf1\xe5\xa1\x28\xc2\x64\x12\x45\x98\xec\x48\x51\xb5\xce\xc1\x29\x82\xdf\x94\xc9\x00\x3f\x85\x3f\xf5\x79\xc2\xfe\xc7\x9d\x11\xe1\x84\xd3\xae\x92\xe0\x6f\xe6\x4b\xc6\xbd\x91\x6b\x32\x02\xbe\x38\x7b\x7b\x1a\x4b\x19\x41\x05\x0d\xac\x17\x7b\x96\xd2\x98\x16\xee\x53\x66\x5a\x49\xa7\xcb\xac\xc4\x62\xff\x32\xf3\x44\x79\x17\x91\xd5\x48\x7f\x1f\x91\x55\x8f\xd0\x43\xd0\xfa\xf4\x5a\x3d\x52\x57\x9f\x3e\x0f\x7d\x7f\x5d\x33\x42\x82\xdb\x94\xc4\x2a\xca\xad\x77\xe7\x35\xf9\xba\xd1\xd7\xad\xf4\x25\xf3\xb8\xc7\x8f\xf4\x60\x82\x43\xef\x4d\x45\xa9\x42\xac\xca\xb3\xea\x84\x43\xcd\x07\xfb\xe3\xe9\x97\x39\x17\x74\x53\x5e\xa2\xd7\x10\xc2\x3a\x6b\xbb\x81\x59\x86\xc9\x5a\x7d\x81\xad\xde\x79\xd5\x90\xde\xe9\xae\xd0\xfc\x1f\xcc\xeb\x8f\xf3\x5b\xe8\x34\x72\xba\x25\xd4\x6e\x51\x18\xb8\xa5\x40\xe8\xde\x58\xdc\xc5\x71\x73\x1f\xef\x3a\xfd\x0b\xf0\x0f\xeb\x5a\xde\x64\xe1\xdc\x21\x66\x05\x1b\x06\xea\x98\x6b\xbf\x76\x6c\xcc\xda\xe5\x91\x9f\xec\xc1\x09\x8e\x14\x67\xdf\x50\x56\xe5\x71\x9c\x27\x14\x55\xab\x33\xbc\x7a\x63\xa5\x53\x83\xf5\x75\x87\xfa\x18\xfe\x2b\xba\x2b\xf3\x55\x63\x4f\x99\xfa\x70\x08\x14\xa0\xf6\x63\x88\x1e\x74\xea\x0c\xea\xcd\x12\xd0\xaf\x46\xfc\xbd\x0b\xd7\x29\xab\x77\x30\xfb\x2c\x97\xfa\xf2\xb3\x9c\xd6\xe2\xf4\x8d\xcd\xe4\x93\x13\xf0\x2f\x04\x22\x9a\xa7\xb1\xe2\x6d\x82\x49\x0c\xb0\x58\x02\x4e\x41\x8a\xc4\x33\x0e\xa2\x6b\x14\x7d\x05\xd4\x7c\x8c\x47\x6f\x11\xd3\xf7\xe9\x98\xc4\xe8\x1b\x8a\x01\xcf\x50\x04\x36\x30\xb3\x65\x36\x84\xe7\x5b\x09\xe2\x25\xe4\xa8\x03\xe1\xf2\xfb\xe0\x4e\x86\x70\x47\x86\x49\x9e\xa6\x96\x8c\xb8\x3b\x72\x03\x33\x4f\x69\xf5\xac\x15\x2c\x24\x8c\xcf\x5a\x58\x5f\x7c\x65\xe5\x41\xbe\x43\x75\x9d\x4a\xcd\x51\xaf\xb6\xea\x4b\xab\x1e\xe5\x74\x5e\x54\x43\x70\x83\xd8\x1d\x80\xf1\x0d\x24\x11\x8a\x81\x64\x80\x42\x4f\x5c\x43\x01\xee\x68\x6e\xde\x72\x29\x49\x13\x84\x62\x70\x95\x0b\x80\x09\xe0\x74\x83\x24\x20\x35\xbd\x64\x25\xc8\x39\x52\xa2\xf6\x7b\x9c\x69\x12\xb5\x2e\x21\xae\xca\x5b\xdf\x03\x94\x1c\x33\x4f\x3d\xd4\xb0\x6d\xc3\x6e\x7b\xa6\x62\x87\x5e\x77\x0c\x3f\x76\xeb\x30\x7f\x7d\xb7\xd3\xde\x22\x6d\x7d\x40\x08\x33\x75\xe1\x58\x89\x56\x0a\x72\xf8\xd6\x61\xac\x8a\x85\xbb\xde\xe9\x14\x45\xdd\x36\xcf\xc6\x49\xe9\x6e\xeb\x63\xb4\x6a\x86\x44\xa1\xf1\x18\xf0\xd8\xae\xe0\xd2\x64\xfa\x7c\xd5\xbe\xb6\xeb\x8c\x55\xe5\x3f\xbb\x7d\xd5\x8e\x2d\x65\x4c\xe9\xc0\x6a\xbc\x63\x71\xc3\xf0\x55\x4f\x7a\xe0\xb8\x28\x26\x85\xf0\xb5\xc3\x58\x4d\x2b\x2a\xdf\x63\xd9\xa6\xad\x11\xeb\xd7\xd8\xd9\x1d\xab\xce\xbc\xc0\x20\x75\xe5\x02\xbd\x57\x8a\x4e\xc7\xaa\x47\x1c\x7e\x4b\x30\xa4\xd4\xe7\x03\x49\xef\x9c\x15\xac\x76\x4d\x41\x63\xa4\x17\x74\x93\x49\x2a\x1d\xe4\xaa\xdf\x6e\x57\xd0\xd5\xf7\x8a\xce\x68\xe7\x83\xc5\xd0\x7f\xc1\x8c\xa1\xa8\x29\x8f\xba\x55\x93\xe2\x8c\xf2\x84\xeb\x3c\xea\xae\x01\x57\xcd\xab\x8e\x77\xd7\x8d\xc7\xd6\x5e\x2b\xb5\x9e\x89\xc8\x7f\x55\xa3\xc6\xdf\x1e\xe3\x07\xb4\xbe\xb0\xaa\x40\xea\x26\x03\xb0\xea\xf7\x02\xf7\x06\xa7\x02\xb1\xf2\x81\x58\xd9\x5b\xb7\x6a\xa0\xce\x28\x3f\xb8\x94\x21\xbc\x26\xff\x8d\x1c\x55\xac\x5b\x0d\x5c\x7b\x94\x17\x5c\xf3\xb8\xdb\xea\xd1\x2d\x1a\x5e\xd5\xeb\x05\xab\xfd\x79\x8d\xfc\x57\xb7\x6a\x98\xce\x28\x2f\xb8\x76\xca\xa9\xea\xac\x1a\x57\xed\xb4\x94\x27\xd0\xd6\xde\x2b\xdb\x56\xad\x3c\x89\x17\x44\x2b\x8f\x54\x83\x2c\x1b\x57\xed\x5c\x93\x27\xd0\x36\x9a\xa6\x6d\xd5\x4a\x0d\xf8\x40\x6c\x1a\x4c\xcb\x4e\x4e\x32\x8f\xea\x23\x92\xa6\xa2\x57\x8d\x1a\x37\x7b\x8c\x17\xd0\x73\x86\x37\x90\xdd\x35\xd4\xbc\x6e\xd5\x60\x9d\x51\x5e\x70\x3f\x22\x18\x37\xed\x78\xd9\xb6\x32\x19\xda\x6a\x84\x27\x44\xf7\x46\x5b\x43\xd4\x6d\xab\x66\xce\xd7\x0b\xe2\x05\x8a\x18\x72\xbe\x27\xd7\x2d\xe5\xf7\xf8\xa6\xd7\x13\x56\x73\x5b\x5f\x58\xdb\xfa\x62\xd2\xb6\xbe\xc0\x6b\xe2\xd2\xa9\x5b\x0c\xac\xb2\xd7\x0f\x96\xf3\x5d\x9e\x82\xa5\x5a\x0c\xac\xb2\xd7\x0f\x56\x7e\x65\x3e\x5e\xa8\x81\xe9\xa6\xb2\x34\x59\x35\xc0\x4f\xa3\x5b\x6f\x1a\xe4\xbf\xaa\x51\xa3\x68\x8f\xf1\x03\xda\x40\xd1\xc2\x6f\x14\x39\xb3\x42\x7f\xb2\x62\xdc\x4d\xef\x0e\x3c\xbf\x83\xbf\xde\xb3\xf0\xc3\x76\xdc\x75\xa4\x1f\x3e\x7a\xed\x8f\x5e\xfb\xa3\xd7\xfe\xe8\xb5\x3f\x7a\xed\x8f\x5e\xfb\xa3\xd7\x3e\x0e\xf4\xd1\x6b\x7f\xf4\xda\x1f\xbd\xf6\x09\x5e\xfb\xb6\x5d\x4e\xe0\x1e\x1f\x4d\xeb\x3b\x40\xff\x82\x8b\xdd\xa5\x8c\x7d\x21\xe8\xba\x7f\x93\xd6\xfb\xfc\x65\xec\xc3\x98\xbd\x15\x37\x9e\x82\xd7\x9f\x5a\xe2\x78\x5a\x85\xd0\xdd\xc8\xdb\xa9\xdc\xf1\x94\x25\x0e\x52\xf4\xf8\x7b\x70\xe6\x40\x05\x90\x77\xe7\xdd\xfd\xca\x20\x8f\xee\xae\xef\x50\x0c\x79\x9a\x00\xfe\x53\x4b\x22\x4f\xe3\xd2\x83\x28\xfb\xa3\x8b\x72\x9c\xa7\x10\xbb\x75\xa4\x26\x1d\x01\x4e\x75\xe4\xc3\xee\x6d\x83\xeb\xc3\xd2\xae\xb0\xc2\x6a\x50\xcf\x76\xa9\x39\x3c\x8d\x3b\xbb\x57\x1e\x1e\xf7\x32\x3a\x4a\x09\xb7\x0b\xb7\x0c\xd5\x14\x0e\xbd\x9c\x0b\xa7\xb4\x70\x87\xba\xfb\x94\x89\xeb\xa8\x13\xac\xca\x09\x87\x1d\xfe\x9a\xfa\x7d\x3f\xa5\x82\xbd\x18\x68\x17\x0c\xf6\xe0\x45\x58\xd5\x0d\x1e\x1f\x1b\x2c\xbc\xbe\x01\xdf\x3a\xe7\xf9\xf8\x84\x6d\xa9\x20\x5e\x15\x6a\x8d\x4e\xb4\xbe\x30\xf7\xf8\x78\xfa\x10\xe5\x6a\x7d\x0b\xd0\x3a\x68\x8f\x94\x55\xf5\x27\x65\x97\x6a\xb4\xd5\xd7\x42\xe3\xa5\x55\x7b\x8a\x1e\x98\x2a\xb4\x85\xcd\xae\xe7\x07\xaa\x6e\xeb\x5b\xaf\x76\xdf\xfc\xbd\x67\xf1\x5a\x9c\x00\x1c\xb7\xca\x99\x7a\x97\xb4\x7d\x6a\x6a\xda\x16\x9a\x4f\x3b\x40\xa8\x2b\x71\x76\xb2\xf4\xcf\x2e\x88\xeb\x23\x82\xef\x59\x16\xd7\xcb\x4a\xed\x78\x5a\xec\xb1\x38\xae\x4b\xa4\x7f\x75\xdc\xe7\xd3\xcb\xe3\xf6\x55\x51\x71\xd1\xd8\xa5\x88\x6e\x4b\x25\x87\x7f\xb5\x2d\xc4\x83\x2a\xa5\xeb\x69\x46\x0e\x5e\x50\xd7\x5b\x77\xff\x22\x65\x75\x71\xac\x3e\x80\x1c\xab\xa1\xdb\x5b\x38\xe4\xa9\x5d\x10\xdc\xd5\xe8\x0e\xa0\x4d\x8d\xde\xa5\x14\xef\x3d\x34\xba\x56\xe7\x7b\x55\xdb\xf5\xd1\xc5\x7b\x39\xef\x5d\x31\x5c\xa6\x5a\x1a\x98\x99\x83\x78\x17\x04\xbb\x62\x2f\xd5\xd2\x55\xf5\xd0\xc4\x5e\xa3\x7f\xf0\xa2\x62\xfe\xa4\x47\x00\x63\xd6\x6f\x4f\x85\x74\x72\xb1\x83\xfd\xec\xd2\xed\x11\x68\x3d\xf6\xbf\x37\x09\x3d\xa4\xb1\x26\xde\xdc\xfe\xc9\x75\xed\x7c\x14\xea\xa1\x56\xb7\xeb\x12\xc7\x78\x79\xbb\xee\x59\x87\xaf\x6f\x57\x71\xfa\xff\x6f\x95\x3b\x1f\x65\x7a\xa8\xb5\xee\x7c\x95\xc9\x2d\x76\xb7\x47\x65\x9a\x54\xed\xee\x4f\x54\xa6\xed\x5f\xb7\xac\xb5\x07\xd7\x86\x8a\x5b\xab\x0a\x5b\xdb\xa1\xc2\xcd\xfd\x51\x47\x8f\x67\xda\x91\x88\xec\x3b\x64\x1e\x66\x81\x6c\x18\xc7\x0c\xf1\x2a\xd7\xde\x5d\x2f\xdb\x8b\xef\x87\xab\x9a\xfd\x74\x3b\x5e\x36\xdb\xb3\x94\x9d\x77\x26\xd0\xdf\x1c\xd6\x65\xed\xbc\x92\x82\x1d\xc1\x4e\x7f\x71\x3b\xaf\x98\xe6\x60\x25\xee\x0e\xc6\xac\xba\xdc\x9d\xcf\xac\xc3\x17\xbd\x1b\xc7\xc2\xa3\xf4\x9d\x4f\xd9\x4a\x47\x65\x55\x9a\x7e\xc2\x5f\x85\x69\xe6\xeb\x27\xfc\x39\xc0\xc1\x00\xd4\xa4\xf1\x55\x98\xdf\xf3\x37\xff\x86\xcf\x98\xbd\x3c\xc0\xe8\xe2\xc6\xf4\x3b\x8c\x43\xf3\xa4\xf7\x7e\xa3\xc5\x12\xeb\x97\xff\x0b\x00\x00\xff\xff\x3f\xf9\xae\xe2\xa9\x7f\x00\x00")

func templatesModelGotplBytes() ([]byte, error) {
	return bindataRead(
		_templatesModelGotpl,
		"templates/model.gotpl",
	)
}

func templatesModelGotpl() (*asset, error) {
	bytes, err := templatesModelGotplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/model.gotpl", size: 32681, mode: os.FileMode(420), modTime: time.Unix(1592953271, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesRelationships_registryGotpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x98\xef\x6f\xf2\x36\x10\xc7\xdf\xe7\xaf\x38\xa1\x68\x6a\x27\x96\xad\xed\x3b\xa4\xbe\x98\xca\x54\xf5\x45\xab\x8a\xfd\x78\x83\x78\xe1\x91\x03\xbc\x19\x27\x35\x0e\x15\x8a\xf2\xbf\x4f\x26\x09\x84\x60\xc7\xe9\x36\x1e\x3d\x7d\x74\x57\xa9\x12\xf6\xdd\xf7\xce\xf6\xd9\x7c\x44\xca\xe6\x7f\xb3\x25\x42\x9e\x43\xf4\x2b\xea\xe8\x21\x91\x0b\xbe\xcc\x14\xd3\x3c\x91\xd1\x0b\x5b\x23\x14\x45\x10\xf0\x75\x9a\x28\x0d\x83\x65\x12\xb1\x34\x51\xa8\x93\x88\x27\x3f\xa2\xc0\x35\x4a\xcd\xc4\x20\x08\xb6\x4c\x81\x42\xb1\x8f\xdb\xac\x78\xba\x99\xe0\x92\x6f\xb4\xda\xc1\xc1\x2b\x9a\xd8\xe6\x83\x60\x91\xc9\x39\x70\xc9\xf5\xd5\x35\xe4\x41\x00\x00\x0e\xa5\x7b\x9f\x56\x5e\x94\xe1\x79\x0e\x8a\xc9\x25\x42\x88\x52\x73\xbd\x33\xeb\x18\x42\x58\xab\xc2\xe8\xbe\x5c\xed\x89\x88\x59\x68\x19\xfc\x03\xf0\x05\x6c\x56\x49\x26\xe2\x52\x19\x55\xd3\x13\x42\x13\xdc\xd4\x86\x30\x7a\xcd\xfe\x14\x7c\xfe\x9c\xc4\x58\xc9\x58\x97\x30\xcd\xf3\x93\xb8\xa2\x78\x8a\xcb\x8f\x33\xb8\x87\xef\xec\xcb\xcb\xf7\x7a\x8d\xd2\x96\x1a\xae\x04\xca\xe3\x82\xa2\x07\x85\x4c\xe3\x35\xfc\x54\x2f\xc2\x58\x39\x38\x82\x35\x4b\xa7\x1b\xad\xb8\x5c\xce\xbe\xb7\x67\x78\x92\x8b\x04\x8e\x69\xea\x54\xd5\x1e\xa6\x4c\xa1\xd4\x43\x08\xd9\xbc\xde\xbd\x76\xe6\x66\x5a\xe7\x1e\x3e\x49\x69\xdf\xc8\x32\xc1\xc9\x26\xb6\x05\x07\x66\xe3\x2a\xbf\xa2\x18\xc0\xe8\xb4\xda\x46\xce\xaa\xca\x68\x8c\xa9\xc2\x39\xd3\x18\xb7\xb5\x8c\x1d\x67\x47\xa0\x55\x86\x43\xab\x1c\x4a\x6b\x70\x2b\xd3\x2b\x53\x6c\x8d\x1a\xd5\x18\x17\xa6\x8d\xcd\x1e\xb9\xa3\x0e\x87\xe7\x8e\x8e\x26\xf8\x96\x71\x85\x71\xeb\x40\x6b\xab\xa7\x0f\xa1\x9b\x51\xe3\x66\xbc\xe0\xfb\x71\xa2\x72\x35\x53\x57\x67\x3a\xc6\xa6\x33\xf3\x57\xb6\xc7\xf9\x9e\x36\x6b\xaf\xba\x81\x0f\x21\x14\x37\xfb\x26\xe8\xb1\x02\x5b\xf9\xae\x0d\x11\x37\x8e\xf5\x1e\x02\x9c\x33\xb6\x1a\x6f\xf7\x35\x8a\x9b\x2e\x45\x6b\x19\xb7\x9e\x32\xfc\xa5\xd8\xca\xb9\x2b\xcb\xb9\xf5\x29\x43\xdd\xed\xe2\xce\x74\xfa\x79\x63\xda\x12\xa1\x8c\x3d\xb2\x45\xb7\x50\x3f\x11\xbf\x57\x47\x1a\x7f\x70\xb7\x87\x45\xfa\xfa\xdf\x5c\xdb\x3e\x17\xf0\x17\xa9\x15\xc7\x8d\xa3\x11\x9a\xf7\x6e\x3a\x3b\xde\x3c\x8b\x92\xfd\x99\x6a\xf4\x45\xea\xbb\x49\x55\x29\xce\xb6\x71\x77\xa2\xf9\x86\x19\x55\x2f\x67\xfd\x55\xde\xd1\x4f\xbf\xed\xd2\xa3\xbb\xf9\xd0\xed\x5e\x3f\x82\x69\x34\xc6\x05\xcb\x84\xfe\x83\x89\xec\xec\xdd\x6e\x5a\xd3\xef\x90\xa8\x15\xec\x49\xe8\xef\x1f\xbe\x00\x7c\x3b\xac\x60\x80\x32\x5b\x0f\xba\x8a\xfa\x59\x88\xe4\x1d\xe3\x87\x55\xc2\xe7\xb8\x3f\x4f\xdf\x5b\x08\xa7\x87\xf8\xd7\x10\xc2\xed\xfe\x10\xd3\xe8\x54\xcc\x77\xd3\xf7\x3b\xb0\xf5\x5f\xf2\x8e\x96\xae\xcd\x7f\xed\x7a\x3c\xc4\x61\x1a\x3d\x67\x42\xf3\x54\x74\x1e\x63\xed\xe3\xfa\xe2\xec\x99\xd8\x52\x72\x47\xc4\xc7\xbc\x1d\x53\x2d\x11\x87\x97\x65\xb8\x11\x68\x99\x75\x52\xd9\xef\x69\x7c\x4e\x65\xe5\xe0\x85\xa9\xac\x4c\xf2\xe5\xa9\xcc\xf2\x1a\x11\x96\x11\x96\x11\x96\xb9\x8c\xb0\xec\xdb\xc4\x32\xa2\xb2\x76\x49\x44\x65\x1f\x8c\x26\x2a\xf3\x0c\x37\x02\x5f\x99\x9e\xaf\x88\xa9\x88\xa9\x88\xa9\x7c\x71\xc4\x54\xc4\x54\xc4\x54\xc4\x54\xc4\x54\xc4\x54\xff\xdb\x2f\x5d\x63\x14\x78\xf6\x4b\x57\x39\x78\x61\x2a\x2b\x93\x10\x95\x11\x95\x11\x95\x75\x97\x43\x54\x46\x54\x06\x44\x65\x44\x65\xb5\x11\x95\xf5\x8f\xf8\x94\x54\xf6\x88\xba\xf5\x80\x4c\xd0\xdc\xe5\xed\xa5\xa1\xec\x11\x35\x11\x19\x11\x19\x11\x59\x77\x39\x44\x64\x44\x64\x40\x44\x46\x44\x56\x1b\x11\x59\xff\x88\xcf\x4a\x64\xcf\x4c\xee\x1c\x54\x66\xa6\x2e\x4f\x66\x26\x0b\xd1\x19\xd1\x19\xd1\x59\x77\x39\x44\x67\x44\x67\x40\x74\x46\x74\x56\x1b\xd1\x59\xff\x88\xaf\x86\xce\x0c\x2b\x11\x53\x11\x53\x11\x53\x11\x53\x11\x53\x11\x53\x11\x53\x9d\x1b\x31\x15\x31\xd5\x7f\xf8\xc5\xab\xfc\xdf\x1a\xcc\xf3\xfa\x53\x11\xfc\x13\x00\x00\xff\xff\x08\x39\x9d\x04\x65\x44\x00\x00")

func templatesRelationships_registryGotplBytes() ([]byte, error) {
	return bindataRead(
		_templatesRelationships_registryGotpl,
		"templates/relationships_registry.gotpl",
	)
}

func templatesRelationships_registryGotpl() (*asset, error) {
	bytes, err := templatesRelationships_registryGotplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/relationships_registry.gotpl", size: 17509, mode: os.FileMode(420), modTime: time.Unix(1585855989, 0)}
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
	"templates/README.md": templatesReadmeMd,
	"templates/identities_registry.gotpl": templatesIdentities_registryGotpl,
	"templates/model.gotpl": templatesModelGotpl,
	"templates/relationships_registry.gotpl": templatesRelationships_registryGotpl,
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
	"templates": {nil, map[string]*bintree{
		"README.md": {templatesReadmeMd, map[string]*bintree{}},
		"identities_registry.gotpl": {templatesIdentities_registryGotpl, map[string]*bintree{}},
		"model.gotpl": {templatesModelGotpl, map[string]*bintree{}},
		"relationships_registry.gotpl": {templatesRelationships_registryGotpl, map[string]*bintree{}},
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
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

