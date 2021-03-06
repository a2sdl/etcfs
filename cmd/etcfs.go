package main

import (
	"log"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
	"github.com/a2sdl/etcfs"
)

func main() {
	mountpoint := "/Users/gtarcea/fuse/nginx"
	fuse.Unmount(mountpoint)
	conn, err := fuse.Mount(
		mountpoint,
		fuse.FSName("etcfs"),
		fuse.Subtype("etcfs"),
		fuse.LocalVolume(),
		fuse.VolumeName("etcfs"),
	)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	err = fs.Serve(conn, etcfs.FS{})

	<-conn.Ready
	if conn.MountError != nil {
		log.Fatal(conn.MountError)
	}
}
