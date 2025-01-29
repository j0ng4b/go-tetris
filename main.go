package main

import (
	"github.com/j0ng4b/go-tetris/game"

	"github.com/gen2brain/raylib-go/raylib"
)


func main() {
    rl.InitWindow(game.GAME_WINDOW_WIDTH, game.GAME_WINDOW_HEIGHT, game.GAME_WINDOW_TITLE)
    defer rl.CloseWindow()

    rl.SetTargetFPS(game.GAME_FPS)

    game := game.NewGame()
    for !rl.WindowShouldClose() {
        rl.BeginDrawing()
        rl.ClearBackground(rl.Black)

        game.Draw()
        game.Update()

        rl.EndDrawing()
    }
}

