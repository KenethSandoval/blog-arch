import { IPost } from "../types/post";
import type { GetStaticProps } from 'next';
import { getAllFiles } from '../lib/mdx';
import Thumbnail from "../components/Thumbnail";

// props type
type Props = {
  posts: [IPost]
};

export default function Home ({ posts }: Props) {
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
										stacks={post.stacks}
									/>
								</div>
							</div>
					))}
				</div>
			</div>
    );
}
// export default Home

export const getStaticProps: GetStaticProps = async () => {
	const posts = getAllFiles([
		'title',
		'slug',
		'date',
		'description',
		'stacks'
	]);

	return {
			props: {
					posts
			}
	};
}
