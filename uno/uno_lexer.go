// Code generated from uno.g4 by ANTLR 4.12.0. DO NOT EDIT.

package uno

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type unoLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var unolexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func unolexerLexerInit() {
	staticData := &unolexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "'('", "')'", "'['", "']'", "'.'", "','", "'\"'", "'+'", "'-'",
		"'*'", "'/'", "'%'", "'int64'", "'int64s'", "'float32'", "'float32s'",
		"'string'", "'strings'", "'ON'", "'and'", "'or'", "'not'", "'in'", "'true'",
		"'false'", "", "'='", "'=='", "'<>'", "'!='", "'>'", "'>='", "'<'",
		"'<='",
	}
	staticData.symbolicNames = []string{
		"", "BRACKET_OPEN", "BRACKET_CLOSE", "SQUARE_OPEN", "SQUARE_CLOSE",
		"DOT", "COMMA", "QUOTA", "T_ADD", "T_SUB", "T_MUL", "T_DIV", "T_MOD",
		"T_INT", "T_INTS", "T_FLOAT", "T_FLOATS", "T_STRING", "T_STRINGS", "T_ON",
		"T_AND", "T_OR", "T_NOT", "T_IN", "T_TRUE", "T_FALSE", "T_COMPARE",
		"T_EQUAL", "T_EQUAL2", "T_NOTEQUAL", "T_NOTEQUAL2", "T_GREATER", "T_GREATEREQUAL",
		"T_LESS", "T_LESSEQUAL", "IDENTIFIER", "INTEGER_LIST", "INTEGER", "DECIMAL_LIST",
		"DECIMAL", "STRING_LIST", "STRING", "WS",
	}
	staticData.ruleNames = []string{
		"BRACKET_OPEN", "BRACKET_CLOSE", "SQUARE_OPEN", "SQUARE_CLOSE", "DOT",
		"COMMA", "QUOTA", "T_ADD", "T_SUB", "T_MUL", "T_DIV", "T_MOD", "T_INT",
		"T_INTS", "T_FLOAT", "T_FLOATS", "T_STRING", "T_STRINGS", "T_ON", "T_AND",
		"T_OR", "T_NOT", "T_IN", "T_TRUE", "T_FALSE", "T_COMPARE", "T_EQUAL",
		"T_EQUAL2", "T_NOTEQUAL", "T_NOTEQUAL2", "T_GREATER", "T_GREATEREQUAL",
		"T_LESS", "T_LESSEQUAL", "IDENTIFIER", "INTEGER_LIST", "INTEGER", "DECIMAL_LIST",
		"DECIMAL", "STRING_LIST", "STRING", "ESC", "UNICODE", "HEX", "SAFECODEPOINT",
		"WS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 42, 329, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 1, 0, 1, 0,
		1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6,
		1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12,
		1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15,
		1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1,
		16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17,
		1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1,
		20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23,
		1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 1,
		25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 3, 25, 199, 8, 25, 1, 26,
		1, 26, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1,
		30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 34,
		1, 34, 5, 34, 224, 8, 34, 10, 34, 12, 34, 227, 9, 34, 1, 35, 1, 35, 1,
		35, 1, 35, 5, 35, 233, 8, 35, 10, 35, 12, 35, 236, 9, 35, 1, 35, 1, 35,
		1, 35, 1, 36, 3, 36, 242, 8, 36, 1, 36, 1, 36, 1, 36, 5, 36, 247, 8, 36,
		10, 36, 12, 36, 250, 9, 36, 3, 36, 252, 8, 36, 1, 37, 1, 37, 1, 37, 1,
		37, 5, 37, 258, 8, 37, 10, 37, 12, 37, 261, 9, 37, 1, 37, 1, 37, 1, 37,
		1, 38, 3, 38, 267, 8, 38, 1, 38, 1, 38, 1, 38, 5, 38, 272, 8, 38, 10, 38,
		12, 38, 275, 9, 38, 3, 38, 277, 8, 38, 1, 38, 1, 38, 5, 38, 281, 8, 38,
		10, 38, 12, 38, 284, 9, 38, 1, 39, 1, 39, 1, 39, 1, 39, 5, 39, 290, 8,
		39, 10, 39, 12, 39, 293, 9, 39, 1, 39, 1, 39, 1, 39, 1, 40, 1, 40, 1, 40,
		5, 40, 301, 8, 40, 10, 40, 12, 40, 304, 9, 40, 1, 40, 1, 40, 1, 41, 1,
		41, 1, 41, 3, 41, 311, 8, 41, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42,
		1, 43, 1, 43, 1, 44, 1, 44, 1, 45, 4, 45, 324, 8, 45, 11, 45, 12, 45, 325,
		1, 45, 1, 45, 0, 0, 46, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15,
		8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17,
		35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26,
		53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 35,
		71, 36, 73, 37, 75, 38, 77, 39, 79, 40, 81, 41, 83, 0, 85, 0, 87, 0, 89,
		0, 91, 42, 1, 0, 21, 2, 0, 73, 73, 105, 105, 2, 0, 78, 78, 110, 110, 2,
		0, 84, 84, 116, 116, 2, 0, 83, 83, 115, 115, 2, 0, 70, 70, 102, 102, 2,
		0, 76, 76, 108, 108, 2, 0, 79, 79, 111, 111, 2, 0, 65, 65, 97, 97, 2, 0,
		82, 82, 114, 114, 2, 0, 71, 71, 103, 103, 2, 0, 68, 68, 100, 100, 2, 0,
		85, 85, 117, 117, 2, 0, 69, 69, 101, 101, 3, 0, 65, 90, 95, 95, 97, 122,
		4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 1, 0, 49, 57, 1, 0, 48, 57, 13,
		0, 34, 34, 47, 47, 66, 66, 70, 70, 78, 78, 82, 82, 84, 84, 92, 92, 98,
		98, 102, 102, 110, 110, 114, 114, 116, 116, 3, 0, 48, 57, 65, 70, 97, 102,
		3, 0, 0, 31, 34, 34, 92, 92, 3, 0, 9, 10, 13, 13, 32, 32, 346, 0, 1, 1,
		0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1,
		0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17,
		1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0,
		25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0,
		0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0,
		0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0,
		0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1,
		0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63,
		1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0,
		71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0,
		0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 91, 1, 0, 0, 0, 1, 93, 1, 0, 0,
		0, 3, 95, 1, 0, 0, 0, 5, 97, 1, 0, 0, 0, 7, 99, 1, 0, 0, 0, 9, 101, 1,
		0, 0, 0, 11, 103, 1, 0, 0, 0, 13, 105, 1, 0, 0, 0, 15, 107, 1, 0, 0, 0,
		17, 109, 1, 0, 0, 0, 19, 111, 1, 0, 0, 0, 21, 113, 1, 0, 0, 0, 23, 115,
		1, 0, 0, 0, 25, 117, 1, 0, 0, 0, 27, 123, 1, 0, 0, 0, 29, 130, 1, 0, 0,
		0, 31, 138, 1, 0, 0, 0, 33, 147, 1, 0, 0, 0, 35, 154, 1, 0, 0, 0, 37, 162,
		1, 0, 0, 0, 39, 165, 1, 0, 0, 0, 41, 169, 1, 0, 0, 0, 43, 172, 1, 0, 0,
		0, 45, 176, 1, 0, 0, 0, 47, 179, 1, 0, 0, 0, 49, 184, 1, 0, 0, 0, 51, 198,
		1, 0, 0, 0, 53, 200, 1, 0, 0, 0, 55, 202, 1, 0, 0, 0, 57, 205, 1, 0, 0,
		0, 59, 208, 1, 0, 0, 0, 61, 211, 1, 0, 0, 0, 63, 213, 1, 0, 0, 0, 65, 216,
		1, 0, 0, 0, 67, 218, 1, 0, 0, 0, 69, 221, 1, 0, 0, 0, 71, 228, 1, 0, 0,
		0, 73, 251, 1, 0, 0, 0, 75, 253, 1, 0, 0, 0, 77, 266, 1, 0, 0, 0, 79, 285,
		1, 0, 0, 0, 81, 297, 1, 0, 0, 0, 83, 307, 1, 0, 0, 0, 85, 312, 1, 0, 0,
		0, 87, 318, 1, 0, 0, 0, 89, 320, 1, 0, 0, 0, 91, 323, 1, 0, 0, 0, 93, 94,
		5, 40, 0, 0, 94, 2, 1, 0, 0, 0, 95, 96, 5, 41, 0, 0, 96, 4, 1, 0, 0, 0,
		97, 98, 5, 91, 0, 0, 98, 6, 1, 0, 0, 0, 99, 100, 5, 93, 0, 0, 100, 8, 1,
		0, 0, 0, 101, 102, 5, 46, 0, 0, 102, 10, 1, 0, 0, 0, 103, 104, 5, 44, 0,
		0, 104, 12, 1, 0, 0, 0, 105, 106, 5, 34, 0, 0, 106, 14, 1, 0, 0, 0, 107,
		108, 5, 43, 0, 0, 108, 16, 1, 0, 0, 0, 109, 110, 5, 45, 0, 0, 110, 18,
		1, 0, 0, 0, 111, 112, 5, 42, 0, 0, 112, 20, 1, 0, 0, 0, 113, 114, 5, 47,
		0, 0, 114, 22, 1, 0, 0, 0, 115, 116, 5, 37, 0, 0, 116, 24, 1, 0, 0, 0,
		117, 118, 7, 0, 0, 0, 118, 119, 7, 1, 0, 0, 119, 120, 7, 2, 0, 0, 120,
		121, 5, 54, 0, 0, 121, 122, 5, 52, 0, 0, 122, 26, 1, 0, 0, 0, 123, 124,
		7, 0, 0, 0, 124, 125, 7, 1, 0, 0, 125, 126, 7, 2, 0, 0, 126, 127, 5, 54,
		0, 0, 127, 128, 5, 52, 0, 0, 128, 129, 7, 3, 0, 0, 129, 28, 1, 0, 0, 0,
		130, 131, 7, 4, 0, 0, 131, 132, 7, 5, 0, 0, 132, 133, 7, 6, 0, 0, 133,
		134, 7, 7, 0, 0, 134, 135, 7, 2, 0, 0, 135, 136, 5, 51, 0, 0, 136, 137,
		5, 50, 0, 0, 137, 30, 1, 0, 0, 0, 138, 139, 7, 4, 0, 0, 139, 140, 7, 5,
		0, 0, 140, 141, 7, 6, 0, 0, 141, 142, 7, 7, 0, 0, 142, 143, 7, 2, 0, 0,
		143, 144, 5, 51, 0, 0, 144, 145, 5, 50, 0, 0, 145, 146, 7, 3, 0, 0, 146,
		32, 1, 0, 0, 0, 147, 148, 7, 3, 0, 0, 148, 149, 7, 2, 0, 0, 149, 150, 7,
		8, 0, 0, 150, 151, 7, 0, 0, 0, 151, 152, 7, 1, 0, 0, 152, 153, 7, 9, 0,
		0, 153, 34, 1, 0, 0, 0, 154, 155, 7, 3, 0, 0, 155, 156, 7, 2, 0, 0, 156,
		157, 7, 8, 0, 0, 157, 158, 7, 0, 0, 0, 158, 159, 7, 1, 0, 0, 159, 160,
		7, 9, 0, 0, 160, 161, 7, 3, 0, 0, 161, 36, 1, 0, 0, 0, 162, 163, 7, 6,
		0, 0, 163, 164, 7, 1, 0, 0, 164, 38, 1, 0, 0, 0, 165, 166, 7, 7, 0, 0,
		166, 167, 7, 1, 0, 0, 167, 168, 7, 10, 0, 0, 168, 40, 1, 0, 0, 0, 169,
		170, 7, 6, 0, 0, 170, 171, 7, 8, 0, 0, 171, 42, 1, 0, 0, 0, 172, 173, 7,
		1, 0, 0, 173, 174, 7, 6, 0, 0, 174, 175, 7, 2, 0, 0, 175, 44, 1, 0, 0,
		0, 176, 177, 7, 0, 0, 0, 177, 178, 7, 1, 0, 0, 178, 46, 1, 0, 0, 0, 179,
		180, 7, 2, 0, 0, 180, 181, 7, 8, 0, 0, 181, 182, 7, 11, 0, 0, 182, 183,
		7, 12, 0, 0, 183, 48, 1, 0, 0, 0, 184, 185, 7, 4, 0, 0, 185, 186, 7, 7,
		0, 0, 186, 187, 7, 5, 0, 0, 187, 188, 7, 3, 0, 0, 188, 189, 7, 12, 0, 0,
		189, 50, 1, 0, 0, 0, 190, 199, 3, 53, 26, 0, 191, 199, 3, 55, 27, 0, 192,
		199, 3, 57, 28, 0, 193, 199, 3, 59, 29, 0, 194, 199, 3, 61, 30, 0, 195,
		199, 3, 63, 31, 0, 196, 199, 3, 65, 32, 0, 197, 199, 3, 67, 33, 0, 198,
		190, 1, 0, 0, 0, 198, 191, 1, 0, 0, 0, 198, 192, 1, 0, 0, 0, 198, 193,
		1, 0, 0, 0, 198, 194, 1, 0, 0, 0, 198, 195, 1, 0, 0, 0, 198, 196, 1, 0,
		0, 0, 198, 197, 1, 0, 0, 0, 199, 52, 1, 0, 0, 0, 200, 201, 5, 61, 0, 0,
		201, 54, 1, 0, 0, 0, 202, 203, 5, 61, 0, 0, 203, 204, 5, 61, 0, 0, 204,
		56, 1, 0, 0, 0, 205, 206, 5, 60, 0, 0, 206, 207, 5, 62, 0, 0, 207, 58,
		1, 0, 0, 0, 208, 209, 5, 33, 0, 0, 209, 210, 5, 61, 0, 0, 210, 60, 1, 0,
		0, 0, 211, 212, 5, 62, 0, 0, 212, 62, 1, 0, 0, 0, 213, 214, 5, 62, 0, 0,
		214, 215, 5, 61, 0, 0, 215, 64, 1, 0, 0, 0, 216, 217, 5, 60, 0, 0, 217,
		66, 1, 0, 0, 0, 218, 219, 5, 60, 0, 0, 219, 220, 5, 61, 0, 0, 220, 68,
		1, 0, 0, 0, 221, 225, 7, 13, 0, 0, 222, 224, 7, 14, 0, 0, 223, 222, 1,
		0, 0, 0, 224, 227, 1, 0, 0, 0, 225, 223, 1, 0, 0, 0, 225, 226, 1, 0, 0,
		0, 226, 70, 1, 0, 0, 0, 227, 225, 1, 0, 0, 0, 228, 234, 3, 1, 0, 0, 229,
		230, 3, 73, 36, 0, 230, 231, 3, 11, 5, 0, 231, 233, 1, 0, 0, 0, 232, 229,
		1, 0, 0, 0, 233, 236, 1, 0, 0, 0, 234, 232, 1, 0, 0, 0, 234, 235, 1, 0,
		0, 0, 235, 237, 1, 0, 0, 0, 236, 234, 1, 0, 0, 0, 237, 238, 3, 73, 36,
		0, 238, 239, 3, 3, 1, 0, 239, 72, 1, 0, 0, 0, 240, 242, 5, 45, 0, 0, 241,
		240, 1, 0, 0, 0, 241, 242, 1, 0, 0, 0, 242, 243, 1, 0, 0, 0, 243, 252,
		5, 48, 0, 0, 244, 248, 7, 15, 0, 0, 245, 247, 7, 16, 0, 0, 246, 245, 1,
		0, 0, 0, 247, 250, 1, 0, 0, 0, 248, 246, 1, 0, 0, 0, 248, 249, 1, 0, 0,
		0, 249, 252, 1, 0, 0, 0, 250, 248, 1, 0, 0, 0, 251, 241, 1, 0, 0, 0, 251,
		244, 1, 0, 0, 0, 252, 74, 1, 0, 0, 0, 253, 259, 3, 1, 0, 0, 254, 255, 3,
		77, 38, 0, 255, 256, 3, 11, 5, 0, 256, 258, 1, 0, 0, 0, 257, 254, 1, 0,
		0, 0, 258, 261, 1, 0, 0, 0, 259, 257, 1, 0, 0, 0, 259, 260, 1, 0, 0, 0,
		260, 262, 1, 0, 0, 0, 261, 259, 1, 0, 0, 0, 262, 263, 3, 77, 38, 0, 263,
		264, 3, 3, 1, 0, 264, 76, 1, 0, 0, 0, 265, 267, 5, 45, 0, 0, 266, 265,
		1, 0, 0, 0, 266, 267, 1, 0, 0, 0, 267, 276, 1, 0, 0, 0, 268, 277, 5, 48,
		0, 0, 269, 273, 7, 15, 0, 0, 270, 272, 7, 16, 0, 0, 271, 270, 1, 0, 0,
		0, 272, 275, 1, 0, 0, 0, 273, 271, 1, 0, 0, 0, 273, 274, 1, 0, 0, 0, 274,
		277, 1, 0, 0, 0, 275, 273, 1, 0, 0, 0, 276, 268, 1, 0, 0, 0, 276, 269,
		1, 0, 0, 0, 277, 278, 1, 0, 0, 0, 278, 282, 5, 46, 0, 0, 279, 281, 7, 16,
		0, 0, 280, 279, 1, 0, 0, 0, 281, 284, 1, 0, 0, 0, 282, 280, 1, 0, 0, 0,
		282, 283, 1, 0, 0, 0, 283, 78, 1, 0, 0, 0, 284, 282, 1, 0, 0, 0, 285, 291,
		3, 1, 0, 0, 286, 287, 3, 81, 40, 0, 287, 288, 3, 11, 5, 0, 288, 290, 1,
		0, 0, 0, 289, 286, 1, 0, 0, 0, 290, 293, 1, 0, 0, 0, 291, 289, 1, 0, 0,
		0, 291, 292, 1, 0, 0, 0, 292, 294, 1, 0, 0, 0, 293, 291, 1, 0, 0, 0, 294,
		295, 3, 81, 40, 0, 295, 296, 3, 3, 1, 0, 296, 80, 1, 0, 0, 0, 297, 302,
		3, 13, 6, 0, 298, 301, 3, 83, 41, 0, 299, 301, 3, 89, 44, 0, 300, 298,
		1, 0, 0, 0, 300, 299, 1, 0, 0, 0, 301, 304, 1, 0, 0, 0, 302, 300, 1, 0,
		0, 0, 302, 303, 1, 0, 0, 0, 303, 305, 1, 0, 0, 0, 304, 302, 1, 0, 0, 0,
		305, 306, 3, 13, 6, 0, 306, 82, 1, 0, 0, 0, 307, 310, 5, 92, 0, 0, 308,
		311, 7, 17, 0, 0, 309, 311, 3, 85, 42, 0, 310, 308, 1, 0, 0, 0, 310, 309,
		1, 0, 0, 0, 311, 84, 1, 0, 0, 0, 312, 313, 7, 11, 0, 0, 313, 314, 3, 87,
		43, 0, 314, 315, 3, 87, 43, 0, 315, 316, 3, 87, 43, 0, 316, 317, 3, 87,
		43, 0, 317, 86, 1, 0, 0, 0, 318, 319, 7, 18, 0, 0, 319, 88, 1, 0, 0, 0,
		320, 321, 8, 19, 0, 0, 321, 90, 1, 0, 0, 0, 322, 324, 7, 20, 0, 0, 323,
		322, 1, 0, 0, 0, 324, 325, 1, 0, 0, 0, 325, 323, 1, 0, 0, 0, 325, 326,
		1, 0, 0, 0, 326, 327, 1, 0, 0, 0, 327, 328, 6, 45, 0, 0, 328, 92, 1, 0,
		0, 0, 17, 0, 198, 225, 234, 241, 248, 251, 259, 266, 273, 276, 282, 291,
		300, 302, 310, 325, 1, 6, 0, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// unoLexerInit initializes any static state used to implement unoLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewunoLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func UnoLexerInit() {
	staticData := &unolexerLexerStaticData
	staticData.once.Do(unolexerLexerInit)
}

// NewunoLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewunoLexer(input antlr.CharStream) *unoLexer {
	UnoLexerInit()
	l := new(unoLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &unolexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "uno.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// unoLexer tokens.
const (
	unoLexerBRACKET_OPEN   = 1
	unoLexerBRACKET_CLOSE  = 2
	unoLexerSQUARE_OPEN    = 3
	unoLexerSQUARE_CLOSE   = 4
	unoLexerDOT            = 5
	unoLexerCOMMA          = 6
	unoLexerQUOTA          = 7
	unoLexerT_ADD          = 8
	unoLexerT_SUB          = 9
	unoLexerT_MUL          = 10
	unoLexerT_DIV          = 11
	unoLexerT_MOD          = 12
	unoLexerT_INT          = 13
	unoLexerT_INTS         = 14
	unoLexerT_FLOAT        = 15
	unoLexerT_FLOATS       = 16
	unoLexerT_STRING       = 17
	unoLexerT_STRINGS      = 18
	unoLexerT_ON           = 19
	unoLexerT_AND          = 20
	unoLexerT_OR           = 21
	unoLexerT_NOT          = 22
	unoLexerT_IN           = 23
	unoLexerT_TRUE         = 24
	unoLexerT_FALSE        = 25
	unoLexerT_COMPARE      = 26
	unoLexerT_EQUAL        = 27
	unoLexerT_EQUAL2       = 28
	unoLexerT_NOTEQUAL     = 29
	unoLexerT_NOTEQUAL2    = 30
	unoLexerT_GREATER      = 31
	unoLexerT_GREATEREQUAL = 32
	unoLexerT_LESS         = 33
	unoLexerT_LESSEQUAL    = 34
	unoLexerIDENTIFIER     = 35
	unoLexerINTEGER_LIST   = 36
	unoLexerINTEGER        = 37
	unoLexerDECIMAL_LIST   = 38
	unoLexerDECIMAL        = 39
	unoLexerSTRING_LIST    = 40
	unoLexerSTRING         = 41
	unoLexerWS             = 42
)
