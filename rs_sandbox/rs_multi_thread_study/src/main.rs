use std::{
    error,
    thread::{self},
};

fn main() {
    let handle = thread::spawn(|| {
        println!("hello world");
    });
    match dbg!(handle.join()) {
        Err(e) => println!("err: {:?}", e),
        _ => (),
    }
}
