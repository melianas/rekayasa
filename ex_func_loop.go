package main

import(
  "fmt"
  "math"
)

func Sqrt(x float64) float64{
  z := float64(1)
  s := 1.0

  for {
    z = z-(z*z - x)/(2*z)
    if math.Abs(s-z) < 1e-15{
      break
    }
    s=z
  }
  return s

}

func main(){
  fmt.Println("sqrt dengan function Sqrt:")
  fmt.Println(Sqrt(2))

  fmt.Println("sqrt dengan math.Sqrt:")
  fmt.Println(math.Sqrt(2))
}
