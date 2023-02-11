import {restApi} from "../../api";
import {useEffect, useState} from "react";
import {Mailbox} from "../types/mailbox";
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
import CreateDialog from "../components/mail-box/CreateDialog";
const MailBox = () => {
    const [mailBoxes, setMailBoxes] = useState<Mailbox[]>([]);
    const [dialogOpen,setDialogOpen] = useState(true);
    useEffect(() => {
        fetchMailbox().finally()
    }, []);
    const fetchMailbox = async () => {
        restApi.get<Mailbox[]>("/mail-boxes").then((res) => {
            setMailBoxes(res.data)
        })
    }
    return (
        <Grid container spacing={4}>
            <Grid item xs={12}>
                <Button size={"small"} color={"primary"} onClick={()=>{
                    setDialogOpen(true)
                }}>Add new inbox</Button>
                <CreateDialog dialogOpen={dialogOpen} setDialogOpen={setDialogOpen}/>
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
                                    <Button size="small">View Inbox</Button>
                                </CardActions>
                            </Card>
                        </Grid>
                    )
                })
            }
        </Grid>
    );
};

export default MailBox;