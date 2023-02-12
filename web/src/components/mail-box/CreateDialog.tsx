import Dialog from '@mui/material/Dialog';
import DialogTitle from '@mui/material/DialogTitle';
import React from 'react';
import {Checkbox, DialogActions, DialogContent, FormControlLabel, Grid} from "@mui/material";
import TextField from "@mui/material/TextField";
import {useForm} from "react-hook-form";
import {restApi} from "../../../api";
import Button from "@mui/material/Button";
import {Controller} from "react-hook-form";
import useNotificationStore from "../../store/notification";

interface Props {
    dialogOpen: boolean,
    setDialogOpen: React.Dispatch<React.SetStateAction<boolean>>
    fetchMailBox: () => Promise<any>
}

const CreateDialog: React.FC<Props> = ({fetchMailBox,dialogOpen, setDialogOpen}) => {
    const {showSuccess,showError} = useNotificationStore((state) =>state)
    const {register,control, handleSubmit, formState: {errors}} = useForm<{
        name: string
        username: string
        password: string
        max_size: number
        tls_enabled: boolean
    }>();

    const onSubmit = handleSubmit(async data => {
        try {
            const response = await restApi.post("/mail-boxes", data)
            showSuccess(response.data.message)
            setDialogOpen(false)
            fetchMailBox().finally()
        }catch (e:any) {
            showError(e.response.data.message)
        }
    })

    return (
        <Dialog open={dialogOpen} aia-labelledby="form-dialog-title" maxWidth={"sm"} fullWidth={true}>
            <form onSubmit={onSubmit}>
                <DialogTitle id="form-dialog-title">Add new mailbox </DialogTitle>
                <DialogContent>
                    <Grid container spacing={2} >
                        <Grid item xs={12} style={{marginTop:10}}>
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
                                    required: "Username is required",
                                    pattern: {
                                        message:"Username can only contains number and alphabet",
                                        value:/^[a-zA-Z\d]+$/
                                    }
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
                                {...register("password", {
                                    required: "Password is required",
                                    pattern: {
                                        message:"Password cannot contains :",
                                        value:/^[^:]*$/
                                    }
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
                                    required: "Maximum mail size is required",
                                    valueAsNumber:true
                                })}
                                helperText={errors.max_size && errors.max_size.message}
                                error={!!errors.max_size}
                                variant={"outlined"}
                            />
                        </Grid>
                        <Grid item xs={12}>
                            <FormControlLabel control={
                                <Controller
                                    name="tls_enabled"
                                    control={control}
                                    render={({ field: props }) => (
                                        <Checkbox
                                            {...props}
                                            onChange={(e) => props.onChange(e.target.checked)}
                                        />
                                    )}
                                />
                            } label="TLS enabled" />
                        </Grid>
                    </Grid>
                </DialogContent>
                <DialogActions>
                    <Button color={"error"} size={"small"} onClick={()=>{
                        setDialogOpen(false)
                    }}>Close</Button>
                    <Button color={"primary"} type={"submit"} size={"small"}>Create</Button>
                </DialogActions>
            </form>
        </Dialog>
    );
};

export default CreateDialog;