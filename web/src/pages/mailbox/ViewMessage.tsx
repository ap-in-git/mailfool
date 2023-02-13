import React, {useEffect, useState} from 'react';
import {useParams} from "react-router";
import {restApi} from "../../../api";
import {MailMessage} from "../../types/mailbox";
import {Divider, List, ListItem, ListItemText} from "@mui/material";
import Box from "@mui/material/Box";
import Markdown from "markdown-to-jsx";
const decodeQuotedPrintable = (data:string):string => {
    // normalise end-of-line signals
    data = data.replace(/(\r\n|\n|\r)/g, "\n");

    // replace equals sign at end-of-line with nothing
    data = data.replace(/=\n/g, "");

    // encoded text might contain percent signs
    // decode each section separately
    let bits = data.split("%");
    for (let i = 0; i < bits.length; i ++)
    {
        // replace equals sign with percent sign
        bits[i] = bits[i].replace(/=/g, "%");

        // decode the section
        bits[i] = decodeURIComponent(bits[i]);
    }

    // join the sections back together
    return(bits.join("%"));
}
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
        console.log(mailMessage.message)
        const messageLines = mailMessage.message.split("\n");
        let bodyStartIndex = -1;
        for (let i = 0; i < messageLines.length; i++) {
            const line = messageLines[i];
            if (line == "") {
                bodyStartIndex = i + 1;
                break;
            }
        }
        const bodies = messageLines.slice(bodyStartIndex)
        let blockStartCharacter = "";
        let contents:any = {};
        let contentBodies = [];
        let contentType = "";
        let bodyStart = -1;
        for (let i = 0; i < bodies.length; i++) {
            const line = bodies[i];
            if(line == blockStartCharacter + "--"){
                contents[contentType] = contentBodies.join("\n")
            }
            //Mail contains plain text or html text
            if (line.startsWith("--")) {
                if (blockStartCharacter == line) {
                    contents[contentType] = contentBodies.join("\n")
                    contentBodies = [];
                    bodyStart = -1;
                }else{
                    blockStartCharacter = line
                }
            }
            if (blockStartCharacter != "") {
                if (bodyStart == -1){
                    const headers = line.split(":")
                    if (headers[0] == "Content-Type") {
                        contentType = headers[1].trim().split(";")[0]
                    }
                    if (line == "" && contentType != "") {
                        bodyStart = 1;
                    }
                }
                if (bodyStart > -1) {
                    contentBodies.push(line)
                }
            }
        }

        // const message = messageLines.slice(bodyStartIndex).join("\n")
        return (
            <div dangerouslySetInnerHTML={{__html: decodeQuotedPrintable(contents['text/html'])}}>
            </div>
        )
    }
    return (
        <div>
            <List dense={true}>
                <ListItem>
                    <ListItemText primary="From" secondary={mailMessage.headers.From}/>
                </ListItem>
                <ListItem>
                    <ListItemText primary="To" secondary={mailMessage.headers.To}/>
                </ListItem>
                <ListItem>
                    <ListItemText primary="Subject" secondary={mailMessage.subject}/>
                </ListItem>
            </List>
            <Divider/>
            {generateHtmlContent()}
        </div>
    );
};

export default ViewMessage;