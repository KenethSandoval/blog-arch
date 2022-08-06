import 'tailwindcss/tailwind.css';
import '../../styles/globals.css';
import type { AppProps } from 'next/app';
import { ThemeProvider } from "next-themes"
import Router from 'next/router';
import { NextPage } from 'next';

import NProgress from "nprogress"

import Layout from '../components/Layout';
import { MdxComponentsProvider } from '../context/mdxContext';
import { AuthGuard } from "../guard/AuthGuard";
import { AuthProvider } from "../guard/AuthProvider";

/*Router.events.on("routeChangeStart", () => {
	NProgress.start();
});*/

export type NextApplicationPage<P = {}, IP = P> = NextPage<P, IP> & {
		auth?: boolean;
};

export default function MyApp(props: AppProps) {
		const { Component, pageProps } : { Component: NextApplicationPage; pageProps: any } = props;
    return (
			<>
				<ThemeProvider enableSystem={true} attribute="class">
					<MdxComponentsProvider>
				  	<Layout>
						<AuthProvider>
							{Component.auth ? (
								<AuthGuard>
			    				<Component {...pageProps} />
								</AuthGuard>
							) : (
			    				<Component {...pageProps} />
							)
							}
						</AuthProvider>
			  		</Layout>
					</MdxComponentsProvider>
				</ThemeProvider>
			</>
		);
}
