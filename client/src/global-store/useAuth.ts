import { combine, devtools, persist } from 'zustand/middleware';
import create from 'zustand';
// add toastify
import { RegisterData, LoginData } from '@types/auth.response';
// add service

const redirectKey = "@md/redirect";

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
}

interface AuthMethods {
	startAuth: (dataLogin: LoginData) => Promise<void>;
	startRegister: (dataRegister: RegisterData) => Promise<void>;
	startLogout: () => void;
}

export const useAuth = create(
	devtools(
		persist(
			combine<AuthStateI, AuthMethods> (
				{
					auth: false,
					checking: false,
					// set user
				},
				(set) => ({
					startAuth: async (dataLogin: LoginData) => {
						
					},
					startRegister: async (data: RegisterData) => {

					},
					startLogout: () => {

					},
				})
			),
			{
				name: "@md/info",
			}
		)
	)
);
