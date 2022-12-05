import React from 'react'
import Modal from 'react-modal';
import MachineList from '../components/MachineList';
import Search from '../components/Search.js';

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


export default function Admin() {
    const [show, setShow] = React.useState(false);
    const [currentPage, setCurrentPage] = React.useState("");
    const student = JSON.parse(localStorage.getItem("student"));
    function viewMachines(){
        setCurrentPage("machines");
        renderPages();
    }
    
    function viewSearch(){
        setCurrentPage("search");
        setShow(true);
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
        else{
            return (
                <React.Fragment>
                    <div>a</div>
                </React.Fragment>
            )
        }
    }

    function renderStudent(){
        if(localStorage.getItem("student") != null){
            return(
                <React.Fragment>
                    <div>{student.first}</div>
                    <div>{student.last}</div>
                    <div>{student.username}</div>
                    <div>{student.grade}</div>
                    <div>{student.code}</div>
                </React.Fragment>
            )
        }
        else{
            return(
                <React.Fragment>
                    <div>b</div>
                </React.Fragment>
            )
        }
    }
    
    return(
        <div>
            <button onClick={() => viewSearch()}>Search</button>
            <Search show={show} onHide={() => setShow(false)} />
            <button onClick={viewMachines}>Machines</button>
            <button>Logs</button>
            
            {renderPages()}
            {/* {renderStudent()} */}
        </div>
    )
}
