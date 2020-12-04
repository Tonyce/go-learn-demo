# go-learn-demo

## Debug

**VS Code**

## Test

![](./unit_test_tweet.png)

- unit test
  - ~~[x] method mock `go test -tags mock`。 `+build mock / !mock` 见 `user model`~~
  - [x] gomock
  - [https://deployeveryday.com/2019/10/08/golang-auth-mock.html](https://deployeveryday.com/2019/10/08/golang-auth-mock.html)
- integrated test
  - io mock
    - [x] http mock [https://www.ru-rocker.com/2017/07/10/how-to-unit-test-and-mocking-a-function-in-go/](https://www.ru-rocker.com/2017/07/10/how-to-unit-test-and-mocking-a-function-in-go/)
    - [ ] file mock
  - before / after 
    - [http://cs-guy.com/blog/2015/01/test-main/](http://cs-guy.com/blog/2015/01/test-main/)
    - [https://deepzz.com/post/study-golang-test.html](https://deepzz.com/post/study-golang-test.html)

## benchmark

## custom error

[https://itnext.io/golang-error-handling-best-practice-a36f47b0b94c](https://itnext.io/golang-error-handling-best-practice-a36f47b0b94c)

## middleware

## logger

## other

- [https://talks.golang.org/2012/10things.slide#12](https://talks.golang.org/2012/10things.slide#12)

- embed
  - [https://travix.io/type-embedding-in-go-ba40dd4264df](https://travix.io/type-embedding-in-go-ba40dd4264df)
  - [https://eli.thegreenplace.net/2020/embedding-in-go-part-1-structs-in-structs/](https://eli.thegreenplace.net/2020/embedding-in-go-part-1-structs-in-structs/)
  
## git hook

- [https://medium.com/@radlinskii/writing-the-pre-commit-git-hook-for-go-files-810f8d5f1c6f](https://medium.com/@radlinskii/writing-the-pre-commit-git-hook-for-go-files-810f8d5f1c6f)

## grpc

```shell
$ protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/helloworld.proto
```

```shell
😒  ~/17zuoye/go-examples/logical-example/proto > protoc --go_out=paths=source_relative:../internal/pb helloworld.proto
😒  ~/17zuoye/go-examples/logical-example/proto > protoc --go-grpc_out=paths=source_relative:../internal/pb helloworld.proto
 ```

## Cross build

[参考 https://studygolang.com/articles/14376](https://studygolang.com/articles/14376)

```bash
$ CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
...
```
