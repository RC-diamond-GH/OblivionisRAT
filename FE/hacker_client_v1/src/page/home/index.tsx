import { FC, useState } from "react";
import { useNavigate } from "react-router-dom";

import { Input, Card } from "antd";
import { Divider, List, Typography } from "antd";
import shell from "@/api/Shell";
import newBeacon from "@/api/NewBeacon";
import { fetch, Body } from "@tauri-apps/api/http";
import { http } from "@tauri-apps/api";

const Home: FC = () => {
    const navigate = useNavigate();
    const [inputValue, setInputValue] = useState("");
    const handleClick = () => {
        navigate("/demo");
    };
    const handleInput = (e: any) => {
        setInputValue(e.target.value);
    };
    const handleInputEnter = () => {
        console.log(inputValue);
        setInputValue("");
    };
    const handleClickTest = async () => {
        const res = await shell();
        console.log(
            res,
            "<--res",
            new TextDecoder().decode(new TextEncoder().encode("hello"))
        );
    };
    const handleClickTest1 = async () => {
        const res = await fetch("http://localhost:9090/api/book/getBook", {
            method: "GET",
            body: http.Body.bytes(new Uint8Array([0x00, 0x00, 0x00, 0x02])),
        });
        console.log(res, "<--res");
    };
    const handleClickTest2 = async () => {
        const res = await newBeacon();
        console.log(res, "<--res");
    };
    const data = [
        "Racing car sprays burning fuel into crowd.",
        "Japanese princess to wed commoner.",
        "Australian walks 100km after outback crash.",
        "Man charged over missing weddingharged over missing weddingharged over missing weddingharged over missing wedding girl.",
        "Los Angeles battles huge wildfires.",
    ];
    return (
        <div style={{ width: "100%" }}>
            {/* <h1 onClick={handleClickTest}>TestShell</h1>
            <h1 onClick={handleClickTest1}>TestBeacons</h1>
            <h1 onClick={handleClickTest2}>TestNewBeacon</h1>
            <button onClick={handleClick}>goto</button> */}
                        
            <List
                style={{
                    position: "fixed",
                    bottom: 0,
                    left: 0,
                    right: 0,
                    boxSizing: "border-box",
                }}
                header={<div>TERMINAL</div>}
                footer={
                    <Input
                        style={{ backgroundColor: "#111317", color: "#0dbc79" }}
                        prefix={">#"}
                        value={inputValue}
                        onChange={handleInput}
                        onPressEnter={handleInputEnter}
                    />
                }
                split={false}
                bordered
                dataSource={data}
                renderItem={(item) => (
                    <List.Item style={{ color: "#ffffff" }}>
                        <Typography.Text mark>[SHELL]</Typography.Text> {item}
                    </List.Item>
                )}
            />
        </div>
    );
};
export default Home;
