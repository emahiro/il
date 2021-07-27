use std::fs::read_to_string;

fn main() {
    let path = "./src/main.rs".to_string();
    match read_to_string(path) {
        Ok(content) => print!("{}", content),
        Err(reason) => print!("error: {}", reason),
    }
}
