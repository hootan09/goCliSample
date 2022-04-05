### go CLI Application

Up & Running
```sh
#Install cobra-cli
$ go install github.com/spf13/cobra-cli@latest
#Create Project
$ go mod init go_cli
$ cobra-cli init .
#Adding new Command for cli
$ cobra-cli add myip

#runinng
$ go run . getjoke
$ go run . myip
```

OutPut:
```sh
$ go run . myip   
IP Addres: 2.181.44.46
Country Code: IR

$ go run . getjoke
A steak pun is a rare medium well done.
```