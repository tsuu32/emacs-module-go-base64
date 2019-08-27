;;; go-base64-test.el --- Tests for gmodule

(ert-deftest go-base64-test/b64-encode-ascii ()
  (let ((actual (b64-encode "hello"))
        (expected "aGVsbG8="))
    (should (equal actual expected))))

(ert-deftest go-base64-test/b64-encode-japanese ()
  (let ((actual (b64-encode "こんにちは"))
        (expected "44GT44KT44Gr44Gh44Gv"))
    (should (equal actual expected))))

(ert-deftest go-base64-test/b64-encode-emoji ()
  (let ((actual (b64-encode "🤔"))
        (expected "8J+klA=="))
    (should (equal actual expected))))

(ert-deftest go-base64-test/b64-decode-ascii ()
  (let ((actual (b64-decode "d29ybGQh"))
        (expected "world!"))
    (should (equal actual expected))))

(ert-deftest go-base64-test/b64-decode-chinese ()
  (let ((actual (b64-decode "5L2g5aW9"))
        (expected "你好"))
    (should (equal actual expected))))

(ert-deftest go-base64-test/b64-decode-emoji ()
  (let ((actual (b64-decode "8J+krg=="))
        (expected "🤮"))
    (should (equal actual expected))))

;;; go-base64-test.el ends here
