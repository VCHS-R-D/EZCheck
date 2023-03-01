import React from 'react'
import { useCookies } from 'react-cookie';
import { useNavigate } from "react-router-dom";
import "../styles/Student.css"
import Grid from '@mui/material/Grid'

export default function Student() {
    // eslint-disable-next-line
    const [cookie, setCookie,removeCookie] = useCookies('user');
    const student = JSON.parse(localStorage.getItem("student"));
    const navigate = useNavigate();
    
    function logout(){
        localStorage.removeItem("student");
        removeCookie("authToken", { path: '/'});
        removeCookie("studentID", { path: '/'});
        removeCookie("adminID", { path: '/'});
        navigate("/")
    }

    function renderStudent(){
        if(localStorage.getItem("student") != null){
            return(
                <React.Fragment>
                    <React.Fragment>
                        <br></br><br></br>
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
                        </div>
                </React.Fragment>
                </React.Fragment>
            )
        }
        else{
            return(
                <React.Fragment>
                    <div>N/A</div>
                </React.Fragment>
            )
        }
    }
    
    return(
        <div>
            <button className="logout" onClick={logout}>Logout</button>
            {renderStudent()}
        </div>
    )
}
