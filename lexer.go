package main

import (
	"fmt"
	"os"
	"log"
	"unicode"
)

type Lexer struct {
	input []rune // input being lexed
	line_number int // the current line number the lexer is on
	index int // the current character the lexer is on
	show_debug bool // true to print debug output, false to disable.
}

// loads in the contents from a file into input field of Lexer struct
func (lexer *Lexer) LoadInput(filename string) {
	// read in contents from file
	file_content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	// add to lexer input
	lexer.input = []rune(string(file_content))

	if lexer.show_debug {
		// print contents
		fmt.Printf("From file %s:\n%s\n", filename, string(lexer.input))
		for _, char := range lexer.input {
			fmt.Printf("%q, ", char)
		}
		fmt.Println()
	}
}

// returns true if s is a keyword and stores the TokenID in *TokenID.
// returns false otherwise.
func is_keyword(s string, token_type *TokenID) bool {
	switch s {
	case "def": 
		*token_type = KEYW_DEF; return true
	case "ret":
		*token_type = KEYW_RET; return true
    case "dbl":
        *token_type = KEYW_DBL; return true
	}
	return false
}

func (lexer Lexer) CurrentChar() rune {
	if lexer.index < len(lexer.input) {
		return lexer.input[lexer.index]
	} 
	return 0 // EOF!
}

func (lexer Lexer) PeekChar() rune {
	if lexer.index + 1 < len(lexer.input) {
		return lexer.input[lexer.index + 1]
	} 
	return 0 // EOF!
}

// returns the next token in the Lexer.
func (lexer *Lexer) NextToken() Token {
	if lexer.input == nil {
		log.Fatal("Lexer input is not defined... remember to call LoadInput()")
	}

	value_buffer := "" // buffer to store token value
	current_char := lexer.CurrentChar() // current token of lexer
	num_decimals := 0

	// skip whitespace
	for unicode.IsSpace(current_char) {
		lexer.index++; current_char = lexer.CurrentChar()
	}

	if lexer.index >= len(lexer.input) {
		// if we are trying to index after the array, this is the EOF token.
		return Token{id: EOF}
	}

    if current_char == '(' {
        lexer.index++; return Token{id: OPEN_PAREN}
    } else if current_char == ')' {
        lexer.index++; return Token{id: CLOSE_PAREN}
    } else if current_char == '{' {
        lexer.index++; return Token{id: OPEN_BRACE}
    } else if current_char == '}' {
        lexer.index++; return Token{id: CLOSE_BRACE}
    } else if current_char == '+' {
        lexer.index++; return Token{id: PLUS}
    } else if current_char == '-' {
        lexer.index++; return Token{id: MINUS}
    } else if current_char == '/' {
        lexer.index++; return Token{id: SLASH}
    } else if current_char == '*' {
        lexer.index++; return Token{id: ASTERIK}
    } else if current_char == ',' {
        lexer.index++; return Token{id: COMMA}
    } else if current_char == '=' {
        if (lexer.PeekChar() == '>') {
            lexer.index += 2; return Token{id: ARROW}
        } else {
            lexer.index++; return Token{id: EQUAL}
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
			lexer.index++; current_char = lexer.CurrentChar()
		}
		// we now have populated the value field, let's return a double.
		return Token{id: DOUBLE, value: value_buffer}

	// collect something alphanumeric (keyword? identifier? etc...)
	} else if unicode.IsLetter(current_char) {
		for unicode.IsLetter(current_char) || unicode.IsDigit(current_char) {
			value_buffer += string(current_char)
			lexer.index++; current_char = lexer.CurrentChar()
		}
		// we now have populated the value field, let's check to see if its a keyword, then return that.
		var token_id TokenID 
		if is_keyword(value_buffer, &token_id) {
			return Token{id: token_id, value: value_buffer}
		} else {
			// just return an identifier if it's not matched.
			return Token{id: IDENTIFIER, value: value_buffer}
		}
	}
	
	// if we don't know this token, return unknown
	return Token{id: UNKNOWN}
}

// tokenize the input and return a slice of Tokens.
func (lexer *Lexer) Tokenize() []Token {
	var tokens []Token
	next_token := lexer.NextToken()
	// collect all tokens and append to the result slice
	for next_token.id != EOF {
		fmt.Println(next_token.Repr())
		tokens = append(tokens, next_token)
		next_token = lexer.NextToken()
	}
	// append EOF token
	tokens = append(tokens, next_token)

	if lexer.show_debug {
		fmt.Print(tokens)
	}

	return tokens
}