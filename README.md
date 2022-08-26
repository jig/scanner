Lisp lexer used by [`jig/lisp`](https://github.com/jig/lisp).

Code adapted from the Go `text/scanner` standard package, and mostly compatible with it (beside the fact this is for Lisp and that is for Go syntax).

Major differences:
- identifier allowed characters is more expansive than Go is (e.g. `+`, `-`, `*host-name*`, or `<=` are valid identifiers)
- negative integers or negative floats are supported as a single token (of type `Int` or `Float` respectively)
- character tokens are not supported anymore (e.g. `'A'`) as `'` is used as synonym of `quote`
- parsing errors are not printed to `os.Stdout` by default
- raw strings in `jig/lisp` are quoted with `¬` character (instead of `` ` ``) as `` ` `` is used as synonym of `quasiquote`. Raw strings might include `¬` by doublind them `¬¬`
- `#{` is specially handled to support Lisp set literals
- `~`, `@` and `~@` are specially handled to support them as synonyms of `unquote`, `deref` and `splice-unquote` respectively
- support of Lisp keywords (e.g. `:key`) as a single token of type `Keyword`
