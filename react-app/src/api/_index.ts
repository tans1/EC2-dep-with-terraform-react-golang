import axios from "axios";


const getToken = () =>{
    const res = localStorage.getItem('token');
    return res;
}
export const apiInstance = axios.create({
    baseURL: import.meta.env.VITE_BASE_URL,
    timeout: 1000,
    headers: {
        'Authorization': `Bearer ${getToken()}`
    }
});