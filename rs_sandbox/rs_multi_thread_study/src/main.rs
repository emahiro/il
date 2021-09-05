use std::{
    sync::{Arc, Mutex},
    thread::{self},
};

fn main() {
    // create 10 threads
    let mut handles = Vec::new();
    let data = Arc::new(Mutex::new(vec![1; 10]));

    for x in 0..10 {
        let data_ref = data.clone();
        let handle = thread::spawn(move || {
            let mut data = data_ref.lock().unwrap();
            data[x] += 1;
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
