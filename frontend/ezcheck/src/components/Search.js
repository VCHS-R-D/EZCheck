import React, { useEffect } from 'react'
import axios from "axios";
import {Cookies, useCookies} from 'react-cookie';
import { Modal, Button } from "react-bootstrap";

function Search(props) {
    const [cookie] = useCookies('user');
    const [studentDict, setStudentDict] = React.useState([]);

    useEffect(() => {handleSearch()}, [])

    function handleSelectStudent(student) {
        props.onHide();
        localStorage.setItem("student", JSON.stringify(student));
    }
    
    async function handleSearch(){
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