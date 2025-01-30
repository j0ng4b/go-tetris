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
    bag *piecesBag
    board *board

    currentPiece *piece
    nextPiece *piece

    moveSpeed float32
    softDropSpeed float32
}

func NewGame() *Game {
    game := Game{
        bag: newPiecesBag(pieceBagDefaultSize),
        board: newBoard(boardDefaultRows, boardDefaultColumns),

        currentPiece: nil,
        nextPiece: nil,

        moveSpeed: 15.0,
        softDropSpeed: 30.0,
    }

    game.spawnNewPiece()
    return &game
}

func (g *Game) Draw() {
    g.board.draw()
    g.currentPiece.draw()
}

func (g *Game) Update() {
    dt := rl.GetFrameTime()

    g.board.update()

    // Input handle
    if rl.IsKeyDown(rl.KeyLeft) {
        g.currentPiece.move(-g.moveSpeed * dt)
    }
    if rl.IsKeyDown(rl.KeyRight) {
        g.currentPiece.move(g.moveSpeed * dt)
    }
    if rl.IsKeyPressed(rl.KeyUp) {
        g.currentPiece.rotate(true)
    }
    if rl.IsKeyDown(rl.KeyDown) {
        if !g.currentPiece.softDrop(g.softDropSpeed * dt) {
            g.currentPiece.lock()
            g.spawnNewPiece()
        }
    }
    if rl.IsKeyPressed(rl.KeySpace) {
        g.currentPiece.hardDrop()
        g.spawnNewPiece()
    }
}

func (g *Game) spawnNewPiece() {
    if g.nextPiece == nil {
        g.nextPiece = newPiece(g.bag.next(), g.board)
    }

    g.currentPiece = g.nextPiece
    g.nextPiece = newPiece(g.bag.next(), g.board)
}

