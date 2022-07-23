import Link from "next/link";

type Props = {
	title: string;
	slug?: string;
	stacks:string[];
};

const Thumbnail: React.FC<Props> = ({ title, slug, stacks }: Props) => {
	return (
		<>
			<div key={slug}>
							<div className="bg-gray-800 text-gray-50 mb-4">
								<div className="container grid grid-cols-12 mx-auto bg-gray-900">
									<div className="flex flex-col p-6 col-span-full row-span-full lg:col-span-8 lg:p-10">
										<div className="flex justify-start">
											{stacks.map((stack, index) => (
												<div key={index}>
													<div className="mx-1">
														<span className="px-2 py-1 text-xs rounded-full bg-violet-400 text-gray-900">{stack}</span> 
													</div>
												</div>	
											))}	
										</div>
										<h1 className="text-3xl font-semibold">{title}</h1>
										<Link href={`/posts/${slug}`}>
											<a className="inline-flex items-center pt-2 pb-6 space-x-2 text-sm text-violet-400">
												<span>Read more</span>
												<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor" className="w-4 h-4">
												<path fillRule="evenodd" d="M12.293 5.293a1 1 0 011.414 0l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414-1.414L14.586 11H3a1 1 0 110-2h11.586l-2.293-2.293a1 1 0 010-1.414z" clipRule="evenodd"></path>
												</svg>
											</a>
										</Link>	
									</div>
								</div> 
							</div> 	
				</div>	
		</>
	);
}

export default Thumbnail;
