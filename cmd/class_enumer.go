// Code generated by "enumer -type Class -trimprefix Class -transform=lower -text"; DO NOT EDIT.

package cmd

import (
	"fmt"
	"strings"
)

const _ClassName = "meterchargervehicletariffcircuitsitemqttdatabaseeebusjavascriptgohemsinfluxmessenger"

var _ClassIndex = [...]uint8{0, 5, 12, 19, 25, 32, 36, 40, 48, 53, 63, 65, 69, 75, 84}

const _ClassLowerName = "meterchargervehicletariffcircuitsitemqttdatabaseeebusjavascriptgohemsinfluxmessenger"

func (i Class) String() string {
	i -= 1
	if i < 0 || i >= Class(len(_ClassIndex)-1) {
		return fmt.Sprintf("Class(%d)", i+1)
	}
	return _ClassName[_ClassIndex[i]:_ClassIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _ClassNoOp() {
	var x [1]struct{}
	_ = x[ClassMeter-(1)]
	_ = x[ClassCharger-(2)]
	_ = x[ClassVehicle-(3)]
	_ = x[ClassTariff-(4)]
	_ = x[ClassCircuit-(5)]
	_ = x[ClassSite-(6)]
	_ = x[ClassMqtt-(7)]
	_ = x[ClassDatabase-(8)]
	_ = x[ClassEEBus-(9)]
	_ = x[ClassJavascript-(10)]
	_ = x[ClassGo-(11)]
	_ = x[ClassHEMS-(12)]
	_ = x[ClassInflux-(13)]
	_ = x[ClassMessenger-(14)]
}

var _ClassValues = []Class{ClassMeter, ClassCharger, ClassVehicle, ClassTariff, ClassCircuit, ClassSite, ClassMqtt, ClassDatabase, ClassEEBus, ClassJavascript, ClassGo, ClassHEMS, ClassInflux, ClassMessenger}

var _ClassNameToValueMap = map[string]Class{
	_ClassName[0:5]:        ClassMeter,
	_ClassLowerName[0:5]:   ClassMeter,
	_ClassName[5:12]:       ClassCharger,
	_ClassLowerName[5:12]:  ClassCharger,
	_ClassName[12:19]:      ClassVehicle,
	_ClassLowerName[12:19]: ClassVehicle,
	_ClassName[19:25]:      ClassTariff,
	_ClassLowerName[19:25]: ClassTariff,
	_ClassName[25:32]:      ClassCircuit,
	_ClassLowerName[25:32]: ClassCircuit,
	_ClassName[32:36]:      ClassSite,
	_ClassLowerName[32:36]: ClassSite,
	_ClassName[36:40]:      ClassMqtt,
	_ClassLowerName[36:40]: ClassMqtt,
	_ClassName[40:48]:      ClassDatabase,
	_ClassLowerName[40:48]: ClassDatabase,
	_ClassName[48:53]:      ClassEEBus,
	_ClassLowerName[48:53]: ClassEEBus,
	_ClassName[53:63]:      ClassJavascript,
	_ClassLowerName[53:63]: ClassJavascript,
	_ClassName[63:65]:      ClassGo,
	_ClassLowerName[63:65]: ClassGo,
	_ClassName[65:69]:      ClassHEMS,
	_ClassLowerName[65:69]: ClassHEMS,
	_ClassName[69:75]:      ClassInflux,
	_ClassLowerName[69:75]: ClassInflux,
	_ClassName[75:84]:      ClassMessenger,
	_ClassLowerName[75:84]: ClassMessenger,
}

var _ClassNames = []string{
	_ClassName[0:5],
	_ClassName[5:12],
	_ClassName[12:19],
	_ClassName[19:25],
	_ClassName[25:32],
	_ClassName[32:36],
	_ClassName[36:40],
	_ClassName[40:48],
	_ClassName[48:53],
	_ClassName[53:63],
	_ClassName[63:65],
	_ClassName[65:69],
	_ClassName[69:75],
	_ClassName[75:84],
}

// ClassString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func ClassString(s string) (Class, error) {
	if val, ok := _ClassNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _ClassNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Class values", s)
}

// ClassValues returns all values of the enum
func ClassValues() []Class {
	return _ClassValues
}

// ClassStrings returns a slice of all String values of the enum
func ClassStrings() []string {
	strs := make([]string, len(_ClassNames))
	copy(strs, _ClassNames)
	return strs
}

// IsAClass returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Class) IsAClass() bool {
	for _, v := range _ClassValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalText implements the encoding.TextMarshaler interface for Class
func (i Class) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for Class
func (i *Class) UnmarshalText(text []byte) error {
	var err error
	*i, err = ClassString(string(text))
	return err
}
