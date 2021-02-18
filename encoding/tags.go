/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/tags.go                                         |
|                                                          |
| LastModified: Feb 18, 2021                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

// Hprose Tags.
const (
	// Serialize Type.
	TagInteger  byte = 'i'
	TagLong     byte = 'l'
	TagDouble   byte = 'd'
	TagNull     byte = 'n'
	TagEmpty    byte = 'e'
	TagTrue     byte = 't'
	TagFalse    byte = 'f'
	TagNaN      byte = 'N'
	TagInfinity byte = 'I'
	TagDate     byte = 'D'
	TagTime     byte = 'T'
	TagUTC      byte = 'Z'
	TagBytes    byte = 'b'
	TagUTF8Char byte = 'u'
	TagString   byte = 's'
	TagGUID     byte = 'g'
	TagList     byte = 'a'
	TagMap      byte = 'm'
	TagClass    byte = 'c'
	TagObject   byte = 'o'
	TagRef      byte = 'r'

	// Serialize Marks.
	TagPos        byte = '+'
	TagNeg        byte = '-'
	TagSemicolon  byte = ';'
	TagOpenbrace  byte = '{'
	TagClosebrace byte = '}'
	TagQuote      byte = '"'
	TagPoint      byte = '.'

	// Protocol Tags.
	TagHeader byte = 'H'
	TagCall   byte = 'C'
	TagResult byte = 'R'
	TagError  byte = 'E'
	TagEnd    byte = 'z'
)
