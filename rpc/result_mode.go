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
 * rpc/result_mode.go                                     *
 *                                                        *
 * hprose ResultMode enum for Go.                         *
 *                                                        *
 * LastModified: Sep 10, 2016                             *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package rpc

// ResultMode is result mode
type ResultMode int

const (
	// Normal is default mode
	Normal = ResultMode(iota)
	// Serialized means the result is serialized
	Serialized
	// Raw means the result is the raw bytes data
	Raw
	// RawWithEndTag means the result is the raw bytes data with the end tag
	RawWithEndTag
)

func (mode ResultMode) String() string {
	switch mode {
	case Normal:
		return "Normal"
	case Serialized:
		return "Serialized"
	case Raw:
		return "Raw"
	case RawWithEndTag:
		return "RawWithEndTag"
	}
	panic("unknown value of ResultMode")
}
