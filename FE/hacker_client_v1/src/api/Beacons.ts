import { POST } from "@/http";

const beacons = ()=>{
    return POST("/", new Uint8Array([0x00,0x00,0x00,0x01]));
}

export default beacons;