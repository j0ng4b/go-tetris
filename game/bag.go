package game

import (
    "math/rand"
)

const (
    pieceBagDefaultSize byte = 7
)

type piecesBag struct {
    pieces []pieceShapeType
    size byte
    cur byte
}

func newPiecesBag(size byte) *piecesBag {
    pb := piecesBag{
        cur: 0,
        size: size,
        pieces: make([]pieceShapeType, size),
    }

    for i := 0; i < len(pb.pieces); i++ {
        pb.pieces[i] = pieceShapeType(int(iPieceShapeType) + i)
    }

    pb.suffle()
    return &pb
}

func (pb *piecesBag) suffle() {
    rand.Shuffle(int(pb.size), func(i, j int) {
        pb.pieces[i], pb.pieces[j] = pb.pieces[j], pb.pieces[i]
    })
}

func (pb *piecesBag) next() pieceShapeType {
    p := pb.pieces[pb.cur]

    pb.cur += 1
    if pb.cur >= pb.size {
        pb.suffle()
        pb.cur = 0
    }

    return p
}

