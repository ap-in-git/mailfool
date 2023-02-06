import * as React from 'react';
import Avatar from '@mui/material/Avatar';
import Button from '@mui/material/Button';
import CssBaseline from '@mui/material/CssBaseline';
import TextField from '@mui/material/TextField';
import Box from '@mui/material/Box';
import LockOutlinedIcon from '@mui/icons-material/LockOutlined';
import Typography from '@mui/material/Typography';
import Container from '@mui/material/Container';
import {createTheme, ThemeProvider} from '@mui/material/styles';
import {useForm} from "react-hook-form";
import {restApi} from "../../api";


const theme = createTheme();

export default function SignIn() {
    const {register,handleSubmit,formState:{errors}} = useForm<{
        email: string
        password: string
    }>();

    const onSubmit = handleSubmit( data => {
        restApi.post("/auth/login",data)
    })
    return (
        <ThemeProvider theme={theme}>
            <Container component="main" maxWidth="xs">
                <CssBaseline />
                <Box
                    sx={{
                        marginTop: 8,
                        display: 'flex',
                        flexDirection: 'column',
                        alignItems: 'center',
                    }}
                >
                    <Avatar sx={{ m: 1, bgcolor: 'secondary.main' }}>
                        <LockOutlinedIcon />
                    </Avatar>
                    <Typography component="h1" variant="h5">
                        Sign in
                    </Typography>
                    <Box component="form" onSubmit={onSubmit} noValidate sx={{ mt: 1 }}>
                        <TextField
                            fullWidth={true}
                            id="email"
                            label="Email *"
                            size={"small"}
                            type={"text"}
                            variant={"outlined"}
                            {...register('email',{
                                required: "Email is required"
                            })}
                            error={!!errors.email}
                            helperText={errors.email && errors.email.message}
                            InputLabelProps={{
                                shrink: true,
                            }}
                        />
                        <TextField
                            fullWidth={true}
                            id="password"
                            label="Password *"
                            size={"small"}
                            sx={{mt:3}}
                            type={"password"}
                            variant={"outlined"}
                            {...register('password',{
                                required: "Password is required"
                            })
                            }
                            helperText={errors.password && errors.password.message}
                            error={!!errors.password}
                            InputLabelProps={{
                                shrink: true,
                            }}
                        />
                        <Button
                            type="submit"
                            fullWidth
                            variant="contained"
                            sx={{ mt: 3, mb: 2 }}
                        >
                            Sign In
                        </Button>
                    </Box>
                </Box>
            </Container>
        </ThemeProvider>
    );
}