package autoeq

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	preampPos   int = 1
	bandTypePos int = 3
	freqPos     int = 5
	gainPos     int = 8
	qualityPos  int = 11
)

func CreateBandMap(bands []band) map[string]band {
	m := make(map[string]band)

	for i := 0; i < len(bands); i++ {
		m["band"+strconv.Itoa(i)] = bands[i]
	}
	return m
}

func OpenParametricData(path string) []string {
	data, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()

	sc := bufio.NewScanner(data)
	var lines []string

	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines
}

func GetPreamp(configHeadline string) float32 {
	lSplit := strings.Split(configHeadline, " ")
	preamp := lSplit[preampPos]
	preampFloat, _ := strconv.ParseFloat(preamp, 32)

	return float32(preampFloat)
}

func GenerateBands(lines []string) []band {
	var b band
	var bands []band
	var freq, gain, quality float64
	var easyEffectsBandTypes = map[string]string{
		"PK": "Bell",
		"LS": "Low Shelf",
		"HS": "High Shelf",
	}

	for _, l := range lines {
		lSplit := strings.Split(l, " ")
		freq, _ = strconv.ParseFloat(lSplit[freqPos], 32)
		gain, _ = strconv.ParseFloat(lSplit[gainPos], 32)
		quality, _ = strconv.ParseFloat(lSplit[qualityPos], 32)

		b = NewBand(float32(freq), float32(gain), float32(quality), easyEffectsBandTypes[lSplit[bandTypePos]])
		bands = append(bands, b)
	}
	return bands
}
