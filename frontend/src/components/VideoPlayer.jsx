export default function VideoPlayer({ title, author, participants, createdAt }) {
  return (
    <section className="player-col">
      <div className="player-frame">
        <div className="play-glyph"></div>
      </div>

      <div className="player-controls">
        <div className="scrub">
          <div className="scrub-fill"></div>
        </div>
        <div className="controls-row">
          <div className="controls-left">
            <span>42:17 / 1:51:04</span>
            <span>{author} управляет плеером</span>
          </div>
        </div>
      </div>

      <div className="player-meta">
        <h2>{title}</h2>
        <p>
          Загружено {author} · комната создана {createdAt} · {participants} участника
        </p>
      </div>
    </section>
  );
}
