package commands

import "testing"
import "time"

func checkExpected(t *testing.T, actual string, expected string) {
	if actual != expected {
		t.Fatalf("Expected '%s' but got '%s'", expected, actual)
	}
}

func checkTime(t *testing.T, actual time.Time, expected time.Time) {
	if actual != expected {
		t.Fatalf("Expected '%s' but got '%s'", expected, actual)
	}
}

func checkNilTime(t *testing.T, actual *time.Time) {
	if !actual.IsZero() {
		t.Fatalf("Expected nil but got '%s'", actual)
	}
}

var cmds = NewCommandList()
var cmd string
var txt string
var dds time.Time

func TestRemindWithColon(t *testing.T) {
	expected := time.Date(2017, 6, 9, 22, 0, 0, 0, time.Now().Location())

	cmd, txt, dds = cmds.Extract("remind do this:9jun 10pm")
	checkExpected(t, cmd, "remind")
	checkExpected(t, txt, "do this")
	checkTime(t, dds, expected)

	cmd, txt, dds = cmds.Extract("remind do this: 9jun 10pm")
	checkExpected(t, cmd, "remind")
	checkExpected(t, txt, "do this")
	checkTime(t, dds, expected)

	cmd, txt, dds = cmds.Extract("remind do this : 9jun 10pm")
	checkExpected(t, cmd, "remind")
	checkExpected(t, txt, "do this")
	checkTime(t, dds, expected)

	cmd, txt, dds = cmds.Extract("remind me to do this:9jun 10pm")
	checkExpected(t, cmd, "remind")
	checkExpected(t, txt, "do this")
	checkTime(t, dds, expected)

	cmd, txt, dds = cmds.Extract("remind me to do this: 9jun 10pm")
	checkExpected(t, cmd, "remind")
	checkExpected(t, txt, "do this")
	checkTime(t, dds, expected)

	cmd, txt, dds = cmds.Extract("remind me to do this : 9jun 10pm")
	checkExpected(t, cmd, "remind")
	checkExpected(t, txt, "do this")
	checkTime(t, dds, expected)
}

func TestRemindWithoutColon(t *testing.T) {
	cmd, txt, dds = cmds.Extract("remind me to do this")
	checkExpected(t, cmd, "remind")
	checkExpected(t, txt, "do this")
	checkNilTime(t, &dds)

	cmd, txt, dds = cmds.Extract("remind do this")
	checkExpected(t, cmd, "remind")
	checkExpected(t, txt, "do this")
	checkNilTime(t, &dds)
}

func TestList(t *testing.T) {
	cmd, txt, dds = cmds.Extract("list")
	checkExpected(t, cmd, "list")
	checkExpected(t, txt, "")
	checkNilTime(t, &dds)

	cmd, txt, dds = cmds.Extract("List")
	checkExpected(t, cmd, "list")
	checkExpected(t, txt, "")
	checkNilTime(t, &dds)

	cmd, txt, dds = cmds.Extract("listen this is not a list")
	checkExpected(t, cmd, "")
	checkExpected(t, txt, "")
	checkNilTime(t, &dds)
}

func TestClear(t *testing.T) {
	cmd, txt, dds = cmds.Extract("clear 2")
	checkExpected(t, cmd, "clear")
	checkExpected(t, txt, "2")
	checkNilTime(t, &dds)

	cmd, txt, dds = cmds.Extract("clearance sale")
	checkExpected(t, cmd, "")
	checkExpected(t, txt, "")
	checkNilTime(t, &dds)
}

func TestRenum(t *testing.T) {
	cmd, txt, dds = cmds.Extract("renum")
	checkExpected(t, cmd, "renum")
	checkExpected(t, txt, "")
	checkNilTime(t, &dds)

	cmd, txt, dds = cmds.Extract("renum-extra-random-characters")
	checkExpected(t, cmd, "")
	checkExpected(t, txt, "")
	checkNilTime(t, &dds)
}

func TestHazel(t *testing.T) {
	cmd, txt, dds = cmds.Extract("hazel")
	checkExpected(t, cmd, "hazel")
	checkExpected(t, txt, "")
	checkNilTime(t, &dds)

	cmd, txt, dds = cmds.Extract("hazel~")
	checkExpected(t, cmd, "hazel")
	checkExpected(t, txt, "")
	checkNilTime(t, &dds)

	cmd, txt, dds = cmds.Extract("hazel!")
	checkExpected(t, cmd, "hazel")
	checkExpected(t, txt, "")
	checkNilTime(t, &dds)

	cmd, txt, dds = cmds.Extract("hazelnut")
	checkExpected(t, cmd, "")
	checkExpected(t, txt, "")
	checkNilTime(t, &dds)
}

func TestClearall(t *testing.T) {
	cmd, txt, dds = cmds.Extract("clearall")
	checkExpected(t, cmd, "clearall")
	checkExpected(t, txt, "")
	checkNilTime(t, &dds)

	cmd, txt, dds = cmds.Extract("clearallrandomchar")
	checkExpected(t, cmd, "")
	checkExpected(t, txt, "")
	checkNilTime(t, &dds)
}
