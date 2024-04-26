use std::net::UdpSocket;

#[tauri::command]
pub fn c2toc1() {
    let socket = UdpSocket::bind("127.0.0.1:9091").expect("bind failed");

    socket.send_to(b"hello", "127.0.0.1:8080").expect("send_to failed");

    let mut buf = [0u8; 1024];
    let amt = socket.recv(&mut buf).expect("recv failed");
    let buf = &mut buf[..amt];
    println!("{}", std::str::from_utf8(buf).unwrap());
}
