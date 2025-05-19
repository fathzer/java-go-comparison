package hellogo

import "fmt"

type ACoolThing struct {
    *AbstractThing
}

func NewACoolThing(name string) *ACoolThing {
    return &ACoolThing{
        AbstractThing: NewAbstractThing(name),
    }
}

func (t *ACoolThing) GetName() string {
    return t.AbstractThing.GetName()
}

func (t *ACoolThing) DoSomething() {
    fmt.Println(t.GetName() + " does something")
}

func (t *ACoolThing) DoSomethingCool() {
    fmt.Println(t.GetName() + " does something cool")
}
