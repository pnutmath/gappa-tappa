import React, { Component } from 'react';
import './App.scss';
import { connect, sendMsg } from "./api"
import Header from "./components/Header"
import ChatHistory from './components/ChatHistory';
import ChatInput from './components/ChatInput';

class App extends Component {

  constructor(props) {
    super(props);
    this.state = {
      chatHistory: []
    };
  }

  componentDidMount() {
    connect((msg) => {
      console.log("New Message");
      this.setState(prevState => ({
        chatHistory: [...this.state.chatHistory, msg]
      }))
      console.log(this.state);
    });
  }

  send(event) {
    if (event.keyCode === 13) {
      sendMsg(event.target.value);
      event.target.value = "";
    }
  }

  render() {
    return (
      <div className="App">
        <Header />
        <header className="App-header">
          <button onClick={this.send} className="btn">Send Message</button>
        </header>
        <ChatHistory chatHistory={this.state.chatHistory} />
        <ChatInput send={this.send} />
      </div>
    );
  }
}

export default App;
