// Copyright 2017 Gregory Siems. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

// Package naco is an attempt to implement NACO authority normalization
// rules as specified by http://www.loc.gov/aba/pcc/naco/normrule-2.html
package naco

import (
	"strings"

	"golang.org/x/text/unicode/norm"
)

/*
http://www.loc.gov/aba/pcc/naco/normrule-2.html
http://memory.loc.gov/diglib/codetables/45.html
*/

// Normalize implements the normalization rules according to Appendix A
// of the Authority File Comparison Rules web-page.
func Normalize(s string, keepFirstComma bool) (ns string) {

	const blankChar = " "
	const deleteChar = ""

	var tr = map[string]string{
		"⁰": "0",  // SUPERSCRIPT 0
		"¹": "1",  // SUPERSCRIPT 1
		"²": "2",  // SUPERSCRIPT 2
		"³": "3",  // SUPERSCRIPT 3
		"⁴": "4",  // SUPERSCRIPT 4
		"⁵": "5",  // SUPERSCRIPT 5
		"⁶": "6",  // SUPERSCRIPT 6
		"⁷": "7",  // SUPERSCRIPT 7
		"⁸": "8",  // SUPERSCRIPT 8
		"⁹": "9",  // SUPERSCRIPT 9
		"₀": "0",  // SUBSCRIPT 0
		"₁": "1",  // SUBSCRIPT 1
		"₂": "2",  // SUBSCRIPT 2
		"₃": "3",  // SUBSCRIPT 3
		"₄": "4",  // SUBSCRIPT 4
		"₅": "5",  // SUBSCRIPT 5
		"₆": "6",  // SUBSCRIPT 6
		"₇": "7",  // SUBSCRIPT 7
		"₈": "8",  // SUBSCRIPT 8
		"₉": "9",  // SUBSCRIPT 9
		"æ": "AE", // LOWERCASE DIGRAPH AE / LATIN SMALL LIGATURE AE
		"Æ": "AE", // UPPERCASE DIGRAPH AE / LATIN CAPITAL LIGATURE AE
		"œ": "OE", // LOWERCASE DIGRAPH OE / LATIN SMALL LIGATURE OE
		"Œ": "OE", // UPPERCASE DIGRAPH OE / LATIN CAPITAL LIGATURE OE
		"đ": "D",  // LOWERCASE D WITH CROSSBAR / LATIN SMALL LETTER D WITH STROKE
		"Đ": "D",  // UPPERCASE D WITH CROSSBAR / LATIN CAPITAL LETTER D WITH STROKE
		"ð": "D",  // LOWERCASE ETH / LATIN SMALL LETTER ETH (Icelandic)
		"Ð": "D",  // UPPERCASE ETH / LATIN CAPITAL LETTER ETH (Icelandic)
		"ı": "I",  // LOWERCASE TURKISH I / LATIN SMALL LETTER DOTLESS I
		"ł": "L",  // LOWERCASE POLISH L / LATIN SMALL LETTER L WITH STROKE
		"Ł": "L",  // UPPERCASE POLISH L / LATIN CAPITAL LETTER L WITH STROKE
		"ℓ": "L",  // SCRIPT SMALL L
		"ơ": "U",  // LOWERCASE O-HOOK / LATIN SMALL LETTER O WITH HORN
		"Ơ": "U",  // UPPERCASE O-HOOK / LATIN CAPITAL LETTER O WITH HORN
		"ư": "U",  // LOWERCASE U-HOOK / LATIN SMALL LETTER U WITH HORN
		"Ư": "U",  // UPPERCASE U-HOOK / LATIN CAPITAL LETTER U WITH HORN
		"ø": "O",  // LOWERCASE SCANDINAVIAN O / LATIN SMALL LETTER O WITH STROKE
		"Ø": "O",  // UPPERCASE SCANDINAVIAN O / LATIN CAPITAL LETTER O WITH STROKE
		"þ": "TH", // LOWERCASE ICELANDIC THORN / LATIN SMALL LETTER THORN (Icelandic)
		"Þ": "TH", // UPPERCASE ICELANDIC THORN / LATIN CAPITAL LETTER THORN (Icelandic)
		"ß": "SS", // ESZETT SYMBOL
		"α": "Α",  // Greek alpha-> Uppercase Greek alpha (Do not use in NACO 1XX fields)
		"β": "Β",  // Greek beta-> Uppercase Greek beta   (Do not use in NACO 1XX fields)
		"γ": "Γ",  // Greek gamma-> Uppercase Greek gamma (Do not use in NACO 1XX fields)
		////////////////////////////////////////////////////////////////
		// Deleted chars
		"[":            deleteChar, // Opening square bracket
		"]":            deleteChar, // Closing square bracket
		string(0x0027): deleteChar, // APOSTROPHE
		string(0x200D): deleteChar, // JOINER / ZERO WIDTH JOINER
		string(0x200C): deleteChar, // NON-JOINER / ZERO WIDTH NON-JOINER
		string(0x02B9): deleteChar, // SOFT SIGN, PRIME / MODIFIER LETTER PRIME
		string(0x02BC): deleteChar, // ALIF / MODIFIER LETTER APOSTROPHE
		string(0x02BB): deleteChar, // AYN / MODIFIER LETTER TURNED COMMA
		string(0x02BA): deleteChar, // HARD SIGN, DOUBLE PRIME / MODIFIER LETTER DOUBLE PRIME
		string(0x0309): deleteChar, // PSEUDO QUESTION MARK / COMBINING HOOK ABOVE
		string(0x0300): deleteChar, // GRAVE / COMBINING GRAVE ACCENT (Varia)
		string(0x0301): deleteChar, // ACUTE / COMBINING ACUTE ACCENT (Oxia)
		string(0x0302): deleteChar, // CIRCUMFLEX / COMBINING CIRCUMFLEX ACCENT
		string(0x0303): deleteChar, // TILDE / COMBINING TILDE
		string(0x0304): deleteChar, // MACRON / COMBINING MACRON
		string(0x0306): deleteChar, // BREVE / COMBINING BREVE (Vrachy)
		string(0x0307): deleteChar, // SUPERIOR DOT / COMBINING DOT ABOVE
		string(0x0308): deleteChar, // UMLAUT, DIAERESIS / COMBINING DIAERESIS (Dialytika)
		string(0x030C): deleteChar, // HACEK / COMBINING CARON
		string(0x030A): deleteChar, // CIRCLE ABOVE, ANGSTROM / COMBINING RING ABOVE
		string(0x0361): deleteChar, // LIGATURE, FIRST HALF / COMBINING DOUBLE INVERTED BREVE	FE20	EFB8A0
		// Note 1				6C	EC	C	LIGATURE, SECOND HALF / COMBINING LIGATURE RIGHT HALF	FE21	EFB8A1
		string(0x0315): deleteChar, //	HIGH COMMA, OFF CENTER / COMBINING COMMA ABOVE RIGHT
		string(0x030B): deleteChar, //	DOUBLE ACUTE / COMBINING DOUBLE ACUTE ACCENT
		string(0x0310): deleteChar, //	CANDRABINDU / COMBINING CANDRABINDU
		string(0x0327): deleteChar, //	CEDILLA / COMBINING CEDILLA
		string(0x0328): deleteChar, //	RIGHT HOOK, OGONEK / COMBINING OGONEK
		string(0x0323): deleteChar, //	DOT BELOW / COMBINING DOT BELOW
		string(0x0324): deleteChar, //	DOUBLE DOT BELOW / COMBINING DIAERESIS BELOW
		string(0x0325): deleteChar, //	CIRCLE BELOW / COMBINING RING BELOW
		string(0x0333): deleteChar, //	DOUBLE UNDERSCORE / COMBINING DOUBLE LOW LINE
		string(0x0332): deleteChar, //	UNDERSCORE / COMBINING LOW LINE
		string(0x0326): deleteChar, //	LEFT HOOK (COMMA BELOW) / COMBINING COMMA BELOW
		string(0x031C): deleteChar, //	RIGHT CEDILLA / COMBINING LEFT HALF RING BELOW
		string(0x032E): deleteChar, //	UPADHMANIYA / COMBINING BREVE BELOW
		string(0x0360): deleteChar, //	DOUBLE TILDE, FIRST HALF / COMBINING DOUBLE TILDE	FE22	EFB8A2
		// Note 2				7B	FB	C	DOUBLE TILDE, SECOND HALF / COMBINING DOUBLE TILDE RIGHT HALF	FE23	EFB8A3
		string(0x0313): deleteChar, //	Delete	CC93	̓	7E	FE	C	HIGH COMMA, CENTERED / COMBINING COMMA ABOVE (Psili)
		////////////////////////////////////////////////////////////////
		// Blanked-out chars
		"!":  blankChar, // Exclamation mark
		"\"": blankChar, // Quotation mark
		"(":  blankChar, // Opening parenthesis
		")":  blankChar, // Closing parenthesis
		"-":  blankChar, // Hyphen, minus-
		"{":  blankChar, // Opening curly bracket
		"}":  blankChar, // Closing curly bracket
		"<":  blankChar, // Less-than sign
		">":  blankChar, // Greater-than sign
		";":  blankChar, // Semicolon
		":":  blankChar, // Colon
		".":  blankChar, // Period, decimal point
		"?":  blankChar, // Question mark
		"¿":  blankChar, // Inverted question mark
		"¡":  blankChar, // Inverted exclamation mark
		"/":  blankChar, // Slash
		"\\": blankChar, // Reverse slash
		"*":  blankChar, // Asterisk
		"|":  blankChar, // Vertical bar (fill)
		"%":  blankChar, // Percent
		"=":  blankChar, // Equals sign
		"±":  blankChar, // Plus or minus
		"⁺":  blankChar, // Superscript plus
		"⁻":  blankChar, // Superscript minus
		"®":  blankChar, // Patent mark
		"℗":  blankChar, // Sound recording copyright
		"©":  blankChar, // Copyright sign
		"°":  blankChar, // Degree sign
		"^":  blankChar, // Spacing circumflex
		"_":  blankChar, // Spacing underscore
		"`":  blankChar, // Spacing grave
		"~":  blankChar, // Spacing tilde
		"·":  blankChar, // Middle dot
	}

	var nc []string
	commaSeen := false

	for _, c := range strings.Split(s, "") {

		cp := string(norm.NFC.Bytes([]byte(c)))

		if val, ok := tr[cp]; ok {
			// Deleted chars
			if val == deleteChar {
				continue
			}
			// Blanks
			if val == blankChar && nc[len(nc)-1] == blankChar {
				continue
			}

			nc = append(nc, val)
			continue
		}

		// Commas: Comma or blank-- The first comma in $a is retained; all other converted to blank
		if cp == "," {
			if keepFirstComma && !commaSeen {
				nc = append(nc, cp)
				commaSeen = true
			}
			continue
		}

		// Blanks: there should be no duplicate blanks
		if cp == blankChar && len(nc) > 0 && nc[len(nc)-1] == blankChar {
			continue
		}

		// Modifiers
		nb := norm.NFD.Bytes([]byte(cp))
		if len(nb) > 1 {
			nc = append(nc, string(nb[0]))
		} else {
			nc = append(nc, cp)
		}
	}

	ns = strings.TrimSpace(strings.ToUpper(strings.Join(nc, "")))

	// Ensure there are no trailing commas
	if keepFirstComma && strings.HasSuffix(ns, ",") {
		ns = strings.TrimSpace(strings.Replace(ns, ",", "", -1))
	}

	return ns
}
