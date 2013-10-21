package libauth

import (
	"testing"
)

var (
	attrstr      = "proto=pass dom=example.org user=foo"
	attrexpected = map[string]string{
		"proto": "pass",
		"dom":   "example.org",
		"user":  "foo",
	}
)

func TestAttrMap(t *testing.T) {
	m := attrmap(attrstr)

	for a, v := range m {
		if ev, ok := attrexpected[a]; !ok {
			t.Errorf("attr %q not found", a)
		} else if ev != v {
			t.Errorf("got value %q expected %q", v, ev)
		}
	}
}
