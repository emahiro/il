trait Animal {
    fn bark(&self);
    fn life_span(&self) -> String;
}

struct Dog {}

impl Animal for Dog {
    fn bark(&self) {
        println!("BowWow");
    }

    fn life_span(&self) -> String {
        return "13".to_string();
    }
}

struct Cat {}

impl Animal for Cat {
    fn bark(&self) {
        println!("NyauNayu");
    }

    fn life_span(&self) -> String {
        return "30".to_string();
    }
}

trait Car {
    fn new() -> Self;
    fn name(&self) -> String;
    fn made_from(&self) -> String;
}

struct Prius {
    kind: String,
    maker: String,
}

impl Car for Prius {
    fn new() -> Self {
        return Prius {
            kind: "Prius".to_string(),
            maker: "Toyota".to_string(),
        };
    }

    fn name(&self) -> String {
        return format!("{}", self.kind);
    }

    fn made_from(&self) -> String {
        return format!("{}", self.maker);
    }
}

fn main() {
    let dog = Dog {};
    dog.bark();

    let cat = Cat {};
    cat.bark();

    println!("{}", dog.life_span());
    println!("{}", cat.life_span());

    let prius = Prius::new();
    println!("{}", prius.kind);
    println!("made from {}", prius.made_from())
}
