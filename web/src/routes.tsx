import {RouteObject} from "react-router";
import Home from "./pages/Home";
import SignIn from "./pages/SignIn";
import MailBox from "./pages/MailBox";

export const routes: RouteObject[] = [

    {
        path: "/",
        element: <Home/>
    },
    {
        path: "/login",
        element: <SignIn/>
    },
    {
        path: "/mail-boxes",
        element: <MailBox/>
    }

]
