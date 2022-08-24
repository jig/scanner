// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package scanner_test

import (
	"fmt"
	"strings"

	"github.com/jig/scanner"
)

func Example() {
	const src = `; This is scanned code
(def a '(list 1 2 "hel\"lo" ¬hel¬¬lo¬ 3.14))
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
	// example:2:15: (Int) 1
	// example:2:17: (Int) 2
	// example:2:19: (String) "hel\"lo"
	// example:2:29: (RawString) ¬hel¬¬lo¬
	// example:2:39: (Float) 3.14
	// example:2:43: (")") )
	// example:2:44: (")") )

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
