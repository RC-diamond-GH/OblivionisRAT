import { POST } from "@/http";

const newBeacon = (data: { port: string; ip: string; sleep: string }) => {
    let type_bytes = new Uint8Array([0x00, 0x00, 0x00, 0x03]);
    const tmp_arr = data.ip.split(".").map((x) => parseInt(x));
    let ip_bytes = new Uint8Array(tmp_arr);

    let tmp_num = parseInt(data.port);
    let num1 = tmp_num >> 8;
    let num2 = tmp_num & 0xff;
    let port_bytes = new Uint8Array([num1, num2]);

    tmp_num = parseInt(data.sleep);
    num1 = (tmp_num >> 24) & 0xff;
    num2 = (tmp_num >> 16) & 0xff;
    let num3 = (tmp_num >> 8) & 0xff;
    let num4 = tmp_num & 0xff;
    let sleep_bytes = new Uint8Array([num1, num2, num3, num4]);

    let full_bytes = new Uint8Array([
        ...port_bytes,
        ...type_bytes,
        ...ip_bytes,
        ...port_bytes,
        ...sleep_bytes,
    ]);
    return POST("", full_bytes);
};

export default newBeacon;
