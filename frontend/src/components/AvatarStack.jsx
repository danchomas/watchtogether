export default function AvatarStack({ colors = ["a1", "a2", "a3"] }) {
  return (
    <div className="avatar-stack">
      {colors.map((colorClass, index) => (
        <span key={index} className={colorClass}></span>
      ))}
    </div>
  );
}
