export const USERTOKEN = "@doc/token";

let item:string= "";
if (typeof window !== 'undefined') {
  // Perform localStorage action
  item = localStorage.getItem("@doc/token") || "";
}
export const getToken = (token?: string) => item;
