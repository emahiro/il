use std::{env::args, fs::read_to_string};

fn run_cat(path: String) {
    match read_to_string(path) {
        Ok(content) => print!("{}\n", content),
        Err (err) => print!("error: {}\n", err)
    };
}

fn main() {
    match args().nth(1) {
        Some(path) => run_cat(path),
        None => print!("No path is specified")
    }
}
