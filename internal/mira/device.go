package mira

import (
	"fmt"
	"time"

	"github.com/sstallion/go-hid"
)

type Device struct {
	dev *hid.Device
}

// NewDevice creates a new Mira device instance
func NewDevice() (*Device, error) {
	// Initialize HID
	if err := hid.Init(); err != nil {
		return nil, fmt.Errorf("failed to initialize HID: %w", err)
	}

	dev, err := hid.OpenFirst(VID, PID)
	if err != nil {
		hid.Exit()
		return nil, fmt.Errorf("failed to open Mira device: %w", err)
	}

	return &Device{dev: dev}, nil
}

// Close closes the device and cleans up
func (d *Device) Close() error {
	if err := d.dev.Close(); err != nil {
		return fmt.Errorf("failed to close device: %w", err)
	}
	hid.Exit()
	return nil
}

func (d *Device) write(data []byte) error {
	_, err := d.dev.Write(data)
	if err != nil {
		return fmt.Errorf("failed to write to device: %w", err)
	}
	time.Sleep(SleepAfterWriteMS * time.Millisecond)
	return nil
}

func (d *Device) Refresh() error {
	return d.write([]byte{USBReportID, byte(OpCodeRefresh)})
}

func (d *Device) SetAutoDitherMode(mode AutoDitherMode) error {
	data := []byte{USBReportID, byte(OpCodeSetAutoDitherMode)}
	data = append(data, mode[:]...)
	return d.write(data)
}

func (d *Device) SetSpeed(speed int) error {
	if speed < 1 || speed > 7 {
		return fmt.Errorf("speed must be between 1 and 7")
	}
	adjustedSpeed := 11 - speed
	return d.write([]byte{USBReportID, byte(OpCodeSetSpeed), byte(adjustedSpeed)})
}

func (d *Device) SetContrast(contrast int) error {
	if contrast < 0 || contrast > 15 {
		return fmt.Errorf("contrast must be between 0 and 15")
	}
	return d.write([]byte{USBReportID, byte(OpCodeSetContrast), byte(contrast)})
}

func (d *Device) SetRefreshMode(mode RefreshMode) error {
	return d.write([]byte{USBReportID, byte(OpCodeSetRefreshMode), byte(mode)})
}

func (d *Device) SetDitherMode(mode int) error {
	if mode < 0 || mode > 3 {
		return fmt.Errorf("dither mode must be between 0 and 3")
	}
	return d.write([]byte{USBReportID, byte(OpCodeSetDitherMode), byte(mode)})
}

func (d *Device) SetColorFilter(whiteFilter, blackFilter int) error {
	if whiteFilter < 0 || whiteFilter > 254 || blackFilter < 0 || blackFilter > 254 {
		return fmt.Errorf("filter values must be between 0 and 254")
	}
	adjustedWhite := 255 - whiteFilter
	return d.write([]byte{USBReportID, byte(OpCodeSetColorFilter), byte(adjustedWhite), byte(blackFilter)})
}

func (d *Device) SetColdLight(brightness int) error {
	if brightness < 0 || brightness > 254 {
		return fmt.Errorf("brightness must be between 0 and 254")
	}
	return d.write([]byte{USBReportID, byte(OpCodeSetColdLight), byte(brightness)})
}

func (d *Device) SetWarmLight(brightness int) error {
	if brightness < 0 || brightness > 254 {
		return fmt.Errorf("brightness must be between 0 and 254")
	}
	return d.write([]byte{USBReportID, byte(OpCodeSetWarmLight), byte(brightness)})
}

// DeviceInfo contains information about an HID device
type DeviceInfo struct {
	VID          uint16
	PID          uint16
	Manufacturer string
	Product      string
}

// ListDevices returns a list of all HID devices
func ListDevices() ([]DeviceInfo, error) {
	if err := hid.Init(); err != nil {
		return nil, fmt.Errorf("failed to initialize HID: %w", err)
	}
	defer hid.Exit()

	var devices []DeviceInfo
	hid.Enumerate(0, 0, func(info *hid.DeviceInfo) error {
		devices = append(devices, DeviceInfo{
			VID:          info.VendorID,
			PID:          info.ProductID,
			Manufacturer: info.MfrStr,
			Product:      info.ProductStr,
		})
		return nil
	})

	return devices, nil
}
