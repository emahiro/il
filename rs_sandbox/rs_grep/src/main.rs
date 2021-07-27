use std::{env::args, fs::read_to_string};

struct GrepArgs {
    pattern: String,
    filepath: String,
}

impl GrepArgs {
    fn new(filepath: String, pattern: String) -> GrepArgs {
        GrepArgs {
            pattern: pattern,
            filepath: filepath,
        }
    }
}

fn grep(content: String, pattern: String) {
    for line in content.lines() {
        if line.contains(pattern.as_str()) {
            println!("{}\n", line);
        }
    }
}

fn run(state: GrepArgs) {
    match read_to_string(state.filepath) {
        Ok(content) => grep(content, state.pattern),
        Err(error) => println!("error: {}\n", error),
    };
}

fn main() {
    let pattern = args().nth(1);
    let filepath = args().nth(2);

    match (pattern, filepath) {
        (Some(pattern), Some(filepath)) => run(GrepArgs::new(filepath, pattern)),
        _ => println!("Pattern or path is not specified.\n"),
    };
}
