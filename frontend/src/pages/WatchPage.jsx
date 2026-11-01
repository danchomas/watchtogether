import { useEffect, useState } from "react";
import { Link, useParams } from "react-router-dom";
import { api, roomSession } from "../api.js";
import RoomHeader from "../components/RoomHeader.jsx";

export default function WatchPage() {
  const { code } = useParams();
  const session = roomSession(code);
  const [room, setRoom] = useState(null);
  const [error, setError] = useState("");
  const [guestName, setGuestName] = useState("");
  const [joining, setJoining] = useState(false);

  useEffect(() => {
    setRoom(null);
    setError("");
    if (!session?.token) return;
    let active = true;
    api(`/rooms/${encodeURIComponent(code)}`, {}, session.token)
      .then((data) => { if (active) setRoom(data); })
      .catch((err) => { if (active) setError(err.message); });
    return () => { active = false; };
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

  return <>
    <RoomHeader title={room.title} code={room.code} participants={room.participants} />
    <main className="room-notice">Комната готова. Пригласите друзей по ссылке из шапки.</main>
  </>;
}
