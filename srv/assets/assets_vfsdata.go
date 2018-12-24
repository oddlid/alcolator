// Code generated by vfsgen; DO NOT EDIT.

// +build !dev

package assets

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// Assets statically implements the virtual filesystem provided to vfsgen.
var Assets = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Time{},
		},
		"/static": &vfsgen۰DirInfo{
			name:    "static",
			modTime: time.Date(2018, 12, 24, 0, 21, 30, 336003127, time.UTC),
		},
		"/static/alc.css": &vfsgen۰CompressedFileInfo{
			name:             "alc.css",
			modTime:          time.Date(2018, 12, 24, 0, 21, 30, 336003127, time.UTC),
			uncompressedSize: 824,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x92\xdb\x6e\xb3\x30\x10\x84\xaf\xe3\xa7\xb0\xc4\x35\x08\xfe\x5f\xa5\x8a\x79\x9a\x85\x5d\x1b\xab\xc6\x46\xb6\x49\x4b\x2b\xde\xbd\xe2\xa0\x34\x09\x3d\xa8\x87\x4b\xb0\xf7\x9b\xd9\x19\xd7\x0e\x47\xfe\xc2\x0e\xd2\xd9\x98\x4a\xe8\xb4\x19\x05\x0f\x60\x43\x1a\xc8\x6b\x59\x6d\x27\x41\x3f\x93\xe0\x45\xf6\x8f\xba\x8a\x1d\x6a\x68\x1e\x94\x77\x83\x45\xc1\x93\x3c\xcf\x2b\x36\x31\xd4\xa7\xa4\x71\x36\x92\x8d\x33\xef\xea\xca\xff\xf2\x38\x4f\x39\x8f\xe4\x53\x0f\xa8\x87\x20\xf8\x8a\xea\x01\x51\x5b\xb5\x7d\x4e\xac\x2d\x76\xd3\xe5\xb1\xa9\xd8\xa1\x71\xc6\x79\xc1\x13\x22\xda\xb3\x8a\x6b\x56\x9e\xdd\xad\x34\xd4\xa7\x4c\x3a\xdf\xed\x98\x44\xf5\x97\x8e\x8a\x1b\x46\x84\xda\xd0\x4c\xea\xc0\x2b\x6d\x05\xcf\x39\x0c\xd1\xdd\xe6\xf1\x13\xbb\xe5\x9b\x54\x0f\x6a\x11\x41\x1d\x7a\x03\xa3\x58\x65\xab\xbd\xea\xc4\x32\x30\x5e\xab\x76\xc9\x3b\xd2\x53\x4c\xc1\x68\x65\x05\x5f\x7e\x5e\x78\x90\x12\xef\xcf\x25\x65\x9e\xc2\x60\xe2\xa7\x12\x05\x75\xef\xae\xf6\x8d\xd4\x42\x0f\x76\x93\x4a\x5b\xf4\xe7\x27\xf6\x48\xb3\x39\xc1\x6b\x67\xf0\x2f\x92\xbb\x28\x7a\x5b\x6c\x57\xd3\x47\xdb\xfc\xa2\x28\xc6\x5e\x03\x00\x00\xff\xff\xa2\xde\xbc\xa3\x38\x03\x00\x00"),
		},
		"/templates": &vfsgen۰DirInfo{
			name:    "templates",
			modTime: time.Date(2018, 12, 24, 0, 22, 33, 756843530, time.UTC),
		},
		"/templates/apkform.html": &vfsgen۰CompressedFileInfo{
			name:             "apkform.html",
			modTime:          time.Date(2018, 12, 24, 0, 22, 33, 756843530, time.UTC),
			uncompressedSize: 2338,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x96\xdf\x6f\xdb\x20\x10\xc7\xdf\xfb\x57\x5c\x91\x2a\x6d\x0f\x0e\x6a\x9f\xa6\x0e\x47\xaa\xd2\x49\x7b\xd8\x8f\x6a\xaa\x36\xed\x91\xe0\x73\x8c\x86\x8d\x07\xe7\x68\x51\x94\xff\x7d\x02\xe3\x36\x76\xb2\xad\x5d\xdb\xa7\x5c\x38\x38\x3e\x77\xc7\x17\x23\x4e\xaf\x3f\x2f\x6e\xbf\xdf\xbc\x83\x8a\x6a\x33\x3f\x11\xfd\x0f\x80\xa8\x50\x16\xc1\x00\x10\x35\x92\x84\x8a\xa8\xcd\xf0\x67\xa7\xd7\x39\x5b\xd8\x86\xb0\xa1\xec\x76\xd3\x22\x03\xd5\xff\xcb\x19\xe1\x2f\xe2\x21\xc0\x5b\x50\x95\x74\x1e\x29\xef\xa8\xcc\xde\xb0\x14\x87\x34\x19\x9c\x5f\x19\x65\x2b\x6b\x60\x21\x8d\xea\x8c\x24\xeb\x04\xef\x3d\xfd\x2c\xa3\x9b\x1f\xe0\xd0\xe4\xcc\xd3\xc6\xa0\xaf\x10\x89\x01\x6d\x5a\x4c\x3b\x28\xef\x19\xd4\x58\x68\x99\x33\xaf\x1c\x62\xc3\xa0\x72\x58\xe6\x8c\x7b\x92\xa4\x15\x97\x46\xcd\xe2\x2c\x1e\x53\xe1\x43\x2e\x62\x69\x8b\x4d\xda\xa6\xd0\x6b\xd0\x45\xce\x12\x7d\x62\x4c\x0e\x65\xa4\xf7\x39\x6b\xe5\x0a\xef\x1c\x63\x57\x69\x5d\xbd\xe7\x0a\x05\x3b\x9f\x0f\x29\x21\xc8\x94\xa4\xa7\xae\x2c\x05\xaf\xce\x47\x53\xc3\x62\xa8\x91\x2a\x5b\xe4\x6c\x85\x34\x0a\x14\x0a\x25\x97\x43\x39\xf6\x47\xdd\x74\x28\x0c\x16\x73\x61\xe4\x12\xcd\xfc\xda\x85\xc2\x35\xb2\xc6\x4b\xc1\xfb\x21\xc1\xa9\xf8\xc3\x1a\xdd\xb4\x1d\xed\x55\x95\xc5\x95\x39\x2b\xc2\x0f\x8b\x95\x49\xe6\x5a\x9a\x0e\x73\xb6\xdd\xc2\x2c\x6e\xf1\x49\xd6\x08\xbb\x1d\x3b\x1a\x5d\xf0\x43\xca\x87\x81\xaf\xad\xe9\x6a\x84\x57\xb5\x79\xfd\x24\xfe\xb5\x35\xb5\x49\x09\x24\x7b\x9a\xc1\x57\x6b\x3e\x7e\x78\x81\x14\x5a\x74\x0a\x1b\x92\xab\xa7\x75\xa0\x55\x94\xf8\xa3\x35\xa5\xbf\x59\xdc\xbe\x04\xbb\xd3\xea\x89\xd8\x21\xc2\x00\xde\xdb\x07\xe8\x61\xf8\x59\xe0\x41\x59\xe3\x5b\xd9\xe4\xec\x82\x0d\x92\x94\xc6\xe9\x55\x45\x6c\xcc\xe8\xbb\x65\xad\xef\x29\x7d\xb7\x4c\x8c\xd1\x4a\x84\x41\xb8\x0f\x85\x12\xfc\x40\x9e\x82\x07\x45\xef\x5d\x14\xbc\xd0\xeb\x39\x88\xd3\x2c\x83\xa8\xf5\x2c\xbb\x77\x6e\xb7\xa0\x4b\x98\xbd\x97\xfe\x5a\x92\x84\xdd\xee\xe8\xfd\xe2\xd0\x77\x66\x7c\x31\x88\x90\xf0\xd8\x9f\x55\x85\x9b\x5c\x1e\x07\xa7\x1c\x6a\x03\xb6\x84\xa9\x7e\x03\x18\x1c\x74\x06\x24\xc1\xe4\xa4\x9d\x81\xac\x6d\xd7\x90\x07\xb2\x97\xa3\xa4\x03\xcf\x08\xf0\xc8\xbd\x75\xa4\x81\xe1\x18\x7d\x93\x84\xee\xf2\x68\xc1\x43\x73\x27\x0d\x0d\x44\x5f\x62\xc2\x71\xdd\x90\xd6\xe1\xf2\x23\xdd\x3a\xbe\x7f\xfa\x08\xfd\x07\xc1\x95\x51\xcf\xb0\x7f\x5f\xef\x16\x5d\xe8\x4e\x11\x8a\xfd\x18\x94\xd6\xe9\x86\x4a\x60\x67\xb3\x8b\x92\x0d\x60\x31\xe4\x0d\xba\xd8\xba\x88\xf8\x5c\x7c\xf2\xf1\xc5\xfa\x2b\xe1\x50\xc2\x7f\xf3\x1d\x68\xad\x57\x56\x14\x56\xaf\x80\xa9\xb4\xb0\x29\xee\x25\xb5\xaf\xc3\xf0\x2d\xbf\x9b\xbc\xef\x48\x0f\x80\xe4\x13\xbc\x7f\x23\x08\xde\xbf\x84\x4e\x7e\x07\x00\x00\xff\xff\x69\xe6\x73\xba\x22\x09\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/static"].(os.FileInfo),
		fs["/templates"].(os.FileInfo),
	}
	fs["/static"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/static/alc.css"].(os.FileInfo),
	}
	fs["/templates"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/templates/apkform.html"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}
