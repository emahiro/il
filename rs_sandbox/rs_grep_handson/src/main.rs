use std::{env::args, fs::read_to_string};


fn grep(content: String, pattern: String){

    for line in content.lines(){
        if line.contains(pattern.as_str()){
            println!("{}", line)
        }
    }
}

fn run (path:String, pattern: String){
    match read_to_string(path){
        Ok(content) =>grep(content,pattern),
        Err(err) => println!("{}",err)
    }
}

fn main() {
    let pattern = args().nth(1);
    let path = args().nth(2);
    match (pattern, path) {
        (Some(pattern),Some(path)) => run(path, pattern),
        _ => println!("pattern or path is not specified")
    };
}
