import { FC, useState } from "react";
import { useNavigate } from "react-router-dom";

import { Input } from "@arco-design/web-react";
import "@arco-design/web-react/dist/css/arco.css";

const Home: FC = () => {
    const navigate = useNavigate();
    const [inputValue, setInputValue] = useState("");
    const handleClick = () => {
        navigate("/demo");
    };
    const handleInput = (val: string) => {
        setInputValue(val);
    };
    const handleInputEnter = () => {
        console.log(inputValue);
        setInputValue("");
    };
    return (
        <div>
            <h1>Home</h1>
            <button onClick={handleClick}>goto</button>
            <Input
                value={inputValue}
                onChange={handleInput}
                onPressEnter={handleInputEnter}
            />
        </div>
    );
};
export default Home;
