use std::env::{self};

fn main() {
    let args: Vec<String> = env::args().collect();
    let numstr = &args[1].to_string();
    let num: i32 = numstr.parse().unwrap();
    let fizzbuss = if num % 15 == 0 {
        "FizzBuzz".to_string()
    } else if num % 5 == 0 {
        "Buzz".to_string()
    } else {
        "Fizz".to_string()
    };
    println!("{}", fizzbuss);
}
