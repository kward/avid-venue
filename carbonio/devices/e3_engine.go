package devices

const (
	e3Engine_numMicInputs   = 4
	e3Engine_numLineOutputs = 4
	e3Engine_numAESOutputs  = 4 // Each physical connector supports two signals.
)

// TODO(2020-02-25) Add support for the E3 Engine. Determine the device type by
// looking at the board_id files.
type E3Engine struct{}
