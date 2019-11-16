// Code generated from Calscript.g4 by ANTLR 4.7.2. DO NOT EDIT.

package calscript_lang // Calscript
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 74, 195,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 3, 2, 3, 2, 6, 2,
	67, 10, 2, 13, 2, 14, 2, 68, 3, 3, 3, 3, 3, 3, 3, 4, 6, 4, 75, 10, 4, 13,
	4, 14, 4, 76, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 5, 6, 84, 10, 6, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 92, 10, 7, 3, 8, 5, 8, 95, 10, 8, 3, 8, 3,
	8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 105, 10, 10, 3, 11, 3,
	11, 5, 11, 109, 10, 11, 3, 11, 3, 11, 5, 11, 113, 10, 11, 5, 11, 115, 10,
	11, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15,
	3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 6, 16, 132, 10, 16, 13, 16, 14, 16,
	133, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 20, 3,
	20, 5, 20, 146, 10, 20, 3, 21, 3, 21, 5, 21, 150, 10, 21, 3, 21, 3, 21,
	3, 22, 3, 22, 3, 23, 3, 23, 5, 23, 158, 10, 23, 3, 23, 3, 23, 3, 24, 3,
	24, 5, 24, 164, 10, 24, 3, 24, 3, 24, 3, 25, 6, 25, 169, 10, 25, 13, 25,
	14, 25, 170, 3, 26, 3, 26, 5, 26, 175, 10, 26, 3, 27, 3, 27, 3, 28, 6,
	28, 180, 10, 28, 13, 28, 14, 28, 181, 3, 29, 6, 29, 185, 10, 29, 13, 29,
	14, 29, 186, 3, 30, 3, 30, 3, 31, 3, 31, 3, 32, 3, 32, 3, 32, 2, 2, 33,
	2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38,
	40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 2, 6, 3, 2, 16, 22, 3,
	2, 27, 38, 3, 2, 65, 69, 3, 2, 23, 26, 2, 182, 2, 64, 3, 2, 2, 2, 4, 70,
	3, 2, 2, 2, 6, 74, 3, 2, 2, 2, 8, 78, 3, 2, 2, 2, 10, 80, 3, 2, 2, 2, 12,
	91, 3, 2, 2, 2, 14, 94, 3, 2, 2, 2, 16, 98, 3, 2, 2, 2, 18, 104, 3, 2,
	2, 2, 20, 114, 3, 2, 2, 2, 22, 116, 3, 2, 2, 2, 24, 119, 3, 2, 2, 2, 26,
	122, 3, 2, 2, 2, 28, 124, 3, 2, 2, 2, 30, 126, 3, 2, 2, 2, 32, 135, 3,
	2, 2, 2, 34, 137, 3, 2, 2, 2, 36, 141, 3, 2, 2, 2, 38, 145, 3, 2, 2, 2,
	40, 147, 3, 2, 2, 2, 42, 153, 3, 2, 2, 2, 44, 155, 3, 2, 2, 2, 46, 161,
	3, 2, 2, 2, 48, 168, 3, 2, 2, 2, 50, 174, 3, 2, 2, 2, 52, 176, 3, 2, 2,
	2, 54, 179, 3, 2, 2, 2, 56, 184, 3, 2, 2, 2, 58, 188, 3, 2, 2, 2, 60, 190,
	3, 2, 2, 2, 62, 192, 3, 2, 2, 2, 64, 66, 5, 4, 3, 2, 65, 67, 5, 10, 6,
	2, 66, 65, 3, 2, 2, 2, 67, 68, 3, 2, 2, 2, 68, 66, 3, 2, 2, 2, 68, 69,
	3, 2, 2, 2, 69, 3, 3, 2, 2, 2, 70, 71, 7, 3, 2, 2, 71, 72, 5, 6, 4, 2,
	72, 5, 3, 2, 2, 2, 73, 75, 7, 70, 2, 2, 74, 73, 3, 2, 2, 2, 75, 76, 3,
	2, 2, 2, 76, 74, 3, 2, 2, 2, 76, 77, 3, 2, 2, 2, 77, 7, 3, 2, 2, 2, 78,
	79, 7, 70, 2, 2, 79, 9, 3, 2, 2, 2, 80, 81, 5, 12, 7, 2, 81, 83, 5, 28,
	15, 2, 82, 84, 5, 30, 16, 2, 83, 82, 3, 2, 2, 2, 83, 84, 3, 2, 2, 2, 84,
	11, 3, 2, 2, 2, 85, 86, 7, 8, 2, 2, 86, 92, 5, 38, 20, 2, 87, 88, 7, 9,
	2, 2, 88, 92, 5, 14, 8, 2, 89, 90, 7, 11, 2, 2, 90, 92, 5, 14, 8, 2, 91,
	85, 3, 2, 2, 2, 91, 87, 3, 2, 2, 2, 91, 89, 3, 2, 2, 2, 92, 13, 3, 2, 2,
	2, 93, 95, 5, 16, 9, 2, 94, 93, 3, 2, 2, 2, 94, 95, 3, 2, 2, 2, 95, 96,
	3, 2, 2, 2, 96, 97, 5, 18, 10, 2, 97, 15, 3, 2, 2, 2, 98, 99, 5, 38, 20,
	2, 99, 17, 3, 2, 2, 2, 100, 101, 7, 12, 2, 2, 101, 105, 5, 20, 11, 2, 102,
	103, 7, 10, 2, 2, 103, 105, 5, 20, 11, 2, 104, 100, 3, 2, 2, 2, 104, 102,
	3, 2, 2, 2, 105, 19, 3, 2, 2, 2, 106, 108, 5, 24, 13, 2, 107, 109, 5, 22,
	12, 2, 108, 107, 3, 2, 2, 2, 108, 109, 3, 2, 2, 2, 109, 115, 3, 2, 2, 2,
	110, 112, 5, 38, 20, 2, 111, 113, 5, 22, 12, 2, 112, 111, 3, 2, 2, 2, 112,
	113, 3, 2, 2, 2, 113, 115, 3, 2, 2, 2, 114, 106, 3, 2, 2, 2, 114, 110,
	3, 2, 2, 2, 115, 21, 3, 2, 2, 2, 116, 117, 7, 13, 2, 2, 117, 118, 5, 24,
	13, 2, 118, 23, 3, 2, 2, 2, 119, 120, 7, 72, 2, 2, 120, 121, 5, 26, 14,
	2, 121, 25, 3, 2, 2, 2, 122, 123, 9, 2, 2, 2, 123, 27, 3, 2, 2, 2, 124,
	125, 7, 71, 2, 2, 125, 29, 3, 2, 2, 2, 126, 127, 7, 14, 2, 2, 127, 128,
	7, 66, 2, 2, 128, 131, 5, 32, 17, 2, 129, 130, 7, 7, 2, 2, 130, 132, 5,
	32, 17, 2, 131, 129, 3, 2, 2, 2, 132, 133, 3, 2, 2, 2, 133, 131, 3, 2,
	2, 2, 133, 134, 3, 2, 2, 2, 134, 31, 3, 2, 2, 2, 135, 136, 7, 70, 2, 2,
	136, 33, 3, 2, 2, 2, 137, 138, 7, 15, 2, 2, 138, 139, 7, 66, 2, 2, 139,
	140, 5, 36, 19, 2, 140, 35, 3, 2, 2, 2, 141, 142, 7, 70, 2, 2, 142, 37,
	3, 2, 2, 2, 143, 146, 5, 62, 32, 2, 144, 146, 5, 42, 22, 2, 145, 143, 3,
	2, 2, 2, 145, 144, 3, 2, 2, 2, 146, 39, 3, 2, 2, 2, 147, 149, 5, 42, 22,
	2, 148, 150, 11, 2, 2, 2, 149, 148, 3, 2, 2, 2, 149, 150, 3, 2, 2, 2, 150,
	151, 3, 2, 2, 2, 151, 152, 5, 60, 31, 2, 152, 41, 3, 2, 2, 2, 153, 154,
	5, 44, 23, 2, 154, 43, 3, 2, 2, 2, 155, 157, 5, 48, 25, 2, 156, 158, 5,
	58, 30, 2, 157, 156, 3, 2, 2, 2, 157, 158, 3, 2, 2, 2, 158, 159, 3, 2,
	2, 2, 159, 160, 5, 46, 24, 2, 160, 45, 3, 2, 2, 2, 161, 163, 5, 50, 26,
	2, 162, 164, 5, 58, 30, 2, 163, 162, 3, 2, 2, 2, 163, 164, 3, 2, 2, 2,
	164, 165, 3, 2, 2, 2, 165, 166, 5, 56, 29, 2, 166, 47, 3, 2, 2, 2, 167,
	169, 7, 72, 2, 2, 168, 167, 3, 2, 2, 2, 169, 170, 3, 2, 2, 2, 170, 168,
	3, 2, 2, 2, 170, 171, 3, 2, 2, 2, 171, 49, 3, 2, 2, 2, 172, 175, 5, 52,
	27, 2, 173, 175, 5, 54, 28, 2, 174, 172, 3, 2, 2, 2, 174, 173, 3, 2, 2,
	2, 175, 51, 3, 2, 2, 2, 176, 177, 9, 3, 2, 2, 177, 53, 3, 2, 2, 2, 178,
	180, 7, 72, 2, 2, 179, 178, 3, 2, 2, 2, 180, 181, 3, 2, 2, 2, 181, 179,
	3, 2, 2, 2, 181, 182, 3, 2, 2, 2, 182, 55, 3, 2, 2, 2, 183, 185, 7, 72,
	2, 2, 184, 183, 3, 2, 2, 2, 185, 186, 3, 2, 2, 2, 186, 184, 3, 2, 2, 2,
	186, 187, 3, 2, 2, 2, 187, 57, 3, 2, 2, 2, 188, 189, 9, 4, 2, 2, 189, 59,
	3, 2, 2, 2, 190, 191, 3, 2, 2, 2, 191, 61, 3, 2, 2, 2, 192, 193, 9, 5,
	2, 2, 193, 63, 3, 2, 2, 2, 20, 68, 76, 83, 91, 94, 104, 108, 112, 114,
	133, 145, 149, 157, 163, 170, 174, 181, 186,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'calscript'", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "','", "':'", "'/'", "'-'",
	"'.'",
}
var symbolicNames = []string{
	"", "", "WHITESPACE", "COMMENT", "LINE_COMMENT", "LIST_SEPARATOR", "ON",
	"STARTING", "UNTIL", "THEN", "FOR", "EVERY", "WITH", "AT", "DAY", "WEEK",
	"MONTH", "YEAR", "HOUR", "MINUTE", "SECOND", "YESTERDAY", "NOW", "TOMORROW",
	"TODAY", "JANUARY", "FEBRUARY", "MARCH", "APRIL", "MAY", "JUNE", "JULY",
	"AUGUST", "SEPTEMBER", "OCTOBER", "NOVEMBER", "DECEMBER", "A", "B", "C",
	"D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R",
	"S", "T", "U", "V", "W", "X", "Y", "Z", "COMMA", "COLON", "SLASH", "DASH",
	"DOT", "IDENTIFIER", "STRING", "UNSIGNED_INT", "SIGNED_INT", "FLOAT",
}

var ruleNames = []string{
	"calscript", "begin", "script_name", "name", "event", "start_end", "date_period",
	"date_period_start", "date_period_end", "recurrence", "reminders", "recurrence_value",
	"recurrence_value_unit", "description", "attendees", "attendee", "location",
	"location_string", "date", "formal_date_time", "formal_date", "year_month_day",
	"month_day", "year", "month", "vocal_month", "numeric_month", "day_of_month",
	"date_separator", "time", "natural_date",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type CalscriptParser struct {
	*antlr.BaseParser
}

func NewCalscriptParser(input antlr.TokenStream) *CalscriptParser {
	this := new(CalscriptParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Calscript.g4"

	return this
}

// CalscriptParser tokens.
const (
	CalscriptParserEOF            = antlr.TokenEOF
	CalscriptParserT__0           = 1
	CalscriptParserWHITESPACE     = 2
	CalscriptParserCOMMENT        = 3
	CalscriptParserLINE_COMMENT   = 4
	CalscriptParserLIST_SEPARATOR = 5
	CalscriptParserON             = 6
	CalscriptParserSTARTING       = 7
	CalscriptParserUNTIL          = 8
	CalscriptParserTHEN           = 9
	CalscriptParserFOR            = 10
	CalscriptParserEVERY          = 11
	CalscriptParserWITH           = 12
	CalscriptParserAT             = 13
	CalscriptParserDAY            = 14
	CalscriptParserWEEK           = 15
	CalscriptParserMONTH          = 16
	CalscriptParserYEAR           = 17
	CalscriptParserHOUR           = 18
	CalscriptParserMINUTE         = 19
	CalscriptParserSECOND         = 20
	CalscriptParserYESTERDAY      = 21
	CalscriptParserNOW            = 22
	CalscriptParserTOMORROW       = 23
	CalscriptParserTODAY          = 24
	CalscriptParserJANUARY        = 25
	CalscriptParserFEBRUARY       = 26
	CalscriptParserMARCH          = 27
	CalscriptParserAPRIL          = 28
	CalscriptParserMAY            = 29
	CalscriptParserJUNE           = 30
	CalscriptParserJULY           = 31
	CalscriptParserAUGUST         = 32
	CalscriptParserSEPTEMBER      = 33
	CalscriptParserOCTOBER        = 34
	CalscriptParserNOVEMBER       = 35
	CalscriptParserDECEMBER       = 36
	CalscriptParserA              = 37
	CalscriptParserB              = 38
	CalscriptParserC              = 39
	CalscriptParserD              = 40
	CalscriptParserE              = 41
	CalscriptParserF              = 42
	CalscriptParserG              = 43
	CalscriptParserH              = 44
	CalscriptParserI              = 45
	CalscriptParserJ              = 46
	CalscriptParserK              = 47
	CalscriptParserL              = 48
	CalscriptParserM              = 49
	CalscriptParserN              = 50
	CalscriptParserO              = 51
	CalscriptParserP              = 52
	CalscriptParserQ              = 53
	CalscriptParserR              = 54
	CalscriptParserS              = 55
	CalscriptParserT              = 56
	CalscriptParserU              = 57
	CalscriptParserV              = 58
	CalscriptParserW              = 59
	CalscriptParserX              = 60
	CalscriptParserY              = 61
	CalscriptParserZ              = 62
	CalscriptParserCOMMA          = 63
	CalscriptParserCOLON          = 64
	CalscriptParserSLASH          = 65
	CalscriptParserDASH           = 66
	CalscriptParserDOT            = 67
	CalscriptParserIDENTIFIER     = 68
	CalscriptParserSTRING         = 69
	CalscriptParserUNSIGNED_INT   = 70
	CalscriptParserSIGNED_INT     = 71
	CalscriptParserFLOAT          = 72
)

// CalscriptParser rules.
const (
	CalscriptParserRULE_calscript             = 0
	CalscriptParserRULE_begin                 = 1
	CalscriptParserRULE_script_name           = 2
	CalscriptParserRULE_name                  = 3
	CalscriptParserRULE_event                 = 4
	CalscriptParserRULE_start_end             = 5
	CalscriptParserRULE_date_period           = 6
	CalscriptParserRULE_date_period_start     = 7
	CalscriptParserRULE_date_period_end       = 8
	CalscriptParserRULE_recurrence            = 9
	CalscriptParserRULE_reminders             = 10
	CalscriptParserRULE_recurrence_value      = 11
	CalscriptParserRULE_recurrence_value_unit = 12
	CalscriptParserRULE_description           = 13
	CalscriptParserRULE_attendees             = 14
	CalscriptParserRULE_attendee              = 15
	CalscriptParserRULE_location              = 16
	CalscriptParserRULE_location_string       = 17
	CalscriptParserRULE_date                  = 18
	CalscriptParserRULE_formal_date_time      = 19
	CalscriptParserRULE_formal_date           = 20
	CalscriptParserRULE_year_month_day        = 21
	CalscriptParserRULE_month_day             = 22
	CalscriptParserRULE_year                  = 23
	CalscriptParserRULE_month                 = 24
	CalscriptParserRULE_vocal_month           = 25
	CalscriptParserRULE_numeric_month         = 26
	CalscriptParserRULE_day_of_month          = 27
	CalscriptParserRULE_date_separator        = 28
	CalscriptParserRULE_time                  = 29
	CalscriptParserRULE_natural_date          = 30
)

// ICalscriptContext is an interface to support dynamic dispatch.
type ICalscriptContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCalscriptContext differentiates from other interfaces.
	IsCalscriptContext()
}

type CalscriptContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCalscriptContext() *CalscriptContext {
	var p = new(CalscriptContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalscriptParserRULE_calscript
	return p
}

func (*CalscriptContext) IsCalscriptContext() {}

func NewCalscriptContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CalscriptContext {
	var p = new(CalscriptContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalscriptParserRULE_calscript

	return p
}

func (s *CalscriptContext) GetParser() antlr.Parser { return s.parser }

func (s *CalscriptContext) Begin() IBeginContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBeginContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBeginContext)
}

func (s *CalscriptContext) AllEvent() []IEventContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEventContext)(nil)).Elem())
	var tst = make([]IEventContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEventContext)
		}
	}

	return tst
}

func (s *CalscriptContext) Event(i int) IEventContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEventContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEventContext)
}

func (s *CalscriptContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CalscriptContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CalscriptContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalscriptListener); ok {
		listenerT.EnterCalscript(s)
	}
}

func (s *CalscriptContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalscriptListener); ok {
		listenerT.ExitCalscript(s)
	}
}

func (p *CalscriptParser) Calscript() (localctx ICalscriptContext) {
	localctx = NewCalscriptContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CalscriptParserRULE_calscript)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(62)
		p.Begin()
	}
	p.SetState(64)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<CalscriptParserON)|(1<<CalscriptParserSTARTING)|(1<<CalscriptParserTHEN))) != 0) {
		{
			p.SetState(63)
			p.Event()
		}

		p.SetState(66)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IBeginContext is an interface to support dynamic dispatch.
type IBeginContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBeginContext differentiates from other interfaces.
	IsBeginContext()
}

type BeginContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBeginContext() *BeginContext {
	var p = new(BeginContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalscriptParserRULE_begin
	return p
}

func (*BeginContext) IsBeginContext() {}

func NewBeginContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BeginContext {
	var p = new(BeginContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalscriptParserRULE_begin

	return p
}

func (s *BeginContext) GetParser() antlr.Parser { return s.parser }

func (s *BeginContext) Script_name() IScript_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IScript_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IScript_nameContext)
}

func (s *BeginContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BeginContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BeginContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalscriptListener); ok {
		listenerT.EnterBegin(s)
	}
}

func (s *BeginContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalscriptListener); ok {
		listenerT.ExitBegin(s)
	}
}

func (p *CalscriptParser) Begin() (localctx IBeginContext) {
	localctx = NewBeginContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CalscriptParserRULE_begin)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(68)
		p.Match(CalscriptParserT__0)
	}
	{
		p.SetState(69)
		p.Script_name()
	}

	return localctx
}

// IScript_nameContext is an interface to support dynamic dispatch.
type IScript_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsScript_nameContext differentiates from other interfaces.
	IsScript_nameContext()
}

type Script_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyScript_nameContext() *Script_nameContext {
	var p = new(Script_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalscriptParserRULE_script_name
	return p
}

func (*Script_nameContext) IsScript_nameContext() {}

func NewScript_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Script_nameContext {
	var p = new(Script_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalscriptParserRULE_script_name

	return p
}

func (s *Script_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Script_nameContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(CalscriptParserIDENTIFIER)
}

func (s *Script_nameContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(CalscriptParserIDENTIFIER, i)
}

func (s *Script_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Script_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Script_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalscriptListener); ok {
		listenerT.EnterScript_name(s)
	}
}

func (s *Script_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalscriptListener); ok {
		listenerT.ExitScript_name(s)
	}
}

func (p *CalscriptParser) Script_name() (localctx IScript_nameContext) {
	localctx = NewScript_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, CalscriptParserRULE_script_name)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(72)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == CalscriptParserIDENTIFIER {
		{
			p.SetState(71)
			p.Match(CalscriptParserIDENTIFIER)
		}

		p.SetState(74)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// INameContext is an interface to support dynamic dispatch.
type INameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNameContext differentiates from other interfaces.
	IsNameContext()
}

type NameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNameContext() *NameContext {
	var p = new(NameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalscriptParserRULE_name
	return p
}

func (*NameContext) IsNameContext() {}

func NewNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NameContext {
	var p = new(NameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalscriptParserRULE_name

	return p
}

func (s *NameContext) GetParser() antlr.Parser { return s.parser }

func (s *NameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CalscriptParserIDENTIFIER, 0)
}

func (s *NameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalscriptListener); ok {
		listenerT.EnterName(s)
	}
}

func (s *NameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalscriptListener); ok {
		listenerT.ExitName(s)
	}
}

func (p *CalscriptParser) Name() (localctx INameContext) {
	localctx = NewNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, CalscriptParserRULE_name)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(76)
		p.Match(CalscriptParserIDENTIFIER)
	}

	return localctx
}

// IEventContext is an interface to support dynamic dispatch.
type IEventContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEventContext differentiates from other interfaces.
	IsEventContext()
}

type EventContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEventContext() *EventContext {
	var p = new(EventContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalscriptParserRULE_event
	return p
}

func (*EventContext) IsEventContext() {}

func NewEventContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EventContext {
	var p = new(EventContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalscriptParserRULE_event

	return p
}

func (s *EventContext) GetParser() antlr.Parser { return s.parser }

func (s *EventContext) Start_end() IStart_endContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStart_endContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStart_endContext)
}

func (s *EventContext) Description() IDescriptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDescriptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDescriptionContext)
}

func (s *EventContext) Attendees() IAttendeesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttendeesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAttendeesContext)
}

func (s *EventContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EventContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EventContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalscriptListener); ok {
		listenerT.EnterEvent(s)
	}
}

func (s *EventContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalscriptListener); ok {
		listenerT.ExitEvent(s)
	}
}

func (p *CalscriptParser) Event() (localctx IEventContext) {
	localctx = NewEventContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, CalscriptParserRULE_event)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(78)
		p.Start_end()
	}
	{
		p.SetState(79)
		p.Description()
	}
	p.SetState(81)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CalscriptParserWITH {
		{
			p.SetState(80)
			p.Attendees()
		}

	}

	return localctx
}

// IStart_endContext is an interface to support dynamic dispatch.
type IStart_endContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStart_endContext differentiates from other interfaces.
	IsStart_endContext()
}

type Start_endContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStart_endContext() *Start_endContext {
	var p = new(Start_endContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalscriptParserRULE_start_end
	return p
}

func (*Start_endContext) IsStart_endContext() {}

func NewStart_endContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Start_endContext {
	var p = new(Start_endContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalscriptParserRULE_start_end

	return p
}

func (s *Start_endContext) GetParser() antlr.Parser { return s.parser }

func (s *Start_endContext) ON() antlr.TerminalNode {
	return s.GetToken(CalscriptParserON, 0)
}

func (s *Start_endContext) Date() IDateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDateContext)
}

func (s *Start_endContext) STARTING() antlr.TerminalNode {
	return s.GetToken(CalscriptParserSTARTING, 0)
}

func (s *Start_endContext) Date_period() IDate_periodContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDate_periodContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDate_periodContext)
}

func (s *Start_endContext) THEN() antlr.TerminalNode {
	return s.GetToken(CalscriptParserTHEN, 0)
}

func (s *Start_endContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Start_endContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Start_endContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalscriptListener); ok {
		listenerT.EnterStart_end(s)
	}
}

func (s *Start_endContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalscriptListener); ok {
		listenerT.ExitStart_end(s)
	}
}

func (p *CalscriptParser) Start_end() (localctx IStart_endContext) {
	localctx = NewStart_endContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, CalscriptParserRULE_start_end)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(89)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CalscriptParserON:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(83)
			p.Match(CalscriptParserON)
		}
		{
			p.SetState(84)
			p.Date()
		}

	case CalscriptParserSTARTING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(85)
			p.Match(CalscriptParserSTARTING)
		}
		{
			p.SetState(86)
			p.Date_period()
		}

	case CalscriptParserTHEN:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(87)
			p.Match(CalscriptParserTHEN)
		}
		{
			p.SetState(88)
			p.Date_period()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IDate_periodContext is an interface to support dynamic dispatch.
type IDate_periodContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDate_periodContext differentiates from other interfaces.
	IsDate_periodContext()
}

type Date_periodContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDate_periodContext() *Date_periodContext {
	var p = new(Date_periodContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalscriptParserRULE_date_period
	return p
}

func (*Date_periodContext) IsDate_periodContext() {}

func NewDate_periodContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Date_periodContext {
	var p = new(Date_periodContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalscriptParserRULE_date_period

	return p
}

func (s *Date_periodContext) GetParser() antlr.Parser { return s.parser }

func (s *Date_periodContext) Date_period_end() IDate_period_endContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDate_period_endContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDate_period_endContext)
}

func (s *Date_periodContext) Date_period_start() IDate_period_startContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDate_period_startContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDate_period_startContext)
}

func (s *Date_periodContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Date_periodContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Date_periodContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalscriptListener); ok {
		listenerT.EnterDate_period(s)
	}
}

func (s *Date_periodContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalscriptListener); ok {
		listenerT.ExitDate_period(s)
	}
}

func (p *CalscriptParser) Date_period() (localctx IDate_periodContext) {
	localctx = NewDate_periodContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, CalscriptParserRULE_date_period)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(92)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<CalscriptParserYESTERDAY)|(1<<CalscriptParserNOW)|(1<<CalscriptParserTOMORROW)|(1<<CalscriptParserTODAY))) != 0) || _la == CalscriptParserUNSIGNED_INT {
		{
			p.SetState(91)
			p.Date_period_start()
		}

	}
	{
		p.SetState(94)
		p.Date_period_end()
	}

	return localctx
}

// IDate_period_startContext is an interface to support dynamic dispatch.
type IDate_period_startContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDate_period_startContext differentiates from other interfaces.
	IsDate_period_startContext()
}

type Date_period_startContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDate_period_startContext() *Date_period_startContext {
	var p = new(Date_period_startContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalscriptParserRULE_date_period_start
	return p
}

func (*Date_period_startContext) IsDate_period_startContext() {}

func NewDate_period_startContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Date_period_startContext {
	var p = new(Date_period_startContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalscriptParserRULE_date_period_start

	return p
}

func (s *Date_period_startContext) GetParser() antlr.Parser { return s.parser }

func (s *Date_period_startContext) Date() IDateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDateContext)
}

func (s *Date_period_startContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Date_period_startContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Date_period_startContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalscriptListener); ok {
		listenerT.EnterDate_period_start(s)
	}
}

func (s *Date_period_startContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalscriptListener); ok {
		listenerT.ExitDate_period_start(s)
	}
}

func (p *CalscriptParser) Date_period_start() (localctx IDate_period_startContext) {
	localctx = NewDate_period_startContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, CalscriptParserRULE_date_period_start)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(96)
		p.Date()
	}

	return localctx
}

// IDate_period_endContext is an interface to support dynamic dispatch.
type IDate_period_endContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDate_period_endContext differentiates from other interfaces.
	IsDate_period_endContext()
}

type Date_period_endContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDate_period_endContext() *Date_period_endContext {
	var p = new(Date_period_endContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalscriptParserRULE_date_period_end
	return p
}

func (*Date_period_endContext) IsDate_period_endContext() {}

func NewDate_period_endContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Date_period_endContext {
	var p = new(Date_period_endContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalscriptParserRULE_date_period_end

	return p
}

func (s *Date_period_endContext) GetParser() antlr.Parser { return s.parser }

func (s *Date_period_endContext) FOR() antlr.TerminalNode {
	return s.GetToken(CalscriptParserFOR, 0)
}

func (s *Date_period_endContext) Recurrence() IRecurrenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRecurrenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRecurrenceContext)
}

func (s *Date_period_endContext) UNTIL() antlr.TerminalNode {
	return s.GetToken(CalscriptParserUNTIL, 0)
}

func (s *Date_period_endContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Date_period_endContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Date_period_endContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalscriptListener); ok {
		listenerT.EnterDate_period_end(s)
	}
}

func (s *Date_period_endContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalscriptListener); ok {
		listenerT.ExitDate_period_end(s)
	}
}

func (p *CalscriptParser) Date_period_end() (localctx IDate_period_endContext) {
	localctx = NewDate_period_endContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, CalscriptParserRULE_date_period_end)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(102)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CalscriptParserFOR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(98)
			p.Match(CalscriptParserFOR)
		}
		{
			p.SetState(99)
			p.Recurrence()
		}

	case CalscriptParserUNTIL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(100)
			p.Match(CalscriptParserUNTIL)
		}
		{
			p.SetState(101)
			p.Recurrence()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IRecurrenceContext is an interface to support dynamic dispatch.
type IRecurrenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRecurrenceContext differentiates from other interfaces.
	IsRecurrenceContext()
}

type RecurrenceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRecurrenceContext() *RecurrenceContext {
	var p = new(RecurrenceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalscriptParserRULE_recurrence
	return p
}

func (*RecurrenceContext) IsRecurrenceContext() {}

func NewRecurrenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RecurrenceContext {
	var p = new(RecurrenceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalscriptParserRULE_recurrence

	return p
}

func (s *RecurrenceContext) GetParser() antlr.Parser { return s.parser }

func (s *RecurrenceContext) Recurrence_value() IRecurrence_valueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRecurrence_valueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRecurrence_valueContext)
}

func (s *RecurrenceContext) Reminders() IRemindersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRemindersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRemindersContext)
}

func (s *RecurrenceContext) Date() IDateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDateContext)
}

func (s *RecurrenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RecurrenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RecurrenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalscriptListener); ok {
		listenerT.EnterRecurrence(s)
	}
}

func (s *RecurrenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalscriptListener); ok {
		listenerT.ExitRecurrence(s)
	}
}

func (p *CalscriptParser) Recurrence() (localctx IRecurrenceContext) {
	localctx = NewRecurrenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, CalscriptParserRULE_recurrence)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(112)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(104)
			p.Recurrence_value()
		}
		p.SetState(106)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CalscriptParserEVERY {
			{
				p.SetState(105)
				p.Reminders()
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(108)
			p.Date()
		}
		p.SetState(110)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CalscriptParserEVERY {
			{
				p.SetState(109)
				p.Reminders()
			}

		}

	}

	return localctx
}

// IRemindersContext is an interface to support dynamic dispatch.
type IRemindersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRemindersContext differentiates from other interfaces.
	IsRemindersContext()
}

type RemindersContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRemindersContext() *RemindersContext {
	var p = new(RemindersContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalscriptParserRULE_reminders
	return p
}

func (*RemindersContext) IsRemindersContext() {}

func NewRemindersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RemindersContext {
	var p = new(RemindersContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalscriptParserRULE_reminders

	return p
}

func (s *RemindersContext) GetParser() antlr.Parser { return s.parser }

func (s *RemindersContext) EVERY() antlr.TerminalNode {
	return s.GetToken(CalscriptParserEVERY, 0)
}

func (s *RemindersContext) Recurrence_value() IRecurrence_valueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRecurrence_valueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRecurrence_valueContext)
}

func (s *RemindersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RemindersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RemindersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalscriptListener); ok {
		listenerT.EnterReminders(s)
	}
}

func (s *RemindersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalscriptListener); ok {
		listenerT.ExitReminders(s)
	}
}

func (p *CalscriptParser) Reminders() (localctx IRemindersContext) {
	localctx = NewRemindersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, CalscriptParserRULE_reminders)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(114)
		p.Match(CalscriptParserEVERY)
	}
	{
		p.SetState(115)
		p.Recurrence_value()
	}

	return localctx
}

// IRecurrence_valueContext is an interface to support dynamic dispatch.
type IRecurrence_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRecurrence_valueContext differentiates from other interfaces.
	IsRecurrence_valueContext()
}

type Recurrence_valueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRecurrence_valueContext() *Recurrence_valueContext {
	var p = new(Recurrence_valueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalscriptParserRULE_recurrence_value
	return p
}

func (*Recurrence_valueContext) IsRecurrence_valueContext() {}

func NewRecurrence_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Recurrence_valueContext {
	var p = new(Recurrence_valueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalscriptParserRULE_recurrence_value

	return p
}

func (s *Recurrence_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Recurrence_valueContext) UNSIGNED_INT() antlr.TerminalNode {
	return s.GetToken(CalscriptParserUNSIGNED_INT, 0)
}

func (s *Recurrence_valueContext) Recurrence_value_unit() IRecurrence_value_unitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRecurrence_value_unitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRecurrence_value_unitContext)
}

func (s *Recurrence_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Recurrence_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Recurrence_valueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalscriptListener); ok {
		listenerT.EnterRecurrence_value(s)
	}
}

func (s *Recurrence_valueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalscriptListener); ok {
		listenerT.ExitRecurrence_value(s)
	}
}

func (p *CalscriptParser) Recurrence_value() (localctx IRecurrence_valueContext) {
	localctx = NewRecurrence_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, CalscriptParserRULE_recurrence_value)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(117)
		p.Match(CalscriptParserUNSIGNED_INT)
	}
	{
		p.SetState(118)
		p.Recurrence_value_unit()
	}

	return localctx
}

// IRecurrence_value_unitContext is an interface to support dynamic dispatch.
type IRecurrence_value_unitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRecurrence_value_unitContext differentiates from other interfaces.
	IsRecurrence_value_unitContext()
}

type Recurrence_value_unitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRecurrence_value_unitContext() *Recurrence_value_unitContext {
	var p = new(Recurrence_value_unitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalscriptParserRULE_recurrence_value_unit
	return p
}

func (*Recurrence_value_unitContext) IsRecurrence_value_unitContext() {}

func NewRecurrence_value_unitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Recurrence_value_unitContext {
	var p = new(Recurrence_value_unitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalscriptParserRULE_recurrence_value_unit

	return p
}

func (s *Recurrence_value_unitContext) GetParser() antlr.Parser { return s.parser }

func (s *Recurrence_value_unitContext) DAY() antlr.TerminalNode {
	return s.GetToken(CalscriptParserDAY, 0)
}

func (s *Recurrence_value_unitContext) WEEK() antlr.TerminalNode {
	return s.GetToken(CalscriptParserWEEK, 0)
}

func (s *Recurrence_value_unitContext) MONTH() antlr.TerminalNode {
	return s.GetToken(CalscriptParserMONTH, 0)
}

func (s *Recurrence_value_unitContext) YEAR() antlr.TerminalNode {
	return s.GetToken(CalscriptParserYEAR, 0)
}

func (s *Recurrence_value_unitContext) HOUR() antlr.TerminalNode {
	return s.GetToken(CalscriptParserHOUR, 0)
}

func (s *Recurrence_value_unitContext) MINUTE() antlr.TerminalNode {
	return s.GetToken(CalscriptParserMINUTE, 0)
}

func (s *Recurrence_value_unitContext) SECOND() antlr.TerminalNode {
	return s.GetToken(CalscriptParserSECOND, 0)
}

func (s *Recurrence_value_unitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Recurrence_value_unitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Recurrence_value_unitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalscriptListener); ok {
		listenerT.EnterRecurrence_value_unit(s)
	}
}

func (s *Recurrence_value_unitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalscriptListener); ok {
		listenerT.ExitRecurrence_value_unit(s)
	}
}

func (p *CalscriptParser) Recurrence_value_unit() (localctx IRecurrence_value_unitContext) {
	localctx = NewRecurrence_value_unitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, CalscriptParserRULE_recurrence_value_unit)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(120)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<CalscriptParserDAY)|(1<<CalscriptParserWEEK)|(1<<CalscriptParserMONTH)|(1<<CalscriptParserYEAR)|(1<<CalscriptParserHOUR)|(1<<CalscriptParserMINUTE)|(1<<CalscriptParserSECOND))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IDescriptionContext is an interface to support dynamic dispatch.
type IDescriptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDescriptionContext differentiates from other interfaces.
	IsDescriptionContext()
}

type DescriptionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDescriptionContext() *DescriptionContext {
	var p = new(DescriptionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalscriptParserRULE_description
	return p
}

func (*DescriptionContext) IsDescriptionContext() {}

func NewDescriptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DescriptionContext {
	var p = new(DescriptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalscriptParserRULE_description

	return p
}

func (s *DescriptionContext) GetParser() antlr.Parser { return s.parser }

func (s *DescriptionContext) STRING() antlr.TerminalNode {
	return s.GetToken(CalscriptParserSTRING, 0)
}

func (s *DescriptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DescriptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DescriptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalscriptListener); ok {
		listenerT.EnterDescription(s)
	}
}

func (s *DescriptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalscriptListener); ok {
		listenerT.ExitDescription(s)
	}
}

func (p *CalscriptParser) Description() (localctx IDescriptionContext) {
	localctx = NewDescriptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, CalscriptParserRULE_description)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(122)
		p.Match(CalscriptParserSTRING)
	}

	return localctx
}

// IAttendeesContext is an interface to support dynamic dispatch.
type IAttendeesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAttendeesContext differentiates from other interfaces.
	IsAttendeesContext()
}

type AttendeesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttendeesContext() *AttendeesContext {
	var p = new(AttendeesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalscriptParserRULE_attendees
	return p
}

func (*AttendeesContext) IsAttendeesContext() {}

func NewAttendeesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttendeesContext {
	var p = new(AttendeesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalscriptParserRULE_attendees

	return p
}

func (s *AttendeesContext) GetParser() antlr.Parser { return s.parser }

func (s *AttendeesContext) WITH() antlr.TerminalNode {
	return s.GetToken(CalscriptParserWITH, 0)
}

func (s *AttendeesContext) COLON() antlr.TerminalNode {
	return s.GetToken(CalscriptParserCOLON, 0)
}

func (s *AttendeesContext) AllAttendee() []IAttendeeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAttendeeContext)(nil)).Elem())
	var tst = make([]IAttendeeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAttendeeContext)
		}
	}

	return tst
}

func (s *AttendeesContext) Attendee(i int) IAttendeeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttendeeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAttendeeContext)
}

func (s *AttendeesContext) AllLIST_SEPARATOR() []antlr.TerminalNode {
	return s.GetTokens(CalscriptParserLIST_SEPARATOR)
}

func (s *AttendeesContext) LIST_SEPARATOR(i int) antlr.TerminalNode {
	return s.GetToken(CalscriptParserLIST_SEPARATOR, i)
}

func (s *AttendeesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttendeesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttendeesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalscriptListener); ok {
		listenerT.EnterAttendees(s)
	}
}

func (s *AttendeesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalscriptListener); ok {
		listenerT.ExitAttendees(s)
	}
}

func (p *CalscriptParser) Attendees() (localctx IAttendeesContext) {
	localctx = NewAttendeesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, CalscriptParserRULE_attendees)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(124)
		p.Match(CalscriptParserWITH)
	}
	{
		p.SetState(125)
		p.Match(CalscriptParserCOLON)
	}
	{
		p.SetState(126)
		p.Attendee()
	}
	p.SetState(129)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == CalscriptParserLIST_SEPARATOR {
		{
			p.SetState(127)
			p.Match(CalscriptParserLIST_SEPARATOR)
		}
		{
			p.SetState(128)
			p.Attendee()
		}

		p.SetState(131)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IAttendeeContext is an interface to support dynamic dispatch.
type IAttendeeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAttendeeContext differentiates from other interfaces.
	IsAttendeeContext()
}

type AttendeeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttendeeContext() *AttendeeContext {
	var p = new(AttendeeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalscriptParserRULE_attendee
	return p
}

func (*AttendeeContext) IsAttendeeContext() {}

func NewAttendeeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttendeeContext {
	var p = new(AttendeeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalscriptParserRULE_attendee

	return p
}

func (s *AttendeeContext) GetParser() antlr.Parser { return s.parser }

func (s *AttendeeContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CalscriptParserIDENTIFIER, 0)
}

func (s *AttendeeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttendeeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttendeeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalscriptListener); ok {
		listenerT.EnterAttendee(s)
	}
}

func (s *AttendeeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalscriptListener); ok {
		listenerT.ExitAttendee(s)
	}
}

func (p *CalscriptParser) Attendee() (localctx IAttendeeContext) {
	localctx = NewAttendeeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, CalscriptParserRULE_attendee)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(133)
		p.Match(CalscriptParserIDENTIFIER)
	}

	return localctx
}

// ILocationContext is an interface to support dynamic dispatch.
type ILocationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLocationContext differentiates from other interfaces.
	IsLocationContext()
}

type LocationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLocationContext() *LocationContext {
	var p = new(LocationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalscriptParserRULE_location
	return p
}

func (*LocationContext) IsLocationContext() {}

func NewLocationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LocationContext {
	var p = new(LocationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalscriptParserRULE_location

	return p
}

func (s *LocationContext) GetParser() antlr.Parser { return s.parser }

func (s *LocationContext) AT() antlr.TerminalNode {
	return s.GetToken(CalscriptParserAT, 0)
}

func (s *LocationContext) COLON() antlr.TerminalNode {
	return s.GetToken(CalscriptParserCOLON, 0)
}

func (s *LocationContext) Location_string() ILocation_stringContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILocation_stringContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILocation_stringContext)
}

func (s *LocationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LocationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LocationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalscriptListener); ok {
		listenerT.EnterLocation(s)
	}
}

func (s *LocationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalscriptListener); ok {
		listenerT.ExitLocation(s)
	}
}

func (p *CalscriptParser) Location() (localctx ILocationContext) {
	localctx = NewLocationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, CalscriptParserRULE_location)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(135)
		p.Match(CalscriptParserAT)
	}
	{
		p.SetState(136)
		p.Match(CalscriptParserCOLON)
	}
	{
		p.SetState(137)
		p.Location_string()
	}

	return localctx
}

// ILocation_stringContext is an interface to support dynamic dispatch.
type ILocation_stringContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLocation_stringContext differentiates from other interfaces.
	IsLocation_stringContext()
}

type Location_stringContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLocation_stringContext() *Location_stringContext {
	var p = new(Location_stringContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalscriptParserRULE_location_string
	return p
}

func (*Location_stringContext) IsLocation_stringContext() {}

func NewLocation_stringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Location_stringContext {
	var p = new(Location_stringContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalscriptParserRULE_location_string

	return p
}

func (s *Location_stringContext) GetParser() antlr.Parser { return s.parser }

func (s *Location_stringContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CalscriptParserIDENTIFIER, 0)
}

func (s *Location_stringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Location_stringContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Location_stringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalscriptListener); ok {
		listenerT.EnterLocation_string(s)
	}
}

func (s *Location_stringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalscriptListener); ok {
		listenerT.ExitLocation_string(s)
	}
}

func (p *CalscriptParser) Location_string() (localctx ILocation_stringContext) {
	localctx = NewLocation_stringContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, CalscriptParserRULE_location_string)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(139)
		p.Match(CalscriptParserIDENTIFIER)
	}

	return localctx
}

// IDateContext is an interface to support dynamic dispatch.
type IDateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDateContext differentiates from other interfaces.
	IsDateContext()
}

type DateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDateContext() *DateContext {
	var p = new(DateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalscriptParserRULE_date
	return p
}

func (*DateContext) IsDateContext() {}

func NewDateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DateContext {
	var p = new(DateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalscriptParserRULE_date

	return p
}

func (s *DateContext) GetParser() antlr.Parser { return s.parser }

func (s *DateContext) Natural_date() INatural_dateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INatural_dateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INatural_dateContext)
}

func (s *DateContext) Formal_date() IFormal_dateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFormal_dateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFormal_dateContext)
}

func (s *DateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalscriptListener); ok {
		listenerT.EnterDate(s)
	}
}

func (s *DateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalscriptListener); ok {
		listenerT.ExitDate(s)
	}
}

func (p *CalscriptParser) Date() (localctx IDateContext) {
	localctx = NewDateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, CalscriptParserRULE_date)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(143)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CalscriptParserYESTERDAY, CalscriptParserNOW, CalscriptParserTOMORROW, CalscriptParserTODAY:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(141)
			p.Natural_date()
		}

	case CalscriptParserUNSIGNED_INT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(142)
			p.Formal_date()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFormal_date_timeContext is an interface to support dynamic dispatch.
type IFormal_date_timeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFormal_date_timeContext differentiates from other interfaces.
	IsFormal_date_timeContext()
}

type Formal_date_timeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFormal_date_timeContext() *Formal_date_timeContext {
	var p = new(Formal_date_timeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalscriptParserRULE_formal_date_time
	return p
}

func (*Formal_date_timeContext) IsFormal_date_timeContext() {}

func NewFormal_date_timeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Formal_date_timeContext {
	var p = new(Formal_date_timeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalscriptParserRULE_formal_date_time

	return p
}

func (s *Formal_date_timeContext) GetParser() antlr.Parser { return s.parser }

func (s *Formal_date_timeContext) Formal_date() IFormal_dateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFormal_dateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFormal_dateContext)
}

func (s *Formal_date_timeContext) Time() ITimeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITimeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITimeContext)
}

func (s *Formal_date_timeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Formal_date_timeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Formal_date_timeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalscriptListener); ok {
		listenerT.EnterFormal_date_time(s)
	}
}

func (s *Formal_date_timeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalscriptListener); ok {
		listenerT.ExitFormal_date_time(s)
	}
}

func (p *CalscriptParser) Formal_date_time() (localctx IFormal_date_timeContext) {
	localctx = NewFormal_date_timeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, CalscriptParserRULE_formal_date_time)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(145)
		p.Formal_date()
	}
	p.SetState(147)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<CalscriptParserT__0)|(1<<CalscriptParserWHITESPACE)|(1<<CalscriptParserCOMMENT)|(1<<CalscriptParserLINE_COMMENT)|(1<<CalscriptParserLIST_SEPARATOR)|(1<<CalscriptParserON)|(1<<CalscriptParserSTARTING)|(1<<CalscriptParserUNTIL)|(1<<CalscriptParserTHEN)|(1<<CalscriptParserFOR)|(1<<CalscriptParserEVERY)|(1<<CalscriptParserWITH)|(1<<CalscriptParserAT)|(1<<CalscriptParserDAY)|(1<<CalscriptParserWEEK)|(1<<CalscriptParserMONTH)|(1<<CalscriptParserYEAR)|(1<<CalscriptParserHOUR)|(1<<CalscriptParserMINUTE)|(1<<CalscriptParserSECOND)|(1<<CalscriptParserYESTERDAY)|(1<<CalscriptParserNOW)|(1<<CalscriptParserTOMORROW)|(1<<CalscriptParserTODAY)|(1<<CalscriptParserJANUARY)|(1<<CalscriptParserFEBRUARY)|(1<<CalscriptParserMARCH)|(1<<CalscriptParserAPRIL)|(1<<CalscriptParserMAY)|(1<<CalscriptParserJUNE)|(1<<CalscriptParserJULY))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(CalscriptParserAUGUST-32))|(1<<(CalscriptParserSEPTEMBER-32))|(1<<(CalscriptParserOCTOBER-32))|(1<<(CalscriptParserNOVEMBER-32))|(1<<(CalscriptParserDECEMBER-32))|(1<<(CalscriptParserA-32))|(1<<(CalscriptParserB-32))|(1<<(CalscriptParserC-32))|(1<<(CalscriptParserD-32))|(1<<(CalscriptParserE-32))|(1<<(CalscriptParserF-32))|(1<<(CalscriptParserG-32))|(1<<(CalscriptParserH-32))|(1<<(CalscriptParserI-32))|(1<<(CalscriptParserJ-32))|(1<<(CalscriptParserK-32))|(1<<(CalscriptParserL-32))|(1<<(CalscriptParserM-32))|(1<<(CalscriptParserN-32))|(1<<(CalscriptParserO-32))|(1<<(CalscriptParserP-32))|(1<<(CalscriptParserQ-32))|(1<<(CalscriptParserR-32))|(1<<(CalscriptParserS-32))|(1<<(CalscriptParserT-32))|(1<<(CalscriptParserU-32))|(1<<(CalscriptParserV-32))|(1<<(CalscriptParserW-32))|(1<<(CalscriptParserX-32))|(1<<(CalscriptParserY-32))|(1<<(CalscriptParserZ-32))|(1<<(CalscriptParserCOMMA-32)))) != 0) || (((_la-64)&-(0x1f+1)) == 0 && ((1<<uint((_la-64)))&((1<<(CalscriptParserCOLON-64))|(1<<(CalscriptParserSLASH-64))|(1<<(CalscriptParserDASH-64))|(1<<(CalscriptParserDOT-64))|(1<<(CalscriptParserIDENTIFIER-64))|(1<<(CalscriptParserSTRING-64))|(1<<(CalscriptParserUNSIGNED_INT-64))|(1<<(CalscriptParserSIGNED_INT-64))|(1<<(CalscriptParserFLOAT-64)))) != 0) {
		p.SetState(146)
		p.MatchWildcard()

	}
	{
		p.SetState(149)
		p.Time()
	}

	return localctx
}

// IFormal_dateContext is an interface to support dynamic dispatch.
type IFormal_dateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFormal_dateContext differentiates from other interfaces.
	IsFormal_dateContext()
}

type Formal_dateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFormal_dateContext() *Formal_dateContext {
	var p = new(Formal_dateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalscriptParserRULE_formal_date
	return p
}

func (*Formal_dateContext) IsFormal_dateContext() {}

func NewFormal_dateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Formal_dateContext {
	var p = new(Formal_dateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalscriptParserRULE_formal_date

	return p
}

func (s *Formal_dateContext) GetParser() antlr.Parser { return s.parser }

func (s *Formal_dateContext) Year_month_day() IYear_month_dayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IYear_month_dayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IYear_month_dayContext)
}

func (s *Formal_dateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Formal_dateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Formal_dateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalscriptListener); ok {
		listenerT.EnterFormal_date(s)
	}
}

func (s *Formal_dateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalscriptListener); ok {
		listenerT.ExitFormal_date(s)
	}
}

func (p *CalscriptParser) Formal_date() (localctx IFormal_dateContext) {
	localctx = NewFormal_dateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, CalscriptParserRULE_formal_date)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(151)
		p.Year_month_day()
	}

	return localctx
}

// IYear_month_dayContext is an interface to support dynamic dispatch.
type IYear_month_dayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsYear_month_dayContext differentiates from other interfaces.
	IsYear_month_dayContext()
}

type Year_month_dayContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyYear_month_dayContext() *Year_month_dayContext {
	var p = new(Year_month_dayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalscriptParserRULE_year_month_day
	return p
}

func (*Year_month_dayContext) IsYear_month_dayContext() {}

func NewYear_month_dayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Year_month_dayContext {
	var p = new(Year_month_dayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalscriptParserRULE_year_month_day

	return p
}

func (s *Year_month_dayContext) GetParser() antlr.Parser { return s.parser }

func (s *Year_month_dayContext) Year() IYearContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IYearContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IYearContext)
}

func (s *Year_month_dayContext) Month_day() IMonth_dayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMonth_dayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMonth_dayContext)
}

func (s *Year_month_dayContext) Date_separator() IDate_separatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDate_separatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDate_separatorContext)
}

func (s *Year_month_dayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Year_month_dayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Year_month_dayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalscriptListener); ok {
		listenerT.EnterYear_month_day(s)
	}
}

func (s *Year_month_dayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalscriptListener); ok {
		listenerT.ExitYear_month_day(s)
	}
}

func (p *CalscriptParser) Year_month_day() (localctx IYear_month_dayContext) {
	localctx = NewYear_month_dayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, CalscriptParserRULE_year_month_day)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(153)
		p.Year()
	}
	p.SetState(155)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la-63)&-(0x1f+1)) == 0 && ((1<<uint((_la-63)))&((1<<(CalscriptParserCOMMA-63))|(1<<(CalscriptParserCOLON-63))|(1<<(CalscriptParserSLASH-63))|(1<<(CalscriptParserDASH-63))|(1<<(CalscriptParserDOT-63)))) != 0 {
		{
			p.SetState(154)
			p.Date_separator()
		}

	}
	{
		p.SetState(157)
		p.Month_day()
	}

	return localctx
}

// IMonth_dayContext is an interface to support dynamic dispatch.
type IMonth_dayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMonth_dayContext differentiates from other interfaces.
	IsMonth_dayContext()
}

type Month_dayContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMonth_dayContext() *Month_dayContext {
	var p = new(Month_dayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalscriptParserRULE_month_day
	return p
}

func (*Month_dayContext) IsMonth_dayContext() {}

func NewMonth_dayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Month_dayContext {
	var p = new(Month_dayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalscriptParserRULE_month_day

	return p
}

func (s *Month_dayContext) GetParser() antlr.Parser { return s.parser }

func (s *Month_dayContext) Month() IMonthContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMonthContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMonthContext)
}

func (s *Month_dayContext) Day_of_month() IDay_of_monthContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDay_of_monthContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDay_of_monthContext)
}

func (s *Month_dayContext) Date_separator() IDate_separatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDate_separatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDate_separatorContext)
}

func (s *Month_dayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Month_dayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Month_dayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalscriptListener); ok {
		listenerT.EnterMonth_day(s)
	}
}

func (s *Month_dayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalscriptListener); ok {
		listenerT.ExitMonth_day(s)
	}
}

func (p *CalscriptParser) Month_day() (localctx IMonth_dayContext) {
	localctx = NewMonth_dayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, CalscriptParserRULE_month_day)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(159)
		p.Month()
	}
	p.SetState(161)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la-63)&-(0x1f+1)) == 0 && ((1<<uint((_la-63)))&((1<<(CalscriptParserCOMMA-63))|(1<<(CalscriptParserCOLON-63))|(1<<(CalscriptParserSLASH-63))|(1<<(CalscriptParserDASH-63))|(1<<(CalscriptParserDOT-63)))) != 0 {
		{
			p.SetState(160)
			p.Date_separator()
		}

	}
	{
		p.SetState(163)
		p.Day_of_month()
	}

	return localctx
}

// IYearContext is an interface to support dynamic dispatch.
type IYearContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsYearContext differentiates from other interfaces.
	IsYearContext()
}

type YearContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyYearContext() *YearContext {
	var p = new(YearContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalscriptParserRULE_year
	return p
}

func (*YearContext) IsYearContext() {}

func NewYearContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *YearContext {
	var p = new(YearContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalscriptParserRULE_year

	return p
}

func (s *YearContext) GetParser() antlr.Parser { return s.parser }

func (s *YearContext) AllUNSIGNED_INT() []antlr.TerminalNode {
	return s.GetTokens(CalscriptParserUNSIGNED_INT)
}

func (s *YearContext) UNSIGNED_INT(i int) antlr.TerminalNode {
	return s.GetToken(CalscriptParserUNSIGNED_INT, i)
}

func (s *YearContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *YearContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *YearContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalscriptListener); ok {
		listenerT.EnterYear(s)
	}
}

func (s *YearContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalscriptListener); ok {
		listenerT.ExitYear(s)
	}
}

func (p *CalscriptParser) Year() (localctx IYearContext) {
	localctx = NewYearContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, CalscriptParserRULE_year)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(166)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(165)
				p.Match(CalscriptParserUNSIGNED_INT)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(168)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext())
	}

	return localctx
}

// IMonthContext is an interface to support dynamic dispatch.
type IMonthContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMonthContext differentiates from other interfaces.
	IsMonthContext()
}

type MonthContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMonthContext() *MonthContext {
	var p = new(MonthContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalscriptParserRULE_month
	return p
}

func (*MonthContext) IsMonthContext() {}

func NewMonthContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MonthContext {
	var p = new(MonthContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalscriptParserRULE_month

	return p
}

func (s *MonthContext) GetParser() antlr.Parser { return s.parser }

func (s *MonthContext) Vocal_month() IVocal_monthContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVocal_monthContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVocal_monthContext)
}

func (s *MonthContext) Numeric_month() INumeric_monthContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumeric_monthContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumeric_monthContext)
}

func (s *MonthContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MonthContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MonthContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalscriptListener); ok {
		listenerT.EnterMonth(s)
	}
}

func (s *MonthContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalscriptListener); ok {
		listenerT.ExitMonth(s)
	}
}

func (p *CalscriptParser) Month() (localctx IMonthContext) {
	localctx = NewMonthContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, CalscriptParserRULE_month)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(172)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CalscriptParserJANUARY, CalscriptParserFEBRUARY, CalscriptParserMARCH, CalscriptParserAPRIL, CalscriptParserMAY, CalscriptParserJUNE, CalscriptParserJULY, CalscriptParserAUGUST, CalscriptParserSEPTEMBER, CalscriptParserOCTOBER, CalscriptParserNOVEMBER, CalscriptParserDECEMBER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(170)
			p.Vocal_month()
		}

	case CalscriptParserUNSIGNED_INT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(171)
			p.Numeric_month()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IVocal_monthContext is an interface to support dynamic dispatch.
type IVocal_monthContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVocal_monthContext differentiates from other interfaces.
	IsVocal_monthContext()
}

type Vocal_monthContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVocal_monthContext() *Vocal_monthContext {
	var p = new(Vocal_monthContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalscriptParserRULE_vocal_month
	return p
}

func (*Vocal_monthContext) IsVocal_monthContext() {}

func NewVocal_monthContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Vocal_monthContext {
	var p = new(Vocal_monthContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalscriptParserRULE_vocal_month

	return p
}

func (s *Vocal_monthContext) GetParser() antlr.Parser { return s.parser }

func (s *Vocal_monthContext) JANUARY() antlr.TerminalNode {
	return s.GetToken(CalscriptParserJANUARY, 0)
}

func (s *Vocal_monthContext) FEBRUARY() antlr.TerminalNode {
	return s.GetToken(CalscriptParserFEBRUARY, 0)
}

func (s *Vocal_monthContext) MARCH() antlr.TerminalNode {
	return s.GetToken(CalscriptParserMARCH, 0)
}

func (s *Vocal_monthContext) APRIL() antlr.TerminalNode {
	return s.GetToken(CalscriptParserAPRIL, 0)
}

func (s *Vocal_monthContext) MAY() antlr.TerminalNode {
	return s.GetToken(CalscriptParserMAY, 0)
}

func (s *Vocal_monthContext) JUNE() antlr.TerminalNode {
	return s.GetToken(CalscriptParserJUNE, 0)
}

func (s *Vocal_monthContext) JULY() antlr.TerminalNode {
	return s.GetToken(CalscriptParserJULY, 0)
}

func (s *Vocal_monthContext) AUGUST() antlr.TerminalNode {
	return s.GetToken(CalscriptParserAUGUST, 0)
}

func (s *Vocal_monthContext) SEPTEMBER() antlr.TerminalNode {
	return s.GetToken(CalscriptParserSEPTEMBER, 0)
}

func (s *Vocal_monthContext) OCTOBER() antlr.TerminalNode {
	return s.GetToken(CalscriptParserOCTOBER, 0)
}

func (s *Vocal_monthContext) NOVEMBER() antlr.TerminalNode {
	return s.GetToken(CalscriptParserNOVEMBER, 0)
}

func (s *Vocal_monthContext) DECEMBER() antlr.TerminalNode {
	return s.GetToken(CalscriptParserDECEMBER, 0)
}

func (s *Vocal_monthContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Vocal_monthContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Vocal_monthContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalscriptListener); ok {
		listenerT.EnterVocal_month(s)
	}
}

func (s *Vocal_monthContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalscriptListener); ok {
		listenerT.ExitVocal_month(s)
	}
}

func (p *CalscriptParser) Vocal_month() (localctx IVocal_monthContext) {
	localctx = NewVocal_monthContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, CalscriptParserRULE_vocal_month)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(174)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-25)&-(0x1f+1)) == 0 && ((1<<uint((_la-25)))&((1<<(CalscriptParserJANUARY-25))|(1<<(CalscriptParserFEBRUARY-25))|(1<<(CalscriptParserMARCH-25))|(1<<(CalscriptParserAPRIL-25))|(1<<(CalscriptParserMAY-25))|(1<<(CalscriptParserJUNE-25))|(1<<(CalscriptParserJULY-25))|(1<<(CalscriptParserAUGUST-25))|(1<<(CalscriptParserSEPTEMBER-25))|(1<<(CalscriptParserOCTOBER-25))|(1<<(CalscriptParserNOVEMBER-25))|(1<<(CalscriptParserDECEMBER-25)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// INumeric_monthContext is an interface to support dynamic dispatch.
type INumeric_monthContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumeric_monthContext differentiates from other interfaces.
	IsNumeric_monthContext()
}

type Numeric_monthContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumeric_monthContext() *Numeric_monthContext {
	var p = new(Numeric_monthContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalscriptParserRULE_numeric_month
	return p
}

func (*Numeric_monthContext) IsNumeric_monthContext() {}

func NewNumeric_monthContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Numeric_monthContext {
	var p = new(Numeric_monthContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalscriptParserRULE_numeric_month

	return p
}

func (s *Numeric_monthContext) GetParser() antlr.Parser { return s.parser }

func (s *Numeric_monthContext) AllUNSIGNED_INT() []antlr.TerminalNode {
	return s.GetTokens(CalscriptParserUNSIGNED_INT)
}

func (s *Numeric_monthContext) UNSIGNED_INT(i int) antlr.TerminalNode {
	return s.GetToken(CalscriptParserUNSIGNED_INT, i)
}

func (s *Numeric_monthContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Numeric_monthContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Numeric_monthContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalscriptListener); ok {
		listenerT.EnterNumeric_month(s)
	}
}

func (s *Numeric_monthContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalscriptListener); ok {
		listenerT.ExitNumeric_month(s)
	}
}

func (p *CalscriptParser) Numeric_month() (localctx INumeric_monthContext) {
	localctx = NewNumeric_monthContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, CalscriptParserRULE_numeric_month)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(177)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(176)
				p.Match(CalscriptParserUNSIGNED_INT)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(179)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext())
	}

	return localctx
}

// IDay_of_monthContext is an interface to support dynamic dispatch.
type IDay_of_monthContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDay_of_monthContext differentiates from other interfaces.
	IsDay_of_monthContext()
}

type Day_of_monthContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDay_of_monthContext() *Day_of_monthContext {
	var p = new(Day_of_monthContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalscriptParserRULE_day_of_month
	return p
}

func (*Day_of_monthContext) IsDay_of_monthContext() {}

func NewDay_of_monthContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Day_of_monthContext {
	var p = new(Day_of_monthContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalscriptParserRULE_day_of_month

	return p
}

func (s *Day_of_monthContext) GetParser() antlr.Parser { return s.parser }

func (s *Day_of_monthContext) AllUNSIGNED_INT() []antlr.TerminalNode {
	return s.GetTokens(CalscriptParserUNSIGNED_INT)
}

func (s *Day_of_monthContext) UNSIGNED_INT(i int) antlr.TerminalNode {
	return s.GetToken(CalscriptParserUNSIGNED_INT, i)
}

func (s *Day_of_monthContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Day_of_monthContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Day_of_monthContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalscriptListener); ok {
		listenerT.EnterDay_of_month(s)
	}
}

func (s *Day_of_monthContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalscriptListener); ok {
		listenerT.ExitDay_of_month(s)
	}
}

func (p *CalscriptParser) Day_of_month() (localctx IDay_of_monthContext) {
	localctx = NewDay_of_monthContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, CalscriptParserRULE_day_of_month)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(182)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(181)
				p.Match(CalscriptParserUNSIGNED_INT)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(184)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext())
	}

	return localctx
}

// IDate_separatorContext is an interface to support dynamic dispatch.
type IDate_separatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDate_separatorContext differentiates from other interfaces.
	IsDate_separatorContext()
}

type Date_separatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDate_separatorContext() *Date_separatorContext {
	var p = new(Date_separatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalscriptParserRULE_date_separator
	return p
}

func (*Date_separatorContext) IsDate_separatorContext() {}

func NewDate_separatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Date_separatorContext {
	var p = new(Date_separatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalscriptParserRULE_date_separator

	return p
}

func (s *Date_separatorContext) GetParser() antlr.Parser { return s.parser }

func (s *Date_separatorContext) COMMA() antlr.TerminalNode {
	return s.GetToken(CalscriptParserCOMMA, 0)
}

func (s *Date_separatorContext) DOT() antlr.TerminalNode {
	return s.GetToken(CalscriptParserDOT, 0)
}

func (s *Date_separatorContext) COLON() antlr.TerminalNode {
	return s.GetToken(CalscriptParserCOLON, 0)
}

func (s *Date_separatorContext) SLASH() antlr.TerminalNode {
	return s.GetToken(CalscriptParserSLASH, 0)
}

func (s *Date_separatorContext) DASH() antlr.TerminalNode {
	return s.GetToken(CalscriptParserDASH, 0)
}

func (s *Date_separatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Date_separatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Date_separatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalscriptListener); ok {
		listenerT.EnterDate_separator(s)
	}
}

func (s *Date_separatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalscriptListener); ok {
		listenerT.ExitDate_separator(s)
	}
}

func (p *CalscriptParser) Date_separator() (localctx IDate_separatorContext) {
	localctx = NewDate_separatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, CalscriptParserRULE_date_separator)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(186)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-63)&-(0x1f+1)) == 0 && ((1<<uint((_la-63)))&((1<<(CalscriptParserCOMMA-63))|(1<<(CalscriptParserCOLON-63))|(1<<(CalscriptParserSLASH-63))|(1<<(CalscriptParserDASH-63))|(1<<(CalscriptParserDOT-63)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ITimeContext is an interface to support dynamic dispatch.
type ITimeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTimeContext differentiates from other interfaces.
	IsTimeContext()
}

type TimeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTimeContext() *TimeContext {
	var p = new(TimeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalscriptParserRULE_time
	return p
}

func (*TimeContext) IsTimeContext() {}

func NewTimeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TimeContext {
	var p = new(TimeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalscriptParserRULE_time

	return p
}

func (s *TimeContext) GetParser() antlr.Parser { return s.parser }
func (s *TimeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TimeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalscriptListener); ok {
		listenerT.EnterTime(s)
	}
}

func (s *TimeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalscriptListener); ok {
		listenerT.ExitTime(s)
	}
}

func (p *CalscriptParser) Time() (localctx ITimeContext) {
	localctx = NewTimeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, CalscriptParserRULE_time)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)

	return localctx
}

// INatural_dateContext is an interface to support dynamic dispatch.
type INatural_dateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNatural_dateContext differentiates from other interfaces.
	IsNatural_dateContext()
}

type Natural_dateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNatural_dateContext() *Natural_dateContext {
	var p = new(Natural_dateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalscriptParserRULE_natural_date
	return p
}

func (*Natural_dateContext) IsNatural_dateContext() {}

func NewNatural_dateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Natural_dateContext {
	var p = new(Natural_dateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalscriptParserRULE_natural_date

	return p
}

func (s *Natural_dateContext) GetParser() antlr.Parser { return s.parser }

func (s *Natural_dateContext) YESTERDAY() antlr.TerminalNode {
	return s.GetToken(CalscriptParserYESTERDAY, 0)
}

func (s *Natural_dateContext) NOW() antlr.TerminalNode {
	return s.GetToken(CalscriptParserNOW, 0)
}

func (s *Natural_dateContext) TOMORROW() antlr.TerminalNode {
	return s.GetToken(CalscriptParserTOMORROW, 0)
}

func (s *Natural_dateContext) TODAY() antlr.TerminalNode {
	return s.GetToken(CalscriptParserTODAY, 0)
}

func (s *Natural_dateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Natural_dateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Natural_dateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalscriptListener); ok {
		listenerT.EnterNatural_date(s)
	}
}

func (s *Natural_dateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalscriptListener); ok {
		listenerT.ExitNatural_date(s)
	}
}

func (p *CalscriptParser) Natural_date() (localctx INatural_dateContext) {
	localctx = NewNatural_dateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, CalscriptParserRULE_natural_date)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(190)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<CalscriptParserYESTERDAY)|(1<<CalscriptParserNOW)|(1<<CalscriptParserTOMORROW)|(1<<CalscriptParserTODAY))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
