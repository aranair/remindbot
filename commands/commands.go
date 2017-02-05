package commands

import (
	"regexp"
)

type Commands struct {
	rmt   *regexp.Regexp
	r     *regexp.Regexp
	l     *regexp.Regexp
	rn    *regexp.Regexp
	c     *regexp.Regexp
	cl    *regexp.Regexp
	hazel *regexp.Regexp
}

func NewCommandList() Commands {
	return Commands{
		rmt:   compileRegexp(`(?i)^(remind) me to (.+)`),
		r:     compileRegexp(`(?i)^(remind) (.+)`),
		l:     compileRegexp(`(?i)^(list)`),
		c:     compileRegexp(`(?i)^(clear) (\d+)`),
		rn:    compileRegexp(`(?i)^(renum)`),
		cl:    compileRegexp(`(?i)^(clearall)`),
		hazel: compileRegexp(`(?i)(hazel)`),
	}
}

func compileRegexp(s string) *regexp.Regexp {
	r, _ := regexp.Compile(s)
	return r
}

func (c *Commands) Extract(t string) (string, string) {
	var a []string

	a = c.rmt.FindStringSubmatch(t)
	if len(a) == 3 {
		return a[1], a[2]
	}

	a = c.r.FindStringSubmatch(t)
	if len(a) == 3 {
		return a[1], a[2]
	}

	a = c.l.FindStringSubmatch(t)
	if len(a) == 2 {
		return a[0], a[1]
	}

	a = c.c.FindStringSubmatch(t)
	if len(a) == 3 {
		return a[1], a[2]
	}

	a = c.rn.FindStringSubmatch(t)
	if len(a) == 2 {
		return a[0], a[1]
	}

	a = c.cl.FindStringSubmatch(t)
	if len(a) == 2 {
		return a[0], a[1]
	}

	a = c.hazel.FindStringSubmatch(t)
	if len(a) == 2 {
		return a[0], a[1]
	}

	return "", ""
}
