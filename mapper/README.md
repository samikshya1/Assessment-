# Aspipration mapper

## RUN-PROGRAM

`make run-program`

----

output:
```shell
go run cmd/main.go
asPirAtiOn.cOm
```

## RUN-TESTS
`make run-test`

---
output:

```shell
go test -v ./...
=== RUN   TestCapitalizeEveryThirdAlphanumericChar
=== RUN   TestCapitalizeEveryThirdAlphanumericChar/test_blabla
=== RUN   TestCapitalizeEveryThirdAlphanumericChar/test_BLaBLa
=== RUN   TestCapitalizeEveryThirdAlphanumericChar/test_Aspiration.com
=== RUN   TestCapitalizeEveryThirdAlphanumericChar/test_special#12#123test
--- PASS: TestCapitalizeEveryThirdAlphanumericChar (0.00s)
    --- PASS: TestCapitalizeEveryThirdAlphanumericChar/test_blabla (0.00s)
    --- PASS: TestCapitalizeEveryThirdAlphanumericChar/test_BLaBLa (0.00s)
    --- PASS: TestCapitalizeEveryThirdAlphanumericChar/test_Aspiration.com (0.00s)
    --- PASS: TestCapitalizeEveryThirdAlphanumericChar/test_special#12#123test (0.00s)
PASS
ok  	github.com/samikshya/mapper	(cached)
?   	github.com/samikshya/mapper/cmd	[no test files]

```
