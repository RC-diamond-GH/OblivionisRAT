import { FC, useEffect, useState, useRef } from "react";

import { Input, Button, Modal, List, Typography, message } from "antd";
import shell from "@/api/Shell";
import newBeacon from "@/api/NewBeacon";
import beacons from "@/api/Beacons";
import ping from "@/api/ping";
import createListener from "@/api/Listener";

interface IData {
    id: number;
    msg: string;
    file?: { name: string; content: string };
}

const Home: FC = () => {
    const [messageApi, contextHolder] = message.useMessage();
    const [inputValue, setInputValue] = useState("");
    const [isModalOpen, setIsModalOpen] = useState(false);
    const [isModalOpenBeacon, setIsModalOpenBeacon] = useState(false);
    const [hosts, setHosts] = useState(0);
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
        if (res === true) {
            fetchBeacons();
        } else if (res === false) {
            messageApi.open({
                type: "error",
                content: "cmd not found",
                duration: 1.5,
            });
        }
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
    const handleOkListener = async () => {
        if (listener.name && listener.port) {
            console.log(listener, "<--listener");
            const res = await createListener(listener);
            console.log(res, "<--listener");
            setIsModalOpen(false);
            messageApi.open({
                type: "success",
                content: "Successfully created a new listener",
                duration: 1.5,
            });
        }
    };

    const [beacon, setBeacon] = useState({
        ip: "",
        port: "",
        sleep: "",
    });
    const handleBeacon = (e: any, type: string) => {
        setBeacon((pre) => ({
            ...pre,
            [type]: e.target.value,
        }));
    };
    const handleOkBeacon = async () => {
        if (beacon.ip && beacon.port && beacon.sleep) {
            const res = await newBeacon(beacon);
            console.log(res, "<--new beacon");
            setIsModalOpenBeacon(false);
            // flush the beacon number
            fetchBeacons();

            messageApi.open({
                type: "success",
                content: "Successfully created a new beacon",
                duration: 1.5,
            });
        }
    };
    const handleCancelBeacon = () => {
        setIsModalOpenBeacon(false);
    };

    const handleCancel = () => {
        setIsModalOpen(false);
    };

    /**
     * @description: create a new beacon
     */
    const showModalNewBeacon = async () => {
        setIsModalOpenBeacon(true);
    };

    const [listener, setListener] = useState({
        name: "",
        port: "",
    });
    const handleListener = (e: any, type: string) => {
        setListener((pre) => ({
            ...pre,
            [type]: e.target.value,
        }));
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
                        onClick={showModalNewBeacon}
                    >
                        + New Beacon
                    </Button>
                    <Button
                        style={{ backgroundColor: "#20242a", color: "#fff" }}
                        onClick={showModal}
                    >
                        + New Listener
                    </Button>
                </section>
                <Modal
                    title="Create a new Beacon"
                    open={isModalOpenBeacon}
                    onOk={handleOkBeacon}
                    onCancel={handleCancelBeacon}
                    okButtonProps={{ style: { backgroundColor: "#20242a" } }}
                >
                    <section
                        style={{
                            display: "flex",
                            flexDirection: "column",
                            gap: "1em",
                        }}
                    >
                        <Input
                            value={beacon.port}
                            placeholder="Beacon Port"
                            onChange={(e) => handleBeacon(e, "port")}
                        />
                        <Input
                            placeholder="Beacon IP"
                            value={beacon.ip}
                            onChange={(e) => handleBeacon(e, "ip")}
                        />
                        <Input
                            placeholder="Beacon Sleep Time"
                            value={beacon.sleep}
                            onChange={(e) => handleBeacon(e, "sleep")}
                        />
                    </section>
                </Modal>
                <Modal
                    title="Create a new Listener"
                    open={isModalOpen}
                    onOk={handleOkListener}
                    onCancel={handleCancel}
                    okButtonProps={{ style: { backgroundColor: "#20242a" } }}
                >
                    <section
                        style={{
                            display: "flex",
                            flexDirection: "column",
                            gap: "1em",
                        }}
                    >
                        <Input
                            value={listener.name}
                            placeholder="Listener Name"
                            onChange={(e) => handleListener(e, "name")}
                        />
                        <Input
                            placeholder="Listener Port"
                            value={listener.port}
                            onChange={(e) => handleListener(e, "port")}
                        />
                    </section>
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
                        <span
                            style={{
                                whiteSpace: "pre-line",
                            }}
                        >
                            {item.msg}
                        </span>
                    </List.Item>
                )}
            />
        </div>
    );
};
export default Home;
