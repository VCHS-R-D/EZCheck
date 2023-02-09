import React from 'react'
// eslint-disable-next-line
import {Cookies, useCookies} from 'react-cookie';
import axios from "axios";
import "../styles/Machine.css";

function MachineList(){
  const [machineList, setMachineList] = React.useState([]);
  const [machineID, setMachineID] = React.useState("");
  const delMachine = React.useRef("");
  const [cookie] = useCookies('user');

  React.useEffect(() => {getMachines()}, [])

  function getMachines() {
    var config = {
      method: 'get',
      url: 'http://localhost:8080/machines',
      headers: { }
    };
    
    axios(config)
    .then(function (response) {
      setMachineList(response.data);  
    })
    .catch(function (error) {
      console.log(error);
    });
  }

  function createMachine() {
    var formdata = new FormData();
    formdata.append("machineID", machineID);
    var config = {
    method: 'post',
    url: 'http://localhost:8080/admin/machines/create',
    headers: {
        'Authorization': `Basic ${cookie.authToken}`,
    },
    data : formdata
    };
    axios(config)
    .then(function (response) {
        alert(response.data);
        getMachines();        
    })
    .catch(function (error) {
      console.log(error);
    });
    }

    function deleteMachine() {
      var formdata = new FormData();
      formdata.append("machineID", String(delMachine.current));
      var config = {
      method: 'delete',
      url: 'http://localhost:8080/admin/machines/delete',
      headers: {
          'Authorization': `Basic ${cookie.authToken}`,
      },
      data : formdata
      };
      axios(config)
      .then(function (response) {
        alert(response.data);
        getMachines();        
      })
      .catch(function (error) {
        console.log(error);
      });
      }

    return( 
      <>
      <div>
      <form>
        <input className="input" placeholder="Machine ID" onChange={(event) => {setMachineID(event.target.value)}}></input>
      </form>
      <button className="create" onClick={() => {createMachine()}}>Create Machine</button>
      {machineList.map(machine => (<div className="machineItem" key={machine.id}>{machine.id} in-use: {<span style={String(machine.in_use) === "true" ? {backgroundColor: "#b5ffc0", color:"#00ff26"} : {backgroundColor: "#ffb8b8", color:"#ff0000"}}className="inUse">{String(machine.in_use)}</span>} actions: ({machine.actions}) <button className="delete" key={machine.id} onClick={() => {delMachine.current = String(machine.id); deleteMachine();}}>Delete</button></div>))}
      </div>
    </>
    )
  }

export default MachineList;