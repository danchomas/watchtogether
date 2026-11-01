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
          <a href="/#start" className="btn btn-primary">Начать сеанс</a>
        </div>
      </div>
    </header>
  );
}
