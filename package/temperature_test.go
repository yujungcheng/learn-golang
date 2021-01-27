package temperature

import (
  "testing"
)

type temperatureTest struct {
  input float64
  expected float64
}

var CtoFTests = []temperatureTest {
  {4.1, 39.38},
  {10, 50},
  {-10, 14},
}

var FtoCTests = []temperatureTest {
  {32, 0},
  {50, 10},
  {5, -15},
}

func TestCtoF(t *testing.T) {
  for _, tt := range CtoFTests {
    actual := CtoF(tt.input)
    if actual != tt.expected {
      t.Errorf("expected %v, actual %v", tt.expected, actual)
    }
  }
}

func TestFtoC(t *testing.T) {
  for _, tt := range FtoCTests {
    actual := FtoC(tt.input)
    if actual != tt.expected {
      t.Errorf("expected %v, actual %v", tt.expected, actual)
    }
  }
}
