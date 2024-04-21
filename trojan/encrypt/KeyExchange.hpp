#include<windows.h>
#include<stdio.h>

class Uint128 {
    public:
    size_t LW; // 低 8 个字节
    size_t HG; // 高 8 个字节
    public:
    Uint128(size_t low, size_t high);
    Uint128(PBYTE data);

    PBYTE toBytes();

    // 实现两个 Uint128 的模 2^127 - 1 乘法
    Uint128 modMulti(Uint128 num);

    // 快速模 2^127 - 1 幂
    Uint128 modPow(Uint128 amount);

    int popBit();
    void printHex();
};