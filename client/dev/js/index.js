import "babel-polyfill";
import React from "react";
import ReactDOM from "react-dom";
import {Provider} from "react-redux";
//import { Router, browserHistory } from 'react-router';
import configureStore from "./store/configure_store";
import App from "./components/app.jsx";

const store = configureStore();

ReactDOM.render(
    <Provider store={store}>
        <App />
    </Provider>
    , document.getElementById('app')
);
