import React from "react";
import ReactDOM from "react-dom";
import { Route, Routes, BrowserRouter as Router } from "react-router-dom";
import "./index.css";
import Welcome from "./Welcome";
import Status from "./Status";
import Ping from "./Ping";

const routing = (
  <Router>
    <Routes>
      <Route exact path="/welcome" element={<Welcome />}></Route>
      <Route exact path="/status" element={<Status />}></Route>
      <Route exact path="/ping" element={<Ping />}></Route>
    </Routes>
  </Router>
);

ReactDOM.render(routing, document.getElementById("root"));
