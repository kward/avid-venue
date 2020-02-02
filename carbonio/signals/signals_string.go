// Code generated by "stringer -output=signals_string.go -type=Connector,Format,Level"; DO NOT EDIT.

package signals

import "strconv"

const _Connector_name = "XLRConnectorJackConnector"

var _Connector_index = [...]uint8{0, 12, 25}

func (i Connector) String() string {
	if i < 0 || i >= Connector(len(_Connector_index)-1) {
		return "Connector(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Connector_name[_Connector_index[i]:_Connector_index[i+1]]
}

const _Format_name = "AnalogFormatAESFormat"

var _Format_index = [...]uint8{0, 12, 21}

func (i Format) String() string {
	if i < 0 || i >= Format(len(_Format_index)-1) {
		return "Format(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Format_name[_Format_index[i]:_Format_index[i+1]]
}

const _Level_name = "LineLevelMicLevel"

var _Level_index = [...]uint8{0, 9, 17}

func (i Level) String() string {
	if i < 0 || i >= Level(len(_Level_index)-1) {
		return "Level(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Level_name[_Level_index[i]:_Level_index[i+1]]
}