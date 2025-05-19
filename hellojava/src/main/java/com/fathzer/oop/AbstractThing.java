package com.fathzer.oop;

public abstract class AbstractThing {
    private String name;

    protected AbstractThing(String name) {
        if (name==null) {
            throw new IllegalArgumentException("name");
        }
        this.name = name;
    }
    
    public abstract void doSomething();

    public String getName() {
        return name;
    }
}

