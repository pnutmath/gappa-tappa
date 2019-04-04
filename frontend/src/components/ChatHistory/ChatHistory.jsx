import React, { Component } from "react"
import "./ChatHistory.scss";
import Message from '../Message'

class ChatHistory extends Component {
    render() {
        const messages = this.props.chatHistory.map((msg, index) => (
            <Message key={index} message={msg.data}/>  ));

        return (
            <div className="chatHistory">
                <h3>Chat History</h3>
                {messages}
            </div>
        );
    }
}

export default ChatHistory;