// Code generated by "stringer -output=spi_string.go -type=Enum spi_types.go"; DO NOT EDIT.

package spi

import "strconv"

const _Enum_name = "unknownSPIPowerStatusMuteGainPadPhantomBlinky"

var _Enum_index = [...]uint8{0, 10, 15, 21, 25, 29, 32, 39, 45}

func (i Enum) String() string {
	if i < 0 || i >= Enum(len(_Enum_index)-1) {
		return "Enum(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Enum_name[_Enum_index[i]:_Enum_index[i+1]]
}
