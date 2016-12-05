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

const yyNprod = 207
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 634

var yyAct = [...]int{

	99, 310, 168, 378, 346, 171, 266, 95, 97, 213,
	258, 224, 254, 96, 193, 68, 170, 3, 184, 85,
	117, 145, 144, 86, 30, 31, 32, 33, 67, 359,
	201, 41, 81, 43, 73, 90, 53, 44, 273, 386,
	386, 70, 386, 133, 75, 139, 56, 78, 245, 139,
	139, 82, 264, 245, 69, 243, 46, 357, 47, 356,
	355, 91, 74, 255, 54, 55, 77, 52, 108, 48,
	327, 329, 145, 144, 244, 129, 277, 278, 279, 280,
	281, 288, 282, 283, 137, 388, 387, 340, 385, 141,
	128, 336, 338, 143, 333, 298, 296, 172, 263, 246,
	328, 173, 49, 50, 51, 255, 126, 302, 157, 158,
	159, 144, 130, 120, 352, 132, 259, 181, 269, 70,
	70, 331, 70, 70, 233, 70, 136, 197, 196, 167,
	169, 76, 69, 189, 186, 69, 69, 182, 69, 354,
	195, 353, 91, 219, 197, 325, 207, 187, 324, 223,
	191, 192, 231, 232, 323, 235, 236, 237, 238, 239,
	240, 241, 242, 218, 198, 205, 124, 245, 234, 208,
	220, 177, 145, 144, 211, 62, 138, 247, 91, 91,
	221, 222, 124, 363, 70, 307, 299, 252, 155, 156,
	157, 158, 159, 259, 261, 249, 251, 69, 262, 186,
	321, 256, 270, 319, 125, 322, 265, 217, 320, 15,
	365, 366, 257, 64, 65, 373, 226, 194, 194, 204,
	206, 203, 139, 183, 216, 287, 372, 274, 247, 80,
	268, 289, 291, 292, 215, 371, 119, 116, 174, 271,
	290, 216, 123, 277, 278, 279, 280, 281, 295, 282,
	283, 215, 179, 91, 152, 153, 154, 155, 156, 157,
	158, 159, 275, 124, 178, 70, 301, 304, 305, 118,
	297, 227, 309, 30, 31, 32, 33, 225, 308, 176,
	175, 83, 119, 115, 217, 317, 318, 107, 61, 286,
	76, 303, 142, 71, 122, 63, 335, 226, 121, 104,
	105, 106, 332, 330, 314, 339, 285, 337, 313, 76,
	210, 343, 209, 362, 344, 347, 342, 190, 134, 131,
	127, 63, 84, 15, 16, 17, 19, 18, 152, 153,
	154, 155, 156, 157, 158, 159, 79, 358, 341, 306,
	383, 294, 390, 360, 217, 217, 15, 199, 135, 59,
	348, 57, 20, 311, 247, 351, 367, 369, 384, 312,
	267, 228, 70, 229, 230, 375, 347, 350, 376, 377,
	185, 316, 379, 379, 379, 69, 380, 381, 194, 66,
	389, 361, 250, 15, 102, 35, 368, 391, 370, 107,
	374, 392, 113, 393, 200, 42, 272, 202, 45, 103,
	89, 104, 105, 106, 21, 22, 24, 23, 25, 72,
	94, 188, 382, 364, 111, 345, 349, 26, 27, 28,
	334, 315, 152, 153, 154, 155, 156, 157, 158, 159,
	300, 180, 253, 93, 101, 98, 100, 109, 110, 87,
	260, 146, 92, 326, 114, 102, 214, 276, 212, 15,
	107, 88, 284, 113, 140, 58, 29, 60, 14, 112,
	103, 89, 104, 105, 106, 248, 13, 12, 34, 11,
	107, 94, 10, 113, 9, 111, 8, 7, 6, 5,
	15, 71, 104, 105, 106, 36, 37, 38, 39, 40,
	4, 174, 2, 1, 93, 111, 102, 0, 109, 110,
	87, 107, 0, 0, 113, 114, 0, 0, 0, 0,
	0, 103, 71, 104, 105, 106, 102, 0, 109, 110,
	112, 107, 94, 0, 113, 114, 111, 0, 0, 0,
	0, 103, 71, 104, 105, 106, 0, 0, 0, 0,
	112, 107, 94, 0, 113, 93, 111, 0, 0, 109,
	110, 0, 71, 104, 105, 106, 114, 0, 0, 0,
	0, 0, 174, 0, 0, 93, 111, 0, 0, 109,
	110, 112, 0, 0, 0, 0, 114, 147, 151, 149,
	150, 152, 153, 154, 155, 156, 157, 158, 159, 109,
	110, 112, 0, 0, 0, 0, 114, 0, 163, 164,
	165, 166, 0, 160, 161, 162, 0, 0, 0, 0,
	293, 112, 152, 153, 154, 155, 156, 157, 158, 159,
	0, 0, 0, 0, 0, 148, 152, 153, 154, 155,
	156, 157, 158, 159,
}
var yyPact = [...]int{

	318, -1000, -1000, 221, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -60, -37, -22, 11, -24, -1000, -1000, -1000, 378,
	333, -1000, -1000, -1000, 330, -1000, 258, 284, 284, 369,
	256, -62, -30, 253, -1000, -25, 253, -1000, 299, -64,
	253, -64, 285, -1000, -1000, -1000, -1000, -1000, 424, -1000,
	241, 284, 235, 34, 264, 260, 284, 110, -1000, 156,
	-1000, 27, 283, 20, 253, -1000, -1000, 282, -1000, -51,
	281, 327, 59, 253, -1000, 166, -1000, -1000, 272, 14,
	104, 555, -1000, 495, 475, -1000, -1000, -1000, 515, 233,
	232, -1000, 217, 205, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 515, -1000, 189, 341, 256, 256,
	280, 256, 256, 367, 256, 515, 253, -1000, 326, -68,
	-1000, 132, -1000, 275, -1000, -1000, 273, -1000, 187, 424,
	-1000, -1000, 253, 94, 495, 495, 515, 230, 339, 515,
	515, 98, 515, 515, 515, 515, 515, 515, 515, 515,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 555, -47,
	-28, -3, 555, -1000, 444, 363, 424, -1000, 378, 261,
	-19, 510, 341, 256, 49, 191, 221, 126, -4, -1000,
	-1000, 207, 110, 346, 495, -1000, 510, -1000, -1000, -1000,
	51, 253, -1000, -56, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 206, 186, 269, 204, 2, -1000, -1000, -1000,
	-1000, -1000, 42, 510, -1000, 444, -1000, -1000, 230, 515,
	515, 510, 541, -1000, 315, 114, 114, 114, 32, 32,
	-1000, -1000, -1000, -1000, -1000, 515, -1000, 510, -1000, -6,
	424, -7, 130, 23, -1000, 495, 49, 126, -1000, 308,
	129, -1000, -1000, -1000, 256, 346, 337, 344, 104, 271,
	-1000, -1000, 267, -1000, 359, 187, 187, -1000, -1000, 146,
	143, 97, 91, 88, 5, -1000, 266, 19, 265, -8,
	-1000, 510, 351, 515, -1000, 510, -1000, -11, -1000, 261,
	7, -1000, 515, 4, -1000, -1000, 306, 191, -1000, 337,
	-1000, 515, 515, -1000, -1000, 354, 340, 186, 47, -1000,
	84, -1000, 82, -1000, -1000, -1000, -1000, -32, -33, -35,
	-1000, -1000, -1000, -1000, 515, 510, -1000, -73, -1000, 510,
	515, 374, -1000, -1000, 257, 127, -1000, 183, -1000, 346,
	495, 515, 495, -1000, -1000, 188, 179, 168, 510, -1000,
	510, 256, 515, 515, -1000, -1000, -1000, 337, 104, 111,
	104, 253, 253, 253, 110, 510, -1000, 323, -14, -1000,
	-16, -17, -1000, 373, 320, -1000, 253, -1000, -1000, -1000,
	253, -1000, 253, -1000,
}
var yyPgo = [...]int{

	0, 493, 492, 16, 490, 479, 478, 477, 476, 474,
	472, 469, 467, 466, 458, 468, 457, 456, 455, 19,
	23, 454, 452, 451, 448, 9, 447, 446, 175, 443,
	3, 14, 35, 442, 441, 18, 7, 2, 11, 5,
	440, 8, 436, 68, 435, 13, 434, 432, 12, 431,
	430, 421, 416, 6, 415, 4, 413, 1, 412, 20,
	411, 10, 28, 15, 229, 409, 398, 397, 396, 395,
	394, 0, 36, 385,
}
var yyR1 = [...]int{

	0, 1, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 3, 3, 4, 4, 4, 4,
	5, 6, 7, 8, 9, 9, 9, 10, 10, 10,
	11, 12, 12, 12, 13, 14, 14, 14, 73, 15,
	16, 16, 17, 17, 17, 17, 17, 18, 18, 19,
	19, 20, 20, 20, 23, 23, 21, 21, 21, 24,
	24, 25, 25, 25, 25, 22, 22, 22, 26, 26,
	26, 26, 26, 26, 26, 26, 26, 27, 27, 27,
	28, 28, 29, 29, 29, 29, 30, 30, 31, 31,
	32, 32, 32, 32, 32, 33, 33, 33, 33, 33,
	33, 33, 33, 33, 33, 33, 34, 34, 34, 34,
	34, 34, 34, 38, 38, 38, 43, 39, 39, 37,
	37, 37, 37, 37, 37, 37, 37, 37, 37, 37,
	37, 37, 37, 37, 37, 37, 42, 42, 44, 44,
	44, 46, 49, 49, 47, 47, 48, 50, 50, 45,
	45, 36, 36, 36, 36, 51, 51, 52, 52, 53,
	53, 54, 54, 55, 56, 56, 56, 57, 57, 57,
	58, 58, 58, 59, 59, 60, 60, 61, 61, 35,
	35, 40, 40, 41, 41, 62, 62, 63, 64, 64,
	65, 65, 66, 66, 67, 67, 67, 67, 67, 68,
	68, 69, 69, 70, 70, 71, 72,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 12, 3, 7, 6, 7, 6,
	8, 5, 7, 3, 5, 8, 4, 6, 7, 4,
	5, 4, 5, 5, 3, 2, 2, 2, 0, 2,
	0, 2, 1, 2, 1, 1, 1, 0, 1, 1,
	3, 1, 2, 3, 1, 1, 0, 1, 2, 1,
	3, 3, 3, 3, 5, 0, 1, 2, 1, 1,
	2, 3, 2, 3, 2, 2, 2, 1, 3, 1,
	1, 3, 0, 5, 5, 5, 1, 3, 0, 2,
	1, 3, 3, 2, 3, 3, 3, 4, 3, 4,
	5, 6, 3, 4, 2, 6, 1, 1, 1, 1,
	1, 1, 1, 3, 1, 1, 3, 1, 3, 1,
	1, 1, 3, 3, 3, 3, 3, 3, 3, 3,
	2, 3, 4, 5, 4, 1, 1, 1, 1, 1,
	1, 5, 0, 1, 1, 2, 4, 0, 2, 1,
	3, 1, 1, 1, 1, 0, 3, 0, 2, 0,
	3, 1, 3, 2, 0, 1, 1, 0, 2, 4,
	0, 2, 4, 0, 3, 1, 3, 0, 5, 2,
	1, 1, 3, 3, 1, 1, 3, 3, 0, 2,
	0, 3, 0, 1, 1, 1, 1, 1, 1, 0,
	1, 0, 1, 0, 2, 1, 0,
}
var yyChk = [...]int{

	-1000, -1, -2, -3, -4, -5, -6, -7, -8, -9,
	-10, -11, -12, -13, -14, 5, 6, 7, 9, 8,
	34, 86, 87, 89, 88, 90, 99, 100, 101, -17,
	52, 53, 54, 55, -15, -73, -15, -15, -15, -15,
	-15, 91, -69, 93, 97, -66, 93, 95, 91, 91,
	92, 93, 91, -72, -72, -72, -3, 18, -18, 19,
	-16, 30, -28, 37, -28, -28, 10, -62, -63, -45,
	-71, 37, -65, 96, 92, -71, 37, 91, -71, 37,
	-64, 96, -71, -64, 37, -19, -20, 76, -23, 37,
	-32, -37, -33, 70, 47, -36, -45, -41, -44, -71,
	-42, -46, 21, 36, 38, 39, 40, 26, -43, 74,
	75, 51, 96, 29, 81, 42, -28, -59, 34, 47,
	79, 34, 34, -28, 56, 48, 79, 37, 70, -71,
	-72, 37, -72, 94, 37, 21, 67, -71, 10, 56,
	-21, -71, 20, 79, 69, 68, -34, 22, 70, 24,
	25, 23, 71, 72, 73, 74, 75, 76, 77, 78,
	48, 49, 50, 43, 44, 45, 46, -32, -37, -32,
	-3, -39, -37, -37, 47, 47, 47, -43, 47, 47,
	-49, -37, -59, 34, -35, 29, -3, -62, -60, -45,
	37, -62, -62, -31, 11, -63, -37, -71, -72, 21,
	-70, 98, -67, 89, 87, 33, 88, 14, 37, 37,
	37, -72, -24, -25, -27, 47, 37, -43, -20, -71,
	76, -32, -32, -37, -38, 47, -43, 41, 22, 24,
	25, -37, -37, 26, 70, -37, -37, -37, -37, -37,
	-37, -37, -37, 102, 102, 56, 102, -37, 102, -19,
	19, -19, -36, -47, -48, 82, -35, -62, -61, 67,
	-40, -41, -61, 102, 56, -31, -53, 14, -32, 67,
	-71, -72, -68, 94, -31, 56, -26, 57, 58, 59,
	60, 61, 63, 64, -22, 37, 20, -25, 79, -39,
	-38, -37, -37, 69, 26, -37, 102, -19, 102, 56,
	-50, -48, 84, -32, -61, -61, 31, 56, -45, -53,
	-57, 16, 15, 37, 37, -51, 12, -25, -25, 57,
	62, 57, 62, 57, 57, 57, -29, 65, 95, 66,
	37, 102, 37, 102, 69, -37, 102, -36, 85, -37,
	83, 32, -41, -57, -37, -54, -55, -37, -72, -52,
	13, 15, 67, 57, 57, 92, 92, 92, -37, 102,
	-37, 7, 56, 56, -56, 27, 28, -53, -32, -39,
	-32, 47, 47, 47, -62, -37, -55, -57, -30, -71,
	-30, -30, -58, 17, 35, 102, 56, 102, 102, 7,
	22, -71, -71, -71,
}
var yyDef = [...]int{

	0, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 12, 13, 38, 38, 38, 38, 38,
	38, 201, 192, 0, 0, 0, 206, 206, 206, 0,
	42, 44, 45, 46, 47, 40, 0, 0, 0, 0,
	0, 190, 0, 0, 202, 0, 0, 193, 0, 188,
	0, 188, 0, 35, 36, 37, 15, 43, 0, 48,
	39, 0, 173, 80, 0, 0, 0, 23, 185, 0,
	149, 205, 0, 0, 0, 206, 205, 0, 206, 0,
	0, 0, 0, 0, 34, 0, 49, 51, 56, 205,
	54, 55, 90, 0, 0, 119, 120, 121, 0, 149,
	0, 135, 0, 0, 151, 152, 153, 154, 184, 138,
	139, 140, 136, 137, 142, 41, 173, 0, 0, 0,
	0, 0, 0, 88, 0, 0, 0, 206, 0, 203,
	26, 0, 29, 0, 31, 189, 0, 206, 0, 0,
	52, 57, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	106, 107, 108, 109, 110, 111, 112, 93, 0, 0,
	0, 0, 117, 130, 0, 0, 0, 104, 0, 0,
	0, 143, 0, 0, 177, 0, 180, 177, 0, 175,
	81, 88, 21, 159, 0, 186, 187, 150, 24, 191,
	0, 0, 206, 199, 194, 195, 196, 197, 198, 30,
	32, 33, 88, 59, 65, 0, 77, 79, 50, 58,
	53, 91, 92, 95, 96, 0, 114, 115, 0, 0,
	0, 98, 0, 102, 0, 122, 123, 124, 125, 126,
	127, 128, 129, 94, 116, 0, 183, 117, 131, 0,
	0, 0, 0, 147, 144, 0, 177, 177, 17, 0,
	179, 181, 19, 174, 0, 159, 167, 0, 89, 0,
	204, 27, 0, 200, 155, 0, 0, 68, 69, 0,
	0, 0, 0, 0, 82, 66, 0, 0, 0, 0,
	97, 99, 0, 0, 103, 118, 132, 0, 134, 0,
	0, 145, 0, 0, 16, 18, 0, 0, 176, 167,
	22, 0, 0, 206, 28, 157, 0, 60, 63, 70,
	0, 72, 0, 74, 75, 76, 61, 0, 0, 0,
	67, 62, 78, 113, 0, 100, 133, 0, 141, 148,
	0, 0, 182, 20, 168, 160, 161, 164, 25, 159,
	0, 0, 0, 71, 73, 0, 0, 0, 101, 105,
	146, 0, 0, 0, 163, 165, 166, 167, 158, 156,
	64, 0, 0, 0, 178, 169, 162, 170, 0, 86,
	0, 0, 14, 0, 0, 83, 0, 84, 85, 171,
	0, 87, 0, 172,
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
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:188
		{
			yyVAL.statement = &Insert{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[3].tableName, Columns: yyDollar[4].columns, Rows: yyDollar[5].insRows, OnDup: OnDup(yyDollar[6].updateExprs)}
		}
	case 18:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:192
		{
			cols := make(Columns, 0, len(yyDollar[6].updateExprs))
			vals := make(ValTuple, 0, len(yyDollar[6].updateExprs))
			for _, col := range yyDollar[6].updateExprs {
				cols = append(cols, &NonStarExpr{Expr: col.Name})
				vals = append(vals, col.Expr)
			}
			yyVAL.statement = &Insert{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Columns: cols, Rows: Values{vals}, OnDup: OnDup(yyDollar[7].updateExprs)}
		}
	case 19:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:202
		{
			cols := make(Columns, 0, len(yyDollar[5].updateExprs))
			vals := make(ValTuple, 0, len(yyDollar[5].updateExprs))
			for _, col := range yyDollar[5].updateExprs {
				cols = append(cols, &NonStarExpr{Expr: col.Name})
				vals = append(vals, col.Expr)
			}
			yyVAL.statement = &Insert{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[3].tableName, Columns: cols, Rows: Values{vals}, OnDup: OnDup(yyDollar[6].updateExprs)}
		}
	case 20:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line sql.y:214
		{
			yyVAL.statement = &Update{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[3].tableName, Exprs: yyDollar[5].updateExprs, Where: NewWhere(AST_WHERE, yyDollar[6].boolExpr), OrderBy: yyDollar[7].orderBy, Limit: yyDollar[8].limit}
		}
	case 21:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:220
		{
			yyVAL.statement = &Replace{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[3].tableName, Exprs: yyDollar[5].updateExprs}
		}
	case 22:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:226
		{
			yyVAL.statement = &Delete{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Where: NewWhere(AST_WHERE, yyDollar[5].boolExpr), OrderBy: yyDollar[6].orderBy, Limit: yyDollar[7].limit}
		}
	case 23:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:232
		{
			yyVAL.statement = &Set{Comments: Comments(yyDollar[2].bytes2), Exprs: yyDollar[3].updateExprs}
		}
	case 24:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:238
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyDollar[4].bytes}
		}
	case 25:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line sql.y:242
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[7].bytes, NewName: yyDollar[7].bytes}
		}
	case 26:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:247
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyDollar[3].bytes}
		}
	case 27:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:253
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[4].bytes, NewName: yyDollar[4].bytes}
		}
	case 28:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:257
		{
			// Change this to a rename statement
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyDollar[4].bytes, NewName: yyDollar[7].bytes}
		}
	case 29:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:262
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[3].bytes, NewName: yyDollar[3].bytes}
		}
	case 30:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:268
		{
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyDollar[3].bytes, NewName: yyDollar[5].bytes}
		}
	case 31:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:274
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyDollar[4].bytes}
		}
	case 32:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:278
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[5].bytes, NewName: yyDollar[5].bytes}
		}
	case 33:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:283
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyDollar[4].bytes}
		}
	case 34:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:289
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[3].bytes, NewName: yyDollar[3].bytes}
		}
	case 35:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:295
		{
			yyVAL.statement = &Other{}
		}
	case 36:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:299
		{
			yyVAL.statement = &Other{}
		}
	case 37:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:303
		{
			yyVAL.statement = &Other{}
		}
	case 38:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:308
		{
			SetAllowComments(yylex, true)
		}
	case 39:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:312
		{
			yyVAL.bytes2 = yyDollar[2].bytes2
			SetAllowComments(yylex, false)
		}
	case 40:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:318
		{
			yyVAL.bytes2 = nil
		}
	case 41:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:322
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[2].bytes)
		}
	case 42:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:328
		{
			yyVAL.str = AST_UNION
		}
	case 43:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:332
		{
			yyVAL.str = AST_UNION_ALL
		}
	case 44:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:336
		{
			yyVAL.str = AST_SET_MINUS
		}
	case 45:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:340
		{
			yyVAL.str = AST_EXCEPT
		}
	case 46:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:344
		{
			yyVAL.str = AST_INTERSECT
		}
	case 47:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:349
		{
			yyVAL.str = ""
		}
	case 48:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:353
		{
			yyVAL.str = AST_DISTINCT
		}
	case 49:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:359
		{
			yyVAL.selectExprs = SelectExprs{yyDollar[1].selectExpr}
		}
	case 50:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:363
		{
			yyVAL.selectExprs = append(yyVAL.selectExprs, yyDollar[3].selectExpr)
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:369
		{
			yyVAL.selectExpr = &StarExpr{}
		}
	case 52:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:373
		{
			yyVAL.selectExpr = &NonStarExpr{Expr: yyDollar[1].expr, As: yyDollar[2].bytes}
		}
	case 53:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:377
		{
			yyVAL.selectExpr = &StarExpr{TableName: yyDollar[1].bytes}
		}
	case 54:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:383
		{
			yyVAL.expr = yyDollar[1].boolExpr
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:387
		{
			yyVAL.expr = yyDollar[1].valExpr
		}
	case 56:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:392
		{
			yyVAL.bytes = nil
		}
	case 57:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:396
		{
			yyVAL.bytes = yyDollar[1].bytes
		}
	case 58:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:400
		{
			yyVAL.bytes = yyDollar[2].bytes
		}
	case 59:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:406
		{
			yyVAL.tableExprs = TableExprs{yyDollar[1].tableExpr}
		}
	case 60:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:410
		{
			yyVAL.tableExprs = append(yyVAL.tableExprs, yyDollar[3].tableExpr)
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:416
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyDollar[1].smTableExpr, As: yyDollar[2].bytes, Hints: yyDollar[3].indexHints}
		}
	case 62:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:420
		{
			yyVAL.tableExpr = &ParenTableExpr{Expr: yyDollar[2].tableExpr}
		}
	case 63:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:424
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr}
		}
	case 64:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:428
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr, On: yyDollar[5].boolExpr}
		}
	case 65:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:433
		{
			yyVAL.bytes = nil
		}
	case 66:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:437
		{
			yyVAL.bytes = yyDollar[1].bytes
		}
	case 67:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:441
		{
			yyVAL.bytes = yyDollar[2].bytes
		}
	case 68:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:447
		{
			yyVAL.str = AST_JOIN
		}
	case 69:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:451
		{
			yyVAL.str = AST_STRAIGHT_JOIN
		}
	case 70:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:455
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 71:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:459
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 72:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:463
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 73:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:467
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 74:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:471
		{
			yyVAL.str = AST_JOIN
		}
	case 75:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:475
		{
			yyVAL.str = AST_CROSS_JOIN
		}
	case 76:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:479
		{
			yyVAL.str = AST_NATURAL_JOIN
		}
	case 77:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:485
		{
			yyVAL.smTableExpr = &TableName{Name: yyDollar[1].bytes}
		}
	case 78:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:489
		{
			yyVAL.smTableExpr = &TableName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 79:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:493
		{
			yyVAL.smTableExpr = yyDollar[1].subquery
		}
	case 80:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:499
		{
			yyVAL.tableName = &TableName{Name: yyDollar[1].bytes}
		}
	case 81:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:503
		{
			yyVAL.tableName = &TableName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 82:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:508
		{
			yyVAL.indexHints = nil
		}
	case 83:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:512
		{
			yyVAL.indexHints = &IndexHints{Type: AST_USE, Indexes: yyDollar[4].bytes2}
		}
	case 84:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:516
		{
			yyVAL.indexHints = &IndexHints{Type: AST_IGNORE, Indexes: yyDollar[4].bytes2}
		}
	case 85:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:520
		{
			yyVAL.indexHints = &IndexHints{Type: AST_FORCE, Indexes: yyDollar[4].bytes2}
		}
	case 86:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:526
		{
			yyVAL.bytes2 = [][]byte{yyDollar[1].bytes}
		}
	case 87:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:530
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[3].bytes)
		}
	case 88:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:535
		{
			yyVAL.boolExpr = nil
		}
	case 89:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:539
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 91:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:546
		{
			yyVAL.boolExpr = &AndExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 92:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:550
		{
			yyVAL.boolExpr = &OrExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 93:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:554
		{
			yyVAL.boolExpr = &NotExpr{Expr: yyDollar[2].boolExpr}
		}
	case 94:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:558
		{
			yyVAL.boolExpr = &ParenBoolExpr{Expr: yyDollar[2].boolExpr}
		}
	case 95:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:564
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: yyDollar[2].str, Right: yyDollar[3].valExpr}
		}
	case 96:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:568
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_IN, Right: yyDollar[3].colTuple}
		}
	case 97:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:572
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_NOT_IN, Right: yyDollar[4].colTuple}
		}
	case 98:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:576
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_LIKE, Right: yyDollar[3].valExpr}
		}
	case 99:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:580
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_NOT_LIKE, Right: yyDollar[4].valExpr}
		}
	case 100:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:584
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: AST_BETWEEN, From: yyDollar[3].valExpr, To: yyDollar[5].valExpr}
		}
	case 101:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:588
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: AST_NOT_BETWEEN, From: yyDollar[4].valExpr, To: yyDollar[6].valExpr}
		}
	case 102:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:592
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NULL, Expr: yyDollar[1].valExpr}
		}
	case 103:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:596
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NOT_NULL, Expr: yyDollar[1].valExpr}
		}
	case 104:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:600
		{
			yyVAL.boolExpr = &ExistsExpr{Subquery: yyDollar[2].subquery}
		}
	case 105:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:604
		{
			yyVAL.boolExpr = &KeyrangeExpr{Start: yyDollar[3].valExpr, End: yyDollar[5].valExpr}
		}
	case 106:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:610
		{
			yyVAL.str = AST_EQ
		}
	case 107:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:614
		{
			yyVAL.str = AST_LT
		}
	case 108:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:618
		{
			yyVAL.str = AST_GT
		}
	case 109:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:622
		{
			yyVAL.str = AST_LE
		}
	case 110:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:626
		{
			yyVAL.str = AST_GE
		}
	case 111:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:630
		{
			yyVAL.str = AST_NE
		}
	case 112:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:634
		{
			yyVAL.str = AST_NSE
		}
	case 113:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:640
		{
			yyVAL.colTuple = ValTuple(yyDollar[2].valExprs)
		}
	case 114:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:644
		{
			yyVAL.colTuple = yyDollar[1].subquery
		}
	case 115:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:648
		{
			yyVAL.colTuple = ListArg(yyDollar[1].bytes)
		}
	case 116:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:654
		{
			yyVAL.subquery = &Subquery{yyDollar[2].selStmt}
		}
	case 117:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:660
		{
			yyVAL.valExprs = ValExprs{yyDollar[1].valExpr}
		}
	case 118:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:664
		{
			yyVAL.valExprs = append(yyDollar[1].valExprs, yyDollar[3].valExpr)
		}
	case 119:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:670
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 120:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:674
		{
			yyVAL.valExpr = yyDollar[1].colName
		}
	case 121:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:678
		{
			yyVAL.valExpr = yyDollar[1].rowTuple
		}
	case 122:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:682
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITAND, Right: yyDollar[3].valExpr}
		}
	case 123:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:686
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITOR, Right: yyDollar[3].valExpr}
		}
	case 124:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:690
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITXOR, Right: yyDollar[3].valExpr}
		}
	case 125:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:694
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_PLUS, Right: yyDollar[3].valExpr}
		}
	case 126:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:698
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MINUS, Right: yyDollar[3].valExpr}
		}
	case 127:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:702
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MULT, Right: yyDollar[3].valExpr}
		}
	case 128:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:706
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_DIV, Right: yyDollar[3].valExpr}
		}
	case 129:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:710
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MOD, Right: yyDollar[3].valExpr}
		}
	case 130:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:714
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
	case 131:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:729
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes}
		}
	case 132:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:733
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Exprs: yyDollar[3].selectExprs}
		}
	case 133:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:737
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Distinct: true, Exprs: yyDollar[4].selectExprs}
		}
	case 134:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:741
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Exprs: yyDollar[3].selectExprs}
		}
	case 135:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:745
		{
			yyVAL.valExpr = yyDollar[1].caseExpr
		}
	case 136:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:751
		{
			yyVAL.bytes = IF_BYTES
		}
	case 137:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:755
		{
			yyVAL.bytes = VALUES_BYTES
		}
	case 138:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:761
		{
			yyVAL.byt = AST_UPLUS
		}
	case 139:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:765
		{
			yyVAL.byt = AST_UMINUS
		}
	case 140:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:769
		{
			yyVAL.byt = AST_TILDA
		}
	case 141:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:775
		{
			yyVAL.caseExpr = &CaseExpr{Expr: yyDollar[2].valExpr, Whens: yyDollar[3].whens, Else: yyDollar[4].valExpr}
		}
	case 142:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:780
		{
			yyVAL.valExpr = nil
		}
	case 143:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:784
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 144:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:790
		{
			yyVAL.whens = []*When{yyDollar[1].when}
		}
	case 145:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:794
		{
			yyVAL.whens = append(yyDollar[1].whens, yyDollar[2].when)
		}
	case 146:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:800
		{
			yyVAL.when = &When{Cond: yyDollar[2].boolExpr, Val: yyDollar[4].valExpr}
		}
	case 147:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:805
		{
			yyVAL.valExpr = nil
		}
	case 148:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:809
		{
			yyVAL.valExpr = yyDollar[2].valExpr
		}
	case 149:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:815
		{
			yyVAL.colName = &ColName{Name: yyDollar[1].bytes}
		}
	case 150:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:819
		{
			yyVAL.colName = &ColName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 151:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:825
		{
			yyVAL.valExpr = StrVal(yyDollar[1].bytes)
		}
	case 152:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:829
		{
			yyVAL.valExpr = NumVal(yyDollar[1].bytes)
		}
	case 153:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:833
		{
			yyVAL.valExpr = ValArg(yyDollar[1].bytes)
		}
	case 154:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:837
		{
			yyVAL.valExpr = &NullVal{}
		}
	case 155:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:842
		{
			yyVAL.valExprs = nil
		}
	case 156:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:846
		{
			yyVAL.valExprs = yyDollar[3].valExprs
		}
	case 157:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:851
		{
			yyVAL.boolExpr = nil
		}
	case 158:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:855
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 159:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:860
		{
			yyVAL.orderBy = nil
		}
	case 160:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:864
		{
			yyVAL.orderBy = yyDollar[3].orderBy
		}
	case 161:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:870
		{
			yyVAL.orderBy = OrderBy{yyDollar[1].order}
		}
	case 162:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:874
		{
			yyVAL.orderBy = append(yyDollar[1].orderBy, yyDollar[3].order)
		}
	case 163:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:880
		{
			yyVAL.order = &Order{Expr: yyDollar[1].valExpr, Direction: yyDollar[2].str}
		}
	case 164:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:885
		{
			yyVAL.str = AST_ASC
		}
	case 165:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:889
		{
			yyVAL.str = AST_ASC
		}
	case 166:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:893
		{
			yyVAL.str = AST_DESC
		}
	case 167:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:898
		{
			yyVAL.limit = nil
		}
	case 168:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:902
		{
			yyVAL.limit = &Limit{Rowcount: yyDollar[2].valExpr}
		}
	case 169:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:906
		{
			yyVAL.limit = &Limit{Offset: yyDollar[2].valExpr, Rowcount: yyDollar[4].valExpr}
		}
	case 170:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:911
		{
			yyVAL.str = ""
		}
	case 171:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:915
		{
			yyVAL.str = AST_FOR_UPDATE
		}
	case 172:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:919
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
	case 173:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:932
		{
			yyVAL.columns = nil
		}
	case 174:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:936
		{
			yyVAL.columns = yyDollar[2].columns
		}
	case 175:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:942
		{
			yyVAL.columns = Columns{&NonStarExpr{Expr: yyDollar[1].colName}}
		}
	case 176:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:946
		{
			yyVAL.columns = append(yyVAL.columns, &NonStarExpr{Expr: yyDollar[3].colName})
		}
	case 177:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:951
		{
			yyVAL.updateExprs = nil
		}
	case 178:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:955
		{
			yyVAL.updateExprs = yyDollar[5].updateExprs
		}
	case 179:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:961
		{
			yyVAL.insRows = yyDollar[2].values
		}
	case 180:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:965
		{
			yyVAL.insRows = yyDollar[1].selStmt
		}
	case 181:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:971
		{
			yyVAL.values = Values{yyDollar[1].rowTuple}
		}
	case 182:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:975
		{
			yyVAL.values = append(yyDollar[1].values, yyDollar[3].rowTuple)
		}
	case 183:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:981
		{
			yyVAL.rowTuple = ValTuple(yyDollar[2].valExprs)
		}
	case 184:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:985
		{
			yyVAL.rowTuple = yyDollar[1].subquery
		}
	case 185:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:991
		{
			yyVAL.updateExprs = UpdateExprs{yyDollar[1].updateExpr}
		}
	case 186:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:995
		{
			yyVAL.updateExprs = append(yyDollar[1].updateExprs, yyDollar[3].updateExpr)
		}
	case 187:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1001
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyDollar[1].colName, Expr: yyDollar[3].valExpr}
		}
	case 188:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1006
		{
			yyVAL.empty = struct{}{}
		}
	case 189:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1008
		{
			yyVAL.empty = struct{}{}
		}
	case 190:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1011
		{
			yyVAL.empty = struct{}{}
		}
	case 191:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1013
		{
			yyVAL.empty = struct{}{}
		}
	case 192:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1016
		{
			yyVAL.empty = struct{}{}
		}
	case 193:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1018
		{
			yyVAL.empty = struct{}{}
		}
	case 194:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1022
		{
			yyVAL.empty = struct{}{}
		}
	case 195:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1024
		{
			yyVAL.empty = struct{}{}
		}
	case 196:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1026
		{
			yyVAL.empty = struct{}{}
		}
	case 197:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1028
		{
			yyVAL.empty = struct{}{}
		}
	case 198:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1030
		{
			yyVAL.empty = struct{}{}
		}
	case 199:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1033
		{
			yyVAL.empty = struct{}{}
		}
	case 200:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1035
		{
			yyVAL.empty = struct{}{}
		}
	case 201:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1038
		{
			yyVAL.empty = struct{}{}
		}
	case 202:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1040
		{
			yyVAL.empty = struct{}{}
		}
	case 203:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1043
		{
			yyVAL.empty = struct{}{}
		}
	case 204:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1045
		{
			yyVAL.empty = struct{}{}
		}
	case 205:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1049
		{
			yyVAL.bytes = bytes.ToLower(yyDollar[1].bytes)
		}
	case 206:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1054
		{
			ForceEOF(yylex)
		}
	}
	goto yystack /* stack new state and value */
}
