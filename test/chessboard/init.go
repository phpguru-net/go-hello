package chessboard

// Declare a type named Rank which stores if a square is occupied by a piece - this will be a slice of bools
type Rank []bool

// Declare a type named Chessboard contains a map of eight Ranks, accessed with values from "A" to "H"
type Chessboard map[string]Rank

func NewChessBoard() Chessboard {
	chessboard := Chessboard{}
	for c := 'A'; c <= 'H'; c++ {
		chessboard[string(c)] = Rank{}
		for file := 0; file < 8; file++ {
			chessboard[string(c)][file] = false
		}
	}
	chessboard["A"] = []bool{true, false, true, false, false, false, false, false}
	chessboard["B"] = []bool{true, false, true, false, false, false, false, false}
	chessboard["C"] = []bool{true, false, true, false, false, false, false, false}
	chessboard["E"] = []bool{true, false, true, false, false, false, false, false}
	chessboard["G"] = []bool{true, false, true, false, false, false, false, false}
	chessboard["H"] = []bool{true, true, true, true, true, true, false, true}
	return chessboard
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) int {
	ranks, ok := cb[rank]
	if !ok {
		return 0
	}
	count := 0
	for _, occupied := range ranks {
		if occupied {
			count++
		}
	}
	return count
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) int {

	count := 0
	for _, ranks := range cb {
		for index, occupied := range ranks {
			if occupied && index+1 == file {
				count++
			}
		}
	}
	return count
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {
	return len(cb) * len(cb["A"])
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) int {
	count := 0
	for _, ranks := range cb {
		for _, occupied := range ranks {
			if occupied {
				count++
			}
		}
	}
	return count
}
