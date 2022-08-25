# gocp

Simple file copy implementation.
This is an exercise to write go application.
Developed in go version 1.17.1


#### usage
```
$ go run main.go -help
$ go run main.go <source file path> <destination directory/file path>
```

#### todo features:
- changeable copy speed limit - initial implemented
- multiple files copy, able to copy one to many with naming format
- able to pause and resume by pressing specific key
- open socket for querying copy process, pid, control from remote server
- direct IO - done


#### References
```
https://zetcode.com/golang/copyfile/
https://www.codetd.com/en/article/8276908
https://www.pixelstech.net/article/1596946473-A-simple-example-on-implementing-progress-bar-in-GoLang
https://segmentfault.com/a/1190000023375330
https://segmentfault.com/a/1190000020927821
https://pkg.go.dev/github.com/ncw/directio

https://eklitzke.org/efficient-file-copying-on-linux
```