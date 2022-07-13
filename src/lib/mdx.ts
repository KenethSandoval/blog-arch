import fs from "fs";
import { join } from "path";
import matter from "gray-matter";
import { serialize } from "next-mdx-remote/serialize";

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
const POST_PATH = join(process.cwd(), '__posts');