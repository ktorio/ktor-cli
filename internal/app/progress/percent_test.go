package progress

import (
	"fmt"
	"strings"
	"testing"
	"unicode"
)

func TestWrites(t *testing.T) {
	var b strings.Builder
	checkAllWrites(t, &Percent{prefix: "___", total: 5, enabled: true, Writer: &b}, &b, []testCase{
		{writeLen: 1, expected: withClearLine("___20%")},
		{writeLen: 1, expected: withClearLine("___40%")},
		{writeLen: 1, expected: withClearLine("___60%")},
		{writeLen: 1, expected: withClearLine("___80%")},
		{writeLen: 1, expected: withClearLine("___100%")},
	})

	checkAllWrites(t, &Percent{prefix: "", total: 3, enabled: true, Writer: &b}, &b, []testCase{
		{writeLen: 1, expected: withClearLine("33%")},
		{writeLen: 1, expected: withClearLine("66%")},
		{writeLen: 1, expected: withClearLine("100%")},
	})

	checkAllWrites(t, &Percent{prefix: "", total: 3, enabled: true, Writer: &b}, &b, []testCase{
		{writeLen: 1, expected: withClearLine("33%")},
		{writeLen: 1, expected: withClearLine("66%")},
		{writeLen: 1, expected: withClearLine("100%")},
		{writeLen: -1, expected: withClearLine("100%\n")},
	})

	checkAllWrites(t, &Percent{prefix: "", total: 10, enabled: true, Writer: &b}, &b, []testCase{
		{writeLen: 3, expected: withClearLine("30%")},
		{writeLen: 3, expected: withClearLine("60%")},
		{writeLen: 3, expected: withClearLine("90%")},
		{writeLen: -1, expected: withClearLine("100%\n")},
	})

	checkAllWrites(t, &Percent{prefix: "", total: 10, enabled: false, Writer: &b}, &b, []testCase{
		{writeLen: 3, expected: ""},
		{writeLen: 3, expected: ""},
		{writeLen: 3, expected: ""},
		{writeLen: -1, expected: ""},
	})

	checkAllWrites(t, &Percent{prefix: "___", total: 5, enabled: true, Writer: &b}, &b, []testCase{
		{writeLen: 1, offset: 0, expected: withClearLine("___20%")},
		{writeLen: 1, offset: 1, expected: withClearLine("___40%")},
		{writeLen: 1, offset: 2, expected: withClearLine("___60%")},
		{writeLen: 1, offset: 3, expected: withClearLine("___80%")},
		{writeLen: 1, offset: 4, expected: withClearLine("___100%")},
	})

	checkAllWrites(t, &Percent{prefix: "___", total: 100, enabled: true, Writer: &b}, &b, []testCase{
		{writeLen: 20, offset: 0, expected: withClearLine("___20%")},
		{writeLen: 20, offset: 40, expected: withClearLine("___40%")},
		{writeLen: 20, offset: 80, expected: withClearLine("___60%")},
		{writeLen: 20, offset: 20, expected: withClearLine("___80%")},
		{writeLen: 20, offset: 60, expected: withClearLine("___100%")},
	})

	checkAllWrites(t, &Percent{prefix: "", total: 3, enabled: true, Writer: &b}, &b, []testCase{
		{writeLen: 1, offset: 0, expected: withClearLine("33%")},
		{writeLen: 1, offset: 1, expected: withClearLine("66%")},
		{writeLen: 1, offset: 2, expected: withClearLine("100%")},
		{writeLen: -1, offset: 0, expected: withClearLine("100%\n")},
	})
}

type testCase struct {
	writeLen int
	offset   int64
	expected string
}

func checkAllWrites(t *testing.T, p *Percent, b *strings.Builder, cases []testCase) {
	for _, test := range cases {
		var err error
		if test.writeLen == -1 {
			err = p.Done()
		} else {
			_, err = p.Write(make([]byte, test.writeLen))
		}

		assert(t, test.expected, err, b)

		b.Reset()
	}

	p.reset()

	for _, test := range cases {
		var err error
		if test.writeLen == -1 {
			err = p.Done()
		} else {
			_, err = p.WriteAt(make([]byte, test.writeLen), test.offset)
		}

		assert(t, test.expected, err, b)

		b.Reset()
	}
}

func assert(t *testing.T, expected string, err error, b *strings.Builder) {
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	if expected != b.String() {
		t.Fatalf("expected %s, got %s", esc(expected), esc(b.String()))
	}
}

func withClearLine(s string) string {
	return fmt.Sprintf("\u001B[2K\r%s", s)
}

func esc(s string) string {
	var b strings.Builder
	for _, r := range s {
		if !unicode.IsGraphic(r) {
			b.WriteString(fmt.Sprintf("\\u00%02d", r))
		} else {
			b.WriteRune(r)
		}
	}
	return b.String()
}
