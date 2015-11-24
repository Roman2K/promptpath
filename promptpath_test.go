package main

import "testing"

func assertShortens(t *testing.T, s *shortener, from, to string) {
	if actual := s.Shorten(from); actual != to {
		t.Fatalf("%#v shortened to %#v instead of %#v", from, actual, to)
	}
}

func TestShorten(t *testing.T) {
	home := "/home/my"
	s := newShortener(map[string]string{
		home + "/code/go/src": "go",
		home + "/code/map":    "map",
		home + "/code":        "code",
		home:                  "~",
	})

	assertShortens(t, s,
		"a",
		"a")

	assertShortens(t, s,
		"/a",
		"/a")

	assertShortens(t, s,
		"/home/my",
		"~")

	assertShortens(t, s,
		"/home/my/a",
		"\\[\033[90m\\]~/\\[\033[0m\\]a")

	assertShortens(t, s,
		"/home/my/code/go/src/abc/def",
		"\\[\033[90m\\]go/\\[\033[0m\\]abc/def")

	assertShortens(t, s,
		"/home/my/code/map/abc/def",
		"\\[\033[90m\\]map/\\[\033[0m\\]abc/def")

	assertShortens(t, s,
		"/home/my/code/abc/def",
		"\\[\033[90m\\]code/\\[\033[0m\\]abc/def")
}
