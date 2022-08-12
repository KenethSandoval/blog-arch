import axios from "axios";
import { getToken } from "@libs/localStorageItems";

const Authorization = getToken();

//TODO: change environment
const LOCAL_URL = "http://localhost:3000";

export const Axios = axios.create({
	baseURL: LOCAL_URL,
	headers: {
		Authorization
	},
});
