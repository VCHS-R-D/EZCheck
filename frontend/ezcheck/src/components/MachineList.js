import React from 'react'
import axios from "axios";

function MachineList(){
  
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
    })
    .catch(function (error) {
      console.log(error);
    });
    
  }
    return( <>
      <div>
        test
      </div>
    </>
  )
    }

export default MachineList;