package grains

import (
    "math"
    "errors"
)

func Square(n int) (uint64, error) {
    if n < 1 || n > 64 {
        return 0, errors.New("Input must be within 1 and 64 inclusively")
    } else {
        n -= 1
        return uint64(math.Pow(2, float64(n))), nil
    }
}

func Total() uint64 {
    var ret uint64 = 0
    for i := 1; i < 65; i++ {
        v, _ := Square(i)
        ret += v
    }
    return ret
}
