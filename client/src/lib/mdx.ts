import fs from "fs";
import matter from "gray-matter";
import { join } from "path";

// Items struct of items
type Items = {
	[key: string]: string;
};

// Post struct of posts
type Post = {
	data: {
		[key: string]: string;
	};
	// each post will include the post content associated with its parameter key
	content: string;
};

// path to our list posts available
const POST_PATH = join(process.cwd(), 'src/_posts');

export const getFiles = (): string[] => {
	return fs.readdirSync(POST_PATH)
		.filter((path) => /\.mdx?$/.test(path));
}

export const getFileBySlug = (slug: string): Post => {
	const fullPath = join(POST_PATH, `${slug}.mdx`);

	const fileContents = fs.readFileSync(fullPath, 'utf-8');

	const { data, content } = matter(fileContents);

	return { data, content }
}

export const getAllFilesMetadata = (filePath: string, fields: string[] = []): Items => {
	const slug = filePath.replace(/\.mdx?$/, "");

	const { data, content } = getFileBySlug(slug);

	const items: Items = {};

	fields.forEach((field) => {
		if (field === 'slug') {
			items[field] = slug;
		}

		if (field === 'content') {
			items[field] = content;
		}

		if (data[field]) {
			items[field] = data[field];
		}
	});

	return items;
}

export const getAllFiles = (fields: string[]): Items[] => {
	const filePaths = getFiles();

	const posts = filePaths.map((filepath) => getAllFilesMetadata(filepath, fields)).sort((post1, post2) => post1.date > post2.date ? 1 : -1);

	return posts;
}



