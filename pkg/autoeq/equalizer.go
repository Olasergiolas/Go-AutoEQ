package autoeq

const (
	defaultEqMode        string  = "IIR"
	defaultSplitChannels bool    = false
	defaultOutputGain    float32 = 0.0
)

type equalizer struct {
	InputGain     float32         `json:"input-gain"`
	OutputGain    float32         `json:"output-gain"`
	Left          map[string]band `json:"left"`
	Right         map[string]band `json:"right"`
	Mode          string          `json:"mode"`
	NumBands      int             `json:"num-bands"`
	SplitChannels bool            `json:"split-channels"`
}

func NewEq(inputGain float32, bands []band) equalizer {
	bandMap := CreateBandMap(bands)
	eq := equalizer{inputGain, defaultOutputGain, bandMap, bandMap, defaultEqMode, len(bandMap), defaultSplitChannels}
	return eq
}
