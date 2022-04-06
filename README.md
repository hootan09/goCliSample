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

$ go run . getjoke --term=hipster
How did the hipster burn the roof of his mouth? He ate the pizza before it was cool.
```

We will be able to distribute this **go_cli** tool as a Go package. In order to do this:

- Upload your code to a public repo
- Install it by running **go get \<link-to-your-public-repo\>** e.g. go get github.com/example/go_cli

Then you will be able to run your tool like this:
```sh
go_cli random

# random joke with term
go_cli getjoke --term=hipster
```