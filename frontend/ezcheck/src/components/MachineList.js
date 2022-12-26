import React from 'react'
import axios from "axios";

function MachineList(){
  const [machineList, setMachineList] = React.useState([]);
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
    return( <>
      <div>
      {machineList.map(machine => (<div key={machine.id}>{machine.name} in-use: {String(machine.in_use)} actions: ({machine.actions})</div>))}
      </div>
    </>
  )
    }

export default MachineList;