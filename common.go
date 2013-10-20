package libauth

import (
	"os"
	"strings"
)

func getkey(params string) error {
	p, e := os.StartProcess(factotum, []string{"getkey", "-g", params},
		&os.ProcAttr{Files: []*os.File{os.Stdin, os.Stdout, os.Stderr}})
	if e != nil {
		return e
	}
	_, e = p.Wait()
	return e
}

func tokenize(s string) []string {
	ss := []string{}
	tmp := []byte{}
	quot := false
	for i := 0; i < len(s); i++ {
		if (s[i] == ' ' || s[i] == '\t' || s[i] == '\n') && !quot {
			ss = append(ss, string(tmp))
			tmp = []byte{}
			continue
		}
		if s[i] != '\'' {
			tmp = append(tmp, s[i])
			continue
		}
		if !quot {
			quot = true
			continue
		}
		if i+1 == len(s) || s[i+1] != '\'' {
			quot = false
			continue
		}
		tmp = append(tmp, '\'')
		i++
	}
	if len(tmp) > 0 {
		ss = append(ss, string(tmp))
	}
	return ss
}

// given a string of the form 'proto=foo service=bar user=baz', tokenize it into a map.
func attrmap(s string) map[string]string {
	attrmap := make(map[string]string)

	strs := tokenize(s)
	for _, av := range strs {
		a := strings.Split(av, "=")
		if len(a) == 1 {
			attrmap[a[0]] = ""
		} else {
			attrmap[a[0]] = a[1]
		}
	}

	return attrmap
}
