use std::fs::read_to_string;
use structopt::StructOpt;

#[derive(StructOpt)]
#[structopt(name = "rsgrep")]
struct GrepArgs {
    #[structopt(name = "PATTERN")]
    pattern: String,
    #[structopt(name = "FILEPATH")]
    filepath: Vec<String>,
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
    for filepath in state.filepath.iter() {
        match read_to_string(filepath) {
            Ok(content) => grep(content, &state),
            Err(err) => println!("error: {}\n", err),
        }
    }
}

fn main() {
    run(GrepArgs::from_args());
}
