export interface Mailbox {
    id: number;
    created_at: string;
    updated_at: string;
    deleted_at?: string;
    name: string;
    user_name: string;
    password: string;
    tls_enabled: boolean;
    maximum_size: number;
}

export interface MailMessage {
    id: number;
    created_at: Date;
    updated_at: Date;
    sender: string;
    subject: string;
    receiver: string;
    message: string;
    mail_box_id: number;
    headers: any
}