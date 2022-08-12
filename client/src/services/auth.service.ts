import { Axios } from "./config.service";
import { LoginData } from "@interface/auth.response";
// import { USERTOKEN } from "@libs/localStorageItems";

export const login = async (credential: LoginData) => {
	const response = await Axios.post<string>('/auth/login', credential);
	console.log(response);
}


