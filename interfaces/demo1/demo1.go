package interfaces

import (
	"fmt"
	"math"
)

type shape interface {
    area() float64
}

type rectangle struct{
    width , height float64
}

type circle struct{

}