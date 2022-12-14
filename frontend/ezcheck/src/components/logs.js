import React from 'react'
import axios from "axios";

function Logs() {
  const [logDisplay, setLogDisplay] = React.useState("")
    React.useEffect(() => {getLogs()}, [])
    function getLogs(){
        var config = {
        method: 'get',
        url: 'http://localhost:8080/log',
        headers: { }
        };

        axios(config)
        .then(function (response) {
        setLogDisplay(JSON.stringify(response.data));
        })
        .catch(function (error) {
        console.log(error);
        });

    }
  return (
    <>
    <div>logs</div>
    <div>
      {logDisplay}
    </div>
    </>
  )
}

export default Logs;