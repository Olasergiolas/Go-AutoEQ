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
	freq      float32
	gain      float32
	quality   float32
	band_type BandType
	mode      string
	silenced  bool
	exclusive bool
	slope     string
}
