import React, {useEffect, useState} from 'react';
import {useParams} from "react-router";
import {restApi} from "../../../api";
import {MailMessage} from "../../types/mailbox";
import {List, ListItem, ListItemText} from "@mui/material";

const ViewMessage = () => {
    const params = useParams();
    const [mailMessage, setMailMessage] = useState<MailMessage | null>(null);
    useEffect(() => {
        restApi.get<MailMessage>(`/mail-messages/${params.messageId}/detail`).then((res) => {
            setMailMessage(res.data)
        })
    }, [params.messageId])

    if (mailMessage == null) {
        return (
            <div></div>
        )
    }
    const generateHtmlContent = () => {
        return (
            <div>
                {mailMessage.message}
                Lorem ipsum dolor sit amet, consectetur adipisicing elit. Adipisci, autem commodi deleniti dolor eaque
                maiores minus molestiae mollitia non, pariatur perferendis quae quia quis recusandae repudiandae
                sapiente sed sit tenetur!
            </div>
        )
    }
    return (
        <div>
            <List dense={true}>
                <ListItem>
                    <ListItemText primary="From" secondary={mailMessage.sender}/>
                </ListItem>
                <ListItem>
                    <ListItemText primary="To" secondary={mailMessage.receiver}/>
                </ListItem>
                <ListItem>
                    <ListItemText primary="Subject" secondary={mailMessage.subject}/>
                </ListItem>
            </List>
            {generateHtmlContent()}
        </div>
    );
};

export default ViewMessage;