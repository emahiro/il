use std::{fs::read_to_string};
use structopt::StructOpt;

#[derive(StructOpt)]
#[structopt(name = "rsgrep")]
struct GrepArgs {
    #[structopt(name = "PATTERN")]
    pattern: String,
    #[structopt(name = "FILEPATH")]
    filepath: String,
}

impl GrepArgs {}

fn grep(content: String, state: &GrepArgs) {
    for line in content.lines() {
        if line.contains(state.pattern.as_str()) {
            println!("{}\n", line);
        }
    }
}

fn run(state: GrepArgs) {
    match read_to_string(&state.filepath) {
        Ok(content) => grep(content, &state),
        Err(error) => println!("error: {}\n", error),
    };
}

fn main() {
    run(GrepArgs::from_args());
}
