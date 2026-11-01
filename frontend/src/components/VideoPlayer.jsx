import { useEffect, useRef, useState } from "react";

export default function VideoPlayer({ room, onPlayback }) {
  const video = useRef(null);
  const suppress = useRef(false);
  const sync = useRef(() => {});
  const [notice, setNotice] = useState("");
  const { playback } = room;

  useEffect(() => {
    const player = video.current;
    if (!player) return;
    let release;
    let active = true;
    const synchronize = () => {
      if (player.readyState === 0) return;
      const elapsed = playback.playing ? Math.max(0, (Date.now() - Date.parse(playback.updatedAt)) / 1000) : 0;
      const target = Math.min(playback.position + elapsed, Number.isFinite(player.duration) ? player.duration : Infinity);
      suppress.current = true;
      if (Math.abs(player.currentTime - target) > 1.5) player.currentTime = target;
      if (playback.playing && player.paused) {
        player.play().then(() => { if (active) setNotice(""); }).catch(() => { if (active) setNotice("Нажмите на видео, чтобы разрешить воспроизведение в браузере."); });
      } else if (!playback.playing && !player.paused) player.pause();
      window.clearTimeout(release);
      release = window.setTimeout(() => { suppress.current = false; }, 150);
    };
    sync.current = synchronize;
    synchronize();
    const timer = window.setInterval(synchronize, 5000);
    return () => { active = false; window.clearInterval(timer); window.clearTimeout(release); suppress.current = false; };
  }, [playback.position, playback.playing, playback.updatedAt]);

  function change() {
    if (!suppress.current && video.current) onPlayback(video.current.currentTime, !video.current.paused);
  }

  return <section className="player-col">
    <video ref={video} className="video-element" src={room.videoUrl} controls playsInline preload="metadata"
      onPlay={change} onPause={change} onSeeked={change}
      onLoadedMetadata={() => sync.current()}
      onError={() => setNotice("Не удалось открыть видео. Проверьте прямую ссылку и доступность файла в браузере.")} />
    {notice && <p className="form-error" role="alert">{notice}</p>}
    <div className="player-meta">
      <h2>{room.title}</h2>
      <p>Участников: {room.participants} · Управлять воспроизведением может каждый участник</p>
    </div>
  </section>;
}
