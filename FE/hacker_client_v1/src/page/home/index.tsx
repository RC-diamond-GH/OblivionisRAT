import { FC, useEffect, useState, useRef } from "react";

import { Input, Button, Modal, List, Typography, message } from "antd";
import shell from "@/api/Shell";
import newBeacon from "@/api/NewBeacon";
import beacons from "@/api/Beacons";
import ping from "@/api/ping";

import { fetch, Body as _ } from "@tauri-apps/api/http";
import { http } from "@tauri-apps/api";

interface IData {
    id: number;
    msg: string;
    file?: { name: string; content: string };
}

const Home: FC = () => {
    const [messageApi, contextHolder] = message.useMessage();
    const [inputValue, setInputValue] = useState("");
    const [isModalOpen, setIsModalOpen] = useState(false);
    const [hosts, setHosts] = useState(0);
    const [newBeaconNum, setNewBeaconNum] = useState(0);
    const [data, setData] = useState<IData[]>([]);
    const [beaconList, setBeaconList] = useState<number[]>([0, 1, 2]);
    const [targetId, setTargetId] = useState(0);

    const pingLock = useRef(true);
    const isFirstRender = useRef(true);
    useEffect(() => {
        const fetchData = async () => {
            const res = (await beacons()) as { c2: number };
            setHosts(res.c2);
            const arr = [];
            for (let i = 0; i < res.c2; i++) {
                arr.push(i);
            }
            setBeaconList(arr);
        };
        if (isFirstRender.current) {
            fetchData();
        }
        return () => {
            isFirstRender.current = false;
        };
    }, []);
    const fetchBeacons = async () => {
        console.log("flush");
        const res = (await beacons()) as { c2: number };
        setHosts(res.c2);
        const arr = [];
        for (let i = 0; i < res.c2; i++) {
            arr.push(i);
        }
        setBeaconList(arr);
    };

    const handleInput = (e: any) => {
        setInputValue(e.target.value);
    };
    /**
     * @description: send the input value to the shell
     */
    const handleInputEnter = async () => {
        console.log(inputValue);
        const cmd_arr = inputValue.split(" ");
        let cmd = cmd_arr[0];
        let param = cmd_arr[1];
        const res = await shell(targetId, cmd, param);
        console.log(res, "<--res");
        setInputValue("");
    };
    useEffect(() => {
        if (pingLock.current) {
            setInterval(async () => {
                console.log("ping");
                const res = (await ping()) as {
                    c2: string;
                    message: string;
                    file?: { name: string; content: string };
                };

                new Promise<{ name: string; content: string }>((resolve) => {
                    if (!!res.file) {
                        resolve(res.file);
                        res.message = `Downloading ${res.file.name}, please wait a moment...`;
                    }
                }).then((file_obj) => {
                    downloadFileFromBase64(file_obj);
                });

                setData((prev) => [...prev, { id: +res.c2, msg: res.message }]);
            }, 1000);
        }
        return () => {
            pingLock.current = false;
        };
    }, []);
    function downloadFileFromBase64(file: { name: string; content: string }) {
        // console.log(file, "<--file");
        // 将 Base64 编码的字符串解码为二进制数据
        const base64String = file.content;
        const fileName = file.name;
        const binaryString = atob(base64String);
        const length = binaryString.length;
        const bytes = new Uint8Array(length);
        for (let i = 0; i < length; i++) {
            bytes[i] = binaryString.charCodeAt(i);
        }

        // 创建 Blob 对象
        const blob = new Blob([bytes], { type: "application/octet-stream" });

        // 创建 Blob URL
        const blobUrl = URL.createObjectURL(blob);

        // 创建链接元素并设置属性
        const link = document.createElement("a");
        link.href = blobUrl;
        link.download = fileName;

        // 模拟点击链接以触发下载
        link.click();

        // 释放 Blob URL
        URL.revokeObjectURL(blobUrl);
    }

    const showModal = () => {
        setIsModalOpen(true);
    };

    /**
     * @description: connect to the beacon
     */
    const handleOk = async () => {
        const res = await fetch("http://localhost:9090/api/book/getBook", {
            method: "GET",
            body: http.Body.bytes(new Uint8Array([0x00, 0x00, 0x00, 0x02])),
        });
        if (res) {
            setHosts(hosts + 1);
            setIsModalOpen(false);
        }
    };

    const handleCancel = () => {
        setIsModalOpen(false);
    };

    const handleConnectNum = (e: any) => {
        setNewBeaconNum(e.target.value);
    };

    /**
     * @description: create a new beacon
     */
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
                    {/* <Button
                        style={{ backgroundColor: "#20242a", color: "#fff" }}
                        onClick={showModal}
                    >
                        Connect to Beacon
                    </Button> */}
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
                    maxHeight: "65vh",
                    overflow: "auto",
                }}
                header={
                    <div style={{ display: "flex", gap: "1em" }}>
                        TERMINAL
                        {beaconList.map((item) => {
                            return (
                                <Button
                                    style={{
                                        backgroundColor: "#20242a",
                                        color: "#fff",
                                    }}
                                    onClick={() => {
                                        fetchBeacons();
                                        setTargetId(item);
                                    }}
                                    key={item}
                                >
                                    {item}
                                </Button>
                            );
                        })}
                    </div>
                }
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
                dataSource={data.filter((item) => item.id === targetId)}
                renderItem={(item) => (
                    <List.Item style={{ color: "#ffffff" }}>
                        <Typography.Text type="success">
                            [SHELL ${item.id}]
                        </Typography.Text>{" "}
                        {item.msg}
                    </List.Item>
                )}
            />
        </div>
    );
};
export default Home;
