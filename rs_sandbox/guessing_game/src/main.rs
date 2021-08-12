use std::io;

use rand::Rng;

fn main() {
    println!("Guess the number.");

    let secret_number = rand::thread_rng().gen_range(1..101);

    println!("The secret number is {}", secret_number);

    loop {
        println!("Please input your guess.");

        let mut guess = String::new();

        io::stdin()
            .read_line(&mut guess)
            .expect("failed to read line.");

        let guess: u32 = guess.trim().parse().expect("failed to parse guess number");

        println!("You guessed: {}", guess);

        match guess.cmp(&secret_number) {
            std::cmp::Ordering::Less => println!("Too small"),
            std::cmp::Ordering::Equal => println!("You win !!!"),
            std::cmp::Ordering::Greater => println!("Too Big"),
        };
    }
}
