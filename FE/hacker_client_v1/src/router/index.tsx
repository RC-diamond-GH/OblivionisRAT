import { useRoutes, BrowserRouter } from "react-router-dom";
import { lazy } from "react";
import LazyWrap from "@/component/lazyWrap";
import NotFound from "@/page/notfound";

/**
 * @fileoverview Router configuration
 */
const router = [
    {
        path: "/",
        element: <LazyWrap Component={lazy(() => import("@/page/home"))} />,
    },
    {
        path: "/demo",
        element: (
            <LazyWrap Component={lazy(() => import("@/page/count/counter"))} />
        ),
    },
    {
        path: "*",
        element: <NotFound />,
    },
];

/**
 * Generate router
 * @todo useRoutes returns the elements that match the current location.
 * @description returned elements are rendered by the caller.
 */
const GenRouter = (): React.ReactNode => {
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
