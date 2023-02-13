import React, {Fragment, useEffect, useState} from 'react';
import {useParams} from "react-router";
import {restApi} from "../../../api";
import {MailMessage} from "../../types/mailbox";
import {Divider, List, ListItem, ListItemText, Tab, Tabs} from "@mui/material";
import {decodeQuotedPrintable} from "../../utils/quote-printable";
import Box from "@mui/material/Box";
import TabPanel, {a11yProps} from "../../components/tab/tab-panel";

const ViewMessage = () => {
    const [value, setValue] = React.useState(1);

    const params = useParams();
    const [mailMessage, setMailMessage] = useState<MailMessage | null>(null);
    useEffect(() => {
        restApi.get<MailMessage>(`/mail-messages/${params.messageId}/detail`).then((res) => {
            setMailMessage(res.data)
        })
    }, [params.messageId])
    const handleChange = (_:any, newValue:number) => {
        setValue(newValue);
    };

    if (mailMessage == null) {
        return (
            <div></div>
        )
    }
    const generateHtmlContent = () => {
        let contents:any = {};
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
        contents["original"] = decodeQuotedPrintable(bodies.join("\n"))
        let blockStartCharacter = "";
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
        return contents;
    }
    const contents = generateHtmlContent();
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
            <Box sx={{ borderBottom: 1, borderColor: 'divider' }}>
                <Tabs value={value} onChange={handleChange} aria-label="basic tabs example">
                    <Tab label="HTML" {...a11yProps(0)} />
                    <Tab label="HTML source" {...a11yProps(1)} />
                    <Tab label="Plain" {...a11yProps(2)} />
                    <Tab label="Raw Content" {...a11yProps(3)} />
                    <Tab label="Headers" {...a11yProps(4)} />
                </Tabs>
                <TabPanel value={value} index={0}>
                    {contents['text/html'] && <div dangerouslySetInnerHTML={{__html:decodeQuotedPrintable(contents['text/html'])}}/>}
                </TabPanel>
                <TabPanel index={1} value={value}>
                    {contents['text/html'] && <div>{
                        decodeQuotedPrintable(contents['text/html']).split("\n")
                            .map((item,idx)=>{
                                    return (
                                        <Fragment key={idx}>
                                            {item}<br/>
                                        </Fragment>
                                    )
                                }
                            )
                    }</div>}
                </TabPanel>
                <TabPanel index={2} value={value}>
                    {contents['text/plain'] && <div>{
                        decodeQuotedPrintable(contents['text/plain']).split("\n")
                            .map((item,idx)=>{
                                    return (
                                        <Fragment key={idx}>
                                            {item}<br/>
                                        </Fragment>
                                    )
                                }
                            )
                    }</div>
                    }
                </TabPanel>
                <TabPanel index={3} value={value}>
                        {
                            mailMessage.message.split("\n")
                            .map((item,idx)=>{
                                    return (
                                        <Fragment key={idx}>
                                            {item}<br/>
                                        </Fragment>
                                    )
                                }
                            )
                    }
                </TabPanel>
                <TabPanel index={4} value={value}>
                    <List dense={true}>
                    {
                        Object.keys(mailMessage.headers).map((value) =>{
                            return (
                                <ListItem key={"header"+value}>
                                    <ListItemText primary={value} secondary={mailMessage?.headers[value]}/>
                                </ListItem>
                            )
                        })
                    }
                    </List>
                </TabPanel>
            </Box>
        </div>
    );
};
export default ViewMessage;