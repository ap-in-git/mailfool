import {create} from "zustand";

interface Notification {
    message: string | null
    type: string
    clearMessage: () => void
    showSuccess: (message: string) => void
    showError: (message: string) => void
}

const useNotificationStore = create<Notification>()((set) => ({
    message: null,
    type: "",
    showSuccess: (message: string) => set(() => ({message: message, type: "success"})),
    showError: (message: string) => set(() => ({message: message, type: "error"})),
    clearMessage: () => set(() => ({message: null, type: ""}))
}));

export default useNotificationStore;