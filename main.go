package main

type Turn struct {
	direction string
	heading   int
	speed     Speed
}
type Speed struct {
	unit     string
	velocity int
}

type Runway struct {
	number int
}

type Taxi struct {
	holdingPoint string
	runway       Runway
}

type Cross struct {
	runway   Runway
	location string
}
type Contact struct {
	frequency int
	target    string
}

func main() {

}
