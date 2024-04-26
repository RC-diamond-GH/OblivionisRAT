// Prevents additional console window on Windows in release, DO NOT REMOVE!!
#![cfg_attr(not(debug_assertions), windows_subsystem = "windows")]

mod command;
// mod tcp;

use tokio::net::{TcpListener, TcpStream};
use tokio::io::{AsyncReadExt, AsyncWriteExt};

async fn handle_tcp_connections() -> Result<(), Box<dyn std::error::Error>> {
    let listener = TcpListener::bind("127.0.0.1:8080").await?;
    println!("Server listening on port 8080");

    loop {
        let (socket, _) = listener.accept().await?;
        tokio::spawn(async move {
            let _ = handle_tcp_connection(socket).await;
        });
    }
}

async fn handle_tcp_connection(mut socket: TcpStream) -> Result<(), Box<dyn std::error::Error>> {
    // 处理来自客户端的数据，发送响应
    let mut buf = [0; 1024];
    let bytes_read = socket.read(&mut buf).await?;
    println!("Received: {}", String::from_utf8_lossy(&buf[..bytes_read]));

    // 假设这里是对数据的处理过程

    // 发送响应给客户端
    socket.write_all(b"Hello from server").await?;

    Ok(())
}

#[tauri::command]
fn tcp() {
  let _ = handle_tcp_connections();
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {  

    tauri::Builder::default()
        .invoke_handler(tauri::generate_handler![
            command::hello,
            tcp
        ])
        .run(tauri::generate_context!())
        .expect("error while running tauri application");
    Ok(())
}
