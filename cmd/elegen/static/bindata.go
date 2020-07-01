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

	info := bindataFileInfo{name: "templates/README.md", size: 77, mode: os.FileMode(420), modTime: time.Unix(1578606980, 0)}
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

	info := bindataFileInfo{name: "templates/identities_registry.gotpl", size: 5305, mode: os.FileMode(420), modTime: time.Unix(1586905041, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesModelGotpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x3d\x69\x73\xdc\x36\xb2\x9f\x3d\xbf\x02\x99\x52\x6c\x8e\x77\x44\xf9\xf3\x64\xb5\x55\x3e\xf3\x54\xcf\x87\xca\xf2\x66\xab\x9e\xcb\xb5\x86\x48\x70\x84\x98\x03\x30\x00\x28\x59\x6f\x8a\xff\xfd\x15\x0e\x92\x00\x4f\x70\x34\xe3\x28\x6f\xe5\x0f\x89\x84\xa3\xd1\x17\x1a\xdd\x0d\xb0\x95\xc1\xe8\x1b\x5c\x23\xb0\xdd\x82\xf0\x02\x89\xf0\x25\x25\x09\x5e\xe7\x0c\x0a\x4c\x49\xf8\x1e\x6e\x10\x28\x8a\xd9\x0c\x6f\x32\xca\x04\x08\x66\x00\x00\x30\x4f\x36\x62\xae\x7f\x5a\xd3\x10\x66\x94\x21\x41\x43\x4c\x4f\x50\x8a\x36\x88\x08\x98\x96\xbd\x58\x5c\xe5\x97\x61\x44\x37\x27\xeb\x94\x5e\xc2\x94\xe3\x35\x39\xd9\xac\xe9\xc9\x25\xa7\xa4\x3d\x68\x83\x45\x74\x85\xd2\xf4\xea\x24\xa2\xd9\x2d\x17\x2c\x8f\x44\xce\x90\x1e\xb8\xdd\x1e\x03\x06\xc9\x1a\x81\xf0\x22\x43\x51\xf8\xe9\x36\x43\xe7\x8c\x5e\xe3\x18\x31\x2e\x91\x54\xd0\x24\x1d\xa0\x28\xea\x29\x88\xc4\xe0\xd8\xf4\x36\x41\xfc\x06\x53\x1c\x2b\x4a\x3d\x01\x15\xc5\x6c\x31\x9b\x6d\xb7\xe0\x28\x85\x02\x71\xf1\x1b\x62\x1c\x53\x02\x56\xa7\x06\xe2\x5b\xd5\xfc\x5c\x08\x86\x2f\x73\x81\x78\x39\x40\xf2\x70\xbb\x35\x8b\x1f\xe1\x25\x38\x42\x24\xdf\xc8\x79\x97\x39\x4e\xe3\xd7\x24\xdf\x70\x0d\xa2\x09\xba\x28\x66\x27\x27\x52\x3c\x6a\x86\xa2\x1a\x14\x05\x60\x28\x63\x88\x23\x22\x38\x10\x57\x08\x64\x94\x73\x7c\x99\x22\x70\x0d\xd3\x1c\x71\x90\x50\x06\x60\x89\x85\x22\x46\x4f\xaf\x30\x33\x92\x9d\x87\x33\x21\x21\xb6\xe0\x73\xc1\x30\x59\xcf\x66\x11\x25\xbc\x94\x7b\xcd\xbe\x23\x02\x37\x68\x09\x8e\xd4\x6a\x92\x0a\x3d\xf9\x37\xbd\xb8\x61\xa1\x41\x9b\xe8\x95\x9a\x18\xeb\xa9\x72\x80\xfe\xa9\x28\x42\xb3\x48\x3d\xa5\x85\xd5\xa9\x26\xa5\x9c\x51\x0a\xa7\x96\x4d\xfd\xf3\x4c\x62\x8b\x13\x40\xa8\x30\xb2\x79\x47\x63\x94\x86\xaf\x90\x80\xd1\x15\x8a\x6b\xc6\xda\xbd\xaf\x89\xc0\xe2\xd6\x30\xe7\x2c\x46\xea\xd7\x26\xea\x55\x3b\x4d\xd4\xef\xf4\xf2\x77\x14\x89\x70\x76\x0d\x99\x1f\xbc\x53\x50\xed\x94\xb0\x6a\xdc\x2a\x62\xe4\xd0\x15\xa8\x34\xd0\x02\xf5\x11\x71\x21\x7b\x8b\x62\xbe\x54\x43\x5f\x42\x81\xd6\x94\xdd\xae\xba\x86\xd2\x9c\x45\x95\x90\xf5\xf8\x73\xbd\xd5\x57\x6d\xd0\xa6\xa7\x1e\xc9\xf0\x35\x14\x72\x64\x73\xa0\xee\x28\x8a\xe5\xac\x98\xc8\xeb\xed\xb6\x6b\xc4\x19\xff\x48\xa9\x18\x93\xc5\x79\x9a\x33\x98\x82\xa2\x78\x8b\xb9\xb0\xa5\x01\x41\x2a\x5b\x68\xe2\x31\xb7\x52\x74\x9f\x35\x3e\x7f\x79\xda\x3b\x52\x12\x7c\x72\x02\x2c\xed\x10\x39\x23\x5a\x35\x70\xa7\x6a\x70\x80\x89\xfa\x55\x62\x1b\xce\x92\x9c\x44\x20\xa0\x9e\xb8\x2c\xaa\x95\x82\x45\xb7\xde\x28\x99\x69\x2c\xfa\x61\xd6\xea\x37\xd3\xf8\xbf\xa4\x59\x8d\x3b\x04\x19\xc5\x44\x20\x06\x04\x05\x10\x48\xf3\xab\x10\xf6\x43\x71\x3a\x49\x72\xf1\x0e\x72\x12\x0c\x2f\x53\xc4\x4b\x9a\x14\x1a\xab\x53\x00\xb3\x0c\x91\x38\xf0\x03\xbe\x2d\x96\x80\x86\x61\xb8\xb0\xd9\xf2\x58\x82\x32\x84\x3f\x57\xd0\x0c\x50\xee\x88\x49\x50\xf5\x2b\x04\x04\xdd\xe8\xd5\x8d\x1c\x0f\xc5\x07\x8d\x4b\x50\xae\x1f\x86\x61\x37\x4b\x46\x59\x45\x73\x71\x47\x4e\xc9\x23\xe3\xdf\x4b\xc9\x0a\x09\x48\xdb\xf9\x12\x2f\x6d\x9b\xca\x75\xaa\x65\x68\x2e\xd4\x84\x30\x18\xda\x2d\x0b\x0d\xbf\x70\xf4\x94\xe6\xc2\x88\x43\xed\xb7\x88\x92\x6b\xc4\x84\x2d\x0d\xa5\x89\xa4\x8f\xee\xdd\xd8\x2d\xff\xdb\xaf\x76\x0a\x13\x97\x9f\x1b\xf8\x0d\x05\x03\xc3\x97\x20\x45\x24\xa0\x8b\x9a\x85\x58\x4e\x7b\xf6\x0b\xc0\xe0\xef\xa6\xef\x17\x80\xff\xf6\x37\x97\x85\x9f\xf1\x17\x70\x0a\xe8\x67\xfc\x65\x90\x35\xaf\x50\x02\xf3\x54\x7c\x60\x31\x62\x8e\x99\x89\x75\x07\xa0\xb2\x07\x93\x35\x48\x30\x4a\x63\x5e\x6a\x6b\x44\x89\x40\x64\x07\xfe\xd8\x0b\x06\x0b\xf0\xf9\x8b\x76\x03\x1a\x36\xa6\x6c\xae\x49\xaa\x5c\x1b\xbd\x8a\x83\x77\xe9\x7c\x39\x5e\xd5\xd2\x9e\x6a\x4e\x11\xcd\x09\x4d\xf9\x27\x7a\x91\x41\xc6\x91\x43\xb5\xa7\xed\x36\xba\x84\x62\xa9\x41\x1a\x8c\xef\xf6\x3d\x39\x01\x1f\xda\x16\x1b\xdc\xe0\x34\x05\x94\xa4\xb7\x8a\xb3\xd0\x74\xad\xf1\x35\x22\x86\xf3\x21\x78\x4f\xf5\x8f\x60\x83\x20\xe1\x40\xea\x09\x43\xa6\x89\xa3\x1d\x64\x51\xb2\x20\x30\xb2\x0d\xc3\x50\xb3\xdd\xd7\x16\x28\xdd\x9d\x42\xff\x9d\x95\x39\x6c\xe0\x2c\x6d\x4b\x18\x3c\x1d\xc1\x01\x14\xc5\xb0\x85\x28\x5d\x61\x5b\x17\xae\x4d\xdb\x5d\x35\xde\xc0\x0e\x16\x00\x13\xd1\x75\x96\x22\x11\x3e\x3f\x3f\x3b\x23\x09\x0d\x2d\x97\x5c\xbb\xf3\x46\x71\xad\xe8\x40\xfd\x7c\x84\xf9\x6b\x12\xb1\xdb\x4c\x48\xb1\x48\x16\x26\x30\xe5\xc8\xc3\xe3\x6c\x7a\x9a\x1b\x39\x44\xd2\x08\x9b\xd3\x4a\x6f\x70\xdc\xb1\x31\xde\x7c\x1e\x49\xf2\x9a\x11\x50\x1d\xa9\x34\xe3\x8e\xe3\xa2\x30\x5e\x5d\x68\x88\x51\x7e\x5c\x07\x7d\xa7\x40\x30\xe5\x8e\x5b\x7c\xd8\x6e\x55\xf8\xf1\x89\xbe\x51\x1b\xe0\x48\xf2\xd1\x70\x21\x6c\xb2\x4c\x72\x5c\x21\x5d\x2e\x2d\x25\xf1\xf5\x77\x4e\xc9\x6a\x7e\x3c\x07\x1b\xbe\x96\x41\xaa\xfa\xf9\x52\x35\xfe\x5b\xb1\xc5\x68\xc0\xfc\xab\xd1\x92\xf7\xe8\x66\x84\xb5\xa5\xab\x23\x0f\xf7\xfe\x03\x4b\xe2\xa4\xb4\x68\x04\x60\xb0\x18\x06\xd2\x50\xa6\xc7\x43\x63\xeb\xfd\x64\x33\x62\x35\xa0\x81\xcb\x99\x65\x41\x8f\xed\xc0\x52\xf2\x5d\xea\x1c\xa7\xcc\x0a\x44\x41\x30\x22\xf0\x85\x63\xa8\x8d\xe8\xf9\x15\xcd\xd3\xf8\x5f\x0c\x0b\x74\x46\xb0\xc0\x30\xc5\xff\x8b\x98\x14\xa7\x8a\x54\xe5\x52\x3a\x47\xd0\x50\x9e\xa3\xf0\x3c\xbf\x4c\x71\x24\xa9\x01\x0d\xb0\x47\x98\x60\x65\x9f\x6e\x3a\xc0\x22\xe1\x00\x6f\xce\xc5\x89\x99\xee\xb4\x77\xb5\x1d\xdb\xa7\x8a\x5f\x93\xd9\xd5\x3e\x51\x63\xa7\xef\xdf\x17\x16\x96\x26\x69\x50\x5b\xf6\xe4\xe5\x03\xc7\xcd\x6f\x1a\x26\x5f\xc2\x12\xdc\xf0\x37\x74\xc0\xed\xd0\xf5\x84\x83\x9c\xe0\x3f\xf2\x32\xe6\x91\x73\x26\xd2\x2a\xa7\x04\x0b\xe0\xfa\x18\x3a\x4e\xd4\x73\x2d\x6c\x4a\xe5\x2c\x0f\x87\xb0\x5a\xa0\x1e\x14\xbe\x2c\x4f\xfe\x72\x1f\x57\x52\x96\x86\xa7\x01\x62\xde\x4a\xed\xec\xc2\xb0\x0b\x24\x2c\x2c\x39\x12\x87\x61\x98\xb3\x4c\x80\x63\x50\xba\x02\x7e\x5c\xf3\x63\x17\x38\x05\x38\x1e\x66\xca\xc9\x09\xf8\x15\x89\x17\x17\x1f\xde\x03\xbc\xc9\xb4\x9a\x6a\x8a\xa5\x69\x06\x1b\xc8\xf8\x15\x4c\xa5\x38\x55\x34\x99\xc0\x08\x29\xaf\xea\xd3\x15\xe6\x00\x73\x90\x73\xed\x96\x09\x06\x09\xcf\x20\x43\x44\x68\xaf\x4a\x22\x02\xce\x5e\xc9\xbe\x77\x94\xac\xe9\xab\x17\x67\xaf\x00\xe4\xe0\xc3\x25\x8a\xc4\xd9\x2b\x6f\x46\x19\xec\x82\x05\x08\x2a\x0c\x64\x9c\x83\x18\xa3\xac\x62\x17\x4e\x00\x05\xa7\xa7\x80\xe0\xd4\xf2\x65\x8c\x62\x10\x9c\x2e\xe5\x7f\x6c\x9f\x84\x4b\x83\xf5\x78\x23\x31\xab\x2d\xe8\xa0\x45\x2f\x95\xcf\xff\xb8\x9d\x59\x56\x2e\xbc\x10\x94\xa1\x18\x34\x5a\x2d\xd1\x9a\x1e\x49\x89\x12\x6e\x4b\x98\x3f\x9d\x82\xf9\xdc\xa2\x8e\x77\x0f\x3b\x55\x92\x0b\xb5\xdb\x7b\x16\xff\x17\xfa\x1e\x74\x03\x2c\x7d\x34\x67\x4f\x19\x2c\x7a\x61\x77\x83\x6a\xea\xd8\xf0\xaf\xf6\xa6\xe5\x5a\x32\x5a\x13\x2f\xee\xb5\x26\x1a\xec\x02\x06\x6f\x34\x8b\x3f\xc2\x9b\x85\xd6\x43\x5f\x35\xdc\x87\x06\xe2\x44\xae\xa9\x43\xfa\x9b\xf0\x9f\xc4\x30\x26\xe0\x8b\x5f\x54\xc7\x4f\x3d\xcb\x23\xc6\x1c\x81\x1f\x58\x8f\x7b\x94\xf8\xb4\x47\xb5\x42\xa9\xa7\x8b\x4e\x5d\x9c\x08\x69\x77\x5d\x34\x8a\xe8\x77\x4e\x74\xc5\x30\x57\x90\xc5\x11\x8d\x51\xdc\x8c\x66\x94\x7f\xeb\xad\x68\x3b\x87\x30\x0d\xc3\xfe\x22\x45\xd7\x48\x65\xda\x9b\x1b\x4a\x76\x84\x2f\x53\xc8\xb9\x96\xd9\x59\xbd\xa3\x3c\x71\xac\x60\xb7\xce\xfb\xf2\x34\xee\x8f\x6f\xe6\xfe\x5c\xee\x4d\x98\x94\x69\xe2\x9e\xc4\x89\x17\x1d\x8a\x90\x7b\x91\x21\x99\xe8\xa6\xd8\xce\x41\xd9\xcb\x23\x86\x33\x51\x5f\x30\xbd\xa2\x91\x9b\x61\xa2\x51\xae\x7c\x50\x35\x26\xa1\xcc\xf2\x64\x7c\x85\xfe\x8a\x46\x7d\xe2\xfe\x2a\x89\xe2\xd1\x0b\x18\x7d\x13\x38\xfa\xc6\x07\xb0\xfb\xea\x5c\x35\x4c\x24\xdd\xd7\x56\x2b\x1c\xfb\x90\x4d\x36\x22\xbc\xc8\x18\x26\x22\x09\xe6\x7f\xff\x99\xaf\x7e\xe6\xff\x98\x2f\x01\x0d\x6b\x97\x5d\x45\x41\x75\x93\xf6\x6c\x17\x2d\x51\x79\x1a\xd1\x4a\x66\x3a\xfe\xfa\x15\x11\xc4\xa0\x40\xbf\x22\x21\x10\x03\x61\x2b\xbc\xd2\x5e\x59\xa7\xd9\x6b\xe6\xcf\x5a\x03\x8c\xc9\x61\x28\x42\xf8\xba\xe9\x91\x1e\x0d\x7b\x5a\x5d\x00\x83\x85\xbb\x4e\x79\x75\xe7\xb2\xb4\xc7\x2f\x68\xa4\x55\xda\x2c\xb8\x18\x60\xc1\x45\x0f\x0b\x2a\xa7\x3c\x63\x34\x43\x4c\x06\x53\x1e\x8c\x00\x39\x97\x9a\x50\x27\xfa\x94\x4b\xef\xcf\x9e\x1e\x6c\x54\x6a\xfe\x7d\x7d\xbf\xd9\x62\x54\xe5\xa3\xf6\x9e\x63\x16\x84\xc6\xd6\x68\xee\x0c\x48\x62\x10\xf4\x6d\x8f\x45\xbb\x4b\xdf\xc6\x2d\x0c\x3f\x3b\x73\xb0\x5c\x37\x75\x1f\x58\xca\xbd\x2a\xc7\xa3\xb8\x4c\xe4\xef\x35\x7d\x3a\xb2\x93\xbd\xb2\xa6\x7a\x88\x9d\x3b\xb5\x3c\xb2\x14\x11\x33\x79\x21\x7d\xb3\x67\x96\x6b\x74\x72\x02\x08\x4d\x31\x11\x2b\xb0\xa6\xfa\x49\x04\x6f\xfa\x4d\x8f\xc7\xb3\x9d\x16\x44\xd0\xf1\x2a\x61\xd0\x2e\x34\x27\xe2\x04\x5c\x41\x7e\xce\x50\x82\xbf\x83\x40\xe7\xdc\x94\x26\x39\x29\xb7\x05\x98\x3f\x9d\xb7\xa7\xb7\xd5\x6b\xd5\xa3\x76\xcb\xd6\xc2\xb6\xcb\x35\x0c\xf1\xb1\x37\x48\x37\x3d\xd3\xd3\x5c\x38\x5e\x71\xa6\xdc\x62\x1f\x9e\x17\xf6\x2d\x57\x52\xdf\x71\x19\x45\xb1\x02\xa5\x1b\x2c\xa2\x2b\x90\xec\x4b\x4c\x11\xe4\xfa\x09\x46\xb9\x6b\xe7\x2b\xa7\x7f\x0f\xa2\xd4\xac\x68\xb3\x79\x2c\x06\xf3\x11\xea\x20\xec\xc7\x83\xc1\xe2\x1e\x04\x5c\xc6\x7d\x99\xe5\x04\x36\xf2\xcf\xda\x5a\x99\x16\x4b\x2a\x48\xb7\x68\xbb\x05\xeb\xf6\x0d\x64\xdf\x50\x2c\x43\xba\xaf\xa8\xcc\x6c\x7f\x6d\x99\xfb\xb2\xcb\x3f\x47\xd3\xc2\x20\xa8\x60\x58\xb6\xa7\xea\x2e\xb3\xea\x6c\x01\x02\x19\x88\x55\x19\x0a\x30\x25\xde\x6a\x04\x56\x76\xaa\x7e\x20\x3f\x50\xe8\x8c\x08\x38\xb5\xc8\x34\x53\x8d\x2b\xd4\x39\x69\x24\x64\x94\x7e\xd2\x6b\x49\x45\x12\xcc\x73\xa2\x64\x23\x68\xb9\x82\xf5\x1c\xe9\x49\x07\xe8\x27\x6a\x67\x3e\x19\x3c\x54\x9f\x80\xe0\x67\xbe\x58\x81\x9f\xf9\xbc\xe9\x6a\x29\x72\x5a\x19\x0a\xdf\x10\x4e\x45\x0e\x4d\xf5\x89\xd1\x5d\xd4\xc7\xcc\x9e\xa0\x3e\x2d\x0c\xfe\x5a\xea\x63\xd0\xdf\xbb\xfa\x18\x46\xde\x67\xf5\xb1\x7b\xa5\x2e\x9d\x43\x79\x7e\xc0\x2c\x4b\xf5\x23\x1a\x42\xd5\xd0\x3a\x29\x0c\x81\xc7\x9d\x68\xf9\x18\x65\xe2\x35\x82\x5a\x3c\x30\x6e\xda\x90\xcb\x63\xe5\x8e\x8f\xf5\x6d\x0b\xaf\xdf\x31\xfa\x6a\x8c\x99\x57\x6b\xcb\x4f\x7a\x65\x3b\x34\x3a\xe3\xaf\xff\xc8\x61\x1a\xd8\xf1\xd2\xc2\x12\x7f\x06\x09\x8e\x82\x79\x04\x89\xf4\x47\x33\xc5\xbc\x84\xd1\x0d\x80\x40\x53\x71\x83\xc5\x15\x88\x71\x92\x20\x86\x88\xa8\xde\x58\xcd\x9d\x5b\x63\x4e\xd5\xa5\x97\x5e\xdd\xef\xce\x79\xe6\x1e\xeb\x2d\x5a\x78\x7f\x66\xd5\x55\xe0\x3b\x9c\xde\xfd\xd9\xaa\x91\x63\xbb\xeb\xb8\xee\x05\xf6\xd4\x0b\x9a\x9d\x64\x68\xb5\x0d\x5d\x09\xbc\x42\x28\x6b\x3c\x27\x8b\x11\xca\xf4\x0b\x2a\x3c\xf2\x82\x4a\xbd\xfc\xf4\xb6\x91\x7a\x21\xdf\xbb\x57\x37\xc1\xfa\xe8\x91\xb5\x6f\x1f\x15\xb3\xd9\x23\xf3\x52\x62\xf8\x6e\xb6\x98\x3d\xa2\x61\xb9\xf2\x19\x11\x34\xa0\xb9\x58\xcc\x66\x8f\x3a\xde\xeb\xd4\x83\x24\xf1\x18\x71\x37\xa6\xc4\xc4\x6c\x6a\x7d\x48\x0c\xd2\x30\x99\x29\x25\x6a\x63\xe3\xb7\xb3\xd9\x23\x01\xd9\x1a\x89\x65\x99\x1a\x76\x9e\x5b\x87\x8a\xc3\x74\x31\x7b\x64\x72\xc7\x3f\xd5\x0c\xd4\x7b\xd5\x49\x88\xfc\xd3\x32\xd5\x28\x53\x22\x1f\x5a\xdf\xd8\x5f\x69\x6f\x17\x5a\x08\x4f\xf5\x9b\xb2\xa7\x1a\xa7\xa1\xb7\x64\x6a\xd7\x9a\x37\x21\xfa\xe9\xb6\xba\x69\xc3\xb1\xe1\x73\x94\x33\x6d\x21\x48\x42\xd9\x46\xa7\xae\xb8\x4e\x40\x57\x9c\xaf\xc9\xf4\xce\xaf\x9a\xa5\x82\x46\xf6\x5e\xfd\xa2\x6c\x66\x6d\x66\xd5\xf9\xc5\xb7\xe5\x45\xe3\x1f\x39\x66\x28\x7e\x3d\x34\x70\xc7\xf3\xba\x4a\x7c\x7d\x62\x90\x70\x2c\xa9\x76\xfa\xc2\xd7\xdf\x33\xca\x51\x7d\x64\x99\xe6\x8f\x06\x27\x77\x34\xfa\x03\xe8\x37\xd6\x73\x1d\x2c\xcf\x2d\x2b\x68\x54\xa4\x46\xbd\xe4\x47\x09\xca\x1c\xf9\x4e\x84\xb3\xec\xb1\x45\xfd\x2e\x80\xc3\xaa\xd3\x46\x43\x68\xde\x49\xb6\x4e\x69\xe7\x54\xd6\xb4\x50\x06\x82\x9a\x1e\x44\xf2\xcd\x7c\x71\x77\x72\x78\xaf\x5f\x23\xa9\xfa\x01\x64\xd5\x24\x09\xbc\x41\x93\x04\xf4\x09\x6f\xd0\x7d\x15\xcf\x77\x81\x18\x81\xe9\x7c\x61\xb7\xa6\x98\x8b\xf9\x62\x02\x85\xaf\x0d\x98\x7b\x43\x65\x4d\x0b\x26\x02\xad\x11\x9b\x24\xb0\x33\x22\xee\x21\x25\x49\x4a\xa1\x98\x44\xc7\x1b\x39\xe3\x7e\x50\x32\x44\x18\x43\xc9\xdc\xeb\x3e\xdd\x45\xad\xa6\xfb\x23\xe2\x48\x98\x2b\x9d\x37\x94\xfd\x0f\x62\x54\x7f\x0a\x33\x9a\x1d\xa9\xb9\xd8\x3d\x32\xac\x0f\x9f\x1e\x06\x59\x27\xd1\xa9\xf9\xa1\xc5\x14\x60\x65\x55\x3c\x37\x26\x43\xc9\x5b\xb5\x0b\x1b\x8d\xef\x60\x56\x9b\x53\x93\x4c\xe3\xf9\xa5\xf5\x64\xbc\x9b\x7b\x5b\x9b\x64\x39\xa1\x75\xed\x0d\xd4\x83\x7f\x22\x30\xc9\x51\x03\x6b\x5f\x6e\xf3\xfc\xb2\x8b\xb5\x3c\xbf\xfc\x81\x7c\x0c\x9f\xa7\x29\xbd\x41\xf1\xcb\x2b\x8a\x23\xc4\x7d\xf6\x8b\x3e\x72\xce\x88\x7a\x9e\x3e\xe9\xe0\x59\xd6\x57\x8d\x72\xde\xef\x14\x93\x16\x02\x5f\xe7\x4b\x30\xff\x2a\xa1\x15\x4b\xe5\x9a\x3d\xcf\x05\x5d\x9b\x0b\x95\x78\x60\xef\x8d\xb1\xc3\x8b\x09\x90\x79\xb1\xe0\x1c\x0a\x69\xc2\xfd\x8c\xc5\x52\xdd\x1f\x36\xd7\xf8\xda\xd1\xfc\x0e\x71\x0e\xd7\x48\xf7\xca\x4e\xcb\xff\x39\x00\xd9\x6b\x01\xc2\x77\xf0\xfb\x5b\x44\xd6\xe2\x0a\x3c\xf3\x21\xfc\x1d\xfc\x8e\x37\xf9\x46\x4f\xf1\x25\x5f\xb6\xd6\xeb\xc8\x16\x15\x5f\x1e\x8a\x22\x4c\x26\x51\x84\xc9\x8e\x14\x55\xeb\x1c\x9c\x22\xf8\x5d\x99\x0c\xf0\x2c\x7c\xd6\xe7\x09\xfb\x1f\x77\x46\x84\x13\x4e\xbb\x4a\x82\xbf\x99\x2f\x19\xf7\x46\xae\xc9\x08\xf8\xe2\xec\xed\x69\x2c\x65\x04\x15\x34\xb0\x5e\xec\x59\x4a\x63\x5a\xb8\x4f\x99\x69\x25\x9d\x2e\xb3\x12\x8b\xfd\xcb\xcc\x13\xe5\x5d\x44\x56\x23\xfd\x63\x44\x56\x3d\x42\x0f\x41\xeb\xd3\x6b\xf5\x48\x5d\x7d\xfa\x3c\xf4\xfd\x75\xcd\x08\x09\x6e\x53\x12\xab\x28\xb7\xde\x9d\xd7\xe4\xeb\x46\x5f\xb7\xd2\x97\xcc\xe3\x1e\x3f\xd2\x83\x09\x0e\xbd\xd7\x15\xa5\x0a\xb1\x2a\xcf\xaa\x13\x0e\x35\x1f\xec\x8f\xa7\x5f\xe6\x5c\xd0\x4d\x79\x89\x5e\x43\x08\xeb\xac\xed\x06\x66\x19\x26\x6b\xf5\x05\xb6\x7a\xe7\x55\x43\x7a\xa7\xbb\x42\xf3\x7f\x30\xaf\x3f\xce\x6f\xa1\xd3\xc8\xe9\x96\x50\xbb\x45\x61\xe0\x96\x02\xa1\x7b\x63\x71\x17\xc7\xcd\x7d\xbc\xeb\xf4\x2f\xc0\x3f\xac\x6b\x79\x93\x85\x73\x87\x98\x15\x6c\x18\xa8\x63\xae\xfd\xda\xb1\x31\x6b\x97\x47\x7e\xb2\x07\x27\x38\x52\x9c\x7d\x43\x59\x95\xc7\x71\x9e\x50\x54\xad\xce\xf0\xea\x8d\x95\x4e\x0d\xd6\xd7\x1d\xea\x63\xf8\x6f\xe8\xb6\xcc\x57\x8d\x3d\x65\xea\xc3\x21\x50\x80\xda\x8f\x21\x7a\xd0\xa9\x33\xa8\xd7\x4b\x40\xbf\x19\xf1\xf7\x2e\x5c\xa7\xac\xde\xc1\xec\xb3\x5c\xea\xcb\x2f\x72\x5a\x8b\xd3\xd7\x36\x93\x4f\x4e\xc0\xbf\x10\x88\x68\x9e\xc6\x8a\xb7\x09\x26\x31\xc0\x62\x09\x38\x05\x29\x12\x4f\x38\x88\xae\x50\xf4\x0d\x50\xf3\x31\x1e\xbd\x41\x4c\xdf\xa7\x63\x12\xa3\xef\x28\x06\x3c\x43\x11\xd8\xc0\xcc\x96\xd9\x10\x9e\x6f\x25\x88\x97\x90\xa3\x0e\x84\xcb\xef\x83\x3b\x19\xc2\x1d\x19\x26\x79\x9a\x5a\x32\xe2\xee\xc8\x0d\xcc\x3c\xa5\xd5\xb3\x56\xb0\x90\x30\x3e\x6b\x61\x7d\xf1\x95\x95\x07\xf9\x0e\xd5\x75\x2a\x35\x47\xbd\xda\xaa\x2f\xad\x7a\x94\xd3\x79\x51\x0d\xc1\x35\x62\xb7\x00\xc6\xd7\x90\x44\x28\x06\x92\x01\x0a\x3d\x71\x05\x05\xb8\xa5\xb9\x79\xcb\xa5\x24\x4d\x10\x8a\xc1\x65\x2e\x00\x26\x80\xd3\x0d\x92\x80\xd4\xf4\x92\x95\x20\xe7\x48\x89\xda\xef\x71\xa6\x49\xd4\xba\x84\xb8\x2a\x6f\x7d\x0f\x50\x72\xcc\x3c\xf5\x50\xc3\xb6\x0d\xbb\xed\x99\x8a\x1d\x7a\xdd\x31\xfc\xd8\xad\xc3\xfc\xf5\xdd\x4e\x7b\x8b\xb4\xf5\x01\x21\xcc\xd4\x85\x63\x25\x5a\x29\xc8\xe1\x5b\x87\xb1\x2a\x16\xee\x7a\xa7\x53\x14\x75\xdb\x3c\x1b\x27\xa5\xbb\xad\x8f\xd1\xaa\x19\x12\x85\xc6\x63\xc0\x63\xbb\x82\x4b\x93\xe9\xf3\x55\xfb\xda\xae\x33\x56\x95\xff\xec\xf6\x55\x3b\xb6\x94\x31\xa5\x03\xab\xf1\x8e\xc5\x0d\xc3\x57\x3d\xe9\x81\xe3\xa2\x98\x14\xc2\xd7\x0e\x63\x35\xad\xa8\x7c\x8f\x65\x9b\xb6\x46\xac\x5f\x63\x67\x77\xac\x3a\xf3\x02\x83\xd4\x95\x0b\xf4\x5e\x29\x3a\x1d\xab\x1e\x71\xf8\x2d\xc1\x90\x52\x9f\x0f\x24\xbd\x75\x56\xb0\xda\x35\x05\x8d\x91\x5e\xd0\x4d\x26\xa9\x74\x90\xab\x7e\xbb\x5d\x41\x57\xdf\x2b\x3a\xa3\x9d\x0f\x16\x43\xff\x05\x33\x86\xa2\xa6\x3c\xea\x56\x4d\x8a\x33\xca\x13\xae\xf3\xa8\xbb\x06\x5c\x35\xaf\x3a\xde\x5d\x37\x1e\x5b\x7b\xad\xd4\x7a\x26\x22\xff\x55\x8d\x1a\x7f\x7b\x8c\x1f\xd0\xfa\xc2\xaa\x02\xa9\x9b\x0c\xc0\xaa\xdf\x0b\xdc\x1b\x9c\x0a\xc4\xca\x07\x62\x65\x6f\xdd\xaa\x81\x3a\xa3\xfc\xe0\x52\x86\xf0\x9a\xfc\x37\x72\x54\xb1\x6e\x35\x70\xed\x51\x5e\x70\xcd\xe3\x6e\xab\x47\xb7\x68\x78\x55\xaf\x17\xac\xf6\xe7\x35\xf2\x5f\xdd\xaa\x61\x3a\xa3\xbc\xe0\xda\x29\xa7\xaa\xb3\x6a\x5c\xb5\xd3\x52\x9e\x40\x5b\x7b\xaf\x6c\x5b\xb5\xf2\x24\x5e\x10\xad\x3c\x52\x0d\xb2\x6c\x5c\xb5\x73\x4d\x9e\x40\xdb\x68\x9a\xb6\x55\x2b\x35\xe0\x03\xb1\x69\x30\x2d\x3b\x39\xc9\x3c\xaa\x8f\x48\x9a\x8a\x5e\x35\x6a\xdc\xec\x31\x5e\x40\xcf\x19\xde\x40\x76\xdb\x50\xf3\xba\x55\x83\x75\x46\x79\xc1\xfd\x88\x60\xdc\xb4\xe3\x65\xdb\xca\x64\x68\xab\x11\x9e\x10\xdd\x1b\x6d\x0d\x51\xb7\xad\x9a\x39\x5f\x2f\x88\x17\x28\x62\xc8\xf9\x9e\x5c\xb7\x94\xdf\xe3\x9b\x5e\x4f\x58\xcd\x6d\x7d\x61\x6d\xeb\x8b\x49\xdb\xfa\x02\xaf\x89\x4b\xa7\x6e\x31\xb0\xca\x5e\x3f\x58\xce\x77\x79\x0a\x96\x6a\x31\xb0\xca\x5e\x3f\x58\xf9\xa5\xf9\x78\xa1\x06\xa6\x9b\xca\xd2\x64\xd5\x00\x3f\x8d\x6e\xbd\x69\x90\xff\xaa\x46\x8d\xa2\x3d\xc6\x0f\x68\x03\x45\x0b\xbf\x51\xe4\xcc\x0a\xfd\xc9\x8a\x71\x37\xbd\x3b\xf0\xfc\x01\xfe\x7a\xcf\xc2\xf7\xdb\x71\xd7\x91\x7e\xf8\xe0\xb5\x1f\xc8\x6b\x6f\x6f\xff\x17\x17\x1f\xde\xab\x62\x2d\xf5\x31\x24\xb5\xd0\x69\x6e\x7f\x99\xf8\x10\x16\x3c\x84\x05\x0f\x61\xc1\x43\x58\xf0\x10\x16\x3c\x84\x05\x0f\x61\xc1\x43\x58\x00\x1e\xc2\x02\x37\x2c\xd8\xb6\xeb\x15\xdc\xe1\xab\x6c\x7d\xc9\xe8\x5f\xd1\xb1\xbb\x56\xb2\x2f\x04\x5d\x58\x70\xd2\x7a\x9f\xbf\x8c\x7d\x79\xb3\xb7\xea\xc9\x53\xf0\xfa\x53\x6b\x28\x4f\x2b\x41\xba\x1b\x79\x3b\xd5\x53\x9e\xb2\xc4\x41\xaa\x2a\xff\x08\xce\x1c\xa8\xc2\xf2\xee\xbc\xbb\x5b\x9d\xe5\xd1\xdd\xf5\x03\xaa\x2d\x4f\x13\xc0\x7f\x6a\xcd\xe5\x69\x5c\xba\x17\x75\x85\x74\xd5\x8f\xf3\x14\x62\xb7\x50\xd5\xa4\x23\xc0\x29\xbf\x7c\xd8\xbd\x6d\x70\xbd\x5f\xda\x15\x56\x58\x0d\xea\xd9\x2e\x45\x8d\xa7\x71\x67\xf7\xd2\xc6\xe3\x5e\x46\x47\xad\xe2\x76\x65\x98\xa1\xa2\xc5\xa1\x97\x73\xe1\xd4\x2e\xee\x50\x77\x9f\x3a\x74\x1d\x85\x88\x55\xbd\xe2\xb0\xc3\x5f\x53\xbf\xef\xa7\x16\xb1\x17\x03\xed\x8a\xc4\x1e\xbc\x08\xab\xc2\xc4\xe3\x63\x83\x85\xd7\x47\xe6\x5b\xe7\x3c\x1f\x9f\xb0\x2d\x15\xc4\xab\x04\xae\xd1\x89\xd6\x27\xec\x1e\x5f\x67\x1f\xa2\x1e\xae\x6f\x85\x5b\x07\xed\x91\xba\xad\xfe\xa4\xec\x52\xee\xb6\xfa\x1c\x69\xbc\x76\x6b\x4f\x55\x05\x53\xe6\xb6\xb0\xd9\xf5\xf4\x40\xe5\x73\x7d\x0b\xe2\xee\x9b\xbf\x77\xac\x8e\x8b\x13\x80\xe3\x56\xbd\x54\xef\x9a\xb9\x8f\x4d\xd1\xdc\x42\xf3\x69\x07\x08\x75\xa9\xcf\x4e\x96\xfe\xd9\x15\x77\x7d\x44\xf0\x23\xeb\xee\x7a\x59\xa9\x1d\x4f\x8b\x3d\x56\xdf\x75\x89\xf4\x2f\xbf\xfb\x74\x7a\xfd\xdd\xbe\x32\x2d\x2e\x1a\xbb\x54\xe9\x6d\xa9\xe4\xf0\xaf\xb6\x85\xb8\x57\xb5\x7a\x3d\xcd\xc8\xc1\x2b\xf6\x7a\xeb\xee\x5f\xa4\x6e\x2f\x8e\xd5\x17\x96\x63\x45\x7a\x7b\x2b\x93\x3c\xb6\x2b\x8e\xbb\x1a\xdd\x01\xb4\xa9\xd1\xbb\xd4\xfa\xbd\x83\x46\xd7\xea\x7c\xa7\x72\xbe\x3e\xba\x78\x27\xe7\xbd\x2b\x86\xcb\x54\x4b\x03\x33\x73\x10\xef\x82\x60\x57\xec\xa5\x5a\xba\xca\x2a\x9a\xd8\x6b\xf4\x2f\x6a\x54\xcc\x9f\xf4\xca\x60\xcc\xfa\xed\xa9\x52\x4f\x2e\x76\xb0\x9f\x5d\xba\x3d\x02\xad\xc7\xfe\xf7\x26\xa1\x87\x34\xd6\xc4\x9b\xdb\x3f\xb9\x70\x9e\x8f\x42\xdd\xd7\xf2\x79\x5d\xe2\x18\xaf\x9f\xd7\x3d\xeb\xf0\x05\xf4\x2a\x4e\xff\xff\x2d\xa3\xe7\xa3\x4c\xf7\xb5\x98\x9e\xaf\x32\xb9\xd5\xf4\xf6\xa8\x4c\x93\xca\xe9\xfd\x89\xca\xb4\xfd\xeb\xd6\xcd\xf6\xe0\xda\x50\xf5\x6c\x55\xc2\x6b\x3b\x54\x19\xba\x3f\xea\xe8\xf1\x4c\x3b\x12\x91\x7d\x87\xcc\xfd\xac\xc0\x0d\xe3\x98\x21\x5e\xe5\xda\xbb\x0b\x72\x7b\xf1\xfd\x70\x65\xb9\x1f\x6f\xc7\xeb\x72\x7b\xd6\xca\xf3\xce\x04\xfa\x9b\xc3\xba\x6e\x9e\x57\x52\xb0\x23\xd8\xe9\xaf\x9e\xe7\x15\xd3\x1c\xac\x86\xde\xc1\x98\x55\xd7\xd3\xf3\x99\x75\xf8\xaa\x7a\xe3\x58\x78\xd4\xd6\xf3\xa9\x8b\xe9\xa8\xac\x4a\xd3\x4f\xf8\xb3\x33\xcd\x7c\xfd\x84\xbf\x37\x38\x18\x80\x9a\x34\xbe\x0a\xf3\x7b\xfe\xa8\xe0\xf0\x19\xb3\x97\x07\x18\x5d\xdc\x98\x7e\x87\x71\x68\x9e\xf4\xde\x6f\xb4\x58\x62\xfd\xf2\x7f\x01\x00\x00\xff\xff\x44\x61\x87\xfc\x0a\x80\x00\x00")

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

	info := bindataFileInfo{name: "templates/model.gotpl", size: 32778, mode: os.FileMode(420), modTime: time.Unix(1593566059, 0)}
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

	info := bindataFileInfo{name: "templates/relationships_registry.gotpl", size: 17509, mode: os.FileMode(420), modTime: time.Unix(1586905041, 0)}
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

