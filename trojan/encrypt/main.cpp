#include<cstdio>
#include<cstdlib>
#include "OblivionisAES.hpp"

void hexDump(PBYTE data, int len) {
    for(int i = 0; i < len; i++) {
        if(i != 0 && i % 16 == 0) {
            printf("\n");
        }
        printf("%02x ", data[i]);
    }
    printf("\n");
}
void AESTest() {
    PBYTE key = (PBYTE)malloc(16);
    memmove(key, "1234567887654321", 16);
    const char *message = "abcdefghabcdefghabcdefghabcdefgh1";
    DWORD mesLen = strlen(message);
    PBYTE mes = (PBYTE)malloc(mesLen + 16);
    memmove(mes, message, mesLen);

    OblivionisAES aes = OblivionisAES(key);
    printf("key = \n");
    hexDump(aes.g_Key, 176);

    printf("to Encrypt = \n");
    hexDump(mes, mesLen);
    aes.EncryptData(mes, &mesLen);
    printf("encrypt = \n");
    hexDump(mes, mesLen);

    aes.DecryptData(mes, &mesLen);
    printf("decrypt = \n");
    hexDump(mes, mesLen);

    printf("\n%s\n", mes);
}
int main() {
    
}