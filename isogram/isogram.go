package isogram

func IsIsogram(input string) bool {
    m := make(map[byte]struct{})
    for i:=0; i < len(input); i++ {
        v := input[i]
        if v > 90 {
            v -= 32
        }

        if 65 <= v && v <= 90 {
            _, exist := m[v]
            if exist {
                return false
            } else {
                var s struct{}
                m[v] = s
            }
        }
    }
    return true
}
