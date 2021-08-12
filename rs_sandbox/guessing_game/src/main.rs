use std::{io, process::exit};

use rand::Rng;

fn main() {
    println!("Guess the number.");

    let secret_number = rand::thread_rng().gen_range(1..101);

    loop {
        println!("Please input your guess.");

        let mut guess = String::new();

        io::stdin()
            .read_line(&mut guess)
            .expect("failed to read line.");

        let guess: u32 = match guess.trim().parse(){
            Ok(num) => num,
            Err(_) => continue, // ignore invalid number
        };

        println!("You guessed: {}", guess);

        match guess.cmp(&secret_number) {
            std::cmp::Ordering::Less => println!("Too small"),
            std::cmp::Ordering::Greater => println!("Too Big"),
            std::cmp::Ordering::Equal => {
                println!("You win !!!");
                break;
            },
        };
    }
    exit(0)
}
