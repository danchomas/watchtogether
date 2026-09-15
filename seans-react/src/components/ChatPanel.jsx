import { useState } from "react";
import ChatMessage from "./ChatMessage.jsx";

// Стартовая история чата — та же переписка, что была захардкожена
const initialMessages = [
  { id: 1, type: "system", text: "Игорь присоединился к комнате" },
  { id: 2, type: "message", name: "Аня", time: "20:14", text: "так, все на месте? запускаю через 10 секунд" },
  { id: 3, type: "message", name: "Игорь", time: "20:14", text: "готов, звук проверил" },
  { id: 4, type: "system", text: "Аня поставила видео на паузу — 41:02" },
  { id: 5, type: "message", name: "Марта", time: "20:53", text: "подождите, а это точно тот самый брат из первой серии?" },
  { id: 6, type: "message", name: "Аня", time: "20:53", text: "да! я тоже не сразу узнала" },
  { id: 7, type: "system", text: "Аня возобновила воспроизведение" },
];

export default function ChatPanel() {
  const [messages, setMessages] = useState(initialMessages);
  const [draft, setDraft] = useState("");

  function handleSend() {
    const text = draft.trim();
    if (!text) return;

    const time = new Date().toLocaleTimeString("ru-RU", {
      hour: "2-digit",
      minute: "2-digit",
    });

    setMessages((prev) => [
      ...prev,
      { id: prev.length + 1, type: "message", name: "Вы", time, text },
    ]);
    setDraft("");
  }

  function handleKeyDown(event) {
    if (event.key === "Enter") handleSend();
  }

  return (
    <aside className="chat-col">
      <div className="chat-head">Чат комнаты</div>

      <div className="chat-messages">
        {messages.map((message) => (
          <ChatMessage key={message.id} {...message} />
        ))}
      </div>

      <div className="chat-input">
        <input
          type="text"
          placeholder="Написать в чат…"
          value={draft}
          onChange={(event) => setDraft(event.target.value)}
          onKeyDown={handleKeyDown}
        />
        <button className="btn btn-primary" onClick={handleSend}>
          Отправить
        </button>
      </div>
    </aside>
  );
}
