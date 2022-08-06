import { combine, devtools, persist } from "zustand/middleware";
import create from "zustand";
import { LoginData, RegisterData } from "@types/auth.response";

interface AuthStateI {
  auth: boolean;
  checking: boolean;
}

interface AuthMethods {
  startAuth: (dataLogin: LoginData) => Promise<void>;
  startRegister: (dataRegister: RegisterData) => Promise<void>;
  startLogout: () => void;
  startChecking: () => void;
}

export const useAuth = create(
  devtools(
    persist(
      combine<AuthStateI, AuthMethods>(
        {
          auth: false,
          checking: false,
        },
        (set) => ({
          startAuth: async (dataLogin: LoginData) => {
            try {
							console.log("start auth")
            } catch (error) {
              console.log("error in start Login", error);
            }
          },
          startRegister: async (dataRegister: RegisterData) => {
            try {
							console.log("start register")
            } catch(error) {
              console.log("error in start Login", error);
            }
          },
          startLogout: () => {
						console.log("logout")
          },
          startChecking: () => {
            set((state) => ({
              ...state,
              auth: false,
              checking: false,
            }));
          },
        })
      ),
      {
        name: "@fh/info",
      }
    )
  )
);
