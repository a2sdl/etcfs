package etcfs

import (
	"bazil.org/fuse"
	"golang.org/x/net/context"
)

type File struct{
	KV *KVClient
}

const fileContents = "nginx config\n"

func (f File) Attr(ctx context.Context, attr *fuse.Attr) error {
	val, err := f.KV.Get("web/nginxkey")
	if err != nil {
		return err
	}
	attr.Inode = 2
	attr.Mode = 0444
	attr.Size = uint64(len(val))
	return nil
}

func (f File) ReadAll(ctx context.Context) ([]byte, error) {
	if val, err := f.KV.Get("web/nginxkey"); err != nil {
		return nil, err
	} else {
		return []byte(val), nil
	}
}
