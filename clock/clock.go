package clock

import (
    "fmt"
)

type Clock struct {
    hour int
    min int
}

func New(_hour int, _min int) Clock {
    v := _hour * 60 + _min
    v %= 24 * 60
    if v < 0 {
        v += 24 * 60
    }
    h, m := v / 60, v % 60
    return Clock{h, m}
}

func (c Clock) Add(min int) Clock {
    return New(c.hour, c.min + min)
}

func (c Clock) Subtract(min int) Clock {
    return New(c.hour, c.min-min)
}

func (c Clock) String() string {
    return fmt.Sprintf("%02d:%02d", c.hour, c.min)
}
