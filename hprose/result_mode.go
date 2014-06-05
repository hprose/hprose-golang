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
 * hprose/result_mode.go                                  *
 *                                                        *
 * hprose ResultMode enum for Go.                         *
 *                                                        *
 * LastModified: Jan 21, 2014                             *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package hprose

type ResultMode int

const (
	Normal = ResultMode(iota)
	Serialized
	Raw
	RawWithEndTag
)

func (result_mode ResultMode) String() string {
	switch result_mode {
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
