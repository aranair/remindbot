package main

import (
	"fmt"
	"github.com/jinzhu/now"
)

func main() {
	now.TimeFormats = append(now.TimeFormats, "2Jan 2006 15:04")
	now.TimeFormats = append(now.TimeFormats, "2Jan 2006 3:04pm")
	now.TimeFormats = append(now.TimeFormats, "2Jan 2006 3pm")
	now.TimeFormats = append(now.TimeFormats, "2Jan 15:04")
	now.TimeFormats = append(now.TimeFormats, "2Jan 3:04pm")
	now.TimeFormats = append(now.TimeFormats, "2Jan 3pm")

	p := fmt.Println
	t, _ := now.Parse("9jun 10pm")
	p(t)

	t, _ = now.Parse("9jun 10pm")
	p(t)

	t, _ = now.Parse("9jun 2017 9pm")
	p(t)

}
