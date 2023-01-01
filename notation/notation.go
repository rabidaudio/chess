package notation

type Annotation string

const (
	Mistake     Annotation = "?"
	Blunder     Annotation = "??"
	Dubious     Annotation = "?!"
	Interesting Annotation = "!?"
	Good        Annotation = "!"
	Brilliant   Annotation = "!!"
)

// type piece struct {
// 	Piece rune `"R" | "N" | "B" | "Q" | "K"`
// }

// type rank struct {
// 	Rank rune `"1" | "2" | "3" | "4" | "5" | "6" | "7" | "8"`
// }

// type file struct {
// 	File rune `"a" | "b" | "c" | "d" | "e" | "f" | "g"`
// }

// type position struct {
// 	Rank *rank `@@`
// 	File *file `@@`
// }

// type Disambiguation struct {
// 	Rank     *rank     `@@`
// 	File     *file     `| @@`
// 	Position *position `| @@`
// }

// type Move struct {
// 	KingSideCastle  bool            `(  "O-O"`
// 	QueenSideCastle bool            ` | "O-O-O"`
// 	Piece           *piece          ` | (@@?`
// 	Disambiguation  *Disambiguation `   @@?`
// 	Capture         bool            `   "x"?`
// 	Destination     *position       `   @@?`
// 	Promotion       *piece          `   @@?))`
// 	Check           bool            `(  "+"`
// 	Mate            bool            ` | "#")?`
// 	Annotation      Annotation      `("?" | "??" | "?!" | "!?" | "!" | "!!")?`
// }

// type Round struct {
// 	Number    int  `@Int": "`
// 	WhiteMove Move `@@" "`
// 	BlackMove Move `@@`
// }

// type Game struct {
// 	Rounds []Round `(@@("\n" | " "))*`
// }

// piece: R | N | B | Q | K | <unicode>
// rank: 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8
// file: a | b | c | d | e | f | g | h
// queenside_castle: O-O-O
// kingside_castle: O-O
// castle: {kingside_castle}|{queenside_castle}
// position: {rank}{file}
// destination: {position}
// disambiguation: {rank}|{file}|{position}
// captures: x
// check: +
// mate: #
// promotion: {piece}
// move: ({castle}|({piece}{disambiguation}?)?{captures}?{destination}{promotion}?({check}|{mate})?){annotation}?
// move_number: <int>
// round: {move_number}: {move} {move}
// game: {round}*\n

// func parse() {
// 	move := Move{}
// 	_, err := participle.Build(&move)
// 	if err != nil {
// 		panic(err)
// 	}
// }
