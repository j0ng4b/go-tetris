package game

import (
    "github.com/gen2brain/raylib-go/raylib"
)

const (
    GAME_WINDOW_TITLE = "Go-tetris - Tetris clone on Go"

    GAME_WINDOW_WIDTH = 800
    GAME_WINDOW_HEIGHT = 600

    GAME_FPS = 60
)

type Game struct {
    *board

    currentPiece *piece
    nextPiece *piece
}

func NewGame() *Game {
    game := Game{
        board: newBoard(boardDefaultRows, boardDefaultColumns),
    }

    game.currentPiece = newPiece(oPieceShapeType, game.board)

    return &game
}

func (game *Game) Draw() {
    game.board.draw()
    game.currentPiece.draw()
}

func (game *Game) Update() {
    game.board.update()
}

