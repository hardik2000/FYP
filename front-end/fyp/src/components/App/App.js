import React from 'react';
import './App.css';
// import { BrowserRouter, Route, Routes } from 'react-router-dom';

import { BrowserRouter, Route, Routes } from 'react-router-dom';
import Dashboard from '../Dashboard/Dashboard';
import Login from '../Login/Login';
import Add from '../Add/Add';
import Query from '../Query/Query';
import Update from '../Update/Update';
import History from '../History/History';
import Delete from '../Delete/Delete';
import Verify from '../Verify/Verify';
// import Register from '../Register/Register';

import getToken from './getToken';

function App() {
    const token = getToken();
    console.log(token);
    if (!token) {
        console.log("Inside if of App")
        return <Login />
    }
    console.log("Hello")
    return (
        <div className="wrapper">
            <h1>Application</h1>
            <BrowserRouter>
                <Routes>
                    <Route exact path="" element={<Dashboard />}>
                    </Route>
                    <Route exact path="/Add" element={<Add />}>
                    </Route>
                    <Route exact path="/Query" element={<Query />}>
                    </Route>
                    <Route exact path="/Update" element={<Update />}>
                    </Route>
                    <Route exact path="/History" element={<History />}>
                    </Route>
                    <Route exact path="/Delete" element={<Delete />}>
                    </Route>
                    <Route exact path="/Verify" element={<Verify />}>
                    </Route>
                </Routes>
            </BrowserRouter>
        </div>
    );
}

export default App;