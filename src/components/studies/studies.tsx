import React, { useState, useEffect } from "react";
import StudiesTAble from "./studiesTable";

import StudyModal from "../../modals/studyModal";

import "./studies.css";

export default function Studies() {
  const [socket, setSocket]: any = useState(null);
  const [studies, setStudies]: any = useState([]);
  const openViewerWindow = () => {
    console.log(window);

    // const appAddr = window.location.href;
    const windowHeight = window.screen.availHeight;
    const windowWidth = window.screen.availWidth;
    const win = window.open(
      "/viewer",
      "",
      `channelmode=yes,height=${windowHeight + 100},width=${windowWidth + 100}`
    );
  };

  const wsConnect = () => {
    const ws: WebSocket = new WebSocket("ws://localhost:8080/ws/studyList");
    setSocket(ws);
    // socket = new WebSocket("ws://localhost:8080/ws");
    console.log("Attempting Connection...");

    ws.onopen = () => {
      console.log("Successfully Connected!");
      var msg = {
        Type: "getStudies",
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
      const msg: { Type: string; Data: StudyModal[] } = JSON.parse(mesEvt.data);
      switch (msg.Type) {
        case "studyListUpdate":
          setStudies(msg.Data);
          break;
        default:
          break;
      }
    };
  };

  const rowClicked = (row: any) => {
    // console.dir(row);
    var msg = {
      Type: "selectedStudyChanged",
      Data: row.accession,
    };
    socket.send(JSON.stringify(msg));
  };

  useEffect(() => {
    wsConnect();
    openViewerWindow();

    return () => socket.Close();
  }, []);

  return (
    <div className="studies-wrap">
      <h1 className="studies-title">Study List</h1>
      <StudiesTAble rows={studies} onClick={rowClicked} />
    </div>
  );
}
