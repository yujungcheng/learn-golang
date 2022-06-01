## Study book:
Go in 24 Hours Sams Teach Yourself
My first book for learning golang, but there are lots mistakes in its examples.
So I corrected and tested them.
Also make some improvements.


#### System Version
Ubuntu 18.04.5 LTS
go version go1.15.6 linux/amd64


#### install golang
1. wget latest golang
2. tar -xvf <downloaded golang tar file>
3. mv go /usr/local/
4. setup go env (edit ~/.bashrc file)
```shell
export GOROOT=/usr/local/go
export GOPATH=$HOME/go
export PATH=$PATH:$GOROOT:$GOPATH:$GOROOT/bin:$GOPATH/bin
```

#### build executable file
go build


#### compile and run (no executable file created)
go run <source code file>

#### clean up executable files
go clean

#### format code
gofmt -s -w .

#### go env
```shell
ycheng@nuc:~$ go env
GO111MODULE=""
GOARCH="amd64"
GOBIN=""
GOCACHE="/home/ycheng/.cache/go-build"
GOENV="/home/ycheng/.config/go/env"
GOEXE=""
GOFLAGS=""
GOHOSTARCH="amd64"
GOHOSTOS="linux"
GOINSECURE=""
GOMODCACHE="/home/ycheng/go/pkg/mod"
GONOPROXY=""
GONOSUMDB=""
GOOS="linux"
GOPATH="/home/ycheng/go"
GOPRIVATE=""
GOPROXY="https://proxy.golang.org,direct"
GOROOT="/usr/local/go"
GOSUMDB="sum.golang.org"
GOTMPDIR=""
GOTOOLDIR="/usr/local/go/pkg/tool/linux_amd64"
GCCGO="gccgo"
AR="ar"
CC="gcc"
CXX="g++"
CGO_ENABLED="1"
GOMOD=""
CGO_CFLAGS="-g -O2"
CGO_CPPFLAGS=""
CGO_CXXFLAGS="-g -O2"
CGO_FFLAGS="-g -O2"
CGO_LDFLAGS="-g -O2"
PKG_CONFIG="pkg-config"
GOGCCFLAGS="-fPIC -m64 -pthread -fmessage-length=0 -fdebug-prefix-map=/tmp/go-build380422097=/tmp/go-build -gno-record-gcc-switches"
```

#### compile flags to reduce size of compiled binaries. (omit symbol table,debug info and DWARF symbol table)
go build -ldflags="-s -w" <go file>


#### Produce checksum file
$ go build hello_world.go
$ sha1sum hello_world
24516517554d610749459ebdaaeaca339457977d  hello_world


#### go get (-u option for update dependencies)
go get -u <package url>


#### reference.
https://www.calhoun.io/
https://opensource.com/tags/go-programming-language
https://github.com/mactsouk/opensource.com