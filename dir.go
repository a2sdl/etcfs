package etcfs

import (
	"os"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
	"golang.org/x/net/context"
)

type Dir struct{}

func (_ Dir) Attr(ctx context.Context, attr *fuse.Attr) error {
	attr.Inode = 1
	attr.Mode = os.ModeDir | 0555
	return nil
}

func (_ Dir) Lookup(ctx context.Context, name string) (fs.Node, error) {
	if name != "nginx.conf" {
		return nil, fuse.ENOENT
	}
	if kv, err := NewKVClient(); err != nil {
		return nil, err
	} else {
		f := File{
			KV: kv,
		}
		return f, nil
	}
}

func (_ Dir) ReadDirAll(ctx context.Context) ([]fuse.Dirent, error) {
	return []fuse.Dirent{
		{
			Inode: 2,
			Name:  "nginx.conf",
			Type:  fuse.DT_File,
		},
	}, nil
}
