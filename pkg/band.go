package autoeq

type BandType string

const (
	Bell      BandType = "Bell"
	HighPass  BandType = "High Pass"
	HighShelf BandType = "High Shelf"
	LowPass   BandType = "Low Pass"
	LowShelf  BandType = "Low Shelf"
	Notch     BandType = "Notch"
	Resonance BandType = "Resonance"

	DefaultMode      string = "RLC (BT)"
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

type banda struct {
	freq      float32
	gain      float32
	mode      string
	silenced  bool
	quality   float32
	exclusive bool
	slope     string
	band_type BandType
}
