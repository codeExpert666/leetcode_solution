package Simulation

func SquareIsWhite(coordinates string) bool {
	x, y := coordinates[0]-'a'+1, coordinates[1]-'0'
	return (x&1)^(y&1) == 1
}
