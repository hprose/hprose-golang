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
 * io/map_encoder.go                                      *
 *                                                        *
 * hprose map encoder for Go.                             *
 *                                                        *
 * LastModified: Oct 6, 2016                              *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package io

var mapBodyEncoders = map[uintptr]func(*Writer, interface{}){
	getType((map[string]string)(nil)):           stringStringMapEncoder,
	getType((map[string]interface{})(nil)):      stringInterfaceMapEncoder,
	getType((map[string]int)(nil)):              stringIntMapEncoder,
	getType((map[int]int)(nil)):                 intIntMapEncoder,
	getType((map[int]string)(nil)):              intStringMapEncoder,
	getType((map[int]interface{})(nil)):         intInterfaceMapEncoder,
	getType((map[interface{}]interface{})(nil)): interfaceInterfaceMapEncoder,
	getType((map[interface{}]int)(nil)):         interfaceIntMapEncoder,
	getType((map[interface{}]string)(nil)):      interfaceStringMapEncoder,
}

// RegisterMapEncoder for fast serialize custom map type.
// This function is usually used for code generators.
// This function should be called in package init function.
func RegisterMapEncoder(m interface{}, encoder func(*Writer, interface{})) {
	mapBodyEncoders[getType(m)] = encoder
}

func stringStringMapEncoder(w *Writer, v interface{}) {
	m := v.(map[string]string)
	for k, v := range m {
		w.WriteString(k)
		w.WriteString(v)
	}
}

func stringInterfaceMapEncoder(w *Writer, v interface{}) {
	m := v.(map[string]interface{})
	for k, v := range m {
		w.WriteString(k)
		w.Serialize(v)
	}
}

func stringIntMapEncoder(w *Writer, v interface{}) {
	m := v.(map[string]int)
	for k, v := range m {
		w.WriteString(k)
		w.WriteInt(int64(v))
	}
}

func intIntMapEncoder(w *Writer, v interface{}) {
	m := v.(map[int]int)
	for k, v := range m {
		w.WriteInt(int64(k))
		w.WriteInt(int64(v))
	}
}

func intStringMapEncoder(w *Writer, v interface{}) {
	m := v.(map[int]string)
	for k, v := range m {
		w.WriteInt(int64(k))
		w.WriteString(v)
	}
}

func intInterfaceMapEncoder(w *Writer, v interface{}) {
	m := v.(map[int]interface{})
	for k, v := range m {
		w.WriteInt(int64(k))
		w.Serialize(v)
	}
}

func interfaceInterfaceMapEncoder(w *Writer, v interface{}) {
	m := v.(map[interface{}]interface{})
	for k, v := range m {
		w.Serialize(k)
		w.Serialize(v)
	}
}

func interfaceIntMapEncoder(w *Writer, v interface{}) {
	m := v.(map[interface{}]int)
	for k, v := range m {
		w.Serialize(k)
		w.WriteInt(int64(v))
	}
}

func interfaceStringMapEncoder(w *Writer, v interface{}) {
	m := v.(map[interface{}]string)
	for k, v := range m {
		w.Serialize(k)
		w.WriteString(v)
	}
}
