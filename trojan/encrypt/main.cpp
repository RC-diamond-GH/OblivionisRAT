#include<cstdio>
#include<cstdlib>
#include "OblivionisAES.hpp"

int main() {
    PBYTE key = (PBYTE)malloc(16);
    memmove(key, "1234567887654321", 16);
    const char *message = "public class HelloWorld{public static void main(String[] args){System.out.println()}}";
    DWORD mesLen = strlen(message);
    PBYTE mes = (PBYTE)malloc(mesLen);
    memmove(mes, message, mesLen);

    OblivionisAES aes = OblivionisAES(key);
    aes.EncryptData(mes, &mesLen);
    aes.DecryptData(mes, mesLen);

    printf("%s\n", mes);
}