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
 * hprose/writer.go                                       *
 *                                                        *
 * hprose Writer for Go.                                  *
 *                                                        *
 * LastModified: Jun 3, 2015                              *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package hprose

import (
	"bytes"
	"container/list"
	"errors"
	"math"
	"math/big"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"
)

const rune3Max = 1<<16 - 1

var serializeType = [...]bool{
	false, // Invalid
	true,  // Bool
	true,  // Int
	true,  // Int8
	true,  // Int16
	true,  // Int32
	true,  // Int64
	true,  // Uint
	true,  // Uint8
	true,  // Uint16
	true,  // Uint32
	true,  // Uint64
	false, // Uintptr
	true,  // Float32
	true,  // Float64
	false, // Complex64
	false, // Complex128
	true,  // Array
	false, // Chan
	false, // Func
	true,  // Interface
	true,  // Map
	true,  // Ptr
	true,  // Slice
	true,  // String
	true,  // Struct
	false, // UnsafePointer
}

var minInt64Buf = [...]byte{
	'-', '9', '2', '2', '3',
	'3', '7', '2', '0', '3',
	'6', '8', '5', '4', '7',
	'7', '5', '8', '0', '8'}

// BufWriter is buffer writer interface, Hprose Writer use it as output stream.
type BufWriter interface {
	Write(p []byte) (n int, err error)
	WriteByte(c byte) error
	WriteRune(r rune) (n int, err error)
	WriteString(s string) (n int, err error)
}

type field struct {
	Name  string
	Index []int
}

type cacheType struct {
	fields            []*field
	hasAnonymousField bool
}

var fieldCache struct {
	sync.RWMutex
	cache map[reflect.Type]*cacheType
}

type writerRefer interface {
	setRef(v interface{})
	writeRef(w *Writer, v interface{}) (success bool, err error)
	resetRef()
}

type fakeWriterRefer struct{}

func (r fakeWriterRefer) setRef(interface{}) {}

func (r fakeWriterRefer) writeRef(w *Writer, v interface{}) (success bool, err error) {
	return false, nil
}

func (r fakeWriterRefer) resetRef() {}

type realWriterRefer struct {
	ref      map[interface{}]int
	refcount int
}

func (r *realWriterRefer) setRef(v interface{}) {
	if r.ref == nil {
		r.ref = make(map[interface{}]int)
		r.refcount = 0
	}
	r.ref[v] = r.refcount
	r.refcount++
}

func (r *realWriterRefer) writeRef(w *Writer, v interface{}) (success bool, err error) {
	if n, found := r.ref[v]; found {
		s := w.Stream
		if err = s.WriteByte(TagRef); err == nil {
			if err = w.writeInt(n); err == nil {
				err = s.WriteByte(TagSemicolon)
			}
		}
		return true, err
	}
	return false, nil
}

func (r *realWriterRefer) resetRef() {
	if r.ref != nil {
		for k := range r.ref {
			delete(r.ref, k)
		}
		r.refcount = 0
	}
}

// Writer is a fine-grained operation struct for Hprose serialization
type Writer struct {
	Stream    BufWriter
	classref  map[string]int
	fieldsref [][]*field
	writerRefer
	numbuf [20]byte
}

// NewWriter is the constructor for Hprose Writer
func NewWriter(stream BufWriter, simple bool) (writer *Writer) {
	writer = new(Writer)
	writer.Stream = stream
	if simple {
		writer.writerRefer = fakeWriterRefer{}
	} else {
		writer.writerRefer = new(realWriterRefer)
	}
	return
}

// Serialize a data to stream
func (w *Writer) Serialize(v interface{}) (err error) {
	return w.fastSerialize(v, reflect.ValueOf(v), 0)
}

// WriteValue to stream
func (w *Writer) WriteValue(v reflect.Value) (err error) {
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return w.WriteNull()
	}
	return w.fastSerialize(v.Interface(), v, 0)
}

// WriteNull to stream
func (w *Writer) WriteNull() error {
	return w.Stream.WriteByte(TagNull)
}

// WriteInt64 to stream
func (w *Writer) WriteInt64(v int64) (err error) {
	s := w.Stream
	if v >= 0 && v <= 9 {
		err = s.WriteByte(byte(v + '0'))
	} else {
		if v >= math.MinInt32 && v <= math.MaxInt32 {
			err = s.WriteByte(TagInteger)
		} else {
			err = s.WriteByte(TagLong)
		}
		if err == nil {
			if err = w.writeInt64(v); err == nil {
				err = s.WriteByte(TagSemicolon)
			}
		}
	}
	return err
}

// WriteUint64 to stream
func (w *Writer) WriteUint64(v uint64) (err error) {
	s := w.Stream
	if v >= 0 && v <= 9 {
		err = s.WriteByte(byte(v + '0'))
	} else {
		if v <= math.MaxInt32 {
			err = s.WriteByte(TagInteger)
		} else {
			err = s.WriteByte(TagLong)
		}
		if err == nil {
			if err = w.writeUint64(v); err == nil {
				err = s.WriteByte(TagSemicolon)
			}
		}
	}
	return err
}

// WriteBigInt to stream
func (w *Writer) WriteBigInt(v *big.Int) (err error) {
	s := w.Stream
	if err = s.WriteByte(TagLong); err == nil {
		if _, err = s.WriteString(v.String()); err == nil {
			err = s.WriteByte(TagSemicolon)
		}
	}
	return err
}

// WriteFloat64 to stream
func (w *Writer) WriteFloat64(v float64) (err error) {
	s := w.Stream
	if math.IsNaN(v) {
		return w.Stream.WriteByte(TagNaN)
	} else if math.IsInf(v, 0) {
		if err = s.WriteByte(TagInfinity); err == nil {
			if v > 0 {
				err = s.WriteByte(TagPos)
			} else {
				err = s.WriteByte(TagNeg)
			}
		}
	} else if err = s.WriteByte(TagDouble); err == nil {
		if _, err = s.WriteString(strconv.FormatFloat(v, 'g', -1, 64)); err == nil {
			err = s.WriteByte(TagSemicolon)
		}
	}
	return err
}

// WriteBool to stream
func (w *Writer) WriteBool(v bool) error {
	s := w.Stream
	if v {
		return s.WriteByte(TagTrue)
	}
	return s.WriteByte(TagFalse)
}

// WriteTime to stream
func (w *Writer) WriteTime(t time.Time) (err error) {
	return w.writeTime(t, t)
}

// WriteString to stream
func (w *Writer) WriteString(str string) (err error) {
	return w.writeString(str, str)
}

// WriteStringWithRef to stream
func (w *Writer) WriteStringWithRef(str string) (err error) {
	s := w.Stream
	if length := len(str); length == 0 {
		err = s.WriteByte(TagEmpty)
	} else if length < 4 && ulen(str) == 1 {
		if err = s.WriteByte(TagUTF8Char); err == nil {
			_, err = s.WriteString(str)
		}
	} else {
		err = w.writeStringWithRef(str, str)
	}
	return err
}

// WriteBytes to stream
func (w *Writer) WriteBytes(bytes []byte) (err error) {
	return w.writeBytes(&bytes, bytes)
}

// WriteBytesWithRef to stream
func (w *Writer) WriteBytesWithRef(bytes []byte) (err error) {
	return w.writeBytesWithRef(&bytes, bytes)
}

// WriteArray to stream
func (w *Writer) WriteArray(v []reflect.Value) (err error) {
	w.setRef(&v)
	s := w.Stream
	count := len(v)
	if err = s.WriteByte(TagList); err == nil {
		if count > 0 {
			if err = w.writeInt(count); err == nil {
				if err = s.WriteByte(TagOpenbrace); err == nil {
					for i := 0; i < count; i++ {
						if err = w.WriteValue(v[i]); err != nil {
							return err
						}
					}
					err = s.WriteByte(TagClosebrace)
				}
			}
		} else if err = s.WriteByte(TagOpenbrace); err == nil {
			err = s.WriteByte(TagClosebrace)
		}
	}
	return err
}

// Reset the serialize reference count
func (w *Writer) Reset() {
	if w.classref != nil {
		for k := range w.classref {
			delete(w.classref, k)
		}
		w.fieldsref = w.fieldsref[:0]
	}
	w.resetRef()
}

// private methods

func (w *Writer) fastSerialize(v interface{}, rv reflect.Value, n int) error {
	switch v := v.(type) {
	case nil:
		return w.WriteNull()
	case int:
		return w.WriteInt64(int64(v))
	case *int:
		return w.WriteInt64(int64(*v))
	case int8:
		return w.WriteInt64(int64(v))
	case *int8:
		return w.WriteInt64(int64(*v))
	case int16:
		return w.WriteInt64(int64(v))
	case *int16:
		return w.WriteInt64(int64(*v))
	case int32:
		return w.WriteInt64(int64(v))
	case *int32:
		return w.WriteInt64(int64(*v))
	case int64:
		return w.WriteInt64(v)
	case *int64:
		return w.WriteInt64(*v)
	case uint:
		return w.WriteUint64(uint64(v))
	case *uint:
		return w.WriteUint64(uint64(*v))
	case uint8:
		return w.WriteUint64(uint64(v))
	case *uint8:
		return w.WriteUint64(uint64(*v))
	case uint16:
		return w.WriteUint64(uint64(v))
	case *uint16:
		return w.WriteUint64(uint64(*v))
	case uint32:
		return w.WriteUint64(uint64(v))
	case *uint32:
		return w.WriteUint64(uint64(*v))
	case uint64:
		return w.WriteUint64(v)
	case *uint64:
		return w.WriteUint64(*v)
	case float32:
		return w.WriteFloat64(float64(v))
	case *float32:
		return w.WriteFloat64(float64(*v))
	case float64:
		return w.WriteFloat64(v)
	case *float64:
		return w.WriteFloat64(*v)
	case bool:
		return w.WriteBool(v)
	case *bool:
		return w.WriteBool(*v)
	case big.Int:
		return w.WriteBigInt(&v)
	case *big.Int:
		return w.WriteBigInt(v)
	case string:
		return w.WriteStringWithRef(v)
	case *string:
		return w.writeStringWithRef(v, *v)
	case time.Time:
		return w.writeTimeWithRef(v, v)
	case *time.Time:
		return w.writeTimeWithRef(v, *v)
	case UUID:
		return w.writeUUIDWithRef(&v, v)
	case *UUID:
		return w.writeUUIDWithRef(v, *v)
	case list.List:
		return w.writeListWithRef(&v, &v)
	case *list.List:
		return w.writeListWithRef(v, v)
	case []byte:
		return w.WriteBytesWithRef(v)
	case *[]byte:
		return w.writeBytesWithRef(v, *v)
	case []int:
		return w.writeIntSliceWithRef(&v, v)
	case *[]int:
		return w.writeIntSliceWithRef(v, *v)
	case []int8:
		return w.writeInt8SliceWithRef(&v, v)
	case *[]int8:
		return w.writeInt8SliceWithRef(v, *v)
	case []int16:
		return w.writeInt16SliceWithRef(&v, v)
	case *[]int16:
		return w.writeInt16SliceWithRef(v, *v)
	case []int32:
		return w.writeInt32SliceWithRef(&v, v)
	case *[]int32:
		return w.writeInt32SliceWithRef(v, *v)
	case []int64:
		return w.writeInt64SliceWithRef(&v, v)
	case *[]int64:
		return w.writeInt64SliceWithRef(v, *v)
	case []uint:
		return w.writeUintSliceWithRef(&v, v)
	case *[]uint:
		return w.writeUintSliceWithRef(v, *v)
	case []uint16:
		return w.writeUint16SliceWithRef(&v, v)
	case *[]uint16:
		return w.writeUint16SliceWithRef(v, *v)
	case []uint32:
		return w.writeUint32SliceWithRef(&v, v)
	case *[]uint32:
		return w.writeUint32SliceWithRef(v, *v)
	case []uint64:
		return w.writeUint64SliceWithRef(&v, v)
	case *[]uint64:
		return w.writeUint64SliceWithRef(v, *v)
	case []float32:
		return w.writeFloat32SliceWithRef(&v, v)
	case *[]float32:
		return w.writeFloat32SliceWithRef(v, *v)
	case []float64:
		return w.writeFloat64SliceWithRef(&v, v)
	case *[]float64:
		return w.writeFloat64SliceWithRef(v, *v)
	case []bool:
		return w.writeBoolSliceWithRef(&v, v)
	case *[]bool:
		return w.writeBoolSliceWithRef(v, *v)
	case []string:
		return w.writeStringSliceWithRef(&v, v)
	case *[]string:
		return w.writeStringSliceWithRef(v, *v)
	case []interface{}:
		return w.writeObjectSliceWithRef(&v, v)
	case *[]interface{}:
		return w.writeObjectSliceWithRef(v, *v)
	case map[string]string:
		return w.writeStringMapWithRef(&v, v)
	case *map[string]string:
		return w.writeStringMapWithRef(v, *v)
	case map[string]interface{}:
		return w.writeStrObjMapWithRef(&v, v)
	case *map[string]interface{}:
		return w.writeStrObjMapWithRef(v, *v)
	case map[interface{}]interface{}:
		return w.writeObjectMapWithRef(&v, v)
	case *map[interface{}]interface{}:
		return w.writeObjectMapWithRef(v, *v)
	}
	return w.slowSerialize(v, rv, n)
}

func (w *Writer) slowSerialize(v interface{}, rv reflect.Value, n int) error {
	kind := rv.Kind()
	switch kind {
	case reflect.Ptr:
		if rv.IsNil() {
			return w.WriteNull()
		}
		return w.slowSerialize(v, rv.Elem(), n+1)
	case reflect.Interface:
		if rv.IsNil() {
			return w.WriteNull()
		}
		return w.fastSerialize(rv.Elem().Interface(), rv.Elem(), n)
	case reflect.Struct:
		switch x := rv.Interface().(type) {
		case big.Int:
			return w.WriteBigInt(&x)
		case time.Time:
			return w.writeTimeWithRef(v, x)
		case list.List:
			return w.writeListWithRef(v, &x)
		default:
			if n == 0 {
				v = &v
			}
			return w.writeObjectWithRef(v, rv)
		}
	case reflect.Map:
		if rv.IsNil() {
			return w.WriteNull()
		}
		if n == 0 {
			v = &v
		}
		return w.writeMapWithRef(v, rv)
	case reflect.Slice:
		if rv.IsNil() {
			return w.WriteNull()
		}
		switch x := rv.Interface().(type) {
		case []byte:
			return w.writeBytesWithRef(v, x)
		case UUID:
			return w.writeUUIDWithRef(v, x)
		default:
			if n == 0 {
				v = &v
			}
			return w.writeSliceWithRef(v, rv)
		}
	case reflect.Array:
		if n == 0 {
			v = &v
		}
		return w.writeSliceWithRef(v, rv)
	case reflect.String:
		return w.writeStringWithRef(v, rv.String())
	case reflect.Bool:
		return w.WriteBool(rv.Bool())
	case reflect.Float32, reflect.Float64:
		return w.WriteFloat64(rv.Float())
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint:
		return w.WriteUint64(rv.Uint())
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
		return w.WriteInt64(rv.Int())
	default:
		return errors.New("This type is not supported: " + reflect.TypeOf(v).String())
	}
}

func (w *Writer) writeTime(v interface{}, t time.Time) (err error) {
	w.setRef(v)
	s := w.Stream
	year, month, day := t.Date()
	hour, min, sec := t.Clock()
	nsec := t.Nanosecond()
	tag := TagSemicolon
	if t.Location() == time.UTC {
		tag = TagUTC
	}
	if hour == 0 && min == 0 && sec == 0 && nsec == 0 {
		if _, err = s.Write(formatDate(year, int(month), day)); err == nil {
			err = s.WriteByte(tag)
		}
	} else if year == 1970 && month == 1 && day == 1 {
		if _, err = s.Write(formatTime(hour, min, sec, nsec)); err == nil {
			err = s.WriteByte(tag)
		}
	} else if _, err = s.Write(formatDate(year, int(month), day)); err == nil {
		if _, err = s.Write(formatTime(hour, min, sec, nsec)); err == nil {
			err = s.WriteByte(tag)
		}
	}
	return err
}

func (w *Writer) writeTimeWithRef(v interface{}, t time.Time) error {
	success, err := w.writeRef(w, v)
	if err == nil && !success {
		return w.writeTime(v, t)
	}
	return err
}

func (w *Writer) writeString(v interface{}, str string) (err error) {
	length := ulen(str)
	if length < 0 {
		return w.writeBytes(v, []byte(str))
	}
	w.setRef(v)
	s := w.Stream
	if err = s.WriteByte(TagString); err == nil {
		if length > 0 {
			if err = w.writeInt(length); err == nil {
				if err = s.WriteByte(TagQuote); err == nil {
					if _, err = s.WriteString(str); err == nil {
						err = s.WriteByte(TagQuote)
					}
				}
			}
		} else if err = s.WriteByte(TagQuote); err == nil {
			err = s.WriteByte(TagQuote)
		}
	}
	return err
}

func (w *Writer) writeStringWithRef(v interface{}, str string) error {
	success, err := w.writeRef(w, v)
	if err == nil && !success {
		return w.writeString(v, str)
	}
	return err
}

func (w *Writer) writeBytes(v interface{}, bytes []byte) (err error) {
	w.setRef(v)
	s := w.Stream
	if err = s.WriteByte(TagBytes); err == nil {
		if length := len(bytes); length > 0 {
			if err = w.writeInt(length); err == nil {
				if err = s.WriteByte(TagQuote); err == nil {
					if _, err = s.Write(bytes); err == nil {
						err = s.WriteByte(TagQuote)
					}
				}
			}
		} else if err = s.WriteByte(TagQuote); err == nil {
			err = s.WriteByte(TagQuote)
		}
	}
	return err
}

func (w *Writer) writeBytesWithRef(v interface{}, bytes []byte) error {
	success, err := w.writeRef(w, v)
	if err == nil && !success {
		return w.writeBytes(v, bytes)
	}
	return err
}

func (w *Writer) writeUUID(v interface{}, uuid UUID) (err error) {
	w.setRef(v)
	s := w.Stream
	if err = s.WriteByte(TagGuid); err == nil {
		if err = s.WriteByte(TagOpenbrace); err == nil {
			if _, err = s.WriteString(uuid.String()); err == nil {
				err = s.WriteByte(TagClosebrace)
			}
		}
	}
	return err
}

func (w *Writer) writeUUIDWithRef(v interface{}, uuid UUID) error {
	success, err := w.writeRef(w, v)
	if err == nil && !success {
		return w.writeUUID(v, uuid)
	}
	return err
}

func (w *Writer) writeList(v interface{}, l *list.List) (err error) {
	w.setRef(v)
	s := w.Stream
	count := l.Len()
	if err = s.WriteByte(TagList); err == nil {
		if count > 0 {
			if err = w.writeInt(count); err == nil {
				if err = s.WriteByte(TagOpenbrace); err == nil {
					for e := l.Front(); e != nil; e = e.Next() {
						if err = w.Serialize(e.Value); err != nil {
							return err
						}
					}
					err = s.WriteByte(TagClosebrace)
				}
			}
		} else if err = s.WriteByte(TagOpenbrace); err == nil {
			err = s.WriteByte(TagClosebrace)
		}
	}
	return err
}

func (w *Writer) writeListWithRef(v interface{}, l *list.List) error {
	success, err := w.writeRef(w, v)
	if err == nil && !success {
		return w.writeList(v, l)
	}
	return err
}

func (w *Writer) writeIntSlice(v interface{}, a []int) (err error) {
	w.setRef(v)
	s := w.Stream
	if err = s.WriteByte(TagList); err == nil {
		if count := len(a); count > 0 {
			if err = w.writeInt(count); err == nil {
				if err = s.WriteByte(TagOpenbrace); err == nil {
					for i := 0; i < count; i++ {
						if err = w.WriteInt64(int64(a[i])); err != nil {
							return err
						}
					}
					err = s.WriteByte(TagClosebrace)
				}
			}
		} else if err = s.WriteByte(TagOpenbrace); err == nil {
			err = s.WriteByte(TagClosebrace)
		}
	}
	return err
}

func (w *Writer) writeIntSliceWithRef(v interface{}, a []int) error {
	success, err := w.writeRef(w, v)
	if err == nil && !success {
		return w.writeIntSlice(v, a)
	}
	return err
}

func (w *Writer) writeInt8Slice(v interface{}, a []int8) (err error) {
	w.setRef(v)
	s := w.Stream
	if err = s.WriteByte(TagList); err == nil {
		if count := len(a); count > 0 {
			if err = w.writeInt(count); err == nil {
				if err = s.WriteByte(TagOpenbrace); err == nil {
					for i := 0; i < count; i++ {
						if err = w.WriteInt64(int64(a[i])); err != nil {
							return err
						}
					}
					err = s.WriteByte(TagClosebrace)
				}
			}
		} else if err = s.WriteByte(TagOpenbrace); err == nil {
			err = s.WriteByte(TagClosebrace)
		}
	}
	return err
}

func (w *Writer) writeInt8SliceWithRef(v interface{}, a []int8) error {
	success, err := w.writeRef(w, v)
	if err == nil && !success {
		return w.writeInt8Slice(v, a)
	}
	return err
}

func (w *Writer) writeInt16Slice(v interface{}, a []int16) (err error) {
	w.setRef(v)
	s := w.Stream
	if err = s.WriteByte(TagList); err == nil {
		if count := len(a); count > 0 {
			if err = w.writeInt(count); err == nil {
				if err = s.WriteByte(TagOpenbrace); err == nil {
					for i := 0; i < count; i++ {
						if err = w.WriteInt64(int64(a[i])); err != nil {
							return err
						}
					}
					err = s.WriteByte(TagClosebrace)
				}
			}
		} else if err = s.WriteByte(TagOpenbrace); err == nil {
			err = s.WriteByte(TagClosebrace)
		}
	}
	return err
}

func (w *Writer) writeInt16SliceWithRef(v interface{}, a []int16) error {
	success, err := w.writeRef(w, v)
	if err == nil && !success {
		return w.writeInt16Slice(v, a)
	}
	return err
}

func (w *Writer) writeInt32Slice(v interface{}, a []int32) (err error) {
	w.setRef(v)
	s := w.Stream
	if err = s.WriteByte(TagList); err == nil {
		if count := len(a); count > 0 {
			if err = w.writeInt(count); err == nil {
				if err = s.WriteByte(TagOpenbrace); err == nil {
					for i := 0; i < count; i++ {
						if err = w.WriteInt64(int64(a[i])); err != nil {
							return err
						}
					}
					err = s.WriteByte(TagClosebrace)
				}
			}
		} else if err = s.WriteByte(TagOpenbrace); err == nil {
			err = s.WriteByte(TagClosebrace)
		}
	}
	return err
}

func (w *Writer) writeInt32SliceWithRef(v interface{}, a []int32) error {
	success, err := w.writeRef(w, v)
	if err == nil && !success {
		return w.writeInt32Slice(v, a)
	}
	return err
}

func (w *Writer) writeInt64Slice(v interface{}, a []int64) (err error) {
	w.setRef(v)
	s := w.Stream
	if err = s.WriteByte(TagList); err == nil {
		if count := len(a); count > 0 {
			if err = w.writeInt(count); err == nil {
				if err = s.WriteByte(TagOpenbrace); err == nil {
					for i := 0; i < count; i++ {
						if err = w.WriteInt64(a[i]); err != nil {
							return err
						}
					}
					err = s.WriteByte(TagClosebrace)
				}
			}
		} else if err = s.WriteByte(TagOpenbrace); err == nil {
			err = s.WriteByte(TagClosebrace)
		}
	}
	return err
}

func (w *Writer) writeInt64SliceWithRef(v interface{}, a []int64) error {
	success, err := w.writeRef(w, v)
	if err == nil && !success {
		return w.writeInt64Slice(v, a)
	}
	return err
}

func (w *Writer) writeUintSlice(v interface{}, a []uint) (err error) {
	w.setRef(v)
	s := w.Stream
	if err = s.WriteByte(TagList); err == nil {
		if count := len(a); count > 0 {
			if err = w.writeInt(count); err == nil {
				if err = s.WriteByte(TagOpenbrace); err == nil {
					for i := 0; i < count; i++ {
						if err = w.WriteUint64(uint64(a[i])); err != nil {
							return err
						}
					}
					err = s.WriteByte(TagClosebrace)
				}
			}
		} else if err = s.WriteByte(TagOpenbrace); err == nil {
			err = s.WriteByte(TagClosebrace)
		}
	}
	return err
}

func (w *Writer) writeUintSliceWithRef(v interface{}, a []uint) error {
	success, err := w.writeRef(w, v)
	if err == nil && !success {
		return w.writeUintSlice(v, a)
	}
	return err
}

func (w *Writer) writeUint16Slice(v interface{}, a []uint16) (err error) {
	w.setRef(v)
	s := w.Stream
	if err = s.WriteByte(TagList); err == nil {
		if count := len(a); count > 0 {
			if err = w.writeInt(count); err == nil {
				if err = s.WriteByte(TagOpenbrace); err == nil {
					for i := 0; i < count; i++ {
						if err = w.WriteUint64(uint64(a[i])); err != nil {
							return err
						}
					}
					err = s.WriteByte(TagClosebrace)
				}
			}
		} else if err = s.WriteByte(TagOpenbrace); err == nil {
			err = s.WriteByte(TagClosebrace)
		}
	}
	return err
}

func (w *Writer) writeUint16SliceWithRef(v interface{}, a []uint16) error {
	success, err := w.writeRef(w, v)
	if err == nil && !success {
		return w.writeUint16Slice(v, a)
	}
	return err
}

func (w *Writer) writeUint32Slice(v interface{}, a []uint32) (err error) {
	w.setRef(v)
	s := w.Stream
	if err = s.WriteByte(TagList); err == nil {
		if count := len(a); count > 0 {
			if err = w.writeInt(count); err == nil {
				if err = s.WriteByte(TagOpenbrace); err == nil {
					for i := 0; i < count; i++ {
						if err = w.WriteUint64(uint64(a[i])); err != nil {
							return err
						}
					}
					err = s.WriteByte(TagClosebrace)
				}
			}
		} else if err = s.WriteByte(TagOpenbrace); err == nil {
			err = s.WriteByte(TagClosebrace)
		}
	}
	return err
}

func (w *Writer) writeUint32SliceWithRef(v interface{}, a []uint32) error {
	success, err := w.writeRef(w, v)
	if err == nil && !success {
		return w.writeUint32Slice(v, a)
	}
	return err
}

func (w *Writer) writeUint64Slice(v interface{}, a []uint64) (err error) {
	w.setRef(v)
	s := w.Stream
	if err = s.WriteByte(TagList); err == nil {
		if count := len(a); count > 0 {
			if err = w.writeInt(count); err == nil {
				if err = s.WriteByte(TagOpenbrace); err == nil {
					for i := 0; i < count; i++ {
						if err = w.WriteUint64(a[i]); err != nil {
							return err
						}
					}
					err = s.WriteByte(TagClosebrace)
				}
			}
		} else if err = s.WriteByte(TagOpenbrace); err == nil {
			err = s.WriteByte(TagClosebrace)
		}
	}
	return err
}

func (w *Writer) writeUint64SliceWithRef(v interface{}, a []uint64) error {
	success, err := w.writeRef(w, v)
	if err == nil && !success {
		return w.writeUint64Slice(v, a)
	}
	return err
}

func (w *Writer) writeFloat32Slice(v interface{}, a []float32) (err error) {
	w.setRef(v)
	s := w.Stream
	if err = s.WriteByte(TagList); err == nil {
		if count := len(a); count > 0 {
			if err = w.writeInt(count); err == nil {
				if err = s.WriteByte(TagOpenbrace); err == nil {
					for i := 0; i < count; i++ {
						if err = w.WriteFloat64(float64(a[i])); err != nil {
							return err
						}
					}
					err = s.WriteByte(TagClosebrace)
				}
			}
		} else if err = s.WriteByte(TagOpenbrace); err == nil {
			err = s.WriteByte(TagClosebrace)
		}
	}
	return err
}

func (w *Writer) writeFloat32SliceWithRef(v interface{}, a []float32) error {
	success, err := w.writeRef(w, v)
	if err == nil && !success {
		return w.writeFloat32Slice(v, a)
	}
	return err
}

func (w *Writer) writeFloat64Slice(v interface{}, a []float64) (err error) {
	w.setRef(v)
	s := w.Stream
	if err = s.WriteByte(TagList); err == nil {
		if count := len(a); count > 0 {
			if err = w.writeInt(count); err == nil {
				if err = s.WriteByte(TagOpenbrace); err == nil {
					for i := 0; i < count; i++ {
						if err = w.WriteFloat64(a[i]); err != nil {
							return err
						}
					}
					err = s.WriteByte(TagClosebrace)
				}
			}
		} else if err = s.WriteByte(TagOpenbrace); err == nil {
			err = s.WriteByte(TagClosebrace)
		}
	}
	return err
}

func (w *Writer) writeFloat64SliceWithRef(v interface{}, a []float64) error {
	success, err := w.writeRef(w, v)
	if err == nil && !success {
		return w.writeFloat64Slice(v, a)
	}
	return err
}

func (w *Writer) writeBoolSlice(v interface{}, a []bool) (err error) {
	w.setRef(v)
	s := w.Stream
	if err = s.WriteByte(TagList); err == nil {
		if count := len(a); count > 0 {
			if err = w.writeInt(count); err == nil {
				if err = s.WriteByte(TagOpenbrace); err == nil {
					for i := 0; i < count; i++ {
						if err = w.WriteBool(a[i]); err != nil {
							return err
						}
					}
					err = s.WriteByte(TagClosebrace)
				}
			}
		} else if err = s.WriteByte(TagOpenbrace); err == nil {
			err = s.WriteByte(TagClosebrace)
		}
	}
	return err
}

func (w *Writer) writeBoolSliceWithRef(v interface{}, a []bool) error {
	success, err := w.writeRef(w, v)
	if err == nil && !success {
		return w.writeBoolSlice(v, a)
	}
	return err
}

func (w *Writer) writeStringSlice(v interface{}, a []string) (err error) {
	w.setRef(v)
	s := w.Stream
	if err = s.WriteByte(TagList); err == nil {
		if count := len(a); count > 0 {
			if err = w.writeInt(count); err == nil {
				if err = s.WriteByte(TagOpenbrace); err == nil {
					for i := 0; i < count; i++ {
						if err = w.WriteStringWithRef(a[i]); err != nil {
							return err
						}
					}
					err = s.WriteByte(TagClosebrace)
				}
			}
		} else if err = s.WriteByte(TagOpenbrace); err == nil {
			err = s.WriteByte(TagClosebrace)
		}
	}
	return err
}

func (w *Writer) writeStringSliceWithRef(v interface{}, a []string) error {
	success, err := w.writeRef(w, v)
	if err == nil && !success {
		return w.writeStringSlice(v, a)
	}
	return err
}

func (w *Writer) writeObjectSlice(v interface{}, a []interface{}) (err error) {
	w.setRef(v)
	s := w.Stream
	if err = s.WriteByte(TagList); err == nil {
		if count := len(a); count > 0 {
			if err = w.writeInt(count); err == nil {
				if err = s.WriteByte(TagOpenbrace); err == nil {
					for i := 0; i < count; i++ {
						if err = w.Serialize(a[i]); err != nil {
							return err
						}
					}
					err = s.WriteByte(TagClosebrace)
				}
			}
		} else if err = s.WriteByte(TagOpenbrace); err == nil {
			err = s.WriteByte(TagClosebrace)
		}
	}
	return err
}

func (w *Writer) writeObjectSliceWithRef(v interface{}, a []interface{}) error {
	success, err := w.writeRef(w, v)
	if err == nil && !success {
		return w.writeObjectSlice(v, a)
	}
	return err
}

func (w *Writer) writeSlice(v interface{}, rv reflect.Value) (err error) {
	w.setRef(v)
	s := w.Stream
	if err = s.WriteByte(TagList); err == nil {
		if count := rv.Len(); count > 0 {
			if err = w.writeInt(count); err == nil {
				if err = s.WriteByte(TagOpenbrace); err == nil {
					for i := 0; i < count; i++ {
						if err = w.WriteValue(rv.Index(i)); err != nil {
							return err
						}
					}
					err = s.WriteByte(TagClosebrace)
				}
			}
		} else if err = s.WriteByte(TagOpenbrace); err == nil {
			err = s.WriteByte(TagClosebrace)
		}
	}
	return err
}

func (w *Writer) writeSliceWithRef(v interface{}, rv reflect.Value) error {
	success, err := w.writeRef(w, v)
	if err == nil && !success {
		return w.writeSlice(v, rv)
	}
	return err
}

func (w *Writer) writeStringMap(v interface{}, m map[string]string) (err error) {
	w.setRef(v)
	s := w.Stream
	if err = s.WriteByte(TagMap); err == nil {
		if count := len(m); count > 0 {
			if err = w.writeInt(count); err == nil {
				if err = s.WriteByte(TagOpenbrace); err == nil {
					for k, v := range m {
						if err = w.writeStringWithRef(k, k); err != nil {
							return err
						}
						if err = w.writeStringWithRef(v, v); err != nil {
							return err
						}
					}
					err = s.WriteByte(TagClosebrace)
				}
			}
		} else if err = s.WriteByte(TagOpenbrace); err == nil {
			err = s.WriteByte(TagClosebrace)
		}
	}
	return err
}

func (w *Writer) writeStringMapWithRef(v interface{}, m map[string]string) error {
	success, err := w.writeRef(w, v)
	if err == nil && !success {
		return w.writeStringMap(v, m)
	}
	return err
}

func (w *Writer) writeStrObjMap(v interface{}, m map[string]interface{}) (err error) {
	w.setRef(v)
	s := w.Stream
	if err = s.WriteByte(TagMap); err == nil {
		if count := len(m); count > 0 {
			if err = w.writeInt(count); err == nil {
				if err = s.WriteByte(TagOpenbrace); err == nil {
					for k, v := range m {
						if err = w.writeStringWithRef(k, k); err != nil {
							return err
						}
						if err = w.Serialize(v); err != nil {
							return err
						}
					}
					err = s.WriteByte(TagClosebrace)
				}
			}
		} else if err = s.WriteByte(TagOpenbrace); err == nil {
			err = s.WriteByte(TagClosebrace)
		}
	}
	return err
}

func (w *Writer) writeStrObjMapWithRef(v interface{}, m map[string]interface{}) error {
	success, err := w.writeRef(w, v)
	if err == nil && !success {
		return w.writeStrObjMap(v, m)
	}
	return err
}

func (w *Writer) writeObjectMap(v interface{}, m map[interface{}]interface{}) (err error) {
	w.setRef(v)
	s := w.Stream
	if err = s.WriteByte(TagMap); err == nil {
		if count := len(m); count > 0 {
			if err = w.writeInt(count); err == nil {
				if err = s.WriteByte(TagOpenbrace); err == nil {
					for k, v := range m {
						if err = w.Serialize(k); err != nil {
							return err
						}
						if err = w.Serialize(v); err != nil {
							return err
						}
					}
					err = s.WriteByte(TagClosebrace)
				}
			}
		} else if err = s.WriteByte(TagOpenbrace); err == nil {
			err = s.WriteByte(TagClosebrace)
		}
	}
	return err
}

func (w *Writer) writeObjectMapWithRef(v interface{}, m map[interface{}]interface{}) error {
	success, err := w.writeRef(w, v)
	if err == nil && !success {
		return w.writeObjectMap(v, m)
	}
	return err
}

func (w *Writer) writeMap(v interface{}, rv reflect.Value) (err error) {
	w.setRef(v)
	s := w.Stream
	if err = s.WriteByte(TagMap); err == nil {
		if count := rv.Len(); count > 0 {
			if err = w.writeInt(count); err == nil {
				if err = s.WriteByte(TagOpenbrace); err == nil {
					keys := rv.MapKeys()
					for i := range keys {
						if err = w.WriteValue(keys[i]); err != nil {
							return err
						}
						if err = w.WriteValue(rv.MapIndex(keys[i])); err != nil {
							return err
						}
					}
					err = s.WriteByte(TagClosebrace)
				}
			}
		} else if err = s.WriteByte(TagOpenbrace); err == nil {
			err = s.WriteByte(TagClosebrace)
		}
	}
	return err
}

func (w *Writer) writeMapWithRef(v interface{}, rv reflect.Value) error {
	success, err := w.writeRef(w, v)
	if err == nil && !success {
		return w.writeMap(v, rv)
	}
	return err
}

func (w *Writer) writeObjectAsMap(v reflect.Value, fields []*field) (err error) {
	s := w.Stream
	buf := new(bytes.Buffer)
	w.Stream = buf
	count := 0
NEXT:
	for _, f := range fields {
		e := v.Field(f.Index[0])
		n := len(f.Index)
		if n > 1 {
			for i := 1; i < n; i++ {
				if e.Kind() == reflect.Ptr && e.IsNil() {
					continue NEXT
				}
				e = reflect.Indirect(e).Field(f.Index[i])
			}
		}
		if err = w.writeStringWithRef(f.Name, f.Name); err != nil {
			return err
		}
		if err = w.WriteValue(e); err != nil {
			return err
		}
		count++
	}
	w.Stream = s
	if err = s.WriteByte(TagMap); err == nil {
		if count > 0 {
			if err = w.writeInt(count); err == nil {
				if err = s.WriteByte(TagOpenbrace); err == nil {
					buf.WriteTo(s)
					err = s.WriteByte(TagClosebrace)
				}
			}
		} else if err = s.WriteByte(TagOpenbrace); err == nil {
			err = s.WriteByte(TagClosebrace)
		}
	}
	return err
}

func (w *Writer) writeObject(v interface{}, rv reflect.Value) (err error) {
	s := w.Stream
	t := rv.Type()
	classname := ClassManager.GetClassAlias(t)
	if classname == "" {
		classname = t.Name()
		ClassManager.Register(t, classname)
	}
	if w.classref == nil {
		w.classref = make(map[string]int)
		w.fieldsref = make([][]*field, 0)
	}
	index, found := w.classref[classname]
	var fields []*field
	if found {
		fields = w.fieldsref[index]
	} else {
		fieldCache.RLock()
		cache, found := fieldCache.cache[t]
		fieldCache.RUnlock()
		if !found {
			fieldCache.Lock()
			if fieldCache.cache == nil {
				fieldCache.cache = make(map[reflect.Type]*cacheType)
			}
			fields = make([]*field, 0)
			hasAnonymousField := false
			getFieldsFunc(t, func(f *reflect.StructField) {
				if len(f.Index) > 1 {
					hasAnonymousField = true
				}
				tag := ClassManager.GetTag(t)
				if tag == "" {
					fields = append(fields, &field{firstLetterToLower(f.Name), f.Index})
				} else {
					name := strings.SplitN(f.Tag.Get(tag), ",", 2)[0]
					name = strings.TrimSpace(strings.SplitN(name, ">", 2)[0])
					if name == "" {
						fields = append(fields, &field{firstLetterToLower(f.Name), f.Index})
					} else if name != "-" {
						fields = append(fields, &field{name, f.Index})
					}

				}
			})
			cache = &cacheType{fields, hasAnonymousField}
			fieldCache.cache[t] = cache
			fieldCache.Unlock()
		} else {
			fields = cache.fields
		}
		if !cache.hasAnonymousField {
			if index, err = w.writeClass(classname, fields); err != nil {
				return err
			}
		} else {
			w.setRef(v)
			return w.writeObjectAsMap(rv, fields)
		}
	}
	w.setRef(v)
	if err = s.WriteByte(TagObject); err == nil {
		if err = w.writeInt(index); err == nil {
			if err = s.WriteByte(TagOpenbrace); err == nil {
				for i := range fields {
					if err = w.WriteValue(rv.FieldByIndex(fields[i].Index)); err != nil {
						return err
					}
				}
				err = w.Stream.WriteByte(TagClosebrace)
			}
		}
	}
	return err
}

func (w *Writer) writeObjectWithRef(v interface{}, rv reflect.Value) error {
	success, err := w.writeRef(w, v)
	if err == nil && !success {
		return w.writeObject(v, rv)
	}
	return err
}

func (w *Writer) writeClass(classname string, fields []*field) (index int, err error) {
	s := w.Stream
	count := len(fields)
	if err = s.WriteByte(TagClass); err != nil {
		return -1, err
	}
	if err = w.writeInt(ulen(classname)); err != nil {
		return -1, err
	}
	if err = s.WriteByte(TagQuote); err != nil {
		return -1, err
	}
	if _, err = s.WriteString(classname); err != nil {
		return -1, err
	}
	if err = s.WriteByte(TagQuote); err != nil {
		return -1, err
	}
	if count > 0 {
		if err = w.writeInt(count); err != nil {
			return -1, err
		}
		if err = s.WriteByte(TagOpenbrace); err != nil {
			return -1, err
		}
		for i := range fields {
			if err = w.WriteString(fields[i].Name); err != nil {
				return -1, err
			}
		}
		if err = s.WriteByte(TagClosebrace); err != nil {
			return -1, err
		}
	} else {
		if err = s.WriteByte(TagOpenbrace); err != nil {
			return -1, err
		}
		if err = s.WriteByte(TagClosebrace); err != nil {
			return -1, err
		}
	}
	index = len(w.fieldsref)
	w.classref[classname] = index
	w.fieldsref = append(w.fieldsref, fields)
	return index, nil
}

func (w *Writer) writeInt64(i int64) error {
	if i >= 0 && i <= 9 {
		return w.Stream.WriteByte((byte)(i + '0'))
	} else if i == math.MinInt64 {
		_, err := w.Stream.Write(minInt64Buf[:])
		return err
	}
	off := 20
	sign := int64(1)
	if i < 0 {
		sign = -sign
		i = -i
	}
	for i != 0 {
		off--
		w.numbuf[off] = (byte)((i % 10) + '0')
		i /= 10
	}
	if sign == -1 {
		off--
		w.numbuf[off] = '-'
	}
	_, err := w.Stream.Write(w.numbuf[off:])
	return err
}

func (w *Writer) writeUint64(i uint64) error {
	if i >= 0 && i <= 9 {
		return w.Stream.WriteByte((byte)(i + '0'))
	}
	off := 20
	for i != 0 {
		off--
		w.numbuf[off] = (byte)((i % 10) + '0')
		i /= 10
	}
	_, err := w.Stream.Write(w.numbuf[off:])
	return err
}

func (w *Writer) writeInt(i int) error {
	return w.writeInt64(int64(i))
}

// private functions

func ulen(str string) (n int) {
	length := len(str)
	n = length
	p := 0
	for p < length {
		a := str[p]
		if a < 0x80 {
			p++
		} else if (a & 0xE0) == 0xC0 {
			p += 2
			n--
		} else if (a & 0xF0) == 0xE0 {
			p += 3
			n -= 2
		} else if (a * 0xF8) == 0xF0 {
			p += 4
			n -= 2
		} else {
			return -1
		}
	}
	return n
}

func formatDate(year int, month int, day int) []byte {
	var date [9]byte
	date[0] = TagDate
	date[1] = byte('0' + (year / 1000 % 10))
	date[2] = byte('0' + (year / 100 % 10))
	date[3] = byte('0' + (year / 10 % 10))
	date[4] = byte('0' + (year % 10))
	date[5] = byte('0' + (month / 10 % 10))
	date[6] = byte('0' + (month % 10))
	date[7] = byte('0' + (day / 10 % 10))
	date[8] = byte('0' + (day % 10))
	return date[:]
}

func formatTime(hour int, min int, sec int, nsec int) []byte {
	var time [7]byte
	time[0] = TagTime
	time[1] = byte('0' + (hour / 10 % 10))
	time[2] = byte('0' + (hour % 10))
	time[3] = byte('0' + (min / 10 % 10))
	time[4] = byte('0' + (min % 10))
	time[5] = byte('0' + (sec / 10 % 10))
	time[6] = byte('0' + (sec % 10))
	if nsec > 0 {
		if nsec%1000000 == 0 {
			var nanoSecond [4]byte
			nanoSecond[0] = TagPoint
			nanoSecond[1] = (byte)('0' + (nsec / 100000000 % 10))
			nanoSecond[2] = (byte)('0' + (nsec / 10000000 % 10))
			nanoSecond[3] = (byte)('0' + (nsec / 1000000 % 10))
			return append(time[:], nanoSecond[:]...)
		} else if nsec%1000 == 0 {
			var nanoSecond [7]byte
			nanoSecond[0] = TagPoint
			nanoSecond[1] = (byte)('0' + (nsec / 100000000 % 10))
			nanoSecond[2] = (byte)('0' + (nsec / 10000000 % 10))
			nanoSecond[3] = (byte)('0' + (nsec / 1000000 % 10))
			nanoSecond[4] = (byte)('0' + (nsec / 100000 % 10))
			nanoSecond[5] = (byte)('0' + (nsec / 10000 % 10))
			nanoSecond[6] = (byte)('0' + (nsec / 1000 % 10))
			return append(time[:], nanoSecond[:]...)

		} else {
			var nanoSecond [10]byte
			nanoSecond[0] = TagPoint
			nanoSecond[1] = (byte)('0' + (nsec / 100000000 % 10))
			nanoSecond[2] = (byte)('0' + (nsec / 10000000 % 10))
			nanoSecond[3] = (byte)('0' + (nsec / 1000000 % 10))
			nanoSecond[4] = (byte)('0' + (nsec / 100000 % 10))
			nanoSecond[5] = (byte)('0' + (nsec / 10000 % 10))
			nanoSecond[6] = (byte)('0' + (nsec / 1000 % 10))
			nanoSecond[7] = (byte)('0' + (nsec / 100 % 10))
			nanoSecond[8] = (byte)('0' + (nsec / 10 % 10))
			nanoSecond[9] = (byte)('0' + (nsec % 10))
			return append(time[:], nanoSecond[:]...)
		}
	}
	return time[:]
}

func firstLetterToLower(s string) string {
	if s == "" || s[0] < 'A' || s[0] > 'Z' {
		return s
	}
	b := ([]byte)(s)
	b[0] = b[0] - 'A' + 'a'
	return string(b)
}

func getFieldsFunc(class reflect.Type, set func(*reflect.StructField)) {
	count := class.NumField()
	for i := 0; i < count; i++ {
		if f := class.Field(i); serializeType[f.Type.Kind()] {
			if !f.Anonymous {
				b := f.Name[0]
				if 'A' <= b && b <= 'Z' {
					set(&f)
				}
			} else {
				ft := f.Type
				if ft.Name() == "" && ft.Kind() == reflect.Ptr {
					ft = ft.Elem()
				}
				if ft.Kind() == reflect.Struct {
					getAnonymousFieldsFunc(ft, f.Index, set)
				}
			}
		}
	}
}

func getAnonymousFieldsFunc(class reflect.Type, index []int, set func(*reflect.StructField)) {
	count := class.NumField()
	for i := 0; i < count; i++ {
		if f := class.Field(i); serializeType[f.Type.Kind()] {
			f.Index = append(index, f.Index[0])
			if !f.Anonymous {
				b := f.Name[0]
				if 'A' <= b && b <= 'Z' {
					set(&f)
				}
			} else {
				ft := f.Type
				if ft.Name() == "" && ft.Kind() == reflect.Ptr {
					ft = ft.Elem()
				}
				if ft.Kind() == reflect.Struct {
					getAnonymousFieldsFunc(ft, f.Index, set)
				}
			}
		}
	}
}
