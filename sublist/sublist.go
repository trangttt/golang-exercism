package sublist


type Relation string

func isSublist(a []int, b []int) bool {
    if len(a) == 0 {
        return true
    }

    for i:=0; i < len(b) - len(a) + 1; i++ {
        count := 0
        for j := i; j < i + len(a); j ++ {
            if b[j] != a[j-i] { break }
            count += 1
        }
        if count == len(a) {
            return true
        }
    }
    return false
}

func Sublist(a []int, b []int) Relation {
    if len(a) == len(b) && isSublist(a, b) {
            return "equal"
    }
    if len(a) < len(b) && isSublist(a, b) {
            return "sublist"
    }
    if len(a) > len(b) && isSublist(b, a) {
            return "superlist"
    }
    return "unequal"
}
