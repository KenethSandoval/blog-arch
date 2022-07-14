import type { AppProps } from 'next/app';
import Layout from '../components/Layout';
import 'tailwindcss/tailwind.css';
import { ThemeProvider } from "next-themes"

export default function MyApp({ Component, pageProps }: AppProps) {
    return (
			<ThemeProvider enableSystem={true} attribute="class">
		    <Layout>
			    <Component {...pageProps} />
			  </Layout>
			</ThemeProvider>
		);
}
