package game

import rl "github.com/gen2brain/raylib-go/raylib"

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

var (
    piecesShapes = [countPieceShapeType][4]pos{
        { {0, 1}, {1, 1}, {2, 1}, {3, 1} }, // iPieceShapeType
        { {1, 0}, {2, 0}, {1, 1}, {2, 1} }, // oPieceShapeType
        { {1, 0}, {0, 1}, {1, 1}, {2, 1} }, // tPieceShapeType
        { {0, 0}, {0, 1}, {1, 1}, {2, 1} }, // jPieceShapeType
        { {2, 0}, {0, 1}, {1, 1}, {2, 1} }, // lPieceShapeType
        { {1, 0}, {2, 0}, {0, 1}, {1, 1} }, // sPieceShapeType
        { {0, 0}, {1, 0}, {1, 1}, {2, 1} }, // zPieceShapeType
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
)

type pieceShapeType byte

type pos struct {
    x float32
    y float32
}

type piece struct {
    pos
    shape pieceShapeType

    *board
}

func newPiece(shape pieceShapeType, b *board) *piece {
    p := piece{
        pos: pos{
            x: float32(b.columns / 2 - 2),
            y: 0,
        },

        shape: shape,

        board: b,
    }

    return &p
}

func (p *piece) draw() {
    for _, cell := range piecesShapes[p.shape] {
        rl.DrawRectangle(
            int32(p.board.offsetX) + int32((p.x + cell.x) * boardCellPixels),
            int32(p.board.offsetY) + int32((p.y + cell.y) * boardCellPixels),
            boardCellPixels,
            boardCellPixels,
            piecesColors[p.shape],
        )
    }
}

func (p *piece) update() {

}

