package lazytz

import (
	"fmt"
	"strings"
	"time"
	_ "time/tzdata"
)

const (
	ZoneAbbreviationCST  = "CST"
	ZoneAbbreviationPDT  = "PDT"
	ZoneAbbreviationAKDT = "AKDT"
	ZoneAbbreviationHAST = "HAST"
	ZoneAbbreviationHADT = "HADT"
	ZoneAbbreviationEST  = "EST"
	ZoneAbbreviationEDT  = "EDT"
	ZoneAbbreviationCDT  = "CDT"
	ZoneAbbreviationMST  = "MST"
	ZoneAbbreviationMDT  = "MDT"
	ZoneAbbreviationPST  = "PST"
	ZoneAbbreviationAKST = "AKST"
)

type AbbreviatedUSAZone struct {
	Name     string
	Offset   int
	Location time.Location

	IsDaylight           bool
	IsStandard           bool
	DaylightAbbreviation string
	StandardAbbreviation string
}

var (
	abbreviationList = []string{
		ZoneAbbreviationMDT,
		ZoneAbbreviationPST,
		ZoneAbbreviationAKST,
		ZoneAbbreviationEST,
		ZoneAbbreviationEDT,
		ZoneAbbreviationCDT,
		ZoneAbbreviationMST,
		ZoneAbbreviationHADT,
		ZoneAbbreviationCST,
		ZoneAbbreviationPDT,
		ZoneAbbreviationAKDT,
		ZoneAbbreviationHAST,
	}

	abbreviationMap = make(map[string]AbbreviatedUSAZone)
)

func init() {

	if tl, err := time.LoadLocation("America/Boise"); err != nil {
		panic(fmt.Sprintf("error parsing \"America/Boise\" into time.Location: %v", err))
	} else {
		abbreviationMap[ZoneAbbreviationMST] = AbbreviatedUSAZone{
			Name:                 "America/Boise",
			Offset:               -25200,
			Location:             *tl,
			IsDaylight:           false,
			IsStandard:           true,
			DaylightAbbreviation: "MDT",
			StandardAbbreviation: "MST",
		}
	}

	if tl, err := time.LoadLocation("America/Boise"); err != nil {
		panic(fmt.Sprintf("error parsing \"America/Boise\" into time.Location: %v", err))
	} else {
		abbreviationMap[ZoneAbbreviationMDT] = AbbreviatedUSAZone{
			Name:                 "America/Boise",
			Offset:               -21600,
			Location:             *tl,
			IsDaylight:           true,
			IsStandard:           false,
			DaylightAbbreviation: "MDT",
			StandardAbbreviation: "MST",
		}
	}

	if tl, err := time.LoadLocation("America/Los_Angeles"); err != nil {
		panic(fmt.Sprintf("error parsing \"America/Los_Angeles\" into time.Location: %v", err))
	} else {
		abbreviationMap[ZoneAbbreviationPST] = AbbreviatedUSAZone{
			Name:                 "America/Los_Angeles",
			Offset:               -28800,
			Location:             *tl,
			IsDaylight:           false,
			IsStandard:           true,
			DaylightAbbreviation: "PDT",
			StandardAbbreviation: "PST",
		}
	}

	if tl, err := time.LoadLocation("America/Anchorage"); err != nil {
		panic(fmt.Sprintf("error parsing \"America/Anchorage\" into time.Location: %v", err))
	} else {
		abbreviationMap[ZoneAbbreviationAKST] = AbbreviatedUSAZone{
			Name:                 "America/Anchorage",
			Offset:               -32400,
			Location:             *tl,
			IsDaylight:           false,
			IsStandard:           true,
			DaylightAbbreviation: "AKDT",
			StandardAbbreviation: "AKST",
		}
	}

	if tl, err := time.LoadLocation("America/Detroit"); err != nil {
		panic(fmt.Sprintf("error parsing \"America/Detroit\" into time.Location: %v", err))
	} else {
		abbreviationMap[ZoneAbbreviationEST] = AbbreviatedUSAZone{
			Name:                 "America/Detroit",
			Offset:               -18000,
			Location:             *tl,
			IsDaylight:           false,
			IsStandard:           true,
			DaylightAbbreviation: "EDT",
			StandardAbbreviation: "EST",
		}
	}

	if tl, err := time.LoadLocation("America/Detroit"); err != nil {
		panic(fmt.Sprintf("error parsing \"America/Detroit\" into time.Location: %v", err))
	} else {
		abbreviationMap[ZoneAbbreviationEDT] = AbbreviatedUSAZone{
			Name:                 "America/Detroit",
			Offset:               -14400,
			Location:             *tl,
			IsDaylight:           true,
			IsStandard:           false,
			DaylightAbbreviation: "EDT",
			StandardAbbreviation: "EST",
		}
	}

	if tl, err := time.LoadLocation("America/Chicago"); err != nil {
		panic(fmt.Sprintf("error parsing \"America/Chicago\" into time.Location: %v", err))
	} else {
		abbreviationMap[ZoneAbbreviationCDT] = AbbreviatedUSAZone{
			Name:                 "America/Chicago",
			Offset:               -18000,
			Location:             *tl,
			IsDaylight:           true,
			IsStandard:           false,
			DaylightAbbreviation: "CDT",
			StandardAbbreviation: "CST",
		}
	}

	if tl, err := time.LoadLocation("America/Adak"); err != nil {
		panic(fmt.Sprintf("error parsing \"America/Adak\" into time.Location: %v", err))
	} else {
		abbreviationMap[ZoneAbbreviationHAST] = AbbreviatedUSAZone{
			Name:                 "America/Adak",
			Offset:               -36000,
			Location:             *tl,
			IsDaylight:           false,
			IsStandard:           true,
			DaylightAbbreviation: "HADT",
			StandardAbbreviation: "HAST",
		}
	}

	if tl, err := time.LoadLocation("America/Adak"); err != nil {
		panic(fmt.Sprintf("error parsing \"America/Adak\" into time.Location: %v", err))
	} else {
		abbreviationMap[ZoneAbbreviationHADT] = AbbreviatedUSAZone{
			Name:                 "America/Adak",
			Offset:               -32400,
			Location:             *tl,
			IsDaylight:           true,
			IsStandard:           false,
			DaylightAbbreviation: "HADT",
			StandardAbbreviation: "HAST",
		}
	}

	if tl, err := time.LoadLocation("America/Chicago"); err != nil {
		panic(fmt.Sprintf("error parsing \"America/Chicago\" into time.Location: %v", err))
	} else {
		abbreviationMap[ZoneAbbreviationCST] = AbbreviatedUSAZone{
			Name:                 "America/Chicago",
			Offset:               -21600,
			Location:             *tl,
			IsDaylight:           false,
			IsStandard:           true,
			DaylightAbbreviation: "CDT",
			StandardAbbreviation: "CST",
		}
	}

	if tl, err := time.LoadLocation("America/Los_Angeles"); err != nil {
		panic(fmt.Sprintf("error parsing \"America/Los_Angeles\" into time.Location: %v", err))
	} else {
		abbreviationMap[ZoneAbbreviationPDT] = AbbreviatedUSAZone{
			Name:                 "America/Los_Angeles",
			Offset:               -25200,
			Location:             *tl,
			IsDaylight:           true,
			IsStandard:           false,
			DaylightAbbreviation: "PDT",
			StandardAbbreviation: "PST",
		}
	}

	if tl, err := time.LoadLocation("America/Anchorage"); err != nil {
		panic(fmt.Sprintf("error parsing \"America/Anchorage\" into time.Location: %v", err))
	} else {
		abbreviationMap[ZoneAbbreviationAKDT] = AbbreviatedUSAZone{
			Name:                 "America/Anchorage",
			Offset:               -28800,
			Location:             *tl,
			IsDaylight:           true,
			IsStandard:           false,
			DaylightAbbreviation: "AKDT",
			StandardAbbreviation: "AKST",
		}
	}
}

// Get attempts to return to you timezone info based on an abbreviation with the understanding that this only handles
// abbreviations within the United States
func Get(abbr string) (AbbreviatedUSAZone, bool) {
	abbr = strings.ToUpper(abbr)
	info, ok := abbreviationMap[abbr]
	return info, ok
}

// GetAbbreviations returns all currently known USA timezone abbreviations to you
func GetAbbreviations() []string {
	out := make([]string, len(abbreviationList))
	copy(out, abbreviationList)
	return out
}

// GetAll returns a keyed map of all known USA timezone abbreviations and their details
func GetAll() map[string]AbbreviatedUSAZone {
	out := make(map[string]AbbreviatedUSAZone, len(abbreviationMap))
	for k, v := range abbreviationMap {
		out[k] = v
	}
	return out
}
