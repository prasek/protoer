package desc

import (
	"strconv"
	"unicode/utf8"
)

const (
	MaxTag = 536870911 // 2^29 - 1

	SpecialReservedStart = 19000
	SpecialReservedEnd   = 19999

	// NB: It would be nice to use constants from generated code instead of hard-coding these here.
	// But code-gen does not emit these as constants anywhere. The only places they appear in generated
	// code are struct tags on fields of the generated descriptor protos.
	File_packageTag           = 2
	File_dependencyTag        = 3
	File_messagesTag          = 4
	File_enumsTag             = 5
	File_servicesTag          = 6
	File_extensionsTag        = 7
	File_optionsTag           = 8
	File_syntaxTag            = 12
	Message_nameTag           = 1
	Message_fieldsTag         = 2
	Message_nestedMessagesTag = 3
	Message_enumsTag          = 4
	Message_extensionRangeTag = 5
	Message_extensionsTag     = 6
	Message_optionsTag        = 7
	Message_oneOfsTag         = 8
	Message_reservedRangeTag  = 9
	Message_reservedNameTag   = 10
	ExtensionRange_startTag   = 1
	ExtensionRange_endTag     = 2
	ExtensionRange_optionsTag = 3
	ReservedRange_startTag    = 1
	ReservedRange_endTag      = 2
	Field_nameTag             = 1
	Field_extendeeTag         = 2
	Field_numberTag           = 3
	Field_labelTag            = 4
	Field_typeTag             = 5
	Field_defaultTag          = 7
	Field_optionsTag          = 8
	Field_jsonNameTag         = 10
	OneOf_nameTag             = 1
	OneOf_optionsTag          = 2
	Enum_nameTag              = 1
	Enum_valuesTag            = 2
	Enum_optionsTag           = 3
	EnumVal_nameTag           = 1
	EnumVal_numberTag         = 2
	EnumVal_optionsTag        = 3
	Service_nameTag           = 1
	Service_methodsTag        = 2
	Service_optionsTag        = 3
	Method_nameTag            = 1
	Method_inputTag           = 2
	Method_outputTag          = 3
	Method_optionsTag         = 4
	Method_inputStreamTag     = 5
	Method_outputStreamTag    = 6

	// All *Options messages use the same tag for the field
	// that stores uninterpreted options
	UninterpretedOptionsTag = 999

	Uninterpreted_nameTag      = 2
	Uninterpreted_identTag     = 3
	Uninterpreted_posIntTag    = 4
	Uninterpreted_negIntTag    = 5
	Uninterpreted_doubleTag    = 6
	Uninterpreted_stringTag    = 7
	Uninterpreted_aggregateTag = 8
	UninterpretedName_nameTag  = 1
)

// createPrefixList returns a list of package prefixes to search when resolving
// a symbol name. If the given package is blank, it returns only the empty
// string. If the given package contains only one token, e.g. "foo", it returns
// that token and the empty string, e.g. ["foo", ""]. Otherwise, it returns
// successively shorter prefixes of the package and then the empty string. For
// example, for a package named "foo.bar.baz" it will return the following list:
//   ["foo.bar.baz", "foo.bar", "foo", ""]
func createPrefixList(pkg string) []string {
	if pkg == "" {
		return []string{""}
	}

	numDots := 0
	// one pass to pre-allocate the returned slice
	for i := 0; i < len(pkg); i++ {
		if pkg[i] == '.' {
			numDots++
		}
	}
	if numDots == 0 {
		return []string{pkg, ""}
	}

	prefixes := make([]string, numDots+2)
	// second pass to fill in returned slice
	for i := 0; i < len(pkg); i++ {
		if pkg[i] == '.' {
			prefixes[numDots] = pkg[:i]
			numDots--
		}
	}
	prefixes[0] = pkg

	return prefixes
}

func unescape(s string) string {
	// protoc encodes default values for 'bytes' fields using C escaping,
	// so this function reverses that escaping
	out := make([]byte, 0, len(s))
	var buf [4]byte
	for len(s) > 0 {
		if s[0] != '\\' || len(s) < 2 {
			// not escape sequence, or too short to be well-formed escape
			out = append(out, s[0])
			s = s[1:]
		} else if s[1] == 'x' || s[1] == 'X' {
			n := matchPrefix(s[2:], 2, isHex)
			if n == 0 {
				// bad escape
				out = append(out, s[:2]...)
				s = s[2:]
			} else {
				c, err := strconv.ParseUint(s[2:2+n], 16, 8)
				if err != nil {
					// shouldn't really happen...
					out = append(out, s[:2+n]...)
				} else {
					out = append(out, byte(c))
				}
				s = s[2+n:]
			}
		} else if s[1] >= '0' && s[1] <= '7' {
			n := 1 + matchPrefix(s[2:], 2, isOctal)
			c, err := strconv.ParseUint(s[1:1+n], 8, 8)
			if err != nil || c > 0xff {
				out = append(out, s[:1+n]...)
			} else {
				out = append(out, byte(c))
			}
			s = s[1+n:]
		} else if s[1] == 'u' {
			if len(s) < 6 {
				// bad escape
				out = append(out, s...)
				s = s[len(s):]
			} else {
				c, err := strconv.ParseUint(s[2:6], 16, 16)
				if err != nil {
					// bad escape
					out = append(out, s[:6]...)
				} else {
					w := utf8.EncodeRune(buf[:], rune(c))
					out = append(out, buf[:w]...)
				}
				s = s[6:]
			}
		} else if s[1] == 'U' {
			if len(s) < 10 {
				// bad escape
				out = append(out, s...)
				s = s[len(s):]
			} else {
				c, err := strconv.ParseUint(s[2:10], 16, 32)
				if err != nil || c > 0x10ffff {
					// bad escape
					out = append(out, s[:10]...)
				} else {
					w := utf8.EncodeRune(buf[:], rune(c))
					out = append(out, buf[:w]...)
				}
				s = s[10:]
			}
		} else {
			switch s[1] {
			case 'a':
				out = append(out, '\a')
			case 'b':
				out = append(out, '\b')
			case 'f':
				out = append(out, '\f')
			case 'n':
				out = append(out, '\n')
			case 'r':
				out = append(out, '\r')
			case 't':
				out = append(out, '\t')
			case 'v':
				out = append(out, '\v')
			case '\\':
				out = append(out, '\\')
			case '\'':
				out = append(out, '\'')
			case '"':
				out = append(out, '"')
			case '?':
				out = append(out, '?')
			default:
				// invalid escape, just copy it as-is
				out = append(out, s[:2]...)
			}
			s = s[2:]
		}
	}
	return string(out)
}

func isOctal(b byte) bool { return b >= '0' && b <= '7' }

func isHex(b byte) bool {
	return (b >= '0' && b <= '9') || (b >= 'a' && b <= 'f') || (b >= 'A' && b <= 'F')
}

func matchPrefix(s string, limit int, fn func(byte) bool) int {
	l := len(s)
	if l > limit {
		l = limit
	}
	i := 0
	for ; i < l; i++ {
		if !fn(s[i]) {
			return i
		}
	}
	return i
}

func merge(a, b string) string {
	if a == "" {
		return b
	} else {
		return a + "." + b
	}
}
