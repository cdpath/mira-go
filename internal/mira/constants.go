package mira

const (
	VID = 0x0416
	PID = 0x5020
)

const (
	USBReportID       = 0x00
	SleepAfterWriteMS = 500
)

// OpCode represents different operation codes for Mira commands
type OpCode byte

const (
	OpCodeRefresh           OpCode = 0x01
	OpCodeSetRefreshMode    OpCode = 0x02
	OpCodeSetSpeed          OpCode = 0x04
	OpCodeSetContrast       OpCode = 0x05
	OpCodeSetColdLight      OpCode = 0x06
	OpCodeSetWarmLight      OpCode = 0x07
	OpCodeSetDitherMode     OpCode = 0x09
	OpCodeSetColorFilter    OpCode = 0x11
	OpCodeSetAutoDitherMode OpCode = 0x12
)

// RefreshMode represents different screen refresh modes
type RefreshMode byte

const (
	RefreshModeDirectUpdate RefreshMode = 0x01 // black/white, fast
	RefreshModeGrayUpdate   RefreshMode = 0x02 // gray scale, slow
	RefreshModeA2           RefreshMode = 0x03 // fast
)

// AutoDitherMode represents different auto dither modes
type AutoDitherMode [4]byte

var (
	AutoDitherModeDisable = AutoDitherMode{0, 0, 0, 0}
	AutoDitherModeLow     = AutoDitherMode{1, 0, 30, 10}
	AutoDitherModeMiddle  = AutoDitherMode{1, 0, 40, 10}
	AutoDitherModeHigh    = AutoDitherMode{1, 0, 50, 30}
)
