import { Link } from "react-router-dom";
import AvatarStack from "./AvatarStack.jsx";

export default function RoomHeader({ title, code }) {
  return (
    <header className="room-header">
      <div className="wrap">
        <div className="room-title-block">
          <h1>{title}</h1>
          <span className="room-code">{code}</span>
        </div>
        <div className="room-people">
          <AvatarStack />
          <Link to="/" className="btn btn-ghost">Выйти</Link>
        </div>
      </div>
    </header>
  );
}
