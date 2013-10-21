package libauth

import (
	"os"
)

func openRPC() (*os.File, error) {
	return os.OpenFile("/mnt/factotum/rpc", os.O_RDWR, 0)
}

func openCtl() (*os.File, error) {
	return os.OpenFile("/mnt/factotum/ctl", os.O_RDWR, 0)
}

var factotum = "/boot/factotum"
