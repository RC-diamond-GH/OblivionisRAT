import { FC, useEffect, useState } from "react";
import { Input } from "antd";
import { useNavigate } from "react-router-dom";

const Login: FC = () => {
    const [userAgent, setUserAgent] = useState("");
    const [host, setHost] = useState("");
    const [name, setName] = useState("");
    const [pass, setPass] = useState("");
    const navigate = useNavigate();

    const handleInputUserAgent = (e: any) => {
        setUserAgent(e.target.value.trim());
    };
    const handleInputHost = (e: any) => {
        setHost(e.target.value.trim());
    };
    const handleInputName = (e: any) => {
        setName(e.target.value.trim());
    };
    const handleInputPass = (e: any) => {
        setPass(e.target.value.trim());
    };
    const handleLogin = () => {
        if (name === "" || pass === "") {
            console.log("name or password is empty");
        }
        if (userAgent === "" || host === "") {
            console.log("user-agent or host is empty");
        }
        localStorage.setItem("user-agent", JSON.stringify(userAgent));
        localStorage.setItem("host", JSON.stringify(host));
        localStorage.setItem("user", JSON.stringify({ name, pass }));
        navigate("/home");
    };

    useEffect(() => {
        const user = localStorage.getItem("user");
        if (user) {
            const userAgent = JSON.parse(localStorage.getItem("user-agent")!);
            const host = JSON.parse(localStorage.getItem("host")!);
            setUserAgent(userAgent);
            setHost(host);
            const { name, pass } = JSON.parse(user);
            setName(name);
            setPass(pass);
        }
    }, []);
    return (
        <main
            style={{
                display: "flex",
                flexDirection: "column",
                justifyContent: "center",
                alignItems: "center",
                gap: "20px",
                border: "1px solid #0dbc79",
                borderRadius: "5px",
                color: "#0dbc79",
            }}
            className="card"
        >
            <p>Login</p>
            <Input
                style={{ backgroundColor: "#111317", color: "#0dbc79" }}
                prefix={"user-agent> "}
                value={userAgent}
                onChange={handleInputUserAgent}
                // onPressEnter={handleLogin}
            />
            <Input
                style={{ backgroundColor: "#111317", color: "#0dbc79" }}
                prefix={"host> "}
                value={host}
                onChange={handleInputHost}
                // onPressEnter={handleLogin}
            />
            <Input
                style={{ backgroundColor: "#111317", color: "#0dbc79" }}
                prefix={"name> "}
                value={name}
                onChange={handleInputName}
                // onPressEnter={handleLogin}
            />
            <Input
                style={{ backgroundColor: "#111317", color: "#0dbc79" }}
                prefix={"password> "}
                value={pass}
                onChange={handleInputPass}
                onPressEnter={handleLogin}
            />
        </main>
    );
};
export default Login;
