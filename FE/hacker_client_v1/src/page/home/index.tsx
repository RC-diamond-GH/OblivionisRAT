import { FC } from "react";
import { useNavigate } from "react-router-dom";

const Home: FC = () => {
    const navigate = useNavigate();
    const handleClick = () => {
        navigate("/demo");
    };
    return (
        <div>
            <h1>Home</h1>
            <button onClick={handleClick}>goto</button>
        </div>
    );
};
export default Home;
