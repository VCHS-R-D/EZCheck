import React, { useEffect } from 'react'
import axios from "axios";
import {Cookies, useCookies} from 'react-cookie';
import { Modal, Button } from "react-bootstrap";

var FormData = require('form-data');


function Search(props) {
    const [cookie, setCookie, removeCookie] = useCookies('user');
    const [studentDict, setStudentDict] = React.useState([]);
    const [show, setShow] = React.useState(props.show);
    const [studentID, setStudentID] = React.useState("");

    useEffect(() => {handleSearch()}, [])

    function handleSelectStudent(student) {
        props.onHide();
        localStorage.setItem("student", JSON.stringify(student));
        console.log(student);
    }
    
    async function handleSearch(){
        console.log(cookie.authToken);
        var config = {
            method: 'post',
            url: 'http://localhost:8080/admin/search/users',
            headers: { 
              'Authorization': `Basic ${cookie.authToken}`, 
            }
          };
          
        await axios(config)
        .then(function (response) {
            const res = async () => {
                setStudentDict(response.data);
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
    return (
        <Modal
      {...props}
      size="lg"
      aria-labelledby="contained-modal-title-vcenter"
      centered
        >

        <Modal.Body>
            {studentDict.map(student => (<button key={student.id} onClick={() => handleSelectStudent(student)}>{student.first} {student.last} ({student.grade}) </button>))}
        </Modal.Body>
        </Modal>
    )
}

export default Search;