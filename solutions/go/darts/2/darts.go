package darts
import(
    "math"
)

func Score(x, y float64) int {
	//panic("Please implement the Score function")
    dis := math.Sqrt((x*x) + (y*y))
    
    switch{
        case dis <= 1 : return 10
        case dis <= 5 : return 5
        case dis <= 10 : return 1
        default : return 0
    }

    return -1
}
