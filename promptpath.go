package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

var mainShortener *shortener

func init() {
	home := os.Getenv("HOME")
	if home == "" {
		fmt.Fprintln(os.Stderr, "$HOME not found, can't initialize")
		os.Exit(1)
	}
	mainShortener = newShortener(map[string]string{
		home + "/code/go/src": "go",
		home + "/code/map":    "map",
		home + "/code":        "code",
		home:                  "~",
	})
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <path>\n", os.Args[0])
		os.Exit(2)
	}
	path := os.Args[1]
	os.Stdout.WriteString(mainShortener.Shorten(path))
}

type shortener struct {
	shortcuts map[string]string
	longsRe   *regexp.Regexp
}

func newShortener(shortcuts map[string]string) (s *shortener) {
	s = &shortener{shortcuts, nil}
	longs := make([]string, 0, len(s.shortcuts))
	for long := range s.shortcuts {
		longs = append(longs, long)
	}
	sort.Sort(byLongest(longs))
	s.longsRe = regexp.MustCompile(`^(` + strings.Join(longs, "|") + `)(?:/(.+)|$)`)
	return
}

func (s *shortener) Shorten(path string) string {
	if m := s.longsRe.FindStringSubmatch(path); m != nil {
		if m[2] == "" {
			return s.shortcuts[m[1]]
		}
		return "\033[90m" + s.shortcuts[m[1]] + "/\033[0m" + m[2]
	}
	return path
}

type byLongest []string

func (s byLongest) Len() int {
	return len(s)
}

func (s byLongest) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byLongest) Less(i, j int) bool {
	return len(s[i]) > len(s[j])
}
