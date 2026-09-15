// Системные записи ("Игорь присоединился к комнате") и обычные реплики
// выглядят по-разному, поэтому компонент сам выбирает разметку по типу.
export default function ChatMessage({ type, name, time, text }) {
  if (type === "system") {
    return <div className="msg msg-system">{text}</div>;
  }

  return (
    <div className="msg">
      <div className="msg-meta">
        <span className="msg-name">{name}</span>
        <span className="msg-time">{time}</span>
      </div>
      <div className="msg-body">{text}</div>
    </div>
  );
}
