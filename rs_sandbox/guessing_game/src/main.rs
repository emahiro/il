use std::io;

fn main() {
    println!("Guess the number.");
    println!("Please input your guess.");

    let mut guess = String::new();

    match io::stdin().read_line(&mut guess) {
        Ok(input) => {
            println!("input is {}", input);
        }
        Err(error) => println!("failed to read line. error {}", error),
    }

    println!("You guessed: {}", guess)
}
