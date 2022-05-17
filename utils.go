package main

import (
	"strings"
)

type PathInfo struct {
	num      int
	elements []string
}

func getPathInfo(path string) (pathInfo PathInfo) {
	path = strings.Trim(path, "/")
	elements := strings.Split(path, "/")
	pathInfo.num = len(elements)
	pathInfo.elements = elements
	if pathInfo.elements[0] == "" { // even if url is http://localhost:8000/, pathInfo.num is 1 because pathInfo.elements include "" as a element.
		pathInfo.num = 0
		pathInfo.elements = []string{}
	}
	return
}
