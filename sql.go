//line sql.y:6
package sqlparser

import __yyfmt__ "fmt"

//line sql.y:6
import "bytes"

func SetParseTree(yylex interface{}, stmt Statement) {
	yylex.(*Tokenizer).ParseTree = stmt
}

func SetAllowComments(yylex interface{}, allow bool) {
	yylex.(*Tokenizer).AllowComments = allow
}

func ForceEOF(yylex interface{}) {
	yylex.(*Tokenizer).ForceEOF = true
}

var (
	SHARE        = []byte("share")
	MODE         = []byte("mode")
	IF_BYTES     = []byte("if")
	VALUES_BYTES = []byte("values")
)

//line sql.y:31
type yySymType struct {
	yys         int
	empty       struct{}
	statement   Statement
	selStmt     SelectStatement
	byt         byte
	bytes       []byte
	bytes2      [][]byte
	str         string
	selectExprs SelectExprs
	selectExpr  SelectExpr
	columns     Columns
	colName     *ColName
	tableExprs  TableExprs
	tableExpr   TableExpr
	smTableExpr SimpleTableExpr
	tableName   *TableName
	indexHints  *IndexHints
	expr        Expr
	boolExpr    BoolExpr
	valExpr     ValExpr
	colTuple    ColTuple
	valExprs    ValExprs
	values      Values
	rowTuple    RowTuple
	subquery    *Subquery
	caseExpr    *CaseExpr
	whens       []*When
	when        *When
	orderBy     OrderBy
	order       *Order
	limit       *Limit
	insRows     InsertRows
	updateExprs UpdateExprs
	updateExpr  *UpdateExpr
}

const LEX_ERROR = 57346
const SELECT = 57347
const INSERT = 57348
const UPDATE = 57349
const DELETE = 57350
const REPLACE = 57351
const FROM = 57352
const WHERE = 57353
const GROUP = 57354
const HAVING = 57355
const ORDER = 57356
const BY = 57357
const LIMIT = 57358
const FOR = 57359
const ALL = 57360
const DISTINCT = 57361
const AS = 57362
const EXISTS = 57363
const IN = 57364
const IS = 57365
const LIKE = 57366
const BETWEEN = 57367
const NULL = 57368
const ASC = 57369
const DESC = 57370
const VALUES = 57371
const INTO = 57372
const DUPLICATE = 57373
const KEY = 57374
const DEFAULT = 57375
const SET = 57376
const LOCK = 57377
const KEYRANGE = 57378
const ID = 57379
const STRING = 57380
const NUMBER = 57381
const VALUE_ARG = 57382
const LIST_ARG = 57383
const COMMENT = 57384
const LE = 57385
const GE = 57386
const NE = 57387
const NULL_SAFE_EQUAL = 57388
const UNION = 57389
const MINUS = 57390
const EXCEPT = 57391
const INTERSECT = 57392
const JOIN = 57393
const STRAIGHT_JOIN = 57394
const LEFT = 57395
const RIGHT = 57396
const INNER = 57397
const OUTER = 57398
const CROSS = 57399
const NATURAL = 57400
const USE = 57401
const FORCE = 57402
const ON = 57403
const OR = 57404
const AND = 57405
const NOT = 57406
const UNARY = 57407
const CASE = 57408
const WHEN = 57409
const THEN = 57410
const ELSE = 57411
const END = 57412
const CREATE = 57413
const ALTER = 57414
const DROP = 57415
const RENAME = 57416
const ANALYZE = 57417
const TABLE = 57418
const INDEX = 57419
const VIEW = 57420
const TO = 57421
const IGNORE = 57422
const IF = 57423
const UNIQUE = 57424
const USING = 57425
const SHOW = 57426
const DESCRIBE = 57427
const EXPLAIN = 57428

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"LEX_ERROR",
	"SELECT",
	"INSERT",
	"UPDATE",
	"DELETE",
	"REPLACE",
	"FROM",
	"WHERE",
	"GROUP",
	"HAVING",
	"ORDER",
	"BY",
	"LIMIT",
	"FOR",
	"ALL",
	"DISTINCT",
	"AS",
	"EXISTS",
	"IN",
	"IS",
	"LIKE",
	"BETWEEN",
	"NULL",
	"ASC",
	"DESC",
	"VALUES",
	"INTO",
	"DUPLICATE",
	"KEY",
	"DEFAULT",
	"SET",
	"LOCK",
	"KEYRANGE",
	"ID",
	"STRING",
	"NUMBER",
	"VALUE_ARG",
	"LIST_ARG",
	"COMMENT",
	"LE",
	"GE",
	"NE",
	"NULL_SAFE_EQUAL",
	"'('",
	"'='",
	"'<'",
	"'>'",
	"'~'",
	"UNION",
	"MINUS",
	"EXCEPT",
	"INTERSECT",
	"','",
	"JOIN",
	"STRAIGHT_JOIN",
	"LEFT",
	"RIGHT",
	"INNER",
	"OUTER",
	"CROSS",
	"NATURAL",
	"USE",
	"FORCE",
	"ON",
	"OR",
	"AND",
	"NOT",
	"'&'",
	"'|'",
	"'^'",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'%'",
	"'.'",
	"UNARY",
	"CASE",
	"WHEN",
	"THEN",
	"ELSE",
	"END",
	"CREATE",
	"ALTER",
	"DROP",
	"RENAME",
	"ANALYZE",
	"TABLE",
	"INDEX",
	"VIEW",
	"TO",
	"IGNORE",
	"IF",
	"UNIQUE",
	"USING",
	"SHOW",
	"DESCRIBE",
	"EXPLAIN",
	"')'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 205
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 615

var yyAct = [...]int{

	98, 300, 164, 66, 337, 89, 370, 254, 95, 167,
	96, 204, 292, 245, 94, 215, 67, 184, 166, 3,
	379, 379, 84, 85, 141, 140, 53, 350, 265, 266,
	267, 268, 269, 107, 270, 271, 30, 31, 32, 33,
	192, 69, 80, 379, 74, 135, 236, 77, 56, 68,
	72, 81, 261, 129, 54, 55, 298, 135, 234, 135,
	236, 90, 46, 41, 47, 43, 381, 380, 348, 44,
	49, 50, 51, 321, 125, 347, 76, 317, 319, 246,
	346, 328, 52, 133, 73, 48, 235, 276, 137, 378,
	198, 326, 323, 141, 140, 246, 168, 290, 163, 165,
	169, 126, 297, 286, 128, 284, 237, 318, 330, 196,
	153, 154, 155, 199, 139, 122, 177, 69, 117, 69,
	181, 69, 183, 188, 187, 68, 224, 68, 75, 68,
	151, 152, 153, 154, 155, 173, 124, 186, 90, 210,
	188, 141, 140, 140, 343, 214, 212, 213, 222, 223,
	189, 226, 227, 228, 229, 230, 231, 232, 233, 209,
	202, 357, 358, 195, 197, 194, 293, 211, 208, 257,
	225, 120, 132, 238, 90, 90, 62, 217, 345, 311,
	69, 69, 293, 250, 312, 344, 315, 309, 68, 252,
	243, 256, 310, 258, 240, 242, 314, 249, 313, 253,
	30, 31, 32, 33, 185, 148, 149, 150, 151, 152,
	153, 154, 155, 120, 134, 64, 236, 185, 275, 238,
	259, 262, 355, 279, 280, 332, 277, 15, 16, 17,
	19, 18, 287, 179, 121, 278, 15, 207, 115, 283,
	208, 365, 119, 79, 90, 364, 180, 206, 15, 263,
	363, 218, 291, 217, 170, 175, 20, 216, 289, 295,
	135, 299, 120, 296, 285, 174, 172, 171, 207, 106,
	75, 114, 112, 70, 322, 307, 308, 320, 206, 274,
	70, 103, 104, 105, 325, 265, 266, 267, 268, 269,
	170, 270, 271, 329, 110, 82, 273, 208, 208, 69,
	304, 334, 327, 138, 335, 338, 303, 333, 21, 22,
	24, 23, 25, 201, 200, 182, 130, 108, 109, 127,
	75, 26, 27, 28, 113, 123, 63, 349, 106, 83,
	339, 78, 376, 351, 118, 116, 352, 331, 15, 111,
	103, 104, 105, 353, 61, 238, 282, 360, 359, 362,
	377, 219, 361, 220, 221, 59, 383, 367, 338, 190,
	368, 369, 248, 131, 371, 371, 371, 69, 57, 301,
	374, 372, 373, 342, 302, 68, 34, 255, 341, 306,
	384, 241, 185, 101, 385, 65, 386, 382, 106, 366,
	15, 112, 35, 36, 37, 38, 39, 40, 102, 88,
	103, 104, 105, 354, 191, 42, 260, 193, 45, 93,
	71, 251, 178, 110, 375, 356, 336, 340, 148, 149,
	150, 151, 152, 153, 154, 155, 305, 288, 176, 244,
	100, 97, 92, 99, 294, 247, 108, 109, 86, 142,
	91, 316, 205, 113, 101, 264, 203, 87, 272, 106,
	136, 58, 112, 29, 60, 14, 13, 12, 111, 102,
	88, 103, 104, 105, 239, 11, 10, 9, 8, 106,
	93, 7, 112, 6, 110, 5, 4, 2, 1, 15,
	70, 103, 104, 105, 0, 0, 0, 0, 0, 0,
	170, 0, 0, 92, 110, 101, 0, 108, 109, 86,
	106, 0, 0, 112, 113, 0, 0, 0, 0, 0,
	102, 70, 103, 104, 105, 101, 0, 108, 109, 111,
	106, 93, 0, 112, 113, 110, 0, 0, 0, 0,
	102, 70, 103, 104, 105, 0, 0, 0, 0, 111,
	0, 93, 0, 0, 92, 110, 0, 0, 108, 109,
	143, 147, 145, 146, 324, 113, 148, 149, 150, 151,
	152, 153, 154, 155, 92, 0, 0, 0, 108, 109,
	111, 159, 160, 161, 162, 113, 156, 157, 158, 281,
	0, 148, 149, 150, 151, 152, 153, 154, 155, 0,
	111, 0, 0, 0, 0, 0, 0, 0, 144, 148,
	149, 150, 151, 152, 153, 154, 155, 148, 149, 150,
	151, 152, 153, 154, 155,
}
var yyPact = [...]int{

	222, -1000, -1000, 148, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -28, -31, -6, -21, -9, -1000, -1000, -1000, 385,
	350, -1000, -1000, -1000, 336, -1000, 314, 289, 289, 375,
	236, -46, -8, 233, -1000, -15, 233, -1000, 294, -54,
	233, -54, 292, -1000, -1000, -1000, -1000, -1000, 423, -1000,
	229, 289, 301, 39, 300, 289, 157, -1000, 186, -1000,
	36, 288, 66, 233, -1000, -1000, 282, -1000, -41, 279,
	342, 105, 233, -1000, 204, -1000, -1000, 283, 35, 73,
	528, -1000, 494, 474, -1000, -1000, -1000, 443, 220, 219,
	-1000, 218, 208, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 443, -1000, 199, 236, 278, 236, 371,
	236, 443, 233, -1000, 338, -58, -1000, 76, -1000, 277,
	-1000, -1000, 276, -1000, 200, 423, -1000, -1000, 233, 91,
	494, 494, 443, 210, 329, 443, 443, 100, 443, 443,
	443, 443, 443, 443, 443, 443, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 528, -44, -16, 4, 528, -1000,
	243, 362, 423, -1000, 385, 302, -3, 536, 333, 236,
	236, 206, -1000, 157, 363, 494, -1000, 536, -1000, -1000,
	-1000, 102, 233, -1000, -42, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 193, 228, 259, 231, 8, -1000, -1000,
	-1000, -1000, -1000, 74, 536, -1000, 243, -1000, -1000, 210,
	443, 443, 536, 510, -1000, 320, 56, 56, 56, 34,
	34, -1000, -1000, -1000, -1000, -1000, 443, -1000, 536, -1000,
	3, 423, 1, 176, 13, -1000, 494, 99, 207, 148,
	115, 0, -1000, 363, 353, 359, 73, 269, -1000, -1000,
	263, -1000, 367, 200, 200, -1000, -1000, 130, 122, 141,
	139, 129, 12, -1000, 240, -29, 237, -10, -1000, 536,
	485, 443, -1000, 536, -1000, -11, -1000, 302, -4, -1000,
	443, 25, -1000, 306, 169, -1000, -1000, -1000, 236, 353,
	-1000, 443, 443, -1000, -1000, 365, 358, 228, 77, -1000,
	128, -1000, 121, -1000, -1000, -1000, -1000, -12, -17, -24,
	-1000, -1000, -1000, -1000, 443, 536, -1000, -75, -1000, 536,
	443, 304, 207, -1000, -1000, 347, 166, -1000, 134, -1000,
	363, 494, 443, 494, -1000, -1000, 203, 198, 194, 536,
	-1000, 536, 382, -1000, 443, 443, -1000, -1000, -1000, 353,
	73, 160, 73, 233, 233, 233, 236, 536, -1000, 315,
	-13, -1000, -35, -36, 157, -1000, 380, 334, -1000, 233,
	-1000, -1000, -1000, 233, -1000, 233, -1000,
}
var yyPgo = [...]int{

	0, 478, 477, 18, 476, 475, 473, 471, 468, 467,
	466, 465, 457, 456, 455, 376, 454, 453, 451, 22,
	23, 450, 448, 447, 446, 11, 445, 442, 176, 441,
	6, 17, 5, 440, 439, 435, 14, 2, 15, 9,
	434, 10, 433, 33, 431, 8, 430, 429, 13, 428,
	427, 426, 417, 7, 416, 4, 415, 1, 414, 412,
	411, 12, 3, 16, 243, 410, 408, 407, 406, 405,
	404, 0, 26, 392,
}
var yyR1 = [...]int{

	0, 1, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 3, 3, 4, 4, 5, 6,
	7, 8, 9, 9, 9, 10, 10, 10, 11, 12,
	12, 12, 13, 14, 14, 14, 73, 15, 16, 16,
	17, 17, 17, 17, 17, 18, 18, 19, 19, 20,
	20, 20, 23, 23, 21, 21, 21, 24, 24, 25,
	25, 25, 25, 22, 22, 22, 26, 26, 26, 26,
	26, 26, 26, 26, 26, 27, 27, 27, 28, 28,
	29, 29, 29, 29, 30, 30, 31, 31, 32, 32,
	32, 32, 32, 33, 33, 33, 33, 33, 33, 33,
	33, 33, 33, 33, 34, 34, 34, 34, 34, 34,
	34, 38, 38, 38, 43, 39, 39, 37, 37, 37,
	37, 37, 37, 37, 37, 37, 37, 37, 37, 37,
	37, 37, 37, 37, 42, 42, 44, 44, 44, 46,
	49, 49, 47, 47, 48, 50, 50, 45, 45, 36,
	36, 36, 36, 51, 51, 52, 52, 53, 53, 54,
	54, 55, 56, 56, 56, 57, 57, 57, 58, 58,
	58, 59, 59, 60, 60, 61, 61, 35, 35, 40,
	40, 41, 41, 62, 62, 63, 64, 64, 65, 65,
	66, 66, 67, 67, 67, 67, 67, 68, 68, 69,
	69, 70, 70, 71, 72,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 12, 3, 7, 7, 8, 5,
	7, 3, 5, 8, 4, 6, 7, 4, 5, 4,
	5, 5, 3, 2, 2, 2, 0, 2, 0, 2,
	1, 2, 1, 1, 1, 0, 1, 1, 3, 1,
	2, 3, 1, 1, 0, 1, 2, 1, 3, 3,
	3, 3, 5, 0, 1, 2, 1, 1, 2, 3,
	2, 3, 2, 2, 2, 1, 3, 1, 1, 3,
	0, 5, 5, 5, 1, 3, 0, 2, 1, 3,
	3, 2, 3, 3, 3, 4, 3, 4, 5, 6,
	3, 4, 2, 6, 1, 1, 1, 1, 1, 1,
	1, 3, 1, 1, 3, 1, 3, 1, 1, 1,
	3, 3, 3, 3, 3, 3, 3, 3, 2, 3,
	4, 5, 4, 1, 1, 1, 1, 1, 1, 5,
	0, 1, 1, 2, 4, 0, 2, 1, 3, 1,
	1, 1, 1, 0, 3, 0, 2, 0, 3, 1,
	3, 2, 0, 1, 1, 0, 2, 4, 0, 2,
	4, 0, 3, 1, 3, 0, 5, 2, 1, 1,
	3, 3, 1, 1, 3, 3, 0, 2, 0, 3,
	0, 1, 1, 1, 1, 1, 1, 0, 1, 0,
	1, 0, 2, 1, 0,
}
var yyChk = [...]int{

	-1000, -1, -2, -3, -4, -5, -6, -7, -8, -9,
	-10, -11, -12, -13, -14, 5, 6, 7, 9, 8,
	34, 86, 87, 89, 88, 90, 99, 100, 101, -17,
	52, 53, 54, 55, -15, -73, -15, -15, -15, -15,
	-15, 91, -69, 93, 97, -66, 93, 95, 91, 91,
	92, 93, 91, -72, -72, -72, -3, 18, -18, 19,
	-16, 30, -28, 37, -28, 10, -62, -63, -45, -71,
	37, -65, 96, 92, -71, 37, 91, -71, 37, -64,
	96, -71, -64, 37, -19, -20, 76, -23, 37, -32,
	-37, -33, 70, 47, -36, -45, -41, -44, -71, -42,
	-46, 21, 36, 38, 39, 40, 26, -43, 74, 75,
	51, 96, 29, 81, 42, -28, 34, 79, 34, -28,
	56, 48, 79, 37, 70, -71, -72, 37, -72, 94,
	37, 21, 67, -71, 10, 56, -21, -71, 20, 79,
	69, 68, -34, 22, 70, 24, 25, 23, 71, 72,
	73, 74, 75, 76, 77, 78, 48, 49, 50, 43,
	44, 45, 46, -32, -37, -32, -3, -39, -37, -37,
	47, 47, 47, -43, 47, 47, -49, -37, -59, 34,
	47, -62, 37, -62, -31, 11, -63, -37, -71, -72,
	21, -70, 98, -67, 89, 87, 33, 88, 14, 37,
	37, 37, -72, -24, -25, -27, 47, 37, -43, -20,
	-71, 76, -32, -32, -37, -38, 47, -43, 41, 22,
	24, 25, -37, -37, 26, 70, -37, -37, -37, -37,
	-37, -37, -37, -37, 102, 102, 56, 102, -37, 102,
	-19, 19, -19, -36, -47, -48, 82, -35, 29, -3,
	-62, -60, -45, -31, -53, 14, -32, 67, -71, -72,
	-68, 94, -31, 56, -26, 57, 58, 59, 60, 61,
	63, 64, -22, 37, 20, -25, 79, -39, -38, -37,
	-37, 69, 26, -37, 102, -19, 102, 56, -50, -48,
	84, -32, -61, 67, -40, -41, -61, 102, 56, -53,
	-57, 16, 15, 37, 37, -51, 12, -25, -25, 57,
	62, 57, 62, 57, 57, 57, -29, 65, 95, 66,
	37, 102, 37, 102, 69, -37, 102, -36, 85, -37,
	83, 31, 56, -45, -57, -37, -54, -55, -37, -72,
	-52, 13, 15, 67, 57, 57, 92, 92, 92, -37,
	102, -37, 32, -41, 56, 56, -56, 27, 28, -53,
	-32, -39, -32, 47, 47, 47, 7, -37, -55, -57,
	-30, -71, -30, -30, -62, -58, 17, 35, 102, 56,
	102, 102, 7, 22, -71, -71, -71,
}
var yyDef = [...]int{

	0, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 12, 13, 36, 36, 36, 36, 36,
	36, 199, 190, 0, 0, 0, 204, 204, 204, 0,
	40, 42, 43, 44, 45, 38, 0, 0, 0, 0,
	0, 188, 0, 0, 200, 0, 0, 191, 0, 186,
	0, 186, 0, 33, 34, 35, 15, 41, 0, 46,
	37, 0, 0, 78, 0, 0, 21, 183, 0, 147,
	203, 0, 0, 0, 204, 203, 0, 204, 0, 0,
	0, 0, 0, 32, 0, 47, 49, 54, 203, 52,
	53, 88, 0, 0, 117, 118, 119, 0, 147, 0,
	133, 0, 0, 149, 150, 151, 152, 182, 136, 137,
	138, 134, 135, 140, 39, 171, 0, 0, 0, 86,
	0, 0, 0, 204, 0, 201, 24, 0, 27, 0,
	29, 187, 0, 204, 0, 0, 50, 55, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 104, 105, 106, 107,
	108, 109, 110, 91, 0, 0, 0, 0, 115, 128,
	0, 0, 0, 102, 0, 0, 0, 141, 0, 0,
	0, 86, 79, 19, 157, 0, 184, 185, 148, 22,
	189, 0, 0, 204, 197, 192, 193, 194, 195, 196,
	28, 30, 31, 86, 57, 63, 0, 75, 77, 48,
	56, 51, 89, 90, 93, 94, 0, 112, 113, 0,
	0, 0, 96, 0, 100, 0, 120, 121, 122, 123,
	124, 125, 126, 127, 92, 114, 0, 181, 115, 129,
	0, 0, 0, 0, 145, 142, 0, 175, 0, 178,
	175, 0, 173, 157, 165, 0, 87, 0, 202, 25,
	0, 198, 153, 0, 0, 66, 67, 0, 0, 0,
	0, 0, 80, 64, 0, 0, 0, 0, 95, 97,
	0, 0, 101, 116, 130, 0, 132, 0, 0, 143,
	0, 0, 16, 0, 177, 179, 17, 172, 0, 165,
	20, 0, 0, 204, 26, 155, 0, 58, 61, 68,
	0, 70, 0, 72, 73, 74, 59, 0, 0, 0,
	65, 60, 76, 111, 0, 98, 131, 0, 139, 146,
	0, 0, 0, 174, 18, 166, 158, 159, 162, 23,
	157, 0, 0, 0, 69, 71, 0, 0, 0, 99,
	103, 144, 0, 180, 0, 0, 161, 163, 164, 165,
	156, 154, 62, 0, 0, 0, 0, 167, 160, 168,
	0, 84, 0, 0, 176, 14, 0, 0, 81, 0,
	82, 83, 169, 0, 85, 0, 170,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 78, 71, 3,
	47, 102, 76, 74, 56, 75, 79, 77, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	49, 48, 50, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 73, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 72, 3, 51,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 52, 53, 54, 55, 57,
	58, 59, 60, 61, 62, 63, 64, 65, 66, 67,
	68, 69, 70, 80, 81, 82, 83, 84, 85, 86,
	87, 88, 89, 90, 91, 92, 93, 94, 95, 96,
	97, 98, 99, 100, 101,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:151
		{
			SetParseTree(yylex, yyDollar[1].statement)
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:157
		{
			yyVAL.statement = yyDollar[1].selStmt
		}
	case 14:
		yyDollar = yyS[yypt-12 : yypt+1]
		//line sql.y:174
		{
			yyVAL.selStmt = &Select{Comments: Comments(yyDollar[2].bytes2), Distinct: yyDollar[3].str, SelectExprs: yyDollar[4].selectExprs, From: yyDollar[6].tableExprs, Where: NewWhere(AST_WHERE, yyDollar[7].boolExpr), GroupBy: GroupBy(yyDollar[8].valExprs), Having: NewWhere(AST_HAVING, yyDollar[9].boolExpr), OrderBy: yyDollar[10].orderBy, Limit: yyDollar[11].limit, Lock: yyDollar[12].str}
		}
	case 15:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:178
		{
			yyVAL.selStmt = &Union{Type: yyDollar[2].str, Left: yyDollar[1].selStmt, Right: yyDollar[3].selStmt}
		}
	case 16:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:184
		{
			yyVAL.statement = &Insert{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Columns: yyDollar[5].columns, Rows: yyDollar[6].insRows, OnDup: OnDup(yyDollar[7].updateExprs)}
		}
	case 17:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:188
		{
			cols := make(Columns, 0, len(yyDollar[6].updateExprs))
			vals := make(ValTuple, 0, len(yyDollar[6].updateExprs))
			for _, col := range yyDollar[6].updateExprs {
				cols = append(cols, &NonStarExpr{Expr: col.Name})
				vals = append(vals, col.Expr)
			}
			yyVAL.statement = &Insert{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Columns: cols, Rows: Values{vals}, OnDup: OnDup(yyDollar[7].updateExprs)}
		}
	case 18:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line sql.y:200
		{
			yyVAL.statement = &Update{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[3].tableName, Exprs: yyDollar[5].updateExprs, Where: NewWhere(AST_WHERE, yyDollar[6].boolExpr), OrderBy: yyDollar[7].orderBy, Limit: yyDollar[8].limit}
		}
	case 19:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:206
		{
			yyVAL.statement = &Replace{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[3].tableName, Exprs: yyDollar[5].updateExprs}
		}
	case 20:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:212
		{
			yyVAL.statement = &Delete{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Where: NewWhere(AST_WHERE, yyDollar[5].boolExpr), OrderBy: yyDollar[6].orderBy, Limit: yyDollar[7].limit}
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:218
		{
			yyVAL.statement = &Set{Comments: Comments(yyDollar[2].bytes2), Exprs: yyDollar[3].updateExprs}
		}
	case 22:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:224
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyDollar[4].bytes}
		}
	case 23:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line sql.y:228
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[7].bytes, NewName: yyDollar[7].bytes}
		}
	case 24:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:233
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyDollar[3].bytes}
		}
	case 25:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:239
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[4].bytes, NewName: yyDollar[4].bytes}
		}
	case 26:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:243
		{
			// Change this to a rename statement
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyDollar[4].bytes, NewName: yyDollar[7].bytes}
		}
	case 27:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:248
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[3].bytes, NewName: yyDollar[3].bytes}
		}
	case 28:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:254
		{
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyDollar[3].bytes, NewName: yyDollar[5].bytes}
		}
	case 29:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:260
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyDollar[4].bytes}
		}
	case 30:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:264
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[5].bytes, NewName: yyDollar[5].bytes}
		}
	case 31:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:269
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyDollar[4].bytes}
		}
	case 32:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:275
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[3].bytes, NewName: yyDollar[3].bytes}
		}
	case 33:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:281
		{
			yyVAL.statement = &Other{}
		}
	case 34:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:285
		{
			yyVAL.statement = &Other{}
		}
	case 35:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:289
		{
			yyVAL.statement = &Other{}
		}
	case 36:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:294
		{
			SetAllowComments(yylex, true)
		}
	case 37:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:298
		{
			yyVAL.bytes2 = yyDollar[2].bytes2
			SetAllowComments(yylex, false)
		}
	case 38:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:304
		{
			yyVAL.bytes2 = nil
		}
	case 39:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:308
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[2].bytes)
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:314
		{
			yyVAL.str = AST_UNION
		}
	case 41:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:318
		{
			yyVAL.str = AST_UNION_ALL
		}
	case 42:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:322
		{
			yyVAL.str = AST_SET_MINUS
		}
	case 43:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:326
		{
			yyVAL.str = AST_EXCEPT
		}
	case 44:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:330
		{
			yyVAL.str = AST_INTERSECT
		}
	case 45:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:335
		{
			yyVAL.str = ""
		}
	case 46:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:339
		{
			yyVAL.str = AST_DISTINCT
		}
	case 47:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:345
		{
			yyVAL.selectExprs = SelectExprs{yyDollar[1].selectExpr}
		}
	case 48:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:349
		{
			yyVAL.selectExprs = append(yyVAL.selectExprs, yyDollar[3].selectExpr)
		}
	case 49:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:355
		{
			yyVAL.selectExpr = &StarExpr{}
		}
	case 50:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:359
		{
			yyVAL.selectExpr = &NonStarExpr{Expr: yyDollar[1].expr, As: yyDollar[2].bytes}
		}
	case 51:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:363
		{
			yyVAL.selectExpr = &StarExpr{TableName: yyDollar[1].bytes}
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:369
		{
			yyVAL.expr = yyDollar[1].boolExpr
		}
	case 53:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:373
		{
			yyVAL.expr = yyDollar[1].valExpr
		}
	case 54:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:378
		{
			yyVAL.bytes = nil
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:382
		{
			yyVAL.bytes = yyDollar[1].bytes
		}
	case 56:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:386
		{
			yyVAL.bytes = yyDollar[2].bytes
		}
	case 57:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:392
		{
			yyVAL.tableExprs = TableExprs{yyDollar[1].tableExpr}
		}
	case 58:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:396
		{
			yyVAL.tableExprs = append(yyVAL.tableExprs, yyDollar[3].tableExpr)
		}
	case 59:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:402
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyDollar[1].smTableExpr, As: yyDollar[2].bytes, Hints: yyDollar[3].indexHints}
		}
	case 60:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:406
		{
			yyVAL.tableExpr = &ParenTableExpr{Expr: yyDollar[2].tableExpr}
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:410
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr}
		}
	case 62:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:414
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr, On: yyDollar[5].boolExpr}
		}
	case 63:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:419
		{
			yyVAL.bytes = nil
		}
	case 64:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:423
		{
			yyVAL.bytes = yyDollar[1].bytes
		}
	case 65:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:427
		{
			yyVAL.bytes = yyDollar[2].bytes
		}
	case 66:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:433
		{
			yyVAL.str = AST_JOIN
		}
	case 67:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:437
		{
			yyVAL.str = AST_STRAIGHT_JOIN
		}
	case 68:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:441
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 69:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:445
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 70:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:449
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 71:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:453
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 72:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:457
		{
			yyVAL.str = AST_JOIN
		}
	case 73:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:461
		{
			yyVAL.str = AST_CROSS_JOIN
		}
	case 74:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:465
		{
			yyVAL.str = AST_NATURAL_JOIN
		}
	case 75:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:471
		{
			yyVAL.smTableExpr = &TableName{Name: yyDollar[1].bytes}
		}
	case 76:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:475
		{
			yyVAL.smTableExpr = &TableName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 77:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:479
		{
			yyVAL.smTableExpr = yyDollar[1].subquery
		}
	case 78:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:485
		{
			yyVAL.tableName = &TableName{Name: yyDollar[1].bytes}
		}
	case 79:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:489
		{
			yyVAL.tableName = &TableName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 80:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:494
		{
			yyVAL.indexHints = nil
		}
	case 81:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:498
		{
			yyVAL.indexHints = &IndexHints{Type: AST_USE, Indexes: yyDollar[4].bytes2}
		}
	case 82:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:502
		{
			yyVAL.indexHints = &IndexHints{Type: AST_IGNORE, Indexes: yyDollar[4].bytes2}
		}
	case 83:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:506
		{
			yyVAL.indexHints = &IndexHints{Type: AST_FORCE, Indexes: yyDollar[4].bytes2}
		}
	case 84:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:512
		{
			yyVAL.bytes2 = [][]byte{yyDollar[1].bytes}
		}
	case 85:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:516
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[3].bytes)
		}
	case 86:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:521
		{
			yyVAL.boolExpr = nil
		}
	case 87:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:525
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 89:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:532
		{
			yyVAL.boolExpr = &AndExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 90:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:536
		{
			yyVAL.boolExpr = &OrExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 91:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:540
		{
			yyVAL.boolExpr = &NotExpr{Expr: yyDollar[2].boolExpr}
		}
	case 92:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:544
		{
			yyVAL.boolExpr = &ParenBoolExpr{Expr: yyDollar[2].boolExpr}
		}
	case 93:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:550
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: yyDollar[2].str, Right: yyDollar[3].valExpr}
		}
	case 94:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:554
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_IN, Right: yyDollar[3].colTuple}
		}
	case 95:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:558
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_NOT_IN, Right: yyDollar[4].colTuple}
		}
	case 96:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:562
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_LIKE, Right: yyDollar[3].valExpr}
		}
	case 97:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:566
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_NOT_LIKE, Right: yyDollar[4].valExpr}
		}
	case 98:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:570
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: AST_BETWEEN, From: yyDollar[3].valExpr, To: yyDollar[5].valExpr}
		}
	case 99:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:574
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: AST_NOT_BETWEEN, From: yyDollar[4].valExpr, To: yyDollar[6].valExpr}
		}
	case 100:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:578
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NULL, Expr: yyDollar[1].valExpr}
		}
	case 101:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:582
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NOT_NULL, Expr: yyDollar[1].valExpr}
		}
	case 102:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:586
		{
			yyVAL.boolExpr = &ExistsExpr{Subquery: yyDollar[2].subquery}
		}
	case 103:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:590
		{
			yyVAL.boolExpr = &KeyrangeExpr{Start: yyDollar[3].valExpr, End: yyDollar[5].valExpr}
		}
	case 104:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:596
		{
			yyVAL.str = AST_EQ
		}
	case 105:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:600
		{
			yyVAL.str = AST_LT
		}
	case 106:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:604
		{
			yyVAL.str = AST_GT
		}
	case 107:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:608
		{
			yyVAL.str = AST_LE
		}
	case 108:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:612
		{
			yyVAL.str = AST_GE
		}
	case 109:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:616
		{
			yyVAL.str = AST_NE
		}
	case 110:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:620
		{
			yyVAL.str = AST_NSE
		}
	case 111:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:626
		{
			yyVAL.colTuple = ValTuple(yyDollar[2].valExprs)
		}
	case 112:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:630
		{
			yyVAL.colTuple = yyDollar[1].subquery
		}
	case 113:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:634
		{
			yyVAL.colTuple = ListArg(yyDollar[1].bytes)
		}
	case 114:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:640
		{
			yyVAL.subquery = &Subquery{yyDollar[2].selStmt}
		}
	case 115:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:646
		{
			yyVAL.valExprs = ValExprs{yyDollar[1].valExpr}
		}
	case 116:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:650
		{
			yyVAL.valExprs = append(yyDollar[1].valExprs, yyDollar[3].valExpr)
		}
	case 117:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:656
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 118:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:660
		{
			yyVAL.valExpr = yyDollar[1].colName
		}
	case 119:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:664
		{
			yyVAL.valExpr = yyDollar[1].rowTuple
		}
	case 120:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:668
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITAND, Right: yyDollar[3].valExpr}
		}
	case 121:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:672
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITOR, Right: yyDollar[3].valExpr}
		}
	case 122:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:676
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITXOR, Right: yyDollar[3].valExpr}
		}
	case 123:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:680
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_PLUS, Right: yyDollar[3].valExpr}
		}
	case 124:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:684
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MINUS, Right: yyDollar[3].valExpr}
		}
	case 125:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:688
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MULT, Right: yyDollar[3].valExpr}
		}
	case 126:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:692
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_DIV, Right: yyDollar[3].valExpr}
		}
	case 127:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:696
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MOD, Right: yyDollar[3].valExpr}
		}
	case 128:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:700
		{
			if num, ok := yyDollar[2].valExpr.(NumVal); ok {
				switch yyDollar[1].byt {
				case '-':
					yyVAL.valExpr = append(NumVal("-"), num...)
				case '+':
					yyVAL.valExpr = num
				default:
					yyVAL.valExpr = &UnaryExpr{Operator: yyDollar[1].byt, Expr: yyDollar[2].valExpr}
				}
			} else {
				yyVAL.valExpr = &UnaryExpr{Operator: yyDollar[1].byt, Expr: yyDollar[2].valExpr}
			}
		}
	case 129:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:715
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes}
		}
	case 130:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:719
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Exprs: yyDollar[3].selectExprs}
		}
	case 131:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:723
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Distinct: true, Exprs: yyDollar[4].selectExprs}
		}
	case 132:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:727
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Exprs: yyDollar[3].selectExprs}
		}
	case 133:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:731
		{
			yyVAL.valExpr = yyDollar[1].caseExpr
		}
	case 134:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:737
		{
			yyVAL.bytes = IF_BYTES
		}
	case 135:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:741
		{
			yyVAL.bytes = VALUES_BYTES
		}
	case 136:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:747
		{
			yyVAL.byt = AST_UPLUS
		}
	case 137:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:751
		{
			yyVAL.byt = AST_UMINUS
		}
	case 138:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:755
		{
			yyVAL.byt = AST_TILDA
		}
	case 139:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:761
		{
			yyVAL.caseExpr = &CaseExpr{Expr: yyDollar[2].valExpr, Whens: yyDollar[3].whens, Else: yyDollar[4].valExpr}
		}
	case 140:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:766
		{
			yyVAL.valExpr = nil
		}
	case 141:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:770
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 142:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:776
		{
			yyVAL.whens = []*When{yyDollar[1].when}
		}
	case 143:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:780
		{
			yyVAL.whens = append(yyDollar[1].whens, yyDollar[2].when)
		}
	case 144:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:786
		{
			yyVAL.when = &When{Cond: yyDollar[2].boolExpr, Val: yyDollar[4].valExpr}
		}
	case 145:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:791
		{
			yyVAL.valExpr = nil
		}
	case 146:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:795
		{
			yyVAL.valExpr = yyDollar[2].valExpr
		}
	case 147:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:801
		{
			yyVAL.colName = &ColName{Name: yyDollar[1].bytes}
		}
	case 148:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:805
		{
			yyVAL.colName = &ColName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 149:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:811
		{
			yyVAL.valExpr = StrVal(yyDollar[1].bytes)
		}
	case 150:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:815
		{
			yyVAL.valExpr = NumVal(yyDollar[1].bytes)
		}
	case 151:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:819
		{
			yyVAL.valExpr = ValArg(yyDollar[1].bytes)
		}
	case 152:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:823
		{
			yyVAL.valExpr = &NullVal{}
		}
	case 153:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:828
		{
			yyVAL.valExprs = nil
		}
	case 154:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:832
		{
			yyVAL.valExprs = yyDollar[3].valExprs
		}
	case 155:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:837
		{
			yyVAL.boolExpr = nil
		}
	case 156:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:841
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 157:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:846
		{
			yyVAL.orderBy = nil
		}
	case 158:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:850
		{
			yyVAL.orderBy = yyDollar[3].orderBy
		}
	case 159:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:856
		{
			yyVAL.orderBy = OrderBy{yyDollar[1].order}
		}
	case 160:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:860
		{
			yyVAL.orderBy = append(yyDollar[1].orderBy, yyDollar[3].order)
		}
	case 161:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:866
		{
			yyVAL.order = &Order{Expr: yyDollar[1].valExpr, Direction: yyDollar[2].str}
		}
	case 162:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:871
		{
			yyVAL.str = AST_ASC
		}
	case 163:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:875
		{
			yyVAL.str = AST_ASC
		}
	case 164:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:879
		{
			yyVAL.str = AST_DESC
		}
	case 165:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:884
		{
			yyVAL.limit = nil
		}
	case 166:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:888
		{
			yyVAL.limit = &Limit{Rowcount: yyDollar[2].valExpr}
		}
	case 167:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:892
		{
			yyVAL.limit = &Limit{Offset: yyDollar[2].valExpr, Rowcount: yyDollar[4].valExpr}
		}
	case 168:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:897
		{
			yyVAL.str = ""
		}
	case 169:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:901
		{
			yyVAL.str = AST_FOR_UPDATE
		}
	case 170:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:905
		{
			if !bytes.Equal(yyDollar[3].bytes, SHARE) {
				yylex.Error("expecting share")
				return 1
			}
			if !bytes.Equal(yyDollar[4].bytes, MODE) {
				yylex.Error("expecting mode")
				return 1
			}
			yyVAL.str = AST_SHARE_MODE
		}
	case 171:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:918
		{
			yyVAL.columns = nil
		}
	case 172:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:922
		{
			yyVAL.columns = yyDollar[2].columns
		}
	case 173:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:928
		{
			yyVAL.columns = Columns{&NonStarExpr{Expr: yyDollar[1].colName}}
		}
	case 174:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:932
		{
			yyVAL.columns = append(yyVAL.columns, &NonStarExpr{Expr: yyDollar[3].colName})
		}
	case 175:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:937
		{
			yyVAL.updateExprs = nil
		}
	case 176:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:941
		{
			yyVAL.updateExprs = yyDollar[5].updateExprs
		}
	case 177:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:947
		{
			yyVAL.insRows = yyDollar[2].values
		}
	case 178:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:951
		{
			yyVAL.insRows = yyDollar[1].selStmt
		}
	case 179:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:957
		{
			yyVAL.values = Values{yyDollar[1].rowTuple}
		}
	case 180:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:961
		{
			yyVAL.values = append(yyDollar[1].values, yyDollar[3].rowTuple)
		}
	case 181:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:967
		{
			yyVAL.rowTuple = ValTuple(yyDollar[2].valExprs)
		}
	case 182:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:971
		{
			yyVAL.rowTuple = yyDollar[1].subquery
		}
	case 183:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:977
		{
			yyVAL.updateExprs = UpdateExprs{yyDollar[1].updateExpr}
		}
	case 184:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:981
		{
			yyVAL.updateExprs = append(yyDollar[1].updateExprs, yyDollar[3].updateExpr)
		}
	case 185:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:987
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyDollar[1].colName, Expr: yyDollar[3].valExpr}
		}
	case 186:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:992
		{
			yyVAL.empty = struct{}{}
		}
	case 187:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:994
		{
			yyVAL.empty = struct{}{}
		}
	case 188:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:997
		{
			yyVAL.empty = struct{}{}
		}
	case 189:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:999
		{
			yyVAL.empty = struct{}{}
		}
	case 190:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1002
		{
			yyVAL.empty = struct{}{}
		}
	case 191:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1004
		{
			yyVAL.empty = struct{}{}
		}
	case 192:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1008
		{
			yyVAL.empty = struct{}{}
		}
	case 193:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1010
		{
			yyVAL.empty = struct{}{}
		}
	case 194:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1012
		{
			yyVAL.empty = struct{}{}
		}
	case 195:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1014
		{
			yyVAL.empty = struct{}{}
		}
	case 196:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1016
		{
			yyVAL.empty = struct{}{}
		}
	case 197:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1019
		{
			yyVAL.empty = struct{}{}
		}
	case 198:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1021
		{
			yyVAL.empty = struct{}{}
		}
	case 199:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1024
		{
			yyVAL.empty = struct{}{}
		}
	case 200:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1026
		{
			yyVAL.empty = struct{}{}
		}
	case 201:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1029
		{
			yyVAL.empty = struct{}{}
		}
	case 202:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1031
		{
			yyVAL.empty = struct{}{}
		}
	case 203:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1035
		{
			yyVAL.bytes = bytes.ToLower(yyDollar[1].bytes)
		}
	case 204:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1040
		{
			ForceEOF(yylex)
		}
	}
	goto yystack /* stack new state and value */
}
