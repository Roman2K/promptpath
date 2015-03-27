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
		"\033[90m~/\033[0ma")

	assertShortens(t, s,
		"/home/my/code/go/src/abc/def",
		"\033[90mgo/\033[0mabc/def")

	assertShortens(t, s,
		"/home/my/code/map/abc/def",
		"\033[90mmap/\033[0mabc/def")

	assertShortens(t, s,
		"/home/my/code/abc/def",
		"\033[90mcode/\033[0mabc/def")
}
