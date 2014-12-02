// +build !plan9

package libauth

import (
	"code.google.com/p/goplan9/plan9"
	"code.google.com/p/goplan9/plan9/client"
	"io"
	"os"
)

func openRPC() (io.ReadWriteCloser, error) {
	fsys, err := client.MountService("factotum")
	if err != nil {
		return nil, err
	}

	fid, err := fsys.Open("rpc", plan9.ORDWR)
	if err != nil {
		return nil, err
	}

	return fid, nil
}

func openCtl() (io.ReadWriteCloser, error) {
	fsys, err := client.MountService("factotum")
	if err != nil {
		return nil, err
	}

	fid, err := fsys.Open("ctl", plan9.ORDWR)
	if err != nil {
		return nil, err
	}

	return fid, nil
}

var factotum = os.Getenv("PLAN9") + "/bin/factotum"
