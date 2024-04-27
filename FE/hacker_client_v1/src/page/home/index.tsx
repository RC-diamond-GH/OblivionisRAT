import { FC, useEffect, useState, useRef } from "react";

import { Input, Button, Modal, List, Typography, message } from "antd";
import shell from "@/api/Shell";
import newBeacon from "@/api/NewBeacon";
import beacons from "@/api/Beacons";
import ping from "@/api/ping";

import { fetch, Body as _ } from "@tauri-apps/api/http";
import { http } from "@tauri-apps/api";

const Home: FC = () => {
    const [messageApi, contextHolder] = message.useMessage();
    const [inputValue, setInputValue] = useState("");
    const [isModalOpen, setIsModalOpen] = useState(false);
    const [hosts, setHosts] = useState(0);
    const [newBeaconNum, setNewBeaconNum] = useState(0);
    const [data, setData] = useState([
        { id: 0, msg: "Racing car sprays burning fuel into crowd." },
        { id: 1, msg: "Racing car sprays burning fuel into crowd." },
        { id: 2, msg: "Japanese princess to wed commoner." },
        { id: 3, msg: "Australian walks 100km after outback crash." },
        { id: 4, msg: "Australian walks 100km after outbac" },
    ]);
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
        let id = +cmd_arr[0];
        let cmd = cmd_arr[1];
        let param = cmd_arr[2];
        const res = await shell(id, cmd, param);
        console.log(res, "<--res");
        setInputValue("");
    };
    useEffect(() => {
        if (pingLock.current) {
            setInterval(async () => {
                console.log("ping");
                const res = (await ping()) as { c2: string; message: string };
                setData((prev) => [...prev, { id: +res.c2, msg: res.message }]);
            }, 1000);
        }
        return () => {
            pingLock.current = false;
        };
    }, []);

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