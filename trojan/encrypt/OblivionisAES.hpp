#pragma once
#include <stdio.h>
#include <windows.h>

class OblivionisAES
{
public:
    OblivionisAES(PBYTE pbKey);
    VOID EncryptData(PBYTE pbData, DWORD* pdwLength);
    VOID DecryptData(PBYTE pbData, DWORD* pdwLength);

    unsigned char g_Key[176] = { 0 };
    unsigned char g_TempKey[4] = { 0 };
};