import axios from "axios";

export const restApi = axios.create({
    baseURL:  "127.0.0.1:5000/api/v1"
});
