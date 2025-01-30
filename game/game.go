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

    // Speeds are blocks per seconds
    fallSpeed float32
    moveSpeed float32
    softDropSpeed float32

    drawGhost bool
}

func NewGame() *Game {
    game := Game{
        bag: newPiecesBag(pieceBagDefaultSize),
        board: newBoard(boardDefaultRows, boardDefaultColumns),

        currentPiece: nil,
        nextPiece: nil,

        fallSpeed: 1.0,
        moveSpeed: 11.0,
        softDropSpeed: 30.0,

        drawGhost: true,
    }

    game.spawnNewPiece()
    return &game
}

func (g *Game) Draw() {
    g.board.draw()
    g.currentPiece.draw(g.drawGhost)
}

func (g *Game) Update() {
    dt := rl.GetFrameTime()

    g.board.update()
    g.updatePiece(g.fallSpeed * dt)

    // Input handle
    if rl.IsKeyDown(rl.KeyLeft) || rl.IsKeyDown(rl.KeyA) {
        g.currentPiece.move(-g.moveSpeed * dt)
    }
    if rl.IsKeyDown(rl.KeyRight) || rl.IsKeyDown(rl.KeyD) {
        g.currentPiece.move(g.moveSpeed * dt)
    }
    if rl.IsKeyPressed(rl.KeyUp) || rl.IsKeyPressed(rl.KeyW) {
        g.currentPiece.rotate(true)
    }
    if rl.IsKeyDown(rl.KeyDown) || rl.IsKeyDown(rl.KeyS) {
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

func (g *Game) updatePiece(dy float32) {
    g.currentPiece.y += dy
    g.currentPiece.pos.y = int(g.currentPiece.y)

    if g.currentPiece.isCollision() {
        g.currentPiece.y -= dy
        g.currentPiece.pos.y = int(g.currentPiece.y)

        g.currentPiece.lock()
        g.spawnNewPiece()
    }
}

