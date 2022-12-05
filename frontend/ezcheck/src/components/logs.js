import React from 'react'

function logs() {
    function getLogs(){
        var axios = require('axios');

        var config = {
        method: 'get',
        url: 'localhost:8080/log',
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
  return (
    <div>logs</div>
  )
}

export default logs;