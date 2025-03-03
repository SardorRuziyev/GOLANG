package transporter

import "fmt"

func RecommendTransporter(destination string, transporters []Transporter) {
	fmt.Printf("%sga yuk jo'natish uchun tavsiyalar:\n", destination)
	for _, t := range transporters {
		switch t.(type) {
		case *Truck:
			fmt.Println(" - Truck: Yirik yuklar uchun yaxshi tanlov.")
		case *Drone:
			fmt.Println(" - Drone: Tez va engil yuklar uchun yaxshi tanlov.")
		case *Ship:
			fmt.Println(" - Ship: Dengiz orqali yuk jo'natish uchun yaxshi tanlov.")
		default:
			fmt.Printf(" - Noma'lum transport turi: %T\n", t)
		}
	}
}
