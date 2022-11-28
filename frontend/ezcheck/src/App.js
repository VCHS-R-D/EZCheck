import './App.css';
import Landing from './pages/Landing.js';
import { BrowserRouter, Routes, Route } from "react-router-dom";
import React from "react";
import Admin from './pages/Admin.js';
import Student from './pages/Student.js';


function App() {

  return (
    <div className="App">
       <BrowserRouter>
        <Routes>
          <Route path="/" element={<Landing/>} />
          <Route path="/student" element={<Student/>} />
          <Route path="/admin" element={<Admin/>} />
        </Routes>
       </BrowserRouter>
    </div>
  );
}

export default App;