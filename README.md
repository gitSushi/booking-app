# Following along a tutorial

## Learning Go

> [!NOTE]
> Go the language, not the board game.

[Golang Tutorial for Beginners | Full Go Course](https://www.youtube.com/watch?v=yyUHQIec83I)

## Basic commands

```shell
go help <command>
go run main.go

go run main.go helper.go
go run .
```

## Mental notes

If two files are in the same package, no imports (like in JS) are necessary, however the command to run the code properly is now : `go run main.go helper.go` (or whatever other name you've given).

So if there is a lot of files it would become quite bothersome. Luckily one can run a folder. So if you are in the current directory of your files just run `go run .` (By the way, the README file in the folder isn't an issue to the compiler).