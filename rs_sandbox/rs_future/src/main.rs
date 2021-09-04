use std::{pin::Pin, process::Output, task::{Context, Poll}};

use futures::{executor, future::join_all, Future};

struct CountDown(u32);

impl Future for CountDown {
    type Output = String;

    fn poll(mut self: Pin<&mut Self>, cx: &mut Context) -> Poll<String> {
        if self.0 == 0 {
            Poll::Ready("ZERO!!!".to_string())
        } else {
            // println!("{}", self.0);
            self.0 -= 1;
            cx.waker().wake_by_ref();
            Poll::Pending
        }
    }
}

async fn async_add(left: i32, right: i32) -> i32 {
    return left + right;
}

async fn something_great_async_function() -> i32 {
    let ans = async_add(1,1).await;
    println!("answer = {}\n", ans);
    return ans
}

fn move_to_async_block() -> impl Future<Output = ()> {
    let outside_variable = "this is outside".to_string();
    async move {
        println!("this is {}", outside_variable);
    }
}

fn main() {
    // let countdown_future1 = CountDown(10);
    // let countdown_future2 = CountDown(20);
    // let cd_set = join_all(vec![countdown_future1, countdown_future2]);
    // let _ = executor::block_on(cd_set);
    // for (i, s) in res.iter().enumerate(){
    //     println!("{}: {}", i, s);
    // }

    let _ = executor::block_on(something_great_async_function());
    let _ = executor::block_on(move_to_async_block());

}
