import { Axios } from "./config.service";
import { LoginData } from "@interface/auth.response";
import { USERTOKEN } from "@libs/localStorageItems";

export const login = async (credential: LoginData) => {
	return await Axios.post<string>('/auth/login', credential);
}

export const setTokenApi = (token: string) => {
	try {
		localStorage.setItem(USERTOKEN, token);
	} catch(error) {
		console.log(error, "Error setTokenApi");
	}
}

export const checkUserLogged = () => !!localStorage.getItem(USERTOKEN);

