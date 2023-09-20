package main


//import "fmt"
import "github.com/es0j/topdownshooter2/util"

func main() {


	ge := new(util.GameEngine);
	ge.AddObject(util.NewButton("hello world"))
	ge.MainLoop()


}