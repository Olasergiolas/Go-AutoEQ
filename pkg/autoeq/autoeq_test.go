package autoeq

import (
	"reflect"
	"testing"
)

func TestOpenParametricData(t *testing.T) {
	t.Log("Parametric EQ config reading test")

	fixture := "../../docs/examples/AutoEQ_parametric.txt"
	returnedLines := OpenParametricData(fixture)
	returnedLinesType := reflect.TypeOf(returnedLines).String()
	wantedLen := 11

	if len(returnedLines) != wantedLen {
		t.Error("Wrong number of bands read")
	}

	if returnedLinesType != "[]string" {
		t.Error("Expected []string but found " + returnedLinesType)
	}
}

func TestCreateBandMap(t *testing.T) {
	t.Log("Band map creation test")

	fixture := []band{
		NewBand(10.0, 5.0, 1, "Bell"),
		NewBand(5.0, 10.0, 0.7, "Low Shelf"),
	}
	returnedBandMap := CreateBandMap(fixture)
	wantedKey := "band1"
	wantedType := "Low Shelf"

	if len(returnedBandMap) != len(fixture) {
		t.Error("Wrong number of bands in map")
	}

	if returnedBandMap[wantedKey].Band_type != wantedType {
		t.Error("Error assigning keys to bands")
	}
}

func TestGenerateBands(t *testing.T) {
	t.Log("Testing the generation of bands")

	fixture := []string{
		"Filter 1: ON PK Fc 4871 Hz Gain 22.1 dB Q 0.65",
	}
	returnedBands := GenerateBands(fixture)
	wantedBand := NewBand(4871, 22.1, 0.65, "Bell")
	wantedBandsLen := len(fixture)

	if len(returnedBands) != wantedBandsLen {
		t.Error("Wrong number of bands generated")
	}

	if returnedBands[0] != wantedBand {
		t.Error("Unexpected generated band")
	}
}
