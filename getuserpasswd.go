package libauth

import "fmt"
import "errors"
import "strings"

func Getuserpasswd(params string, args ...interface{}) (string, error) {
	var buf [4096]byte
	f, e := openRPC()
	if e != nil {
		return "", e
	}
	defer f.Close()

retry0:
	_, e = f.Write([]byte(fmt.Sprintf("start "+params, args...)))
	if e != nil {
		return "", e
	}
	n, e := f.Read(buf[:])
	if e != nil {
		return "", e
	}
	s := string(buf[0:n])
	ss := tokenize(s)
	switch ss[0] {
	case "ok":
	case "needkey":
		getkey(strings.Join(ss[1:], " "))
		goto retry0
	default:
		return "", errors.New(s)
	}
retry1:
	_, e = f.Write([]byte("read"))
	if e != nil {
		return "", e
	}
	n, e = f.Read(buf[:])
	if e != nil {
		return "", e
	}
	s = string(buf[0:n])
	ss = tokenize(s)
	switch ss[0] {
	case "needkey":
		getkey(strings.Join(ss[1:], " "))
		goto retry1
	case "ok":
		return ss[2], nil
	default:
		return "", errors.New(s)
	}
	println(s)

	return "FIFI", nil
}
