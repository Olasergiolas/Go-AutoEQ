package autoeq

import (
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOpenParametricData(t *testing.T) {
	assert := assert.New((t))
	t.Log("Parametric EQ config reading test")

	fixture := "../../docs/examples/AutoEQ_parametric.txt"
	returnedLines := OpenParametricData(fixture)
	returnedLinesType := reflect.TypeOf(returnedLines).String()
	wantedLen := 11

	assert.Equal(len(returnedLines), wantedLen, "Wrong number of bands read")
	assert.Equal(returnedLinesType, "[]string", "Expected []string but found "+returnedLinesType)
}

func TestCreateBandMap(t *testing.T) {
	assert := assert.New((t))
	t.Log("Band map creation test")

	fixture := []band{
		NewBand(10.0, 5.0, 1, "Bell"),
		NewBand(5.0, 10.0, 0.7, "Low Shelf"),
	}
	returnedBandMap := CreateBandMap(fixture)
	wantedKey := "band1"
	wantedType := "Low Shelf"

	assert.Equal(len(returnedBandMap), len(fixture), "Wrong number of bands in map")
	assert.Equal(returnedBandMap[wantedKey].Band_type, wantedType, "Error assigning keys to bands")
}

func TestGenerateBands(t *testing.T) {
	assert := assert.New((t))
	t.Log("Testing the generation of bands")

	fixture := []string{
		"Filter 1: ON PK Fc 4871 Hz Gain 22.1 dB Q 0.65",
	}
	returnedBands := GenerateBands(fixture)
	wantedBand := NewBand(4871, 22.1, 0.65, "Bell")
	wantedBandsLen := len(fixture)

	assert.Equal(len(returnedBands), wantedBandsLen, "Wrong number of bands generated")
	assert.Equal(returnedBands[0], wantedBand, "Unexpected generated band")
}

func TestExportEasyEffectsProfile(t *testing.T) {
	assert := assert.New((t))
	o := NewOutput(equalizer{})
	t.Log("Testing the creation of a json file")
	ExportEasyEffectsProfile(o, "test")

	assert.DirExists("profiles/EasyEffects", "Export path not created!")
	assert.FileExists("profiles/EasyEffects/test.json", "Export path not created!")
	err := os.RemoveAll("profiles")
	if err != nil {
		assert.FailNow("Error while cleaning up!")
	}
}
