package lexer

import (
	"fmt"
	"os"
	"log"
	"unicode"
	"github.com/ethan-prime/graphite/tokens"
)

type Lexer struct {
	Input []rune // Input being lexed
	LineNumber int // the current line number the lexer is on
	Index int // the current character the lexer is on
	ShowDebug bool // true to print debug output, false to disable.
}

// loads in the contents from a file into Input field of Lexer struct
func (lexer *Lexer) LoadInput(filename string) {
	// read in contents from file
	file_content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	// add to lexer Input
	lexer.Input = []rune(string(file_content))

	if lexer.ShowDebug {
		// print contents
		fmt.Printf("From file %s:\n%s\n", filename, string(lexer.Input))
		for _, char := range lexer.Input {
			fmt.Printf("%q, ", char)
		}
		fmt.Println()
	}
}

// returns true if s is a keyword and stores the TokenID in *TokenID.
// returns false otherwise.
func is_keyword(s string, token_type *tokens.TokenID) bool {
	switch s {
	case "def": 
		*token_type = tokens.KEYW_DEF; return true
	case "ret":
		*token_type = tokens.KEYW_RET; return true
    case "dbl":
        *token_type = tokens.KEYW_DBL; return true
	}
	return false
}

func (lexer Lexer) CurrentChar() rune {
	if lexer.Index < len(lexer.Input) {
		return lexer.Input[lexer.Index]
	} 
	return 0 // EOF!
}

func (lexer Lexer) PeekChar() rune {
	if lexer.Index + 1 < len(lexer.Input) {
		return lexer.Input[lexer.Index + 1]
	} 
	return 0 // EOF!
}

// returns the next token in the Lexer.
func (lexer *Lexer) NextToken() tokens.Token {
	if lexer.Input == nil {
		log.Fatal("Lexer Input is not defined... remember to call LoadInput()")
	}

	value_buffer := "" // buffer to store token value
	current_char := lexer.CurrentChar() // current token of lexer
	num_decimals := 0

	// skip whitespace
	for unicode.IsSpace(current_char) {
        if current_char == '\n' {
            lexer.LineNumber++ // increment line number if we see a line break
        }
		lexer.Index++; current_char = lexer.CurrentChar()
	}

	if lexer.Index >= len(lexer.Input) {
		// if we are trying to Index after the array, this is the EOF token.
		return tokens.Token{ID: tokens.EOF, LineNumber: lexer.LineNumber}
	}

    if current_char == '(' {
        lexer.Index++; return tokens.Token{ID: tokens.OPEN_PAREN, LineNumber: lexer.LineNumber}
    } else if current_char == ')' {
        lexer.Index++; return tokens.Token{ID: tokens.CLOSE_PAREN, LineNumber: lexer.LineNumber}
    } else if current_char == '{' {
        lexer.Index++; return tokens.Token{ID: tokens.OPEN_BRACE, LineNumber: lexer.LineNumber}
    } else if current_char == '}' {
        lexer.Index++; return tokens.Token{ID: tokens.CLOSE_BRACE, LineNumber: lexer.LineNumber}
    } else if current_char == '+' {
        lexer.Index++; return tokens.Token{ID: tokens.PLUS, LineNumber: lexer.LineNumber}
    } else if current_char == '-' {
        lexer.Index++; return tokens.Token{ID: tokens.MINUS, LineNumber: lexer.LineNumber}
    } else if current_char == '/' {
        lexer.Index++; return tokens.Token{ID: tokens.SLASH, LineNumber: lexer.LineNumber}
    } else if current_char == '*' {
        lexer.Index++; return tokens.Token{ID: tokens.ASTERIK, LineNumber: lexer.LineNumber}
    } else if current_char == ',' {
        lexer.Index++; return tokens.Token{ID: tokens.COMMA, LineNumber: lexer.LineNumber}
    } else if current_char == '=' {
        if (lexer.PeekChar() == '>') {
            lexer.Index += 2; return tokens.Token{ID: tokens.ARROW, LineNumber: lexer.LineNumber}
        } else {
            lexer.Index++; return tokens.Token{ID: tokens.EQUAL, LineNumber: lexer.LineNumber}
        }
    } else if unicode.IsDigit(current_char) || current_char == '.' {
        // collect a double
		for (unicode.IsDigit(current_char) || current_char == '.') && (num_decimals <= 1) {
			if current_char == '.' {
				num_decimals++
			}
			if num_decimals > 1 {
				break
			}
			value_buffer += string(current_char)
			lexer.Index++; current_char = lexer.CurrentChar()
		}
		// we now have populated the value field, let's return a double.
		return tokens.Token{ID: tokens.DOUBLE, Value: value_buffer, LineNumber: lexer.LineNumber}

	// collect something alphanumeric (keyword? identifier? etc...)
	} else if unicode.IsLetter(current_char) {
		for unicode.IsLetter(current_char) || unicode.IsDigit(current_char) {
			value_buffer += string(current_char)
			lexer.Index++; current_char = lexer.CurrentChar()
		}
		// we now have populated the value field, let's check to see if its a keyword, then return that.
		var token_id tokens.TokenID 
		if is_keyword(value_buffer, &token_id) {
			return tokens.Token{ID: token_id, Value: value_buffer, LineNumber: lexer.LineNumber}
		} else {
			// just return an identifier if it's not matched.
			return tokens.Token{ID: tokens.IDENTIFIER, Value: value_buffer, LineNumber: lexer.LineNumber}
		}
	}
	
	// if we don't know this token, return unknown
    lexer.Index++; return tokens.Token{ID: tokens.UNKNOWN, LineNumber: lexer.LineNumber}
}

// tokenize the Input and return a slice of Tokens.
func (lexer *Lexer) Tokenize() []tokens.Token {
	var lexer_tokens []tokens.Token
	next_token := lexer.NextToken()
	// collect all tokens and append to the result slice
	for next_token.ID != tokens.EOF {
		fmt.Printf("%s, (line %d)\n", next_token.Repr(), next_token.LineNumber)
		lexer_tokens = append(lexer_tokens, next_token)
		next_token = lexer.NextToken()
	}
	// append EOF token
	lexer_tokens = append(lexer_tokens, next_token)

	if lexer.ShowDebug {
		fmt.Print(lexer_tokens)
	}

	return lexer_tokens
}