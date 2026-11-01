import { Link } from "react-router-dom";

export default function SiteFooter() {
  return (
    <footer className="site-footer">
      <div className="wrap">
        <nav>
          <Link to="/">Главная</Link>
        </nav>
      </div>
    </footer>
  );
}
