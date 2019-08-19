(defvar go-base64-test/test-path
  (f-dirname load-file-name))

(defvar go-base64-test/root-path
  (f-parent go-base64-test/test-path))

(module-load (f-expand "go-base64.so" go-base64-test/root-path))
