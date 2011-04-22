package libauth

import "os"

func getkey(params string) os.Error {
	p,e := os.StartProcess(factotum, []string{"getkey","-g",params},
		&os.ProcAttr{Files: []*os.File{os.Stdin, os.Stdout, os.Stderr}})
	if e != nil {
		return e
	}
	_,e = p.Wait(0)
	return e
}

func tokenize(s string) []string {
	ss   := []string{}
	tmp  := []byte{}
	quot := false
	for i:=0; i<len(s); i++ {
		if s[i] == ' ' && !quot {
			ss = append(ss, string(tmp))
			tmp= []byte{}
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
		if i+1==len(s) || s[i+1]!='\'' {
			quot = false
			continue
		}
		tmp = append(tmp, '\'')
		i++
	}
	if len(tmp)>0 {
		ss = append(ss, string(tmp))
	}
	return ss
}
