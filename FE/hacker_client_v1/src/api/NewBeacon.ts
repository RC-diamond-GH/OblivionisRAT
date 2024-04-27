import { POST } from "@/http";

const newBeacon = ()=>{
    return POST("", new Uint8Array([0x00,0x00,0x00,0x03]));
}

export default newBeacon;