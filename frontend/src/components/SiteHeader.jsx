import { Link } from "react-router-dom";

export default function SiteHeader() {
  return (
    <header className="site-header">
      <div className="wrap">
        <Link to="/" className="logo">
          <span className="logo-mark"></span>
          Watchroom
        </Link>
        <div className="header-actions">
          <Link to="/auth" className="header-link">Войти</Link>
          <Link to="/auth" className="btn btn-primary">Начать сеанс</Link>
        </div>
      </div>
    </header>
  );
}
