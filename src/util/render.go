package util

import rl "github.com/gen2brain/raylib-go/raylib"
import "fmt"

type GameScript interface {
	Start(ge *GameEngine)
	Update(ge *GameEngine,dTime float64)
}

type Transform struct {
	x,y,rot float64
}


type GameObject struct {
	id int
	transform Transform
	scripts []GameScript 
}

func (gameObj *GameObject) AddScript(gs GameScript){
	gameObj.scripts = append(gameObj.scripts,gs)
}


type GameEngine struct {
	objList []GameObject
}


func (ge *GameEngine) MainLoop() {

	rl.InitWindow(800, 450, "raylib2 [core] example - basic window")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)
	for !rl.WindowShouldClose() {

		//x := rl.GetMousePosition()
		//fmt.Println(x.X)

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		for _,obj :=range(ge.objList){
			//fmt.Println("Update")
			//fmt.Printf("%v", obj)
			for _,s :=range(obj.scripts){
				s.Update(ge,1.0)
			}
		}

		rl.EndDrawing()
	}
}

func (ctx  *GameEngine) AddObject(gameObj *GameObject) {
	ctx.objList = append(ctx.objList,*gameObj)
	fmt.Println("Added")
}