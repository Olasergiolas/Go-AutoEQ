package AutoEQ

const (
	Bell string = "Bell"
)

type banda struct {
	freq      float32
	gain      float32
	mode      string
	silenced  bool
	quality   float32
	exclusive bool
	slope     string
}
