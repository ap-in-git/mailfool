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
