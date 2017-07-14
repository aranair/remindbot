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
		rmt:   compileRegexp(`(?im)^(remind){1}(?: me to)? ([^:\r\n]*)(?::?)(.*)$`),
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

func (c *Commands) Extract(t string) (string, string, string) {
	var a []string

	a = c.rmt.FindStringSubmatch(t)
	if len(a) == 4 {
		return a[1], a[2], a[3]
	}

	a = c.r.FindStringSubmatch(t)
	if len(a) == 3 {
		return a[1], a[2], ""
	}

	a = c.l.FindStringSubmatch(t)
	if len(a) == 2 {
		return a[0], a[1], ""
	}

	a = c.c.FindStringSubmatch(t)
	if len(a) == 3 {
		return a[1], a[2], ""
	}

	a = c.rn.FindStringSubmatch(t)
	if len(a) == 2 {
		return a[0], a[1], ""
	}

	a = c.cl.FindStringSubmatch(t)
	if len(a) == 2 {
		return a[0], a[1], ""
	}

	a = c.hazel.FindStringSubmatch(t)
	if len(a) == 2 {
		return a[0], a[1], ""
	}

	return "", "", ""
}
