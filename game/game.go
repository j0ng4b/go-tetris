package game

import (
	"time"

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

    // Time before lock piece in seconds
    lockDelay float32
    lockLastTime time.Time

    drawGhost bool

    level int
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

        lockDelay: 0.6,
        lockLastTime: time.Time{},

        drawGhost: true,

        level: 0,
    }

    game.spawnNewPiece()
    return &game
}

func (g *Game) Draw() {
    g.board.draw()

    g.currentPiece.draw(g.drawGhost)
    g.nextPiece.draw(false)
}

func (g *Game) Update() {
    dt := rl.GetFrameTime()

    g.board.update()
    g.updateGravity()

    g.updatePiece(g.fallSpeed * dt)
    g.updatePieceLock()

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
            if g.lockLastTime.IsZero() {
                g.lockLastTime = time.Now()
            }
        } else {
            g.lockLastTime = time.Time{}
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

    g.nextPiece.y = 0
    g.nextPiece.pos.y = 0

    g.currentPiece = g.nextPiece
    g.nextPiece = newPiece(g.bag.next(), g.board)

    g.nextPiece.y -= 4
    g.nextPiece.pos.y -= 4
}

func (g *Game) updatePiece(dy float32) {
    g.currentPiece.y += dy
    g.currentPiece.pos.y = int(g.currentPiece.y)

    if g.currentPiece.isCollision() {
        g.currentPiece.y -= dy
        g.currentPiece.pos.y = int(g.currentPiece.y)

        if g.lockLastTime.IsZero() {
            g.lockLastTime = time.Now()
        }
    } else {
        g.lockLastTime = time.Time{}
    }
}

func (g *Game) updateGravity() {
    var frameDelay float32

    if g.level = g.board.clearedRows / 10; g.level >= 29 {
        frameDelay = 1
    } else {
        frameDelay = []float32{
            48, 43, 38, 33, 28, 23, 18, 13, 8, 6, // Levels 0–9
            5, 5, 5, 5, 5, 5, 5, 5, 5, 5,         // Levels 10–19
            5, 5, 5, 5, 5, 5, 5, 5, 5, 5,         // Levels 20–29
        }[g.level]
    }

    g.fallSpeed = (1.0 / frameDelay) * GAME_FPS
}

func (g *Game) updatePieceLock() {
    if g.lockLastTime.IsZero() {
        return
    }

    lockTime := time.Now().Sub(g.lockLastTime)
    if lockTime.Seconds() < float64(g.lockDelay) {
        return
    }

    g.currentPiece.lock()
    g.spawnNewPiece()

    g.lockLastTime = time.Time{}
}

