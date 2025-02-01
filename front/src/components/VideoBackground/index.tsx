import { useRef, useState } from "react";

function VideoBackground() {
  const videoRef = useRef<HTMLVideoElement>(null);
  const [isMuted, setIsMuted] = useState(true);

  const toggleSound = () => {
    if (videoRef.current) {
      videoRef.current.muted = !videoRef.current.muted;
      setIsMuted(!isMuted);

      if (!videoRef.current.muted) {
        videoRef.current
          .play()
          .catch((e) => console.error("Ошибка воспроизведения:", e));
      }
    }
  };

  return (
    <div className="absolute top-0 left-0 w-full h-full">
      <video
        ref={videoRef}
        autoPlay
        loop
        muted
        className="absolute top-0 left-0 w-full h-full object-cover"
      >
        <source src="/background.mp4" type="video/mp4" />
      </video>

      <button
        onClick={toggleSound}
        className="absolute bottom-5 right-5 p-3 bg-black bg-opacity-50 opacity-80 text-white rounded-full"
      >
        {isMuted ? "🔇 Включить звук" : "🔊 Выключить звук"}
      </button>
    </div>
  );
}

export default VideoBackground;
