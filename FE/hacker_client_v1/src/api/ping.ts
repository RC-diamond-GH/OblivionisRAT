import { POST } from "@/http";

const ping =async ()=>{
    return POST("", new Uint8Array([]));
}

export default ping;