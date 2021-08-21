use std::thread::{self};

fn main() {
    thread::spawn(|| {
        println!("hello world");
    });
}
