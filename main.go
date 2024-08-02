package main

import (
	"log"
	"math"
	"sort"

	"github.com/fogleman/gg"
)

func main() {
	const width = 1024
	const height = 1024

	dc := gg.NewContext(width, height)
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	dc.SetRGB(0, 0, 0)

	points := []gg.Point{
		{X: 120, Y: 450},
		{X: 120, Y: 400},
		{X: 150, Y: 400},
		{X: 180, Y: 450},
		{X: 170, Y: 450},
		{X: 200, Y: 400},
		{X: 230, Y: 450},
		{X: 260, Y: 400},
		{X: 260, Y: 350},
		{X: 230, Y: 300},
		{X: 150, Y: 350},
		{X: 666, Y: 432},
		{X: 153, Y: 225},
	}

	dc.SetRGB(1, 0, 0)
	for _, point := range points {
		dc.DrawPoint(point.X, point.Y, 3)
		dc.Fill()
	}

	dc.SetRGB(0, 1, 0)
	dc.NewSubPath()
	for _, p := range sortPoints(points) {
		dc.LineTo(p.X, p.Y)
	}
	dc.ClosePath()
	dc.Stroke()

	c1 := FindMinEnclosingCircle(points)

	circles := []struct {
		X, Y, Radius float64
	}{
		{X: c1.C.X, Y: c1.C.Y, Radius: c1.R},
	}

	dc.SetRGB(0, 0, 1)
	for _, circle := range circles {
		dc.DrawCircle(circle.X, circle.Y, circle.Radius)
		dc.Stroke()
	}

	err := dc.SavePNG("output.png")
	if err != nil {
		log.Fatalf("could not save image: %v", err)
	}

	log.Println("Image successfully saved as output.png")
}

func centroid(points []gg.Point) gg.Point {
	var c gg.Point
	for _, p := range points {
		c.X += p.X
		c.Y += p.Y
	}
	n := float64(len(points))
	c.X /= n
	c.Y /= n
	return c
}

func angle(p, c gg.Point) float64 {
	return math.Atan2(p.Y-c.Y, p.X-c.X)
}

func sortPoints(points []gg.Point) []gg.Point {
	c := centroid(points)
	sort.Slice(points, func(i, j int) bool {
		return angle(points[i], c) < angle(points[j], c)
	})
	return points
}
