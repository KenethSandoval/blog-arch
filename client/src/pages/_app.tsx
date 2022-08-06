import 'tailwindcss/tailwind.css';
import type { AppProps } from 'next/app';
import { ThemeProvider } from "next-themes"
import Router from 'next/router';
import { NextPage } from 'next';
import { AuthGuard } from "src/guard/AuthGuard";

import Layout from '../components/Layout';
import { MdxComponentsProvider } from '../context/mdxContext';


/*Router.events.on("routeChangeStart", () => {
	console.log("start");
});*/

export type NextApplicationPage<P = {}, IP = P> = NextPage<P, IP> & {
		auth?: boolean;
};

export default function MyApp(props: AppProps) {
		const { Component, pageProps } : { Component: NextApplicationPage; pageProps: any } = props;
    return (
			<>
				{Component.auth ? (
						<AuthGuard>
							<ThemeProvider enableSystem={true} attribute="class">
								<MdxComponentsProvider>
				  				<Layout>
			    					<Component {...pageProps} />
			  					</Layout>
								</MdxComponentsProvider>
							</ThemeProvider>
						</AuthGuard>
					) : (
						<>
							<Component {...pageProps} />
						</>
					)
				}	
			</>
		);
}
