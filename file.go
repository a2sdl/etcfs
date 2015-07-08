package etcfs

import (
	"bazil.org/fuse"
	"golang.org/x/net/context"
)

type File struct{}

func (_ File) Attr(ctx context.Context, attr *fuse.Attr) error {
	attr.Inode = 2
	attr.Mode = 0444
	attr.Size = 4
	return nil
}

func (_ File) ReadAll(ctx context.Context) ([]byte, error) {
	return []byte(""), nil
}
