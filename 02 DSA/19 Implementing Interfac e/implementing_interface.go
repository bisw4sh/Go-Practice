package main
import "fmt"

type catalog interface { 
   shipping() float64 
   tax() float64
} 
 
type configurable struct { 
   name string
   price, qty float64   
} 

func (c *configurable) tax() float64{
  return c.price * c.qty * 0.05
}

func (c *configurable) shipping() float64{
  return c.qty * 5
}

func  main() {
  tshirt := configurable{}
  tshirt.price = 250
  tshirt.qty = 2
  fmt.Println("Shipping Charge: ", tshirt.shipping())
  fmt.Println("Tax: ", tshirt.tax())  
}