package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	width  = 800
	height = 600
)

type Sprite struct {
	Pos   rl.Vector2
	Speed rl.Vector2
	Color rl.Color
}

type Game struct {
	ball   []Sprite
	player Sprite
}

func (g *Game) Init() {
	g.ball = []Sprite{
		{
			Pos:   rl.Vector2{X: float32(width / 2), Y: float32(height / 2)},
			Color: rl.Black,
		},
		{
			Pos:   rl.Vector2{X: float32(width / 4), Y: float32(height / 4)},
			Color: rl.Magenta,
		},
	}

	g.player = Sprite{
		Pos:   rl.Vector2{X: float32(width / 6), Y: float32(height / 6)},
		Color: rl.Yellow,
	}
}

func main() {
	rl.InitWindow(width, height, "Raylib-go test")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	game := Game{}
	game.Init()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.DrawText("Congrats, Raylib-go First created", 190, 200, 20, rl.LightGray)

		rl.DrawRectangleV(game.player.Pos, rl.Vector2{X: 50, Y: 50}, game.player.Color)

		for _, ball := range game.ball {
			rl.DrawCircleV(ball.Pos, 20, ball.Color)
		}

		rl.EndDrawing()
	}
}
