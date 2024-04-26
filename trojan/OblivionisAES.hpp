#ifndef AES
#define AES
#include <stdio.h>

typedef unsigned char *PBYTE;
typedef unsigned long DWORD;

class OblivionisAES
{
public:
    OblivionisAES(PBYTE pbKey);
    void EncryptData(PBYTE pbData, int* pdwLength);
    void DecryptData(PBYTE pbData, int* pdwLength);

    unsigned char g_Key[176] = { 0 };
    unsigned char g_TempKey[4] = { 0 };
};

#endif