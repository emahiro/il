fn main() {
    let num = 10;
    let fizzbuss = if num == 15 {
        "FizzBuzz"
    } else if num == 5 {
        "Buzz"
    } else{
        "Fizz"
    };
    println!("{}", fizzbuss);
}
