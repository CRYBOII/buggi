package main

import (
	"regexp"
	"strings"

	"github.com/mitchellh/go-ps"
)

func CheckProcess() string {
	processList, err := ps.Processes()
	if err != nil {
		return ""

	}

	list := []string{}
	for x := range processList {
		pd := processList[x].Executable()
		if strings.Contains(pd, ".exe") {

			list = append(list, pd)
		}
	}
	return difference(blackList, list)

}

func difference(a, b []string) string {

	for _, v := range a {
		for _, o := range b {

			curr := regexp.MustCompile("(?i)" + v)
			if curr.MatchString(o) {

				return v

			}

		}

	}

	return ""
}