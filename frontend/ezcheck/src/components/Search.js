import React, { useEffect } from 'react'
import axios from "axios";
import {Cookies, useCookies} from 'react-cookie';
import { Modal, Button } from "react-bootstrap";

var FormData = require('form-data');

function Search(props) {
    const [cookie, setCookie] = useCookies('user');
    const [studentDict, setStudentDict] = React.useState([]);
    const [show, setShow] = React.useState(props.show);
    const [studentID, setStudentID] = React.useState("");

    useEffect(() => {handleSearch()}, [])

    function handleSelectStudent(id) {
        props.onHide();
        console.log(id);
        setCookie('studentID', id, { path: '/'});
    }
    function handleSearch(){
        console.log(cookie.authToken);
        var config = {
        method: 'post',
        url: 'http://localhost:8080/admin/search/users',
        headers: { 
            'Authorization': `Basic ${cookie.authToken}`
        }
        };

        axios(config)
        .then(function (response) {
        console.log(JSON.parse(("[" + response.data.match(/[^[\]]+(?=])/g)[0]) + "]"[0]));
        setStudentDict(JSON.parse(("[" + response.data.match(/[^[\]]+(?=])/g)[0]) + "]"[0]));

        })
        .catch(function (error) {
        console.log(error);

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
            {studentDict.map(student => (<button key={student.id} onClick={() => handleSelectStudent(student.id)}>{student.first} {student.last} ({student.grade})</button>))}
        </Modal.Body>
        </Modal>
    )
}

export default Search;