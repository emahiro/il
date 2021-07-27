use std::{env::args, fs::read_to_string};

fn run_cat(path: String){
    match read_to_string(path) {
        Ok(content) => println!("{}",content),
        Err(err) => println!("{}", err)
    }
}

fn main() {
    if let Some(path) = args().nth(1){
        run_cat(path)
    }
    // match args().nth(1) {
    //     Some(path) => run_cat(path),
    //     None => println!("{}", "no path is specified")
    // }
}
