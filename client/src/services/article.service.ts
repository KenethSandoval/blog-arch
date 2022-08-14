import { Article } from "@interface/article";
import { Axios } from "./config.service";

export const newArticle = async (data: Article) => {
	return await Axios.post<string>('/articles', data);
}
