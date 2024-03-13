/* @refresh reload */
import { render } from "solid-js/web";
import { Router, Route } from "@solidjs/router";

import "./index.css";
import App from "./App";
import AllBooks from "./routes/все-книги/all-books";
import Debtors from "./routes/должники/debtors";
import Home from "./routes/home/home";

const root = document.getElementById("root");

render(
  () => (
    <Router root={App}>
      <Route path="/" component={Home} />
      <Route path="/все-книги" component={AllBooks} />
      <Route path="/должники" component={Debtors} />
    </Router>
  ),
  root!
);
