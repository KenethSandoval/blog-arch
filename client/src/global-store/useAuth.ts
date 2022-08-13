import create from "zustand";
import { LoginData, RegisterData } from "@interface/auth.response";
import { login, setTokenApi } from '@service/auth.service';

const redirectKey = "@doc/redirect";

export function setRedirect(redirect: string) {
	window.sessionStorage.setItem(redirectKey, redirect);
}

export function getRedirect(): string | null {
  return window.sessionStorage.getItem(redirectKey);
}

export function clearRedirect() {
	window.sessionStorage.removeItem(redirectKey);
}

interface AuthStateI {
  auth: boolean;
  checking: boolean;
	startAuth: (dataLogin: LoginData) => Promise<void>;
  startRegister: (dataRegister: RegisterData) => Promise<void>;
  startLogout: () => void;
  startChecking: () => void;
}

export const useAuth = create<AuthStateI>(
	(set) => ({
			auth: false,
			checking: false,
			startAuth: async (dataLogin: LoginData) => {
				try {
					const response = await login(dataLogin);

					set((state) => ({
						...state,
						auth: !state.auth,
						checking: true,
					}));
					
					setTokenApi(response.data);
				} catch (error) {
					console.log("error is start login", error);
				}
			},
			startRegister: async () => {
			},
			startLogout: () => {
				console.log("logout");
				set((state) => ({...state, auth: false}));
			},
			startChecking: () => {
				console.log("checking");
				set((state) => ({...state, auth: true}));
			}
	})
);
