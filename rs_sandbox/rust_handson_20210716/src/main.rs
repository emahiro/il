fn main() {
    let num = 10;
    let fizzbuss = if num == 15 {
        "FizzBuzz".to_string()
    } else if num == 5 {
        "Buzz".to_string()
    } else{
        "Fizz".to_string()
    };
    println!("{}", fizzbuss);
}
