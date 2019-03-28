import React, { Component } from "react"
import "./ChatHistory.scss";

class ChatHistory extends Component {
    render() {
        const messages = this.props.chatHistory.map((msg, index) => (
            <p key={index}>{msg.data}</p>
        ));

        return (
            <div className="chatHistory">
                <h3>Chat History</h3>
                {messages}
            </div>
        );
    }
}

export default ChatHistory;