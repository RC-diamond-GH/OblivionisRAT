#include "KeyExchange.hpp"

#define B32MASK 0xFFFFFFFF

Uint128::Uint128(size_t low, size_t high) {
    this->LW = low;
    this->HG = high;
}
Uint128::Uint128(PBYTE data) {
    this->LW = *((size_t *)data);
    this->HG = *((size_t *)(data + 8));
}
PBYTE Uint128::toBytes() {
    PBYTE data = (PBYTE)malloc(16);
    *((size_t *)data) = LW;
    *((size_t *)(data + 8)) = HG;
    return data;
}
void Uint128::printHex() {
    printf("0x%llx %016llx", HG, LW);
}

Uint128 uint64Multi(size_t a, size_t b) {
    unsigned char tmpc[16];
    size_t *low = (size_t *)tmpc;
    size_t *high = (size_t *)(tmpc + 8);
    *low = 0;
    *high = 0;

    size_t ah = a >> 32;
    size_t al = a & B32MASK;
    size_t bh = b >> 32;
    size_t bl = b & B32MASK;
    *high += ah * bh;
    *low += al * bl;

    size_t tmp = al * bh;
    *(size_t *)(tmpc + 4) += tmp & B32MASK;
    *(size_t *)(tmpc + 8) += tmp >> 32;
    tmp = ah * bl;
    *(size_t *)(tmpc + 4) += tmp & B32MASK;
    *(size_t *)(tmpc + 8) += tmp >> 32;
    return Uint128(*low, *high);
}

void tmpcAdd(unsigned char *tmpc, Uint128 toAdd) {
    size_t aH = toAdd.HG, aL = toAdd.LW;
    *(size_t *)(tmpc)      += aL & B32MASK;
    *(size_t *)(tmpc + 4)  += aL >> 32;
    *(size_t *)(tmpc + 8)  += aH & B32MASK;
    *(size_t *)(tmpc + 12) += aH >> 32;
}
Uint128 Uint128::modMulti(Uint128 num) {
    unsigned char tmpc[32];
    size_t *low  = (size_t *)tmpc;         // 0 ~ 63
    size_t *mid  = (size_t *)(tmpc + 8);   // 64 ~ 127 
    size_t *high = (size_t *)(tmpc + 16);  // 128 ~ 191
    size_t *supre = (size_t *)(tmpc + 24); // 192 ~ 255

    *low = 0, *mid = 0, *high = 0, *supre = 0;
    size_t aH = HG, aL = LW;
    size_t bH = num.HG, bL = num.LW;

    Uint128 aHbH = uint64Multi(aH, bH);
    Uint128 aLbL = uint64Multi(aL, bL);
    Uint128 aLbH = uint64Multi(aL, bH);
    Uint128 aHbL = uint64Multi(aH, bL);
    tmpcAdd(tmpc, aHbH);
    tmpcAdd(tmpc, aHbH);
    tmpcAdd(tmpc, aLbL);
    tmpcAdd((unsigned char *)((size_t)tmpc + 8), aLbH);
    tmpcAdd((unsigned char *)((size_t)tmpc + 8), aHbL);

    size_t paddLow = *mid >> 63;
    *mid &= 0x7FFFFFFFFFFFFFFF;
    paddLow += *high << 1;
    size_t paddHigh = *high >> 63;
    paddHigh += *supre << 1;
    Uint128 padd(paddLow, paddHigh);
    
    tmpcAdd(tmpc, padd);
    
    return Uint128(*low, *mid);
}

int Uint128::popBit() {
    unsigned int toRet = LW & 1;
    LW >>= 1;
    unsigned long long bit = HG & 1;
    HG >>= 1;
    LW += bit << 63;
    return toRet;
}

Uint128 Uint128::modPow(Uint128 num) {
    if(num.LW == 0 && num.HG == 0) {
        return Uint128(1, 0);
    }
    if(num.LW == 1 && num.HG == 0) {
        return Uint128(LW, HG);
    }
    Uint128 result(1, 0);
    Uint128 base(LW, HG);
    while(num.LW > 0 || num.HG > 0) {
        if(num.popBit() == 1) {
            result = result.modMulti(base);
        }
        base = base.modMulti(base);
    }
    return result;
}
/*
int main() {
    Uint128 a(0x4f4d4b4947454341, 0x5f5d5b5957555351);

    Uint128 b1(0x000018be00006784, 0x2900004823); // C2 的私钥

    Uint128 b2(0x2250024d65c55f0e, 0xb99dd089071ee5d6);
    
    Uint128 a_b1 = a.modPow(b1);
    Uint128 a_b2 = a.modPow(b2);
    a_b1.printHex();
    printf("\n");
    a_b2.printHex();
    printf("\n");
    a_b1.modPow(b2).printHex();
    printf("\n");
    a_b2.modPow(b1).printHex();
}
*/