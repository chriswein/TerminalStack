# Terminal Stack

This project gives you a lightweight and easy way to store data on a stack. You can use pipeing and argument input.

```bash
ls | push
...
pop 
...
-rw-r--r--  1 christoph  staff      701 20 Feb 11:55 db.go
-rw-r--r--  1 christoph  staff      111 20 Feb 12:34 go.mod
-rw-r--r--  1 christoph  staff      177 20 Feb 12:34 go.sum
-rw-r--r--  1 christoph  staff     1060 20 Feb 13:05 main.go
...
```
You can also use it to store links or text

```bash
push Hello World
pop
Hello World
```

## Requirements

Uses Sqlite

## Building

To buid pop and push from source use the build.sh file or run the following commands:

### Push
```bash
go build -ldflags "-X main.pushorpop=push" -o push;
```

### Pop
```bash
go build -ldflags "-X main.pushorpop=pop" -o pop;
```

## Installing

Create aliases for the program in your shell config file.

On MacOS edit:
~/.zshrc


```bash
alias push=/path/to/your/build/push
alias pop=/path/to/your/build/pop
```
For Powershell use: https://stackoverflow.com/questions/24914589/how-to-create-permanent-powershell-aliases
