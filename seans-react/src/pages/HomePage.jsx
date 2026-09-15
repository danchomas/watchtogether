import { Link } from "react-router-dom";
import SiteHeader from "../components/SiteHeader.jsx";
import SiteFooter from "../components/SiteFooter.jsx";
import AvatarStack from "../components/AvatarStack.jsx";

// Контент шагов и возможностей вынесен в данные — так разметка
// в JSX не разрастается и её проще редактировать или расширять.
const steps = [
  {
    num: "01",
    title: "Создайте комнату",
    text: "Вставьте ссылку на видео или загрузите файл — Watchroom подготовит плеер и сгенерирует код комнаты.",
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
    title: "Плеер, синхронный с точностью до кадра",
    text: "Если кто-то ставит видео на паузу - вся комната ждёт вместе с ним. Автоматическая пересинхронизация решает задержки сети сама.",
    large: true,
  },
  {
    tag: "общение",
    title: "Чат рядом с видео",
    text: "Реакции и реплики не перекрывают экран и не теряются в перемотке.",
  },
  {
    tag: "источники",
    title: "Ссылка, файл или трансляция",
    text: "Работает с прямыми ссылками на видео, загруженными файлами и потоковыми источниками.",
  },
];

export default function HomePage() {
  return (
    <>
      <SiteHeader />

      <section className="hero">
        <div className="wrap">
          <div>
            <div className="timecode">
              <span className="dot"></span>
              00:42:17 — синхронно у всех участников
            </div>
            <h1>
              Разные устройства,
              <br />
              Одно время.
            </h1>
            <p className="hero-lede">
              Watchroom синхронизирует воспроизведение видео между всеми, кто в
              комнате — пауза, перемотка и звук совпадают у каждого зрителя с
              точностью до кадра. Рядом — чат, чтобы обсуждать происходящее в
              реальном времени.
            </p>
            <div className="hero-actions">
              <Link to="/auth" className="btn btn-primary">Создать комнату</Link>
              <Link to="/watch" className="btn btn-ghost">Посмотреть демо</Link>
            </div>
            <p className="hero-note">
              Комната открывается по короткому коду, например{" "}
              <code>AX7F92</code>
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
