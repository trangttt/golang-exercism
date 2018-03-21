package letter

type FreqMap map[rune]int

func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(ls []string) FreqMap{
    fc := make(chan FreqMap, len(ls))
    for _, s := range ls {
        go func(_s string){
            m := Frequency(_s)
            fc <- m
        }(s)
    }
    r := FreqMap{}
    for i := 0; i < len(ls); i++{
        rm := <- fc
        for k, v := range rm {
            r[k] += v
        }
    }
    return r
}
