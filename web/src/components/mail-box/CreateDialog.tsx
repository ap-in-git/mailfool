import Dialog from '@mui/material/Dialog';
import DialogTitle from '@mui/material/DialogTitle';
import React from 'react';
import {DialogActions, DialogContent, Grid} from "@mui/material";
import TextField from "@mui/material/TextField";
import {useForm} from "react-hook-form";
import {restApi} from "../../../api";
import Button from "@mui/material/Button";
import Box from "@mui/material/Box";

interface Props {
    dialogOpen: boolean,
    setDialogOpen: React.Dispatch<React.SetStateAction<boolean>>
}

const CreateDialog: React.FC<Props> = ({dialogOpen, setDialogOpen}) => {
    const {register, handleSubmit, formState: {errors}} = useForm<{
        name: string
        username: string
        password: string
        max_size: number
        tls_enabled: boolean
    }>();

    const onSubmit = handleSubmit(data => {
        restApi.post("/auth/login", data)
    })

    return (
        <Dialog open={dialogOpen} aia-labelledby="form-dialog-title" maxWidth={"sm"} fullWidth={true}>
            <form onSubmit={onSubmit}>
                <DialogTitle id="form-dialog-title">Add new mailbox </DialogTitle>
                <DialogContent>
                    <Grid container spacing={2}>
                        <Grid item xs={12}>
                            <TextField
                                size="small"
                                id="name"
                                label="Name * "
                                type="text"
                                fullWidth
                                {...register("name", {
                                    required: "Name is required"
                                })}
                                helperText={errors.name && errors.name.message}
                                error={!!errors.name}
                                variant={"outlined"}
                            />
                        </Grid>
                        <Grid item xs={12}>
                            <TextField
                                size="small"
                                id="name"
                                label="Username * "
                                type="text"
                                fullWidth
                                {...register("username", {
                                    required: "Username is required"
                                })}
                                helperText={errors.username && errors.username.message}
                                error={!!errors.username}
                                variant={"outlined"}
                            />
                        </Grid>
                        <Grid item xs={12}>
                            <TextField
                                size="small"
                                id="password"
                                label="Password * "
                                type="text"
                                fullWidth
                                {...register("username", {
                                    required: "Password is required"
                                })}
                                helperText={errors.password && errors.password.message}
                                error={!!errors.password}
                                variant={"outlined"}
                            />
                        </Grid>
                        <Grid item xs={12}>
                            <TextField
                                size="small"
                                id="maximum_mail_size"
                                label="Maximum mail size * "
                                type="number"
                                fullWidth
                                {...register("max_size", {
                                    required: "Maximum mail is required"
                                })}
                                helperText={errors.max_size && errors.max_size.message}
                                error={!!errors.max_size}
                                variant={"outlined"}
                            />
                        </Grid>
                    </Grid>
                </DialogContent>
                <DialogActions>
                    <Button color={"error"} size={"small"}>Close</Button>
                    <Button color={"primary"} type={"submit"} size={"small"}>Create</Button>
                </DialogActions>
            </form>
        </Dialog>
    );
};

export default CreateDialog;