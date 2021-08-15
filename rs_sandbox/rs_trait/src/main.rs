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

fn main() {
    let dog = Dog {};
    dog.bark();

    let cat = Cat {};
    cat.bark();

    println!("{}", dog.life_span());
    println!("{}", cat.life_span())
}
