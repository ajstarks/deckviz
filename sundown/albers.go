// Package main converts WGS84 (lon/lat) coordinates to the
// USA Contiguous Albers Equal Area Conic projection (EPSG:5070).
//
// Standard parallels: 29.5°N and 45.5°N
// Central meridian:  -96.0°
// Latitude of origin:  23.0°N
// False easting/northing: 0 / 0  (metres)
//
// Reference: Snyder, J.P. (1987) Map Projections – A Working Manual, p. 98-103.

package main

import (
	"fmt"
	"math"
)

// Albers holds the parameters that define an Albers Equal-Area Conic projection.
type Albers struct {
	phi0, phi1, phi2 float64 // origin lat, std parallel 1 & 2 (radians)
	lam0             float64 // central meridian (radians)
	a, e, e2         float64 // semi-major axis, eccentricity, eccentricity²
	n, C, rho0       float64 // projection constants
}

// newAlbers builds an Albers projection from degrees.
// Uses the GRS80 / WGS84 ellipsoid (a = 6 378 137 m, f = 1/298.257222101).
func newAlbers(phi0Deg, phi1Deg, phi2Deg, lam0Deg float64) Albers {
	const (
		aWGS84 = 6_378_137.0
		fWGS84 = 1.0 / 298.257223563
	)
	toRad := math.Pi / 180.0

	a := aWGS84
	f := fWGS84
	e2 := 2*f - f*f
	e := math.Sqrt(e2)

	phi0 := phi0Deg * toRad
	phi1 := phi1Deg * toRad
	phi2 := phi2Deg * toRad
	lam0 := lam0Deg * toRad

	// q(phi) – authalic latitude helper (Snyder eq. 3-12)
	q := func(phi float64) float64 {
		sinPhi := math.Sin(phi)
		eSinPhi := e * sinPhi
		return (1 - e2) * (sinPhi/(1-eSinPhi*eSinPhi) -
			(1/(2*e))*math.Log((1-eSinPhi)/(1+eSinPhi)))
	}

	// m(phi) – (Snyder eq. 14-15)
	m := func(phi float64) float64 {
		sinPhi := math.Sin(phi)
		return math.Cos(phi) / math.Sqrt(1-e2*sinPhi*sinPhi)
	}

	m1 := m(phi1)
	m2 := m(phi2)
	q0 := q(phi0)
	q1 := q(phi1)
	q2 := q(phi2)

	var n float64
	if math.Abs(phi1-phi2) > 1e-10 {
		n = (m1*m1 - m2*m2) / (q2 - q1)
	} else {
		sinPhi1 := math.Sin(phi1)
		n = sinPhi1
	}

	C := m1*m1 + n*q1
	rho0 := a * math.Sqrt(C-n*q0) / n

	return Albers{
		phi0: phi0, phi1: phi1, phi2: phi2, lam0: lam0,
		a: a, e: e, e2: e2,
		n: n, C: C, rho0: rho0,
	}
}

// Forward projects a WGS84 (lon, lat) point in degrees to (easting, northing) in metres.
func (alb Albers) Forward(lonDeg, latDeg float64) (x, y float64) {
	toRad := math.Pi / 180.0
	phi := latDeg * toRad
	lam := lonDeg * toRad

	sinPhi := math.Sin(phi)
	eSinPhi := alb.e * sinPhi
	q := (1 - alb.e2) * (sinPhi/(1-eSinPhi*eSinPhi) -
		(1/(2*alb.e))*math.Log((1-eSinPhi)/(1+eSinPhi)))

	rho := alb.a * math.Sqrt(alb.C-alb.n*q) / alb.n
	theta := alb.n * (lam - alb.lam0)

	x = rho*math.Sin(theta)
	y = alb.rho0 - rho*math.Cos(theta)
	return
}

// Point holds an input WGS84 coordinate and its projected output.
type Point struct {
	Lon, Lat float64
	X, Y     float64
}

func main() {
	// USA Contiguous Albers Equal Area Conic (EPSG:5070)
	proj := newAlbers(
		23.0,  // latitude of origin
		29.5,  // standard parallel 1
		45.5,  // standard parallel 2
		-96.0, // central meridian
	)

	// Example input: WGS84 (longitude, latitude) pairs
	inputs := []Point{
		{Lon: -122.4194, Lat: 37.7749},  // San Francisco, CA
		{Lon: -87.6298, Lat: 41.8781},   // Chicago, IL
		{Lon: -73.9857, Lat: 40.7484},   // New York, NY
		{Lon: -104.9903, Lat: 39.7392},  // Denver, CO
		{Lon: -97.7431, Lat: 30.2672},   // Austin, TX
	}

	fmt.Printf("%-30s %-30s %-20s %-20s\n",
		"Location (lon, lat)", "", "Easting (m)", "Northing (m)")
	fmt.Printf("%-15s %-15s %-20s %-20s\n",
		"Longitude", "Latitude", "X", "Y")
	fmt.Println(repeat("-", 72))

	for _, p := range inputs {
		p.X, p.Y = proj.Forward(p.Lon, p.Lat)
		fmt.Printf("%-15.6f %-15.6f %-20.3f %-20.3f\n",
			p.Lon, p.Lat, p.X, p.Y)
	}

	// --- demonstrate reading from stdin ---
	fmt.Println("\nEnter coordinates as 'lon lat' (Ctrl-D / Ctrl-Z to quit):")
	for {
		var lon, lat float64
		_, err := fmt.Scanf("%f %f", &lon, &lat)
		if err != nil {
			break
		}
		x, y := proj.Forward(lon, lat)
		fmt.Printf("  Albers  X = %.3f m,  Y = %.3f m\n", x, y)
	}
}

func repeat(s string, n int) string {
	out := ""
	for i := 0; i < n; i++ {
		out += s
	}
	return out
}
