package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(int64(time.Now().Unix()))
	mat := MakeMatrix(4, 0)
	image := MakeImage(500, 500)
	for i := 0; i < 2; i++ {
		mat.AddEdge(rand.Float64()*200, 200-rand.Float64()*200, 0.0, rand.Float64()*200, 200-rand.Float64()*200, 0.0)
	}
	image.DrawLines(mat, Color{r: 0, b: 0, g: 0})
	image.Display()
	transform := MakeTranslate(25, 0, 0)
	scale := MakeScale(2, 1, 1)
	transform, _ = transform.Mult(scale)
	fmt.Println(transform)
	mat, _ = mat.Mult(transform)
	image.DrawLines(mat, Color{r: 0, b: 255, g: 0})
	image.Display()
	image.SavePPM("temp.ppm")

	rotz := MakeRotZ(20.0)
	fmt.Println(mat)
	mat, _ = mat.Mult(rotz)
	fmt.Println(mat)
	image.DrawLines(mat, Color{r: 255, b: 0, g: 0})
	image.Display()
	image.Clear()

	t := MakeMatrix(4, 4)
	t.Ident()
	e := MakeMatrix(4, 0)
	ParseFile("script", t, e, image)
}
