package temperature

func CtoF(c float64) float64 {
  return (c * 1.8) + 32
}

func FtoC(f float64) float64 {
  return (f - 32) / 1.8
}
