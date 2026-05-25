package main

import (
	"fmt"
	"math/rand"
)

func randInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func arvorada (pen *Pen, dist float64 ) {

	ang := 25.0
	fator := 0.77
	
	if dist < 10 {
		// pen.DrawCircle(5)
		pen.FillCircle(7.5)
		return

	}

	angulodir := ang - float64(randInt(-5,5))
	anguloesq := ang + float64(randInt(-5,5))
	pen.SetLineWidth(dist / 20)
	pen.Walk(dist)
	pen.Right(angulodir)
	arvorada(pen, dist * fator)
	pen.Left(angulodir + anguloesq)
	arvorada(pen, dist * fator)
	pen.Right(anguloesq)
	pen.Walk(-dist)
	
}

func main() {
	pen := NewPen(1000, 1000)
	pen.SetRGB(70, 30, 255)     
	pen.SetPosition(500, 1000)
	pen.SetHeading(90)        

	arvorada(pen, 200.0)

	pen.SavePNG("tree.png")
	fmt.Println("PNG file created successfully.")
}
