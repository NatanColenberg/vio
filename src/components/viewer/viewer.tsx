import React, { useState, useEffect } from "react";
import { Button } from "@material-ui/core";
import { NavigateNext, NavigateBefore } from "@material-ui/icons";

import "./viewer.css";

export default function Viewer() {
  const [socket, setSocket]: any = useState(null);
  const [img, setImg] = useState("");
  useEffect(() => {
    wsConnect();
    return () => socket.Close();
  }, []);

  const wsConnect = () => {
    const ws: WebSocket = new WebSocket("ws://localhost:8080/ws/viewer");
    setSocket(ws);
    console.log("Attempting Connection...");

    ws.onopen = () => {
      console.log("Successfully Connected!");
      var msg = {
        Type: "getStudyImage",
        Data: "",
      };
      ws.send(JSON.stringify(msg));
    };

    ws.onclose = () => {
      console.log("Connection Closed");
    };

    ws.onerror = (err) => {
      console.log("Socket Error: ", err);
    };

    ws.onmessage = (mesEvt) => {
      console.log(mesEvt);

      const msg: { Type: string; Data: string } = JSON.parse(mesEvt.data);

      switch (msg.Type) {
        case "selectedStudyChanged":
          setImg(msg.Data);
          break;
        default:
          break;
      }
    };
  };

  return (
    <div className="viewer-wrap">
      <h1 className="viewer-title">Viewer</h1>
      {img ? (
        <div className="viewer-cont">
          <div className="viewer-controlButtons">
            <Button
              className="login-navControlButton"
              variant="contained"
              startIcon={<NavigateBefore />}
              onClick={() => {
                var msg = {
                  Type: "prevSelectedStudy",
                  Data: "",
                };
                socket.send(JSON.stringify(msg));
              }}
            >
              Prev
            </Button>

            <Button
              className="login-navControlButton"
              variant="contained"
              endIcon={<NavigateNext />}
              onClick={() => {
                var msg = {
                  Type: "nextSelectedStudy",
                  Data: "",
                };
                socket.send(JSON.stringify(msg));
              }}
            >
              Next
            </Button>
          </div>
          <img
            src={`data:image/jpg;base64,${img}`}
            alt="Red dot"
            width="500px"
          ></img>
        </div>
      ) : null}
    </div>
  );
}
