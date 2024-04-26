#include "debug.hpp"

void hexDump(PBYTE data, int len) {
    for(int i = 0; i < len; i++) {
        if(i != 0 && i % 16 == 0) {
            printf("\n");
        }
        printf("%02x ", data[i]);
    }
    printf("\n");
}
int strlen(char *str) {
    int i;
    while(str[i] != 0) {
        i++;
    }
    return i;
}