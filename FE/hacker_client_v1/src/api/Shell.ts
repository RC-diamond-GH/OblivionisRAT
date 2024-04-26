import { POST } from "@/http";

const shell = ()=>{
    return POST("/", new Uint8Array([0x00,0x00,0x00,0x02]));
}

export default shell;