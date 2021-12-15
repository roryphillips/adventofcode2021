package diagnostics

import (
	"fmt"
	"strconv"
)

type PowerUsage struct {
	Gamma   uint
	Epsilon uint
}

func (s PowerUsage) String() string {
	return fmt.Sprintf("(%v) Gamma: %v, Epsilon: %v",
		s.Gamma*s.Epsilon,
		s.Gamma,
		s.Epsilon)
}

type LifeSupportRating struct {
	OxygenGenerator uint
	CO2Scrubber     uint
}

func (l LifeSupportRating) String() string {
	return fmt.Sprintf("(%v) O2: %v, CO2: %v",
		l.OxygenGenerator*l.CO2Scrubber,
		l.OxygenGenerator,
		l.CO2Scrubber)
}

type Diagnostics interface {
	CalculatePowerUsage(readings ...string) (PowerUsage, error)
	CalculateLifeSupportRating(readings ...string) (LifeSupportRating, error)
}

func NewDiagnostics(bits int) Diagnostics {
	return &diagnosticsImpl{bits: bits}
}

type diagnosticsImpl struct {
	bits int
}

func (s *diagnosticsImpl) CalculatePowerUsage(readings ...string) (PowerUsage, error) {
	onCount := make([]int, s.bits)
	offCount := make([]int, s.bits)

	for _, reading := range readings {
		for i := 0; i < s.bits; i++ {
			if reading[i] == '0' {
				offCount[i] += 1
			} else {
				onCount[i] += 1
			}
		}
	}

	var (
		gammaStr   string = ""
		epsilonStr string = ""
	)

	for i := 0; i < s.bits; i++ {
		// If we have more on than off
		if onCount[i] > offCount[i] {
			gammaStr += "1"
			epsilonStr += "0"
		} else {
			gammaStr += "0"
			epsilonStr += "1"
		}
	}

	gamma, err := strconv.ParseUint(gammaStr, 2, s.bits)
	if err != nil {
		return PowerUsage{}, fmt.Errorf("unable to parse gamma number: %v", err)
	}
	epsilon, err := strconv.ParseUint(epsilonStr, 2, s.bits)
	if err != nil {
		return PowerUsage{}, fmt.Errorf("unable to parse epsilon number: %v", err)
	}

	return PowerUsage{
		Gamma:   uint(gamma),
		Epsilon: uint(epsilon),
	}, nil
}

func (s *diagnosticsImpl) CalculateLifeSupportRating(readings ...string) (LifeSupportRating, error) {
	o2Str := filterCommonality(readings, 0, commonalityFilterMost)[0]
	co2Str := filterCommonality(readings, 0, commonalityFilterLeast)[0]

	o2, err := strconv.ParseUint(o2Str, 2, s.bits)
	if err != nil {
		return LifeSupportRating{}, fmt.Errorf("unable to parse gamma number: %v", err)
	}
	co2, err := strconv.ParseUint(co2Str, 2, s.bits)
	if err != nil {
		return LifeSupportRating{}, fmt.Errorf("unable to parse epsilon number: %v", err)
	}

	return LifeSupportRating{
		OxygenGenerator: uint(o2),
		CO2Scrubber:     uint(co2),
	}, nil
}

type commonalityFilter uint32

const (
	commonalityFilterUnknown commonalityFilter = iota
	commonalityFilterMost
	commonalityFilterLeast
)

func filterCommonality(input []string, bit int, filter commonalityFilter) []string {
	if len(input) == 1 {
		return input
	}

	on := []string{}
	off := []string{}
	for _, reading := range input {
		if reading[bit] == '1' {
			on = append(on, reading)
		} else {
			off = append(off, reading)
		}
	}

	// If we're filtering on the most common values
	if filter == commonalityFilterMost {
		// Default to returning on bits if lengths are level
		if len(on) >= len(off) {
			return filterCommonality(on, bit+1, filter)
		}
		return filterCommonality(off, bit+1, filter)
	} else if filter == commonalityFilterLeast {
		// Default to returning off bits if lengths are level
		if len(on) < len(off) {
			return filterCommonality(on, bit+1, filter)
		}
		return filterCommonality(off, bit+1, filter)
	}

	return []string{}
}
