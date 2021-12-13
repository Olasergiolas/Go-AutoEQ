package autoeq

import (
	"os"
	"reflect"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOpenParametricData(t *testing.T) {
	assert := assert.New((t))
	t.Log("Parametric EQ config reading test")

	fixture := "../../docs/examples/AutoEQ_parametric.txt"
	returnedLines, err := OpenParametricData(fixture)
	assert.EqualValues(Success, err, "Error while opening parametric data")

	returnedLinesType := reflect.TypeOf(returnedLines).String()
	wantedLen := 11

	assert.Equal(wantedLen, len(returnedLines), "Wrong number of bands read")
	assert.Equal("[]string", returnedLinesType, "Expected []string but found "+returnedLinesType)
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
	assert.Equal(returnedBandMap[wantedKey].BandType, wantedType, "Error assigning keys to bands")
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

	tmpDir := "/tmp"
	tmpFile := "easyeffects" + strconv.Itoa(os.Getpid())
	err := os.Chdir(tmpDir)
	if err != nil {
		assert.FailNow("Error while trying to cd to GOPATH")
	}

	ExportEasyEffectsProfile(o, tmpFile)

	assert.DirExists("profiles/EasyEffects", "Export path not created!")
	assert.FileExists("profiles/EasyEffects/"+tmpFile+".json", "Exported file not created!")
	dirErr := os.RemoveAll("profiles")
	if dirErr != nil {
		assert.FailNow("Error while cleaning up!")
	}
}
