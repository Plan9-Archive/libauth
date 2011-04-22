package libauth

import (
	"os"
)

func openRPC() (*os.File,os.Error) {
	return os.OpenFile("/mnt/factotum/rpc", os.O_RDWR, 0)
}

var factotum = "/boot/factotum"


