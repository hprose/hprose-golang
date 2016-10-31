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
 * io/raw_reader.go                                       *
 *                                                        *
 * hprose raw reader for Go.                              *
 *                                                        *
 * LastModified: Oct 15, 2016                             *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package io

import "errors"

// RawReader is the hprose raw reader
type RawReader struct {
	ByteReader
}

// NewRawReader is a constructor for RawReader
func NewRawReader(buf []byte) (reader *RawReader) {
	reader = new(RawReader)
	reader.buf = buf
	return
}

// ReadRaw from stream
func (r *RawReader) ReadRaw() (raw []byte) {
	w := new(ByteWriter)
	r.ReadRawTo(w)
	raw = w.Bytes()
	return
}

// ReadRawTo buffer from stream
func (r *RawReader) ReadRawTo(w *ByteWriter) {
	r.readRaw(w, r.readByte())
}

func (r *RawReader) readRaw(w *ByteWriter, tag byte) {
	w.writeByte(tag)
	switch tag {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
		TagNull, TagEmpty, TagTrue, TagFalse, TagNaN:
	case TagInfinity:
		w.writeByte(r.readByte())
	case TagInteger, TagLong, TagDouble, TagRef:
		r.readNumberRaw(w)
	case TagDate, TagTime:
		r.readDateTimeRaw(w)
	case TagUTF8Char:
		r.readUTF8CharRaw(w)
	case TagBytes:
		r.readBytesRaw(w)
	case TagString:
		r.readStringRaw(w)
	case TagGUID:
		r.readGUIDRaw(w)
	case TagList, TagMap, TagObject:
		r.readComplexRaw(w)
	case TagClass:
		r.readComplexRaw(w)
		r.ReadRawTo(w)
	case TagError:
		r.ReadRawTo(w)
	default:
		unexpectedTag(tag, nil)
	}
	return
}

func (r *RawReader) readNumberRaw(w *ByteWriter) {
	for {
		tag := r.readByte()
		w.writeByte(tag)
		if tag == TagSemicolon {
			return
		}
	}
}

func (r *RawReader) readDateTimeRaw(w *ByteWriter) {
	for {
		tag := r.readByte()
		w.writeByte(tag)
		if tag == TagSemicolon || tag == TagUTC {
			return
		}
	}
}

func (r *RawReader) readUTF8CharRaw(w *ByteWriter) {
	w.write(r.readUTF8Slice(1))
}

func (r *RawReader) readBytesRaw(w *ByteWriter) {
	count := 0
	tag := byte('0')
	for {
		count *= 10
		count += int(tag - '0')
		tag = r.readByte()
		w.writeByte(tag)
		if tag == TagQuote {
			count++
			w.write(r.Next(count))
			return
		}
	}
}

func (r *RawReader) readStringRaw(w *ByteWriter) {
	count := 0
	tag := byte('0')
	for {
		count *= 10
		count += int(tag - '0')
		tag = r.readByte()
		w.writeByte(tag)
		if tag == TagQuote {
			w.write(r.readUTF8Slice(count + 1))
			return
		}
	}
}

func (r *RawReader) readGUIDRaw(w *ByteWriter) {
	w.write(r.Next(38))
}

func (r *RawReader) readComplexRaw(w *ByteWriter) {
	var tag byte
	for tag != TagOpenbrace {
		tag = r.readByte()
		w.writeByte(tag)
	}
	tag = r.readByte()
	for tag != TagClosebrace {
		r.readRaw(w, tag)
		tag = r.readByte()
	}
	w.writeByte(tag)
}

// private functions

func unexpectedTag(tag byte, expectTags []byte) {
	if tag == 0 {
		panic(errors.New("No byte found in stream"))
	} else if expectTags == nil {
		panic(errors.New("Unexpected serialize tag '" + string(rune(tag)) + "' in stream"))
	} else {
		panic(errors.New("Tag '" + string(expectTags) + "' expected, but '" + string(rune(tag)) + "' found in stream"))
	}
}
