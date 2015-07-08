package etcfs

import (
	"bazil.org/fuse"
	"golang.org/x/net/context"
	"os"
	"bazil.org/fuse/fs"
)

type Dir struct{}

func (_ Dir) Attr(ctx context.Context, attr *fuse.Attr) error {
	attr.Inode = 1
	attr.Mode = os.ModeDir | 0555
	return nil
}

func (_ Dir) Lookup(ctx context.Context, name string) (fs.Node, error) {
	return File{}, nil
}

func (_ Dir) ReadDirAll(ctx context.Context) ([]fuse.Dirent, error) {
	return []fuse.Dirent{}, nil
}
