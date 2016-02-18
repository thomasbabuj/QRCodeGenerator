# QRCodeGenerator
Learning TDD by building QRCode Generator based on this ebook
https://leanpub.com/golang-tdd

# t.Errorf

This differs from typical asserts and expectations in other languages. In Go test
will continue to execute, even if t.Errorf was called. Because of this we can see
all test failures at once.

# go test

go test command automatically recompiles your code, just like go run
