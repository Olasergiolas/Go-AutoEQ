package autoeq

const (
	defaultBandMode  string = "RLC (BT)"
	defaultSlope     string = "x1"
	defaultSilenced  bool   = false
	defaultExclusive bool   = false

	maxFreq    float32 = 24000.0
	minFreq    float32 = 10.0
	maxGain    float32 = 20.0
	minGain    float32 = -20.0
	maxQuality float32 = 100.0
	minQuality float32 = 0.0
)

type band struct {
	Freq      float32
	Gain      float32
	Quality   float32
	BandType  string `json:"Band_type"`
	Mode      string
	Silenced  bool
	Exclusive bool
	Slope     string
}

func correctOutOfBounds(arg *float32, lower, upper float32) bool {
	res := true

	if *arg < lower {
		*arg = lower
	} else if *arg > upper {
		*arg = upper
	} else {
		res = false
	}

	return res
}

func NewBand(freq, gain, quality float32, band_type string) band {
	logger := NewLogger()

	if correctOutOfBounds(&freq, minFreq, maxFreq) {
		logger.InvalidArgRangeLog("freq")
	} else if correctOutOfBounds(&gain, minGain, maxGain) {
		logger.InvalidArgRangeLog("gain")
	} else if correctOutOfBounds(&quality, minQuality, maxQuality) {
		logger.InvalidArgRangeLog("quality")
	}

	b := band{freq, gain, quality, band_type, defaultBandMode, defaultSilenced, defaultExclusive, defaultSlope}
	return b
}
