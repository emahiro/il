use rayon::prelude::*;
use std::fs::read_to_string;
use structopt::StructOpt;

#[derive(StructOpt)]
#[structopt(name = "rsgrep")]
struct GrepArgs {
    #[structopt(name = "PATTERN")]
    pattern: String,
    #[structopt(name = "FILEPATH")]
    files: Vec<String>,
}

impl GrepArgs {}

fn grep(content: String, state: &GrepArgs, file_name: &String) {
    let mut n = 0;
    for line in content.lines() {
        n = n + 1;
        if line.contains(state.pattern.as_str()) {
            println!("{}:L{} {}\n", file_name, n, line);
        }
    }
}

fn run(state: GrepArgs) {
    state
        .files
        .par_iter()
        .for_each(|file| match read_to_string(file) {
            Ok(content) => grep(content, &state, file),
            Err(err) => println!("error: {}\n", err),
        })
}

fn main() {
    run(GrepArgs::from_args());
}
