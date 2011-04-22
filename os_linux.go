package libauth

import (
	"goplan9.googlecode.com/hg/plan9"
	"goplan9.googlecode.com/hg/plan9/client"
	"io"
	"os"
)

type RW interface {
	io.ReadCloser
	io.Writer
}

func openRPC() (RW,os.Error) {
	fsys, err := client.MountService("factotum")
	if err != nil {
		return nil,err
	}
	
	fid, err := fsys.Open("rpc", plan9.ORDWR)
	if err != nil {
		return nil,err
	}
	
	return fid, nil
}

var factotum = os.Getenv("PLAN9")+"/bin/factotum"


