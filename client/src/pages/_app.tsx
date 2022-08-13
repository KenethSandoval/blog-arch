import 'tailwindcss/tailwind.css';
import '../../styles/globals.css';
import { useEffect } from 'react';
import type { AppProps } from 'next/app';
import { ThemeProvider } from "next-themes"
import { NextPage } from 'next';
import Router from "next/router";
import NProgress from "nprogress"

import Layout from '../components/Layout';
import { MdxComponentsProvider } from '../context/mdxContext';
import { AuthGuard } from "../guard/AuthGuard";
import { useAuth } from "@global-store/useAuth";
import { useRouter } from "next/router";

Router.events.on("routeChangeStart", () => {
	NProgress.start();
});

export type NextApplicationPage<P = {}, IP = P> = NextPage<P, IP> & {
		auth?: boolean;
};

export default function MyApp(props: AppProps) {
	const { auth } = useAuth((state) => state);
  const router = useRouter();
  useEffect(() => {
    if (auth) {
      router.replace("/");
    }
  }, [auth, router]);
		const { Component, pageProps } : { Component: NextApplicationPage; pageProps: any } = props;
    return (
			<>
				<ThemeProvider enableSystem={true} attribute="class">
					<MdxComponentsProvider>
				  	<Layout>
							{Component.auth ? (
								<AuthGuard>
			    				<Component {...pageProps} />
								</AuthGuard>
							) : (
			    				<Component {...pageProps} />
							)
							}
			  		</Layout>
					</MdxComponentsProvider>
				</ThemeProvider>
			</>
		);
}
