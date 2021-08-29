use std::{rc::Rc, thread::{self}};

fn main() {
    // create 10 threads
    let mut handles = Vec::new();
    let mut data = Rc::new(vec![1; 10]);
    println!("{:?}", data);

    for x in 0..10 {
        let data_ref = data.clone();
        let handle = thread::spawn(move || {
            data_ref[x] += 1;
        });
        handles.push(handle)
    }

    for handle in handles {
        match dbg!(handle.join()) {
            Err(e) => println!("err: {:?}", e),
            _ => (),
        }
    }

    match dbg!(data) {
        _ => (),
    };
}
