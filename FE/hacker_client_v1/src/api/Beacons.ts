import { POST } from "@/http";

const beacons = () => {
    let tmp_num = +localStorage.getItem("port")!;
    let num1 = tmp_num >> 8;
    let num2 = tmp_num & 0xff;
    let port = new Uint8Array([num1, num2]);
    let full_bytes = new Uint8Array([...port, 0x00, 0x00, 0x00, 0x01]);
    return POST("", full_bytes);
};

export default beacons;
