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
    maker: String,
}

impl Car for Prius {
    fn new() -> Self {
        return Prius {
            maker: "Toyota".to_string(),
        };
    }

    fn name(&self) -> String {
        return format!("{}", "Prius");
    }

    fn made_from(&self) -> String {
        return format!("{}", self.maker);
    }
}

struct Note{
    maker:String,
}

impl Car for Note {
    fn new() -> Self {
        return Note{
            maker: "Nissan".to_string(),
        };
    }

    fn name(&self) -> String {
        return "Note".to_string();
    }

    fn made_from(&self) -> String {
        // self.maker だけだと self は借用してるだけで made_from が完了したら
        // 確保してたメモリは解放してしまう。そのため値を copy して再度戻す必要がある。
        // ちなみに clone じゃなく to_string でも可能（to_string の内部で copy が
        // 走ってる？）
        return self.maker.clone();
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
    println!("{}", prius.name());
    println!("made from {}", prius.made_from());

    let note = Note::new();
    println!("{}", note.name());
    println!("{}", note.made_from());
}
