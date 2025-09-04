# install modd untuk hot reload
go install github.com/cortesi/modd/cmd/modd@latest


# tambahkan bin ke path
`export PATH=$PATH:$(go env GOPATH)/bin`
`source ~/.bashrc`

## windows
`go env GOPATH`
C:\Users\USERNAME\go


# run
modd