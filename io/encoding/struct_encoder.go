/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| io/encoding/encoder.go                                   |
|                                                          |
| LastModified: Mar 20, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"

	"github.com/hprose/hprose-golang/v3/io"
	"github.com/modern-go/reflect2"
)

type encodeFunc func(enc *Encoder, v interface{}) error

type field struct {
	typ    reflect2.Type
	field  reflect2.StructField
	encode encodeFunc
}

// StructEncoder is the implementation of ValueEncoder for struct/*struct.
type StructEncoder struct {
	fields   []field
	metadata []byte
}

// Encode writes the hprose encoding of v to stream
// if v is already written to stream, it will writes it as reference
func (valenc *StructEncoder) Encode(enc *Encoder, v interface{}) (err error) {
	if reflect.TypeOf(v).Kind() == reflect.Struct {
		return valenc.Write(enc, v)
	}
	if reflect2.IsNil(v) {
		return WriteNil(enc.Writer)
	}
	var ok bool
	if ok, err = enc.WriteReference(v); !ok && err == nil {
		err = valenc.Write(enc, v)
	}
	return
}

// Write writes the hprose encoding of v to stream
// if v is already written to stream, it will writes it as value
func (valenc *StructEncoder) Write(enc *Encoder, v interface{}) (err error) {
	t := reflect.TypeOf(v)
	st := t
	if t.Kind() == reflect.Ptr {
		st = t.Elem()
	}
	writer := enc.Writer
	fields := valenc.fields
	n := len(fields)
	var r int
	r, err = enc.WriteStructType(st, func() (err error) {
		enc.AddReferenceCount(n)
		_, err = writer.Write(valenc.metadata)
		return
	})
	if err == nil {
		if t.Kind() == reflect.Ptr {
			enc.SetReference(v)
		} else {
			enc.AddReferenceCount(1)
		}
		p := reflect2.PtrOf(v)
		err = WriteObjectHead(writer, r)
		for i := 0; i < n && err == nil; i++ {
			err = fields[i].encode(enc, fields[i].typ.UnsafeIndirect(fields[i].field.UnsafeGet(p)))
		}
		if err == nil {
			err = WriteFoot(writer)
		}
	}
	return
}

func boolEncode(enc *Encoder, v interface{}) error {
	return WriteBool(enc.Writer, *(*bool)(reflect2.PtrOf(v)))
}

func intEncode(enc *Encoder, v interface{}) error {
	return WriteInt(enc.Writer, *(*int)(reflect2.PtrOf(v)))
}

func int8Encode(enc *Encoder, v interface{}) error {
	return WriteInt8(enc.Writer, *(*int8)(reflect2.PtrOf(v)))
}

func int16Encode(enc *Encoder, v interface{}) error {
	return WriteInt16(enc.Writer, *(*int16)(reflect2.PtrOf(v)))
}

func int32Encode(enc *Encoder, v interface{}) error {
	return WriteInt32(enc.Writer, *(*int32)(reflect2.PtrOf(v)))
}

func int64Encode(enc *Encoder, v interface{}) error {
	return WriteInt64(enc.Writer, *(*int64)(reflect2.PtrOf(v)))
}

func uintEncode(enc *Encoder, v interface{}) error {
	return WriteUint(enc.Writer, *(*uint)(reflect2.PtrOf(v)))
}

func uint8Encode(enc *Encoder, v interface{}) error {
	return WriteUint8(enc.Writer, *(*uint8)(reflect2.PtrOf(v)))
}

func uint16Encode(enc *Encoder, v interface{}) error {
	return WriteUint16(enc.Writer, *(*uint16)(reflect2.PtrOf(v)))
}

func uint32Encode(enc *Encoder, v interface{}) error {
	return WriteUint32(enc.Writer, *(*uint32)(reflect2.PtrOf(v)))
}

func uint64Encode(enc *Encoder, v interface{}) error {
	return WriteUint64(enc.Writer, *(*uint64)(reflect2.PtrOf(v)))
}

func float32Encode(enc *Encoder, v interface{}) error {
	return WriteFloat32(enc.Writer, *(*float32)(reflect2.PtrOf(v)))
}

func float64Encode(enc *Encoder, v interface{}) error {
	return WriteFloat64(enc.Writer, *(*float64)(reflect2.PtrOf(v)))
}

func complex64Encode(enc *Encoder, v interface{}) error {
	return WriteComplex64(enc, *(*complex64)(reflect2.PtrOf(v)))
}

func complex128Encode(enc *Encoder, v interface{}) error {
	return WriteComplex128(enc, *(*complex128)(reflect2.PtrOf(v)))
}

func stringEncode(enc *Encoder, v interface{}) error {
	return EncodeString(enc, *(*string)(reflect2.PtrOf(v)))
}

func arrayEncode(enc *Encoder, v interface{}) error {
	return WriteArray(enc, v)
}

func mapEncode(enc *Encoder, v interface{}) error {
	if reflect.ValueOf(v).IsNil() {
		return WriteNil(enc.Writer)
	}
	return WriteMap(enc, v)
}

func sliceEncode(enc *Encoder, v interface{}) error {
	if reflect.ValueOf(v).IsNil() {
		return WriteNil(enc.Writer)
	}
	return WriteSlice(enc, v)
}

func interfaceEncode(enc *Encoder, v interface{}) error {
	return enc.Encode(v)
}

func boolPtrEncode(enc *Encoder, v interface{}) error {
	p := (*bool)(reflect2.PtrOf(v))
	if p == nil {
		return WriteNil(enc.Writer)
	}
	return WriteBool(enc.Writer, *p)
}

func intPtrEncode(enc *Encoder, v interface{}) error {
	p := (*int)(reflect2.PtrOf(v))
	if p == nil {
		return WriteNil(enc.Writer)
	}
	return WriteInt(enc.Writer, *p)
}

func int8PtrEncode(enc *Encoder, v interface{}) error {
	p := (*int8)(reflect2.PtrOf(v))
	if p == nil {
		return WriteNil(enc.Writer)
	}
	return WriteInt8(enc.Writer, *p)
}

func int16PtrEncode(enc *Encoder, v interface{}) error {
	p := (*int16)(reflect2.PtrOf(v))
	if p == nil {
		return WriteNil(enc.Writer)
	}
	return WriteInt16(enc.Writer, *p)
}

func int32PtrEncode(enc *Encoder, v interface{}) error {
	p := (*int32)(reflect2.PtrOf(v))
	if p == nil {
		return WriteNil(enc.Writer)
	}
	return WriteInt32(enc.Writer, *p)
}

func int64PtrEncode(enc *Encoder, v interface{}) error {
	p := (*int64)(reflect2.PtrOf(v))
	if p == nil {
		return WriteNil(enc.Writer)
	}
	return WriteInt64(enc.Writer, *p)
}

func uintPtrEncode(enc *Encoder, v interface{}) error {
	p := (*uint)(reflect2.PtrOf(v))
	if p == nil {
		return WriteNil(enc.Writer)
	}
	return WriteUint(enc.Writer, *p)
}

func uint8PtrEncode(enc *Encoder, v interface{}) error {
	p := (*uint8)(reflect2.PtrOf(v))
	if p == nil {
		return WriteNil(enc.Writer)
	}
	return WriteUint8(enc.Writer, *p)
}

func uint16PtrEncode(enc *Encoder, v interface{}) error {
	p := (*uint16)(reflect2.PtrOf(v))
	if p == nil {
		return WriteNil(enc.Writer)
	}
	return WriteUint16(enc.Writer, *p)
}

func uint32PtrEncode(enc *Encoder, v interface{}) error {
	p := (*uint32)(reflect2.PtrOf(v))
	if p == nil {
		return WriteNil(enc.Writer)
	}
	return WriteUint32(enc.Writer, *p)
}

func uint64PtrEncode(enc *Encoder, v interface{}) error {
	p := (*uint64)(reflect2.PtrOf(v))
	if p == nil {
		return WriteNil(enc.Writer)
	}
	return WriteUint64(enc.Writer, *p)
}

func float32PtrEncode(enc *Encoder, v interface{}) error {
	p := (*float32)(reflect2.PtrOf(v))
	if p == nil {
		return WriteNil(enc.Writer)
	}
	return WriteFloat32(enc.Writer, *p)
}

func float64PtrEncode(enc *Encoder, v interface{}) error {
	p := (*float64)(reflect2.PtrOf(v))
	if p == nil {
		return WriteNil(enc.Writer)
	}
	return WriteFloat64(enc.Writer, *p)
}

func complex64PtrEncode(enc *Encoder, v interface{}) error {
	p := (*complex64)(reflect2.PtrOf(v))
	if p == nil {
		return WriteNil(enc.Writer)
	}
	return WriteComplex64(enc, *p)
}

func complex128PtrEncode(enc *Encoder, v interface{}) error {
	p := (*complex128)(reflect2.PtrOf(v))
	if p == nil {
		return WriteNil(enc.Writer)
	}
	return WriteComplex128(enc, *p)
}

func stringPtrEncode(enc *Encoder, v interface{}) error {
	p := (*string)(reflect2.PtrOf(v))
	if p == nil {
		return WriteNil(enc.Writer)
	}
	return EncodeString(enc, *p)
}

func arrayPtrEncode(enc *Encoder, v interface{}) error {
	if reflect2.IsNil(v) {
		return WriteNil(enc.Writer)
	}
	return arrayEncoder.Encode(enc, v)
}

func mapPtrEncode(enc *Encoder, v interface{}) error {
	if rv := reflect.ValueOf(v); rv.IsNil() || rv.Elem().IsNil() {
		return WriteNil(enc.Writer)
	}
	return mapEncoder.Encode(enc, v)
}

func slicePtrEncode(enc *Encoder, v interface{}) error {
	if rv := reflect.ValueOf(v); rv.IsNil() || rv.Elem().IsNil() {
		return WriteNil(enc.Writer)
	}
	return sliceEncoder.Encode(enc, v)
}

func interfacePtrEncode(enc *Encoder, v interface{}) error {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return WriteNil(enc.Writer)
	}
	return enc.Encode(rv.Elem().Interface())
}

func ptrEncode(enc *Encoder, v interface{}) error {
	if reflect2.IsNil(v) {
		return WriteNil(enc.Writer)
	}
	return ptrEncoder.Encode(enc, v)
}

func writeName(writer io.Writer, s string) (err error) {
	length := utf16Length(s)
	if length < 0 {
		return ErrInvalidUTF8
	}
	return writeBinary(writer, reflect2.UnsafeCastString(s), length)
}

func stripOptions(tag string) string {
	i := strings.Index(tag, ",")
	if i < 0 {
		return tag
	}
	return string(tag[:i])
}

func _fieldAlias(tag reflect.StructTag, tagname string) string {
	return strings.Trim(stripOptions(tag.Get(tagname)), " ")
}

func fieldAlias(tag reflect.StructTag, name string, tagnames []string) string {
	alias := _fieldAlias(tag, defaultTagName)
	if alias != "" {
		return alias
	}
	for _, tagname := range tagnames {
		if tagname != "" && tagname != defaultTagName {
			alias = _fieldAlias(tag, tagname)
			if alias != "" {
				return alias
			}
		}
	}
	if name[0] >= 'A' && name[0] <= 'Z' {
		name = string(name[0]-'A'+'a') + name[1:]
	}
	return name
}

func getEncode(t reflect.Type) encodeFunc {
	if f := getOtherEncode(t); f != nil {
		return f
	}
	switch t.Kind() {
	case reflect.Int:
		return intEncode
	case reflect.Int8:
		return int8Encode
	case reflect.Int16:
		return int16Encode
	case reflect.Int32:
		return int32Encode
	case reflect.Int64:
		return int64Encode
	case reflect.Uint:
		return uintEncode
	case reflect.Uint8:
		return uint8Encode
	case reflect.Uint16:
		return uint16Encode
	case reflect.Uint32:
		return uint32Encode
	case reflect.Uint64, reflect.Uintptr:
		return uint64Encode
	case reflect.Bool:
		return boolEncode
	case reflect.Float32:
		return float32Encode
	case reflect.Float64:
		return float64Encode
	case reflect.Complex64:
		return complex64Encode
	case reflect.Complex128:
		return complex128Encode
	case reflect.Array:
		return arrayEncode
	case reflect.Interface:
		return interfaceEncode
	case reflect.Map:
		return mapEncode
	case reflect.Ptr:
		return getPtrEncode(t.Elem())
	case reflect.Slice:
		return sliceEncode
	case reflect.String:
		return stringEncode
	case reflect.Struct:
		return getStructEncode(t)
	}
	return nil
}

func getFields(t reflect2.StructType, tagnames []string, mapping map[string]bool, names []string, fields []field) ([]string, []field) {
	n := t.NumField()
	for i := 0; i < n; i++ {
		f := t.Field(i)
		ft := f.Type()
		kind := ft.Kind()

		switch kind {
		case reflect.Func, reflect.Chan, reflect.UnsafePointer:
			continue
		case reflect.Struct:
			if f.Anonymous() {
				names, fields = getFields(ft.(reflect2.StructType), tagnames, mapping, names, fields)
				continue
			}
		}

		if f.PkgPath() != "" {
			continue
		}

		name := fieldAlias(f.Tag(), f.Name(), tagnames)
		if name == "-" {
			continue
		}
		if mapping[name] {
			panic(fmt.Sprintf("ambiguous fields with the same name or alias: %s", name))
		}

		var field field
		field.typ = ft
		field.field = f
		typ := ft.Type1()
		if field.encode = getEncode(typ); field.encode == nil {
			continue
		}

		mapping[name] = true
		names = append(names, name)
		fields = append(fields, field)
	}
	return names, fields
}

func newStructEncoder(t reflect.Type, name string, tagnames []string) ValueEncoder {
	encoder := &StructEncoder{}
	registerEncoder(t, encoder)
	var names []string
	names, encoder.fields = getFields(reflect2.Type2(t).(reflect2.StructType), tagnames, map[string]bool{}, nil, nil)
	n := len(names)
	buffer := &bytes.Buffer{}
	buffer.WriteByte(io.TagClass)
	writeName(buffer, name)
	if n > 0 {
		writeUint64(buffer, uint64(n))
	}
	buffer.WriteByte(io.TagOpenbrace)
	for i := 0; i < n; i++ {
		buffer.WriteByte(io.TagString)
		writeName(buffer, names[i])
	}
	buffer.WriteByte(io.TagClosebrace)
	encoder.metadata = buffer.Bytes()
	return encoder
}

func getStructEncode(t reflect.Type) encodeFunc {
	return getStructEncoder(t).Write
}

func getStructPtrEncode(t reflect.Type) encodeFunc {
	return getStructEncoder(t).Encode
}

func getOtherEncode(t reflect.Type) encodeFunc {
	if encoder := getOtherEncoder(t); encoder != nil {
		return encoder.Write
	}
	return nil
}

func getOtherPtrEncode(t reflect.Type) encodeFunc {
	if encoder := getOtherEncoder(t); encoder != nil {
		return encoder.Encode
	}
	return nil
}

func getPtrEncode(t reflect.Type) encodeFunc {
	if f := getOtherPtrEncode(t); f != nil {
		return f
	}
	switch t.Kind() {
	case reflect.Int:
		return intPtrEncode
	case reflect.Int8:
		return int8PtrEncode
	case reflect.Int16:
		return int16PtrEncode
	case reflect.Int32:
		return int32PtrEncode
	case reflect.Int64:
		return int64PtrEncode
	case reflect.Uint:
		return uintPtrEncode
	case reflect.Uint8:
		return uint8PtrEncode
	case reflect.Uint16:
		return uint16PtrEncode
	case reflect.Uint32:
		return uint32PtrEncode
	case reflect.Uint64, reflect.Uintptr:
		return uint64PtrEncode
	case reflect.Bool:
		return boolPtrEncode
	case reflect.Float32:
		return float32PtrEncode
	case reflect.Float64:
		return float64PtrEncode
	case reflect.Complex64:
		return complex64PtrEncode
	case reflect.Complex128:
		return complex128PtrEncode
	case reflect.Array:
		return arrayPtrEncode
	case reflect.Interface:
		return interfacePtrEncode
	case reflect.Map:
		return mapPtrEncode
	case reflect.Ptr:
		for t.Kind() == reflect.Ptr {
			t = t.Elem()
		}
		switch t.Kind() {
		case reflect.Func, reflect.Chan, reflect.UnsafePointer:
			return nil
		}
		return ptrEncode
	case reflect.Slice:
		return slicePtrEncode
	case reflect.String:
		return stringPtrEncode
	case reflect.Struct:
		return getStructPtrEncode(t)
	}
	return nil
}
