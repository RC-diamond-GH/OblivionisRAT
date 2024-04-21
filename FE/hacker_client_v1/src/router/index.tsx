import { useRoutes, BrowserRouter } from "react-router-dom";

import Count from "@/page/count/counter";
import NotFound from "@/page/notfound";

/**
 * @fileoverview Router configuration
 */
const router = [
    {
        path: "/demo",
        element: <Count />,
    },
    {
        path: "*",
        element: <NotFound />,
    },
];
const GenRouter = () => {  
    const routing = useRoutes(router);
    return routing;
};

const App = () => {
    return (
        <BrowserRouter>
            <GenRouter />
        </BrowserRouter>
    );
};

export default App;
