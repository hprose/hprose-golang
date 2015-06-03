/**********************************************************\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: http://www.hprose.com/                 |
|                   http://www.hprose.org/                 |
|                                                          |
\**********************************************************/
/**********************************************************\
 *                                                        *
 * hprose/reader.go                                       *
 *                                                        *
 * hprose Reader for Go.                                  *
 *                                                        *
 * LastModified: Jun 3, 2015                              *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package hprose

import (
	"container/list"
	"errors"
	"math"
	"math/big"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"
	"unicode/utf8"
)

// ErrNil is a error of nil
var ErrNil = errors.New("nil")

var errBadEncode = errors.New("bad utf-8 encoding")
var errRef = unexpectedTag(TagRef, nil)

var bigDigit = [...]*big.Int{
	big.NewInt(0),
	big.NewInt(1),
	big.NewInt(2),
	big.NewInt(3),
	big.NewInt(4),
	big.NewInt(5),
	big.NewInt(6),
	big.NewInt(7),
	big.NewInt(8),
	big.NewInt(9),
}
var bigTen = big.NewInt(10)

const timeStringFormat = "2006-01-02 15:04:05.999999999 -0700 MST"

var timeZero = time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC)

var soMapType = reflect.TypeOf(map[string]interface{}(nil))
var ooMapType = reflect.TypeOf(map[interface{}]interface{}(nil))

var indexCache struct {
	sync.RWMutex
	cache map[reflect.Type]map[string][]int
}

// BufReader is buffer reader interface, Hprose Reader use it as input stream.
type BufReader interface {
	Read(p []byte) (n int, err error)
	ReadByte() (c byte, err error)
	ReadRune() (r rune, size int, err error)
	ReadString(delim byte) (line string, err error)
}

type readerRefer interface {
	setRef(p interface{})
	readRef(i int, err error) (interface{}, error)
	resetRef()
}

type fakeReaderRefer struct{}

func (r fakeReaderRefer) setRef(p interface{}) {}

func (r fakeReaderRefer) readRef(i int, err error) (interface{}, error) {
	return nil, errRef
}

func (r fakeReaderRefer) resetRef() {}

type realReaderRefer struct {
	ref []interface{}
}

func (r *realReaderRefer) setRef(p interface{}) {
	if r.ref == nil {
		r.ref = make([]interface{}, 0)
	}
	r.ref = append(r.ref, p)
}

func (r *realReaderRefer) readRef(i int, err error) (interface{}, error) {
	if err == nil {
		return r.ref[i], nil
	}
	return nil, err
}

func (r *realReaderRefer) resetRef() {
	if r.ref != nil {
		r.ref = r.ref[:0]
	}
}

// Reader is a fine-grained operation struct for Hprose unserialization
// when JSONCompatible is true, the Map data will unserialize to map[string]interface as the default type
type Reader struct {
	*RawReader
	classref  []interface{}
	fieldsref [][]string
	readerRefer
	JSONCompatible bool
}

// NewReader is the constructor for Hprose Reader
func NewReader(stream BufReader, simple bool) (reader *Reader) {
	reader = new(Reader)
	reader.RawReader = NewRawReader(stream)
	if simple {
		reader.readerRefer = fakeReaderRefer{}
	} else {
		reader.readerRefer = new(realReaderRefer)
	}
	return
}

// CheckTag the next byte in stream is the expected tag
func (r *Reader) CheckTag(expectTag byte) error {
	tag, err := r.Stream.ReadByte()
	if err == nil {
		return unexpectedTag(tag, []byte{expectTag})
	}
	return err
}

// CheckTags the next byte in stream in the expected tags
func (r *Reader) CheckTags(expectTags []byte) (tag byte, err error) {
	tag, err = r.Stream.ReadByte()
	if err == nil {
		if err = unexpectedTag(tag, expectTags); err == nil {
			return tag, nil
		}
	}
	return 0, err
}

// Unserialize a data from stream
func (r *Reader) Unserialize(p interface{}) (err error) {
	switch p := p.(type) {
	case nil:
		return errors.New("argument p must be non-null pointer")
	case *int:
		if *p, err = r.ReadInt(); err == ErrNil {
			err = nil
		}
		return err
	case *uint:
		if *p, err = r.ReadUint(); err == ErrNil {
			err = nil
		}
		return err
	case *int8:
		if *p, err = r.ReadInt8(); err == ErrNil {
			err = nil
		}
		return err
	case *uint8:
		if *p, err = r.ReadUint8(); err == ErrNil {
			err = nil
		}
		return err
	case *int16:
		if *p, err = r.ReadInt16(); err == ErrNil {
			err = nil
		}
		return err
	case *uint16:
		if *p, err = r.ReadUint16(); err == ErrNil {
			err = nil
		}
		return err
	case *int32:
		if *p, err = r.ReadInt32(); err == ErrNil {
			err = nil
		}
		return err
	case *uint32:
		if *p, err = r.ReadUint32(); err == ErrNil {
			err = nil
		}
		return err
	case *int64:
		if *p, err = r.ReadInt64(); err == ErrNil {
			err = nil
		}
		return err
	case *uint64:
		if *p, err = r.ReadUint64(); err == ErrNil {
			err = nil
		}
		return err
	case *big.Int:
		if x, err := r.ReadBigInt(); err == ErrNil {
			err = nil
			*p = *x
		} else {
			*p = *x
		}
		return err
	case *float32:
		if *p, err = r.ReadFloat32(); err == ErrNil {
			err = nil
		}
		return err
	case *float64:
		if *p, err = r.ReadFloat64(); err == ErrNil {
			err = nil
		}
		return err
	case *bool:
		if *p, err = r.ReadBool(); err == ErrNil {
			err = nil
		}
		return err
	case *time.Time:
		if *p, err = r.ReadDateTime(); err == ErrNil {
			err = nil
		}
		return err
	case *string:
		if *p, err = r.ReadString(); err == ErrNil {
			err = nil
		}
		return err
	case *[]byte:
		if x, err := r.ReadBytes(); err == ErrNil {
			*p = *x
			err = nil
		} else {
			*p = *x
		}
		return err
	case *UUID:
		if x, err := r.ReadUUID(); err == ErrNil {
			*p = *x
			err = nil
		} else {
			*p = *x
		}
		return err
	case *list.List:
		if x, err := r.ReadList(); err == ErrNil {
			*p = *x
			err = nil
		} else {
			*p = *x
		}
		return err
	case **int:
		if x, err := r.ReadInt(); err == ErrNil {
			*p = nil
			err = nil
		} else {
			*p = &x
		}
		return err
	case **uint:
		if x, err := r.ReadUint(); err == ErrNil {
			*p = nil
			err = nil
		} else {
			*p = &x
		}
		return err
	case **int8:
		if x, err := r.ReadInt8(); err == ErrNil {
			*p = nil
			err = nil
		} else {
			*p = &x
		}
		return err
	case **uint8:
		if x, err := r.ReadUint8(); err == ErrNil {
			*p = nil
			err = nil
		} else {
			*p = &x
		}
		return err
	case **int16:
		if x, err := r.ReadInt16(); err == ErrNil {
			*p = nil
			err = nil
		} else {
			*p = &x
		}
		return err
	case **uint16:
		if x, err := r.ReadUint16(); err == ErrNil {
			*p = nil
			err = nil
		} else {
			*p = &x
		}
		return err
	case **int32:
		if x, err := r.ReadInt32(); err == ErrNil {
			*p = nil
			err = nil
		} else {
			*p = &x
		}
		return err
	case **uint32:
		if x, err := r.ReadUint32(); err == ErrNil {
			*p = nil
			err = nil
		} else {
			*p = &x
		}
		return err
	case **int64:
		if x, err := r.ReadInt64(); err == ErrNil {
			*p = nil
			err = nil
		} else {
			*p = &x
		}
		return err
	case **uint64:
		if x, err := r.ReadUint64(); err == ErrNil {
			*p = nil
			err = nil
		} else {
			*p = &x
		}
		return err
	case **big.Int:
		if *p, err = r.ReadBigInt(); err == ErrNil {
			*p = nil
			err = nil
		}
		return err
	case **float32:
		if x, err := r.ReadFloat32(); err == ErrNil {
			*p = nil
			err = nil
		} else {
			*p = &x
		}
		return err
	case **float64:
		if x, err := r.ReadFloat64(); err == ErrNil {
			*p = nil
			err = nil
		} else {
			*p = &x
		}
		return err
	case **bool:
		if x, err := r.ReadBool(); err == ErrNil {
			*p = nil
			err = nil
		} else {
			*p = &x
		}
		return err
	case **time.Time:
		if x, err := r.ReadDateTime(); err == ErrNil {
			*p = nil
			err = nil
		} else {
			*p = &x
		}
		return err
	case **string:
		if x, err := r.ReadString(); err == ErrNil {
			*p = nil
			err = nil
		} else {
			*p = &x
		}
		return err
	case **[]byte:
		if *p, err = r.ReadBytes(); err == ErrNil {
			*p = nil
			err = nil
		}
		return err
	case **UUID:
		if *p, err = r.ReadUUID(); err == ErrNil {
			*p = nil
			err = nil
		}
		return err
	case **list.List:
		if *p, err = r.ReadList(); err == ErrNil {
			*p = nil
			err = nil
		}
		return err
	case *interface{}:
		if *p, err = r.readInterface(); err == ErrNil {
			*p = nil
			err = nil
		}
		return err
	default:
		v, err := r.checkPointer(p)
		if err == nil {
			return r.ReadValue(v.Elem())
		}
	}
	return err
}

// ReadInteger from stream
func (r *Reader) ReadInteger(tag byte) (int, error) {
	s := r.Stream
	i := 0
	b, err := s.ReadByte()
	if err == nil && b == tag {
		return i, nil
	}
	if err != nil {
		return i, err
	}
	sign := 1
	switch b {
	case '-':
		sign = -1
		fallthrough
	case '+':
		b, err = s.ReadByte()
	}
	for b != tag && err == nil {
		i *= 10
		i += int(b-'0') * sign
		b, err = s.ReadByte()
	}
	return i, err
}

// ReadUinteger from stream
func (r *Reader) ReadUinteger(tag byte) (uint, error) {
	i, err := r.ReadInteger(tag)
	return uint(i), err
}

// ReadInt from stream
func (r *Reader) ReadInt() (int, error) {
	s := r.Stream
	tag, err := s.ReadByte()
	if err == nil {
		switch tag {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			return int(tag - '0'), nil
		case TagInteger, TagLong:
			return r.ReadIntWithoutTag()
		case TagDouble:
			f, err := r.ReadFloat64WithoutTag()
			return int(f), err
		case TagNull:
			return 0, ErrNil
		case TagEmpty, TagFalse:
			return 0, nil
		case TagTrue:
			return 1, nil
		case TagUTF8Char:
			var str string
			if str, err = r.readUTF8String(1); err == nil {
				i, err := strconv.ParseInt(str, 10, 64)
				return int(i), err
			}
		case TagString:
			var str string
			if str, err = r.ReadStringWithoutTag(); err == nil {
				i, err := strconv.ParseInt(str, 10, 64)
				return int(i), err
			}
		case TagRef:
			var ref interface{}
			if ref, err = r.readRef(r.ReadInteger(TagSemicolon)); err == nil {
				if ref, ok := ref.(string); ok {
					i, err := strconv.ParseInt(ref, 10, 64)
					return int(i), err
				}
				return 0, errors.New("cannot convert type " +
					reflect.TypeOf(ref).String() + " to type int")
			}
		default:
			return 0, convertError(tag, "int")
		}
	}
	return 0, err
}

// ReadIntWithoutTag from stream
func (r *Reader) ReadIntWithoutTag() (int, error) {
	return r.ReadInteger(TagSemicolon)
}

// ReadUint from stream
func (r *Reader) ReadUint() (uint, error) {
	s := r.Stream
	tag, err := s.ReadByte()
	if err == nil {
		switch tag {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			return uint(tag - '0'), nil
		case TagInteger, TagLong:
			return r.ReadUintWithoutTag()
		case TagDouble:
			f, err := r.ReadFloat64WithoutTag()
			return uint(f), err
		case TagNull:
			return 0, ErrNil
		case TagEmpty, TagFalse:
			return 0, nil
		case TagTrue:
			return 1, nil
		case TagUTF8Char:
			var str string
			if str, err = r.readUTF8String(1); err == nil {
				i, err := strconv.ParseUint(str, 10, 64)
				return uint(i), err
			}
		case TagString:
			var str string
			if str, err = r.ReadStringWithoutTag(); err == nil {
				i, err := strconv.ParseUint(str, 10, 64)
				return uint(i), err
			}
		case TagRef:
			var ref interface{}
			if ref, err = r.readRef(r.ReadInteger(TagSemicolon)); err == nil {
				if ref, ok := ref.(string); ok {
					i, err := strconv.ParseUint(ref, 10, 64)
					return uint(i), err
				}
				return 0, errors.New("cannot convert type " +
					reflect.TypeOf(ref).String() + " to type uint")
			}
		default:
			return 0, convertError(tag, "uint")
		}
	}
	return 0, err
}

// ReadUintWithoutTag from stream
func (r *Reader) ReadUintWithoutTag() (uint, error) {
	return r.ReadUinteger(TagSemicolon)
}

// ReadInt64 from stream
func (r *Reader) ReadInt64() (int64, error) {
	s := r.Stream
	tag, err := s.ReadByte()
	if err == nil {
		switch tag {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			return int64(tag - '0'), nil
		case TagInteger, TagLong:
			return r.ReadInt64WithoutTag()
		case TagDouble:
			f, err := r.ReadFloat64WithoutTag()
			return int64(f), err
		case TagNull:
			return 0, ErrNil
		case TagEmpty, TagFalse:
			return 0, nil
		case TagTrue:
			return 1, nil
		case TagUTF8Char:
			var str string
			if str, err = r.readUTF8String(1); err == nil {
				return strconv.ParseInt(str, 10, 64)
			}
		case TagString:
			var str string
			if str, err = r.ReadStringWithoutTag(); err == nil {
				return strconv.ParseInt(str, 10, 64)
			}
		case TagRef:
			var ref interface{}
			if ref, err = r.readRef(r.ReadInteger(TagSemicolon)); err == nil {
				if ref, ok := ref.(string); ok {
					return strconv.ParseInt(ref, 10, 64)
				}
				return 0, errors.New("cannot convert type " +
					reflect.TypeOf(ref).String() + " to type int64")
			}
		default:
			return 0, convertError(tag, "int64")
		}
	}
	return 0, err
}

// ReadInt64WithoutTag from stream
func (r *Reader) ReadInt64WithoutTag() (int64, error) {
	return r.readInt(TagSemicolon)
}

// ReadUint64 from stream
func (r *Reader) ReadUint64() (uint64, error) {
	s := r.Stream
	tag, err := s.ReadByte()
	if err == nil {
		switch tag {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			return uint64(tag - '0'), nil
		case TagInteger, TagLong:
			return r.ReadUint64WithoutTag()
		case TagDouble:
			f, err := r.ReadFloat64WithoutTag()
			return uint64(f), err
		case TagNull:
			return 0, ErrNil
		case TagEmpty, TagFalse:
			return 0, nil
		case TagTrue:
			return 1, nil
		case TagUTF8Char:
			var str string
			if str, err = r.readUTF8String(1); err == nil {
				return strconv.ParseUint(str, 10, 64)
			}
		case TagString:
			var str string
			if str, err = r.ReadStringWithoutTag(); err == nil {
				return strconv.ParseUint(str, 10, 64)
			}
		case TagRef:
			var ref interface{}
			if ref, err = r.readRef(r.ReadInteger(TagSemicolon)); err == nil {
				if ref, ok := ref.(string); ok {
					return strconv.ParseUint(ref, 10, 64)
				}
				return 0, errors.New("cannot convert type " +
					reflect.TypeOf(ref).String() + " to type uint64")
			}
		default:
			return 0, convertError(tag, "uint64")
		}
	}
	return 0, err
}

// ReadUint64WithoutTag from stream
func (r *Reader) ReadUint64WithoutTag() (uint64, error) {
	return r.readUint(TagSemicolon)
}

// ReadInt8 from stream
func (r *Reader) ReadInt8() (int8, error) {
	s := r.Stream
	tag, err := s.ReadByte()
	if err == nil {
		switch tag {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			return int8(tag - '0'), nil
		case TagInteger, TagLong:
			return r.ReadInt8WithoutTag()
		case TagDouble:
			f, err := r.ReadFloat64WithoutTag()
			return int8(f), err
		case TagNull:
			return 0, ErrNil
		case TagEmpty, TagFalse:
			return 0, nil
		case TagTrue:
			return 1, nil
		case TagUTF8Char:
			var str string
			if str, err = r.readUTF8String(1); err == nil {
				i, err := strconv.ParseInt(str, 10, 64)
				return int8(i), err
			}
		case TagString:
			var str string
			if str, err = r.ReadStringWithoutTag(); err == nil {
				i, err := strconv.ParseInt(str, 10, 64)
				return int8(i), err
			}
		case TagRef:
			var ref interface{}
			if ref, err = r.readRef(r.ReadInteger(TagSemicolon)); err == nil {
				if ref, ok := ref.(string); ok {
					i, err := strconv.ParseInt(ref, 10, 64)
					return int8(i), err
				}
				return 0, errors.New("cannot convert type " +
					reflect.TypeOf(ref).String() + " to type int8")
			}
		default:
			return 0, convertError(tag, "int8")
		}
	}
	return 0, err
}

// ReadInt8WithoutTag from stream
func (r *Reader) ReadInt8WithoutTag() (int8, error) {
	return r.readInt8(TagSemicolon)
}

// ReadUint8 from stream
func (r *Reader) ReadUint8() (uint8, error) {
	s := r.Stream
	tag, err := s.ReadByte()
	if err == nil {
		switch tag {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			return uint8(tag - '0'), nil
		case TagInteger, TagLong:
			return r.ReadUint8WithoutTag()
		case TagDouble:
			f, err := r.ReadFloat64WithoutTag()
			return uint8(f), err
		case TagNull:
			return 0, ErrNil
		case TagEmpty, TagFalse:
			return 0, nil
		case TagTrue:
			return 1, nil
		case TagUTF8Char:
			var str string
			if str, err = r.readUTF8String(1); err == nil {
				i, err := strconv.ParseUint(str, 10, 64)
				return uint8(i), err
			}
		case TagString:
			var str string
			if str, err = r.ReadStringWithoutTag(); err == nil {
				i, err := strconv.ParseUint(str, 10, 64)
				return uint8(i), err
			}
		case TagRef:
			var ref interface{}
			if ref, err = r.readRef(r.ReadInteger(TagSemicolon)); err == nil {
				if ref, ok := ref.(string); ok {
					i, err := strconv.ParseUint(ref, 10, 64)
					return uint8(i), err
				}
				return 0, errors.New("cannot convert type " +
					reflect.TypeOf(ref).String() + " to type uint8")
			}
		default:
			return 0, convertError(tag, "uint8")
		}
	}
	return 0, err
}

// ReadUint8WithoutTag from stream
func (r *Reader) ReadUint8WithoutTag() (uint8, error) {
	return r.readUint8(TagSemicolon)
}

// ReadInt16 from stream
func (r *Reader) ReadInt16() (int16, error) {
	s := r.Stream
	tag, err := s.ReadByte()
	if err == nil {
		switch tag {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			return int16(tag - '0'), nil
		case TagInteger, TagLong:
			return r.ReadInt16WithoutTag()
		case TagDouble:
			f, err := r.ReadFloat64WithoutTag()
			return int16(f), err
		case TagNull:
			return 0, ErrNil
		case TagEmpty, TagFalse:
			return 0, nil
		case TagTrue:
			return 1, nil
		case TagUTF8Char:
			var str string
			if str, err = r.readUTF8String(1); err == nil {
				i, err := strconv.ParseInt(str, 10, 64)
				return int16(i), err
			}
		case TagString:
			var str string
			if str, err = r.ReadStringWithoutTag(); err == nil {
				i, err := strconv.ParseInt(str, 10, 64)
				return int16(i), err
			}
		case TagRef:
			var ref interface{}
			if ref, err = r.readRef(r.ReadInteger(TagSemicolon)); err == nil {
				if ref, ok := ref.(string); ok {
					i, err := strconv.ParseInt(ref, 10, 64)
					return int16(i), err
				}
				return 0, errors.New("cannot convert type " +
					reflect.TypeOf(ref).String() + " to type int16")
			}
		default:
			return 0, convertError(tag, "int16")
		}
	}
	return 0, err
}

// ReadInt16WithoutTag from stream
func (r *Reader) ReadInt16WithoutTag() (int16, error) {
	return r.readInt16(TagSemicolon)
}

// ReadUint16 from stream
func (r *Reader) ReadUint16() (uint16, error) {
	s := r.Stream
	tag, err := s.ReadByte()
	if err == nil {
		switch tag {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			return uint16(tag - '0'), nil
		case TagInteger, TagLong:
			return r.ReadUint16WithoutTag()
		case TagDouble:
			f, err := r.ReadFloat64WithoutTag()
			return uint16(f), err
		case TagNull:
			return 0, ErrNil
		case TagEmpty, TagFalse:
			return 0, nil
		case TagTrue:
			return 1, nil
		case TagUTF8Char:
			var str string
			if str, err = r.readUTF8String(1); err == nil {
				i, err := strconv.ParseUint(str, 10, 64)
				return uint16(i), err
			}
		case TagString:
			var str string
			if str, err = r.ReadStringWithoutTag(); err == nil {
				i, err := strconv.ParseUint(str, 10, 64)
				return uint16(i), err
			}
		case TagRef:
			var ref interface{}
			if ref, err = r.readRef(r.ReadInteger(TagSemicolon)); err == nil {
				if ref, ok := ref.(string); ok {
					i, err := strconv.ParseUint(ref, 10, 64)
					return uint16(i), err
				}
				return 0, errors.New("cannot convert type " +
					reflect.TypeOf(ref).String() + " to type uint16")
			}
		default:
			return 0, convertError(tag, "uint16")
		}
	}
	return 0, err
}

// ReadUint16WithoutTag from stream
func (r *Reader) ReadUint16WithoutTag() (uint16, error) {
	return r.readUint16(TagSemicolon)
}

// ReadInt32 from stream
func (r *Reader) ReadInt32() (int32, error) {
	s := r.Stream
	tag, err := s.ReadByte()
	if err == nil {
		switch tag {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			return int32(tag - '0'), nil
		case TagInteger, TagLong:
			return r.ReadInt32WithoutTag()
		case TagDouble:
			f, err := r.ReadFloat64WithoutTag()
			return int32(f), err
		case TagNull:
			return 0, ErrNil
		case TagEmpty, TagFalse:
			return 0, nil
		case TagTrue:
			return 1, nil
		case TagUTF8Char:
			var str string
			if str, err = r.readUTF8String(1); err == nil {
				i, err := strconv.ParseInt(str, 10, 64)
				return int32(i), err
			}
		case TagString:
			var str string
			if str, err = r.ReadStringWithoutTag(); err == nil {
				i, err := strconv.ParseInt(str, 10, 64)
				return int32(i), err
			}
		case TagRef:
			var ref interface{}
			if ref, err = r.readRef(r.ReadInteger(TagSemicolon)); err == nil {
				if ref, ok := ref.(string); ok {
					i, err := strconv.ParseInt(ref, 10, 64)
					return int32(i), err
				}
				return 0, errors.New("cannot convert type " +
					reflect.TypeOf(ref).String() + " to type int32")
			}
		default:
			return 0, convertError(tag, "int32")
		}
	}
	return 0, err
}

// ReadInt32WithoutTag from stream
func (r *Reader) ReadInt32WithoutTag() (int32, error) {
	return r.readInt32(TagSemicolon)
}

// ReadUint32 from stream
func (r *Reader) ReadUint32() (uint32, error) {
	s := r.Stream
	tag, err := s.ReadByte()
	if err == nil {
		switch tag {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			return uint32(tag - '0'), nil
		case TagInteger, TagLong:
			return r.ReadUint32WithoutTag()
		case TagDouble:
			f, err := r.ReadFloat64WithoutTag()
			return uint32(f), err
		case TagNull:
			return 0, ErrNil
		case TagEmpty, TagFalse:
			return 0, nil
		case TagTrue:
			return 1, nil
		case TagUTF8Char:
			var str string
			if str, err = r.readUTF8String(1); err == nil {
				i, err := strconv.ParseUint(str, 10, 64)
				return uint32(i), err
			}
		case TagString:
			var str string
			if str, err = r.ReadStringWithoutTag(); err == nil {
				i, err := strconv.ParseUint(str, 10, 64)
				return uint32(i), err
			}
		case TagRef:
			var ref interface{}
			if ref, err = r.readRef(r.ReadInteger(TagSemicolon)); err == nil {
				if ref, ok := ref.(string); ok {
					i, err := strconv.ParseUint(ref, 10, 64)
					return uint32(i), err
				}
				return 0, errors.New("cannot convert type " +
					reflect.TypeOf(ref).String() + " to type uint32")
			}
		default:
			return 0, convertError(tag, "uint32")
		}
	}
	return 0, err
}

// ReadUint32WithoutTag from stream
func (r *Reader) ReadUint32WithoutTag() (uint32, error) {
	return r.readUint32(TagSemicolon)
}

// ReadBigInt from stream
func (r *Reader) ReadBigInt() (*big.Int, error) {
	s := r.Stream
	tag, err := s.ReadByte()
	if err == nil {
		switch tag {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			return big.NewInt(int64(tag - '0')), nil
		case TagInteger, TagLong:
			return r.ReadBigIntWithoutTag()
		case TagDouble:
			f, err := r.ReadFloat64WithoutTag()
			return big.NewInt(int64(f)), err
		case TagNull:
			return big.NewInt(0), ErrNil
		case TagEmpty, TagFalse:
			return big.NewInt(0), nil
		case TagTrue:
			return big.NewInt(1), nil
		case TagUTF8Char:
			var str string
			if str, err = r.readUTF8String(1); err == nil {
				return stringToBigInt(str)
			}
		case TagString:
			var str string
			if str, err = r.ReadStringWithoutTag(); err == nil {
				return stringToBigInt(str)
			}
		case TagRef:
			var ref interface{}
			if ref, err = r.readRef(r.ReadInteger(TagSemicolon)); err == nil {
				if ref, ok := ref.(string); ok {
					return stringToBigInt(ref)
				}
				return nil, errors.New("cannot convert type " +
					reflect.TypeOf(ref).String() + " to type big.Int")
			}
		default:
			return nil, convertError(tag, "big.Int")
		}
	}
	return nil, err
}

// ReadBigIntWithoutTag from stream
func (r *Reader) ReadBigIntWithoutTag() (*big.Int, error) {
	s := r.Stream
	tag := TagSemicolon
	i := big.NewInt(0)
	b, err := s.ReadByte()
	if err == nil && b == tag {
		return i, nil
	}
	if err != nil {
		return i, err
	}
	pos := true
	switch b {
	case '-':
		pos = false
		fallthrough
	case '+':
		b, err = s.ReadByte()
	}
	for b != tag && err == nil {
		i = i.Mul(i, bigTen)
		i = i.Add(i, bigDigit[b-'0'])
		b, err = s.ReadByte()
	}
	if !pos {
		i = i.Neg(i)
	}
	return i, err
}

// ReadFloat32 from stream
func (r *Reader) ReadFloat32() (float32, error) {
	s := r.Stream
	tag, err := s.ReadByte()
	if err == nil {
		switch tag {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			return float32(tag - '0'), nil
		case TagInteger, TagLong:
			return r.readIntAsFloat32(TagSemicolon)
		case TagDouble:
			return r.ReadFloat32WithoutTag()
		case TagNull:
			return 0, ErrNil
		case TagEmpty, TagFalse:
			return 0, nil
		case TagTrue:
			return 1, nil
		case TagNaN:
			return float32(math.NaN()), nil
		case TagInfinity:
			f, err := r.readInfinity()
			return float32(f), err
		case TagUTF8Char:
			var str string
			if str, err = r.readUTF8String(1); err == nil {
				f, err := strconv.ParseFloat(str, 32)
				return float32(f), err
			}
		case TagString:
			var str string
			if str, err = r.ReadStringWithoutTag(); err == nil {
				f, err := strconv.ParseFloat(str, 32)
				return float32(f), err
			}
		case TagRef:
			var ref interface{}
			if ref, err = r.readRef(r.ReadInteger(TagSemicolon)); err == nil {
				if ref, ok := ref.(string); ok {
					f, err := strconv.ParseFloat(ref, 32)
					return float32(f), err
				}
				return 0, errors.New("cannot convert type " +
					reflect.TypeOf(ref).String() + " to type float32")
			}
		default:
			return 0, convertError(tag, "float32")
		}
	}
	return 0, err
}

// ReadFloat32WithoutTag from stream
func (r *Reader) ReadFloat32WithoutTag() (float32, error) {
	str, err := r.readUntil(TagSemicolon)
	if err != nil {
		return float32(math.NaN()), err
	}
	f, _ := strconv.ParseFloat(str, 32)
	return float32(f), nil
}

// ReadFloat64 from stream
func (r *Reader) ReadFloat64() (float64, error) {
	s := r.Stream
	tag, err := s.ReadByte()
	if err == nil {
		switch tag {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			return float64(tag - '0'), nil
		case TagInteger, TagLong:
			return r.readIntAsFloat64(TagSemicolon)
		case TagDouble:
			return r.ReadFloat64WithoutTag()
		case TagNull:
			return 0, ErrNil
		case TagEmpty, TagFalse:
			return 0, nil
		case TagTrue:
			return 1, nil
		case TagNaN:
			return math.NaN(), nil
		case TagInfinity:
			return r.readInfinity()
		case TagUTF8Char:
			var str string
			if str, err = r.readUTF8String(1); err == nil {
				return strconv.ParseFloat(str, 64)
			}
		case TagString:
			var str string
			if str, err = r.ReadStringWithoutTag(); err == nil {
				return strconv.ParseFloat(str, 64)
			}
		case TagRef:
			var ref interface{}
			if ref, err = r.readRef(r.ReadInteger(TagSemicolon)); err == nil {
				if ref, ok := ref.(string); ok {
					return strconv.ParseFloat(ref, 64)
				}
				return 0, errors.New("cannot convert type " +
					reflect.TypeOf(ref).String() + " to type float64")
			}
		default:
			return 0, convertError(tag, "float64")
		}
	}
	return 0, err
}

// ReadFloat64WithoutTag from stream
func (r *Reader) ReadFloat64WithoutTag() (float64, error) {
	str, err := r.readUntil(TagSemicolon)
	if err != nil {
		return math.NaN(), err
	}
	f, _ := strconv.ParseFloat(str, 64)
	return f, nil
}

// ReadBool from stream
func (r *Reader) ReadBool() (bool, error) {
	s := r.Stream
	tag, err := s.ReadByte()
	if err == nil {
		switch tag {
		case '0':
			return false, nil
		case '1', '2', '3', '4', '5', '6', '7', '8', '9':
			return true, nil
		case TagInteger, TagLong:
			i, err := r.ReadInt64WithoutTag()
			return i != 0, err
		case TagDouble:
			f, err := r.ReadFloat64WithoutTag()
			return f != 0, err
		case TagNull:
			return false, ErrNil
		case TagEmpty, TagFalse:
			return false, nil
		case TagTrue, TagNaN:
			return true, nil
		case TagInfinity:
			_, err = r.readInfinity()
			return true, err
		case TagUTF8Char:
			var str string
			if str, err = r.readUTF8String(1); err == nil {
				return strconv.ParseBool(str)
			}
		case TagString:
			var str string
			if str, err = r.ReadStringWithoutTag(); err == nil {
				return strconv.ParseBool(str)
			}
		case TagRef:
			var ref interface{}
			if ref, err = r.readRef(r.ReadInteger(TagSemicolon)); err == nil {
				if ref, ok := ref.(string); ok {
					return strconv.ParseBool(ref)
				}
				return false, errors.New("cannot convert type " +
					reflect.TypeOf(ref).String() + " to type bool")
			}
		default:
			return false, convertError(tag, "bool")
		}
	}
	return false, err
}

// ReadDateTime from stream
func (r *Reader) ReadDateTime() (time.Time, error) {
	s := r.Stream
	tag, err := s.ReadByte()
	if err == nil {
		switch tag {
		case '0', TagEmpty:
			return timeZero, nil
		case TagNull:
			return timeZero, ErrNil
		case TagString:
			var str string
			if str, err = r.ReadStringWithoutTag(); err == nil {
				return time.Parse(timeStringFormat, str)
			}
		case TagDate:
			return r.ReadDateWithoutTag()
		case TagTime:
			return r.ReadTimeWithoutTag()
		case TagRef:
			var ref interface{}
			if ref, err = r.readRef(r.ReadInteger(TagSemicolon)); err == nil {
				switch ref := ref.(type) {
				case time.Time:
					return ref, nil
				case string:
					return time.Parse(timeStringFormat, ref)
				default:
					return timeZero, errors.New("cannot convert type " +
						reflect.TypeOf(ref).String() + " to type time.Time")
				}
			}
		default:
			return timeZero, convertError(tag, "time.Time")
		}
	}
	return timeZero, err
}

// ReadDateWithoutTag from stream
func (r *Reader) ReadDateWithoutTag() (time.Time, error) {
	s := r.Stream
	var year, month, day, hour, min, sec, nsec int
	tag, err := s.ReadByte()
	if err == nil {
		year = int(tag - '0')
		if tag, err = s.ReadByte(); err == nil {
			year *= 10
			year += int(tag - '0')
			if tag, err = s.ReadByte(); err == nil {
				year *= 10
				year += int(tag - '0')
				if tag, err = s.ReadByte(); err == nil {
					year *= 10
					year += int(tag - '0')
					if tag, err = s.ReadByte(); err == nil {
						month = int(tag - '0')
						if tag, err = s.ReadByte(); err == nil {
							month *= 10
							month += int(tag - '0')
							if tag, err = s.ReadByte(); err == nil {
								day = int(tag - '0')
								if tag, err = s.ReadByte(); err == nil {
									day *= 10
									day += int(tag - '0')
									tag, err = s.ReadByte()
								}
							}
						}
					}
				}
			}
		}
	}
	if err != nil {
		return timeZero, err
	}
	if tag == TagTime {
		if hour, min, sec, nsec, tag, err = r.readTime(); err != nil {
			return timeZero, err
		}
	}
	var loc *time.Location
	if tag == TagUTC {
		loc = time.UTC
	} else if tag == TagSemicolon {
		loc = time.Local
	} else {
		return timeZero, unexpectedTag(tag, []byte{TagUTC, TagSemicolon})
	}
	d := time.Date(year, time.Month(month), day, hour, min, sec, nsec, loc)
	r.setRef(d)
	return d, nil
}

// ReadTimeWithoutTag from stream
func (r *Reader) ReadTimeWithoutTag() (time.Time, error) {
	hour, min, sec, nsec, tag, err := r.readTime()
	if err != nil {
		return timeZero, err
	}
	var loc *time.Location
	if tag == TagUTC {
		loc = time.UTC
	} else if tag == TagSemicolon {
		loc = time.Local
	} else {
		return timeZero, unexpectedTag(tag, []byte{TagUTC, TagSemicolon})
	}
	t := time.Date(1970, 1, 1, hour, min, sec, nsec, loc)
	r.setRef(t)
	return t, nil
}

// ReadString from stream
func (r *Reader) ReadString() (string, error) {
	s := r.Stream
	tag, err := s.ReadByte()
	if err == nil {
		switch tag {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			return string([]byte{tag}), nil
		case TagInteger, TagLong, TagDouble:
			return r.readUntil(TagSemicolon)
		case TagNull:
			return "", ErrNil
		case TagEmpty:
			return "", nil
		case TagTrue:
			return "true", nil
		case TagFalse:
			return "false", nil
		case TagNaN:
			return "NaN", nil
		case TagInfinity:
			if sign, err := s.ReadByte(); err == nil {
				return string([]byte{sign}) + "Inf", nil
			}
			return "Inf", err
		case TagDate:
			d, err := r.ReadDateWithoutTag()
			return d.String(), err
		case TagTime:
			t, err := r.ReadTimeWithoutTag()
			return t.String(), err
		case TagUTF8Char:
			return r.readUTF8String(1)
		case TagString:
			return r.ReadStringWithoutTag()
		case TagGuid:
			u, err := r.ReadUUIDWithoutTag()
			return u.String(), err
		case TagBytes:
			if b, err := r.ReadBytesWithoutTag(); err == nil {
				if !utf8.Valid(*b) {
					err = errBadEncode
				}
				return string(*b), err
			}
		case TagRef:
			var ref interface{}
			if ref, err = r.readRef(r.ReadInteger(TagSemicolon)); err == nil {
				if ref, ok := ref.(string); ok {
					return ref, nil
				}
				return "", errors.New("cannot convert type " +
					reflect.TypeOf(ref).String() + " to type string")
			}
		default:
			return "", convertError(tag, "string")
		}
	}
	return "", err
}

// ReadStringWithoutTag from stream
func (r *Reader) ReadStringWithoutTag() (str string, err error) {
	if str, err = r.readStringWithoutTag(); err == nil {
		r.setRef(str)
	}
	return str, err
}

// ReadBytes from stream
func (r *Reader) ReadBytes() (*[]byte, error) {
	bytes := new([]byte)
	s := r.Stream
	tag, err := s.ReadByte()
	if err == nil {
		switch tag {
		case TagNull:
			return bytes, ErrNil
		case TagEmpty:
			return bytes, nil
		case TagUTF8Char:
			c, err := r.readUTF8String(1)
			*bytes = []byte(c)
			return bytes, err
		case TagString:
			str, err := r.ReadStringWithoutTag()
			*bytes = []byte(str)
			return bytes, err
		case TagGuid:
			u, err := r.ReadUUIDWithoutTag()
			*bytes = []byte(*u)
			return bytes, err
		case TagBytes:
			return r.ReadBytesWithoutTag()
		case TagList:
			err = r.ReadSliceWithoutTag(bytes)
			return bytes, err
		case TagRef:
			var ref interface{}
			if ref, err = r.readRef(r.ReadInteger(TagSemicolon)); err == nil {
				if ref, ok := ref.(*[]byte); ok {
					return ref, nil
				}
				return bytes, errors.New("cannot convert type " +
					reflect.TypeOf(ref).String() + " to type bytes")
			}
		default:
			return bytes, convertError(tag, "bytes")
		}
	}
	return bytes, err
}

// ReadBytesWithoutTag from stream
func (r *Reader) ReadBytesWithoutTag() (*[]byte, error) {
	s := r.Stream
	length, err := r.ReadInteger(TagQuote)
	if err != nil {
		return new([]byte), err
	}
	b := make([]byte, length)
	if _, err = s.Read(b); err == nil {
		err = r.CheckTag(TagQuote)
	}
	r.setRef(&b)
	return &b, err
}

// ReadUUID from stream
func (r *Reader) ReadUUID() (*UUID, error) {
	uuid := new(UUID)
	s := r.Stream
	tag, err := s.ReadByte()
	if err == nil {
		switch tag {
		case TagNull:
			return uuid, ErrNil
		case TagString:
			str, err := r.ReadStringWithoutTag()
			*uuid = ToUUID(str)
			return uuid, err
		case TagGuid:
			return r.ReadUUIDWithoutTag()
		case TagBytes:
			if b, err := r.ReadBytesWithoutTag(); err == nil {
				if len(*b) == 16 {
					uuid = (*UUID)(b)
					return uuid, nil
				}
				return uuid, convertError(TagBytes, "UUID")
			}
			return uuid, err
		case TagRef:
			var ref interface{}
			if ref, err = r.readRef(r.ReadInteger(TagSemicolon)); err == nil {
				if ref, ok := ref.(*UUID); ok {
					return ref, nil
				}
				return uuid, errors.New("cannot convert type " +
					reflect.TypeOf(ref).String() + " to type UUID")
			}
		default:
			return uuid, convertError(tag, "UUID")
		}
	}
	return uuid, err
}

// ReadUUIDWithoutTag from stream
func (r *Reader) ReadUUIDWithoutTag() (*UUID, error) {
	s := r.Stream
	err := r.CheckTag(TagOpenbrace)
	if err == nil {
		b := make([]byte, 36)
		if _, err = s.Read(b); err == nil {
			err = r.CheckTag(TagClosebrace)
			u := ToUUID(string(b))
			r.setRef(&u)
			return &u, err
		}
	}
	return new(UUID), err
}

// ReadList from stream
func (r *Reader) ReadList() (*list.List, error) {
	l := list.New()
	s := r.Stream
	tag, err := s.ReadByte()
	if err == nil {
		switch tag {
		case TagNull:
			return l, ErrNil
		case TagList:
			return r.ReadListWithoutTag()
		case TagRef:
			var ref interface{}
			if ref, err = r.readRef(r.ReadInteger(TagSemicolon)); err == nil {
				if ref, ok := ref.(*list.List); ok {
					return ref, nil
				}
				return l, errors.New("cannot convert type " +
					reflect.TypeOf(ref).String() + " to type List")
			}
		default:
			return l, convertError(tag, "List")
		}
	}
	return l, err
}

// ReadListWithoutTag from stream
func (r *Reader) ReadListWithoutTag() (*list.List, error) {
	l := list.New()
	r.setRef(l)
	length, err := r.ReadInteger(TagOpenbrace)
	if err == nil {
		for i := 0; i < length; i++ {
			if e, err := r.readInterface(); err == nil {
				l.PushBack(e)
			} else {
				return l, err
			}
		}
		if err = r.CheckTag(TagClosebrace); err == nil {
			return l, nil
		}
	}
	return l, err
}

// ReadArray from stream
func (r *Reader) ReadArray(a []reflect.Value) error {
	length := len(a)
	r.setRef(&a)
	for i := 0; i < length; i++ {
		if err := r.ReadValue(a[i]); err != nil {
			return err
		}
	}
	return r.CheckTag(TagClosebrace)
}

// ReadSlice from stream
func (r *Reader) ReadSlice(p interface{}) error {
	v, err := r.checkPointer(p)
	if err == nil {
		return r.readSlice(v.Elem())
	}
	return err
}

// ReadSliceWithoutTag from stream
func (r *Reader) ReadSliceWithoutTag(p interface{}) error {
	v, err := r.checkPointer(p)
	if err == nil {
		return r.readSliceWithoutTag(v.Elem())
	}
	return err
}

// ReadMap from stream
func (r *Reader) ReadMap(p interface{}) error {
	v, err := r.checkPointer(p)
	if err == nil {
		return r.readMap(v.Elem())
	}
	return err
}

// ReadMapWithoutTag from stream
func (r *Reader) ReadMapWithoutTag(p interface{}) error {
	v, err := r.checkPointer(p)
	if err == nil {
		return r.readMapWithoutTag(v.Elem())
	}
	return err
}

// ReadObject from stream
func (r *Reader) ReadObject(p interface{}) error {
	v, err := r.checkPointer(p)
	if err == nil {
		return r.readObject(v.Elem())
	}
	return err
}

// ReadObjectWithoutTag from stream
func (r *Reader) ReadObjectWithoutTag(p interface{}) error {
	v, err := r.checkPointer(p)
	if err == nil {
		return r.readObjectWithoutTag(v.Elem())
	}
	return err
}

// Reset the serialize reference count
func (r *Reader) Reset() {
	if r.classref != nil {
		r.classref = r.classref[:0]
		r.fieldsref = r.fieldsref[:0]
	}
	r.resetRef()
}

// ReadValue from stream
func (r *Reader) ReadValue(v reflect.Value) error {
	t := v.Type()
	switch t.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return r.readInt64(v)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return r.readUint64(v)
	case reflect.Bool:
		return r.readBool(v)
	case reflect.Float32:
		return r.readFloat32(v)
	case reflect.Float64:
		return r.readFloat64(v)
	case reflect.String:
		return r.readString(v)
	case reflect.Slice:
		if t.Name() == "UUID" {
			return r.readUUID(v)
		}
		if t.Elem().Kind() == reflect.Uint8 {
			return r.readBytes(v)
		}
		return r.readSlice(v)
	case reflect.Map:
		return r.readMap(v)
	case reflect.Struct:
		switch t.Name() {
		case "Time":
			return r.readDateTime(v)
		case "Int":
			return r.readBigInt(v)
		case "List":
			return r.readList(v)
		}
		return r.readObject(v)
	case reflect.Interface:
		p, err := r.readInterface()
		if err == nil {
			t := v.Type()
			rv := reflect.ValueOf(&p).Elem()
			rt := rv.Type()
			if rt.AssignableTo(t) {
				v.Set(rv)
			} else if rt.ConvertibleTo(t) {
				v.Set(rv.Convert(t))
			}
		}
		return err
	case reflect.Ptr:
		switch t := t.Elem(); t.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return r.readInt64Pointer(v)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return r.readUint64Pointer(v)
		case reflect.Bool:
			return r.readBoolPointer(v)
		case reflect.Float32:
			return r.readFloat32Pointer(v)
		case reflect.Float64:
			return r.readFloat64Pointer(v)
		case reflect.String:
			return r.readStringPointer(v)
		case reflect.Slice:
			if t.Name() == "UUID" {
				return r.readUUIDPointer(v)
			}
			if t.Elem().Kind() == reflect.Uint8 {
				return r.readBytesPointer(v)
			}
			return r.readSlice(v)
		case reflect.Map:
			return r.readMap(v)
		case reflect.Struct:
			switch t.Name() {
			case "Time":
				return r.readDateTimePointer(v)
			case "Int":
				return r.readBigIntPointer(v)
			case "List":
				return r.readListPointer(v)
			}
			return r.readObject(v)
		case reflect.Interface:
			p, err := r.readInterface()
			if err == nil {
				t := v.Type()
				rp := reflect.ValueOf(&p)
				rt := rp.Type()
				if rt.AssignableTo(t) {
					v.Set(rp)
				} else if rt.ConvertibleTo(t) {
					v.Set(rp.Convert(t))
				} else if rt.Elem().ConvertibleTo(t.Elem()) {
					v.Set(reflect.New(t.Elem()))
					v.Elem().Set(rp.Elem().Convert(t.Elem()))
				}
			}
			return err
		}
	}
	return errors.New("unsupported Type:" + t.String())
}

// private methods

func (r *Reader) checkPointer(p interface{}) (v reflect.Value, err error) {
	v = reflect.ValueOf(p)
	if v.Kind() != reflect.Ptr {
		return v, errors.New("argument p must be a pointer")
	}
	return v, nil
}

func (r *Reader) readInterface() (interface{}, error) {
	s := r.Stream
	tag, err := s.ReadByte()
	if err == nil {
		switch tag {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			return int(tag - '0'), nil
		case TagInteger:
			return r.ReadIntWithoutTag()
		case TagLong:
			return r.ReadBigIntWithoutTag()
		case TagDouble:
			return r.ReadFloat64WithoutTag()
		case TagNull:
			return nil, nil
		case TagEmpty:
			return "", nil
		case TagTrue:
			return true, nil
		case TagFalse:
			return false, nil
		case TagNaN:
			return math.NaN(), nil
		case TagInfinity:
			return r.readInfinity()
		case TagDate:
			return r.ReadDateWithoutTag()
		case TagTime:
			return r.ReadTimeWithoutTag()
		case TagBytes:
			return r.ReadBytesWithoutTag()
		case TagUTF8Char:
			return r.readUTF8String(1)
		case TagString:
			return r.ReadStringWithoutTag()
		case TagGuid:
			return r.ReadUUIDWithoutTag()
		case TagList:
			var e *[]interface{}
			err := r.ReadSliceWithoutTag(&e)
			return e, err
		case TagMap:
			if r.JSONCompatible {
				var e *map[string]interface{}
				err := r.ReadMapWithoutTag(&e)
				return e, err
			}
			var e *map[interface{}]interface{}
			err := r.ReadMapWithoutTag(&e)
			return e, err
		case TagClass:
			err := r.readClass()
			if err == nil {
				var e interface{}
				err := r.ReadObject(&e)
				return e, err
			}
			return nil, err
		case TagObject:
			var e interface{}
			err := r.ReadObjectWithoutTag(&e)
			return e, err
		case TagRef:
			return r.readRef(r.ReadInteger(TagSemicolon))
		}
		return nil, unexpectedTag(tag, nil)
	}
	return nil, err
}

func (r *Reader) readUntil(tag byte) (string, error) {
	s, err := r.Stream.ReadString(tag)
	if err != nil {
		return s, err
	}
	return s[:len(s)-1], nil
}

func (r *Reader) readInt(tag byte) (int64, error) {
	s := r.Stream
	i := int64(0)
	b, err := s.ReadByte()
	if err == nil && b == tag {
		return i, nil
	}
	if err != nil {
		return i, err
	}
	sign := int64(1)
	switch b {
	case '-':
		sign = -1
		fallthrough
	case '+':
		b, err = s.ReadByte()
	}
	for b != tag && err == nil {
		i *= 10
		i += int64(b-'0') * sign
		b, err = s.ReadByte()
	}
	return i, err
}

func (r *Reader) readUint(tag byte) (uint64, error) {
	i, err := r.readInt(tag)
	return uint64(i), err
}

func (r *Reader) readInt8(tag byte) (int8, error) {
	s := r.Stream
	i := int8(0)
	b, err := s.ReadByte()
	if err == nil && b == tag {
		return i, nil
	}
	if err != nil {
		return i, err
	}
	sign := int8(1)
	switch b {
	case '-':
		sign = -1
		fallthrough
	case '+':
		b, err = s.ReadByte()
	}
	for b != tag && err == nil {
		i *= 10
		i += int8(b-'0') * sign
		b, err = s.ReadByte()
	}
	return i, err
}

func (r *Reader) readUint8(tag byte) (uint8, error) {
	i, err := r.readInt8(tag)
	return uint8(i), err
}

func (r *Reader) readInt16(tag byte) (int16, error) {
	s := r.Stream
	i := int16(0)
	b, err := s.ReadByte()
	if err == nil && b == tag {
		return i, nil
	}
	if err != nil {
		return i, err
	}
	sign := int16(1)
	switch b {
	case '-':
		sign = -1
		fallthrough
	case '+':
		b, err = s.ReadByte()
	}
	for b != tag && err == nil {
		i *= 10
		i += int16(b-'0') * sign
		b, err = s.ReadByte()
	}
	return i, err
}

func (r *Reader) readUint16(tag byte) (uint16, error) {
	i, err := r.readInt16(tag)
	return uint16(i), err
}

func (r *Reader) readInt32(tag byte) (int32, error) {
	s := r.Stream
	i := int32(0)
	b, err := s.ReadByte()
	if err == nil && b == tag {
		return i, nil
	}
	if err != nil {
		return i, err
	}
	sign := int32(1)
	switch b {
	case '-':
		sign = -1
		fallthrough
	case '+':
		b, err = s.ReadByte()
	}
	for b != tag && err == nil {
		i *= 10
		i += int32(b-'0') * sign
		b, err = s.ReadByte()
	}
	return i, err
}

func (r *Reader) readUint32(tag byte) (uint32, error) {
	i, err := r.readInt32(tag)
	return uint32(i), err
}

func (r *Reader) readIntAsFloat64(tag byte) (float64, error) {
	s := r.Stream
	f := float64(0)
	b, err := s.ReadByte()
	if err == nil && b == tag {
		return f, nil
	}
	if err != nil {
		return f, err
	}
	sign := float64(1)
	switch b {
	case '-':
		sign = -1
		fallthrough
	case '+':
		b, err = s.ReadByte()
	}
	for b != tag && err == nil {
		f *= 10
		f += float64(b-'0') * sign
		b, err = s.ReadByte()
	}
	return f, err
}

func (r *Reader) readIntAsFloat32(tag byte) (float32, error) {
	f, err := r.readIntAsFloat64(tag)
	return float32(f), err
}

func (r *Reader) readInfinity() (float64, error) {
	if sign, err := r.Stream.ReadByte(); err == nil {
		switch sign {
		case '+':
			return math.Inf(1), nil
		case '-':
			return math.Inf(-1), nil
		default:
			return math.NaN(), unexpectedTag(sign, []byte{'+', '-'})
		}
	} else {
		return math.NaN(), err
	}
}

func (r *Reader) readTime() (hour int, min int, sec int, nsec int, tag byte, err error) {
	s := r.Stream
	if tag, err = s.ReadByte(); err == nil {
		hour = int(tag - '0')
		if tag, err = s.ReadByte(); err == nil {
			hour *= 10
			hour += int(tag - '0')
			if tag, err = s.ReadByte(); err == nil {
				min = int(tag - '0')
				if tag, err = s.ReadByte(); err == nil {
					min *= 10
					min += int(tag - '0')
					if tag, err = s.ReadByte(); err == nil {
						sec = int(tag - '0')
						if tag, err = s.ReadByte(); err == nil {
							sec *= 10
							sec += int(tag - '0')
							tag, err = s.ReadByte()
						}
					}
				}
			}
		}
	}
	if err != nil {
		return hour, min, sec, nsec, tag, err
	}
	if tag == TagPoint {
		if tag, err = s.ReadByte(); err == nil {
			nsec = int(tag - '0')
			if tag, err = s.ReadByte(); err == nil {
				nsec *= 10
				nsec += int(tag - '0')
				if tag, err = s.ReadByte(); err == nil {
					nsec *= 10
					nsec += int(tag - '0')
					tag, err = s.ReadByte()
				}
			}
		}
		if err != nil {
			return hour, min, sec, nsec, tag, err
		}
		if tag >= '0' && tag <= '9' {
			nsec *= 10
			nsec += int(tag - '0')
			if tag, err = s.ReadByte(); err == nil {
				nsec *= 10
				nsec += int(tag - '0')
				if tag, err = s.ReadByte(); err == nil {
					nsec *= 10
					nsec += int(tag - '0')
					tag, err = s.ReadByte()
				}
			}
		} else {
			nsec *= 1000
		}
		if err != nil {
			return hour, min, sec, nsec, tag, err
		}
		if tag >= '0' && tag <= '9' {
			nsec *= 10
			nsec += int(tag - '0')
			if tag, err = s.ReadByte(); err == nil {
				nsec *= 10
				nsec += int(tag - '0')
				if tag, err = s.ReadByte(); err == nil {
					nsec *= 10
					nsec += int(tag - '0')
					tag, err = s.ReadByte()
				}
			}
		} else {
			nsec *= 1000
		}
	}
	return hour, min, sec, nsec, tag, err
}

func (r *Reader) readStringWithoutTag() (str string, err error) {
	var length int
	if length, err = r.ReadInteger(TagQuote); err == nil {
		if str, err = r.readUTF8String(length); err == nil {
			err = r.CheckTag(TagQuote)
		}
	}
	return str, err
}

func (r *Reader) readInt64(v reflect.Value) error {
	x, err := r.ReadInt64()
	if err == nil || err == ErrNil {
		v.SetInt(x)
		return nil
	}
	return err
}

func (r *Reader) readUint64(v reflect.Value) error {
	x, err := r.ReadUint64()
	if err == nil || err == ErrNil {
		v.SetUint(x)
		return nil
	}
	return err
}

func (r *Reader) readBool(v reflect.Value) error {
	x, err := r.ReadBool()
	if err == nil || err == ErrNil {
		v.SetBool(x)
		return nil
	}
	return err
}

func (r *Reader) readFloat32(v reflect.Value) error {
	x, err := r.ReadFloat32()
	if err == nil || err == ErrNil {
		v.SetFloat(float64(x))
		return nil
	}
	return err
}

func (r *Reader) readFloat64(v reflect.Value) error {
	x, err := r.ReadFloat64()
	if err == nil || err == ErrNil {
		v.SetFloat(x)
		return nil
	}
	return err
}

func (r *Reader) readBigInt(v reflect.Value) error {
	x, err := r.ReadBigInt()
	if err == nil || err == ErrNil {
		v.Set(reflect.ValueOf(*x))
		return nil
	}
	return err
}

func (r *Reader) readDateTime(v reflect.Value) error {
	x, err := r.ReadDateTime()
	if err == nil || err == ErrNil {
		v.Set(reflect.ValueOf(x))
		return nil
	}
	return err
}

func (r *Reader) readString(v reflect.Value) error {
	x, err := r.ReadString()
	if err == nil || err == ErrNil {
		v.SetString(x)
		return nil
	}
	return err
}

func (r *Reader) readBytes(v reflect.Value) error {
	x, err := r.ReadBytes()
	if err == nil || err == ErrNil {
		v.Set(reflect.ValueOf(*x))
		return nil
	}
	return err
}

func (r *Reader) readUUID(v reflect.Value) error {
	x, err := r.ReadUUID()
	if err == nil || err == ErrNil {
		v.Set(reflect.ValueOf(*x))
		return nil
	}
	return err
}

func (r *Reader) readList(v reflect.Value) error {
	x, err := r.ReadList()
	if err == nil || err == ErrNil {
		v.Set(reflect.ValueOf(*x))
		return nil
	}
	return err
}

func (r *Reader) getRef(v reflect.Value) error {
	t := v.Type()
	ref, err := r.readRef(r.ReadInteger(TagSemicolon))
	if err != nil {
		return err
	}
	refValue := reflect.ValueOf(ref)
	refType := refValue.Type()
	if refType.Kind() == reflect.Ptr {
		if refType.AssignableTo(t) {
			v.Set(refValue)
			return nil
		}
		if refType.Elem().AssignableTo(t) {
			v.Set(refValue.Elem())
			return nil
		}
	}
	return errors.New("cannot convert type " +
		refType.String() + " to type " + t.String())
}

func (r *Reader) readSlice(v reflect.Value) error {
	s := r.Stream
	t := v.Type()
	tag, err := s.ReadByte()
	if err == nil {
		switch tag {
		case TagNull:
			v.Set(reflect.Zero(t))
			return nil
		case TagList:
			return r.readSliceWithoutTag(v)
		case TagRef:
			return r.getRef(v)

		}
		return convertError(tag, v.Type().String())
	}
	return err
}

func (r *Reader) readSliceWithoutTag(v reflect.Value) error {
	t := v.Type()
	switch t.Kind() {
	case reflect.Slice:
	case reflect.Interface:
		t = reflect.TypeOf([]interface{}(nil))
	case reflect.Ptr:
		switch t = t.Elem(); t.Kind() {
		case reflect.Slice:
		case reflect.Interface:
			t = reflect.TypeOf([]interface{}(nil))
		default:
			return errors.New("cannot convert slice to type " + t.String())
		}
	default:
		return errors.New("cannot convert slice to type " + t.String())
	}
	slicePointer := reflect.New(t)
	r.setRef(slicePointer.Interface())
	slice := slicePointer.Elem()
	length, err := r.ReadInteger(TagOpenbrace)
	if err == nil {
		slice.Set(reflect.MakeSlice(t, length, length))
		for i := 0; i < length; i++ {
			elem := slice.Index(i)
			if err := r.ReadValue(elem); err != nil {
				return err
			}
		}
		if err = r.CheckTag(TagClosebrace); err == nil {
			switch t := v.Type(); t.Kind() {
			case reflect.Slice:
				v.Set(slice)
			case reflect.Interface:
				v.Set(slicePointer)
			case reflect.Ptr:
				switch t.Elem().Kind() {
				case reflect.Slice:
					v.Set(slicePointer)
				case reflect.Interface:
					v.Set(reflect.New(t.Elem()))
					v.Elem().Set(slicePointer)
				}
			}
			return nil
		}
	}
	return err
}

func (r *Reader) readSliceAsMap(v reflect.Value) error {
	t := v.Type()
	switch t.Kind() {
	case reflect.Map:
		break
	case reflect.Interface:
		if r.JSONCompatible {
			t = soMapType
		} else {
			t = ooMapType
		}
	case reflect.Ptr:
		switch t = t.Elem(); t.Kind() {
		case reflect.Map:
			break
		case reflect.Interface:
			if r.JSONCompatible {
				t = soMapType
			} else {
				t = ooMapType
			}
		default:
			return errors.New("cannot convert slice to type " + t.String())
		}
	default:
		return errors.New("cannot convert slice to type " + t.String())
	}
	mPointer := reflect.New(t)
	r.setRef(mPointer.Interface())
	m := mPointer.Elem()
	length, err := r.ReadInteger(TagOpenbrace)
	if err == nil {
		m.Set(reflect.MakeMap(t))
		for i := 0; i < length; i++ {
			key := reflect.New(t.Key()).Elem()
			val := reflect.New(t.Elem()).Elem()
			switch t.Key().Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				key.SetInt(int64(i))
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				key.SetUint(uint64(i))
			case reflect.Float32, reflect.Float64:
				key.SetFloat(float64(i))
			case reflect.String:
				key.SetString(strconv.Itoa(i))
			case reflect.Interface:
				key.Set(reflect.ValueOf(i))
			default:
				return errors.New("cannot convert int to type " + t.Key().String())
			}
			if err := r.ReadValue(val); err != nil {
				return err
			}
			m.SetMapIndex(key, val)
		}
		if err = r.CheckTag(TagClosebrace); err == nil {
			switch t := v.Type(); t.Kind() {
			case reflect.Map:
				v.Set(m)
			case reflect.Interface:
				v.Set(mPointer)
			case reflect.Ptr:
				switch t.Elem().Kind() {
				case reflect.Map:
					v.Set(mPointer)
				case reflect.Interface:
					v.Set(reflect.New(t.Elem()))
					v.Elem().Set(mPointer)
				}
			}
			return nil
		}
	}
	return err
}

func (r *Reader) readMap(v reflect.Value) error {
	s := r.Stream
	t := v.Type()
	tag, err := s.ReadByte()
	if err == nil {
		switch tag {
		case TagNull:
			v.Set(reflect.Zero(t))
			return nil
		case TagList:
			return r.readSliceAsMap(v)
		case TagMap:
			return r.readMapWithoutTag(v)
		case TagClass:
			if err = r.readClass(); err == nil {
				return r.readMap(v)
			}
			return err
		case TagObject:
			return r.readObjectWithoutTag(v)
		case TagRef:
			return r.getRef(v)
		}
		return convertError(tag, v.Type().String())
	}
	return err
}

func (r *Reader) readMapWithoutTag(v reflect.Value) error {
	t := v.Type()
	switch t.Kind() {
	case reflect.Struct:
		return r.readMapAsObject(v)
	case reflect.Map:
		break
	case reflect.Interface:
		if r.JSONCompatible {
			t = soMapType
		} else {
			t = ooMapType
		}
	case reflect.Ptr:
		switch t = t.Elem(); t.Kind() {
		case reflect.Struct:
			return r.readMapAsObject(v)
		case reflect.Map:
			break
		case reflect.Interface:
			if r.JSONCompatible {
				t = soMapType
			} else {
				t = ooMapType
			}
		default:
			return errors.New("cannot convert map to type " + t.String())
		}
	default:
		return errors.New("cannot convert map to type " + t.String())
	}
	mPointer := reflect.New(t)
	r.setRef(mPointer.Interface())
	m := mPointer.Elem()
	length, err := r.ReadInteger(TagOpenbrace)
	if err == nil {
		m.Set(reflect.MakeMap(t))
		tk := t.Key()
		tv := t.Elem()
		for i := 0; i < length; i++ {
			key := reflect.New(tk).Elem()
			val := reflect.New(tv).Elem()
			if err := r.ReadValue(key); err != nil {
				return err
			}
			if err := r.ReadValue(val); err != nil {
				return err
			}
			m.SetMapIndex(key, val)
		}
		if err = r.CheckTag(TagClosebrace); err == nil {
			switch t := v.Type(); t.Kind() {
			case reflect.Map:
				v.Set(m)
			case reflect.Interface:
				v.Set(mPointer)
			case reflect.Ptr:
				switch t.Elem().Kind() {
				case reflect.Map:
					v.Set(mPointer)
				case reflect.Interface:
					v.Set(reflect.New(t.Elem()))
					v.Elem().Set(mPointer)
				}
			}
			return nil
		}
	}
	return err
}

func (r *Reader) readMapAsObject(v reflect.Value) error {
	t := v.Type()
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	objPointer := reflect.New(t)
	r.setRef(objPointer.Interface())
	obj := objPointer.Elem()
	count, err := r.ReadInteger(TagOpenbrace)
	if err == nil {
		indexMap := getIndexCache(t)
		for i := 0; i < count; i++ {
			key, err := r.ReadString()
			if err != nil {
				return err
			}
			if index, ok := indexMap[strings.ToLower(key)]; ok {
				f := obj.Field(index[0])
				n := len(index)
				for j := 1; j < n; j++ {
					if f.Kind() == reflect.Ptr {
						f.Set(reflect.New(f.Type().Elem()))
						f = f.Elem()
					}
					f = f.Field(index[j])
				}
				err = r.ReadValue(f)
			} else {
				_, err = r.readInterface()
			}
			if err != nil {
				return err
			}
		}
		if err = r.CheckTag(TagClosebrace); err == nil {
			switch t := v.Type(); t.Kind() {
			case reflect.Struct:
				v.Set(obj)
			case reflect.Interface:
				v.Set(objPointer)
			case reflect.Ptr:
				switch t.Elem().Kind() {
				case reflect.Struct:
					v.Set(objPointer)
				case reflect.Interface:
					v.Set(reflect.New(t.Elem()))
					v.Elem().Set(objPointer)
				}
			}
			return nil
		}
	}
	return err
}

func (r *Reader) readObject(v reflect.Value) error {
	s := r.Stream
	t := v.Type()
	tag, err := s.ReadByte()
	if err == nil {
		switch tag {
		case TagNull:
			v.Set(reflect.Zero(t))
			return nil
		case TagMap:
			return r.readMapWithoutTag(v)
		case TagClass:
			if err = r.readClass(); err == nil {
				return r.readObject(v)
			}
			return err
		case TagObject:
			return r.readObjectWithoutTag(v)
		case TagRef:
			return r.getRef(v)
		}
		return convertError(tag, v.Type().String())
	}
	return err
}

func (r *Reader) readObjectAsMap(v reflect.Value, index int) error {
	t := soMapType
	mPointer := reflect.New(t)
	r.setRef(mPointer.Interface())
	m := mPointer.Elem()
	m.Set(reflect.MakeMap(t))
	fields := r.fieldsref[index]
	length := len(fields)
	tk := t.Key()
	tv := t.Elem()
	for i := 0; i < length; i++ {
		key := reflect.New(tk).Elem()
		val := reflect.New(tv).Elem()
		key.SetString(fields[i])
		if err := r.ReadValue(val); err != nil {
			return err
		}
		m.SetMapIndex(key, val)
	}
	err := r.CheckTag(TagClosebrace)
	if err == nil {
		switch t := v.Type(); t.Kind() {
		case reflect.Map:
			v.Set(m)
		case reflect.Interface:
			v.Set(mPointer)
		case reflect.Ptr:
			switch t.Elem().Kind() {
			case reflect.Map:
				v.Set(mPointer)
			case reflect.Interface:
				v.Set(reflect.New(t.Elem()))
				v.Elem().Set(mPointer)
			}
		}
	}
	return err
}

func (r *Reader) readObjectWithoutTag(v reflect.Value) error {
	t := v.Type()
	kind := t.Kind()
	index, err := r.ReadInteger(TagOpenbrace)
	if err != nil {
		return err
	}
	key := r.classref[index]
	class, ok := key.(reflect.Type)
	if !ok {
		if kind == reflect.Struct {
			class = t
		} else if kind == reflect.Ptr && t.Elem().Kind() == reflect.Struct {
			class = t.Elem()
		} else {
			class = soMapType
		}
	} else {
		if t == soMapType || (kind == reflect.Ptr && t.Elem() == soMapType) {
			class = soMapType
		}
	}
	assignable := class.AssignableTo(t)
	if kind == reflect.Ptr {
		assignable = class.AssignableTo(t.Elem())
	}
	if !assignable {
		return errors.New("cannot convert type " + class.String() + " to type " + t.String())
	}
	if class.Kind() == reflect.Map {
		return r.readObjectAsMap(v, index)
	}
	objPointer := reflect.New(class)
	r.setRef(objPointer.Interface())
	obj := objPointer.Elem()
	fields := r.fieldsref[index]
	indexMap := getIndexCache(class)
	count := len(fields)
	for i := 0; i < count; i++ {
		if index, ok := indexMap[strings.ToLower(fields[i])]; ok {
			f := obj.Field(index[0])
			n := len(index)
			for j := 1; j < n; j++ {
				if f.Kind() == reflect.Ptr {
					f.Set(reflect.New(f.Type().Elem()))
					f = f.Elem()
				}
				f = f.Field(index[j])
			}
			err = r.ReadValue(f)
		} else {
			_, err = r.readInterface()
		}
		if err != nil {
			return err
		}
	}
	if err = r.CheckTag(TagClosebrace); err == nil {
		switch kind {
		case reflect.Struct:
			v.Set(obj)
		case reflect.Interface:
			v.Set(objPointer)
		case reflect.Ptr:
			switch t := t.Elem(); t.Kind() {
			case reflect.Struct:
				v.Set(objPointer)
			case reflect.Interface:
				v.Set(reflect.New(t))
				v.Elem().Set(objPointer)
			}
		}
		return nil
	}
	return err
}

func (r *Reader) readClass() error {
	className, err := r.readStringWithoutTag()
	if err != nil {
		return err
	}
	count, err := r.ReadInteger(TagOpenbrace)
	if err != nil {
		return err
	}
	fields := make([]string, count)
	for i := 0; i < count; i++ {
		if fields[i], err = r.ReadString(); err != nil {
			return err
		}
	}
	if err = r.CheckTag(TagClosebrace); err != nil {
		return err
	}
	class := ClassManager.GetClass(className)
	if r.classref == nil {
		r.classref = make([]interface{}, 0)
		r.fieldsref = make([][]string, 0)
	}
	var key interface{} = class
	if class == nil {
		key = len(r.classref)
	}
	r.classref = append(r.classref, key)
	r.fieldsref = append(r.fieldsref, fields)
	return nil
}

func (r *Reader) readPointer(v reflect.Value, getValue func() (interface{}, error), setValue func(reflect.Value, interface{})) error {
	if x, err := getValue(); err == nil {
		if reflect.TypeOf(x).Kind() != reflect.Ptr {
			v.Set(reflect.New(v.Type().Elem()))
			setValue(v.Elem(), x)
		} else {
			setValue(v, x)
		}
		return nil
	} else if err == ErrNil {
		v.Set(reflect.Zero(v.Type()))
		return nil
	} else {
		return err
	}
}

func (r *Reader) getInt64() (interface{}, error)          { return r.ReadInt64() }
func (r *Reader) setInt64(v reflect.Value, x interface{}) { v.SetInt(x.(int64)) }
func (r *Reader) readInt64Pointer(v reflect.Value) error {
	return r.readPointer(v, r.getInt64, r.setInt64)
}

func (r *Reader) getUint64() (interface{}, error)          { return r.ReadUint64() }
func (r *Reader) setUint64(v reflect.Value, x interface{}) { v.SetUint(x.(uint64)) }
func (r *Reader) readUint64Pointer(v reflect.Value) error {
	return r.readPointer(v, r.getUint64, r.setUint64)
}

func (r *Reader) getBool() (interface{}, error)          { return r.ReadBool() }
func (r *Reader) setBool(v reflect.Value, x interface{}) { v.SetBool(x.(bool)) }
func (r *Reader) readBoolPointer(v reflect.Value) error {
	return r.readPointer(v, r.getBool, r.setBool)
}

func (r *Reader) getFloat32() (interface{}, error)          { return r.ReadFloat32() }
func (r *Reader) setFloat32(v reflect.Value, x interface{}) { v.SetFloat(float64(x.(float32))) }
func (r *Reader) readFloat32Pointer(v reflect.Value) error {
	return r.readPointer(v, r.getFloat32, r.setFloat32)
}

func (r *Reader) getFloat64() (interface{}, error)          { return r.ReadFloat64() }
func (r *Reader) setFloat64(v reflect.Value, x interface{}) { v.SetFloat(x.(float64)) }
func (r *Reader) readFloat64Pointer(v reflect.Value) error {
	return r.readPointer(v, r.getFloat64, r.setFloat64)
}

func (r *Reader) getBigInt() (interface{}, error)          { return r.ReadBigInt() }
func (r *Reader) setBigInt(v reflect.Value, x interface{}) { v.Set(reflect.ValueOf(x)) }
func (r *Reader) readBigIntPointer(v reflect.Value) error {
	return r.readPointer(v, r.getBigInt, r.setBigInt)
}

func (r *Reader) getDateTime() (interface{}, error)          { return r.ReadDateTime() }
func (r *Reader) setDateTime(v reflect.Value, x interface{}) { v.Set(reflect.ValueOf(x)) }
func (r *Reader) readDateTimePointer(v reflect.Value) error {
	return r.readPointer(v, r.getDateTime, r.setDateTime)
}

func (r *Reader) getString() (interface{}, error)          { return r.ReadString() }
func (r *Reader) setString(v reflect.Value, x interface{}) { v.SetString(x.(string)) }
func (r *Reader) readStringPointer(v reflect.Value) error {
	return r.readPointer(v, r.getString, r.setString)
}

func (r *Reader) getBytes() (interface{}, error)          { return r.ReadBytes() }
func (r *Reader) setBytes(v reflect.Value, x interface{}) { v.Set(reflect.ValueOf(x)) }
func (r *Reader) readBytesPointer(v reflect.Value) error {
	return r.readPointer(v, r.getBytes, r.setBytes)
}

func (r *Reader) getUUID() (interface{}, error)          { return r.ReadUUID() }
func (r *Reader) setUUID(v reflect.Value, x interface{}) { v.Set(reflect.ValueOf(x)) }
func (r *Reader) readUUIDPointer(v reflect.Value) error {
	return r.readPointer(v, r.getUUID, r.setUUID)
}

func (r *Reader) getList() (interface{}, error)          { return r.ReadList() }
func (r *Reader) setList(v reflect.Value, x interface{}) { v.Set(reflect.ValueOf(x)) }
func (r *Reader) readListPointer(v reflect.Value) error {
	return r.readPointer(v, r.getList, r.setList)
}

// private functions

func convertError(tag byte, dst string) error {
	src, err := tagToString(tag)
	if err == nil {
		return errors.New("cannot convert type " + src + " to type " + dst)
	}
	return err
}

func tagToString(tag byte) (string, error) {
	switch tag {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', TagInteger:
		return "int", nil
	case TagLong:
		return "big.Int", nil
	case TagDouble:
		return "float64", nil
	case TagNull:
		return "nil", nil
	case TagEmpty:
		return "empty string", nil
	case TagTrue:
		return "bool true", nil
	case TagFalse:
		return "bool false", nil
	case TagNaN:
		return "NaN", nil
	case TagInfinity:
		return "Infinity", nil
	case TagDate:
		return "time.Time", nil
	case TagTime:
		return "time.Time", nil
	case TagBytes:
		return "[]byte", nil
	case TagUTF8Char:
		return "string", nil
	case TagString:
		return "string", nil
	case TagGuid:
		return "UUID", nil
	case TagList:
		return "slice", nil
	case TagMap:
		return "map", nil
	case TagClass:
		return "struct type", nil
	case TagObject:
		return "struct value", nil
	case TagRef:
		return "value reference", nil
	case TagError:
		return "error", nil
	default:
		return "unknown", unexpectedTag(tag, nil)
	}
}

func stringToBigInt(str string) (*big.Int, error) {
	if bigint, success := new(big.Int).SetString(str, 0); success {
		return bigint, nil
	}
	return big.NewInt(0), errors.New(`cannot convert string "` + str + `" to type big.Int`)
}

func getIndexCache(class reflect.Type) map[string][]int {
	indexCache.RLock()
	indexMap, ok := indexCache.cache[class]
	indexCache.RUnlock()
	if !ok {
		indexCache.Lock()
		if indexCache.cache == nil {
			indexCache.cache = make(map[reflect.Type]map[string][]int)
		}
		indexMap = make(map[string][]int)
		getFieldsFunc(class, func(f *reflect.StructField) {
			tag := ClassManager.GetTag(class)
			if tag == "" {
				indexMap[strings.ToLower(f.Name)] = f.Index
			} else {
				name := strings.SplitN(f.Tag.Get(tag), ",", 2)[0]
				name = strings.TrimSpace(strings.SplitN(name, ">", 2)[0])
				if name == "" {
					indexMap[strings.ToLower(f.Name)] = f.Index
				} else if name != "-" {
					indexMap[strings.ToLower(name)] = f.Index
				}
			}
		})
		indexCache.cache[class] = indexMap
		indexCache.Unlock()
	}
	return indexMap
}
