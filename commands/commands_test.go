package commands

import "testing"

func checkExpected(t *testing.T, actual string, expected string) {
	if actual != expected {
		t.Fatalf("Expected %s but got %s", expected, actual)
	}
}

func TestExtract(t *testing.T) {
	cmds := NewCommandList()
	cmd, txt, dds := cmds.Extract("remind me to do this:9jun 10pm")
	checkExpected(t, cmd, "remind")
	checkExpected(t, txt, "do this")
	checkExpected(t, dds, "9jun 10pm")
}
