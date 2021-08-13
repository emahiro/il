use tokio::{
    io::{AsyncReadExt, AsyncWriteExt},
    net::TcpListener,
};

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let listener = TcpListener::bind("127.0.0.1:8080").await?;

    loop {
        let (mut socket, _) = listener.accept().await?;

        tokio::spawn(async move {
            let mut buf = [0; 1024];

            loop {
                let n = match socket.read(&mut buf).await {
                    Ok(n) if n == 0 => return,
                    Ok(n) => n,
                    Err(error) => {
                        eprintln!("failed to read from socket. error: {}", error);
                        return;
                    }
                };

                if let Err(error) = socket.write_all(&buf[0..n]).await {
                    eprintln!("failed to write to socket. error: {}", error);
                    return;
                };
            }
        });
    }
}
