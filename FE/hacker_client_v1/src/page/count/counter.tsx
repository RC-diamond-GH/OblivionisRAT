import typescriptLogo from "@/assets/typescript.svg";
import viteLogo from "@/assets/vite.svg";
import { useState } from "react";
import { invoke } from "@tauri-apps/api";
import { FC } from "react";

const Count: FC = () => {
    const [counter, setCounter] = useState(0);
    const [str, setStr] = useState("");

    const handleUpdate = () => {
        invoke("hello", { name: "world" }).then((res) => {
            console.log(res);
            setStr(res as string);
        });
        setCounter((count) => count + 1);
    };
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
            <p className="read-the-docs">
                Click on the Vite and TypeScript logos to learn more
            </p>
        </div>
    );
};
export default Count;
