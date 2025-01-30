package game

import (
	"github.com/gen2brain/raylib-go/raylib"
)

const (
    iPieceShapeType pieceShapeType = iota
    oPieceShapeType
    tPieceShapeType
    jPieceShapeType
    lPieceShapeType
    sPieceShapeType
    zPieceShapeType
    countPieceShapeType
)

const (
    pieceRotation0 pieceRotation = iota
    pieceRotation90
    pieceRotation180
    pieceRotation270
    countPieceRotations
)

var (
    piecesShapes = [countPieceShapeType][4][4]pos{
        { // iPieceShapeType
            { {0, 1}, {1, 1}, {2, 1}, {3, 1} }, // 0°
            { {2, 0}, {2, 1}, {2, 2}, {2, 3} }, // 90°
            { {0, 2}, {1, 2}, {2, 2}, {3, 2} }, // 180°
            { {1, 0}, {1, 1}, {1, 2}, {1, 3} }, // 270°
        },

        { // oPieceShapeType
            { {1, 0}, {2, 0}, {1, 1}, {2, 1} }, // 0°
        },

        { // tPieceShapeType
            { {1, 0}, {0, 1}, {1, 1}, {2, 1} }, // 0°
            { {1, 0}, {1, 1}, {2, 1}, {1, 2} }, // 90°
            { {0, 1}, {1, 1}, {2, 1}, {1, 2} }, // 180°
            { {1, 0}, {0, 1}, {1, 1}, {1, 2} }, // 270°
        },

        { // jPieceShapeType
            { {0, 0}, {0, 1}, {1, 1}, {2, 1} }, // 0°
            { {1, 0}, {2, 0}, {1, 1}, {1, 2} }, // 90°
            { {0, 1}, {1, 1}, {2, 1}, {2, 2} }, // 180°
            { {1, 0}, {1, 1}, {0, 2}, {1, 2} }, // 270°
        },

        { // lPieceShapeType
            { {2, 0}, {0, 1}, {1, 1}, {2, 1} }, // 0°
            { {1, 0}, {1, 1}, {1, 2}, {2, 2} }, // 90°
            { {0, 1}, {1, 1}, {2, 1}, {0, 2} }, // 180°
            { {0, 0}, {1, 0}, {1, 1}, {1, 2} }, // 270°
        },

        { // sPieceShapeType
            { {1, 0}, {2, 0}, {0, 1}, {1, 1} }, // 0°
            { {1, 0}, {1, 1}, {2, 1}, {2, 2} }, // 90°
            { {1, 1}, {2, 1}, {0, 2}, {1, 2} }, // 180°
            { {0, 0}, {0, 1}, {1, 1}, {1, 2} }, // 270°
        },

        { // zPieceShapeType
            { {0, 0}, {1, 0}, {1, 1}, {2, 1} }, // 0°
            { {2, 0}, {1, 1}, {2, 1}, {1, 2} }, // 90°
            { {0, 1}, {1, 1}, {1, 2}, {2, 2} }, // 180°
            { {1, 0}, {0, 1}, {1, 1}, {0, 2} }, // 270°
        },
    }

    piecesColors = [countPieceShapeType]rl.Color{
        { R: 22,  G: 214, B: 252, A: 255 }, // iPieceShapeType
        { R: 252, G: 202, B: 22,  A: 255 }, // oPieceShapeType
        { R: 150, G: 22,  B: 252, A: 255 }, // tPieceShapeType
        { R: 40,  G: 90,  B: 220, A: 255 }, // jPieceShapeType
        { R: 250, G: 140, B: 10,  A: 255 }, // lPieceShapeType
        { R: 16,  G: 210, B: 40,  A: 255 }, // sPieceShapeType
        { R: 220, G: 40,  B: 40,  A: 255 }, // zPieceShapeType
    }

    pieceWallKick = [countPieceRotations][5]pos {
        { {0, 0}, {-1, 0}, {-1, +1}, {0, -2}, {-1, -2}, }, // 0° → 90°
        { {0, 0}, {+1, 0}, {+1, -1}, {0, +2}, {+1, +2}, }, // 90° → 180°
        { {0, 0}, {+1, 0}, {+1, +1}, {0, -2}, {+1, -2}, }, // 180° → 270°
        { {0, 0}, {-1, 0}, {-1, -1}, {0, +2}, {-1, +2}, }, // 270° → 0°
    }

    iPieceWallKick = [countPieceRotations][5]pos {
        { {0, 0}, {-2, 0}, {+1, 0}, {-2, -1}, {+1, +2}, }, // 0° → 90°
        { {0, 0}, {-1, 0}, {+2, 0}, {-1, +2}, {+2, -1}, }, // 90° → 180°
        { {0, 0}, {+2, 0}, {-1, 0}, {+2, +1}, {-1, -2}, }, // 180° → 270°
        { {0, 0}, {+1, 0}, {-2, 0}, {+1, -2}, {-2, +1}, }, // 270° → 0°
    }
)

type pieceShapeType byte
type pieceRotation byte

type pos struct {
    x int
    y int
}

type piece struct {
    pos pos

    // Real piece position, used on move
    x float32
    y float32

    shape pieceShapeType
    rotation pieceRotation

    *board
}

func newPiece(shape pieceShapeType, b *board) *piece {
    p := piece{
        pos: pos{
            x: b.columns / 2 - 2,
            y: 0,
        },

        shape: shape,
        rotation: pieceRotation0,

        board: b,
    }

    p.x = float32(p.pos.x)
    p.y = float32(p.pos.y)

    return &p
}

func (p *piece) draw() {
    for _, block := range piecesShapes[p.shape][p.rotation] {
        rl.DrawRectangle(
            int32(p.board.offsetX) + int32((p.pos.x + block.x) * boardCellPixels),
            int32(p.board.offsetY) + int32((p.pos.y + block.y) * boardCellPixels),
            boardCellPixels,
            boardCellPixels,
            piecesColors[p.shape],
        )
    }
}

func (p *piece) move(dx float32) {
    p.x += dx
    p.pos.x = int(p.x)

    if p.isCollision() {
        p.x -= dx
        p.pos.x = int(p.x)
    }
}

func (p *piece) rotate(clockwise bool) {
    if p.shape == oPieceShapeType {
        return
    }

    oldRotation := p.rotation

    switch p.rotation {
    case pieceRotation0:
        p.rotation = pieceRotation90
        if !clockwise {
            p.rotation = pieceRotation270
        }

    case pieceRotation90:
        p.rotation = pieceRotation180
        if !clockwise {
            p.rotation = pieceRotation0
        }

    case pieceRotation180:
        p.rotation = pieceRotation270
        if !clockwise {
            p.rotation = pieceRotation90
        }

    case pieceRotation270:
        p.rotation = pieceRotation0
        if !clockwise {
            p.rotation = pieceRotation180
        }
    }

    wallKicks := pieceWallKick[oldRotation]
    if p.shape == iPieceShapeType {
        wallKicks = iPieceWallKick[oldRotation]
    }

    for _, kick := range wallKicks {
        p.pos.x += kick.x
        p.pos.y += kick.y

        if !p.isCollision() {
            p.x = float32(p.pos.x)
            p.y = float32(p.pos.y)

            return
        }

        p.pos.x -= kick.x
        p.pos.y -= kick.y
    }

    p.rotation = oldRotation
}

func (p *piece) lock() {
    for _, block := range piecesShapes[p.shape][p.rotation] {
        x := p.pos.x + block.x
        y := p.pos.y + block.y

        p.board.cells[y][x].empty = false
        p.board.cells[y][x].shape = p.shape
    }
}

func (p *piece) softDrop(dy float32) bool {
    p.y += dy
    p.pos.y = int(p.y)

    if p.isCollision() {
        p.y -= 1
        p.pos.y = int(p.y)
        return false
    }

    return true
}

func (p *piece) hardDrop() {
    for p.softDrop(1) {}
    p.lock()
}

func (p *piece) isCollision() bool {
    for _, block := range piecesShapes[p.shape][p.rotation] {
        x := p.pos.x + block.x
        y := p.pos.y + block.y

        if x < 0 || x >= p.board.columns || y >= p.board.rows {
            return true
        }

        if !p.board.cells[y][x].empty {
            return true
        }
    }

    return false
}

