use std::{env::args, fs::read_to_string};

fn grep(content: String, pattern: String) {
    for line in content.lines() {
        if line.contains(pattern.as_str()) {
            println!("{}\n", line);
        }
    }
}

fn run(filepath: String, pattern: String) {
    match read_to_string(filepath) {
        Ok(content) => grep(content, pattern),
        Err(error) => println!("error: {}\n", error),
    };
}

fn main() {
    let pattern = args().nth(1);
    let filepath = args().nth(2);

    match (pattern, filepath) {
        (Some(pattern), Some(filepath)) => run(filepath, pattern),
        _ => println!("Pattern or path is not specified.\n"),
    };
}
