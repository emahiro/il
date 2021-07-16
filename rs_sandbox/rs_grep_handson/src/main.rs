use std::fs::read_to_string;
use structopt::StructOpt;

#[derive(StructOpt)] // attribute cf. annotation
#[structopt(name = "rsgrep")]
struct GrepArgs {
    #[structopt(name="PATTERN")]
    pattern: String,
    #[structopt(name="NAME")]
    path: String,
}

// impl GrepArgs {
//     fn new(pattern: String, path: String) -> Self {
//         GrepArgs { pattern, path }
//     }
// }

fn grep(content: String, state: &GrepArgs) {
    for line in content.lines() {
        if line.contains(state.pattern.as_str()) {
            println!("{}", line)
        }
    }
}

fn run(state: GrepArgs) {
    match read_to_string(&state.path) {
        Ok(content) => grep(content, &state),
        Err(err) => println!("{}", err),
    }
}

fn main() {
    run(GrepArgs::from_args())
    // let pattern = args().nth(1);
    // let path = args().nth(2);
    // match (pattern, path) {
    //     (Some(pattern), Some(path)) => run(GrepArgs::new(pattern, path)),
    //     _ => println!("pattern or path is not specified"),
    // };
}
