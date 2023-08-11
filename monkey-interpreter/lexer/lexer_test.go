package lexer

import (
	"monkey/token"
	"fmt"
	"testing"
)

func TestNextToken(t *testing.T) {
    input := `
    let five = 5;
let ten = 10;

let add = fn(x, y) {
  x + y;
};

let result = add(five, ten);
!-/*5;
5 < 10 > 5;

if (5 < 10) {
	return true;
} else {
	return false;
}

10 == 10;
10 != 9;
10 <= 9;
9 >= 10;`
    tests := []struct{
        expectedType token.TokenType
        expectedLiteral string
    }{
        {token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.INT, "10"},
		{token.EQ, "=="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.INT, "10"},
		{token.NOT_EQ, "!="},
		{token.INT, "9"},
		{token.SEMICOLON, ";"},
		{token.INT, "10"},
		{token.LT_EQ, "<="},
		{token.INT, "9"},
		{token.SEMICOLON, ";"},
		{token.INT, "9"},
		{token.GT_EQ, ">="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
    }

    l := New(input)

    for i, tt := range(tests){
        tok := l.nextToken()
        fmt.Println(tok)
        if tok.Type != tt.expectedType {
            t.Fatalf("index: %d\n Expected Token:%s\n got Token:%s\n", i, tt.expectedType, tok.Type)
        }

        if tok.Literal != tt.expectedLiteral {
            t.Fatalf("index: %d\n Expected Literal:%s\n got Literal:%s\n", i, tt.expectedLiteral, tok.Literal)
        }
    }

}

//
func TestIsLetter(t *testing.T) {
    tests := []struct{
        expectedByte byte
        expectedValue bool
    }{
            {'A', true},
            {'B', true},
            {'-', false},
            {'!', false},
        }

        for i, tt  := range(tests) {
            gotBool := isLetter(tt.expectedByte)
            if gotBool != tt.expectedValue{
                t.Fatalf("index: %d  Expected: %t Got: %t for %v", i, tt.expectedValue, gotBool, tt.expectedByte)
            }
        }
}



func TestIsDigit(t *testing.T) {
    fmt.Println("first line")
    tests := []struct{
        expectedByte byte
        expectedValue bool
    }{
            {'0', true},
            {'1', true},
            {'2', true},
            {'3', true},
            {'4', true},
            {'5', true},
            {'6', true},
            {'7', true},
            {'8', true},
            {'9', true},
        }

        for i, tt  := range(tests) {
            gotBool := isDigit(tt.expectedByte)
            if gotBool != tt.expectedValue{
                t.Fatalf("index: %d  Expected: %t Got: %t", i, tt.expectedValue, gotBool)
            }
        }

}



