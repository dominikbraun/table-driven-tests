# Table-Driven Tests

> A simple example for table-driven unit tests in Go.

This repository contains a simple [user storage](storage.go) implementation and provides an example for [testing the storage](storage_test.go) using the table-driven test approach. You can experiment with the test cases yourself by cloning the repository, running the tests locally, erroneously changing the expected values, and running them again.

## Running the tests

You can run the tests using `go test -v ./...`.

```
=== RUN   TestStorage_FindUser
=== RUN   TestStorage_FindUser/user_doesn't_exist
=== RUN   TestStorage_FindUser/user_exists
--- PASS: TestStorage_FindUser (0.00s)
    --- PASS: TestStorage_FindUser/user_doesn't_exist (0.00s)
    --- PASS: TestStorage_FindUser/user_exists (0.00s)
PASS
ok      github.com/dominikbraun/table-driven-tests      0.112s
```

After changing an expected value – for example wrongly changing the [expected user name](https://github.com/dominikbraun/table-driven-tests/blob/dc13113ab3276a4c701efcf1cde44261baff8853/storage_test.go#L25) in the first test case from `user #2` to `user #3` – running the tests again will cause them to fail.

```
=== RUN   TestStorage_FindUser
=== RUN   TestStorage_FindUser/user_exists
    storage_test.go:58: user expectancy doesn't match:   main.User{
                ID:   2,
        -       Name: "User #3",
        +       Name: "User #2",
          }
=== RUN   TestStorage_FindUser/user_doesn't_exist
--- FAIL: TestStorage_FindUser (0.00s)
    --- FAIL: TestStorage_FindUser/user_exists (0.00s)
    --- PASS: TestStorage_FindUser/user_doesn't_exist (0.00s)
FAIL
FAIL    github.com/dominikbraun/table-driven-tests      0.241s
FAIL
```

The line starting with `-` shows the expected value, while the line starting with `+` shows the actual value.
