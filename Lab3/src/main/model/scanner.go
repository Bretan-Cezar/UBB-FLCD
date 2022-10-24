package model

import (
	"errors"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Scan(f *os.File) (pif []PIFEntry, st STWrapper, err error) {

	var buf []byte

	buf = make([]byte, 131072) // max. 128 kb source file

	_, err = f.Read(buf)

	if err != nil {
		return nil, STWrapper{}, err
	}

	var m0, m1, m2, m3, m4, m5, m6, tok []byte

	var line, column int

	// Matches newline characters
	reEndl, _ := regexp.Compile(`((\n+)|(\r\n+))`)

	// Matches whitespaces at the beginning of the line
	reIndent, _ := regexp.Compile(`^[ \t]+`)

	// Matches operators
	reOp, _ := regexp.Compile(`(=)|(\+)|(-)|(\*)|(\/)|(\*\*)|(\/\/)|(==)|(<)|(<=)|(>)|(>=)|(\|\|)|(&&)`)

	// Matches reserved words
	reKw, _ := regexp.Compile(`(if)|(else)|(while)|(clread)|(clwrite)|(i64)|(string)`)

	// Matches identifiers
	reId, _ := regexp.Compile(`[a-zA-Z_][a-zA-Z0-9_]{0,255}`)

	// Matches integer constants
	reIc, _ := regexp.Compile(`(0)|(-?[1-9][0-9]{0,18})`)

	// Matches string constants
	reSc, _ := regexp.Compile(`"[^"]*"`)

	// Matches separators: parentheses, brackets, braces, commas, whitespaces that are not for indentation purpose
	reSep, _ := regexp.Compile(`(\()|(\))|(\[)|(\])|(\{)|(\})|(,)|([ \t]+)`)

	for len(buf) != 0 {

		m0 = reEndl.Find(buf)

		if strings.Index(string(buf), string(m0)) == 0 {

			line += len(m0)
			column = 0

			buf = buf[len(m0):]

			pif = append(pif, PIFEntry{"\n", ENDL, -1, -1})
		}

		m0 = reIndent.Find(buf)

		if strings.Index(string(buf), string(m0)) == 0 {

			column += len(m0) //+ 2*strings.Count(string(m0), "\t")

			buf = buf[len(m0):]
		}

		m1 = reOp.Find(buf)
		m2 = reKw.Find(buf)
		m3 = reId.Find(buf)
		m4 = reIc.Find(buf)
		m5 = reSc.Find(buf)
		m6 = reSep.Find(buf)

		if strings.Index(string(buf), string(m1)) == 0 {

			column += len(m1)

			tok = buf[:len(m1)]
			buf = buf[len(m1):]

			pif = append(pif, PIFEntry{string(tok), OPERATOR, -1, -1})

		} else if strings.Index(string(buf), string(m2)) == 0 {

			column += len(m2)

			tok = buf[:len(m2)]
			buf = buf[len(m2):]

			pif = append(pif, PIFEntry{string(tok), KEYWORD, -1, -1})

		} else if strings.Index(string(buf), string(m3)) == 0 {

			column += len(m3)

			tok = buf[:len(m3)]
			buf = buf[len(m3):]

			b, err := st.Ids.HasValue(string(tok))

			if err != nil {
				return nil, st, err
			}

			var hash, index int

			if !b {

				hash, index, err = st.Ids.SetSymbol(string(tok))

				if err != nil {
					return nil, st, err
				}

			} else {

				hash, index, err = st.Ids.GetHashAndIndex(string(tok))

				if err != nil {
					return nil, st, err
				}
			}

			pif = append(pif, PIFEntry{string(tok), ID, hash, index})

		} else if strings.Index(string(buf), string(m4)) == 0 {

			column += len(m4)

			var ic int64

			tok = buf[:len(m4)]
			buf = buf[len(m4):]

			ic, err = strconv.ParseInt(string(tok), 10, 64)

			b, err := st.IntConstants.HasValue(ic)

			if err != nil {
				return nil, st, err
			}

			var hash, index int

			if !b {

				hash, index, err = st.IntConstants.SetSymbol(ic)

				if err != nil {
					return nil, st, err
				}

			} else {

				hash, index, err = st.IntConstants.GetHashAndIndex(ic)

				if err != nil {
					return nil, st, err
				}
			}

			pif = append(pif, PIFEntry{string(tok), INT_CONST, hash, index})

		} else if strings.Index(string(buf), string(m5)) == 0 {

			column += len(m5)

			tok = buf[1 : len(m5)-1]
			buf = buf[len(m5):]

			b, err := st.StringConstants.HasValue(string(tok))

			if err != nil {
				return nil, st, err
			}

			var hash, index int

			if !b {

				hash, index, err = st.StringConstants.SetSymbol(string(tok))

				if err != nil {
					return nil, st, err
				}

			} else {

				hash, index, err = st.StringConstants.GetHashAndIndex(string(tok))

				if err != nil {
					return nil, st, err
				}
			}

			pif = append(pif, PIFEntry{string(tok), STR_CONST, hash, index})

		} else if strings.Index(string(buf), string(m6)) == 0 {

			column += len(m6) //+ 2*strings.Count(string(m0), "\t")

			tok = buf[:len(m6)]
			buf = buf[len(m6):]

			pif = append(pif, PIFEntry{string(tok), SEPARATOR, -1, -1})

		} else {

			return nil, st, errors.New("lexical error at " + strconv.Itoa(line) + ":" + strconv.Itoa(column))
		}
	}

	return
}
