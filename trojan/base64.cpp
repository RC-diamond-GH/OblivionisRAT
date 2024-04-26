#include "base64.hpp"


char *base64_encode(const unsigned char *input, size_t length, int *outLen) {
    char *output = NULL;
    size_t output_length = 4 * ((length + 2) / 3); // 计算Base64编码后的长度
    size_t i, j;

    output = (char *)malloc(output_length + 1);
    if (!output) return NULL;

    for (i = 0, j = 0; i < length;) {
        uint32_t octet_a = i < length ? input[i++] : 0;
        uint32_t octet_b = i < length ? input[i++] : 0;
        uint32_t octet_c = i < length ? input[i++] : 0;

        uint32_t triple = (octet_a << 16) + (octet_b << 8) + octet_c;

        output[j++] = base64_chars[(triple >> 18) & 0x3F];
        output[j++] = base64_chars[(triple >> 12) & 0x3F];
        output[j++] = base64_chars[(triple >> 6) & 0x3F];
        output[j++] = base64_chars[triple & 0x3F];
    }

    for (i = 0; i < (3 - length % 3) % 3; i++) {
        output[output_length - 1 - i] = '=';
    }

    output[output_length] = '\0';
    *outLen = output_length;
    return output;
}

unsigned char *base64_decode(const char *input, size_t *output_length) {
    size_t input_length = strlen(input);
    if (input_length % 4 != 0) return NULL;

    size_t length = input_length / 4 * 3;
    if (input[input_length - 1] == '=') length--;
    if (input[input_length - 2] == '=') length--;

    unsigned char *output = (unsigned char *)malloc(length + 1);
    if (!output) return NULL;

    size_t i, j;
    for (i = 0, j = 0; i < input_length;) {
        uint32_t sextet_a = input[i] == '=' ? 0 & i++ : strchr(base64_chars, input[i++]) - base64_chars;
        uint32_t sextet_b = input[i] == '=' ? 0 & i++ : strchr(base64_chars, input[i++]) - base64_chars;
        uint32_t sextet_c = input[i] == '=' ? 0 & i++ : strchr(base64_chars, input[i++]) - base64_chars;
        uint32_t sextet_d = input[i] == '=' ? 0 & i++ : strchr(base64_chars, input[i++]) - base64_chars;

        uint32_t triple = (sextet_a << 18) + (sextet_b << 12) + (sextet_c << 6) + sextet_d;

        if (j < length) output[j++] = (triple >> 16) & 0xFF;
        if (j < length) output[j++] = (triple >> 8) & 0xFF;
        if (j < length) output[j++] = triple & 0xFF;
    }
    output[length] = '\0';
    *output_length = length;
    return output;
}