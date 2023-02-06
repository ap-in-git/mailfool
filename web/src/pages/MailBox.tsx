import {restApi} from "../../api";
import {useEffect, useState} from "react";
import {Mailbox} from "../types/mailbox";

const MailBox = () => {
    const [mailBoxes,setMailBoxes] = useState<Mailbox[]>([])
    useEffect(() => {
        restApi.get<Mailbox[]>("/mail-boxes").then((res)=>{
            setMailBoxes(res.data)
        })
    },[])
    return (
        <div>
           Lorem ipsum dolor sit amet, consectetur adipisicing elit. Enim, labore, odit! Autem dicta dolore, doloribus dolorum eaque eos esse expedita facere mollitia, perspiciatis praesentium rem, sunt ullam veniam veritatis voluptatem.
        </div>
    );
};

export default MailBox;