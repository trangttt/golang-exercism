package robotname

import (
    "math/rand"
    "fmt"
)

type Robot struct {
    name string
}

var m = make(map[string]struct{})

func getName() string {
    n := int('Z' - 'A' + 1)
    ret := []rune("12345")
    ret[0] = rune(rand.Intn(n) + 'A')
    ret[1] = rune(rand.Intn(n) + 'A')
    ret[2] = rune(rand.Intn(10) + '0')
    ret[3] = rune(rand.Intn(10) + '0')
    ret[4] = rune(rand.Intn(10) + '0')
    return string(ret)
}

func (r *Robot) Name() string {
    if r.name == "" {
        var name string
        exist := true
        for {
            if !exist {
                break
            }
            name = getName()
            _, exist = m[name]
        }
        r.name = name
        var s struct{}
        m[name] = s
    }
    fmt.Println(r.name)
    return r.name
}

func (r *Robot) Reset() {
    exist := true
    var name string
    for {
        if !exist {
            break
        }
        name = getName()
        _, exist = m[name]
    }
    fmt.Println(name)
    r.name = name
    var s struct{}
    m[name] = s
}
