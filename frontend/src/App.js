import React, { Component } from 'react';
import logo from './logo.svg';
import './App.css';
import { connect, sendMsg } from "./api"

class App extends Component {

  constructor(props) {
    super(props);
    connect();
  }

  send() {
    console.log('Hello');
    sendMsg('Hello');
  }

  render() {
    return (
      <div className="App">
        <header className="App-header">
          <img src={logo} className="App-logo" alt="logo" />

          <button onClick={this.send} className="btn">Send Message</button>
        </header>
      </div>
    );
  }
}

export default App;
