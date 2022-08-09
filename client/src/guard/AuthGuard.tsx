import React, { useEffect } from "react";
import { useRouter } from "next/router";
import { setRedirect, useAuth } from "@global-store/useAuth"
// import { useAuth } from "./AuthProvider";

interface Props {
	children: React.ReactNode;
};

export const AuthGuard: React.FC<Props> = ({ children }: Props) => {
	const { checking } = useAuth((state) => state);
	const router = useRouter();

	useEffect(() => {
		if (!checking) {
				setRedirect(router.route);
				router.push("/login");
			}
	}, [checking, router]);

	/* show loading indicator while the auth provider is still isLoading */
  if (checking) {
    return (
      <div className="absolute h-screen w-screen z-50">
        <div className="flex overflow-hidden fixed top-0 right-0 bottom-0 left-0 z-50 flex-col justify-center items-center w-full h-screen bg-opacity-80 bg-gray-300">
          <div className="mb-4 w-12 h-12 rounded-full border-4 border-t-4 border-gray-200 ease-linear loader animate-spinner" />
          <h2 className="text-xl font-semibold text-center text-primary-base">
            Loading...
          </h2>
        </div>
      </div>
    );
  }

	if (!checking) {
		return <>{children}</>;
	}

	return null;
}
