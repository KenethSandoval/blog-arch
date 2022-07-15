import { IPost } from "../types/post";
import type { NextPage, GetStaticProps } from 'next';
import Link from 'next/link';
import { getAllFiles } from '../lib/mdx';
import Thumbnail from "../components/Thumbnail";

// props type
type Props = {
  posts: [IPost]
};


const Home: NextPage<Props> = ({ posts }: Props) => {
    return (
			<div>
				<h1 className="text-4xl font-bold mb-4">Documentaciones</h1>
				<div className="space-y-12">
					{posts.map((post) => (
						<div key={post.slug}>
							<div className="mb-4">
								<Thumbnail
									slug={post.slug}
									title={post.title}
									src={post.thumbnail}
								/>
							</div>

							<h2 className="text-2xl font-bold mb-4">
								<Link href={`/post/${post.slug}`}>
									<a>{post.title}</a>
								</Link>
							</h2>
						</div>
					))}
				</div>
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
