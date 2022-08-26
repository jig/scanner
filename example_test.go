// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Copyright 2022 Jordi Íñigo Griera. All rights reserved.

package scanner_test

import (
	"fmt"
	"strings"

	"github.com/jig/scanner"
)

func Example() {
	const src = `; This is scanned code
(def a '(list 10 3.14 -30 "hello" ¬hello¬ "hel\"lo" ¬hel¬¬lo¬ :a))
`

	var s scanner.Scanner
	s.Init(strings.NewReader(src))
	s.Filename = "example"
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		fmt.Printf("%s: (%s) %s\n", s.Position, scanner.TokenString(tok), s.TokenText())
	}

	// Output:
	// example:2:1: ("(") (
	// example:2:2: (Ident) def
	// example:2:6: (Ident) a
	// example:2:8: ("'") '
	// example:2:9: ("(") (
	// example:2:10: (Ident) list
	// example:2:15: (Int) 10
	// example:2:18: (Float) 3.14
	// example:2:23: (Int) -30
	// example:2:27: (String) "hello"
	// example:2:35: (RawString) ¬hello¬
	// example:2:43: (String) "hel\"lo"
	// example:2:53: (RawString) ¬hel¬¬lo¬
	// example:2:63: (Keyword) :a
	// example:2:65: (")") )
	// example:2:66: (")") )
}

func Example_with_actual_code() {
	const src = `; This is scanned code
(def _iter->
	(fn [acc form]
		(if (list? form)
		` + "`" + `(~(first form) ~acc ~@(rest form))
		(list form acc))))
	`

	var s scanner.Scanner
	s.Init(strings.NewReader(src))
	s.Filename = "actual-code"

	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		fmt.Printf("%s: (%s) %s\n", s.Position, scanner.TokenString(tok), s.TokenText())
	}

	// Output:
	// actual-code:2:1: ("(") (
	// actual-code:2:2: (Ident) def
	// actual-code:2:6: (Ident) _iter->
	// actual-code:3:2: ("(") (
	// actual-code:3:3: (Ident) fn
	// actual-code:3:6: ("[") [
	// actual-code:3:7: (Ident) acc
	// actual-code:3:11: (Ident) form
	// actual-code:3:15: ("]") ]
	// actual-code:4:3: ("(") (
	// actual-code:4:4: (Ident) if
	// actual-code:4:7: ("(") (
	// actual-code:4:8: (Ident) list?
	// actual-code:4:14: (Ident) form
	// actual-code:4:18: (")") )
	// actual-code:5:3: ("`") `
	// actual-code:5:4: ("(") (
	// actual-code:5:5: ("~") ~
	// actual-code:5:6: ("(") (
	// actual-code:5:7: (Ident) first
	// actual-code:5:13: (Ident) form
	// actual-code:5:17: (")") )
	// actual-code:5:19: ("~") ~
	// actual-code:5:20: (Ident) acc
	// actual-code:5:24: (Ident) ~@
	// actual-code:5:26: ("(") (
	// actual-code:5:27: (Ident) rest
	// actual-code:5:32: (Ident) form
	// actual-code:5:36: (")") )
	// actual-code:5:37: (")") )
	// actual-code:6:3: ("(") (
	// actual-code:6:4: (Ident) list
	// actual-code:6:9: (Ident) form
	// actual-code:6:14: (Ident) acc
	// actual-code:6:17: (")") )
	// actual-code:6:18: (")") )
	// actual-code:6:19: (")") )
	// actual-code:6:20: (")") )
}

func Example_mode() {
	const src = `
    ;; Comment begins at column 5

This line should not be included in the output
`

	var s scanner.Scanner
	s.Init(strings.NewReader(src))
	s.Filename = "comments"
	s.Mode ^= scanner.SkipComments // don't skip comments

	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		txt := s.TokenText()
		if strings.HasPrefix(txt, ";") {
			fmt.Printf("%s: %s\n", s.Position, txt)
		}
	}

	// Output:
	// comments:2:5: ;; Comment begins at column 5
}
