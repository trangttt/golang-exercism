package erratum

func Use(o ResourceOpener, input string) (err error) {
    var res Resource
    var isFailed error
    isTransient := true

    // if transient, repeat
    for isTransient {
        res, isFailed = o()
        if isFailed != nil {
            _, isTransient = isFailed.(TransientError)
        } else {
            isTransient = false
        }
    }

    // maybe some other errors
    if isFailed != nil {
        return isFailed
    }

    defer res.Close()

    defer func() {
        if r := recover(); r != nil {
            frobE, isFrobError := r.(FrobError)
            if isFrobError {
                res.Defrob(frobE.defrobTag)
            }
            err, _ = r.(error)
        }
    }()

    res.Frob(input)
    return err
}
