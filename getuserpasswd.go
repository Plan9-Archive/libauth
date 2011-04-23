package libauth

import "fmt"
import "os"

func Getuserpasswd(params string, args ...interface{}) (string,os.Error) {
	var buf [4096]byte
	f,e := openRPC()
	if e!=nil {
		return "",e
	}
	defer f.Close()

retry0:
	_,e  = f.Write([]byte(fmt.Sprintf("start "+params, args...)))
	if e!=nil {
		return "",e
	}
	n,e := f.Read(buf[:])
	if e!=nil {
		return "",e
	}
	s  := string(buf[0:n])
	ss := tokenize(s)
	switch ss[0] {
	case "ok":
	case "needkey":
		getkey(s)
		goto retry0
	default:
		return "",os.NewError(s)
	}
retry1:
	_,e  = f.Write([]byte("read"))
	if e!=nil {
		return "",e
	}
	n,e  = f.Read(buf[:])
	if e!=nil {
		return "",e
	}
	s   = string(buf[0:n])
	ss  = tokenize(s)
	switch ss[0] {
	case "needkey":
		getkey(s)
		goto retry1
	case "ok":
		return ss[2],nil
	default:
		return "",os.NewError(s)
	}
	println(s)

	return "FIFI",nil
}
