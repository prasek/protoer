package desc

import (
	"testing"

	"github.com/prasek/protoer/internal/test/testutil"
)

func TestCreatePrefixList(t *testing.T) {
	list := createPrefixList("")
	testutil.Eq(t, []string{""}, list)

	list = createPrefixList("pkg")
	testutil.Eq(t, []string{"pkg", ""}, list)

	list = createPrefixList("fully.qualified.pkg.name")
	testutil.Eq(t, []string{"fully.qualified.pkg.name", "fully.qualified.pkg", "fully.qualified", "fully", ""}, list)
}

func TestUnescape(t *testing.T) {
	testCases := []struct {
		in, out string
	}{
		// EOF, bad escapes
		{"\\", "\\"},
		{"\\y", "\\y"},
		// octal escapes
		{"\\0", "\000"},
		{"\\7", "\007"},
		{"\\07", "\007"},
		{"\\77", "\077"},
		{"\\78", "\0078"},
		{"\\077", "\077"},
		{"\\377", "\377"},
		{"\\128", "\0128"},
		{"\\0001", "\0001"},
		{"\\0008", "\0008"},
		// bad octal escape
		{"\\8", "\\8"},
		// hex escapes
		{"\\x0", "\x00"},
		{"\\x7", "\x07"},
		{"\\x07", "\x07"},
		{"\\x77", "\x77"},
		{"\\x7g", "\x07g"},
		{"\\xcc", "\xcc"},
		{"\\xfff", "\xfff"},
		// bad hex escape
		{"\\xg7", "\\xg7"},
		{"\\x", "\\x"},
		// short unicode escapes
		{"\\u0020", "\u0020"},
		{"\\u007e", "\u007e"},
		{"\\u1234", "\u1234"},
		{"\\uffff", "\uffff"},
		// long unicode escapes
		{"\\U00000024", "\U00000024"},
		{"\\U00000076", "\U00000076"},
		{"\\U00001234", "\U00001234"},
		{"\\U0010FFFF", "\U0010FFFF"},
		// bad unicode escapes
		{"\\u12", "\\u12"},
		{"\\ug1232", "\\ug1232"},
		{"\\u", "\\u"},
		{"\\U1234567", "\\U1234567"},
		{"\\U12345678", "\\U12345678"},
		{"\\U0010Fghi", "\\U0010Fghi"},
		{"\\U", "\\U"},
	}
	for _, tc := range testCases {
		for _, p := range []string{"", "prefix"} {
			for _, s := range []string{"", "suffix"} {
				i := p + tc.in + s
				o := p + tc.out + s
				u := unescape(i)
				testutil.Eq(t, o, u, "unescaped %q into %q, but should have been %q", i, u, o)
			}
		}
	}
}
