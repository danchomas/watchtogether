import { Link } from "react-router-dom";
import AuthForm from "../components/AuthForm.jsx";

export default function AuthPage() {
  return (
    <div className="auth-shell">
      <section className="auth-side">
        <div className="auth-side-top">
          <Link to="/" className="logo">
            <span className="logo-mark"></span>
            Watchroom
          </Link>
        </div>
        <div className="auth-quote">
          <p>Мы смотрели финал в трёх разных городах — и всё равно вскрикнули в один момент</p>
        </div>
      </section>

      <section className="auth-form-side">
        <AuthForm />
      </section>
    </div>
  );
}
