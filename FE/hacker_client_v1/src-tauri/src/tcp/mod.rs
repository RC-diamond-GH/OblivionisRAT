use tokio::net::{TcpListener, TcpStream};
use tokio::io::{AsyncReadExt, AsyncWriteExt};

pub struct TcpClient {
    socket: TcpStream,
}

impl TcpClient {
    pub async fn connect(addr: &str) -> Result<Self, Box<dyn std::error::Error>> {
        let socket = TcpStream::connect(addr).await?;
        Ok(Self { socket })
    }

    pub async fn send(&mut self, data: &[u8]) -> Result<(), Box<dyn std::error::Error>> {
        self.socket.write_all(data).await?;
        Ok(())
    }

    pub async fn receive(&mut self, buf: &mut [u8]) -> Result<usize, Box<dyn std::error::Error>> {
        let bytes_read = self.socket.read(buf).await?;
        Ok(bytes_read)
    }

    pub async fn close(&mut self) -> Result<(), Box<dyn std::error::Error>> {
        self.socket.shutdown().await?;
        Ok(())
    }
}
