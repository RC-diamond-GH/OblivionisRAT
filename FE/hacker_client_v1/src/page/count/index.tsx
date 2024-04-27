import { invoke } from "@tauri-apps/api";
import { FC, useState } from "react";
// import { confirm, message } from "@tauri-apps/api/dialog";
import { GET } from "@/http";
import { Input } from "antd";

import typescriptLogo from "@/assets/typescript.svg";
import viteLogo from "@/assets/vite.svg";

const Count: FC = () => {
    const [counter, setCounter] = useState(0);
    const [str, setStr] = useState("");
    const [inputValue, setInputValue] = useState("");

    const handleUpdate = () => {
        invoke("hello", { name: "world" }).then((res) => {
            console.log(res);
            setStr(res as string);
        });
        setCounter((count) => count + 1);
    };
    const handleClickText = async () => {
        console.log("click");
        const res = await GET("/api/publisher/getPublisherInfo",{});
        console.log(res,'flag');
        // const confirmed = await confirm("Are you sure?", "Tauri");
        // message(`confirmed`, "" + confirmed);
    };

    // const handleSend1 = async () => {
    //     const res = await invoke("c1toc2");
    //     console.log(res);
    // };
    // const handleSend2 = async () => {
    //     const res = await invoke("c2toc1");
    //     console.log(res);
    // };
    // const handleSend3 = async () => {
    //     const res = await invoke("tcp");
    //     console.log(res);
    // };
    const handleInput = (e:any) => {
        setInputValue(e.target.value);
    };
    const handleInputEnter = () => {
        console.log(inputValue);
        setInputValue("");
    }
    return (
        <div>
            <a href="https://vitejs.dev" target="_blank">
                <img src={viteLogo} className="logo" alt="Vite logo" />
            </a>
            <a href="https://www.typescriptlang.org/" target="_blank">
                <img
                    src={typescriptLogo}
                    className="logo vanilla"
                    alt="TypeScript logo"
                />
            </a>
            <h1>Vite + TypeScript</h1>
            <div className="card">
                <button onClick={handleUpdate} type="button">
                    count is {counter}, str is {str}
                </button>
            </div>
            <p className="read-the-docs" onClick={handleClickText}>
                Click on the Vite and TypeScript logos to learn more
            </p>
            <Input value={inputValue} onChange={handleInput} onPressEnter={handleInputEnter} />
        </div>
    );
};
export default Count;
