package modbus

import (
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/evcc-io/evcc/util"
	"github.com/volkszaehler/mbmd/meters"
	"github.com/volkszaehler/mbmd/meters/rs485"
	"github.com/volkszaehler/mbmd/meters/sunspec"
)

type Protocol int

const (
	Tcp Protocol = iota
	Rtu
	Ascii

	CoilOn uint16 = 0xFF00
)

// Settings contains the ModBus TCP settings
// RTU field is included for compatibility with modbus.tpl which renders rtu: false for TCP
// TODO remove RTU field (https://github.com/evcc-io/evcc/issues/3360)
type TcpSettings struct {
	URI string
	ID  uint8
	RTU *bool `mapstructure:"rtu"`
}

// Settings contains the ModBus settings
type Settings struct {
	ID                  uint8
	SubDevice           int
	URI, Device, Comset string
	Baudrate            int
	UDP                 bool
	RTU                 *bool // indicates RTU over TCP if true
}

func (s *Settings) String() string {
	if s.URI != "" {
		return s.URI
	}
	return s.Device
}

var (
	connections = make(map[string]meters.Connection)
	mu          sync.Mutex
)

func registeredConnection(key string, newConn meters.Connection) meters.Connection {
	mu.Lock()
	defer mu.Unlock()

	if conn, ok := connections[key]; ok {
		return conn
	}

	connections[key] = newConn

	return newConn
}

// ProtocolFromRTU identifies the wire format from the RTU setting
func ProtocolFromRTU(rtu *bool) Protocol {
	if rtu != nil && *rtu {
		return Rtu
	}
	return Tcp
}

// NewConnection creates physical modbus device from config
func NewConnection(uri, device, comset string, baudrate int, proto Protocol, slaveID uint8) (*Connection, error) {
	conn, err := physicalConnection(proto, Settings{
		URI:      uri,
		Device:   device,
		Comset:   comset,
		Baudrate: baudrate,
	})
	if err != nil {
		return nil, err
	}

	res := &Connection{
		Connection: conn.Clone(slaveID),
		delay:      func() {}, // no delay
		logger:     new(logger),
	}

	return res, nil
}

func physicalConnection(proto Protocol, cfg Settings) (meters.Connection, error) {
	var conn meters.Connection

	if (cfg.Device != "") == (cfg.URI != "") {
		return nil, errors.New("invalid modbus configuration: must have either uri or device")
	}

	if cfg.Device != "" {
		switch strings.ToUpper(cfg.Comset) {
		case "8N1", "8E1", "8N2":
		case "80":
			cfg.Comset = "8E1"
		default:
			return nil, fmt.Errorf("invalid comset: %s", cfg.Comset)
		}

		if cfg.Baudrate == 0 {
			return nil, errors.New("invalid modbus configuration: need baudrate and comset")
		}

		if proto == Ascii {
			conn = registeredConnection(cfg.Device, meters.NewASCII(cfg.Device, cfg.Baudrate, cfg.Comset))
		} else {
			conn = registeredConnection(cfg.Device, meters.NewRTU(cfg.Device, cfg.Baudrate, cfg.Comset))
		}
	}

	if cfg.URI != "" {
		cfg.URI = util.DefaultPort(cfg.URI, 502)

		switch proto {
		case Rtu:
			if cfg.UDP {
				conn = registeredConnection(cfg.URI, meters.NewRTUOverUDP(cfg.URI))
			} else {
				conn = registeredConnection(cfg.URI, meters.NewRTUOverTCP(cfg.URI))
			}
		case Ascii:
			conn = registeredConnection(cfg.URI, meters.NewASCIIOverTCP(cfg.URI))
		default:
			conn = registeredConnection(cfg.URI, meters.NewTCP(cfg.URI))
		}
	}

	return conn, nil
}

// NewDevice creates physical modbus device from config
func NewDevice(model string, subdevice int) (device meters.Device, err error) {
	if IsRS485(model) {
		device, err = rs485.NewDevice(strings.ToUpper(model))
	} else {
		device = sunspec.NewDevice(strings.ToUpper(model), subdevice)
	}

	if device == nil {
		err = errors.New("invalid modbus configuration: need either uri or device")
	}

	return device, err
}

// IsRS485 determines if model is a known MBMD rs485 device model
func IsRS485(model string) bool {
	for k := range rs485.Producers {
		if strings.EqualFold(model, k) {
			return true
		}
	}
	return false
}

// RS485FindDeviceOp checks is RS485 device supports operation
func RS485FindDeviceOp(device *rs485.RS485, measurement meters.Measurement) (op rs485.Operation, err error) {
	ops := device.Producer().Produce()

	for _, op := range ops {
		if op.IEC61850 == measurement {
			return op, nil
		}
	}

	return op, fmt.Errorf("unsupported measurement: %s", measurement.String())
}
