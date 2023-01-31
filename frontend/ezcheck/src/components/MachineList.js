import React from 'react'
import {Cookies, useCookies} from 'react-cookie';
import axios from "axios";

function MachineList(){
  const [machineList, setMachineList] = React.useState([]);
  const [machineID, setMachineID] = React.useState("");
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
      formdata.append("machineID", machineID);
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
        <input placeholder="Machine ID" onChange={(event) => {setMachineID(event.target.value)}}></input>
      </form>
      <button onClick={() => {createMachine()}}>Create Machine</button>
      {machineList.map(machine => (<div key={machine.id}>{machine.id} in-use: {String(machine.in_use)} actions: ({machine.actions}) <button key={machine.id} onClick={() => {setMachineID(machine.id); console.log(machine.id); deleteMachine();}}>DELETE</button></div>))}
      </div>
    </>
    )
  }

export default MachineList;