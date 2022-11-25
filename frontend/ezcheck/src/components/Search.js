import React, { useEffect } from 'react'
import axios from "axios";
import {Cookies, useCookies} from 'react-cookie';
var FormData = require('form-data');

export default function Search() {
    const [cookie, setCookie] = useCookies('user');

    useEffect(() => {handleSearch();})
    function handleSearch(){
        console.log(cookie.authToken);
        var config = {
        method: 'post',
        url: 'http://localhost:8080/admin/search',
        headers: { 
            'Authorization': `Basic ${cookie.authToken}`
        }
        };

        axios(config)
        .then(function (response) {
        console.log(response.data);

        })
        .catch(function (error) {
        console.log(error);

        });
    }
    return (
        <div>
            hi
        </div>
    )
}

