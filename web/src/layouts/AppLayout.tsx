import React from 'react';
import {Alert, AppBar, Link, Snackbar, Toolbar} from "@mui/material";
import Typography from "@mui/material/Typography";
import Button from "@mui/material/Button";
import {Outlet} from "react-router";
import Box from "@mui/material/Box";
import useNotificationStore from "../store/notification";

const AppLayout = () => {
    return (
        <div>
            <AppBar
                position="static"
                color="default"
                elevation={0}
            >
                <Toolbar sx={{flexWrap: 'wrap'}}>
                    <Typography variant="h6" color="inherit" noWrap sx={{flexGrow: 1}}>
                        Mail fool
                    </Typography>
                    <nav>
                        <Link
                            variant="button"
                            color="text.primary"
                            href="#"
                            sx={{my: 1, mx: 1.5}}
                        >
                            Features
                        </Link>
                        <Link
                            variant="button"
                            color="text.primary"
                            href="#"
                            sx={{my: 1, mx: 1.5}}
                        >
                            Enterprise
                        </Link>
                        <Link
                            variant="button"
                            color="text.primary"
                            href="#"
                            sx={{my: 1, mx: 1.5}}
                        >
                            Support
                        </Link>
                    </nav>
                    <Button href="#" variant="outlined" sx={{my: 1, mx: 1.5}}>
                        Login
                    </Button>
                </Toolbar>
            </AppBar>
            <NotificationMessage/>
            <Box sx={{p: 2}}>
                <Outlet/>
            </Box>
        </div>
    );
};

const NotificationMessage = () => {
    const {message, type, clearMessage} = useNotificationStore((state) => state)
    return (
        <>
            {
                message && <Snackbar
                    anchorOrigin={{vertical: 'top', horizontal: 'right'}}
                    open={true}
                    autoHideDuration={4000}
                    onClose={() => {
                        clearMessage()
                    }}
                >
                    <Alert severity={type as any}>
                        {message}
                    </Alert>
                </Snackbar>
            }

        </>
    )
}
export default AppLayout;