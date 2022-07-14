import { IPost } from "../types/post";
import type { NextPage, GetStaticProps } from 'next';
import Link from 'next/link';
import { getAllFiles } from '../lib/mdx';

// props type
type Props = {
  posts: [IPost]
};


const Home: NextPage<Props> = ({ posts }: Props) => {
    return (
			<div>
				<h1 className="text-4xl font-bold mb-4">Documentaciones</h1>	
			</div>
    );
}
export default Home

export const getStaticProps: GetStaticProps = async () => {
	const posts = getAllFiles([
		'title',
		'slug',
		'date',
		'description',
		'thumbnail'
	]);

	return {
			props: {
					posts
			}
	};
}
