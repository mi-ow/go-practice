package darts
import(
    "math"
)

func Score(x, y float64) int {
	//panic("Please implement the Score function")
    dis := math.Sqrt((x*x) + (y*y))
    
    if dis <=10 && dis > 5 {
        return 1
    }
    if dis <= 5 && dis > 1 {
        return 5
    }
    if dis <= 1 {
        return 10
    }
    return 0;
}
