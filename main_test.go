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
		color("90")+"~/"+color("0")+"a")

	assertShortens(t, s,
		"/home/my/code/go/src/abc/def",
		color("90")+"go/"+color("0")+"abc/def")

	assertShortens(t, s,
		"/home/my/code/abc/def",
		color("90")+"code/"+color("0")+"abc/def")
}
