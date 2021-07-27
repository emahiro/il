use std::{env::args, fs::read_to_string};

fn run_cat(path: String) {
    match read_to_string(path) {
        Ok(content) => print!("{}\n", content),
        Err (err) => print!("error: {}\n", err)
    };
}

fn main() {
    let x: Option<i32> = Some(10);
    match x {
        Some(n) => print!("found {}\n", n),
        None => print!("none")
    };
    match args().nth(1) {
        Some(path) => run_cat(path),
        None => print!("No path is specified\n")
    }
}
