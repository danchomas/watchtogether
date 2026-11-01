import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { api } from "../api.js";
import SiteHeader from "../components/SiteHeader.jsx";
import SiteFooter from "../components/SiteFooter.jsx";
import AvatarStack from "../components/AvatarStack.jsx";

// Контент шагов и возможностей вынесен в данные — так разметка
// в JSX не разрастается и её проще редактировать или расширять.
const steps = [
  {
    num: "01",
    title: "Создайте комнату",
    text: "Вставьте прямую ссылку на видео — Watchroom подготовит плеер и сгенерирует код комнаты.",
  },
  {
    num: "02",
    title: "Пригласите друзей",
    text: "Отправьте код или ссылку. Гости присоединяются мгновенно",
  },
  {
    num: "03",
    title: "Смотрите синхронно",
    text: "Пауза и перемотка одного участника применяются у всех. Рядом — чат для реплик по ходу просмотра.",
  },
];

const features = [
  {
    tag: "ядро сервиса",
    title: "Общее управление просмотром",
    text: "Пауза и перемотка одного зрителя применяются у всех. Плеер периодически корректирует расхождения из-за задержек сети.",
    large: true,
  },
  {
    tag: "общение",
    title: "Чат рядом с видео",
    text: "Реакции и реплики не перекрывают экран и не теряются в перемотке.",
  },
  {
    tag: "источники",
    title: "Прямые ссылки на видео",
    text: "Укажите общедоступный видеофайл, который открывается в браузере каждого участника.",
  },
];

export default function HomePage() {
  const navigate = useNavigate();
  const [mode, setMode] = useState("create");
  const [name, setName] = useState("");
  const [title, setTitle] = useState("");
  const [videoUrl, setVideoUrl] = useState("");
  const [code, setCode] = useState("");
  const [error, setError] = useState("");
  const [busy, setBusy] = useState(false);

  async function submit(event) {
    event.preventDefault();
    setBusy(true);
    setError("");
    try {
      const result = mode === "create"
        ? await api("/rooms", { method: "POST", body: JSON.stringify({ name, title, videoUrl }) })
        : await api(`/rooms/${encodeURIComponent(code.trim().toUpperCase())}/join`, { method: "POST", body: JSON.stringify({ name }) });
      sessionStorage.setItem(`room:${result.room.code}`, JSON.stringify({ token: result.token, name: name.trim() }));
      navigate(`/watch/${result.room.code}`);
    } catch (err) { setError(err.message); }
    finally { setBusy(false); }
  }
  return (
    <>
      <SiteHeader />

      <section className="hero">
        <div className="wrap">
          <div>
            <div className="timecode">
              <span className="dot"></span>
              00:42:17 — смотрим вместе
            </div>
            <h1>
              Разные устройства,
              <br />
              Одно время.
            </h1>
            <p className="hero-lede">
              Watchroom синхронизирует воспроизведение видео между всеми, кто в
              комнате — пауза и перемотка общие для всех зрителей.
              Рядом — чат, чтобы обсуждать происходящее в
              реальном времени.
            </p>
            <div className="hero-actions">
              <button type="button" className={`btn ${mode === "create" ? "btn-primary" : "btn-ghost"}`} onClick={() => setMode("create")}>Создать комнату</button>
              <button type="button" className={`btn ${mode === "join" ? "btn-primary" : "btn-ghost"}`} onClick={() => setMode("join")}>Войти по коду</button>
            </div>
            <form id="start" className="room-form" onSubmit={submit}>
              <label>Ваше имя<input required maxLength={40} value={name} onChange={(e) => setName(e.target.value)} placeholder="Как к вам обращаться" /></label>
              {mode === "create" ? <>
                <label>Название комнаты<input required maxLength={100} value={title} onChange={(e) => setTitle(e.target.value)} placeholder="Пятничный киновечер" /></label>
                <label>Прямая ссылка на видео<input required type="url" value={videoUrl} onChange={(e) => setVideoUrl(e.target.value)} placeholder="https://example.com/movie.mp4" /></label>
              </> : <label>Код комнаты<input required value={code} onChange={(e) => setCode(e.target.value)} placeholder="AX7F92" /></label>}
              {error && <p className="form-error" role="alert">{error}</p>}
              <button className="btn btn-primary" disabled={busy}>{busy ? "Подождите…" : mode === "create" ? "Начать просмотр" : "Присоединиться"}</button>
              {mode === "create" && <p className="hero-note">Поддерживаются прямые ссылки на видео, доступные браузеру (например MP4). Ссылки на страницы видеохостингов не воспроизводятся.</p>}
            </form>
            <p className="hero-note">
              Комната открывается по коду или ссылке-приглашению.
            </p>
          </div>

          <div className="hero-visual">
            <div className="hero-screen">
              <div className="hero-screen-bars">
                <span></span>
                <span></span>
                <span></span>
              </div>
              <div className="hero-screen-frame"></div>
            </div>
            <div className="hero-viewers">
              <AvatarStack />
              3 зрителя смотрят прямо сейчас
            </div>
          </div>
        </div>
      </section>

      <section className="steps">
        <div className="wrap">
          <div className="steps-head">
            <h2>Как начать смотреть вместе</h2>
            <p>От пустой страницы до общего просмотра — три шага, без регистрации гостей.</p>
          </div>
          <div className="steps-list">
            {steps.map((step) => (
              <div className="step" key={step.num}>
                <div className="step-num">{step.num}</div>
                <h3>{step.title}</h3>
                <p>{step.text}</p>
              </div>
            ))}
          </div>
        </div>
      </section>

      <section className="features">
        <div className="wrap">
          <div className="features-grid">
            {features.map((feature) => (
              <div
                className={feature.large ? "feature feature-lg" : "feature"}
                key={feature.title}
              >
                <div>
                  <div className="feature-tag">{feature.tag}</div>
                  <h3>{feature.title}</h3>
                  <p>{feature.text}</p>
                </div>
              </div>
            ))}
          </div>
        </div>
      </section>

      <SiteFooter />
    </>
  );
}
