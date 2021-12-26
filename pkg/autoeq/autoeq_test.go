package autoeq

import (
	"bufio"
	"encoding/json"
	"fmt"
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
	assert.Nil(err, "Error while opening parametric data")

	returnedLinesType := reflect.TypeOf(returnedLines).String()
	wantedLen := 11

	assert.Equal(wantedLen, len(returnedLines), "Wrong number of bands read")
	assert.Equal("[]string", returnedLinesType, "Expected []string but found "+returnedLinesType)
}

func TestCreateBandMap(t *testing.T) {
	assert := assert.New((t))
	t.Log("Band map creation test")

	fixture := []band{
		NewBand(maxFreq, maxGain, maxQuality, "Bell"),
		NewBand(minFreq, minGain, minQuality, "Low Shelf"),
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
		"Filter 1: ON PK Fc 4871 Hz Gain 20.0 dB Q 0.65",
	}
	returnedBands := GenerateBands(fixture)
	wantedBand := NewBand(4871, maxGain, 0.65, "Bell")
	wantedBandsLen := len(fixture)

	assert.Equal(wantedBandsLen, len(returnedBands), "Wrong number of bands generated")
	assert.Equal(wantedBand, returnedBands[0], "Unexpected generated band")
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

	err2 := ExportEasyEffectsProfile(o, tmpFile)
	assert.Nil(err2, "Exported path/"+tmpFile+".json not created!")
	dirErr := os.RemoveAll("profiles")
	if dirErr != nil {
		assert.FailNow("Error while cleaning up!")
	}
}

func TestGetConfig(t *testing.T) {
	assert := assert.New(t)
	currentLogPathEnv := os.Getenv(logPathEnvVar)
	config := GetConfig()
	var expectedLogPath string

	if currentLogPathEnv != "" {
		expectedLogPath = currentLogPathEnv
	} else {
		expectedLogPath = defaultLogPath
	}

	assert.Equal(expectedLogPath, config.logpath, "Logpath is not being configured as expected")
}

func TestLogWrapper(t *testing.T) {
	assert := assert.New(t)
	logMsg := "testing log msg"
	var testLogEntry map[string]interface{}
	var testLog []map[string]interface{}

	testLogPath := os.Getenv(logPathEnvVar)
	assert.NotEmpty(testLogPath, "Test LogPath env-var not set")

	logger := GetLogger()
	logger.SuccessLog(logMsg)

	f, err := os.Open(testLogPath)
	assert.Nilf(err, "Error while opening test log %s", testLogPath)

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		json.Unmarshal([]byte(sc.Bytes()), &testLogEntry)
		testLog = append(testLog, testLogEntry)
	}

	actualMsg := testLog[len(testLog)-1]["message"]
	expectedMsg := fmt.Sprintf(successEvent.msg, logMsg)
	assert.EqualValuesf(expectedMsg, actualMsg, "Couldn't find expected log in %s", testLogPath)

	defer f.Close()
}
