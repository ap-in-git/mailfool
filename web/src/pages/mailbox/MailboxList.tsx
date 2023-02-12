import {restApi} from "../../../api";
import {useEffect, useState} from "react";
import {Mailbox} from "../../types/mailbox";
import Card from "@mui/material/Card";
import Typography from "@mui/material/Typography";
import CardContent from "@mui/material/CardContent";
import CardActions from "@mui/material/CardActions";
import Button from "@mui/material/Button";
import {Grid,List, ListItem, ListItemAvatar, ListItemText} from "@mui/material";
import Avatar from "@mui/material/Avatar";
import {AiOutlineUser} from "react-icons/ai"
import {RiLockPasswordLine} from "react-icons/ri"
import {MdStorage,MdLock} from "react-icons/md"
import CreateDialog from "../../components/mail-box/CreateDialog";
import useNotificationStore from "../../store/notification";
import {Link} from "react-router-dom";
const MailboxList = () => {
    const [mailBoxes, setMailBoxes] = useState<Mailbox[]>([]);
    const [dialogOpen,setDialogOpen] = useState(false);
    const {showSuccess,showError} = useNotificationStore((state) =>state)
    useEffect(() => {
        fetchMailbox().finally()
    }, []);
    const fetchMailbox = async () => {
        restApi.get<Mailbox[]>("/mail-boxes").then((res) => {
            setMailBoxes(res.data)
        })
    }

    const deleteInbox = async (id: number) => {
       if (confirm("Are you sure?"))  {
           try {
               const response =  await restApi.delete(("/mail-boxes/"+id))
               showSuccess(response.data.message)
               fetchMailbox().finally()
           }catch (e:any) {
               showError(e.response.data.message)
           }

       }
    }

    return (
        <Grid container spacing={4}>
            <Grid item xs={12}>
                <Button variant={"contained"} size={"small"} color={"primary"} onClick={()=>{
                    setDialogOpen(true)
                }}>Add new inbox</Button>
                <CreateDialog dialogOpen={dialogOpen} setDialogOpen={setDialogOpen} fetchMailBox={fetchMailbox}/>
            </Grid>
            {
                mailBoxes.map((mailBox) => {
                    return (
                        <Grid item xs={4} key={"mailbox" + mailBox.id}>
                            <Card >
                                <CardContent>
                                    <Typography sx={{fontSize: 14}} color="text.secondary" gutterBottom>
                                        {mailBox.name}
                                    </Typography>
                                    <List>
                                        <ListItem>
                                            <ListItemAvatar>
                                                <Avatar>
                                                    <AiOutlineUser/>
                                                </Avatar>
                                            </ListItemAvatar>
                                            <ListItemText primary="Username" secondary={mailBox.user_name}/>
                                        </ListItem>
                                        <ListItem>
                                            <ListItemAvatar>
                                                <Avatar>
                                                    <RiLockPasswordLine/>
                                                </Avatar>
                                            </ListItemAvatar>
                                            <ListItemText primary="Password" secondary={mailBox.password}/>
                                        </ListItem>
                                        <ListItem>
                                            <ListItemAvatar>
                                                <Avatar>
                                                    <MdStorage/>
                                                </Avatar>
                                            </ListItemAvatar>
                                            <ListItemText primary="Maximum email size" secondary={mailBox.maximum_size + ' MB'}/>
                                        </ListItem>
                                        <ListItem>
                                            <ListItemAvatar>
                                                <Avatar>
                                                    <MdLock/>
                                                </Avatar>
                                            </ListItemAvatar>
                                            <ListItemText primary="TLS enabled" secondary={mailBox.tls_enabled?"Yes":"No"}/>
                                        </ListItem>
                                    </List>
                                </CardContent>
                                <CardActions>
                                    <Link style={{textDecoration:"none"}} to={"/mail-boxes/"+mailBox.id}>
                                        <Button size="small">View Inbox</Button>
                                    </Link>
                                    <Button size="small" color={"error"} onClick={()=>{
                                        deleteInbox(mailBox.id)
                                    }}>Delete Inbox</Button>
                                </CardActions>
                            </Card>
                        </Grid>
                    )
                })
            }
        </Grid>
    );
};

export default MailboxList;