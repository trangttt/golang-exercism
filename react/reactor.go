package react

type MCell struct {
	val int
	callbacks map[int]func(int)
	deps []MDepCell
	reactor *MReactor
}

type MDepCell struct {
	*MCell
	update func(int) int
}

type MReactor struct {
	oldValues map[*MCell]int
}

func (c *MCell) Value() int {
	return c.val
}

func (c *MCell) SetValue(newVal int) {
	// bookkeeping old values
	c.SaveValue()

	// propagate values changes across all dependants
	c.UpdateValue(newVal)

	// Perform callbacks only when values are all updated
	c.PerformCallbacks()

	//clear bookkeeping
	c.reactor.oldValues = make(map[*MCell]int)
}

func (c *MCell) SaveValue(){
	_, exist := c.reactor.oldValues[c]
	if !exist {
		// bookkeeping old values
		c.reactor.oldValues[c] = c.val
		for _, d := range c.deps {
				d.SaveValue()
		}
	}

}

func (c *MCell) UpdateValue(newVal int){
	c.val = newVal
	for _, d := range c.deps {
		nv := d.update(newVal)
		d.UpdateValue(nv)
	}
}

func (c *MCell) PerformCallbacks() {
	// check if this cell is updated
	oldV, exist := c.reactor.oldValues[c]

	// if updated and value is changed
	if exist && oldV != c.val {

		// call callbacks
		for _, callback := range c.callbacks {
			callback(c.val)
		}

		// clear from bookkeeper to avoid re-running callbacks
		// when certain ComputeCell uses the same InputCell multiple times
		delete(c.reactor.oldValues, c)

		// run callbacks of other dependants
		for _, dep := range c.deps {
			dep.PerformCallbacks()
		}
	}
}

type MCanceler struct {
	cell *MCell
	index int
}

func (c MCanceler) Cancel() {
	delete(c.cell.callbacks, c.index)
}

func (c *MCell) AddCallback(callback func(int)) Canceler {
	index := len(c.callbacks)
	c.callbacks[index] = callback
	return Canceler(MCanceler{cell: c, index: index})
}

func (c *MCell) AddDependant(cell MDepCell) {
	c.deps = append(c.deps, cell)
}


func (r *MReactor) NewCell(v int) *MCell {
	return &MCell{
		val: v,
		deps: []MDepCell{},
		callbacks:  make(map[int]func(int)),
		reactor: r,
	}
}

func (r *MReactor) CreateInput(v int) InputCell {
	return InputCell(r.NewCell(v))
}

func (r *MReactor) CreateCompute1(c Cell, compute func(int) int) ComputeCell {
	v := compute(c.Value())

	cell := r.NewCell(v)

	c.(*MCell).AddDependant(MDepCell{
		MCell: cell,
		update: func(newVal int) int {
			return compute(newVal)
		},
	})
	return ComputeCell(cell)
}

func (r *MReactor) CreateCompute2(c1 Cell, c2 Cell, compute func(int, int) int) ComputeCell {
	v := compute(c1.Value(), c2.Value())

	cell :=  r.NewCell(v)

	c1.(*MCell).AddDependant(MDepCell{
		MCell: cell,
		update: func(newVal int) int {
			return compute(newVal, c2.Value())
		},
	})

	c2.(*MCell).AddDependant(MDepCell{
		MCell: cell,
		update: func(newVal int) int {
			return compute(c1.Value(), newVal)
		},
	})

	return ComputeCell(cell)
}

func New() Reactor {
	return Reactor(&MReactor{ oldValues: make(map[*MCell]int)})
}
