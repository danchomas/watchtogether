import { useState } from "react";

export default function AuthForm() {
  const [tab, setTab] = useState("login"); // "login" | "register"

  return (
    <div className="auth-box">
      <div className="auth-tabs">
        <button
          type="button"
          className={`auth-tab ${tab === "login" ? "active" : ""}`}
          onClick={() => setTab("login")}
        >
          Вход
        </button>
        <button
          type="button"
          className={`auth-tab ${tab === "register" ? "active" : ""}`}
          onClick={() => setTab("register")}
        >
          Регистрация
        </button>
      </div>

      {tab === "login" ? (
        <>
          <h1>С возвращением</h1>
          <p className="auth-sub">
            Войдите, чтобы открыть свои комнаты и историю просмотров.
          </p>

          <form onSubmit={(event) => event.preventDefault()}>
            <div className="field">
              <label htmlFor="email">Электронная почта</label>
              <input id="email" type="email" placeholder="you@example.com" />
            </div>
            <div className="field">
              <label htmlFor="password">Пароль</label>
              <input id="password" type="password" placeholder="Введите пароль" />
            </div>

            <div className="auth-row">
              <span>Запомнить меня</span>
              <a href="#">Забыли пароль?</a>
            </div>

            <button type="submit" className="btn btn-primary btn-block">
              Войти
            </button>
          </form>

          <p className="auth-switch">
            Ещё нет аккаунта?{" "}
            <a href="#" onClick={(event) => { event.preventDefault(); setTab("register"); }}>
              Зарегистрироваться
            </a>
          </p>
        </>
      ) : (
        <>
          <h1>Присоединяйтесь</h1>
          <p className="auth-sub">
            Создайте аккаунт, чтобы открывать комнаты под своим именем.
          </p>

          <form onSubmit={(event) => event.preventDefault()}>
            <div className="field">
              <label htmlFor="name">Имя</label>
              <input id="name" type="text" placeholder="Как вас называть в чате" />
            </div>
            <div className="field">
              <label htmlFor="reg-email">Электронная почта</label>
              <input id="reg-email" type="email" placeholder="you@example.com" />
            </div>
            <div className="field">
              <label htmlFor="reg-password">Пароль</label>
              <input id="reg-password" type="password" placeholder="Не меньше 8 символов" />
            </div>

            <button type="submit" className="btn btn-primary btn-block">
              Создать аккаунт
            </button>
          </form>

          <p className="auth-switch">
            Уже есть аккаунт?{" "}
            <a href="#" onClick={(event) => { event.preventDefault(); setTab("login"); }}>
              Войти
            </a>
          </p>
        </>
      )}
    </div>
  );
}
