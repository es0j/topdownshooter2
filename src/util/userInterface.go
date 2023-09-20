package util

import rl "github.com/gen2brain/raylib-go/raylib"



type ButtonScript struct {
	text string

}

func (b ButtonScript) Start(ge *GameEngine){

}
func (b ButtonScript) Update(ge *GameEngine,dTime float64){
	rl.DrawText(b.text, 190, 200, 20, rl.LightGray)
}
func NewButton(message string) *GameObject{
	p := new(GameObject)
	p.AddScript(ButtonScript{message})
	return p
}