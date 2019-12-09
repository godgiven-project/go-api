# GO INSTALLATION
for download go lang go to the [go website](https://golang.org/dl/) and for install read [this url](https://golang.org/doc/install)<br/>
after installing Go Lang, you should set env variables as follow:
```
export PATH=$PATH:/usr/local/go/bin
export GOPATH=$HOME/go
export GO111MODULE=on
```
-------------
remember to create `/go` under your `home` directory ofcourse :D

As you started working create main.go basically under `/home/YourUserName/go/src/YourProjectDirectory` then run it via `go run main.go` command.<br/>

**notice:** After one run command, by default the process itself creates a file named `go.mod` beside `main.go` that contains dependencies and requirements of your project.<br/>

**tip-1:** If you wanted to install something with `go get PACKAGE_NAME` command consider that, first it is better be under `/home/YourUserName/go/` then run command. 
second, make sure `curl` and `apt-transport-https` packages installed on your PC otherwise you might face SSL and authorization error.<br/>
