package charger

// Code generated by github.com/evcc-io/evcc/cmd/tools/decorate.go. DO NOT EDIT.

import (
	"github.com/evcc-io/evcc/api"
)

func decorateOCPP(base *OCPP, meter func() (float64, error), meterEnergy func() (float64, error), meterCurrent func() (float64, float64, float64, error), phaseSwitcher func(int) (error)) api.Charger {
	switch {
	case meter == nil && meterCurrent == nil && meterEnergy == nil && phaseSwitcher == nil:
		return base

	case meter != nil && meterCurrent == nil && meterEnergy == nil && phaseSwitcher == nil:
		return &struct {
			*OCPP
			api.Meter
		}{
			OCPP: base,
			Meter: &decorateOCPPMeterImpl{
				meter: meter,
			},
		}

	case meter == nil && meterCurrent == nil && meterEnergy != nil && phaseSwitcher == nil:
		return &struct {
			*OCPP
			api.MeterEnergy
		}{
			OCPP: base,
			MeterEnergy: &decorateOCPPMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case meter != nil && meterCurrent == nil && meterEnergy != nil && phaseSwitcher == nil:
		return &struct {
			*OCPP
			api.Meter
			api.MeterEnergy
		}{
			OCPP: base,
			Meter: &decorateOCPPMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decorateOCPPMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case meter == nil && meterCurrent != nil && meterEnergy == nil && phaseSwitcher == nil:
		return &struct {
			*OCPP
			api.MeterCurrent
		}{
			OCPP: base,
			MeterCurrent: &decorateOCPPMeterCurrentImpl{
				meterCurrent: meterCurrent,
			},
		}

	case meter != nil && meterCurrent != nil && meterEnergy == nil && phaseSwitcher == nil:
		return &struct {
			*OCPP
			api.Meter
			api.MeterCurrent
		}{
			OCPP: base,
			Meter: &decorateOCPPMeterImpl{
				meter: meter,
			},
			MeterCurrent: &decorateOCPPMeterCurrentImpl{
				meterCurrent: meterCurrent,
			},
		}

	case meter == nil && meterCurrent != nil && meterEnergy != nil && phaseSwitcher == nil:
		return &struct {
			*OCPP
			api.MeterCurrent
			api.MeterEnergy
		}{
			OCPP: base,
			MeterCurrent: &decorateOCPPMeterCurrentImpl{
				meterCurrent: meterCurrent,
			},
			MeterEnergy: &decorateOCPPMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case meter != nil && meterCurrent != nil && meterEnergy != nil && phaseSwitcher == nil:
		return &struct {
			*OCPP
			api.Meter
			api.MeterCurrent
			api.MeterEnergy
		}{
			OCPP: base,
			Meter: &decorateOCPPMeterImpl{
				meter: meter,
			},
			MeterCurrent: &decorateOCPPMeterCurrentImpl{
				meterCurrent: meterCurrent,
			},
			MeterEnergy: &decorateOCPPMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case meter == nil && meterCurrent == nil && meterEnergy == nil && phaseSwitcher != nil:
		return &struct {
			*OCPP
			api.PhaseSwitcher
		}{
			OCPP: base,
			PhaseSwitcher: &decorateOCPPPhaseSwitcherImpl{
				phaseSwitcher: phaseSwitcher,
			},
		}

	case meter != nil && meterCurrent == nil && meterEnergy == nil && phaseSwitcher != nil:
		return &struct {
			*OCPP
			api.Meter
			api.PhaseSwitcher
		}{
			OCPP: base,
			Meter: &decorateOCPPMeterImpl{
				meter: meter,
			},
			PhaseSwitcher: &decorateOCPPPhaseSwitcherImpl{
				phaseSwitcher: phaseSwitcher,
			},
		}

	case meter == nil && meterCurrent == nil && meterEnergy != nil && phaseSwitcher != nil:
		return &struct {
			*OCPP
			api.MeterEnergy
			api.PhaseSwitcher
		}{
			OCPP: base,
			MeterEnergy: &decorateOCPPMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseSwitcher: &decorateOCPPPhaseSwitcherImpl{
				phaseSwitcher: phaseSwitcher,
			},
		}

	case meter != nil && meterCurrent == nil && meterEnergy != nil && phaseSwitcher != nil:
		return &struct {
			*OCPP
			api.Meter
			api.MeterEnergy
			api.PhaseSwitcher
		}{
			OCPP: base,
			Meter: &decorateOCPPMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decorateOCPPMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseSwitcher: &decorateOCPPPhaseSwitcherImpl{
				phaseSwitcher: phaseSwitcher,
			},
		}

	case meter == nil && meterCurrent != nil && meterEnergy == nil && phaseSwitcher != nil:
		return &struct {
			*OCPP
			api.MeterCurrent
			api.PhaseSwitcher
		}{
			OCPP: base,
			MeterCurrent: &decorateOCPPMeterCurrentImpl{
				meterCurrent: meterCurrent,
			},
			PhaseSwitcher: &decorateOCPPPhaseSwitcherImpl{
				phaseSwitcher: phaseSwitcher,
			},
		}

	case meter != nil && meterCurrent != nil && meterEnergy == nil && phaseSwitcher != nil:
		return &struct {
			*OCPP
			api.Meter
			api.MeterCurrent
			api.PhaseSwitcher
		}{
			OCPP: base,
			Meter: &decorateOCPPMeterImpl{
				meter: meter,
			},
			MeterCurrent: &decorateOCPPMeterCurrentImpl{
				meterCurrent: meterCurrent,
			},
			PhaseSwitcher: &decorateOCPPPhaseSwitcherImpl{
				phaseSwitcher: phaseSwitcher,
			},
		}

	case meter == nil && meterCurrent != nil && meterEnergy != nil && phaseSwitcher != nil:
		return &struct {
			*OCPP
			api.MeterCurrent
			api.MeterEnergy
			api.PhaseSwitcher
		}{
			OCPP: base,
			MeterCurrent: &decorateOCPPMeterCurrentImpl{
				meterCurrent: meterCurrent,
			},
			MeterEnergy: &decorateOCPPMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseSwitcher: &decorateOCPPPhaseSwitcherImpl{
				phaseSwitcher: phaseSwitcher,
			},
		}

	case meter != nil && meterCurrent != nil && meterEnergy != nil && phaseSwitcher != nil:
		return &struct {
			*OCPP
			api.Meter
			api.MeterCurrent
			api.MeterEnergy
			api.PhaseSwitcher
		}{
			OCPP: base,
			Meter: &decorateOCPPMeterImpl{
				meter: meter,
			},
			MeterCurrent: &decorateOCPPMeterCurrentImpl{
				meterCurrent: meterCurrent,
			},
			MeterEnergy: &decorateOCPPMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseSwitcher: &decorateOCPPPhaseSwitcherImpl{
				phaseSwitcher: phaseSwitcher,
			},
		}
	}

	return nil
}

type decorateOCPPMeterImpl struct {
	meter func() (float64, error)
}

func (impl *decorateOCPPMeterImpl) CurrentPower() (float64, error) {
	return impl.meter()
}

type decorateOCPPMeterCurrentImpl struct {
	meterCurrent func() (float64, float64, float64, error)
}

func (impl *decorateOCPPMeterCurrentImpl) Currents() (float64, float64, float64, error) {
	return impl.meterCurrent()
}

type decorateOCPPMeterEnergyImpl struct {
	meterEnergy func() (float64, error)
}

func (impl *decorateOCPPMeterEnergyImpl) TotalEnergy() (float64, error) {
	return impl.meterEnergy()
}

type decorateOCPPPhaseSwitcherImpl struct {
	phaseSwitcher func(int) (error)
}

func (impl *decorateOCPPPhaseSwitcherImpl) Phases1p3p(phases int) (error) {
	return impl.phaseSwitcher(phases)
}
