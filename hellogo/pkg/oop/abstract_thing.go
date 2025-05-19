package hellogo

type AbstractThing struct {
    name string
}

func NewAbstractThing(name string) *AbstractThing {
    if name == "" {
        panic("name cannot be empty")
    }
    return &AbstractThing{name: name}
}

func (t *AbstractThing) GetName() string {
    return t.name
}

// DoSomething must be implemented by concrete types
func (t *AbstractThing) DoSomething() {
    panic("DoSomething must be implemented")
}
