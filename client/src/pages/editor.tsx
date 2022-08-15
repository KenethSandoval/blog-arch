import {  useState} from "react";
import ReactMarkdown from "react-markdown";
import remarkGfm from "remark-gfm";
import FileList from "../components/FileList"
import Nav from "../components/Nav"
import { Prism as SyntaxHighlighter } from 'react-syntax-highlighter';
import atomDark from 'react-syntax-highlighter/dist/cjs/styles/prism/atom-dark';

export default function Editor() {
	const [show, setShow] = useState(false);
	const [displayWindow, setDisplayWindow] = useState(false);
	const [newFile, setNewFile] = useState("");
	const [value, setValue] = useState("");

  const createStory = () => {
			console.log(newFile);
	}

  const setFile = (dat: boolean) => {
		setDisplayWindow(dat)
	}

  const handleChange = (e:any) => {
		setNewFile(e.target.value);
	}

	const handleInputChange = (e:any) => {
			setValue(e.target.value)
	}
	
 	return (
		<>
			<Nav display={setFile}/>
			<div className="relative bg-slate-900">
				<div className="relative h-full flex flex-row pt-28">
					<FileList />
					<div className="w-3/4 border border-slate-900 py-8 px-8">
						<button
							className="cursor-pointer relative text-blue-500 font-semibold text-md"
							onClick={() => {
								setShow(!show);
							}}
						>
							{show ? "Edit" : "Preview"}
						</button>
						{!show ? (
								<textarea 
									className="bg-slate-900 fullheight w-full relative outline-none text-white border-0 pt-6"
									placeholder="Markdown here"
									value={value}
									onChange={handleInputChange}
								/>
							) : (
								// preview window
              	<div className="bg-slate-900 h-full w-full text-white editor">
                	<ReactMarkdown 
									  remarkPlugins={[remarkGfm]}
										components={{
      								code({node, inline, className, children, ...props}) {
        								const match = /language-(\w+)/.exec(className || '')
        								return !inline && match ? (
          								<SyntaxHighlighter
            								children={String(children).replace(/\n$/, '')}
														style={atomDark}
            								language={match[1]}
            								PreTag="div" 
														{...props}
          								/>
          							) : (
            						<code className={className} {...props}>
              						{children} 
												</code>
          						);
        						},
      						}}
									>
                  	{value}
                	</ReactMarkdown>
              	</div>
							)}
					</div>
				</div>
				{/* new file creation window */}
				{displayWindow ? (
					<div className="absolute h-full w-full flex justify-center items-center top-0 backdrop-blur-lg">
						<div className="bg-blue-500 roundend-lg p-4">
							<div className="flex flex-col justify-center items-center gap-8">
                <div className=" relative w-full flex flex-row justify-between">
                  <h1 className=" font-medium uppercase text-white">
                    Create a New File
                  </h1>
                  <button
                    onClick={() => {
                      setDisplayWindow(false);
                    }}
                  >
                    Close
                  </button>
                </div>
								<div className="flex flex-row gap-2 justify-center items-center">
									<div className="flex-1">
										<input
											className="roundend-lg p-4 text-black"
											type="text"
											placeholder="File name"
											value={newFile}
											onChange={handleChange}
										/>
									</div>
									<div className="flex-1">
									   <button
                      className="text-white bg-slate-900 px-4 py-3 rounded-md"
                      onClick={() => {
                        createStory();
                      }}
                    >
                      Create File
                    </button>
									</div>
								</div>
							</div>
						</div>
					</div>
				): null}
			</div>
		</>
  );
}
