import React from 'react'
import { useCookies, removeCookie} from 'react-cookie';
import { useNavigate } from "react-router-dom";

export default function Student() {
    const [cookie, setCookie, removeCookie] = useCookies('user');
    const student = JSON.parse(localStorage.getItem("student"));
    const [isLoading, setLoading] = React.useState(true);
    const navigate = useNavigate();
    
    function logout(){
        localStorage.removeItem("student");
        removeCookie("authToken", { path: '/'});
        removeCookie("studentID", { path: '/'});
        removeCookie("adminID", { path: '/'});
        navigate("/")
    }

    function renderStudent(){
        //TODO: not working on initial login
        if(localStorage.getItem("student") != null){
            return(
                <React.Fragment>
                    <div>{student.first}</div>
                    <div>{student.last}</div>
                    <div>{student.username}</div>
                    <div>{student.grade}</div>
                    <div>{student.code}</div>
                    {student.Machines.map(machine => (<div key={machine.id}>MACHINE ID: {String(machine.id)}</div>))}
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

            <button onClick={logout}>Logout</button>

            {renderStudent()}
        </div>
    )
}
