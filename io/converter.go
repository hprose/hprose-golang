/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| io/converter.go                                          |
|                                                          |
| LastModified: May 17, 2021                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package io

import (
	"fmt"
	"math/big"
	"reflect"
	"sync"
	"time"
	"unsafe"

	"github.com/google/uuid"
	"github.com/modern-go/reflect2"
)

func Convert(src interface{}, t reflect.Type) (interface{}, error) {
	if t == nil || t == interfaceType {
		return src, nil
	}
	t2 := reflect2.Type2(t)
	p := t2.New()
	if converter := GetConverter(reflect.TypeOf(src), t); converter != nil {
		dec := NewDecoder(nil)
		if converter(dec, src, p); dec.Error == nil {
			return t2.Indirect(p), nil
		}
	}
	data, err := Marshal(src)
	if err != nil {
		return nil, err
	}
	if err := Unmarshal(data, p); err != nil {
		return nil, err
	}
	return t2.Indirect(p), nil
}

type fastConverterMapKey struct {
	Source      reflect.Kind
	Destination reflect.Kind
}

func strConverter(dec *Decoder, o interface{}, p interface{}) {
	switch o := o.(type) {
	case string:
		*(*string)(reflect2.PtrOf(p)) = o
	case *string:
		*(*string)(reflect2.PtrOf(p)) = *o
	case fmt.Stringer:
		*(*string)(reflect2.PtrOf(p)) = o.String()
	case fmt.GoStringer:
		*(*string)(reflect2.PtrOf(p)) = o.GoString()
	default:
		*(*string)(reflect2.PtrOf(p)) = fmt.Sprint(o)
	}
}

func assignTo(dec *Decoder, o interface{}, p interface{}) {
	reflect.ValueOf(p).Elem().Set(reflect.ValueOf(o))
}

func ptrCopy(dec *Decoder, o interface{}, p interface{}) {
	*(*unsafe.Pointer)(reflect2.PtrOf(p)) = reflect2.PtrOf(o)
}

func sliceCopy(dec *Decoder, o interface{}, p interface{}) {
	*(*sliceHeader)(reflect2.PtrOf(p)) = *(*sliceHeader)(reflect2.PtrOf(o))
}

func mapCopy(dec *Decoder, o interface{}, p interface{}) {
	reflect2.TypeOf(p).UnsafeSet(reflect2.PtrOf(p), reflect2.PtrOf(o))
}

func dataCopy(dec *Decoder, o interface{}, p interface{}) {
	value := reflect.ValueOf(o)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	reflect.ValueOf(p).Elem().Set(value)
}

func arrayCopy(dec *Decoder, o interface{}, p interface{}) {
	value := reflect.ValueOf(o)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	reflect.Copy(reflect.ValueOf(p).Elem(), value)
}

func ptrConverter(dec *Decoder, o interface{}, p interface{}) {
	t := reflect.TypeOf(p).Elem().Elem()
	t2 := reflect2.Type2(t)
	ptr := (*unsafe.Pointer)(reflect2.PtrOf(p))
	if *ptr == nil {
		*ptr = t2.UnsafeNew()
	}
	if converter := GetConverter(reflect.TypeOf(o), t); converter != nil {
		converter(dec, o, t2.PackEFace(*ptr))
	}
}

var fastConverterMap map[fastConverterMapKey]func(dec *Decoder, o interface{}, p interface{})

type converterMapKey struct {
	Source      reflect.Type
	Destination reflect.Type
}

var converterMap sync.Map

// RegisterConverter for converting src to dest.
func RegisterConverter(src, dest reflect.Type, converter func(dec *Decoder, o interface{}, p interface{})) {
	converterMap.Store(converterMapKey{src, dest}, converter)
}

// GetConverter returns the converter for converting src to dest.
func GetConverter(src, dest reflect.Type) func(dec *Decoder, o interface{}, p interface{}) {
	if converter, ok := converterMap.Load(converterMapKey{src, dest}); ok {
		return converter.(func(dec *Decoder, o interface{}, p interface{}))
	}
	if dest == interfaceType {
		return assignTo
	}
	switch dest.Kind() {
	case reflect.String:
		return strConverter
	case reflect.Ptr:
		if src == dest && src.Elem().Kind() == reflect.Struct ||
			src == dest.Elem() && src.Kind() != reflect.Ptr {
			return ptrCopy
		}
		return ptrConverter
	default:
		if src == dest || (src.Kind() == reflect.Ptr && src.Elem() == dest) {
			switch dest.Kind() {
			case reflect.Array:
				return arrayCopy
			case reflect.Slice:
				return sliceCopy
			case reflect.Map:
				return mapCopy
			default:
				return dataCopy
			}
		}
		return fastConverterMap[fastConverterMapKey{src.Kind(), dest.Kind()}]
	}
}

func init() {
	fastConverterMap = map[fastConverterMapKey]func(dec *Decoder, o interface{}, p interface{}){
		{reflect.String, reflect.Bool}: func(dec *Decoder, o interface{}, p interface{}) {
			*(*bool)(reflect2.PtrOf(p)) = dec.stringToBool(*(*string)(reflect2.PtrOf(o)))
		},
		{reflect.String, reflect.Int}: func(dec *Decoder, o interface{}, p interface{}) {
			*(*int)(reflect2.PtrOf(p)) = int(dec.stringToInt64(*(*string)(reflect2.PtrOf(o)), 0))
		},
		{reflect.String, reflect.Int8}: func(dec *Decoder, o interface{}, p interface{}) {
			*(*int8)(reflect2.PtrOf(p)) = int8(dec.stringToInt64(*(*string)(reflect2.PtrOf(o)), 8))
		},
		{reflect.String, reflect.Int16}: func(dec *Decoder, o interface{}, p interface{}) {
			*(*int16)(reflect2.PtrOf(p)) = int16(dec.stringToInt64(*(*string)(reflect2.PtrOf(o)), 16))
		},
		{reflect.String, reflect.Int32}: func(dec *Decoder, o interface{}, p interface{}) {
			*(*int32)(reflect2.PtrOf(p)) = int32(dec.stringToInt64(*(*string)(reflect2.PtrOf(o)), 32))
		},
		{reflect.String, reflect.Int64}: func(dec *Decoder, o interface{}, p interface{}) {
			*(*int64)(reflect2.PtrOf(p)) = dec.stringToInt64(*(*string)(reflect2.PtrOf(o)), 64)
		},
		{reflect.String, reflect.Uint}: func(dec *Decoder, o interface{}, p interface{}) {
			*(*uint)(reflect2.PtrOf(p)) = uint(dec.stringToUint64(*(*string)(reflect2.PtrOf(o)), 0))
		},
		{reflect.String, reflect.Uint8}: func(dec *Decoder, o interface{}, p interface{}) {
			*(*uint8)(reflect2.PtrOf(p)) = uint8(dec.stringToUint64(*(*string)(reflect2.PtrOf(o)), 8))
		},
		{reflect.String, reflect.Uint16}: func(dec *Decoder, o interface{}, p interface{}) {
			*(*uint16)(reflect2.PtrOf(p)) = uint16(dec.stringToUint64(*(*string)(reflect2.PtrOf(o)), 16))
		},
		{reflect.String, reflect.Uint32}: func(dec *Decoder, o interface{}, p interface{}) {
			*(*uint32)(reflect2.PtrOf(p)) = uint32(dec.stringToUint64(*(*string)(reflect2.PtrOf(o)), 32))
		},
		{reflect.String, reflect.Uint64}: func(dec *Decoder, o interface{}, p interface{}) {
			*(*uint64)(reflect2.PtrOf(p)) = dec.stringToUint64(*(*string)(reflect2.PtrOf(o)), 64)
		},
		{reflect.String, reflect.Uintptr}: func(dec *Decoder, o interface{}, p interface{}) {
			*(*uintptr)(reflect2.PtrOf(p)) = uintptr(dec.stringToUint64(*(*string)(reflect2.PtrOf(o)), 64))
		},
		{reflect.String, reflect.Float32}: func(dec *Decoder, o interface{}, p interface{}) {
			*(*float32)(reflect2.PtrOf(p)) = dec.stringToFloat32(*(*string)(reflect2.PtrOf(o)))
		},
		{reflect.String, reflect.Float64}: func(dec *Decoder, o interface{}, p interface{}) {
			*(*float64)(reflect2.PtrOf(p)) = dec.stringToFloat64(*(*string)(reflect2.PtrOf(o)))
		},
		{reflect.String, reflect.Complex64}: func(dec *Decoder, o interface{}, p interface{}) {
			*(*complex64)(reflect2.PtrOf(p)) = dec.stringToComplex64(*(*string)(reflect2.PtrOf(o)))
		},
		{reflect.String, reflect.Complex128}: func(dec *Decoder, o interface{}, p interface{}) {
			*(*complex128)(reflect2.PtrOf(p)) = dec.stringToComplex128(*(*string)(reflect2.PtrOf(o)))
		},
	}
	RegisterConverter(stringType, bigIntValueType, func(dec *Decoder, o interface{}, p interface{}) {
		s := *(*string)(reflect2.PtrOf(o))
		value := dec.stringToBigInt(s, bigIntValueType)
		if value != nil {
			*(*big.Int)(reflect2.PtrOf(p)) = *value
		} else {
			*(*big.Int)(reflect2.PtrOf(p)) = *bigIntZero
		}
	})
	RegisterConverter(stringType, bigIntType, func(dec *Decoder, o interface{}, p interface{}) {
		*(**big.Int)(reflect2.PtrOf(p)) = dec.stringToBigInt(*(*string)(reflect2.PtrOf(o)), bigIntType)
	})
	RegisterConverter(stringType, bigFloatValueType, func(dec *Decoder, o interface{}, p interface{}) {
		s := *(*string)(reflect2.PtrOf(o))
		value := dec.stringToBigFloat(s, bigFloatValueType)
		if value != nil {
			*(*big.Float)(reflect2.PtrOf(p)) = *value
		} else {
			*(*big.Float)(reflect2.PtrOf(p)) = *bigFloatZero
		}
	})
	RegisterConverter(stringType, bigFloatType, func(dec *Decoder, o interface{}, p interface{}) {
		*(**big.Float)(reflect2.PtrOf(p)) = dec.stringToBigFloat(*(*string)(reflect2.PtrOf(o)), bigFloatType)
	})
	RegisterConverter(stringType, bigRatValueType, func(dec *Decoder, o interface{}, p interface{}) {
		s := *(*string)(reflect2.PtrOf(o))
		value := dec.stringToBigRat(s, bigRatValueType)
		if value != nil {
			*(*big.Rat)(reflect2.PtrOf(p)) = *value
		} else {
			*(*big.Rat)(reflect2.PtrOf(p)) = *bigRatZero
		}
	})
	RegisterConverter(stringType, bigRatType, func(dec *Decoder, o interface{}, p interface{}) {
		*(**big.Rat)(reflect2.PtrOf(p)) = dec.stringToBigRat(*(*string)(reflect2.PtrOf(o)), bigRatType)
	})
	RegisterConverter(stringType, bytesType, func(dec *Decoder, o interface{}, p interface{}) {
		*(*[]byte)(reflect2.PtrOf(p)) = reflect2.UnsafeCastString(*(*string)(reflect2.PtrOf(o)))
	})
	RegisterConverter(bytesType, stringType, func(dec *Decoder, o interface{}, p interface{}) {
		*(*string)(reflect2.PtrOf(p)) = unsafeString(*(*[]byte)(reflect2.PtrOf(o)))
	})
	RegisterConverter(stringType, timeType, func(dec *Decoder, o interface{}, p interface{}) {
		*(*time.Time)(reflect2.PtrOf(p)) = dec.stringToTime(*(*string)(reflect2.PtrOf(o)))
	})
	RegisterConverter(stringType, uuidType, func(dec *Decoder, o interface{}, p interface{}) {
		*(*uuid.UUID)(reflect2.PtrOf(p)) = dec.stringToUUID(*(*string)(reflect2.PtrOf(o)))
	})
}
