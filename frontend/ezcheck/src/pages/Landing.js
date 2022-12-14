import React from 'react'
import Modal from 'react-modal';
import { Link, Navigate, useNavigate } from "react-router-dom";
import axios from "axios";
import {useCookies} from 'react-cookie';
import { Buffer } from 'buffer'
var FormData = require('form-data');

const customStyles = {
    content: {
      top: '50%',
      left: '50%',
      right: 'auto',
      bottom: 'auto',
      marginRight: '-50%',
      transform: 'translate(-50%, -50%)',
    },
  };

export default function Landing() {
    const [userType, setUserType] = React.useState("");
    const [isLogin, setIsLogin] = React.useState(false);
    const [modalIsOpen, setIsOpen] = React.useState(false);
    const [username, setUsername] = React.useState("");
    const [password, setPassword] = React.useState("");
    const [firstName, setFirstName] = React.useState("");
    const [lastName, setLastName] = React.useState("");
    const [code, setCode] = React.useState("");
    const [grade, setGrade] = React.useState("");
    const [adminCode, setAdminCode] = React.useState("");
    const navigate = useNavigate();
    const [cookie, setCookie,removeCookie] = useCookies('user');

    function adminSignup() {
        setUserType("admin");
        setIsLogin(false);
        openModal(true);
    }

    function adminLogin() {
        setUserType("admin");
        setIsLogin(true);
        openModal(true);
    }

    function studentSignup() {
        setUserType("student");
        setIsLogin(false);
        openModal(true);
    }

    function studentLogin() {
        setUserType("student");
        setIsLogin(true);
        openModal(true);
    }

    function openModal() {
        setIsOpen(true);
    }

    function closeModal() {
        setIsOpen(false);
    }
    
    function handleStudentCreate() {
        var formdata = new FormData();
        formdata.append("username", username);
        formdata.append("password", password);
        formdata.append("first", firstName);
        formdata.append("last", lastName);
        formdata.append("code", code);
        formdata.append("grade", grade);
        var config = {
        method: 'post',
        url: 'http://localhost:8080/user/create',
        headers: formdata.getHeaders ? formdata.getHeaders() : { 'Content-Type': 'multipart/form-data' },
        data : formdata
        };
        axios(config)
        .then(function (response) {
        if(String(response.data) === "success"){
            const token = `${username}:${password}`;
            const encodedToken = Buffer.from(token).toString('base64');
            setCookie('authToken', encodedToken, { path: '/'});
            handleUserSignin()
        }
        })
        .catch(function (error) {
        console.log(error);
        });
    }

    function handleAdminCreate() {
        var formdata = new FormData();
        formdata.append("username", username);
        formdata.append("password", password);
        formdata.append("first", firstName);
        formdata.append("last", lastName);
        formdata.append("code", code);
        formdata.append("adminPass", adminCode);
        var config = {
        method: 'post',
        url: 'http://localhost:8080/admin/create',
        headers: formdata.getHeaders ? formdata.getHeaders() : { 'Content-Type': 'multipart/form-data' },
        data : formdata
        };
        axios(config)
        .then(function (response) {;
        if(String(response.data) === "success"){
            const token = `${username}:${password}`;
            const encodedToken = Buffer.from(token).toString('base64');
            setCookie('authToken', encodedToken, { path: '/'});
            handleAdminSignin()
        }
        })
        .catch(function (error) {
        console.log(error);
        });
    }

    function handleUserSignin() {
        var formdata = new FormData();
        formdata.append("username", username);
        formdata.append("password", password);
        
        const token = `${username}:${password}`;
        const encodedToken = Buffer.from(token).toString('base64');
        var config = {
        method: 'post',
        url: 'http://localhost:8080/user/get',
        headers: {
            'Authorization': `Basic ${encodedToken}`,
        },
        data : formdata
        };
        axios(config)
        .then(function (response) {;
            
            setCookie('authToken', encodedToken, { path: '/'});
            setCookie('userID', response.data.id, { path: '/'});
            localStorage.setItem("student", JSON.stringify(response.data));
            navigate("/student")
        })
        .catch(function (error) {
        console.log(error);
        });
    }

    function handleAdminSignin() {
        var formdata = new FormData();
        formdata.append("username", username);
        formdata.append("password", password);
        
        const token = `${username}:${password}`;
        const encodedToken = Buffer.from(token).toString('base64');
        var config = {
        method: 'post',
        url: 'http://localhost:8080/admin/get',
        headers: {
            'Authorization': `Basic ${encodedToken}`,
        },
        data : formdata
        };
        axios(config)
        .then(function (response) {;
            setCookie('authToken', encodedToken, { path: '/'});
            setCookie('adminID', response.data.id, { path: '/'});
            navigate("/admin")
            
        })
        .catch(function (error) {
        console.log(error);
        });
    }

    const renderForm = () => {
        if(userType === "admin"){
            return(
                <React.Fragment>
                    {isLogin ? (
                        <div>
                            <form>
                            <input placeholder="username" onChange={(event) => {setUsername(event.target.value)}}></input>
                                <input placeholder="password" onChange={(event) => {setPassword(event.target.value)}}></input>
                            </form>
                            <button onClick={handleAdminSignin}>Submit</button>
                        </div>
                    ) : (
                        <span>
                            <form>
                                <input placeholder="username" onChange={(event) => {setUsername(event.target.value)}}></input>
                                <input placeholder="password" onChange={(event) => {setPassword(event.target.value)}}></input>
                                <input placeholder="first name" onChange={(event) => {setFirstName(event.target.value)}}></input>
                                <input placeholder="last name" onChange={(event) => {setLastName(event.target.value)}}></input>
                                <input placeholder="code" onChange={(event) => {setCode(event.target.value)}}></input>
                                <input placeholder="admin code" onChange={(event) => {setAdminCode(event.target.value)}}></input>
                            </form>
                            <button onClick={handleAdminCreate}>Submit</button>
                        </span>
                    )
                    }
                   
                </React.Fragment>
            )
        }
        else {
            return(
                <React.Fragment>
                    {isLogin ? (
                        <span>
                            <form>
                                <input placeholder="username" onChange={(event) => {setUsername(event.target.value)}}></input>
                                <input placeholder="password" onChange={(event) => {setPassword(event.target.value)}}></input>
                            </form>
                            <button onClick={handleUserSignin}>Submit</button>
                        </span>
                    ) : (
                        <span>
                            <form>
                                <input placeholder="username" onChange={(event) => {setUsername(event.target.value)}}></input>
                                <input placeholder="password" onChange={(event) => {setPassword(event.target.value)}}></input>
                                <input placeholder="first name" onChange={(event) => {setFirstName(event.target.value)}}></input>
                                <input placeholder="last name" onChange={(event) => {setLastName(event.target.value)}}></input>
                                <input placeholder="code" onChange={(event) => {setCode(event.target.value)}}></input>
                                <input placeholder="grade" onChange={(event) => {setGrade(event.target.value)}}></input>
                            </form>
                            <button onClick={handleStudentCreate}>Submit</button>
                        </span>
                    )
                    }
                </React.Fragment>
            )
        }
    }
    
    return (
        <div>
            <button onClick={adminSignup}>Admin Sign Up</button>
            <button onClick={adminLogin}>Admin Log In</button>
            <button onClick={studentSignup}>Student Sign Up</button>
            <button onClick={studentLogin}>Student Log In</button>
            <Modal
                isOpen={modalIsOpen}
                onRequestClose={closeModal}
                style={customStyles}
                contentLabel="Example Modal"
            >
                <button onClick={closeModal}>close</button>
                <h1>Login</h1>
                {renderForm()}
            </Modal>
                </div>
    )
}
