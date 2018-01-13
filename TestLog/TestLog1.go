package main

import "log"

func init()  {
	log.SetPrefix("Trace :")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)//显示的格式
}
func main()  {

	log.Println("message")

	log.Fatalln("fatal message")

	log.Panicln("panice message")


}


