import { FC, useState } from "react";
import { useNavigate } from "react-router-dom";

import { Input, Button, Modal, List, Typography, message } from "antd";
import shell from "@/api/Shell";
import newBeacon from "@/api/NewBeacon";

import { fetch, Body } from "@tauri-apps/api/http";
import { http } from "@tauri-apps/api";

const Home: FC = () => {
    const navigate = useNavigate();
    const [messageApi, contextHolder] = message.useMessage();
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
        "Los Angeles battles huge wildfires.",
        "Los Angeles battles huge wildfires.",
        "Los Angeles battles huge wildfires.",
        "Los Angeles battles huge wildfires.",
        "Los Angeles battles huge wildfires.1111",
    ];
    const [isModalOpen, setIsModalOpen] = useState(false);
    const [hosts, setHosts] = useState(0);
    const [newBeaconNum, setNewBeaconNum] = useState(0);

    const showModal = () => {
        setIsModalOpen(true);
    };

    const handleOk = () => {
        setIsModalOpen(false);
    };

    const handleCancel = () => {
        setIsModalOpen(false);
    };

    const handleConnectNum = (e: any) => {
        setNewBeaconNum(e.target.value);
    };
    const handleNewBeacon = async () => {
        const res = await newBeacon();
        console.log(res, "<--res");
        messageApi.open({
            type: "success",
            content: "Successfully created a new beacon",
            duration: 1.5,
        });
    };
    return (
        <div style={{ width: "100%" }}>
            <main
                style={{
                    position: "absolute",
                    height: "50vh",
                    top: 0,
                    left: 0,
                    right: 0,
                    display: "flex",
                    flexDirection: "column",
                    alignItems: "start",
                    padding: "0 1em",
                }}
            >
                <p>Current available hosts : {hosts} </p>
                <section style={{ display: "flex", gap: "1em" }}>
                    <Button
                        style={{ backgroundColor: "#20242a", color: "#fff" }}
                        onClick={showModal}
                    >
                        Connect to Beacon
                    </Button>
                    <Button
                        style={{ backgroundColor: "#20242a", color: "#fff" }}
                        onClick={handleNewBeacon}
                    >
                        + New Beacon
                    </Button>
                </section>
                <Modal
                    title="Create a new Beacon"
                    open={isModalOpen}
                    onOk={handleOk}
                    onCancel={handleCancel}
                    okButtonProps={{ style: { backgroundColor: "#20242a" } }}
                >
                    <p>Please input a beacon number : </p>
                    <Input value={newBeaconNum} onChange={handleConnectNum} />
                </Modal>
                {contextHolder}
            </main>

            <List
                style={{
                    position: "fixed",
                    bottom: 0,
                    left: 0,
                    right: 0,
                    boxSizing: "border-box",
                    color: "#ffffff",
                    maxHeight: "50vh",
                    overflow: "auto",
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
                        <Typography.Text type="success">
                            [SHELL]
                        </Typography.Text>{" "}
                        {item}
                    </List.Item>
                )}
            />
        </div>
    );
};
export default Home;
