use std::{env::args, fs::read_to_string};

struct GrepArgs {
    pattern: String,
    path: String,
}

impl GrepArgs {
    fn new(pattern: String, path: String) -> Self {
        GrepArgs { pattern, path }
    }
}

fn grep(content: String, pattern: String) {
    for line in content.lines() {
        if line.contains(pattern.as_str()) {
            println!("{}", line)
        }
    }
}

fn run(state: GrepArgs) {
    match read_to_string(state.path) {
        Ok(content) => grep(content, state.pattern),
        Err(err) => println!("{}", err),
    }
}

fn main() {
    let pattern = args().nth(1);
    let path = args().nth(2);
    match (pattern, path) {
        (Some(pattern), Some(path)) => run(GrepArgs::new(pattern, path)),
        _ => println!("pattern or path is not specified"),
    };
}
