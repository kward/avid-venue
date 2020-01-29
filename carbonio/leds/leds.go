// Package leds enables control of LEDs through the Carbon I/O front-panel LEDs.
//
// The Avid Stage 16 LEDs are controlled through the kernel interface of the
// spi device. The current state of the LED can be read by looking at the
// contents of the file that represents the device interface. Changing the
// state can be done by writing to the same interface.
package leds

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/kward/golib/math"
)

const (
	powerIface  = "/sys/bus/spi/devices/spi4.0/status_led_1_en"
	statusIface = "/sys/bus/spi/devices/spi4.0/status_led_0_en"
	muteIface   = "/sys/bus/spi/devices/spi4.0/mute_led_en"
	maxDataLen  = 40
)

type LED interface {
	State() (LEDState, error)
	SetState(state LEDState) error
	String() string
}

//-----------------------------------------------------------------------------
// The Power LED supports the states On, Off, and Alert.
var Power LED

type powerLED struct{}

var _ LED = new(powerLED)

func (l *powerLED) State() (LEDState, error) {
	s, err := readState(powerIface)
	if err != nil {
		return Unknown, err
	}
	switch s {
	case '0':
		return Off, nil
	case '1':
		return Alert, nil
	case '2':
		return On, nil
	case 't':
		return testLEDState, nil
	default:
		return Unknown, fmt.Errorf("unrecognized LEDState %q [%d]", s, s)
	}
}

func (l *powerLED) SetState(state LEDState) error {
	var s byte
	switch state {
	case Off:
		s = '0'
	case Alert:
		s = '1'
	case On:
		s = '2'
	case testLEDState:
		s = 't'
	default:
		return fmt.Errorf("unrecognized LEDState %q [%d]", state, state)
	}

	return writeState(powerIface, s)
}

func (l *powerLED) String() string {
	state, _ := l.State()
	return fmt.Sprintf("Power LED %s", state)
}

//-----------------------------------------------------------------------------
// The Status LED supports the states On, Off, and Alert.
var Status LED

type statusLED struct{}

var _ LED = new(statusLED)

func (l *statusLED) State() (LEDState, error) {
	s, err := readState(statusIface)
	if err != nil {
		return Unknown, err
	}
	switch s {
	case '0':
		return Off, nil
	case '1':
		return Alert, nil
	case '2':
		return On, nil
	case 255:
		return testLEDState, nil
	default:
		return Unknown, fmt.Errorf("unrecognized LEDState %q [%d]", s, s)
	}
}
func (l *statusLED) SetState(state LEDState) error {
	var s byte
	switch state {
	case Off:
		s = '0'
	case Alert:
		s = '1'
	case On:
		s = '2'
	case testLEDState:
		s = 't'
	default:
		return fmt.Errorf("unrecognized LEDState %q [%d]", state, state)
	}
	return writeState(statusIface, s)
}
func (l *statusLED) String() string {
	state, _ := l.State()
	return fmt.Sprintf("State LED %s", state)
}

//-----------------------------------------------------------------------------
// The Mute LED supports the states On and Off.
var Mute LED

type muteLED struct{}

var _ LED = new(muteLED)

func (l *muteLED) State() (LEDState, error) {
	s, err := readState(muteIface)
	if err != nil {
		return Unknown, err
	}
	switch s {
	case '0':
		return Off, nil
	case '1':
		return On, nil
	case 255:
		return testLEDState, nil
	default:
		return Unknown, fmt.Errorf("unrecognized LEDState %q [%d]", s, s)
	}
}
func (l *muteLED) SetState(state LEDState) error {
	var s byte
	switch state {
	case Off:
		s = '0'
	case On:
		s = '1'
	case testLEDState:
		s = 't'
	default:
		return fmt.Errorf("unrecognized LEDState %q [%d]", state, state)
	}
	return writeState(muteIface, s)
}
func (l *muteLED) String() string {
	state, _ := l.State()
	return fmt.Sprintf("Mute LED %s", state)
}

//=============================================================================

var readFileFn = ioutil.ReadFile

func readState(filename string) (byte, error) {
	data, err := readFileFn(filename)
	if err != nil {
		return 0, err
	}
	if len(data) != 2 || data[1] != '\n' {
		return 0, fmt.Errorf("%q contains unexpected data; %v", filename, data[0:math.MinInt(len(data), maxDataLen)])
	}
	return data[0], nil
}

var writeFileFn = ioutil.WriteFile

func writeState(filename string, s byte) error {
	return writeFileFn(filename, []byte{s, '\n'}, os.FileMode(0644))
}

func init() {
	Power = new(powerLED)
	Status = new(statusLED)
	Mute = new(muteLED)
}