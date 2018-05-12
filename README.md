# gospace
## Installation
Follow this [manuel](https://tecadmin.net/install-go-on-ubuntu/)
- `wget https://dl.google.com/go/go1.10.1.linux-amd64.tar.gz`
- `sudo tar -xvf go1.10.1.linux-amd64.tar.gz`
- `sudo mv go /usr/local`
- `vim ~/.bashrc`
  - `export GOROOT=/usr/local/go`
  - `export GOPATH=$HOME/goworkspace`
  - `export PATH=$GOPATH/bin:$GOROOT/bin:$PATH`
- `go version`: check 
- `go env`
- `go get -u github.com/beego/bee`: install bee
- `go get github.com/tools/godep`: install godep