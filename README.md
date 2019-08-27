# emacs-module-go-base64

This package provides `go-base64-encode` and `go-base64-decode` functions that use go base64 library.

This requires emacs built with `--with-modules`.

## Build
```sh
make
```

## Usage
```sh
emacs -L .
```

and

```emacs-lisp
(require 'go-base64)

(go-base64-encode "hello")
;; => "aGVsbG8="

(go-base64-decode "aGVsbG8=")
;; => "hello"
```

## note
`go-base64-encode` is slower than `base64-encode-string`...🤯

Also `go-base64-decode` is slower than `base64-decode-string`...😱

`go-base64-encode` and `go-base64-decode` allow multibyte characters🤗

```emacs-lisp
(base64-encode-string "🤔")
;; => error

(base64-encode-string (encode-coding-string "🤔" 'raw-text))
;; => "8J+klA=="

(base64-decode-string "8J+klA==")
;; => "\360\237\244\224"

(string-as-multibyte (base64-decode-string "8J+klA=="))
;; => "🤔"

(go-base64-encode "🤔")
;; => "8J+klA=="

(go-base64-decode "8J+klA==")
;; => "🤔"
```
