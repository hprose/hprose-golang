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
 * io/writer.go                                           *
 *                                                        *
 * hprose writer for Go.                                  *
 *                                                        *
 * LastModified: Oct 19, 2016                             *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package io

import (
	"container/list"
	"math"
	"math/big"
	"reflect"
	"strconv"
	"time"
	"unsafe"

	"github.com/hprose/hprose-golang/util"
)

// Writer is a fine-grained operation struct for Hprose serialization
type Writer struct {
	ByteWriter
	Simple    bool
	structRef map[uintptr]int
	ref       map[uintptr]int
	refCount  int
}

// NewWriter is the constructor for Hprose Writer
func NewWriter(simple bool, buf ...byte) (w *Writer) {
	w = new(Writer)
	w.buf = buf
	w.Simple = simple
	return
}

// Serialize a data v to the writer
func (w *Writer) Serialize(v interface{}) *Writer {
	if v == nil {
		w.WriteNil()
	} else if rv, ok := v.(reflect.Value); ok {
		w.WriteValue(rv)
	} else {
		w.WriteValue(reflect.ValueOf(v))
	}
	return w
}

// WriteValue to the writer
func (w *Writer) WriteValue(v reflect.Value) {
	valueEncoders[v.Kind()](w, v)
}

// WriteNil to the writer
func (w *Writer) WriteNil() {
	w.writeByte(TagNull)
}

// WriteBool to the writer
func (w *Writer) WriteBool(b bool) {
	if b {
		w.writeByte(TagTrue)
	} else {
		w.writeByte(TagFalse)
	}
}

// WriteInt to the writer
func (w *Writer) WriteInt(i int64) {
	if i >= 0 && i <= 9 {
		w.writeByte(byte('0' + i))
		return
	}
	if i >= math.MinInt32 && i <= math.MaxInt32 {
		w.writeByte(TagInteger)
	} else {
		w.writeByte(TagLong)
	}
	var buf [20]byte
	w.write(util.GetIntBytes(buf[:], i))
	w.writeByte(TagSemicolon)
}

// WriteUint to the writer
func (w *Writer) WriteUint(i uint64) {
	if i <= 9 {
		w.writeByte(byte('0' + i))
		return
	}
	if i <= math.MaxInt32 {
		w.writeByte(TagInteger)
	} else {
		w.writeByte(TagLong)
	}
	var buf [20]byte
	w.write(util.GetUintBytes(buf[:], i))
	w.writeByte(TagSemicolon)
}

// WriteFloat to the writer
func (w *Writer) WriteFloat(f float64, bitSize int) {
	if f != f {
		w.writeByte(TagNaN)
		return
	}
	if f > math.MaxFloat64 {
		w.write([]byte{TagInfinity, TagPos})
		return
	}
	if f < -math.MaxFloat64 {
		w.write([]byte{TagInfinity, TagNeg})
		return
	}
	w.writeByte(TagDouble)
	var buf [64]byte
	w.write(strconv.AppendFloat(buf[:0], f, 'g', -1, bitSize))
	w.writeByte(TagSemicolon)
}

// WriteComplex64 to the writer
func (w *Writer) WriteComplex64(c complex64) {
	if imag(c) == 0 {
		w.WriteFloat(float64(real(c)), 32)
		return
	}
	setWriterRef(w, nil)
	writeListHeader(w, 2)
	w.WriteFloat(float64(real(c)), 32)
	w.WriteFloat(float64(imag(c)), 32)
	writeListFooter(w)
}

// WriteComplex128 to the writer
func (w *Writer) WriteComplex128(c complex128) {
	if imag(c) == 0 {
		w.WriteFloat(real(c), 64)
		return
	}
	setWriterRef(w, nil)
	writeListHeader(w, 2)
	w.WriteFloat(real(c), 64)
	w.WriteFloat(imag(c), 64)
	writeListFooter(w)
}

// WriteString to the writer
func (w *Writer) WriteString(str string) {
	length := util.UTF16Length(str)
	switch {
	case length == 0:
		w.writeByte(TagEmpty)
	case length < 0:
		w.WriteBytes(*(*[]byte)(unsafe.Pointer(&str)))
	case length == 1:
		w.writeByte(TagUTF8Char)
		w.writeString(str)
	default:
		setWriterRef(w, nil)
		writeString(w, str, length)
	}
}

// WriteBytes to the writer
func (w *Writer) WriteBytes(bytes []byte) {
	setWriterRef(w, nil)
	writeBytes(w, bytes)
}

// WriteBigInt to the writer
func (w *Writer) WriteBigInt(bi *big.Int) {
	w.writeByte(TagLong)
	w.writeString(bi.String())
	w.writeByte(TagSemicolon)
}

// WriteBigRat to the writer
func (w *Writer) WriteBigRat(br *big.Rat) {
	if br.IsInt() {
		w.WriteBigInt(br.Num())
	} else {
		str := br.String()
		setWriterRef(w, nil)
		writeString(w, str, len(str))
	}
}

// WriteBigFloat to the writer
func (w *Writer) WriteBigFloat(bf *big.Float) {
	w.writeByte(TagDouble)
	var buf [64]byte
	w.write(bf.Append(buf[:0], 'g', -1))
	w.writeByte(TagSemicolon)
}

func writeDate(w *Writer, buf []byte, year, month, day int) {
	w.writeByte(TagDate)
	w.write(util.GetDateBytes(buf, year, int(month), day))
}

func writeTime(w *Writer, buf []byte, hour, min, sec, nsec int) {
	w.writeByte(TagTime)
	w.write(util.GetTimeBytes(buf, hour, min, sec))
	if nsec > 0 {
		w.writeByte(TagPoint)
		w.write(util.GetNsecBytes(buf, nsec))
	}
}

// WriteTime to the writer
func (w *Writer) WriteTime(t *time.Time) {
	ptr := unsafe.Pointer(t)
	if writeRef(w, ptr) {
		return
	}
	setWriterRef(w, ptr)
	year, month, day := t.Date()
	hour, min, sec := t.Clock()
	nsec := t.Nanosecond()
	buf := make([]byte, 9)
	if hour == 0 && min == 0 && sec == 0 && nsec == 0 {
		writeDate(w, buf, year, int(month), day)
	} else if year == 1970 && month == 1 && day == 1 {
		writeTime(w, buf, hour, min, sec, nsec)
	} else {
		writeDate(w, buf, year, int(month), day)
		writeTime(w, buf, hour, min, sec, nsec)
	}
	loc := TagSemicolon
	if t.Location() == time.UTC {
		loc = TagUTC
	}
	w.writeByte(loc)
}

// WriteList to the writer
func (w *Writer) WriteList(lst *list.List) {
	ptr := unsafe.Pointer(lst)
	if writeRef(w, ptr) {
		return
	}
	setWriterRef(w, ptr)
	count := lst.Len()
	if count == 0 {
		writeEmptyList(w)
		return
	}
	writeListHeader(w, count)
	for e := lst.Front(); e != nil; e = e.Next() {
		w.Serialize(e.Value)
	}
	writeListFooter(w)
}

// WriteTuple to the writer
func (w *Writer) WriteTuple(tuple ...interface{}) {
	setWriterRef(w, nil)
	count := len(tuple)
	if count == 0 {
		writeEmptyList(w)
		return
	}
	writeListHeader(w, count)
	for _, v := range tuple {
		w.Serialize(v)
	}
	writeListFooter(w)
}

// WriteSlice to the writer
func (w *Writer) WriteSlice(slice []reflect.Value) {
	setWriterRef(w, nil)
	count := len(slice)
	if count == 0 {
		writeEmptyList(w)
		return
	}
	writeListHeader(w, count)
	for i := range slice {
		w.WriteValue(slice[i])
	}
	writeListFooter(w)
}

// WriteStringSlice to the writer
func (w *Writer) WriteStringSlice(slice []string) {
	setWriterRef(w, nil)
	count := len(slice)
	if count == 0 {
		writeEmptyList(w)
		return
	}
	writeListHeader(w, count)
	stringSliceEncoder(w, slice)
	writeListFooter(w)
}

// Reset the reference counter
func (w *Writer) Reset() {
	if w.structRef != nil {
		for k := range w.structRef {
			delete(w.structRef, k)
		}
	}
	if w.Simple {
		return
	}
	w.refCount = 0
	if w.ref != nil {
		for k := range w.ref {
			delete(w.ref, k)
		}
	}
}

// private functions

func writeRef(w *Writer, ref unsafe.Pointer) bool {
	if w.Simple {
		return false
	}
	if w.ref == nil {
		w.ref = map[uintptr]int{}
	}
	n, found := w.ref[uintptr(ref)]
	if found {
		w.writeByte(TagRef)
		var buf [20]byte
		w.write(util.GetIntBytes(buf[:], int64(n)))
		w.writeByte(TagSemicolon)
	}
	return found
}

func setWriterRef(w *Writer, ref unsafe.Pointer) {
	if w.Simple {
		return
	}
	if ref != nil {
		if w.ref == nil {
			w.ref = map[uintptr]int{}
		}
		w.ref[uintptr(ref)] = w.refCount
	}
	w.refCount++
}

func writeString(w *Writer, str string, length int) {
	w.writeByte(TagString)
	var buf [20]byte
	w.write(util.GetIntBytes(buf[:], int64(length)))
	w.writeByte(TagQuote)
	w.writeString(str)
	w.writeByte(TagQuote)
}

func writeBytes(w *Writer, bytes []byte) {
	count := len(bytes)
	if count == 0 {
		w.write([]byte{TagBytes, TagQuote, TagQuote})
		return
	}
	w.writeByte(TagBytes)
	var buf [20]byte
	w.write(util.GetIntBytes(buf[:], int64(count)))
	w.writeByte(TagQuote)
	w.write(bytes)
	w.writeByte(TagQuote)
}

func writeListHeader(w *Writer, count int) {
	w.writeByte(TagList)
	var buf [20]byte
	w.write(util.GetIntBytes(buf[:], int64(count)))
	w.writeByte(TagOpenbrace)
}

func writeListBody(w *Writer, list reflect.Value, count int) {
	for i := 0; i < count; i++ {
		e := list.Index(i)
		valueEncoders[e.Kind()](w, e)
	}
}

func writeListFooter(w *Writer) {
	w.writeByte(TagClosebrace)
}

func writeEmptyList(w *Writer) {
	w.write([]byte{TagList, TagOpenbrace, TagClosebrace})
}

func writeArray(w *Writer, v reflect.Value) {
	st := reflect.SliceOf(v.Type().Elem())
	sliceType := (*emptyInterface)(unsafe.Pointer(&st)).ptr
	count := v.Len()
	if sliceType == bytesType {
		sliceHeader := reflect.SliceHeader{
			Data: (*emptyInterface)(unsafe.Pointer(&v)).ptr,
			Len:  count,
			Cap:  count,
		}
		writeBytes(w, *(*[]byte)(unsafe.Pointer(&sliceHeader)))
		return
	}
	if count == 0 {
		writeEmptyList(w)
		return
	}
	writeListHeader(w, count)
	encoder := sliceBodyEncoders[sliceType]
	if encoder != nil {
		var slice interface{}
		sliceStruct := (*emptyInterface)(unsafe.Pointer(&slice))
		sliceStruct.typ = sliceType
		sliceStruct.ptr = uintptr(unsafe.Pointer(&reflect.SliceHeader{
			Data: (*emptyInterface)(unsafe.Pointer(&v)).ptr,
			Len:  count,
			Cap:  count,
		}))
		encoder(w, slice)
	} else {
		writeListBody(w, v, count)
	}
	writeListFooter(w)
}

func writeSlice(w *Writer, v reflect.Value) {
	val := (*reflectValue)(unsafe.Pointer(&v))
	if val.typ == bytesType {
		writeBytes(w, v.Bytes())
		return
	}
	count := v.Len()
	if count == 0 {
		writeEmptyList(w)
		return
	}
	writeListHeader(w, count)
	encoder := sliceBodyEncoders[val.typ]
	if encoder != nil {
		encoder(w, *(*interface{})(unsafe.Pointer(&v)))
	} else {
		writeListBody(w, v, count)
	}
	writeListFooter(w)
}

func writeEmptyMap(w *Writer) {
	w.write([]byte{TagMap, TagOpenbrace, TagClosebrace})
}

func writeMapHeader(w *Writer, count int) {
	w.writeByte(TagMap)
	var buf [20]byte
	w.write(util.GetIntBytes(buf[:], int64(count)))
	w.writeByte(TagOpenbrace)
}

func writeMapBody(w *Writer, v reflect.Value) {
	mapType := v.Type()
	keyEncoder := valueEncoders[mapType.Key().Kind()]
	valueEncoder := valueEncoders[mapType.Elem().Kind()]
	keys := v.MapKeys()
	for _, key := range keys {
		keyEncoder(w, key)
		valueEncoder(w, v.MapIndex(key))
	}
}

func writeMapFooter(w *Writer) {
	w.writeByte(TagClosebrace)
}

func writeMap(w *Writer, v reflect.Value) {
	count := v.Len()
	if count == 0 {
		writeEmptyMap(w)
		return
	}
	writeMapHeader(w, count)
	val := (*reflectValue)(unsafe.Pointer(&v))
	mapEncoder := mapBodyEncoders[val.typ]
	if mapEncoder != nil {
		mapEncoder(w, v.Interface())
	} else {
		writeMapBody(w, v)
	}
	writeMapFooter(w)
}

func writeStruct(w *Writer, v reflect.Value) {
	val := (*reflectValue)(unsafe.Pointer(&v))
	cache := getStructCache(v.Type())
	if w.structRef == nil {
		w.structRef = map[uintptr]int{}
	}
	index, found := w.structRef[val.typ]
	if !found {
		w.write(cache.Data)
		if !w.Simple {
			w.refCount += len(cache.Fields)
		}
		index = len(w.structRef)
		w.structRef[val.typ] = index
	}
	ptr := val.ptr
	setWriterRef(w, ptr)
	w.writeByte(TagObject)
	var buf [20]byte
	w.write(util.GetIntBytes(buf[:], int64(index)))
	w.writeByte(TagOpenbrace)
	fields := cache.Fields
	for _, field := range fields {
		valueEncoders[field.Kind](w, v.FieldByIndex(field.Index))
	}
	w.writeByte(TagClosebrace)
}
