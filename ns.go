package main

type (
	Namespace struct {
		name     Symbol
		mappings map[*string]*Var
	}
)

func NewNamespace(sym Symbol) *Namespace {
	return &Namespace{
		name:     sym,
		mappings: make(map[*string]*Var),
	}
}

func (ns *Namespace) Refer(sym Symbol, vr *Var) *Var {
	if sym.ns != nil {
		panic(RT.newError("Can't intern namespace-qualified symbol " + sym.ToString(false)))
	}
	ns.mappings[sym.name] = vr
	return vr
}

func (ns *Namespace) ReferAll(other *Namespace) {
	for name, vr := range other.mappings {
		ns.mappings[name] = vr
	}
}

// sym must be not qualified
func (ns *Namespace) intern(sym Symbol) *Var {
	if TYPES[*sym.name] != nil {
		panic(RT.newError("Can't intern type name " + *sym.name + " as a Var"))
	}
	sym.meta = nil
	v, ok := ns.mappings[sym.name]
	if !ok {
		v = &Var{
			ns:   ns,
			name: sym,
		}
		ns.mappings[sym.name] = v
	}
	return v
}
