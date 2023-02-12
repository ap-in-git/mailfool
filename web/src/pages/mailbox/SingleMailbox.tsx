import React, {Fragment, useEffect, useState} from 'react';
import {restApi} from "../../../api";
import {Checkbox, Divider, List, ListItem, ListItemButton, ListItemIcon, ListItemText} from "@mui/material";
import {MailMessage} from "../../types/mailbox";
import {Link} from "react-router-dom";
import {useParams} from "react-router";

const SingleMailbox = () => {
    const params = useParams();
    const mailId = params.id;
    const [mailMessages, setMailMessages] = useState<MailMessage[]>([]);
    useEffect(() => {
        restApi.get<MailMessage[]>(`/mail-messages/${mailId}`).then((res) => {
            setMailMessages(res.data)
        })
    }, [mailId]);
    return (
        <List>
            {mailMessages.map((mailMessage) => {
                const primaryText = `From: ${mailMessage.sender}, To: ${mailMessage.receiver}`
                return (
                    <Fragment
                        key={'mailMessage' + mailMessage.id}
                    >
                        <ListItem
                            secondaryAction={<Fragment>
                                {mailMessage.created_at.toString()}
                            </Fragment>}
                        >
                            <ListItemButton
                                component={Link}
                                to={`/mail-messages/${mailId}/messages/${mailMessage.id}`}
                                role={undefined} dense>
                                <ListItemIcon>
                                    <Checkbox
                                        edge="start"
                                        tabIndex={-1}
                                        disableRipple
                                    />
                                </ListItemIcon>
                                <ListItemText id={mailMessage.id.toString()} primary={primaryText} secondary={<Fragment>
                                  Subject: {mailMessage.subject}
                                </Fragment>}/>
                            </ListItemButton>
                        </ListItem>
                        <Divider/>
                    </Fragment>
                )
            })}
        </List>
    );
};

export default SingleMailbox;