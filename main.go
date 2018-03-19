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
	var (
		gosrc = home + "/code/go/src"
		gogh  = gosrc + "/github.com/Roman2K"
		gogl  = gosrc + "/gitlab.com/Roman2K"
		gomap = gosrc + "/github.com/MonAlbumPhoto"
		glmap = gosrc + "/map-gitlab.bearstech.com/mon-album-photo"
	)
	mainShortener = newShortener(map[string]string{
		home: "~",

		// code
		home + "/code":     "code",
		home + "/code/map": "map",

		// go
		gosrc: "gosrc",

		// go GitHub Roman2K
		gogh:           "gogh",
		gogh + "/scat": "scat",

		// go GitLab Roman2K
		gogl: "gogl",

		// go GitHub MonAlbumPhoto
		gomap: "gomap",

		// go GitLab MonAlbumPhoto
		glmap: "glmap",
		glmap + "/MAP-storage":         "mapstorage",
		glmap + "/MAP-workers":         "mapworkers",
		glmap + "/MAP-download-images": "mapdlimages",
		glmap + "/map-post-image":      "mappostimage",
		glmap + "/map-process-image":   "mapprocessimage",
	})
}

func main() {
	switch len(os.Args) {
	case 1:
		printMapping()
	case 2:
		path := os.Args[1]
		os.Stdout.WriteString(mainShortener.Shorten(path))
	default:
		fmt.Fprintf(os.Stderr, "Usage: %s [path]\n", os.Args[0])
		os.Exit(2)
	}
}

func printMapping() {
	for long, short := range mainShortener.shortcuts {
		fmt.Printf("%s\t%s\n", short, long)
	}
}

type shortener struct {
	shortcuts map[string]string
	longsRe   *regexp.Regexp
}

func newShortener(shortcuts map[string]string) *shortener {
	s := &shortener{shortcuts, nil}
	longs := make([]string, 0, len(s.shortcuts))
	for long := range s.shortcuts {
		longs = append(longs, regexp.QuoteMeta(long))
	}
	sort.Sort(byLongest(longs))
	s.longsRe = regexp.MustCompile(
		`^(` + strings.Join(longs, "|") + `)(?:/(.+)|$)`,
	)
	return s
}

func (s *shortener) Shorten(path string) string {
	if m := s.longsRe.FindStringSubmatch(path); m != nil {
		short := s.shortcuts[m[1]]
		if m[2] == "" {
			return short
		}
		return color("38;5;241") + short + "/" + color("0") + m[2]
	}
	return path
}

func color(attributes string) string {
	return "\\[\033[" + attributes + "m\\]"
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
