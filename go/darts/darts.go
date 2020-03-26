package darts

import "math"

func Score(x, y float64) int {
  radius := math.Sqrt(x*x + y*y)
  if radius <= 1 {
    return 10
  } else if radius <= 5 {
    return 5
  } else if radius <= 10 {
    return 1
  }

  return 0
}
