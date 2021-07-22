fn main() {
    for num in 1..100 {
        let fizzbuss = if num % 15 == 0 {
            "FizzBuzz".to_string()
        } else if num % 5 == 0 {
            "Buzz".to_string()
        } else {
            "Fizz".to_string()
        };
        println!("{}", fizzbuss);
    }
}
