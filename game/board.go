package game

import (
    "github.com/gen2brain/raylib-go/raylib"
)

const (
    boardDefaultRows = 20
    boardDefaultColumns = 10

    boardCellPixels = 20
)

type boardCell struct {
    empty bool
    shape pieceShapeType
}

type board struct {
    rows int
    columns int
    hiddenRows int

    offsetX int
    offsetY int

    cells [][]boardCell
}

func newBoard(rows, columns int) *board {
    b := board{
        rows: rows,
        columns: columns,
        hiddenRows: 2,

        offsetX: (rl.GetScreenWidth() - columns * boardCellPixels) / 2,
        offsetY: (rl.GetScreenHeight() - rows * boardCellPixels) / 2,
    }
    b.rows += b.hiddenRows

    b.cells = make([][]boardCell, b.rows)
    for row := 0; row < b.rows; row++ {
        b.cells[row] = make([]boardCell, columns)
        for column := 0; column < columns; column++ {
            b.cells[row][column].empty = true;
        }
    }

    return &b
}

func (b *board) draw() {
    for row := b.hiddenRows; row < b.rows; row++ {
        for column := 0; column < b.columns; column++ {
            if b.cells[row][column].empty {
                continue
            }

            rl.DrawRectangle(
                int32(b.offsetX + column * boardCellPixels),
                int32(b.offsetY + row * boardCellPixels),
                boardCellPixels,
                boardCellPixels,
                piecesColors[b.cells[row][column].shape],
            )

            rl.DrawRectangleLines(
                int32(b.offsetX + column * boardCellPixels),
                int32(b.offsetY + row * boardCellPixels),
                boardCellPixels,
                boardCellPixels,
                rl.Gray,
            )
        }
    }

    rl.DrawRectangleLines(
        int32(b.offsetX),
        int32(b.offsetY),
        int32(b.columns * boardCellPixels),
        int32(b.rows * boardCellPixels),
        rl.White,
    )
}

func (b *board) update() {
    for row := b.rows - 1; row >= 0; row-- {
        if b.isRowEmpty(row, 0) {
            break;
        }

        if b.isRowFull(row, 0) {
            b.cleanRow(row)
            b.fallRows(row - 1)
        }
    }
}

func (b *board) isRowEmpty(row, column int) bool {
    if column == b.columns {
        return true
    } else if b.cells[row][column].empty {
        return b.isRowEmpty(row, column + 1)
    }

    return false
}

func (b *board) isRowFull(row, column int) bool {
    if column == b.columns {
        return true
    } else if !b.cells[row][column].empty {
        return b.isRowFull(row, column + 1)
    }

    return false
}

func (b *board) cleanRow(row int) {
    for column := 0; column < b.columns; column++ {
        b.cells[row][column].empty = true
    }
}

func (b *board) fallRows(row int) {
    if row < 0 {
        return
    }

    emptyRow := true
    for column := 0; column < b.columns; column++ {
        if emptyRow && !b.cells[row][column].empty {
            emptyRow = false
        }

        b.cells[row + 1][column] = b.cells[row][column]
    }

    if !emptyRow {
        b.fallRows(row - 1)
    }
}

