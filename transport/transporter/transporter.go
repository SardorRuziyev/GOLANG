package transporter

type Transporter interface {
	Deliver(destination string, weight float64) string
	CalculateETA(distance float64) float64
}
