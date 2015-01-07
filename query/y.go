
//line parser.y:16

package query
import __yyfmt__ "fmt"
//line parser.y:17
		
import (
  "strconv"
  "net"
  "strings"
  "unicode"
  "errors"
)


//line parser.y:29
type parserSymType struct {
	yys int
	num int
  ip net.IP
  str string
  query Query
}

const HOST = 57346
const PORT = 57347
const PROTOCOL = 57348
const AND = 57349
const OR = 57350
const IP = 57351
const NUM = 57352

var parserToknames = []string{
	"HOST",
	"PORT",
	"PROTOCOL",
	"AND",
	"OR",
	"IP",
	"NUM",
}
var parserStatenames = []string{}

const parserEofCode = 1
const parserErrCode = 2
const parserMaxDepth = 200

//line parser.y:73


// The parser uses the type <prefix>Lex as a lexer.  It must provide
// the methods Lex(*<prefix>SymType) int and Error(string).
type parserLex struct {
  in string
  pos int
  err error
}

var tokens = map[string]int{
 "host": HOST,
 "port": PORT,
 "protocol": PROTOCOL,
 "and": AND,
 "or": OR,
}

// The parser calls this method to get each new token.  This
// implementation returns operators and NUM.
func (x *parserLex) Lex(yylval *parserSymType) int {
  for x.pos < len(x.in) && unicode.IsSpace(rune(x.in[x.pos])) {
    x.pos++
  }
  s := x.pos
  var isIP bool
L:
  for x.pos < len(x.in) {
    switch c := x.in[x.pos]; c {
    case ':', '.':
      isIP = true
      x.pos++
    case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f':
      x.pos++
    default:
       break L
    }
  }
  if isIP {
    yylval.ip = net.ParseIP(x.in[s:x.pos])
    if yylval.ip == nil {
      return -1
    }
    if ip4 := yylval.ip.To4(); ip4 != nil {
      yylval.ip = ip4
    }
    return IP
  } else if x.pos != s {
    n, err := strconv.Atoi(x.in[s:x.pos])
    if err != nil { return -1 }
    yylval.num = n
    return NUM
  } else if x.pos >= len(x.in) {
    return 0
  }
  for t, i := range tokens {
    if strings.HasPrefix(x.in[x.pos:], t) {
      x.pos += len(t)
      return i
    }
  }
  switch c := x.in[x.pos]; c {
  case ':', '.', '(', ')':
    return int(c)
  }
  return -1
}

// The parser calls this method on a parse error.
func (x *parserLex) Error(s string) {
	x.err = errors.New(s)
}

func parse(in string) (Query, error) {
  lex := &parserLex{in: in}
	return parserParse(lex), lex.err
}

//line yacctab:1
var parserExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const parserNprod = 8
const parserPrivate = 57344

var parserTokenNames []string
var parserStates []string

const parserLast = 21

var parserAct = []int{

	7, 8, 3, 4, 5, 15, 11, 10, 2, 6,
	9, 7, 8, 1, 0, 0, 13, 14, 0, 0,
	12,
}
var parserPact = []int{

	-2, 4, -1000, 1, -3, -4, -2, -2, -2, -1000,
	-1000, -1000, -7, -1000, -1000, -1000,
}
var parserPgo = []int{

	0, 13, 8,
}
var parserR1 = []int{

	0, 1, 1, 1, 2, 2, 2, 2,
}
var parserR2 = []int{

	0, 1, 3, 3, 2, 2, 2, 3,
}
var parserChk = []int{

	-1000, -1, -2, 4, 5, 6, 11, 7, 8, 9,
	10, 10, -1, -2, -2, 12,
}
var parserDef = []int{

	0, -2, 1, 0, 0, 0, 0, 0, 0, 4,
	5, 6, 0, 2, 3, 7,
}
var parserTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	11, 12,
}
var parserTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10,
}
var parserTok3 = []int{
	0,
}

//line yaccpar:1

/*	parser for yacc output	*/

var parserDebug = 0

type parserLexer interface {
	Lex(lval *parserSymType) int
	Error(s string)
}

const parserFlag = -1000

func parserTokname(c int) string {
	// 4 is TOKSTART above
	if c >= 4 && c-4 < len(parserToknames) {
		if parserToknames[c-4] != "" {
			return parserToknames[c-4]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func parserStatname(s int) string {
	if s >= 0 && s < len(parserStatenames) {
		if parserStatenames[s] != "" {
			return parserStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func parserlex1(lex parserLexer, lval *parserSymType) int {
	c := 0
	char := lex.Lex(lval)
	if char <= 0 {
		c = parserTok1[0]
		goto out
	}
	if char < len(parserTok1) {
		c = parserTok1[char]
		goto out
	}
	if char >= parserPrivate {
		if char < parserPrivate+len(parserTok2) {
			c = parserTok2[char-parserPrivate]
			goto out
		}
	}
	for i := 0; i < len(parserTok3); i += 2 {
		c = parserTok3[i+0]
		if c == char {
			c = parserTok3[i+1]
			goto out
		}
	}

out:
	if c == 0 {
		c = parserTok2[1] /* unknown char */
	}
	if parserDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", parserTokname(c), uint(char))
	}
	return c
}

func parserParse(parserlex parserLexer) int {
	var parsern int
	var parserlval parserSymType
	var parserVAL parserSymType
	parserS := make([]parserSymType, parserMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	parserstate := 0
	parserchar := -1
	parserp := -1
	goto parserstack

ret0:
	return 0

ret1:
	return 1

parserstack:
	/* put a state and value onto the stack */
	if parserDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", parserTokname(parserchar), parserStatname(parserstate))
	}

	parserp++
	if parserp >= len(parserS) {
		nyys := make([]parserSymType, len(parserS)*2)
		copy(nyys, parserS)
		parserS = nyys
	}
	parserS[parserp] = parserVAL
	parserS[parserp].yys = parserstate

parsernewstate:
	parsern = parserPact[parserstate]
	if parsern <= parserFlag {
		goto parserdefault /* simple state */
	}
	if parserchar < 0 {
		parserchar = parserlex1(parserlex, &parserlval)
	}
	parsern += parserchar
	if parsern < 0 || parsern >= parserLast {
		goto parserdefault
	}
	parsern = parserAct[parsern]
	if parserChk[parsern] == parserchar { /* valid shift */
		parserchar = -1
		parserVAL = parserlval
		parserstate = parsern
		if Errflag > 0 {
			Errflag--
		}
		goto parserstack
	}

parserdefault:
	/* default state action */
	parsern = parserDef[parserstate]
	if parsern == -2 {
		if parserchar < 0 {
			parserchar = parserlex1(parserlex, &parserlval)
		}

		/* look through exception table */
		xi := 0
		for {
			if parserExca[xi+0] == -1 && parserExca[xi+1] == parserstate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			parsern = parserExca[xi+0]
			if parsern < 0 || parsern == parserchar {
				break
			}
		}
		parsern = parserExca[xi+1]
		if parsern < 0 {
			goto ret0
		}
	}
	if parsern == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			parserlex.Error("syntax error")
			Nerrs++
			if parserDebug >= 1 {
				__yyfmt__.Printf("%s", parserStatname(parserstate))
				__yyfmt__.Printf(" saw %s\n", parserTokname(parserchar))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for parserp >= 0 {
				parsern = parserPact[parserS[parserp].yys] + parserErrCode
				if parsern >= 0 && parsern < parserLast {
					parserstate = parserAct[parsern] /* simulate a shift of "error" */
					if parserChk[parserstate] == parserErrCode {
						goto parserstack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if parserDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", parserS[parserp].yys)
				}
				parserp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if parserDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", parserTokname(parserchar))
			}
			if parserchar == parserEofCode {
				goto ret1
			}
			parserchar = -1
			goto parsernewstate /* try again in the same state */
		}
	}

	/* reduction by production parsern */
	if parserDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", parsern, parserStatname(parserstate))
	}

	parsernt := parsern
	parserpt := parserp
	_ = parserpt // guard against "declared and not used"

	parserp -= parserR2[parsern]
	parserVAL = parserS[parserp+1]

	/* consult goto table to find next state */
	parsern = parserR1[parsern]
	parserg := parserPgo[parsern]
	parserj := parserg + parserS[parserp].yys + 1

	if parserj >= parserLast {
		parserstate = parserAct[parserg]
	} else {
		parserstate = parserAct[parserj]
		if parserChk[parserstate] != -parsern {
			parserstate = parserAct[parserg]
		}
	}
	// dummy call; replaced with literal code
	switch parsernt {

	case 1:
		parserVAL.query = parserS[parserpt-0].query
	case 2:
		//line parser.y:47
		{
	    parserVAL.query = intersectQuery{parserS[parserpt-2].query, parserS[parserpt-0].query}
	}
	case 3:
		//line parser.y:51
		{
	    parserVAL.query = unionQuery{parserS[parserpt-2].query, parserS[parserpt-0].query}
	}
	case 4:
		//line parser.y:57
		{
	  parserVAL.query = ipQuery{parserS[parserpt-0].ip, parserS[parserpt-0].ip}
	}
	case 5:
		//line parser.y:61
		{
	  parserVAL.query = portQuery(parserS[parserpt-0].num)
	}
	case 6:
		//line parser.y:65
		{
	  parserVAL.query = protocolQuery(parserS[parserpt-0].num)
	}
	case 7:
		//line parser.y:69
		{
	  parserVAL.query = parserS[parserpt-1].query
	}
	}
	goto parserstack /* stack new state and value */
}
