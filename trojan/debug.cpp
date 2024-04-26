#include "debug.hpp"

void hexDump(PBYTE data, int len) {
    int i = 0;
    while (i < len) {
        if (i % 16 == 0 && i != 0) {
            printf("        ");
            for (int j = 0; j < 16; j++) {
                if (i - 16 + j < len)
                    printf("%c", (data[i - 16 + j] >= 32 && data[i - 16 + j] <= 126) ? data[i - 16 + j] : '.');
                else
                    printf(" ");
            }
            printf("\n");
        }
        printf("%02x ", data[i]);
        i++;
    }
}