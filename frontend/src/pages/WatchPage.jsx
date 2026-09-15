import RoomHeader from "../components/RoomHeader.jsx";
import VideoPlayer from "../components/VideoPlayer.jsx";
import ChatPanel from "../components/ChatPanel.jsx";

export default function WatchPage() {
  return (
    <>
      <RoomHeader title="Пятничный киновечер" code="AX7F92" />
      <main className="room-layout">
        <VideoPlayer
          title="Погоня за горизонтом"
          author="Аня"
          participants={3}
          createdAt="20:14"
        />
        <ChatPanel />
      </main>
    </>
  );
}
