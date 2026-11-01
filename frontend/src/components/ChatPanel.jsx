import { useEffect, useRef, useState } from "react";
import ChatMessage from "./ChatMessage.jsx";

export default function ChatPanel({ messages, onSend }) {
  const [draft, setDraft] = useState("");
  const [sending, setSending] = useState(false);
  const history = useRef(null);
  const latestMessage = messages.at(-1)?.id;

  useEffect(() => {
    const panel = history.current;
    if (panel) panel.scrollTop = panel.scrollHeight;
  }, [latestMessage]);

  async function submit(event) {
    event.preventDefault();
    const text = draft.trim();
    if (!text || sending) return;
    setSending(true);
    try {
      if (await onSend(text)) setDraft((current) => current.trim() === text ? "" : current);
    } finally { setSending(false); }
  }

  return <aside className="chat-col">
    <div className="chat-head">Чат комнаты</div>
    <div ref={history} className="chat-messages" aria-live="polite">
      {messages.map((message) => <ChatMessage key={message.id} name={message.name}
        text={message.text} time={new Date(message.time).toLocaleTimeString("ru-RU", { hour: "2-digit", minute: "2-digit" })} />)}
    </div>
    <form className="chat-input" onSubmit={submit}>
      <input type="text" aria-label="Сообщение" maxLength={500} placeholder="Написать в чат…" value={draft} onChange={(event) => setDraft(event.target.value)} />
      <button className="btn btn-primary" type="submit" disabled={sending || !draft.trim()}>Отправить</button>
    </form>
  </aside>;
}
