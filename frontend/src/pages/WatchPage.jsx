import { useEffect, useState } from "react";
import { Link, useParams } from "react-router-dom";
import { api, roomSession } from "../api.js";
import RoomHeader from "../components/RoomHeader.jsx";
import VideoPlayer from "../components/VideoPlayer.jsx";
import ChatPanel from "../components/ChatPanel.jsx";

export default function WatchPage() {
  const { code } = useParams();
  const session = roomSession(code);
  const [room, setRoom] = useState(null);
  const [error, setError] = useState("");
  const [actionError, setActionError] = useState("");
  const [connected, setConnected] = useState(false);
  const [guestName, setGuestName] = useState("");
  const [joining, setJoining] = useState(false);

  useEffect(() => {
    setRoom(null);
    setError("");
    setConnected(false);
    if (!session?.token) return;
    let active = true;
    let stream;
    api(`/rooms/${encodeURIComponent(code)}`, {}, session.token)
      .then((data) => {
        if (!active) return;
        setRoom(data);
        stream = new EventSource(`/api/rooms/${encodeURIComponent(code)}/events?token=${encodeURIComponent(session.token)}`);
        stream.onmessage = (event) => {
          if (!active) return;
          const data = JSON.parse(event.data);
          if (data.type === "state") { setRoom(data.snapshot); setConnected(true); }
        };
        stream.onerror = () => { if (active) setConnected(false); };
      })
      .catch((err) => { if (active) setError(err.message); });
    return () => { active = false; stream?.close(); };
  }, [code, session?.token]);

  if (!session?.token) return <form className="room-form room-notice" onSubmit={async (event) => {
    event.preventDefault(); setJoining(true); setError("");
    try {
      const result = await api(`/rooms/${encodeURIComponent(code)}/join`, { method: "POST", body: JSON.stringify({ name: guestName }) });
      sessionStorage.setItem(`room:${result.room.code}`, JSON.stringify({ token: result.token, name: guestName.trim() }));
      window.location.reload();
    } catch (err) { setError(err.message); setJoining(false); }
  }}>
    <h1>Войти в комнату {code.toUpperCase()}</h1>
    <label>Ваше имя<input required maxLength={40} value={guestName} onChange={(e) => setGuestName(e.target.value)} /></label>
    {error && <p className="form-error" role="alert">{error}</p>}
    <button className="btn btn-primary" disabled={joining}>Присоединиться</button>
    <Link to="/">На главную</Link>
  </form>;
  if (error) return <div className="room-notice" role="alert">{error} <Link to="/">На главную</Link></div>;
  if (!room) return <div className="room-notice">Загружаем комнату…</div>;

  async function send(path, payload) {
    try { await api(`/rooms/${encodeURIComponent(code)}/${path}`, { method: "POST", body: JSON.stringify(payload) }, session.token); setActionError(""); return true; }
    catch (err) { setActionError(err.message); return false; }
  }

  return <>
    <RoomHeader title={room.title} code={room.code} participants={room.participants} />
    {!connected && <div className="connection-banner">Связь с комнатой прервана. Повторное подключение…</div>}
    {actionError && <div className="connection-banner" role="alert">{actionError}</div>}
    <main className="room-layout">
      <VideoPlayer room={room} onPlayback={(position, playing) => send("playback", { position, playing })} />
      <ChatPanel messages={room.messages} onSend={(text) => send("messages", { text })} />
    </main>
  </>;
}
