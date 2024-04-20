#pragma once
#include <stdio.h>
#include <windows.h>

class OblivionisAES
{
public:
    OblivionisAES(PBYTE pbKey);
    // 加密函数
    VOID EncryptData(PBYTE pbData, DWORD* pdwLength);
    // 解密函数
    VOID DecryptData(PBYTE pbData, DWORD pdwLength);
    // 加密密钥
    unsigned char g_Key[176] = { 0 };
    // 用于临时存储替换密钥
    unsigned char g_TempKey[4] = { 0 };
    // box替换下标
    int g_replaceIndex = 1;

    void InitXorKeyBox();
    void replaceFourByteKey(unsigned char* fourByteKey);
    void ByteOutOfOrder(unsigned char* data);
    void boxXorData(unsigned char* data);
    // Decrypt
    void DeByteOutOfOrder(unsigned char* data);
    void DereplaceBoxData(unsigned char* data);
    void DeboxXorData(unsigned char* data);
    // ror
    template<class T> T __ROL__(T value, int count)
    {
        const unsigned int nbits = sizeof(T) * 8;

        if (count > 0)
        {
            count %= nbits;
            T high = value >> (nbits - count);
            if (T(-1) < 0) // signed value
                high &= ~((T(-1) << count));
            value <<= count;
            value |= high;
        }
        else
        {
            count = -count % nbits;
            T low = value << (nbits - count);
            value >>= count;
            value |= low;
        }
        return value;
    }

    inline unsigned int __ROR4__(unsigned int value, int count) {
        return __ROL__((unsigned int)value, -count);
    }
};