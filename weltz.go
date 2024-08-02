package main

import (
	"math"
	"math/rand"
	"time"

	"github.com/fogleman/gg"
)

const INF = 1e18

type Circle struct {
	C gg.Point
	R float64
}

func dist(a, b gg.Point) float64 {
	return math.Sqrt(math.Pow(a.X-b.X, 2) + math.Pow(a.Y-b.Y, 2))
}

func isInside(c Circle, p gg.Point) bool {
	return dist(c.C, p) <= c.R
}

func getCircleCenter(bx, by, cx, cy float64) gg.Point {
	B := bx*bx + by*by
	C := cx*cx + cy*cy
	D := bx*cy - by*cx
	return gg.Point{
		X: (cy*B - by*C) / (2 * D),
		Y: (bx*C - cx*B) / (2 * D),
	}
}

func circleFrom3Points(A, B, C gg.Point) Circle {
	I := getCircleCenter(B.X-A.X, B.Y-A.Y, C.X-A.X, C.Y-A.Y)
	I.X += A.X
	I.Y += A.Y
	return Circle{I, dist(I, A)}
}

func circleFrom2Points(A, B gg.Point) Circle {
	C := gg.Point{(A.X + B.X) / 2.0, (A.Y + B.Y) / 2.0}
	return Circle{C, dist(A, B) / 2.0}
}

func isValidCircle(c Circle, P []gg.Point) bool {
	for _, p := range P {
		if !isInside(c, p) {
			return false
		}
	}
	return true
}

func minCircleTrivial(P []gg.Point) Circle {
	switch len(P) {
	case 0:
		return Circle{gg.Point{0, 0}, 0}
	case 1:
		return Circle{P[0], 0}
	case 2:
		return circleFrom2Points(P[0], P[1])
	}

	for i := 0; i < 3; i++ {
		for j := i + 1; j < 3; j++ {
			c := circleFrom2Points(P[i], P[j])
			if isValidCircle(c, P) {
				return c
			}
		}
	}
	return circleFrom3Points(P[0], P[1], P[2])
}

func welzlHelper(P []gg.Point, R []gg.Point, n int) Circle {
	if n == 0 || len(R) == 3 {
		return minCircleTrivial(R)
	}

	idx := rand.Intn(n)
	p := P[idx]
	P[idx], P[n-1] = P[n-1], P[idx]

	d := welzlHelper(P, R, n-1)

	if isInside(d, p) {
		return d
	}

	R = append(R, p)
	return welzlHelper(P, R, n-1)
}

func FindMinEnclosingCircle(points []gg.Point) Circle {
	shuffle(points)
	return welzlHelper(points, nil, len(points))
}

func shuffle(points []gg.Point) {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range points {
		j := rand.Intn(i + 1)
		points[i], points[j] = points[j], points[i]
	}
}
