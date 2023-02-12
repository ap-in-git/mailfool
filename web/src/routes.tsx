import {RouteObject} from "react-router";
import Home from "./pages/Home";
import SignIn from "./pages/SignIn";
import MailboxList from "./pages/mailbox/MailboxList";
import SingleMailbox from "./pages/mailbox/SingleMailbox";
import AppLayout from "./layouts/AppLayout";
import ViewMessage from "./pages/mailbox/ViewMessage";
import MailboxLayout from "./layouts/MailboxLayout";

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
        children: [
            {
                path: "/mail-boxes",
                element: <MailboxList/>
            },
            {
                path: "mail-messages/:id",
                element: <MailboxLayout/>,
                children: [
                    {
                        element: <SingleMailbox/>,
                        path: ""
                    },
                    {
                        element: <ViewMessage/>,
                        path: "messages/:messageId"
                    }
                ]
            }
        ]
    }
]
