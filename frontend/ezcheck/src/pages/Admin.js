import React from 'react'
import Modal from 'react-modal';
import MachineList from '../components/MachineList';
import Search from '../components/Search.js';
import Logs from '../components/logs.js';
import Select from 'react-select';
import axios from "axios";
import {Cookies, useCookies, removeCookie} from 'react-cookie';
import { useNavigate } from "react-router-dom";

export default function Admin() {
    const [show, setShow] = React.useState(false);
    const [cookie, setCookie, removeCookie] = useCookies('user');
    const [currentPage, setCurrentPage] = React.useState("");
    const student = JSON.parse(localStorage.getItem("student"));
    const [options, setOptions] = React.useState([""]);
    const [selectedOption, setSelectedOption] = React.useState("");
    const [isLoading, setLoading] = React.useState(false);
    const navigate = useNavigate();
    React.useEffect(() => {onStart()}, [])
    
    function handleOptionChange(selectedOption)
    {
        setSelectedOption({selectedOption});
    }

    function onStart(){
        getMachines();
        handleStudentSearch();
        setLoading(false);
    }

    function getMachines() {
        const arr = [];
        var config = {
            method: 'get',
            url: 'http://localhost:8080/machines',
            headers: { }
        };
        axios(config)
        .then(function (response) {
            response.data.map((machine) => { return arr.push({label: machine.id, value: machine.id});
        });
            setOptions(arr);
        })
        .catch(function (error) {
            console.log(error);
        });
    }

    function handleCertify() {
        var formdata = new FormData();
        formdata.append("adminID", cookie.adminID);
        formdata.append("userID", student.id);
        formdata.append("machineID", selectedOption.selectedOption.value);
        var config = {
        method: 'post',
        url: 'http://localhost:8080/admin/certify',
        headers: {
            'Authorization': `Basic ${cookie.authToken}`,
        },
        data : formdata
        };
        axios(config)
        .then(function (response) {
            alert(response.data);
            handleStudentSearch();
            renderStudent();
            renderPages();
            
        })
        .catch(function (error) {
        console.log(error);
        });
    }

    function handleUncertify() {
        var formdata = new FormData();
        formdata.append("adminID", cookie.adminID);
        formdata.append("userID", student.id);
        formdata.append("machineID", selectedOption.selectedOption.value);
        var config = {
        method: 'post',
        url: 'http://localhost:8080/admin/uncertify',
        headers: {
            'Authorization': `Basic ${cookie.authToken}`,
        },
        data : formdata
        };
        axios(config)
        .then(function (response) {
            alert(response.data);
            handleStudentSearch();
            renderStudent();
            renderPages();
            
        })
        .catch(function (error) {
        console.log(error);
        });
    }

    function handleDelete() {
        var formdata = new FormData();
        formdata.append("id", student.id);
        var config = {
        method: 'delete',
        url: 'http://localhost:8080/admin/delete/user',
        headers: {
            'Authorization': `Basic ${cookie.authToken}`,
        },
        data : formdata
        };
        axios(config)
        .then(function (response) {
            alert(response.data);
            localStorage.removeItem("student");
            renderStudent();
            renderPages();
        })
        .catch(function (error) {
        console.log(error);
        });
    }
    
    function viewMachines(){
        setShow(false);
        setCurrentPage("machines");
        localStorage.removeItem("student");
        renderStudent();
        renderPages();
    }
    
    function viewLogs(){
        setShow(false);
        setCurrentPage("logs");
        localStorage.removeItem("student");
        renderStudent();
        renderPages();
    }

    async function handleStudentSearch(){
        setLoading(true);
        if(localStorage.getItem("student") != null){
        var formdata = new FormData();
        formdata.append("id", JSON.parse(localStorage.getItem("student")).id);
        var config = {
            method: 'post',
            url: 'http://localhost:8080/admin/search/users',
            headers: { 
              'Authorization': `Basic ${cookie.authToken}`, 
            },
            data: formdata
          };
          
        await axios(config)
        .then(function (response) {
            const res = async () => {
                localStorage.setItem("student", JSON.stringify(response.data[0]))
                setLoading(false);
            }
            res();
        })
        .catch(function (error) {
            const err = async () => {
                console.log(error.stack);
            }
        err();
        });
        }
    }

    function viewSearch(){
        localStorage.removeItem("student");
        
        renderStudent();
        setCurrentPage("search");
        setShow(true);
    }
    
    function logout(){
        localStorage.removeItem("student");
        removeCookie("authToken", { path: '/'});
        removeCookie("studentID", { path: '/'});
        removeCookie("adminID", { path: '/'});
        navigate("/")
    }
    function renderPages(){
        if(currentPage === "machines"){
            return (
                <React.Fragment>
                    <div>
                        <MachineList></MachineList>
                    </div>
                    
                </React.Fragment>
            )
        }
        if(currentPage === "logs"){
            return (
                <React.Fragment>
                    <div>
                        <Logs></Logs>
                    </div>
                </React.Fragment>
            )
        }
        else{
            return (
                <React.Fragment>
                    <div></div>
                </React.Fragment>
            )
        }
    }

    function renderStudent(){
        //TODO: Delete User not Working
        if(localStorage.getItem("student") != null && isLoading == false){
            return(
                <React.Fragment>
                    <div>FIRST NAME: {student.first}</div>
                    <div>LAST NAME: {student.last}</div>
                    <div>USERNAME: {student.username}</div>
                    <div>GRADE: {student.grade}</div>
                    <div>CODE: {student.code}</div>
                    <div>ID: {student.id}</div>
                    {student.Machines.map(machine => (<div key={machine.id}>MACHINE ID: {String(machine.id)}</div>))}
                    <Select options={options} onChange={handleOptionChange} noOptionsMessage={() => "name not found"} />
                    <button onClick={() => {handleCertify()}}>Certify</button>
                    <button onClick={() => {handleUncertify()}}>Uncertify</button>
                    <button onClick={() => {handleDelete()}}>Delete User</button>
                </React.Fragment>
            )
        }
        else if(isLoading){
            return(
                <React.Fragment>
                    <div>
                        Loading...
                    </div>
                </React.Fragment>
            )
        }
        else{
            return(
                <React.Fragment>
                    <div></div>
                </React.Fragment>
            )
        }
    }
    
    return(
        <div>
            <button onClick={() => viewSearch()}>Search</button>
            <Search show={show} onHide={() => {setShow(false);}} />
            <button onClick={viewMachines}>Machines</button>
            <button onClick={viewLogs}>Logs</button>
            <button onClick={logout}>Logout</button>
            {renderPages()}
            {renderStudent()}
        </div>
    )
}
