import React from 'react'
import Modal from 'react-modal';
import { Link, Navigate, useNavigate } from "react-router-dom";
import axios from "axios";
import { useCookies } from 'react-cookie';
import { Buffer } from 'buffer'
var FormData = require('form-data');

const customStyles = {
  content: {
    top: '50%',
    left: '50%',
    marginRight: "auto",
    marginLeft: "auto",
    right: 'auto',
    bottom: 'auto',
    // marginRight: '-50%',
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
  const [cookie, setCookie, removeCookie] = useCookies('user');

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
      data: formdata
    };
    axios(config)
      .then(function(response) {
        if (String(response.data) === "success") {
          const token = `${username}:${password}`;
          const encodedToken = Buffer.from(token).toString('base64');
          setCookie('authToken', encodedToken, { path: '/' });
          handleUserSignin()
        }
      })
      .catch(function(error) {
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
      data: formdata
    };
    axios(config)
      .then(function(response) {
        ;
        if (String(response.data) === "success") {
          const token = `${username}:${password}`;
          const encodedToken = Buffer.from(token).toString('base64');
          setCookie('authToken', encodedToken, { path: '/' });
          handleAdminSignin()
        }
      })
      .catch(function(error) {
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
      data: formdata
    };
    axios(config)
      .then(function(response) {
        ;

        setCookie('authToken', encodedToken, { path: '/' });
        setCookie('userID', response.data.id, { path: '/' });
        localStorage.setItem("student", JSON.stringify(response.data));
        navigate("/student")
      })
      .catch(function(error) {
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
      data: formdata
    };
    axios(config)
      .then(function(response) {
        ;
        setCookie('authToken', encodedToken, { path: '/' });
        setCookie('adminID', response.data.id, { path: '/' });
        navigate("/admin")

      })
      .catch(function(error) {
        console.log(error);
      });
  }

  const renderForm = () => {
    if (userType === "admin") {
      return (
        <React.Fragment>
          {isLogin ? (
            <div>
              <form>
                <div class="grid gap-x-3 mb-6 md:grid-cols-1">
                  <div className='flex flex-row rounded-sm m-2 text-sm text-gray-700'>
                    <label className='m-auto text-gray-500 font-bold uppercase justify-start w-1/2'>Username</label>
                    <input className='w-1/2 justify-start appearance-none rounded-full nm-inset-gray-200 leading-5 px-8 py-4 flex-grow sm:w-2/3' placeholder="Username" onChange={(event) => { setUsername(event.target.value) }}></input>
                  </div>
                  <div className='flex flex-row rounded-sm m-2 text-sm text-gray-700'>
                    <label className='m-auto text-gray-500 font-bold uppercase justify-start w-1/2'>Password</label>
                    <input className='w-1/2 justify-start appearance-none rounded-full nm-inset-gray-200 leading-5 px-8 py-4 flex-grow sm:w-2/3' placeholder="Password" onChange={(event) => { setUsername(event.target.value) }}></input>
                  </div>
                </div>
              </form>
              <button className=' text-gray-700 text-sm rounded-full nm-flat-gray-200 hover:nm-flat-gray-200-lg leading-5 px-8 py-4 uppercase font-bold tracking-widest transition duration-200 ease-in-out transform hover:scale-110' onClick={handleAdminCreate}>Submit</button>
            </div>
          ) : (
            <span>
              <form>
                <div class="grid gap-x-3 mb-6 md:grid-cols-1">
                  <div className='flex flex-row rounded-sm m-2 text-sm text-gray-700'>
                    <label className='m-auto text-gray-500 font-bold uppercase justify-start w-1/2'>Username</label>
                    <input className='w-1/2 justify-start appearance-none rounded-full nm-inset-gray-200 leading-5 px-8 py-4 flex-grow sm:w-2/3' placeholder="Username" onChange={(event) => { setUsername(event.target.value) }}></input>
                  </div>
                  <div className='flex flex-row rounded-sm m-2 text-sm text-gray-700'>
                    <label className='m-auto text-gray-500 font-bold uppercase justify-start w-1/2'>Password</label>
                    <input className='w-1/2 justify-start appearance-none rounded-full nm-inset-gray-200 leading-5 px-8 py-4 flex-grow sm:w-2/3' placeholder="Password" onChange={(event) => { setUsername(event.target.value) }}></input>
                  </div>
                  <div className='flex flex-row rounded-sm m-2 text-sm text-gray-700'>
                    <label className='m-auto text-gray-500 font-bold uppercase justify-start w-1/2'>First Name</label>
                    <input className='w-1/2 justify-start appearance-none rounded-full nm-inset-gray-200 leading-5 px-8 py-4 flex-grow sm:w-2/3' placeholder="First Name" onChange={(event) => { setFirstName(event.target.value) }}></input>
                  </div>
                  <div className='flex flex-row rounded-sm m-2 text-sm text-gray-700'>
                    <label className='m-auto text-gray-500 font-bold uppercase justify-start w-1/2'>Last Name</label>
                    <input className='w-1/2 justify-start appearance-none rounded-full nm-inset-gray-200 leading-5 px-8 py-4 flex-grow sm:w-2/3' placeholder="Last Name" onChange={(event) => { setLastName(event.target.value) }}></input>
                  </div>
                  <div className='flex flex-row rounded-sm m-2 text-sm text-gray-700'>
                    <label className='m-auto text-gray-500 font-bold uppercase justify-start w-1/2'>Admin Code</label>
                    <input className='w-1/2 justify-start appearance-none rounded-full nm-inset-gray-200 leading-5 px-8 py-4 flex-grow sm:w-2/3' placeholder="Admin Code" onChange={(event) => { setAdminCode(event.target.value) }}></input>
                  </div>
                  <div className='flex flex-row rounded-sm m-2 text-sm text-gray-700'>
                    <label className='m-auto text-gray-500 font-bold uppercase justify-start w-1/2'>Code</label>
                    <input className='w-1/2 justify-start appearance-none rounded-full nm-inset-gray-200 leading-5 px-8 py-4 flex-grow sm:w-2/3' placeholder="Code" onChange={(event) => { setCode(event.target.value) }}></input>
                  </div>
                </div>
              </form>
              <button className=' text-gray-700 text-sm rounded-full nm-flat-gray-200 hover:nm-flat-gray-200-lg leading-5 px-8 py-4 uppercase font-bold tracking-widest transition duration-200 ease-in-out transform hover:scale-110' onClick={handleAdminCreate}>Submit</button>
            </span>
          )
          }

        </React.Fragment>
      )
    }
    else {
      return (
        <React.Fragment>
          {isLogin ? (
            <span>
              <form>
                <div class="grid gap-x-3 mb-6 md:grid-cols-1">
                  <div className='flex flex-row rounded-sm m-2 text-sm text-gray-700'>
                    <label className='m-auto text-gray-500 font-bold uppercase justify-start w-1/2'>Username</label>
                    <input className='w-1/2 justify-start appearance-none rounded-full nm-inset-gray-200 leading-5 px-8 py-4 flex-grow sm:w-2/3' placeholder="Username" onChange={(event) => { setUsername(event.target.value) }}></input>
                  </div>
                  <div className='flex flex-row rounded-sm m-2 text-sm text-gray-700'>
                    <label className='m-auto text-gray-500 font-bold uppercase justify-start w-1/2'>Password</label>
                    <input className='w-1/2 justify-start appearance-none rounded-full nm-inset-gray-200 leading-5 px-8 py-4 flex-grow sm:w-2/3' placeholder="Password" onChange={(event) => { setUsername(event.target.value) }}></input>
                  </div>
                </div>
              </form>
              <button className=' text-gray-700 text-sm rounded-full nm-flat-gray-200 hover:nm-flat-gray-200-lg leading-5 px-8 py-4 uppercase font-bold tracking-widest transition duration-200 ease-in-out transform hover:scale-110' onClick={handleUserSignin}>Submit</button>
            </span>
          ) : (
            <span>
              <form>
                <div class="grid gap-x-3 mb-6 md:grid-cols-1">
                  <div className='flex flex-row rounded-sm m-2 text-sm text-gray-700'>
                    <label className='m-auto text-gray-500 font-bold uppercase justify-start w-1/2'>Username</label>
                    <input className='w-1/2 justify-start appearance-none rounded-full nm-inset-gray-200 leading-5 px-8 py-4 flex-grow sm:w-2/3' placeholder="Username" onChange={(event) => { setUsername(event.target.value) }}></input>
                  </div>
                  <div className='flex flex-row rounded-sm m-2 text-sm text-gray-700'>
                    <label className='m-auto text-gray-500 font-bold uppercase justify-start w-1/2'>Password</label>
                    <input className='w-1/2 justify-start appearance-none rounded-full nm-inset-gray-200 leading-5 px-8 py-4 flex-grow sm:w-2/3' placeholder="Password" onChange={(event) => { setUsername(event.target.value) }}></input>
                  </div>
                  <div className='flex flex-row rounded-sm m-2 text-sm text-gray-700'>
                    <label className='m-auto text-gray-500 font-bold uppercase justify-start w-1/2'>First Name</label>
                    <input className='w-1/2 justify-start appearance-none rounded-full nm-inset-gray-200 leading-5 px-8 py-4 flex-grow sm:w-2/3' placeholder="First Name" onChange={(event) => { setFirstName(event.target.value) }}></input>
                  </div>
                  <div className='flex flex-row rounded-sm m-2 text-sm text-gray-700'>
                    <label className='m-auto text-gray-500 font-bold uppercase justify-start w-1/2'>Last Name</label>
                    <input className='w-1/2 justify-start appearance-none rounded-full nm-inset-gray-200 leading-5 px-8 py-4 flex-grow sm:w-2/3' placeholder="Last Name" onChange={(event) => { setLastName(event.target.value) }}></input>
                  </div>
                  <div className='flex flex-row rounded-sm m-2 text-sm text-gray-700'>
                    <label className='m-auto text-gray-500 font-bold uppercase justify-start w-1/2'>Grade</label>
                    <input className='w-1/2 justify-start appearance-none rounded-full nm-inset-gray-200 leading-5 px-8 py-4 flex-grow sm:w-2/3' placeholder="Grade" onChange={(event) => { setGrade(event.target.value) }}></input>
                  </div>
                  <div className='flex flex-row rounded-sm m-2 text-sm text-gray-700'>
                    <label className='m-auto text-gray-500 font-bold uppercase justify-start w-1/2'>Code</label>
                    <input className='w-1/2 justify-start appearance-none rounded-full nm-inset-gray-200 leading-5 px-8 py-4 flex-grow sm:w-2/3' placeholder="Code" onChange={(event) => { setCode(event.target.value) }}></input>
                  </div>
                </div>
              </form>
              <button className=' text-gray-700 text-sm rounded-full nm-flat-gray-200 hover:nm-flat-gray-200-lg leading-5 px-8 py-4 uppercase font-bold tracking-widest transition duration-200 ease-in-out transform hover:scale-110' onClick={handleStudentCreate}>Submit</button>
            </span>
          )
          }
        </React.Fragment >
      )
    }
  }

  return (
    <div className='m-10 flex flex-row w-full gap-5'>
      <div class="nm-flat-gray-200-lg rounded-lg p-8 text-center max-w-sm w-full">
        <h2 class="text-2xl font-bold leading-tight mb-4">Admin Sign Up</h2>
        {/* <button className="font-bold py-2 px-4 rounded bg-blue-500 hover:bg-blue-700 text-white" onClick={adminSignup}>Admin Sign Up</button> */}
        <button onClick={adminSignup} class="rounded-full bg-green-500 shadow-gray-200 leading-5 px-8 py-4 uppercase font-bold tracking-widest text-white inline-block mt-4">Sign Up</button>
      </div>
      <div class="nm-flat-gray-200-lg rounded-lg p-8 text-center max-w-sm w-full">
        <h2 class="text-2xl font-bold leading-tight mb-4">Admin Login</h2>
        {/* <button className="font-bold py-2 px-4 rounded bg-blue-500 hover:bg-blue-700 text-white" onClick={adminSignup}>Admin Sign Up</button> */}
        <button onClick={adminLogin} class="rounded-full bg-green-500 shadow-gray-200 leading-5 px-8 py-4 uppercase font-bold tracking-widest text-white inline-block mt-4">Login</button>
      </div>
      <div class="nm-flat-gray-200-lg rounded-lg p-8 text-center max-w-sm w-full">
        <h2 class="text-2xl font-bold leading-tight mb-4">Student Sign Up</h2>
        {/* <button className="font-bold py-2 px-4 rounded bg-blue-500 hover:bg-blue-700 text-white" onClick={adminSignup}>Admin Sign Up</button> */}
        <button onClick={studentSignup} class="rounded-full bg-green-500 shadow-gray-200 leading-5 px-8 py-4 uppercase font-bold tracking-widest text-white inline-block mt-4">Sign Up</button>
      </div>
      <div class="nm-flat-gray-200-lg rounded-lg p-8 text-center max-w-sm w-full">
        <h2 class="text-2xl font-bold leading-tight mb-4">Student Login</h2>
        {/* <button className="font-bold py-2 px-4 rounded bg-blue-500 hover:bg-blue-700 text-white" onClick={adminSignup}>Admin Sign Up</button> */}
        <button onClick={studentLogin} class="rounded-full bg-green-500 shadow-gray-200 leading-5 px-8 py-4 uppercase font-bold tracking-widest text-white inline-block mt-4">Login</button>
      </div>
      <Modal
        isOpen={modalIsOpen}
        onRequestClose={closeModal}
        style={customStyles}
        className=""
      >
        <button onClick={closeModal} type="button" class="bg-red-500 rounded-lg p-1.5 mt-0 inline-flex items-center justify-center text-black hover:bg-red-600 focus:outline-none focus:ring-2 focus:ring-inset focus:ring-indigo-500">
        </button>
        <div>
          <h1 className='my-2 text-xl text-gray-500 font-bold uppercase'>Login</h1>
          {renderForm()}
        </div>
      </Modal>
    </div >
  )
}
