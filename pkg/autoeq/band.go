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

func NewBand(freq, gain, quality float32, band_type string) band {
	b := band{freq, gain, quality, band_type, defaultBandMode, defaultSilenced, defaultExclusive, defaultSlope}
	return b
}
