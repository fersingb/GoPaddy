package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/fatih/color"
)

var logo = `
 ▄▄ •        ▄▄▄· ▄▄▄· ·▄▄▄▄  ·▄▄▄▄   ▄· ▄▌
▐█ ▀ ▪▪     ▐█ ▄█▐█ ▀█ ██▪ ██ ██▪ ██ ▐█▪██▌
▄█ ▀█▄ ▄█▀▄  ██▀·▄█▀▀█ ▐█· ▐█▌▐█· ▐█▌▐█▌▐█▪
▐█▄▪▐█▐█▌.▐▌▐█▪·•▐█ ▪▐▌██. ██ ██. ██  ▐█▀·.
·▀▀▀▀  ▀█▄▀▪.▀    ▀  ▀ ▀▀▀▀▀• ▀▀▀▀▀•   ▀ • 
`

func Logo() {
	indent := "                 "
	cyan := color.New(color.FgCyan)

	for _, s := range strings.Split(logo, "\n") {
		time.Sleep(50 * time.Millisecond)
		fmt.Fprintf(color.Error, cyan.Sprintf("%s%s\n", indent, s))
	}
}
