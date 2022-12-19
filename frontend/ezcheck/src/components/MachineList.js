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
      console.log(JSON.stringify(response.data));
      setMachineList(response.data);
    })
    .catch(function (error) {
      console.log(error);
    });
    
  }
    return( <>
      <div>
      {machineList.map(machine => (<h1 key={machine.id}>{machine.name} {String(machine.in_use)} ({machine.actions})</h1>))}
      </div>
    </>
  )
    }

export default MachineList;