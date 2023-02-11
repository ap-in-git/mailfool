import {RouteObject} from "react-router";
import Home from "./pages/Home";
import SignIn from "./pages/SignIn";
import MailBox from "./pages/MailBox";
import AppLayout from "./layouts/AppLayout";

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
        element: <AppLayout/>,
        children:[
            {
                path:"/mail-boxes",
                element: <MailBox/>
            }

        ]
    }
]
