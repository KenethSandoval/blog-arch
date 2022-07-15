import type { AppProps } from 'next/app';
import Layout from '../components/Layout';
import 'tailwindcss/tailwind.css';
import { ThemeProvider } from "next-themes"
import { MdxComponentsProvider } from '../context/mdxContext';

export default function MyApp({ Component, pageProps }: AppProps) {
    return (
			<ThemeProvider enableSystem={true} attribute="class">
				<MdxComponentsProvider>
				  <Layout>
			    	<Component {...pageProps} />
			  	</Layout>
				</MdxComponentsProvider>
		    
			</ThemeProvider>
		);
}
