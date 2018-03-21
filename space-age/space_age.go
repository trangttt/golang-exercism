package space

import (
    "fmt"
    "strconv"
)

type Planet string

func Age(time float64, planet Planet) float64 {
    m := make(map[Planet]float64)
    m["Earth"] = 31557600
    m["Mercury"] = 0.2408467 * m["Earth"]
    m["Venus"] = 0.61519726 * m["Earth"]
    m["Mars"] = 1.8808158 * m["Earth"]
    m["Jupiter"] = 11.862615 * m["Earth"]
    m["Saturn"] = 29.447498 * m["Earth"]
    m["Uranus"] = 84.016846 * m["Earth"]
    m["Neptune"] = 164.79132 * m["Earth"]
    result := time / m[planet]
    fmt.Println(time, m[planet])
    result, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", result), 64)
    return result
}
