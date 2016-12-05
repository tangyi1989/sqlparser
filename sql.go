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

const yyNprod = 206
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 591

var yyAct = [...]int{

	99, 305, 166, 375, 342, 90, 97, 259, 169, 96,
	207, 95, 248, 67, 256, 218, 68, 187, 383, 168,
	3, 355, 383, 86, 85, 383, 53, 270, 271, 272,
	273, 274, 195, 275, 276, 30, 31, 32, 33, 143,
	142, 70, 81, 73, 75, 137, 15, 78, 239, 56,
	69, 82, 302, 266, 54, 55, 353, 137, 137, 131,
	108, 91, 102, 239, 385, 322, 324, 107, 384, 352,
	113, 382, 326, 237, 46, 127, 47, 103, 71, 104,
	105, 106, 351, 74, 135, 238, 77, 126, 94, 139,
	52, 331, 111, 333, 328, 323, 48, 170, 301, 165,
	167, 171, 128, 291, 289, 130, 41, 249, 43, 240,
	281, 93, 44, 143, 142, 109, 110, 179, 70, 141,
	70, 70, 114, 70, 142, 191, 190, 69, 335, 69,
	69, 183, 69, 185, 186, 124, 249, 112, 295, 189,
	91, 213, 191, 49, 50, 51, 118, 217, 215, 216,
	225, 226, 192, 229, 230, 231, 232, 233, 234, 235,
	236, 212, 205, 175, 329, 227, 150, 151, 152, 153,
	154, 155, 156, 157, 76, 241, 91, 91, 155, 156,
	157, 348, 70, 70, 153, 154, 155, 156, 157, 246,
	257, 69, 255, 262, 261, 253, 263, 211, 243, 245,
	252, 143, 142, 258, 107, 134, 220, 113, 122, 228,
	62, 350, 349, 214, 320, 71, 104, 105, 106, 257,
	280, 201, 241, 264, 267, 172, 284, 285, 282, 111,
	319, 318, 15, 16, 17, 19, 18, 188, 283, 122,
	199, 316, 288, 239, 202, 314, 317, 91, 64, 65,
	315, 136, 109, 110, 188, 296, 360, 336, 299, 114,
	294, 20, 292, 181, 123, 297, 304, 370, 300, 290,
	211, 369, 116, 76, 112, 368, 182, 121, 172, 312,
	313, 80, 268, 220, 30, 31, 32, 33, 71, 330,
	210, 177, 176, 174, 198, 200, 197, 137, 334, 122,
	209, 221, 15, 70, 332, 173, 339, 219, 115, 340,
	343, 327, 337, 21, 22, 24, 23, 25, 150, 151,
	152, 153, 154, 155, 156, 157, 26, 27, 28, 211,
	211, 279, 354, 83, 210, 344, 325, 309, 356, 61,
	140, 308, 204, 357, 209, 203, 63, 184, 278, 132,
	241, 129, 365, 364, 367, 125, 366, 76, 63, 70,
	84, 79, 372, 343, 120, 373, 374, 119, 69, 376,
	376, 376, 371, 377, 378, 244, 117, 102, 338, 303,
	15, 287, 107, 380, 388, 113, 107, 387, 389, 193,
	390, 133, 103, 89, 104, 105, 106, 359, 104, 105,
	106, 381, 59, 94, 251, 57, 222, 111, 223, 224,
	306, 347, 150, 151, 152, 153, 154, 155, 156, 157,
	307, 270, 271, 272, 273, 274, 93, 275, 276, 260,
	109, 110, 87, 346, 311, 188, 66, 114, 102, 386,
	358, 15, 15, 107, 35, 194, 113, 42, 265, 196,
	45, 72, 112, 103, 89, 104, 105, 106, 242, 254,
	180, 34, 379, 107, 94, 361, 113, 341, 111, 345,
	310, 293, 178, 247, 71, 104, 105, 106, 36, 37,
	38, 39, 40, 101, 172, 98, 100, 93, 111, 102,
	298, 109, 110, 87, 107, 250, 144, 113, 114, 92,
	321, 208, 269, 206, 103, 71, 104, 105, 106, 362,
	363, 109, 110, 112, 88, 94, 277, 138, 114, 111,
	58, 29, 60, 14, 145, 149, 147, 148, 13, 12,
	11, 10, 9, 112, 8, 7, 6, 5, 93, 4,
	2, 1, 109, 110, 0, 161, 162, 163, 164, 114,
	158, 159, 160, 150, 151, 152, 153, 154, 155, 156,
	157, 0, 0, 0, 112, 0, 0, 0, 0, 0,
	0, 0, 146, 150, 151, 152, 153, 154, 155, 156,
	157, 286, 0, 150, 151, 152, 153, 154, 155, 156,
	157,
}
var yyPact = [...]int{

	227, -1000, -1000, 232, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 15, -19, 5, 52, -1, -1000, -1000, -1000, 436,
	387, -1000, -1000, -1000, 383, -1000, 309, 321, 321, 426,
	251, -53, -9, 236, -1000, -5, 236, -1000, 324, -54,
	236, -54, 323, -1000, -1000, -1000, -1000, -1000, 417, -1000,
	266, 321, 342, 67, 333, 330, 321, 183, -1000, 216,
	-1000, 56, 318, 17, 236, -1000, -1000, 314, -1000, -35,
	312, 370, 138, 236, -1000, 241, -1000, -1000, 320, 40,
	133, 502, -1000, 468, 41, -1000, -1000, -1000, 178, 258,
	246, -1000, 245, 244, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 178, -1000, 229, 251, 310, 251,
	251, 424, 251, 178, 236, -1000, 368, -66, -1000, 207,
	-1000, 308, -1000, -1000, 305, -1000, 253, 417, -1000, -1000,
	236, 137, 468, 468, 178, 260, 384, 178, 178, 139,
	178, 178, 178, 178, 178, 178, 178, 178, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 502, -29, -17, 7,
	502, -1000, 437, 356, 417, -1000, 436, 360, 25, 247,
	375, 251, 251, 152, -1000, 243, 183, 415, 468, -1000,
	247, -1000, -1000, -1000, 126, 236, -1000, -41, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 226, 364, 311, 297,
	31, -1000, -1000, -1000, -1000, -1000, 55, 247, -1000, 437,
	-1000, -1000, 260, 178, 178, 247, 512, -1000, 355, 110,
	110, 110, 102, 102, -1000, -1000, -1000, -1000, -1000, 178,
	-1000, 247, -1000, 2, 417, 1, 206, 54, -1000, 468,
	123, 231, 232, 152, -4, -1000, -1000, 348, 415, 394,
	405, 133, 304, -1000, -1000, 300, -1000, 422, 253, 253,
	-1000, -1000, 188, 184, 174, 173, 157, 0, -1000, 299,
	-30, 274, -8, -1000, 247, 95, 178, -1000, 247, -1000,
	-11, -1000, 360, 8, -1000, 178, 45, -1000, 201, -1000,
	-1000, -1000, 251, 346, 394, -1000, 178, 178, -1000, -1000,
	420, 396, 364, 114, -1000, 155, -1000, 154, -1000, -1000,
	-1000, -1000, -10, -23, -36, -1000, -1000, -1000, -1000, 178,
	247, -1000, -81, -1000, 247, 178, 231, -1000, 433, -1000,
	341, 200, -1000, 482, -1000, 415, 468, 178, 468, -1000,
	-1000, 228, 224, 220, 247, -1000, 247, -1000, 251, 178,
	178, -1000, -1000, -1000, 394, 133, 187, 133, 236, 236,
	236, 183, 247, -1000, 366, -31, -1000, -34, -38, -1000,
	432, 365, -1000, 236, -1000, -1000, -1000, 236, -1000, 236,
	-1000,
}
var yyPgo = [...]int{

	0, 541, 540, 19, 539, 537, 536, 535, 534, 532,
	531, 530, 529, 528, 523, 461, 522, 521, 520, 24,
	23, 517, 516, 514, 503, 10, 502, 501, 210, 500,
	3, 17, 5, 499, 496, 495, 11, 2, 15, 8,
	490, 6, 486, 60, 485, 9, 483, 473, 12, 472,
	471, 470, 469, 7, 467, 4, 465, 1, 462, 460,
	459, 14, 13, 16, 281, 451, 450, 449, 448, 447,
	445, 0, 26, 444,
}
var yyR1 = [...]int{

	0, 1, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 3, 3, 4, 4, 4, 5,
	6, 7, 8, 9, 9, 9, 10, 10, 10, 11,
	12, 12, 12, 13, 14, 14, 14, 73, 15, 16,
	16, 17, 17, 17, 17, 17, 18, 18, 19, 19,
	20, 20, 20, 23, 23, 21, 21, 21, 24, 24,
	25, 25, 25, 25, 22, 22, 22, 26, 26, 26,
	26, 26, 26, 26, 26, 26, 27, 27, 27, 28,
	28, 29, 29, 29, 29, 30, 30, 31, 31, 32,
	32, 32, 32, 32, 33, 33, 33, 33, 33, 33,
	33, 33, 33, 33, 33, 34, 34, 34, 34, 34,
	34, 34, 38, 38, 38, 43, 39, 39, 37, 37,
	37, 37, 37, 37, 37, 37, 37, 37, 37, 37,
	37, 37, 37, 37, 37, 42, 42, 44, 44, 44,
	46, 49, 49, 47, 47, 48, 50, 50, 45, 45,
	36, 36, 36, 36, 51, 51, 52, 52, 53, 53,
	54, 54, 55, 56, 56, 56, 57, 57, 57, 58,
	58, 58, 59, 59, 60, 60, 61, 61, 35, 35,
	40, 40, 41, 41, 62, 62, 63, 64, 64, 65,
	65, 66, 66, 67, 67, 67, 67, 67, 68, 68,
	69, 69, 70, 70, 71, 72,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 12, 3, 7, 7, 6, 8,
	5, 7, 3, 5, 8, 4, 6, 7, 4, 5,
	4, 5, 5, 3, 2, 2, 2, 0, 2, 0,
	2, 1, 2, 1, 1, 1, 0, 1, 1, 3,
	1, 2, 3, 1, 1, 0, 1, 2, 1, 3,
	3, 3, 3, 5, 0, 1, 2, 1, 1, 2,
	3, 2, 3, 2, 2, 2, 1, 3, 1, 1,
	3, 0, 5, 5, 5, 1, 3, 0, 2, 1,
	3, 3, 2, 3, 3, 3, 4, 3, 4, 5,
	6, 3, 4, 2, 6, 1, 1, 1, 1, 1,
	1, 1, 3, 1, 1, 3, 1, 3, 1, 1,
	1, 3, 3, 3, 3, 3, 3, 3, 3, 2,
	3, 4, 5, 4, 1, 1, 1, 1, 1, 1,
	5, 0, 1, 1, 2, 4, 0, 2, 1, 3,
	1, 1, 1, 1, 0, 3, 0, 2, 0, 3,
	1, 3, 2, 0, 1, 1, 0, 2, 4, 0,
	2, 4, 0, 3, 1, 3, 0, 5, 2, 1,
	1, 3, 3, 1, 1, 3, 3, 0, 2, 0,
	3, 0, 1, 1, 1, 1, 1, 1, 0, 1,
	0, 1, 0, 2, 1, 0,
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
	75, 51, 96, 29, 81, 42, -28, 34, 79, 34,
	34, -28, 56, 48, 79, 37, 70, -71, -72, 37,
	-72, 94, 37, 21, 67, -71, 10, 56, -21, -71,
	20, 79, 69, 68, -34, 22, 70, 24, 25, 23,
	71, 72, 73, 74, 75, 76, 77, 78, 48, 49,
	50, 43, 44, 45, 46, -32, -37, -32, -3, -39,
	-37, -37, 47, 47, 47, -43, 47, 47, -49, -37,
	-59, 34, 47, -62, 37, -62, -62, -31, 11, -63,
	-37, -71, -72, 21, -70, 98, -67, 89, 87, 33,
	88, 14, 37, 37, 37, -72, -24, -25, -27, 47,
	37, -43, -20, -71, 76, -32, -32, -37, -38, 47,
	-43, 41, 22, 24, 25, -37, -37, 26, 70, -37,
	-37, -37, -37, -37, -37, -37, -37, 102, 102, 56,
	102, -37, 102, -19, 19, -19, -36, -47, -48, 82,
	-35, 29, -3, -62, -60, -45, -61, 67, -31, -53,
	14, -32, 67, -71, -72, -68, 94, -31, 56, -26,
	57, 58, 59, 60, 61, 63, 64, -22, 37, 20,
	-25, 79, -39, -38, -37, -37, 69, 26, -37, 102,
	-19, 102, 56, -50, -48, 84, -32, -61, -40, -41,
	-61, 102, 56, 31, -53, -57, 16, 15, 37, 37,
	-51, 12, -25, -25, 57, 62, 57, 62, 57, 57,
	57, -29, 65, 95, 66, 37, 102, 37, 102, 69,
	-37, 102, -36, 85, -37, 83, 56, -45, 32, -57,
	-37, -54, -55, -37, -72, -52, 13, 15, 67, 57,
	57, 92, 92, 92, -37, 102, -37, -41, 7, 56,
	56, -56, 27, 28, -53, -32, -39, -32, 47, 47,
	47, -62, -37, -55, -57, -30, -71, -30, -30, -58,
	17, 35, 102, 56, 102, 102, 7, 22, -71, -71,
	-71,
}
var yyDef = [...]int{

	0, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 12, 13, 37, 37, 37, 37, 37,
	37, 200, 191, 0, 0, 0, 205, 205, 205, 0,
	41, 43, 44, 45, 46, 39, 0, 0, 0, 0,
	0, 189, 0, 0, 201, 0, 0, 192, 0, 187,
	0, 187, 0, 34, 35, 36, 15, 42, 0, 47,
	38, 0, 0, 79, 0, 0, 0, 22, 184, 0,
	148, 204, 0, 0, 0, 205, 204, 0, 205, 0,
	0, 0, 0, 0, 33, 0, 48, 50, 55, 204,
	53, 54, 89, 0, 0, 118, 119, 120, 0, 148,
	0, 134, 0, 0, 150, 151, 152, 153, 183, 137,
	138, 139, 135, 136, 141, 40, 172, 0, 0, 0,
	0, 87, 0, 0, 0, 205, 0, 202, 25, 0,
	28, 0, 30, 188, 0, 205, 0, 0, 51, 56,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 105, 106,
	107, 108, 109, 110, 111, 92, 0, 0, 0, 0,
	116, 129, 0, 0, 0, 103, 0, 0, 0, 142,
	0, 0, 0, 176, 80, 87, 20, 158, 0, 185,
	186, 149, 23, 190, 0, 0, 205, 198, 193, 194,
	195, 196, 197, 29, 31, 32, 87, 58, 64, 0,
	76, 78, 49, 57, 52, 90, 91, 94, 95, 0,
	113, 114, 0, 0, 0, 97, 0, 101, 0, 121,
	122, 123, 124, 125, 126, 127, 128, 93, 115, 0,
	182, 116, 130, 0, 0, 0, 0, 146, 143, 0,
	176, 0, 179, 176, 0, 174, 18, 0, 158, 166,
	0, 88, 0, 203, 26, 0, 199, 154, 0, 0,
	67, 68, 0, 0, 0, 0, 0, 81, 65, 0,
	0, 0, 0, 96, 98, 0, 0, 102, 117, 131,
	0, 133, 0, 0, 144, 0, 0, 16, 178, 180,
	17, 173, 0, 0, 166, 21, 0, 0, 205, 27,
	156, 0, 59, 62, 69, 0, 71, 0, 73, 74,
	75, 60, 0, 0, 0, 66, 61, 77, 112, 0,
	99, 132, 0, 140, 147, 0, 0, 175, 0, 19,
	167, 159, 160, 163, 24, 158, 0, 0, 0, 70,
	72, 0, 0, 0, 100, 104, 145, 181, 0, 0,
	0, 162, 164, 165, 166, 157, 155, 63, 0, 0,
	0, 177, 168, 161, 169, 0, 85, 0, 0, 14,
	0, 0, 82, 0, 83, 84, 170, 0, 86, 0,
	171,
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
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:198
		{
			cols := make(Columns, 0, len(yyDollar[5].updateExprs))
			vals := make(ValTuple, 0, len(yyDollar[5].updateExprs))
			for _, col := range yyDollar[5].updateExprs {
				cols = append(cols, &NonStarExpr{Expr: col.Name})
				vals = append(vals, col.Expr)
			}
			yyVAL.statement = &Insert{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[3].tableName, Columns: cols, Rows: Values{vals}, OnDup: OnDup(yyDollar[6].updateExprs)}
		}
	case 19:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line sql.y:211
		{
			yyVAL.statement = &Update{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[3].tableName, Exprs: yyDollar[5].updateExprs, Where: NewWhere(AST_WHERE, yyDollar[6].boolExpr), OrderBy: yyDollar[7].orderBy, Limit: yyDollar[8].limit}
		}
	case 20:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:217
		{
			yyVAL.statement = &Replace{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[3].tableName, Exprs: yyDollar[5].updateExprs}
		}
	case 21:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:223
		{
			yyVAL.statement = &Delete{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Where: NewWhere(AST_WHERE, yyDollar[5].boolExpr), OrderBy: yyDollar[6].orderBy, Limit: yyDollar[7].limit}
		}
	case 22:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:229
		{
			yyVAL.statement = &Set{Comments: Comments(yyDollar[2].bytes2), Exprs: yyDollar[3].updateExprs}
		}
	case 23:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:235
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyDollar[4].bytes}
		}
	case 24:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line sql.y:239
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[7].bytes, NewName: yyDollar[7].bytes}
		}
	case 25:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:244
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyDollar[3].bytes}
		}
	case 26:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:250
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[4].bytes, NewName: yyDollar[4].bytes}
		}
	case 27:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:254
		{
			// Change this to a rename statement
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyDollar[4].bytes, NewName: yyDollar[7].bytes}
		}
	case 28:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:259
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[3].bytes, NewName: yyDollar[3].bytes}
		}
	case 29:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:265
		{
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyDollar[3].bytes, NewName: yyDollar[5].bytes}
		}
	case 30:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:271
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyDollar[4].bytes}
		}
	case 31:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:275
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[5].bytes, NewName: yyDollar[5].bytes}
		}
	case 32:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:280
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyDollar[4].bytes}
		}
	case 33:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:286
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[3].bytes, NewName: yyDollar[3].bytes}
		}
	case 34:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:292
		{
			yyVAL.statement = &Other{}
		}
	case 35:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:296
		{
			yyVAL.statement = &Other{}
		}
	case 36:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:300
		{
			yyVAL.statement = &Other{}
		}
	case 37:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:305
		{
			SetAllowComments(yylex, true)
		}
	case 38:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:309
		{
			yyVAL.bytes2 = yyDollar[2].bytes2
			SetAllowComments(yylex, false)
		}
	case 39:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:315
		{
			yyVAL.bytes2 = nil
		}
	case 40:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:319
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[2].bytes)
		}
	case 41:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:325
		{
			yyVAL.str = AST_UNION
		}
	case 42:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:329
		{
			yyVAL.str = AST_UNION_ALL
		}
	case 43:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:333
		{
			yyVAL.str = AST_SET_MINUS
		}
	case 44:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:337
		{
			yyVAL.str = AST_EXCEPT
		}
	case 45:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:341
		{
			yyVAL.str = AST_INTERSECT
		}
	case 46:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:346
		{
			yyVAL.str = ""
		}
	case 47:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:350
		{
			yyVAL.str = AST_DISTINCT
		}
	case 48:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:356
		{
			yyVAL.selectExprs = SelectExprs{yyDollar[1].selectExpr}
		}
	case 49:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:360
		{
			yyVAL.selectExprs = append(yyVAL.selectExprs, yyDollar[3].selectExpr)
		}
	case 50:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:366
		{
			yyVAL.selectExpr = &StarExpr{}
		}
	case 51:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:370
		{
			yyVAL.selectExpr = &NonStarExpr{Expr: yyDollar[1].expr, As: yyDollar[2].bytes}
		}
	case 52:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:374
		{
			yyVAL.selectExpr = &StarExpr{TableName: yyDollar[1].bytes}
		}
	case 53:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:380
		{
			yyVAL.expr = yyDollar[1].boolExpr
		}
	case 54:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:384
		{
			yyVAL.expr = yyDollar[1].valExpr
		}
	case 55:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:389
		{
			yyVAL.bytes = nil
		}
	case 56:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:393
		{
			yyVAL.bytes = yyDollar[1].bytes
		}
	case 57:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:397
		{
			yyVAL.bytes = yyDollar[2].bytes
		}
	case 58:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:403
		{
			yyVAL.tableExprs = TableExprs{yyDollar[1].tableExpr}
		}
	case 59:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:407
		{
			yyVAL.tableExprs = append(yyVAL.tableExprs, yyDollar[3].tableExpr)
		}
	case 60:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:413
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyDollar[1].smTableExpr, As: yyDollar[2].bytes, Hints: yyDollar[3].indexHints}
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:417
		{
			yyVAL.tableExpr = &ParenTableExpr{Expr: yyDollar[2].tableExpr}
		}
	case 62:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:421
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr}
		}
	case 63:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:425
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr, On: yyDollar[5].boolExpr}
		}
	case 64:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:430
		{
			yyVAL.bytes = nil
		}
	case 65:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:434
		{
			yyVAL.bytes = yyDollar[1].bytes
		}
	case 66:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:438
		{
			yyVAL.bytes = yyDollar[2].bytes
		}
	case 67:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:444
		{
			yyVAL.str = AST_JOIN
		}
	case 68:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:448
		{
			yyVAL.str = AST_STRAIGHT_JOIN
		}
	case 69:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:452
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 70:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:456
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 71:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:460
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 72:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:464
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 73:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:468
		{
			yyVAL.str = AST_JOIN
		}
	case 74:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:472
		{
			yyVAL.str = AST_CROSS_JOIN
		}
	case 75:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:476
		{
			yyVAL.str = AST_NATURAL_JOIN
		}
	case 76:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:482
		{
			yyVAL.smTableExpr = &TableName{Name: yyDollar[1].bytes}
		}
	case 77:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:486
		{
			yyVAL.smTableExpr = &TableName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 78:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:490
		{
			yyVAL.smTableExpr = yyDollar[1].subquery
		}
	case 79:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:496
		{
			yyVAL.tableName = &TableName{Name: yyDollar[1].bytes}
		}
	case 80:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:500
		{
			yyVAL.tableName = &TableName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 81:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:505
		{
			yyVAL.indexHints = nil
		}
	case 82:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:509
		{
			yyVAL.indexHints = &IndexHints{Type: AST_USE, Indexes: yyDollar[4].bytes2}
		}
	case 83:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:513
		{
			yyVAL.indexHints = &IndexHints{Type: AST_IGNORE, Indexes: yyDollar[4].bytes2}
		}
	case 84:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:517
		{
			yyVAL.indexHints = &IndexHints{Type: AST_FORCE, Indexes: yyDollar[4].bytes2}
		}
	case 85:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:523
		{
			yyVAL.bytes2 = [][]byte{yyDollar[1].bytes}
		}
	case 86:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:527
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[3].bytes)
		}
	case 87:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:532
		{
			yyVAL.boolExpr = nil
		}
	case 88:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:536
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 90:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:543
		{
			yyVAL.boolExpr = &AndExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 91:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:547
		{
			yyVAL.boolExpr = &OrExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 92:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:551
		{
			yyVAL.boolExpr = &NotExpr{Expr: yyDollar[2].boolExpr}
		}
	case 93:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:555
		{
			yyVAL.boolExpr = &ParenBoolExpr{Expr: yyDollar[2].boolExpr}
		}
	case 94:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:561
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: yyDollar[2].str, Right: yyDollar[3].valExpr}
		}
	case 95:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:565
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_IN, Right: yyDollar[3].colTuple}
		}
	case 96:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:569
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_NOT_IN, Right: yyDollar[4].colTuple}
		}
	case 97:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:573
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_LIKE, Right: yyDollar[3].valExpr}
		}
	case 98:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:577
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_NOT_LIKE, Right: yyDollar[4].valExpr}
		}
	case 99:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:581
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: AST_BETWEEN, From: yyDollar[3].valExpr, To: yyDollar[5].valExpr}
		}
	case 100:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:585
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: AST_NOT_BETWEEN, From: yyDollar[4].valExpr, To: yyDollar[6].valExpr}
		}
	case 101:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:589
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NULL, Expr: yyDollar[1].valExpr}
		}
	case 102:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:593
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NOT_NULL, Expr: yyDollar[1].valExpr}
		}
	case 103:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:597
		{
			yyVAL.boolExpr = &ExistsExpr{Subquery: yyDollar[2].subquery}
		}
	case 104:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:601
		{
			yyVAL.boolExpr = &KeyrangeExpr{Start: yyDollar[3].valExpr, End: yyDollar[5].valExpr}
		}
	case 105:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:607
		{
			yyVAL.str = AST_EQ
		}
	case 106:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:611
		{
			yyVAL.str = AST_LT
		}
	case 107:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:615
		{
			yyVAL.str = AST_GT
		}
	case 108:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:619
		{
			yyVAL.str = AST_LE
		}
	case 109:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:623
		{
			yyVAL.str = AST_GE
		}
	case 110:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:627
		{
			yyVAL.str = AST_NE
		}
	case 111:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:631
		{
			yyVAL.str = AST_NSE
		}
	case 112:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:637
		{
			yyVAL.colTuple = ValTuple(yyDollar[2].valExprs)
		}
	case 113:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:641
		{
			yyVAL.colTuple = yyDollar[1].subquery
		}
	case 114:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:645
		{
			yyVAL.colTuple = ListArg(yyDollar[1].bytes)
		}
	case 115:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:651
		{
			yyVAL.subquery = &Subquery{yyDollar[2].selStmt}
		}
	case 116:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:657
		{
			yyVAL.valExprs = ValExprs{yyDollar[1].valExpr}
		}
	case 117:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:661
		{
			yyVAL.valExprs = append(yyDollar[1].valExprs, yyDollar[3].valExpr)
		}
	case 118:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:667
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 119:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:671
		{
			yyVAL.valExpr = yyDollar[1].colName
		}
	case 120:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:675
		{
			yyVAL.valExpr = yyDollar[1].rowTuple
		}
	case 121:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:679
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITAND, Right: yyDollar[3].valExpr}
		}
	case 122:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:683
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITOR, Right: yyDollar[3].valExpr}
		}
	case 123:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:687
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITXOR, Right: yyDollar[3].valExpr}
		}
	case 124:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:691
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_PLUS, Right: yyDollar[3].valExpr}
		}
	case 125:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:695
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MINUS, Right: yyDollar[3].valExpr}
		}
	case 126:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:699
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MULT, Right: yyDollar[3].valExpr}
		}
	case 127:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:703
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_DIV, Right: yyDollar[3].valExpr}
		}
	case 128:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:707
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MOD, Right: yyDollar[3].valExpr}
		}
	case 129:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:711
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
	case 130:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:726
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes}
		}
	case 131:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:730
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Exprs: yyDollar[3].selectExprs}
		}
	case 132:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:734
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Distinct: true, Exprs: yyDollar[4].selectExprs}
		}
	case 133:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:738
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Exprs: yyDollar[3].selectExprs}
		}
	case 134:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:742
		{
			yyVAL.valExpr = yyDollar[1].caseExpr
		}
	case 135:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:748
		{
			yyVAL.bytes = IF_BYTES
		}
	case 136:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:752
		{
			yyVAL.bytes = VALUES_BYTES
		}
	case 137:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:758
		{
			yyVAL.byt = AST_UPLUS
		}
	case 138:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:762
		{
			yyVAL.byt = AST_UMINUS
		}
	case 139:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:766
		{
			yyVAL.byt = AST_TILDA
		}
	case 140:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:772
		{
			yyVAL.caseExpr = &CaseExpr{Expr: yyDollar[2].valExpr, Whens: yyDollar[3].whens, Else: yyDollar[4].valExpr}
		}
	case 141:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:777
		{
			yyVAL.valExpr = nil
		}
	case 142:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:781
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 143:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:787
		{
			yyVAL.whens = []*When{yyDollar[1].when}
		}
	case 144:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:791
		{
			yyVAL.whens = append(yyDollar[1].whens, yyDollar[2].when)
		}
	case 145:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:797
		{
			yyVAL.when = &When{Cond: yyDollar[2].boolExpr, Val: yyDollar[4].valExpr}
		}
	case 146:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:802
		{
			yyVAL.valExpr = nil
		}
	case 147:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:806
		{
			yyVAL.valExpr = yyDollar[2].valExpr
		}
	case 148:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:812
		{
			yyVAL.colName = &ColName{Name: yyDollar[1].bytes}
		}
	case 149:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:816
		{
			yyVAL.colName = &ColName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 150:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:822
		{
			yyVAL.valExpr = StrVal(yyDollar[1].bytes)
		}
	case 151:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:826
		{
			yyVAL.valExpr = NumVal(yyDollar[1].bytes)
		}
	case 152:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:830
		{
			yyVAL.valExpr = ValArg(yyDollar[1].bytes)
		}
	case 153:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:834
		{
			yyVAL.valExpr = &NullVal{}
		}
	case 154:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:839
		{
			yyVAL.valExprs = nil
		}
	case 155:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:843
		{
			yyVAL.valExprs = yyDollar[3].valExprs
		}
	case 156:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:848
		{
			yyVAL.boolExpr = nil
		}
	case 157:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:852
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 158:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:857
		{
			yyVAL.orderBy = nil
		}
	case 159:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:861
		{
			yyVAL.orderBy = yyDollar[3].orderBy
		}
	case 160:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:867
		{
			yyVAL.orderBy = OrderBy{yyDollar[1].order}
		}
	case 161:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:871
		{
			yyVAL.orderBy = append(yyDollar[1].orderBy, yyDollar[3].order)
		}
	case 162:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:877
		{
			yyVAL.order = &Order{Expr: yyDollar[1].valExpr, Direction: yyDollar[2].str}
		}
	case 163:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:882
		{
			yyVAL.str = AST_ASC
		}
	case 164:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:886
		{
			yyVAL.str = AST_ASC
		}
	case 165:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:890
		{
			yyVAL.str = AST_DESC
		}
	case 166:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:895
		{
			yyVAL.limit = nil
		}
	case 167:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:899
		{
			yyVAL.limit = &Limit{Rowcount: yyDollar[2].valExpr}
		}
	case 168:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:903
		{
			yyVAL.limit = &Limit{Offset: yyDollar[2].valExpr, Rowcount: yyDollar[4].valExpr}
		}
	case 169:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:908
		{
			yyVAL.str = ""
		}
	case 170:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:912
		{
			yyVAL.str = AST_FOR_UPDATE
		}
	case 171:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:916
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
	case 172:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:929
		{
			yyVAL.columns = nil
		}
	case 173:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:933
		{
			yyVAL.columns = yyDollar[2].columns
		}
	case 174:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:939
		{
			yyVAL.columns = Columns{&NonStarExpr{Expr: yyDollar[1].colName}}
		}
	case 175:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:943
		{
			yyVAL.columns = append(yyVAL.columns, &NonStarExpr{Expr: yyDollar[3].colName})
		}
	case 176:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:948
		{
			yyVAL.updateExprs = nil
		}
	case 177:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:952
		{
			yyVAL.updateExprs = yyDollar[5].updateExprs
		}
	case 178:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:958
		{
			yyVAL.insRows = yyDollar[2].values
		}
	case 179:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:962
		{
			yyVAL.insRows = yyDollar[1].selStmt
		}
	case 180:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:968
		{
			yyVAL.values = Values{yyDollar[1].rowTuple}
		}
	case 181:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:972
		{
			yyVAL.values = append(yyDollar[1].values, yyDollar[3].rowTuple)
		}
	case 182:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:978
		{
			yyVAL.rowTuple = ValTuple(yyDollar[2].valExprs)
		}
	case 183:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:982
		{
			yyVAL.rowTuple = yyDollar[1].subquery
		}
	case 184:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:988
		{
			yyVAL.updateExprs = UpdateExprs{yyDollar[1].updateExpr}
		}
	case 185:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:992
		{
			yyVAL.updateExprs = append(yyDollar[1].updateExprs, yyDollar[3].updateExpr)
		}
	case 186:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:998
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyDollar[1].colName, Expr: yyDollar[3].valExpr}
		}
	case 187:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1003
		{
			yyVAL.empty = struct{}{}
		}
	case 188:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1005
		{
			yyVAL.empty = struct{}{}
		}
	case 189:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1008
		{
			yyVAL.empty = struct{}{}
		}
	case 190:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1010
		{
			yyVAL.empty = struct{}{}
		}
	case 191:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1013
		{
			yyVAL.empty = struct{}{}
		}
	case 192:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1015
		{
			yyVAL.empty = struct{}{}
		}
	case 193:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1019
		{
			yyVAL.empty = struct{}{}
		}
	case 194:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1021
		{
			yyVAL.empty = struct{}{}
		}
	case 195:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1023
		{
			yyVAL.empty = struct{}{}
		}
	case 196:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1025
		{
			yyVAL.empty = struct{}{}
		}
	case 197:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1027
		{
			yyVAL.empty = struct{}{}
		}
	case 198:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1030
		{
			yyVAL.empty = struct{}{}
		}
	case 199:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1032
		{
			yyVAL.empty = struct{}{}
		}
	case 200:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1035
		{
			yyVAL.empty = struct{}{}
		}
	case 201:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1037
		{
			yyVAL.empty = struct{}{}
		}
	case 202:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1040
		{
			yyVAL.empty = struct{}{}
		}
	case 203:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1042
		{
			yyVAL.empty = struct{}{}
		}
	case 204:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1046
		{
			yyVAL.bytes = bytes.ToLower(yyDollar[1].bytes)
		}
	case 205:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1051
		{
			ForceEOF(yylex)
		}
	}
	goto yystack /* stack new state and value */
}
