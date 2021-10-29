package autoeq

type BandType string

const (
	Off       BandType = "Off"
	Bell      BandType = "Bell"
	HighPass  BandType = "High Pass"
	HighShelf BandType = "High Shelf"
	LowPass   BandType = "Low Pass"
	LowShelf  BandType = "Low Shelf"
	Notch     BandType = "Notch"
	Resonance BandType = "Resonance"

	DefaultBandMode  string = "RLC (BT)"
	DefaultSlope     string = "x1"
	DefaultSilenced  bool   = false
	DefaultExclusive bool   = false

	MaxFreq    float32 = 24000.0
	MinFreq    float32 = 10.0
	MaxGain    float32 = 20.0
	MinGain    float32 = -20.0
	MaxQuality float32 = 100.0
	MinQuality float32 = 0.0
)

type band struct {
	Freq      float32
	Gain      float32
	Quality   float32
	Band_type BandType
	Mode      string
	Silenced  bool
	Exclusive bool
	Slope     string
}

func NewBand(freq, gain, quality float32, band_type BandType) band {
	b := band{freq, gain, quality, band_type, DefaultBandMode, DefaultSilenced, DefaultExclusive, DefaultSlope}
	return b
}
