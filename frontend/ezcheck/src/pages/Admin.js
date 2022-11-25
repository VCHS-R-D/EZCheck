import React from 'react'
import Modal from 'react-modal';
import MachineList from '../components/MachineList';
import Search from '../components/Search';

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
    const [modalIsOpen, setIsOpen] = React.useState(false);
    const [currentPage, setCurrentPage] = React.useState("");

    function openModal() {
        setIsOpen(true);
    }

    function closeModal() {
        setIsOpen(false);
    }

    function search(){
        openModal();
    }

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
            <button onClick={search}>Search</button>
            <button onClick={viewMachines}>Machines</button>
            <button>Logs</button>
            <Modal
                isOpen={modalIsOpen}
                onRequestClose={closeModal}
                style={customStyles}
                contentLabel="Example Modal"
            >
                <button onClick={closeModal}>close</button>
                <h1><Search/></h1>
            </Modal>
            {renderPages()}
        </div>
    )
}
