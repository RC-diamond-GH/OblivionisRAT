import { POST } from "@/http";

const listener = async (data: { name: string; port: string }) => {
    let type_bytes = new Uint8Array([0x00, 0x00, 0x00, 0x04]); //4
    let name_bytes = new TextEncoder().encode(data.name); //4

    let tmp_num = parseInt(data.port);
    let num1 = tmp_num >> 8;
    let num2 = tmp_num & 0xff;
    let port_bytes = new Uint8Array([num1, num2]); //2

    let a = new Uint8Array([
        0x00, 0x00, 0x00, 0x05, 0x00, 0x00, 0x00, 0x05, 0x00, 0x00, 0x00, 0x05,
        0x00, 0x00, 0x00, 0x05,
    ]);
    a = new TextEncoder().encode("abcdefghijklmnop").slice(0, 16);
    let full_byte = new Uint8Array([
        ...type_bytes,
        ...name_bytes,
        ...port_bytes,
        ...a,
    ]);
    return POST("", full_byte);
};

export default listener;
