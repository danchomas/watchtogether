import { Link } from "react-router-dom";

export default function RoomHeader({ title, code, participants }) {
  return (
    <header className="room-header">
      <div className="wrap">
        <div className="room-title-block">
          <h1>{title}</h1>
          <button className="room-code" type="button" title="Скопировать ссылку" onClick={() => navigator.clipboard.writeText(`${location.origin}/watch/${code}`)}>{code} · копировать ссылку</button>
        </div>
        <div className="room-people">
          <span>{participants} зр.</span>
          <Link to="/" className="btn btn-ghost">Выйти</Link>
        </div>
      </div>
    </header>
  );
}
