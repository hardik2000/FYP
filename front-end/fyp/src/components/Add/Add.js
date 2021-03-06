import React, { useState } from 'react';
import './Add.css';
import ReactDOM from 'react-dom'
import { render } from 'react-dom';
import { JSHash, CONSTANTS } from "react-native-hash";


async function addResources(credentials) {
    let token = JSON.parse(sessionStorage.getItem('token'));
    console.log("TOKEN ", token)
    return fetch('http://localhost:4000/channels/mychannel/chaincodes/basic', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${token}`
        },
        body: JSON.stringify(credentials)
    })
        .then(data => data.json())
}

export default function Add() {
    const [username, setUserName] = useState();
    const [hash, setHash] = useState();
    // const hash;
    // const [file, setFile] = useState()
    async function handleSubmit(e) {
        e.preventDefault();
        const res = await addResources({
            "fcn": "invoke",
            "peers": ["peer0.org1.example.com", "peer0.org2.example.com"],
            "chaincodeName": "basic",
            "channelName": "mychannel",
            "args": [username, hash]
        });
        console.log(res['result'])
        ReactDOM.render(
            <div className="App">
                {res['result']['message']}
                <br />
                ID : {res['result']['result']['id']}
                <br />
                HASH : {res['result']['result']['hash']}
            </div>,
            document.getElementById('response')
        );
    }

    async function getHash(e) {

        // const input = document.getElementById('file_input')
        const reader = new FileReader();
        const fileByteArray = [];

        reader.readAsArrayBuffer(e.target.files[0]);
        reader.onloadend = (evt) => {
            if (evt.target.readyState === FileReader.DONE) {
                const arrayBuffer = evt.target.result,
                    array = new Uint8Array(arrayBuffer);
                for (const a of array) {
                    fileByteArray.push(a);
                }
                console.log(fileByteArray.toString());

                JSHash(fileByteArray.toString(), CONSTANTS.HashAlgorithms.sha256)
                    .then(hash => setHash(hash), console.log(hash))
                    .catch(e => console.log(e));

                // setHash("fileByteArray.toString()")
            }
        }

    }
    return (
        <div className="login-wrapper">
            <h2>Inside Add function</h2>
            <form onSubmit={handleSubmit}>
                <label>
                    <p>
                        <span>ID&nbsp;&nbsp;&nbsp;&nbsp;</span>
                        <input type="text" onChange={e => setUserName(e.target.value)} />
                    </p>
                </label>
                <label>
                    <p>
                        <span>HASH&nbsp;&nbsp;&nbsp;&nbsp;</span>
                        <input type="file" id="file_input" accept="image/png, image/jpeg,application/pdf" onChange={getHash} />
                    </p>
                </label>
                <div><h1> </h1></div>
                <div>
                    <button type="submit" >Submit</button>
                </div>
            </form>
            <br />
            <br />
            <div id="response"></div>
        </div >
    )
}

render(<Add />, document.getElementById('root'));

