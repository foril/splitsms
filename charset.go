package splitsms

/**
* Go library for split SMS.
* This library support SMS in GSM 7, basic extended GSM 7 table and Unicode charset.
* The size of UDH can be defined for concatened SMS.
* https://en.wikipedia.org/wiki/GSM_03.38#GSM_7-bit_default_alphabet_and_extension_table_of_3GPP_TS_23.038_.2F_GSM_03.38
 */

// Basic Character Set
var gsm7Chars = map[rune]string{
	10:   "\n",
	12:   "\f",
	13:   "\r",
	32:   " ",
	33:   "!",
	34:   "\"",
	35:   "#",
	36:   "$",
	37:   "%",
	38:   "&",
	39:   "'",
	40:   "(",
	41:   ")",
	42:   "*",
	43:   "+",
	44:   ",",
	45:   "-",
	46:   ".",
	47:   "/",
	48:   "0",
	49:   "1",
	50:   "2",
	51:   "3",
	52:   "4",
	53:   "5",
	54:   "6",
	55:   "7",
	56:   "8",
	57:   "9",
	58:   ":",
	59:   ";",
	60:   "<",
	61:   "=",
	62:   ">",
	63:   "?",
	64:   "@",
	65:   "A",
	66:   "B",
	67:   "C",
	68:   "D",
	69:   "E",
	70:   "F",
	71:   "G",
	72:   "H",
	73:   "I",
	74:   "J",
	75:   "K",
	76:   "L",
	77:   "M",
	78:   "N",
	79:   "O",
	80:   "P",
	81:   "Q",
	82:   "R",
	83:   "S",
	84:   "T",
	85:   "U",
	86:   "V",
	87:   "W",
	88:   "X",
	89:   "Y",
	90:   "Z",
	91:   "[",
	92:   "\\",
	93:   "]",
	94:   "^",
	95:   "_",
	97:   "a",
	98:   "b",
	99:   "c",
	100:  "d",
	101:  "e",
	102:  "f",
	103:  "g",
	104:  "h",
	105:  "i",
	106:  "j",
	107:  "k",
	108:  "l",
	109:  "m",
	110:  "n",
	111:  "o",
	112:  "p",
	113:  "q",
	114:  "r",
	115:  "s",
	116:  "t",
	117:  "u",
	118:  "v",
	119:  "w",
	120:  "x",
	121:  "y",
	122:  "z",
	123:  "{",
	124:  "|",
	125:  "}",
	126:  "~",
	161:  "¡",
	163:  "£",
	164:  "¤",
	165:  "¥",
	167:  "§",
	191:  "¿",
	196:  "Ä",
	197:  "Å",
	198:  "Æ",
	199:  "Ç",
	201:  "É",
	209:  "Ñ",
	214:  "Ö",
	216:  "Ø",
	220:  "Ü",
	223:  "ß",
	224:  "à",
	228:  "ä",
	229:  "å",
	230:  "æ",
	232:  "è",
	233:  "é",
	236:  "ì",
	241:  "ñ",
	242:  "ò",
	246:  "ö",
	248:  "ø",
	249:  "ù",
	252:  "ü",
	915:  "Γ",
	916:  "Δ",
	920:  "Θ",
	923:  "Λ",
	926:  "Ξ",
	928:  "Π",
	931:  "Σ",
	934:  "Φ",
	936:  "Ψ",
	937:  "Ω",
	8364: "€",
}

// Basic Character Set Extension
var gsm7ExtChars = map[rune]string{
	12:   "\f",
	91:   "[",
	92:   "\\",
	93:   "]",
	94:   "^",
	123:  "{",
	124:  "|",
	125:  "}",
	126:  "~",
	8364: "€",
}