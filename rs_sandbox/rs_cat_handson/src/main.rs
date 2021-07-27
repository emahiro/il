use std::fs::read_to_string;

fn run_cat(path: String) {
    match read_to_string(path) {
        Ok(content) => print!("{}\n", content),
        Err (err) => print!("error: {}\n", err)
    };
}

fn main() {
    let path = "./src/main.rs".to_string();
    run_cat(path);
}
