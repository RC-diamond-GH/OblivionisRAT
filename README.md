# OblivionisRAT
架构如下：

```mermaid
graph LR
A["C2服务端"]
B["前端页面(黑客客户端)"]
C["木马"]

A--->TX1("(黑客专用频道)通信")
TX1--->B
A--->TX2("(不同Listener)通信")
TX2--->C
```

