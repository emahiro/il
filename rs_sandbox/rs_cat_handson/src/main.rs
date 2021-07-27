use std::{env::args, fs::read_to_string};

fn run_cat(path: String) {
    match read_to_string(path) {
        Ok(content) => print!("{}\n", content),
        Err (err) => print!("error: {}\n", err)
    };
}

fn main() {
    if let Some(path) = args().nth(1) {
        run_cat(path)
    }
}
