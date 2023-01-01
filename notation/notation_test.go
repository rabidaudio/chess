package notation_test

import (
	"bytes"
	"fmt"
	"io"
	"regexp"
	"testing"

	"gotest.tools/v3/assert"
)

type Token int

const (
	WS Token = iota
	EOF

	MoveNumber
	KingSideCastle
	QueenSideCastle
	Piece
	Rank
	File
	Capture
	Promotion
	Check
	DoubleCheck
	CheckMate
	Annotation
	WhiteWin
	BlackWin
	Draw

	Tag
	SingleLineComment
	MultilineComment
)

// TODO: this strategy is very inefficent as it scans the whole body every time.
//  building a lexer manually to deal with the hard cases (WS, tags, comments)
//  would be much more efficent
func (t Token) expression() *regexp.Regexp {
	switch t {
	case WS:
		return regexp.MustCompile(`\s+`)
	case EOF:
		return regexp.MustCompile(`\z`)

	case MoveNumber:
		return regexp.MustCompile(`[0-9]+\.`)
	case KingSideCastle:
		return regexp.MustCompile(`O-O`)
	case QueenSideCastle:
		return regexp.MustCompile(`O-O-O`)
	case Piece:
		return regexp.MustCompile(`[QKBNR]`)
	case Rank:
		return regexp.MustCompile(`[1-8]`)
	case File:
		return regexp.MustCompile(`[a-h]`)
	case Capture:
		return regexp.MustCompile(`x`)
	case Promotion:
		return regexp.MustCompile(`=`)
	case Check:
		return regexp.MustCompile(`\+`)
	case DoubleCheck:
		return regexp.MustCompile(`\+\+`)
	case CheckMate:
		return regexp.MustCompile(`#`)
	case Annotation:
		return regexp.MustCompile(`[!?]+`)
	case WhiteWin:
		return regexp.MustCompile(`1-0`)
	case BlackWin:
		return regexp.MustCompile(`0-1`)
	case Draw:
		return regexp.MustCompile(`½–½`)

	case Tag:
		return regexp.MustCompile(`(?m)^(\[)([^ ]+)\s+\"(.*)\"(\])$`)
	case SingleLineComment:
		return regexp.MustCompile(`(?m)^;.*$`)
	case MultilineComment:
		return regexp.MustCompile(`(?s){.*}`)
	}
	return nil
}

type Match struct {
	Token   Token
	Literal string
	Pos     int64
}

type Scanner struct {
	r   *bytes.Reader
	pos int64
}

// note: allowing io.Reader and wrapping in bufio.Reader would be better for large files.
// However to use a regex-based lexer we need to be able to seek
func NewScanner(r *bytes.Reader) *Scanner {
	return &Scanner{r: r, pos: 0}
}

func (s *Scanner) test(t Token) (m *Match, ok bool) {
	sp := s.pos
	loc := t.expression().FindReaderIndex(s.r)
	s.r.Seek(sp, io.SeekStart) // seek back
	if loc != nil && loc[0] == 0 {
		lit := make([]byte, loc[1])
		s.r.Read(lit)
		s.pos += int64(loc[1])
		return &Match{Token: t, Literal: string(lit), Pos: sp}, true
	}
	return nil, false
}

// these are in a particular check order
var tokens = []Token{
	WS,
	SingleLineComment,
	MultilineComment,

	Tag,

	MoveNumber,

	QueenSideCastle,
	KingSideCastle,

	Piece,
	Rank,
	File,
	Capture,
	Promotion,
	DoubleCheck,
	Check,
	CheckMate,
	Annotation,

	WhiteWin,
	BlackWin,
	Draw,

	EOF,
}

func (s *Scanner) Scan() (*Match, error) {
	for _, token := range tokens {
		if m, ok := s.test(token); ok {
			return m, nil
		}
	}
	return nil, fmt.Errorf("invalid token at position %v", s.pos)
}

func readAll(s *Scanner) ([]*Match, error) {
	mm := make([]*Match, 0)
	for {
		m, err := s.Scan()
		if err != nil {
			return mm, err
		}
		mm = append(mm, m)
		if m.Token == EOF {
			return mm, nil
		}
	}
}

func TestScanner(t *testing.T) {
	cases := []struct {
		name    string
		input   string
		matches []*Match
	}{
		{
			name:  "moves",
			input: "1. e5 d6",
			matches: []*Match{
				{Token: MoveNumber, Literal: "1.", Pos: 0},
				{Token: WS, Literal: " ", Pos: 2},
				{Token: File, Literal: "e", Pos: 3},
				{Token: Rank, Literal: "5", Pos: 4},
				{Token: WS, Literal: " ", Pos: 5},
				{Token: File, Literal: "d", Pos: 6},
				{Token: Rank, Literal: "6", Pos: 7},
				{Token: EOF, Literal: "", Pos: 8},
			},
		},
		{
			name: "whitespace",
			input: `
		  	
   e5			
  		`,
			matches: []*Match{
				{Token: WS, Literal: "\n\t\t  \t\n   ", Pos: 0},
				{Token: File, Literal: "e", Pos: 10},
				{Token: Rank, Literal: "5", Pos: 11},
				{Token: WS, Literal: "\t\t\t\n  \t\t", Pos: 12},
				{Token: EOF, Literal: "", Pos: 20},
			},
		},
		{
			name: "tags",
			input: `
[Event "F/S Return Match"]
[Site "Belgrade, Serbia JUG"]
[Date "1992.11.04"]
[Round "29"]
[White "Fischer, Robert J."]
[Black "Spassky, Boris V."]
[Result "1/2-1/2"]`,
			matches: []*Match{
				{Token: WS, Literal: "\n", Pos: 0},
				{Token: Tag, Literal: "[Event \"F/S Return Match\"]", Pos: 1},
				{Token: WS, Literal: "\n", Pos: 27},
				{Token: Tag, Literal: "[Site \"Belgrade, Serbia JUG\"]", Pos: 28},
				{Token: WS, Literal: "\n", Pos: 57},
				{Token: Tag, Literal: "[Date \"1992.11.04\"]", Pos: 58},
				{Token: WS, Literal: "\n", Pos: 77},
				{Token: Tag, Literal: "[Round \"29\"]", Pos: 78},
				{Token: WS, Literal: "\n", Pos: 90},
				{Token: Tag, Literal: "[White \"Fischer, Robert J.\"]", Pos: 91},
				{Token: WS, Literal: "\n", Pos: 119},
				{Token: Tag, Literal: "[Black \"Spassky, Boris V.\"]", Pos: 120},
				{Token: WS, Literal: "\n", Pos: 147},
				{Token: Tag, Literal: "[Result \"1/2-1/2\"]", Pos: 148},
				{Token: EOF, Literal: "", Pos: 166},
			},
		},
		// comments
		// different moves
	}
	t.Parallel()
	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			s := NewScanner(bytes.NewReader([]byte(test.input)))
			mm, err := readAll(s)
			assert.NilError(t, err)
			assert.DeepEqual(t, mm, test.matches)
		})
	}
}

// func (s *Scanner) read() (rune, error) {
// 	r, _, err := s.r.ReadRune()
// 	return r, err
// }

// func (s *Scanner) unread() error {
// 	return s.r.UnreadRune()
// }

// func (s *Scanner) peekString(len int) string {
// 	bb := make([]byte, len)
// 	_, err := s.r.Read(bb)
// 	if err != nil {
// 		s.err = err
// 		return ""
// 	}
// 	return string(bb)
// }

// func (s *Scanner) readUntil(check func(rune) bool, inclusive bool) (string, error) {
// 	bb := make([]rune, 0)
// 	for {
// 		r, err := s.read()
// 		if err != nil {
// 			return "", err
// 		}
// 		if check(r) {
// 			if inclusive {
// 				bb = append(bb, r)
// 			} else {
// 				s.unread()
// 			}
// 			return string(bb), nil
// 		}
// 		bb = append(bb, r)
// 	}
// }

// func (s *Scanner) readUntilTarget(target rune, inclusive bool) (string, error) {
// 	return s.readUntil(func(r rune) bool { return r == target }, inclusive)
// }

// var whitespace = regexp.MustCompile("\\s+")

// func (s *Scanner) Scan() (tok token, lit string, err error) {
// 	r, err := s.read()
// 	if err != nil {
// 		return EOF, "", fmt.Errorf("lexer error: %w", err)
// 	}

// 	if whitespace.MatchString(string(r)) {
// 		err := s.unread()
// 		if err != nil {
// 			return EOF, "", fmt.Errorf("lexer error: %w", err)
// 		}
// 		ws, err := s.readUntil(func(r rune) bool {
// 			return !whitespace.MatchString(string(r))
// 		}, false)
// 		return WS, ws, err
// 	} else if r == '[' {
// 		err := s.unread()
// 		if err != nil {
// 			return EOF, "", fmt.Errorf("lexer error: %w", err)
// 		}
// 		tag, err := s.readUntil(func(r rune) bool {
// 			return r == ']'
// 		}, true)
// 		return Tag, tag, err
// 	}
// }

// func isWhitespace(r rune) bool {
// 	return whitespace.MatchReader()
// }

// func (s *Scanner) readWhitespace() (string, error) {
// 	bb := make([]rune, 0)
// 	for {
// 		r, err := s.read()

// 	}
// }

// func (s *Scanner) read() (rune, error) {
// 	ch, _, err := s.r.ReadRune()
// 	if err != nil {
// 		return -1
// 	}
// 	return ch
// }

// type Piece rune
// type Rank rune
// type File rune

// func (p *Piece) Capture(values []string) error {
// 	*p = Piece([]rune(values[0])[0])
// 	return nil
// }

// func (r *Rank) Capture(values []string) error {
// 	*r = Rank([]rune(values[0])[0])
// 	return nil
// }

// func (f *File) Capture(values []string) error {
// 	*f = File([]rune(values[0])[0])
// 	return nil
// }

// type Position struct {
// 	File File `@("a" | "b" | "c" | "d" | "e" | "f" | "g" | "h")`
// 	Rank Rank `@("1" | "2" | "3" | "4" | "5" | "6" | "7" | "8")`
// }

// // var move = regexp.MustCompile("([RBVQK]?x?")

// type Move struct {
// 	// KingSideCastle  bool   `(  "O-O"`
// 	// QueenSideCastle bool   ` | "O-O-O"`
// 	Piece *Piece ` (@("R" | "N" | "B" | "Q" | "K")?`
// 	// Disambiguation  *Disambiguation `   @@?`
// 	Capture     bool      `   "x"?`
// 	Destination *Position `   @@?`
// 	Promotion   *Piece    `   @("R" | "N" | "B" | "Q" | "K")?)` // )
// 	// Check       bool      `(  "+"`
// 	// Mate        bool      ` | "#")?`
// 	// Annotation  string    `@("?" | "??" | "?!" | "!?" | "!" | "!!")?`
// }

// func TestParse(t *testing.T) {
// 	// cases := []struct {
// 	// 	Name     string
// 	// 	Value    string
// 	// 	Expected Move
// 	// }{
// 	// 	{
// 	// 		Name:  "pawn",
// 	// 		Value: "e4",
// 	// 		Expected: Move{
// 	// 			Piece:       nil,
// 	// 			Destination: &Position{Rank: 'e', File: '4'},
// 	// 		},
// 	// 	},
// 	// }
// 	parser, err := participle.Build(&Position{})
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	d := Position{}
// 	err = parser.ParseString("", "e4\n", &d)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	assert.Equal(t, d.File, 'e')
// 	assert.Equal(t, d.Rank, '4')

// 	// for _, test := range cases {
// 	// 	m := Move{}
// 	// 	err = parser.ParseString(test.Name, test.Value, &m)
// 	// 	if err != nil {
// 	// 		t.Fatal(err)
// 	// 	}
// 	// 	assert.DeepEqual(t, m, test.Expected)
// 	// }
// }
