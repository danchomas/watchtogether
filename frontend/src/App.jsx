import { Routes, Route } from "react-router-dom";
import HomePage from "./pages/HomePage.jsx";
import WatchPage from "./pages/WatchPage.jsx";
import AuthPage from "./pages/AuthPage.jsx";

export default function App() {
  return (
    <Routes>
      <Route path="/" element={<HomePage />} />
      <Route path="/watch" element={<WatchPage />} />
      <Route path="/auth" element={<AuthPage />} />
    </Routes>
  );
}
