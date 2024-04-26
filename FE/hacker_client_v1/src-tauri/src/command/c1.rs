use std::net::UdpSocket;

#[tauri::command]
pub fn c1toc2() {
    let socket = UdpSocket::bind("127.0.0.1:8080").expect("bind failed");

    loop {
        let mut buf = [0u8; 1024];
        let (amt, src) = socket.recv_from(&mut buf).expect("recv_from failed");
        let buf = &mut buf[..amt];
        buf.reverse();
        socket.send_to(buf, &src).expect("send_to failed");
    }
}
