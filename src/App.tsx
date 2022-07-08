import { useState, useEffect, useRef } from "react";
import { Routes, Route } from "react-router-dom";
import "@wcj/dark-mode";
import "./App.css";
import ReadmeStr from './README.md';
import MDEditor, { commands } from "@uiw/react-md-editor";
import mermaid from "mermaid";

const randomid = () => parseInt(String(Math.random() * 1e15), 10).toString(36);

const Code = ({ node, inline, className, children, ...props }: any) => {
  const demoid = useRef(`dome${randomid()}`);
  const code = getCode(children);
  const demo = useRef(null);
  useEffect(() => {
    if (demo.current) {
      try {
        const str = mermaid.render(
          demoid.current,
          code,
          () => null,
          demo.current
        );
        // @ts-ignore
        demo.current.innerHTML = str;
      } catch (error) {
        // @ts-ignore
        demo.current.innerHTML = error;
      }
    }
  }, [code, demo]);

  if (
    typeof code === "string" &&
    typeof className === "string" &&
    /^language-mermaid/.test(className.toLocaleLowerCase())
  ) {
    return (
      <code ref={demo}>
        <code id={demoid.current} style={{ display: "none" }} />
      </code>
    );
  }
  return <code className={String(className)}>{children}</code>;
};

const getCode: string | boolean | any = (arr: any[] = []) =>
  arr
    .map((dt) => {
      if (typeof dt === "string") {
        return dt;
      }
      if (dt.props && dt.props.children) {
        return getCode(dt.props.children);
      }
      return false;
    })
    .filter(Boolean)
    .join("");

const Editor = () => {
  const [value, setValue] = useState("");
  return (
    <MDEditor
      onChange={(newValue = "") => setValue(newValue)}
      textareaProps={{
        placeholder: "Please enter Markdown text",
      }}
      height={500}
      value={value}
      previewOptions={{
        components: {
          code: Code,
        },
      }}
      commands={[commands.codeEdit, commands.codePreview]}
      extraCommands={[
        commands.group(
          [
            commands.title1,
            commands.title2,
            commands.title3,
            commands.title4,
            commands.title5,
            commands.title6,
          ],
          {
            name: "title",
            groupName: "title",
            buttonProps: { "aria-label": "Insert title" },
          }
        ),
        commands.divider,
        commands.divider,
        commands.fullscreen,
      ]}
    />
  );
};

const Documentation = () => { 
  return (
    <>
      <MDEditor
        value={ReadmeStr}
        height={500}
        style={{ whiteSpace: "pre-wrap", margin: "0 100px" }}
        preview={"preview"}
        hideToolbar={true}
        previewOptions={{
          components: {
            code: Code,
          },
        }}
      />
    </>
  );
};

const App = () => {
  return (
    <div className="App">
      <dark-mode permanent light="Light" dark="Dark"></dark-mode>
      <Routes>
        <Route path="/" element={<Editor />} />
        <Route path="doc" element={<Documentation />} />
      </Routes>
    </div>
  );
};

export default App;
