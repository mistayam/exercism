//Package for space functions
package space

//Import  package for error checking
import ( 
	"errors"
)

//Type Planet is a string
type Planet string

const earthOrbPeriod = 31557600

// A function that converts age in seconds and name of planet
//  and returns a given age in years relative to given planet
func Age(ageSeconds float64, planet Planet) float64 {

	//Variable to store each planet's orbital period in relation to
	//Earth's orbital period
	var planetFactor float64

	//Maps orbital period to corresponding planet
	switch planet {
	case "Mercury":
		planetFactor = 0.2408467
		break
	case "Venus":
		planetFactor = 0.61519726
		break
	case "Earth":
		planetFactor = 1
		break
	case "Mars":
		planetFactor = 1.8808158
		break
	case "Jupiter":
		planetFactor = 11.862615
		break
	case "Saturn":
		planetFactor = 29.447498
		break
	case "Uranus":
		planetFactor = 84.016846
		break
	case "Neptune":
		planetFactor = 164.79132
		break
	default:
		panic(errors.New("Invalid Planet detected."))
	}

	//Returns computed age in years for specified planet
	return ageSeconds / (earthOrbPeriod * planetFactor)
}