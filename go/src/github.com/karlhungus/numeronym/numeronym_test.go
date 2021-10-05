package numeronym

import (
	"testing"
)

func TestNumeronym(t *testing.T) {
	got := numeronymUrl("tests.com/thing/bar/customer.maria.do")
	want := "t3s.c1m/t3g/b1r/c6r.m3a.do"
	if got != want {
		t.Errorf("mismatch want: %s , got:%s", got, want)
	}

	got = numeronymUrl("testxxxs.com/thing/bar/customer.maria.dxxxo")
	want = "t6s.c1m/t3g/b1r/c6r.m3a.d3o"
	if got != want {
		t.Errorf("mismatch want: %s , got:%s", got, want)
	}
}
