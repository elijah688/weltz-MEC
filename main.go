package main

import (
	"fmt"
	"log"
	"math"
	"sort"

	"github.com/fogleman/gg"
)

func main() {
	const width = 1000
	const height = 1000

	dc := gg.NewContext(width, height)
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	points := []gg.Point{
		{X: -3.705433, Y: 40.417323},
		{X: -3.702666, Y: 40.419283},
		{X: -3.705621, Y: 40.418924},
		{X: -3.704087, Y: 40.416533},
		{X: -3.701512, Y: 40.417971},
		{X: -3.701319, Y: 40.419196},
		{X: -3.701212, Y: 40.420045},
		{X: -3.703272, Y: 40.418804},
		{X: -3.700118, Y: 40.417382},
		{X: -3.700955, Y: 40.414458},
		{X: -3.700633, Y: 40.412351},
		{X: -3.705417, Y: 40.412939},
		{X: -3.703658, Y: 40.414197},
		{X: -3.713311, Y: 40.413625},
		{X: -3.707605, Y: 40.416223},
		{X: -3.711896, Y: 40.416549},
		{X: -3.712067, Y: 40.42132},
		{X: -3.707562, Y: 40.421336},
		{X: -3.713268, Y: 40.41802},
		{X: -3.717344, Y: 40.416043},
	}

	// Get bounding box
	minX, minY, maxX, maxY := points[0].X, points[0].Y, points[0].X, points[0].Y
	for _, p := range points {
		if p.X < minX {
			minX = p.X
		}
		if p.Y < minY {
			minY = p.Y
		}
		if p.X > maxX {
			maxX = p.X
		}
		if p.Y > maxY {
			maxY = p.Y
		}
	}

	scaleX := float64(width) / (maxX - minX)
	scaleY := float64(height) / (maxY - minY)
	scale := math.Min(scaleX, scaleY)

	centerX := (maxX + minX) / 2.0
	centerY := (maxY + minY) / 2.0

	offsetX := (width / 2.0) - (centerX * scale)
	offsetY := (height / 2.0) - (centerY * scale)

	dc.SetRGB(1, 0, 0)
	for _, point := range points {
		scaledX := point.X*scale + offsetX
		scaledY := point.Y*scale + offsetY
		dc.DrawPoint(scaledX, scaledY, 3)
		dc.Fill()
	}

	dc.SetRGB(0, 1, 0)
	dc.NewSubPath()
	for _, p := range points {
		scaledX := p.X*scale + offsetX
		scaledY := p.Y*scale + offsetY
		dc.LineTo(scaledX, scaledY)
	}
	dc.ClosePath()
	dc.Stroke()

	c1 := FindMinEnclosingCircle(points)

	dc.SetRGB(0, 0, 1)
	scaledCX := c1.C.X*scale + offsetX
	scaledCY := c1.C.Y*scale + offsetY
	scaledRadius := c1.R * scale
	dc.DrawCircle(scaledCX, scaledCY, scaledRadius)
	dc.Stroke()

	err := dc.SavePNG("output.png")
	if err != nil {
		log.Fatalf("could not save image: %v", err)
	}

	log.Println("Image successfully saved as output.png")
	fmt.Println(c1)
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
