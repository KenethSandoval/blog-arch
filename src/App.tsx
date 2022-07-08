import { useState } from "react";
import { Routes, Route } from "react-router-dom";
import "@wcj/dark-mode";
import "./App.css";
import MDEditor from "@uiw/react-md-editor";
// import MDStr from "./README.md";

const Editor = () => {
  const [value, setValue] = useState<string | undefined>("");
  return (
    <>
      <div className="App-editor">
        <MDEditor value={value} onChange={setValue} />
        <MDEditor.Markdown source={value} style={{ whiteSpace: "pre-wrap" }} />
      </div>
    </>
  );
};

const Documentation = () => {
  return (
    <>
      <p>working...</p>
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
