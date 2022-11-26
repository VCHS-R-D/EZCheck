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
    function viewMachines(){
        setCurrentPage("machines");
        renderPages();
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
    
    return(
        <div>
            <button onClick={() => setShow(true)}>Search</button>
            <Search show={show} onHide={() => setShow(false)} />
            <button onClick={viewMachines}>Machines</button>
            <button>Logs</button>
            
            {renderPages()}
        </div>
    )
}
