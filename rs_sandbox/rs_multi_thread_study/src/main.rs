use std::{
    error,
    thread::{self},
};

fn main() {
    // create 10 threads
    let mut handles = Vec::new();

    for x in 0..10{
        let handle = thread::spawn(|| {
            println!("hello world! - {}", x);
        });
        handles.push(handle)
    }

    for handle in handles {
        match dbg!(handle.join()) {
            Err(e) => println!("err: {:?}", e),
            _ => (),
        }
    }
}
