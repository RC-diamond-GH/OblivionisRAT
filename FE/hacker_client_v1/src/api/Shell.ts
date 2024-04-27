import { POST } from "@/http";
// import { message } from "antd";
// const [messageApi] = message.useMessage();

let port = Uint8Array.from([]);
const shell = (id: number, cmd: string, param: string) => {
    let id_byte, cmd_byte, param_byte, type_byte;
    switch (cmd) {
        case "echo":
            cmd_byte = new Uint8Array([0x00, 0x02]);
            break;
        case "ls":
            cmd_byte = new Uint8Array([0x00, 0x03]);
            break;
        case "download":
            cmd_byte = new Uint8Array([0x00, 0x04]);
            break;
        case "port":
            localStorage.setItem("port", param);
            let tmp_num = parseInt(param);
            let num1 = tmp_num >> 8;
            let num2 = tmp_num & 0xff;
            port = new Uint8Array([num1, num2]);
            return;
        default:
            // messageApi.open({
            //     type: "success",
            //     content: "Successfully created a new beacon",
            //     duration: 1.5,
            // });
            console.log("cmd not found");
            return;
    }
    type_byte = new Uint8Array([0x00, 0x00, 0x00, 0x02]);
    id_byte = new Uint8Array([id]);
    param_byte = new TextEncoder().encode(param);

    const full_byte = new Uint8Array([
        ...port,
        ...type_byte,
        ...id_byte,
        ...cmd_byte,
        ...param_byte,
    ]);
    return POST("", full_byte);
};

export default shell;
