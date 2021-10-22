package AutoEQ

type Tipo string

const (
	Bell Tipo = "Bell"
)

type banda struct {
	freq      float32
	gain      float32
	mode      string
	silenced  bool
	quality   float32
	exclusive bool
	slope     string
	band_type Tipo
}
