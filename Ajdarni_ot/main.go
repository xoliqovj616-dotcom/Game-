package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	ScreenX = 1600
	ScreenY = 1100
)

type Gun struct {
	pos   rl.Vector2
	speed float32
}
type Ajdar struct {
	texture rl.Texture2D
	pos     rl.Vector2
	speed   float32
}

var gun []Gun
var ajdar []Ajdar
var GAmeOver bool
var count = 0

func main() {
	rl.InitWindow(ScreenX, ScreenY, "GAME")
	rl.InitAudioDevice()
	rl.SetTargetFPS(60)
	odamcha := rl.LoadTexture("canvas.png")
	dragon := rl.LoadTexture("canvas(1).png")
	music := rl.LoadMusicStream("muzic.mp3")

	rl.PlayMusicStream(music)
	ajdar = append(ajdar,
		Ajdar{dragon, rl.Vector2{X: ScreenX, Y: 100}, 10},
		Ajdar{dragon, rl.Vector2{X: ScreenX, Y: 300}, 10},
		Ajdar{dragon, rl.Vector2{X: ScreenX, Y: 500}, 10},
		Ajdar{dragon, rl.Vector2{X: ScreenX, Y: 700}, 10},
		Ajdar{dragon, rl.Vector2{X: ScreenX, Y: 900}, 10},
	)
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.White)

		mouse := rl.GetMousePosition()
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			guns := Gun{
				pos:   mouse,
				speed: 10,
			}
			gun = append(gun, guns)
		}
		rl.UpdateMusicStream(music)
		for i := range gun {
			gun[i].pos.X += gun[i].speed
			rl.DrawCircleV(gun[i].pos, 5, rl.Black)
		}
		for i := range ajdar {
			ajdar[i].pos.X -= ajdar[i].speed
			if ajdar[i].pos.X < 0 {

				GAmeOver = true
			}

			rl.DrawTextureEx(ajdar[i].texture, ajdar[i].pos, 0, 0.4, rl.White)
		}
		//tuqnashuv
		for i := range ajdar {
			for j := range gun {
				rect := rl.Rectangle{
					X:      ajdar[i].pos.X,
					Y:      ajdar[i].pos.Y,
					Width:  float32(ajdar[i].texture.Width) * 0.4,
					Height: float32(ajdar[i].texture.Height) * 0.4,
				}
				if rl.CheckCollisionCircleRec(gun[j].pos, 5, rect) {
					ajdar[i].pos.X = ScreenX
					count++
					gun = append(gun[:j], gun[j+1:]...)
					break
				}

			}
		}
		text := fmt.Sprint(count)
		rl.DrawText(text, 1500, 60, 50, rl.Red)
		rl.DrawTextureEx(odamcha, mouse, 0, 0.4, rl.White)

		if GAmeOver {
			for i := range ajdar {
				for j := range gun {
					ajdar[i].speed = 0
					gun[j].speed = 0
				}
			}
			rl.StopMusicStream(music)
			rl.DrawText("GAME OVER", ScreenX/2, ScreenY/2, 50, rl.Red)
		}

		rl.EndDrawing()
	}
}
