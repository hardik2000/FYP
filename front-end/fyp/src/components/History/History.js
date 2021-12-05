

import React, { useState } from 'react';
import './History.css';
import ReactDOM from 'react-dom'
import { render } from 'react-dom';


async function HistoryResources(credentials) {
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

export default function History() {
    const [username, setUserName] = useState();

    async function handleSubmit(e) {
        e.preventDefault();
        const res = await HistoryResources({
            "fcn": "history",
            "peers": ["peer0.org1.example.com", "peer0.org2.example.com"],
            "chaincodeName": "basic",
            "channelName": "mychannel",
            "args": [username]
        });
        if (!res['result']['result']) {
            const element = "\n\nHISTORY DOES NOT EXIST";
            ReactDOM.render(element, document.getElementById('response'));
            return;

        }
        ReactDOM.render(
            <div className="App">
                {res['result']['message']}
                <br />
                <table cellpadding="2" cellspacing="2" border="2">
                    <tr>
                        <th>Transaction ID</th>
                        <th>ID</th>
                        <th>Hash</th>

                    </tr>
                    {res['result']['result'].map((value, index) => {
                        return <tr><td>{value['txId']}</td><td>{value['value']['id']}</td><td>{value['value']['hash']}</td></tr>
                    })}
                </table>
                < br />
                <br />
            </div >,
            document.getElementById('response')
        );

    }

    return (
        <div className="login-wrapper">
            <h2>Inside History function</h2>
            <form onSubmit={handleSubmit}>
                <label>
                    <p>
                        <span>ID&nbsp;&nbsp;&nbsp;&nbsp;</span>
                        <input type="text" onChange={e => setUserName(e.target.value)} />
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
        </div>
    )
}

render(<History />, document.getElementById('root'));

