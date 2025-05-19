package com.fathzer.oop;

public class ACoolThing extends AbstractThing implements CoolInterface {
    public ACoolThing(String name) {
        super(name);
    }

    @Override
    public void doSomething() {
        System.out.println(getName() + " does something");
    }

    @Override
    public void doSomethingCool() {
        System.out.println(getName() + " do something cool");
    }
}
