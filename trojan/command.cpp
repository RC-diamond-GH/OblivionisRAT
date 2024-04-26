
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