package util

import "math"

func Latlng2Tile(lat, lng, z float64) (x, y int) {
	x = int((lng / 180 + 1) * math.Pow(2, z) / 2)
	y = int(((-math.Log(math.Tan((45 + lat / 2) * math.Pi / 180)) + math.Pi) * math.Pow(2, z) / (2 * math.Pi)))
	return x, y
}