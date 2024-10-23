# GO LEARNING DAY1

### why go even required?
- I reason of learning go is that, the go language provide its developers the fast compilation of the code and fast deployment. As compared to c++ which takes which may have a faster code execution but long compilation of code.
- Go is static type language and can be run for the backend development also.
- Go provides its developers to deploy the product on production faster than other languages.
- Go have the good amount of libraries avaiable on the open source for developers to use.
- Go provide easy way to do threading via using `Goroutines`.

### how go works?
- Go is a static type programming language and which need to compile on the system, and after the compilation complete we can directly share the result or software to any system without requiring any kind go pre installation.
- Go is a programming language which can handle the web production backend as well as terminal based application to build the product, up and running.

### how does go install on linux? [GO OFFICIAL DOCUMENTATION] (https://go.dev/doc/install)
- To install the go programming language in linux we need to run commands on the terminal.

- This command check if go is avaiable in system on not. If the go exists already then it removes and install the fresh version of go, else it just install :).
`rm -rf /usr/local/go && tar -C /usr/local -xzf go1.23.2.linux-amd64.tar.gz`

- After go install we need to add the GO directory to the environment variables, this is done because of easy access of GO, else we need to always provide the path of the go directory.
`export PATH=$PATH:/usr/local/go/bin`

- After that we just need to check if the go installed or not
`go version`

```
#NOTE: if any commnad not work just check the linux on the heading which refers to the GO official documentations.
```

### code declaration & code strucutre for go..?
- Code declaration is the way to write the code any programming language.
- Go provide 3 things before starting the code in code declaration.
 - 1.) we require the package name if the package name is main then its the first file which execute or else you can say the entry point of the code execution.
 - 2.) we require the packages to import that are required! like fmt for formating & printing like iostream in c++ for std::cout or std::cin;
 - 3.) the main function, main package require the main function from where all the code instructions & code steps written!

### creating hello world program in go
- To create a hello world program in go we first need to create the folder for it.
- After the folder we need to initialize the go module.
`go mod init example/hello`

```
#NOTE: Go modules is nothing but a directory which manages your code, also these module help tracking an dependency for the code.
```

- After everything setup for the go code to run we can now write the code which again going to require the 3 code declaration which are- package name, import files, and main function
- Once code is done we can type `go run .` in the terminal to run the code.
