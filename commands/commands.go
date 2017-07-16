package commands

import (
	"regexp"
	s "strings"
	"time"

	"github.com/jinzhu/now"
)

type Commands struct {
	rmt   *regexp.Regexp
	cd    *regexp.Regexp
	r     *regexp.Regexp
	l     *regexp.Regexp
	rn    *regexp.Regexp
	c     *regexp.Regexp
	cl    *regexp.Regexp
	hazel *regexp.Regexp
}

func NewCommandList() Commands {
	now.TimeFormats = append(now.TimeFormats, "2Jan 15:04 2006")
	now.TimeFormats = append(now.TimeFormats, "2Jan 3:04pm 2006")
	now.TimeFormats = append(now.TimeFormats, "2Jan 3pm 2006")

	now.TimeFormats = append(now.TimeFormats, "2Jan 2006 15:04")
	now.TimeFormats = append(now.TimeFormats, "2Jan 2006 3:04pm")
	now.TimeFormats = append(now.TimeFormats, "2Jan 2006 3pm")

	now.TimeFormats = append(now.TimeFormats, "2Jan 15:04")
	now.TimeFormats = append(now.TimeFormats, "2Jan 3:04pm")
	now.TimeFormats = append(now.TimeFormats, "2Jan 3pm")

	now.TimeFormats = append(now.TimeFormats, "2Jan")

	return Commands{
		rmt:   compileRegexp(`(?im)^(remind){1}(?: me to)? ([^:\r\n]*)(?::?)(.*)$`),
		cd:    compileRegexp(`(?im)^(check due)$`),
		l:     compileRegexp(`(?im)^(list)$`),
		c:     compileRegexp(`(?im)^(clear) (\d+)$`),
		rn:    compileRegexp(`(?im)^(renum)$`),
		cl:    compileRegexp(`(?im)^(clearall)$`),
		hazel: compileRegexp(`(?im)(hazel)(?:!|~)?$`),
	}
}

func compileRegexp(s string) *regexp.Regexp {
	r, _ := regexp.Compile(s)
	return r
}

func (c *Commands) Extract(t string) (string, string, time.Time) {
	var a []string
	var r1, r2, r3 = "", "", ""

	a = c.rmt.FindStringSubmatch(t)
	if len(a) == 4 {
		r1, r2, r3 = a[1], a[2], a[3]
	}

	a = c.cd.FindStringSubmatch(t)
	if len(a) == 2 {
		r1 = a[1]
	}

	a = c.l.FindStringSubmatch(t)
	if len(a) == 2 {
		r1 = a[1]
	}

	a = c.c.FindStringSubmatch(t)
	if len(a) == 3 {
		r1, r2 = a[1], a[2]
	}

	a = c.rn.FindStringSubmatch(t)
	if len(a) == 2 {
		r1 = a[1]
	}

	a = c.cl.FindStringSubmatch(t)
	if len(a) == 2 {
		r1 = a[1]
	}

	a = c.hazel.FindStringSubmatch(t)
	if len(a) == 2 {
		r1 = a[1]
	}

	r1 = s.ToLower(s.TrimSpace(r1))
	r2 = s.ToLower(s.TrimSpace(r2))
	r3 = s.ToLower(s.TrimSpace(r3))

	// ddt = now.Parse(r3 + " " + strconv.Itoa(time.now().Year()))
	ddt, err := now.Parse(r3)

	var r3t time.Time
	if err == nil {
		r3t = ddt
	} else {
		r3t = time.Time{}
	}

	return r1, r2, r3t
}
