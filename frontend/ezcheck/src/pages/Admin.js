import React from 'react'
import MachineList from '../components/MachineList';
import Search from '../components/Search.js';
import Logs from '../components/logs.js';
import Select from 'react-select';
import axios from "axios";
import '../styles/Admin.css';
import {useCookies } from 'react-cookie';
import { useNavigate } from "react-router-dom";
import Grid from '@mui/material/Grid';

export default function Admin() {
    const [show, setShow] = React.useState(false);
    const [cookie, removeCookie] = useCookies('user');
    const [currentPage, setCurrentPage] = React.useState("");
    const student = JSON.parse(localStorage.getItem("student"));
    const [options, setOptions] = React.useState([""]);
    const [selectedOption, setSelectedOption] = React.useState("");
    const [isLoading, setLoading] = React.useState(false);
    const navigate = useNavigate();
    // eslint-disable-next-line
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
        if(localStorage.getItem("student") != null && isLoading === false){
            return(
                <React.Fragment>
                    <div className="userInfo">
                        <h1 className="info">User Information</h1>
                        <div className="grid">
                            <Grid container rowSpacing={3} columnSpacing={{ xs: 1, sm: 2, md: 3 }}>
                                <Grid item xs={6}>
                                    <div className="tagValue">{student.first} {student.last}<br></br><div className="tag">Full Name</div></div>
                                </Grid>
                                <Grid item xs={6}>
                                    <div className="tagValue">{student.username}<br></br><div className="tag">Username</div></div>
                                </Grid>
                                <Grid item xs={6}>
                                    <div className="tagValue">{student.grade}th Grade<br></br><div className="tag">Grade</div></div>
                                </Grid>
                                <Grid item xs={6}>
                                    <div className="tagValue">{student.code}<br></br><div className="tag">Code</div></div>
                                </Grid>
                                </Grid>
                        <br></br>
                        <div className="tagValue">{student.id}<br></br><div className="tag">Student ID</div></div>
                        </div>
                        <h1 className="machineInfo">Machines Added: </h1>
                        {student.Machines.map(machine => (<div className="machine" key={machine.id}>Machine ID: {String(machine.id)}</div>))}
                        <Select className="searchMachine" options={options} onChange={handleOptionChange} noOptionsMessage={() => "name not found"} />
                        <br></br>
                        <button className="certify" onClick={() => {handleCertify()}}>Certify</button>
                        <button className="decertify" onClick={() => {handleUncertify()}}>Uncertify</button>
                        <button className="delete" onClick={() => {handleDelete()}}>Delete User</button>
                    </div>
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
            <button className="adminButton" onClick={() => viewSearch()}>Search</button>
            <Search className="searchModal" show={show} onHide={() => {setShow(false);}} />
            <button className="adminButton" onClick={viewMachines}>Machines</button>
            <button className="adminButton" onClick={viewLogs}>Logs</button>
            <button className="logout" onClick={logout}>Logout</button>
            {renderPages()}
            {renderStudent()}
        </div>
    )
}
